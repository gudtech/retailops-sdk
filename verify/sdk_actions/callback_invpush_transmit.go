package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
  "bytes"
  "net/http"
  "time"
)
// TODO Update struct to match what the caller is sending
type InventoryPushV1Input struct {
  // Action string `json:"action"`
  // Data   struct {
    Channel struct {
      ID     int `json:"id"`
      Params struct {
        AppKey             string `json:"appKey"`
        BreakdownInventory int    `json:"breakdown_inventory"`
        Tenant             string `json:"tenant"`
      } `json:"params"`
      Definition struct {
        	Handle string `json:"handle"`
        	Params struct {
        			Interactions []ChannelInteraction `json:"interactions"`
        	} `json:"params"`
       } `json:"definition"`
    } `json:"channel"`
    ClientID  int `json:"client_id,string"`
    Inventory struct {
      Data []InventoryPushInputDataItem `json:"data"`
    } `json:"inventory"`
  // } `json:"data"`
  // Headers struct {
  //   ClientID int    `json:"client_id"`
  //   Ticket   string `json:"ticket"`
  // } `json:"headers"`
  // Version int `json:"version"`
}

type InventoryPushInputDataItem struct {
  Concept      string `json:"concept"`
  ID           string `json:"id"`
  QtyAvailable int    `json:"qty_available"`
  QtyBreakdown []ROPInputQtyBreakdownItem `json:"qty_breakdown"`
  Sku          string `json:"sku"`
}

//output structs
type InventoryPushV1Output struct {
  Action      string `json:"action"`
  ChannelInfo ChannelInfo `json:"channel_info"`
  ClientID             int    `json:"client_id"`
  IntegrationAuthToken string `json:"integration_auth_token"`
  InventoryUpdates     []InventoryUpdate `json:"inventory_updates"`
  Version int64 `json:"version"`
}

func InventoryPushV1(msg *scamp.Message, client *scamp.Client) {
    var input InventoryPushV1Input
    // scamp.Info.Printf("json: %s", string(msg.Bytes()))

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %+v\n ", err)
        respMsg := scamp.NewResponseMessage()
        respMsg.SetError(err.Error())
        respMsg.Write([]byte(err.Error()))
        respMsg.SetRequestId(msg.RequestId)
        _,err := client.Send(respMsg)
        if err != nil {
            scamp.Info.Printf("SDK: callback_invpush_transmit error %s", err)
        }
        return
    } else {
        //TODO: need to munge actual input data to output format for sdk
        // scamp.Info.Printf("input: %v", &input)
        var output InventoryPushV1Output
        output.Action = msg.Action //input.Action // this may need to be munged
        output.ChannelInfo.ID = input.Channel.ID
        output.ClientID = input.ClientID
        output.IntegrationAuthToken = msg.Ticket
        output.Version = msg.Version


        //ouput.InventoryUpdates
        inventoryArray := make([]InventoryUpdate, len(input.Inventory.Data), (cap(input.Inventory.Data)+1)*2)
        sumQuantity := 0
        for i := range input.Inventory.Data {
            var tempInvUpdate InventoryUpdate

            tempInvUpdate.Sku = input.Inventory.Data[i].Sku
            quantityDetailArray := make([]QuantityDetail, len(input.Inventory.Data[i].QtyBreakdown), (cap(input.Inventory.Data[i].QtyBreakdown)+1)*2)
            for j := range input.Inventory.Data[i].QtyBreakdown {
                var tempQuantityDetail QuantityDetail
                inputBreakdownItem := input.Inventory.Data[i].QtyBreakdown[j]
                // rannge over breakdown items, update sumQuantity
                tempQuantityDetail.AvailableQuantity = inputBreakdownItem.Sellable
                tempQuantityDetail.FacilityName = inputBreakdownItem.Facility
                tempQuantityDetail.TotalQuantity = (inputBreakdownItem.Sellable + inputBreakdownItem.Unclaimed) //correct?
                tempQuantityDetail.VendorName = inputBreakdownItem.Vendor
                sumQuantity += inputBreakdownItem.Sellable
                quantityDetailArray[j] = tempQuantityDetail
            }
            //assign sumQuantity after building details
            tempInvUpdate.QuantityAvailable = sumQuantity//I think this needs to be sum of all detail items
            inventoryArray[i] = tempInvUpdate
        }
        output.InventoryUpdates = inventoryArray

        // TODO: convert all actions to use code below, baseuri not valid, channel def passes endpointurl for each action
        // need to search channel.definition.params.Interactions for correct action and
        // retreive endpointurl
        var endPointURI string

        var version int
        interactions := input.Channel.Definition.Params.Interactions
        // scamp.Info.Printf("interactions: %v", interactions)
        for i := range interactions {
            if interactions[i].Action == "inventory_push" {
                // scamp.Info.Printf("action: %s\n", interactions[i].Action)
                endPointURI = interactions[i].EndpointURL
                version = interactions[i].Version
            }
        }
        // scamp.Info.Printf("endPointURI: %s\n", endPointURI)
        if len(endPointURI) == 0 || version <= 0 {
            scamp.Info.Printf("endpoint or version is blank")
            return
        }
        channelURI := BuildURI(endPointURI, version )
        // scamp.Info.Printf("channelURI: %s\n", channelURI)

        var requestBuffer bytes.Buffer
        err := json.NewEncoder(&requestBuffer).Encode(output)
        if err != nil {
            scamp.Info.Printf("err: %s\n", err)
            return
        }

        var httpClient = &http.Client{
            Timeout: time.Minute * 5,
        }

        scamp.Info.Printf("Making API call to: %s", channelURI)
        response,err := httpClient.Post(channelURI, "application/json", &requestBuffer)
        defer response.Body.Close()
        if err != nil {
            scamp.Info.Printf("err: %s\n", err)
            return
        }

        var apiResp CommonV1Response
        err = json.NewDecoder(response.Body).Decode(&apiResp)
        if err != nil {
            return
        }
        // ./src/github.com/gudtech/retailops-sdk/verify/schema/
        // validResponse, err := ValidateResponse("../verify/schema/inventory_push_v1.json", &apiResp )
        validResponse, err := ValidateResponse("./src/github.com/gudtech/retailops-sdk/verify/schema/inventory_push_v1.json", &apiResp )
        if err != nil {
            scamp.Info.Printf("There was an error validating the response: %+v\n ", err)
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(err.Error())
            respMsg.Write([]byte(err.Error()))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: callback_invpush_transmit error %s", err)
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
                scamp.Info.Printf("SDK: callback_invpush_transmit error %s", err)
            }
            return
        }
        // munge API response back into what perl modules expect and return JSON below.

        var invpushResponse InvPushTransmitResponse
        eventArray := make([]InvPushTransmitResponseEvent, len(apiResp.Events), (cap(apiResp.Events)+1)*2)
        for i := range apiResp.Events {
            var tempEvent InvPushTransmitResponseEvent
            tempEvent.Data.Message = apiResp.Events[i].Message
            tempEvent.Data.Status = apiResp.Events[i].Code
            tempEvent.Handle = apiResp.Events[i].EventType

            secondaryArray := make([]InvPushTransmitResponseSecondary, len(apiResp.Events[i].Associations), (cap(apiResp.Events[i].Associations)+1)*2)
            for j := range apiResp.Events[i].Associations {
                var tempSecondary InvPushTransmitResponseSecondary
                tempSecondary.Concept = apiResp.Events[i].Associations[j].IdentifierType
                tempSecondary.ID = apiResp.Events[i].Associations[j].Identifier
                secondaryArray[j] = tempSecondary
            }
            tempEvent.Secondary = secondaryArray
            eventArray[i] = tempEvent
        }
        invpushResponse.Events = eventArray

        respMsg := scamp.NewResponseMessage()
        respMsg.WriteJson(invpushResponse)
        respMsg.SetRequestId(msg.RequestId)

        _,err = client.Send(respMsg)
        if err != nil {
            return
        }
    }
}
