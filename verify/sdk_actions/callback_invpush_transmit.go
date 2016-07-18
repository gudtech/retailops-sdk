package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
)

type InventoryPushV1Input struct {
  Action string `json:"action"`
  Data   struct {
    Channel struct {
      ID     int `json:"id"`
      Params struct {
        AppKey             string `json:"appKey"`
        BreakdownInventory int    `json:"breakdown_inventory"`
        Tenant             string `json:"tenant"`
      } `json:"params"`
    } `json:"channel"`
    ClientID  int `json:"client_id"`
    Inventory struct {
      Data []struct {
        Concept      string `json:"concept"`
        ID           string `json:"id"`
        QtyAvailable int    `json:"qty_available"`
        QtyBreakdown []struct {
          Facility  string `json:"facility"`
          Sellable  int    `json:"sellable"`
          Sku       string `json:"sku"`
          Unclaimed int    `json:"unclaimed"`
          Zones     []struct {
            Npick int    `json:"npick"`
            Pick  int    `json:"pick"`
            Zone  string `json:"zone"`
          } `json:"zones"`
        } `json:"qty_breakdown"`
        Sku string `json:"sku"`
      } `json:"data"`
    } `json:"inventory"`
  } `json:"data"`
  Headers struct {
    ClientID int    `json:"client_id"`
    Ticket   string `json:"ticket"`
  } `json:"headers"`
  Version int `json:"version"`
}

//output structs
type InventoryPushV1Output struct {
  Action      string `json:"action"`
  ChannelInfo ChannelInfo `json:"channel_info"`
  ClientID             int    `json:"client_id"`
  IntegrationAuthToken string `json:"integration_auth_token"`
  InventoryUpdates     []InventoryUpdate `json:"inventory_updates"`
  Version int `json:"version"`
}

type InventoryUpdate struct {
  QuantityAvailable int `json:"quantity_available"`
  QuantityDetail    []QuantityDetail `json:"quantity_detail"`
  Sku string `json:"sku"`
}

type QuantityDetail struct {
  AvailableQuantity         int    `json:"available_quantity"`
  EstimatedAvailabilityDate string `json:"estimated_availability_date"`
  FacilityName              string `json:"facility_name"`
  Po                        string `json:"po"`
  PoDestination             string `json:"po_destination"`
  QuantityType              string `json:"quantity_type"`
  TotalQuantity             int    `json:"total_quantity"`
  VendorName                string `json:"vendor_name"`
}

type ChannelInfo struct {
  ID int `json:"id"`
}

//service method - receives JSON data from PERL caller, forwards to service endpoint
func InventoryPushV1(msg *scamp.Message, client *scamp.Client) {
  scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
  var input InventoryPushV1Input

  err := json.Unmarshal(msg.Bytes(), &input)
  if err != nil {
      scamp.Info.Printf("Input Data Error: %s ", input)
  }

  //TODO: need to munge actual input data to output format for sdk
  var output InventoryPushV1Output
  output.InventoryUpdates = make([]InventoryUpdate, 0)

  respMsg := scamp.NewResponseMessage()
  respMsg.WriteJson(output)
  respMsg.SetRequestId(msg.RequestId)

  _,err = client.Send(respMsg)
  if err != nil {
    return
  }
}
