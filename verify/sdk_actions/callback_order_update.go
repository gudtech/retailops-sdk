package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
)

type OrderUpdateV1Input struct {
    Action string `json:"action"`
	Data   struct {
		Channel struct {
			ID     int `json:"id,string"`
			Params struct {
				BaseURI                       string      `json:"base_uri"`
				EmailInvoice                  int         `json:"email_invoice"`
				EmailReturn                   int         `json:"email_return"`
				EmailTracking                 int         `json:"email_tracking"`
				ExpressConfigurableSuperLinks int         `json:"express_configurable_super_links"`
				ImportOrderAttrs              string      `json:"import_order_attrs"`
				InvSuspendedInstock           int         `json:"inv_suspended_instock"`
				InvSuspendedMode              interface{} `json:"inv_suspended_mode"`
				OrderWriteback                int         `json:"order_writeback"`
				PushCancel                    int         `json:"push_cancel"`
				UnsetOtherAttributes          int         `json:"unset_other_attributes"`
				UnsetOtherMedia               int         `json:"unset_other_media"`
			} `json:"params"`
		} `json:"channel"`
		LineItems []struct {
			ApportionedShipAmt    string      `json:"apportioned_ship_amt"`
			Corr                  string      `json:"corr"`
			DirectShipAmt         string      `json:"direct_ship_amt"`
			EstimatedCost         float32      `json:"estimated_cost,string"`
			EstimatedExtendedCost float32      `json:"estimated_extended_cost,string"`
			EstimatedShipDate     int         `json:"estimated_ship_date"`
			EstimatedUnitCost     float32      `json:"estimated_unit_cost,string"`
			Quantity              string      `json:"quantity"`
			Removed               interface{} `json:"removed"`
			Sku                   string      `json:"sku"`
			UnitPrice             float32      `json:"unit_price,string"`
		} `json:"line_items"`
		Order struct {
			ChannelPayment struct {
				Available           int    `json:"available"`
				Captured            int    `json:"captured"`
				CapturesAreDeferred int    `json:"captures_are_deferred"`
				CapturesAreExternal int    `json:"captures_are_external"`
				ChannelID           int    `json:"channel_id"`
				Charged             int    `json:"charged"`
				Credited            int    `json:"credited"`
				ID                  string `json:"id"`
				Method              string `json:"method"`
				Module              string `json:"module"`
				Settled             int    `json:"settled"`
				Unsettled           int    `json:"unsettled"`
				Voided              int    `json:"voided"`
			} `json:"channel_payment"`
			ChannelRefnum        string `json:"channel_refnum"`
			FromCounterpartyRate int    `json:"from_counterparty_rate"`
			GrandTotal           float32 `json:"grand_total,string"`
			ID                   int `json:"id,string"`
			PaymentSeriesID      string `json:"payment_series_id"`
			PaymentStatus        struct {
				Authed    int `json:"authed"`
				Available int `json:"available"`
				ByAccount []struct {
					Available           int    `json:"available"`
					Captured            int    `json:"captured"`
					CapturesAreDeferred int    `json:"captures_are_deferred"`
					CapturesAreExternal int    `json:"captures_are_external"`
					ChannelID           int    `json:"channel_id"`
					Charged             int    `json:"charged"`
					Credited            int    `json:"credited"`
					ID                  string `json:"id"`
					Method              string `json:"method"`
					Module              string `json:"module"`
					Settled             int    `json:"settled"`
					Unsettled           int    `json:"unsettled"`
					Voided              int    `json:"voided"`
				} `json:"by_account"`
				Captured          int `json:"captured"`
				Charged           int `json:"charged"`
				Credited          int `json:"credited"`
				Settled           int `json:"settled"`
				Success           int `json:"success"`
				Unsettled         int `json:"unsettled"`
				UnsettledDeferred int `json:"unsettled_deferred"`
				UnsettledExternal int `json:"unsettled_external"`
			} `json:"payment_status"`
			ShipServiceName   string        `json:"ship_service_name"`
			Shipments         []OrderUpdateShipment `json:"shipments"`
			UnshippedItemsRef []OrderUpdateUnshippedItem `json:"unshipped_items_ref"`
		} `json:"order"`
		OrderInfo struct {
			DirectTaxAmt string `json:"direct_tax_amt"`
			DiscountAmt  string `json:"discount_amt"`
			ShippingAmt  string `json:"shipping_amt"`
			TaxAmt       string `json:"tax_amt"`
		} `json:"order_info"`
		Rmas []interface{} `json:"rmas"` //TODO: need structure of RMA?
	} `json:"data"`
	Headers struct {
		ClientID int    `json:"client_id"`
		Ticket   string `json:"ticket"`
	} `json:"headers"`
	Version int `json:"version"`
}

type OrderUpdateShipment struct {
    ID       int `json:"id,string"` //was string
    Packages []OrderUpdatePackage `json:"packages"`
}

type OrderUpdatePackage struct {
    CarrierClassCode string      `json:"carrier_class_code"`
    CarrierClassName string      `json:"carrier_class_name"`
    CarrierCode      string      `json:"carrier_code"`
    CarrierName      string      `json:"carrier_name"`
    ClassName        string      `json:"class_name"`
    DateShipped      string      `json:"date_shipped"`
    ID               int         `json:"id,string"`
    MappedShipcode   string `json:"mapped_shipcode"`
    ShipItems        []OrderUpdateShipItem `json:"ship_items"`
    TrackingNumber string `json:"tracking_number"`
    Weight         float64      `json:"weight,string"`
}

type OrderUpdateShipItem struct {
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

type OrderUpdateUnshippedItem struct {
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

//output structs
type OrderUpdateV1Output struct {
    Action      string `json:"action"`
	ChannelInfo struct {
		ID int `json:"id"`
	} `json:"channel_info"`
	ClientID             int    `json:"client_id"`
	IntegrationAuthToken string `json:"integration_auth_token"`
	Order                struct {
		ChannelOrderRefnum string `json:"channel_order_refnum"`
		GrandTotal         float32    `json:"grand_total"`//Note: should we use float64 or float32
		RetailopsOrderID   int    `json:"retailops_order_id"`
		Shipments          []Shipment `json:"shipments"` //Shipment struct defined in callback_items_returned.go
		UnshippedItems []UnshippedItem `json:"unshipped_items"` //UnshippedItem struct defined in callback_items_returned.go
	} `json:"order"`
	Rmas []RMA `json:"rmas"`
	Version int `json:"version"`
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

func OrderUpdateV1(msg *scamp.Message, client *scamp.Client) {
    var err error
    // scamp.Info.Printf("incoming msg: %s", string(msg.Bytes()))
    var input OrderUpdateV1Input

    err = json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %+v\n ", err)
        respMsg := scamp.NewResponseMessage()
        respMsg.SetError(err.Error())
        respMsg.Write([]byte(err.Error()))
        respMsg.SetRequestId(msg.RequestId)
        // scamp.Info.Printf("respMsg: %s", respMsg.Error)
        _,err = client.Send(respMsg)
        if err != nil {
          return
        }
    } else {
        var output OrderUpdateV1Output
        // scamp.Info.Printf("Data: %v", input)
        output.Action = input.Action
        output.ChannelInfo.ID = input.Data.Channel.ID
        output.Order.ChannelOrderRefnum = input.Data.Order.ChannelRefnum
        output.Order.GrandTotal = input.Data.Order.GrandTotal
        output.Order.RetailopsOrderID = input.Data.Order.ID
        output.Version = input.Version
        output.IntegrationAuthToken = input.Headers.Ticket

        shipmentArray := make([]Shipment, len(input.Data.Order.Shipments), (cap(input.Data.Order.Shipments)+1)*2) //NOTE: do we really need to double?
        for i := range input.Data.Order.Shipments {
                //build new shipment item
                var tempShipment Shipment
                tempShipment.RetailopsShipmentID = input.Data.Order.Shipments[i].ID
                //create package array for shipment object
                tempShipment.Packages = make([]Package, len(input.Data.Order.Shipments[i].Packages), (cap(input.Data.Order.Shipments[i].Packages)+1)*2)
                //range over package array and add packages to package array
                for j := range input.Data.Order.Shipments[i].Packages {
                    var tempPackage Package
                    tempPackage.CarrierClassName = input.Data.Order.Shipments[i].Packages[j].CarrierClassName
                    tempPackage.CarrierName = input.Data.Order.Shipments[i].Packages[j].CarrierName
                    tempPackage.ChannelShipCode = input.Data.Order.Shipments[i].Packages[j].MappedShipcode //is MappedShipcode correct?
                    tempPackage.DateShipped = input.Data.Order.Shipments[i].Packages[j].DateShipped
                    tempPackage.RetailopsPackageID = input.Data.Order.Shipments[i].Packages[j].ID
                    tempPackage.TrackingNumber = input.Data.Order.Shipments[i].Packages[j].TrackingNumber
                    tempPackage.WeightKg = input.Data.Order.Shipments[i].Packages[j].Weight
                    // create array of PackageItems
                    tempPackage.PackageItems = make([]PackageItem, len(input.Data.Order.Shipments[i].Packages[j].ShipItems), (cap(input.Data.Order.Shipments[i].Packages[j].ShipItems)+1)*2)
                    //range over ShipItems
                    for k := range input.Data.Order.Shipments[i].Packages[j].ShipItems {
                        var tempPackageItem PackageItem
                        tempPackageItem.ChannelItemRefnum = input.Data.Order.Shipments[i].Packages[j].ShipItems[k].ChannelRefnum
                        tempPackageItem.Quantity = input.Data.Order.Shipments[i].Packages[j].ShipItems[k].Quantity
                        tempPackageItem.RetailopsOrderItemID = input.Data.Order.Shipments[i].Packages[j].ShipItems[k].ChannelOrderRefnum //is this correct?
                        tempPackageItem.RetailopsShipmentItemID = input.Data.Order.Shipments[i].Packages[j].ShipItems[k].ID
                        tempPackageItem.Sku = input.Data.Order.Shipments[i].Packages[j].ShipItems[k].SkuNum
                        // add item to packageitems
                        tempPackage.PackageItems[k] = tempPackageItem
                    }
                    //add package to shipment
                    tempShipment.Packages[i] = tempPackage
                }
                // add item to array
                shipmentArray[i] = tempShipment
        }
        output.Order.Shipments = shipmentArray

        unshippedItemsArray := make([]UnshippedItem, len(input.Data.Order.UnshippedItemsRef), (cap(input.Data.Order.UnshippedItemsRef)+1)*2)
        for i := range input.Data.Order.UnshippedItemsRef {
            var tempItem UnshippedItem
            tempItem.ChannelItemRefnum = input.Data.Order.UnshippedItemsRef[i].ChannelItemRefnum
            // tempItem.EffectiveExtendedPrice = input.Data.Order.UnshippedItemsRef[i]. //TODO: what field do we map????
            // tempItem.EffectiveUnitPrice = input.Data.Order.UnshippedItemsRef[i].  //TODO: what field do we map????
            // tempItem.OrderedQuantity = input.Data.Order.UnshippedItemsRef[i]. //TODO: what field do we map????
            tempItem.Sku = input.Data.Order.UnshippedItemsRef[i].SkuNum
            tempItem.UnshippedQuantity = input.Data.Order.UnshippedItemsRef[i].Quantity
            unshippedItemsArray[i] = tempItem
        }
        output.Order.UnshippedItems = unshippedItemsArray

        // TODO: send http POST request to API
        baseURI := input.Data.Channel.Params.BaseURI
        if len(baseURI) == 0 {
            return
        }
        
        scamp.Info.Printf("base URI: %s", baseURI)
        // TODO: munge API response back into what perl expects and return JSON below

        respMsg := scamp.NewResponseMessage()
        respMsg.WriteJson(output)
        respMsg.SetRequestId(msg.RequestId)
        _,err = client.Send(respMsg)
        if err != nil {
          return
        }
    }
}
