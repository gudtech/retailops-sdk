package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
)

type OrderPullV1Input struct {
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
		ClientID    int `json:"client_id"`
		MaxPageSize int `json:"max_page_size"`
		Order       struct {
			ChannelRefnum int `json:"channel_refnum"`
		} `json:"order"`
		PageState interface{} `json:"page_state"`
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
        scamp.Info.Printf("Input Data Error: %s ", input)
    }

    //TODO: need to munge actual input data to output format for sdk
    var output OrderPullV1Output

    respMsg := scamp.NewResponseMessage()
    respMsg.WriteJson(output)
    respMsg.SetRequestId(msg.RequestId)

    _,err = client.Send(respMsg)
    if err != nil {
      return
    }
}
