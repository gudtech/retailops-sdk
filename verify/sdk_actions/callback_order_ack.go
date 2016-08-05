package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
  "bytes"
  "net/http"
  "time"
)

type OrderAcknowledgeV1Input struct {
    Action string `json:"action"`
	Data   struct {
		Channel struct {
			ID     int `json:"id"`
			Params struct {
				StoreID                    string `json:"StoreID"`
				NextOrderRefnum            int    `json:"next_order_refnum"`
				OrderAckStatusID           string `json:"order_ack_status_id"`
				OrderFulfilledStatusID     string `json:"order_fulfilled_status_id"`
				OrderInFilfillmentStatusID string `json:"order_in_filfillment_status_id"`
                BaseURI                    string `json:"base_uri"` //need to add to perl JSON
			} `json:"params"`
            Definition struct {
              	Handle string `json:"handle"`
              	Params struct {
              			Interactions []ChannelInteraction `json:"interactions"`
              	} `json:"params"`
             } `json:"definition"`
		} `json:"channel"`
		ClientID int `json:"client_id"`
		Order    struct {
			Acks []string `json:"acks"` //array of channel_ref_nums
		} `json:"order"`
	} `json:"data"`
	Headers struct {
		ClientID int    `json:"client_id"`
		Ticket   string `json:"ticket"`
	} `json:"headers"`
	Version int `json:"version"`
}

type OrderAcknowledgeV1Output struct {
    Action      string `json:"action"`
	ChannelInfo struct {
		ID int `json:"id"`
	} `json:"channel_info"`
	ClientID             int    `json:"client_id"`
	IntegrationAuthToken string `json:"integration_auth_token"`
	Orders               []OrderAcknowledgement `json:"orders"`
	Version int `json:"version"`
}

type OrderAcknowledgement struct {
    ChannelOrderRefnum string `json:"channel_order_refnum"`
    RetailopsOrderID   int    `json:"retailops_order_id"`
}

func OrderAcknowledgeV1(msg *scamp.Message, client *scamp.Client) {
    scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
    var input OrderAcknowledgeV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %+v\n ", err)
        respMsg := scamp.NewResponseMessage()
        respMsg.SetError(err.Error())
        respMsg.Write([]byte(err.Error()))
        respMsg.SetRequestId(msg.RequestId)
        _,err := client.Send(respMsg)
        if err != nil {
            scamp.Info.Printf("SDK: callback_order_update error %s", err)
        }
        return
    } else {
        var output OrderAcknowledgeV1Output
        output.Action = input.Action
        output.Version = input.Version
        output.IntegrationAuthToken = input.Headers.Ticket
        output.ChannelInfo.ID = input.Data.Channel.ID
        output.ClientID = input.Headers.ClientID

        orderArray := make([]OrderAcknowledgement, len(input.Data.Order.Acks), (cap(input.Data.Order.Acks)+1)*2)
        for i := range input.Data.Order.Acks {
            var tempAck OrderAcknowledgement
            tempAck.RetailopsOrderID = 0 //we aren't currently passing rrtailops ordder id from perl
            tempAck.ChannelOrderRefnum = input.Data.Order.Acks[i]
            orderArray[i] = tempAck
        }

        // TODO: convert all actions to use code below, baseuri not valid, channel def passes endpointurl for each action
        // need to search channel.definition.params.Interactions for correct action and
        // retreive endpointurl
        var endPointURI string
        var version int
        interactions := input.Data.Channel.Definition.Params.Interactions
        for i := range interactions {
            if interactions[i].Action == "order_acknowledge" {
                endPointURI = interactions[i].EndpointURL
                version = interactions[i].Version
            }
        }

        if len(endPointURI) == 0 || version <= 0 {
            return
        }
        channelURI := BuildURI(endPointURI, version )
        // baseURI := input.Data.Channel.Params.BaseURI
        // if len(baseURI) == 0 {
        //     return
        // }
        // channelURI := BuildURI(baseURI, "order_acknowledge_v1")

        var requestBuffer bytes.Buffer
        err := json.NewEncoder(&requestBuffer).Encode(output)
        if err != nil {
            return
        }

        var httpClient = &http.Client{
            Timeout: time.Minute * 5,
        }

        scamp.Info.Printf("Making API call to: %s", channelURI)
        response,err := httpClient.Post(channelURI, "application/json", &requestBuffer)
        defer response.Body.Close()
        if err != nil {
            return
        }

        var apiResp CommonV1Response //TODO: check response type in swagger
        err = json.NewDecoder(response.Body).Decode(&apiResp)
        if err != nil {
            return
        }

        validResponse, err := ValidateResponse("../verify/schema/order_acknowledge_v1.json", &apiResp )
        if err != nil {
            scamp.Info.Printf("There was an error validating the response: %+v\n ", err)
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(err.Error())
            respMsg.Write([]byte(err.Error()))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: callback_order_ack error %s", err)
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
                scamp.Info.Printf("SDK: callback_order_ack error %s", err)
            }
            return
        }

        var orderAckResponse CommonResponseV1

        eventArray := make([]CommonResponseEvent, len(apiResp.Events), (cap(apiResp.Events)+1)*2)
        for i := range apiResp.Events {
            var tempEvent CommonResponseEvent
            tempEvent.Data.Message = apiResp.Events[i].Message
            tempEvent.Data.Status = apiResp.Events[i].Code
            tempEvent.Handle = apiResp.Events[i].EventType

            secondaryArray := make([]CommonResponseSecondary, len(apiResp.Events[i].Associations), (cap(apiResp.Events[i].Associations)+1)*2)
            for j := range apiResp.Events[i].Associations {
                var tempSecondary CommonResponseSecondary
                tempSecondary.Concept = apiResp.Events[i].Associations[j].IdentifierType
                tempSecondary.ID = apiResp.Events[i].Associations[j].Identifier
                secondaryArray[j] = tempSecondary
            }
            tempEvent.Secondary = secondaryArray
            eventArray[i] = tempEvent
        }
        orderAckResponse.Events = eventArray

        respMsg := scamp.NewResponseMessage()
        respMsg.WriteJson(orderAckResponse)
        respMsg.SetRequestId(msg.RequestId)

        _,err = client.Send(respMsg)
        if err != nil {
          return
        }
    }
}
