package sdk_actions

//common input structs
type ROPInputShipment struct {
    ID       int `json:"id,string"` //was string
    Packages []ROPInputPackage `json:"packages"`
}

type ROPInputPackage struct {
    CarrierClassCode string      `json:"carrier_class_code"`
    CarrierClassName string      `json:"carrier_class_name"`
    CarrierCode      string      `json:"carrier_code"`
    CarrierName      string      `json:"carrier_name"`
    ClassName        string      `json:"class_name"`
    DateShipped      string      `json:"date_shipped"`
    ID               int         `json:"id,string"`
    MappedShipcode   string `json:"mapped_shipcode"`
    ShipItems        []ROPInputShipItem `json:"ship_items"`
    TrackingNumber string `json:"tracking_number"`
    Weight         float64      `json:"weight,string"`
}

type ROPInputShipItem struct {
    ChannelOrderRefnum int `json:"channel_order_refnum,string"`
    ChannelRefnum      string `json:"channel_refnum"`
    ID                 int `json:"id,string"`
    ItemGiftwrapAmt    int    `json:"item_giftwrap_amt"`
    ItemGiftwrapTaxAmt int    `json:"item_giftwrap_tax_amt"`
    ItemProductAmt     int    `json:"item_product_amt,string"`
    ItemRecyclingAmt   int    `json:"item_recycling_amt"`
    ItemShippingAmt    float32 `json:"item_shipping_amt,string"`
    ItemShippingTaxAmt int    `json:"item_shipping_tax_amt"`
    ItemSubtotalAmt    int    `json:"item_subtotal_amt"`
    ItemTaxAmt         float32 `json:"item_tax_amt,string"`
    Quantity           int    `json:"quantity"`
    SkuNum             string `json:"sku_num"`
}

type ROPInputUnshippedItem struct {
    ChannelItemRefnum  string `json:"channel_item_refnum"`
    ChannelRefnum      string `json:"channel_refnum"`
    ID                 int `json:"id,string"`
    ItemGiftwrapAmt    int    `json:"item_giftwrap_amt"`
    ItemGiftwrapTaxAmt int    `json:"item_giftwrap_tax_amt"`
    ItemProductAmt     int    `json:"item_product_amt,string"`
    ItemRecyclingAmt   int    `json:"item_recycling_amt"`
    ItemShippingAmt    float32 `json:"item_shipping_amt,string"`
    ItemShippingTaxAmt int    `json:"item_shipping_tax_amt"`
    ItemSubtotalAmt    int    `json:"item_subtotal_amt"`
    ItemTaxAmt         float32 `json:"item_tax_amt,string"`
    Quantity           int    `json:"quantity"`
    SkuNum             string `json:"sku"`
    ChannelID          string `json:"channel_id"`
    Reason             string `json:"reason"`
}

type ROPInputQtyBreakdownItem struct {
    EstShip         string   `json:"est_ship"`
	Facility        string   `json:"facility"`
	ReservingOrders []string `json:"reserving_orders"`
	Sellable        int      `json:"sellable"`
	Sku             string   `json:"sku"`
	Unclaimed       int      `json:"unclaimed"`
	Vendor          string   `json:"vendor"`
    Zones     []ROPInputZone `json:"zones"`
}

type ROPInputZone struct {
  Npick int    `json:"npick"`
  Pick  int    `json:"pick"`
  Zone  string `json:"zone"`
}

//common output structs
type Shipment struct {
    Packages []Package `json:"packages"`
    RetailopsShipmentID int `json:"retailops_shipment_id"`
}

type Package struct {
    CarrierClassName string `json:"carrier_class_name"`
    CarrierName      string `json:"carrier_name"`
    ChannelShipCode  string `json:"channel_ship_code"`
    DateShipped      string `json:"date_shipped"`
    PackageItems     []PackageItem `json:"package_items"`
    RetailopsPackageID int    `json:"retailops_package_id"`
    TrackingNumber     string `json:"tracking_number"`
    WeightKg           float64  `json:"weight_kg"`
}

type PackageItem struct {
    ChannelItemRefnum       string `json:"channel_item_refnum"`
    Quantity                int    `json:"quantity"`
    RetailopsOrderItemID    int    `json:"retailops_order_item_id"`
    RetailopsShipmentItemID int    `json:"retailops_shipment_item_id"`
    Sku                     string `json:"sku"`
}

type ReturnItem struct {
    ChannelItemRefnum  string `json:"channel_item_refnum"`
    CreditAmt          float32    `json:"credit_amt"`
    GiftwrapAmt        float32    `json:"giftwrap_amt"`
    GiftwrapTaxAmt     float32    `json:"giftwrap_tax_amt"`
    ItemShippingTaxAmt float32    `json:"item_shipping_tax_amt"`
    ProductAmt         int    `json:"product_amt"`
    Quantity           int    `json:"quantity"`
    Reason             string `json:"reason"`
    RecyclingAmt       int    `json:"recycling_amt"`
    RestockFeeAmt      int    `json:"restock_fee_amt"`
    RetailopsItemID    int    `json:"retailops_item_id"`
    ShippingAmt        int    `json:"shipping_amt"`
    Sku                int    `json:"sku"`
    SubtotalAmt        float32    `json:"subtotal_amt"`
    TaxAmt             float32    `json:"tax_amt"`
}

type UnshippedItem struct {
    ChannelItemRefnum      string `json:"channel_item_refnum"`
    EffectiveExtendedPrice float32    `json:"effective_extended_price"`
    EffectiveUnitPrice     float32    `json:"effective_unit_price"`
    OrderedQuantity        int    `json:"ordered_quantity"`
    Sku                    string `json:"sku"`
    UnshippedQuantity      int    `json:"unshipped_quantity"`
}

type RMA struct {
    DiscountAmt int `json:"discount_amt"`
    Items       []ReturnItem `json:"items"` //ReturnItem struct defined in callback_items_returned.go
    ProductAmt        int    `json:"product_amt"`
    RefundAction      string `json:"refund_action"`
    RefundAmt         int    `json:"refund_amt"`
    RetailopsReturnID int    `json:"retailops_return_id"`
    RetailopsRmaID    string `json:"retailops_rma_id"`
    ShippingAmt       int    `json:"shipping_amt"`
    SubtotalAmt       int    `json:"subtotal_amt"`
    TaxAmt            int    `json:"tax_amt"`
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
  ID int `json:"id, string"`
}
