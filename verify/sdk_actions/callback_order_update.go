package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
  "bytes"
  "net/http"
  "time"
)

func OrderUpdateV1(msg *scamp.Message, client *scamp.Client) {
    // var err error
    var input OrderUpdateV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %+v\n ", err)
        respMsg := scamp.NewResponseMessage()
        respMsg.SetError(err.Error())
        respMsg.Write([]byte(err.Error()))
        respMsg.SetRequestId(msg.RequestId)
        _,err := client.Send(respMsg)
        if err != nil {
          return
        }
    } else {
        var output OrderUpdateV1Output
        output.Action = input.Action
        output.ChannelInfo.ID = input.Data.Channel.ID
        output.Order.ChannelOrderRefnum = input.Data.Order.ChannelRefnum
        output.Order.GrandTotal = input.Data.Order.GrandTotal
        output.Order.RetailopsOrderID = input.Data.Order.ID
        output.Version = input.Version
        output.IntegrationAuthToken = input.Headers.Ticket
        output.Order.Shipments = BuildShipments(input.Data.Order.Shipments)
        output.Order.UnshippedItems = BuildUnshippedItems(input.Data.Order.UnshippedItemsRef)

        // baseURI := input.Data.Channel.Params.BaseURI
        // if len(baseURI) == 0 {
        //     return
        // }
        // channelURI := BuildURI(baseURI, "order_update_v1")
        var endPointURI string
        var version int
        interactions := input.Data.Channel.Definition.Params.Interactions
        for i := range interactions {
            if interactions[i].Action == "order_update" {
                endPointURI = interactions[i].EndpointURL
                version = interactions[i].Version
            }
        }

        if len(endPointURI) == 0 || version <= 0 {
            return
        }
        channelURI := BuildURI(endPointURI, version )

        var requestBuffer bytes.Buffer
        err := json.NewEncoder(&requestBuffer).Encode(output)
        if err != nil {
          return
        }

        var httpClient = &http.Client{
          Timeout: time.Minute * 5, //NOTE: this timeout seems kind of long...
        }

        scamp.Info.Printf("Making API call to: %s", channelURI)
        response,err := httpClient.Post(channelURI, "application/json", &requestBuffer)
        defer response.Body.Close()
        if err != nil {
          return
        }

        var apiResp OrderUpdateV1Response
        err = json.NewDecoder(response.Body).Decode(&apiResp)
        if err != nil {
          return
        }

        validResponse, err := ValidateResponse("../verify/schema/order_update_v1.json", &apiResp )//string(responseJSON)
        if err != nil {
            //TODO see what type of error I should return to ROP when service response fails to validate
            scamp.Info.Printf("There was an error validating the response: %+v\n ", err)
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(err.Error())
            respMsg.Write([]byte(err.Error()))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: callback_order_update error %s", err)
            }
            return
        }
        if !validResponse {
            validationMsg := "API response was invalid"
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(validationMsg)
            respMsg.Write([]byte(validationMsg))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: callback_order_update error %s", err)
            }
            return
        }

        // TODO: munge API response back into what perl modules expect and return JSON below.
        // need to determine what perl modules were doing with the previous response data
        // and if we need to modify schema to support it because the SDK defines a
        // different structure


        // encode and return JSON
        respMsg := scamp.NewResponseMessage()
        respMsg.WriteJson(apiResp)
        respMsg.SetRequestId(msg.RequestId)
        _,err = client.Send(respMsg)
        if err != nil {
            scamp.Info.Printf("SDK: callback_order_update error %s", err)
            return
        }
    }
}
