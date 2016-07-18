package sdk_actions

import (
    "github.com/gudtech/scamp-go/scamp"
    "encoding/json"
)

type CatalogGetConfigV1Input struct {
    Action  string   `json:"action"`
	Data    struct{} `json:"data"`
	Headers struct {
		ClientID int    `json:"client_id"`
		Ticket   string `json:"ticket"`
	} `json:"headers"`
	Version int `json:"version"`
}

//output here is what we send to the webhook
type CatalogGetConfigV1Output struct {
    // TODO ask Daniel if we need to better define the SDK response
    SkuFanout string `json:"sku_fanout"`
}

func CatalogGetConfigV1(msg *scamp.Message, client *scamp.Client) {
    scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
    var input CatalogGetConfigV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %s ", input)
    }

    //TODO: need to munge actual input data to output format for sdk
    // for the time being return an empty output struct
    var output CatalogGetConfigV1Output

    respMsg := scamp.NewResponseMessage()
    respMsg.WriteJson(output)
    respMsg.SetRequestId(msg.RequestId)

    _,err = client.Send(respMsg)
    if err != nil {
      return
    }
}
