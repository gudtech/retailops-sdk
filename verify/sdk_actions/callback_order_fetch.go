package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
  "bytes"
  "net/http"
  "time"
)

type OrderPullV1Input struct {
    Action string `json:"action"`
	Data   struct {
		Channel struct {
			ID     int `json:"id"`
			Params struct {
                BaseURI                    string `json:"base_uri"`//added was not in original perl request
				StoreID                    string `json:"StoreID"`
				NextOrderRefnum            int    `json:"next_order_refnum,"`
				OrderAckStatusID           string `json:"order_ack_status_id"`
				OrderFulfilledStatusID     string `json:"order_fulfilled_status_id"`
				OrderInFilfillmentStatusID string `json:"order_in_filfillment_status_id"`
			} `json:"params"`
            Definition struct {
              	Handle string `json:"handle"`
              	Params struct {
              			Interactions []ChannelInteraction `json:"interactions"`
              	} `json:"params"`
             } `json:"definition"`
		} `json:"channel"`
		ClientID    int `json:"client_id"`
		MaxPageSize int `json:"max_page_size"`
		Order       struct {
			ChannelRefnum int `json:"channel_refnum"`
		} `json:"order"`
		// PageState interface{} `json:"page_state"`
        PageState string `json:"page_state"`
		Single    int         `json:"single"`
	} `json:"data"`
	Headers struct {
		ClientID int    `json:"client_id"`
		Ticket   string `json:"ticket"`
	} `json:"headers"`
	Version int `json:"version"`
}

type OrderPullV1Output struct {
    Action      string `json:"action"`
	ChannelInfo struct {
		ID int `json:"id"`
	} `json:"channel_info"`
	ClientID             int    `json:"client_id"`
	IntegrationAuthToken string `json:"integration_auth_token"`
	PageToken            string `json:"page_token"`
	SpecificOrders       []SpecificOrder `json:"specific_orders"`
	Version int `json:"version"`
}

type SpecificOrder struct {
    ChannelRefnum string `json:"channel_refnum"`
}

func OrderPullV1(msg *scamp.Message, client *scamp.Client) {
    scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
    var input OrderPullV1Input

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
        var output OrderPullV1Output
        //TODO: need to munge actual input data to output format for sdk
        output.Action = input.Action
        output.ChannelInfo.ID = input.Data.Channel.ID
        output.ClientID = input.Headers.ClientID
        output.IntegrationAuthToken = input.Headers.Ticket
        output.Version = input.Version
        output.PageToken = input.Data.PageState

        // baseURI := input.Data.Channel.Params.BaseURI
        // if len(baseURI) == 0 {
        //     return //TODO: return relevant scamp error msg (for all callbacks)
        // }
        // channelURI := BuildURI(baseURI, "inventory_push_v1")

        var endPointURI string
        var version int
        interactions := input.Data.Channel.Definition.Params.Interactions
        for i := range interactions {
            if interactions[i].Action == "order_pull" {
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
            Timeout: time.Minute * 5,
        }

        scamp.Info.Printf("Making API call to: %s", channelURI)
        response,err := httpClient.Post(channelURI, "application/json", &requestBuffer)
        defer response.Body.Close()
        if err != nil {
            return
        }

        var apiResp CommonV1Response
        err = json.NewDecoder(response.Body).Decode(&apiResp)
        if err != nil {
            return
        }

        validResponse, err := ValidateResponse("../verify/schema/order_pull_v1.json", &apiResp )
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

        // TODO: munge API response back into what perl modules expect and return JSON below.

        respMsg := scamp.NewResponseMessage()
        respMsg.WriteJson(output)
        respMsg.SetRequestId(msg.RequestId)

        _,err = client.Send(respMsg)
        if err != nil {
          return
        }
    }
}
