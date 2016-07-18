package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
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
			} `json:"params"`
		} `json:"channel"`
		ClientID int `json:"client_id"`
		Order    struct {
			Acks []string `json:"acks"`
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
	Orders               []Order `json:"orders"`
	Version int `json:"version"`
}

type Order struct {
    ChannelOrderRefnum string `json:"channel_order_refnum"`
    RetailopsOrderID   int    `json:"retailops_order_id"`
}

func OrderAcknowledgeV1(msg *scamp.Message, client *scamp.Client) {
    scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
    var input OrderAcknowledgeV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %s ", input)
    }

    //TODO: need to munge actual input data to output format for sdk
    var output OrderAcknowledgeV1Output

    respMsg := scamp.NewResponseMessage()
    respMsg.WriteJson(output)
    respMsg.SetRequestId(msg.RequestId)

    _,err = client.Send(respMsg)
    if err != nil {
      return
    }
}
