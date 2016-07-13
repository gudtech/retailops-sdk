package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
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

func InvpushTransmitV1(msg *scamp.Message, client *scamp.Client) {

}
