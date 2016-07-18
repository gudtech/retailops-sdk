package sdk_actions

import (
    "github.com/gudtech/scamp-go/scamp"
    "encoding/json"
)

type OrderShipmentSubmitV1Input struct {
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
    		Order struct {
    			ChannelPayment struct {
    				Available           string `json:"available"`
    				Captured            int    `json:"captured"`
    				CapturesAreDeferred int    `json:"captures_are_deferred"`
    				CapturesAreExternal int    `json:"captures_are_external"`
    				ChannelID           int    `json:"channel_id"`
    				Charged             string `json:"charged"`
    				Credited            int    `json:"credited"`
    				ID                  string `json:"id"`
    				Method              string `json:"method"`
    				Module              string `json:"module"`
    				Settled             string `json:"settled"`
    				Unsettled           int    `json:"unsettled"`
    				Voided              int    `json:"voided"`
    			} `json:"channel_payment"`
    			ChannelRefnum        string `json:"channel_refnum"`
    			FromCounterpartyRate int    `json:"from_counterparty_rate"`
    			GrandTotal           string `json:"grand_total"`
    			ID                   string `json:"id"`
    			PaymentSeriesID      string `json:"payment_series_id"`
    			PaymentStatus        struct {
    				Authed    int    `json:"authed"`
    				Available string `json:"available"`
    				ByAccount []struct {
    					Available           string `json:"available"`
    					Captured            int    `json:"captured"`
    					CapturesAreDeferred int    `json:"captures_are_deferred"`
    					CapturesAreExternal int    `json:"captures_are_external"`
    					ChannelID           int    `json:"channel_id"`
    					Charged             string `json:"charged"`
    					Credited            int    `json:"credited"`
    					ID                  string `json:"id"`
    					Method              string `json:"method"`
    					Module              string `json:"module"`
    					Settled             string `json:"settled"`
    					Unsettled           int    `json:"unsettled"`
    					Voided              int    `json:"voided"`
    				} `json:"by_account"`
    				Captured          int    `json:"captured"`
    				Charged           string `json:"charged"`
    				Credited          int    `json:"credited"`
    				Settled           string `json:"settled"`
    				Success           int    `json:"success"`
    				Unsettled         int    `json:"unsettled"`
    				UnsettledDeferred int    `json:"unsettled_deferred"`
    				UnsettledExternal int    `json:"unsettled_external"`
    			} `json:"payment_status"`
    			ShipServiceName string `json:"ship_service_name"`
    			Shipments       []struct {
    				ID       string `json:"id"`
    				Packages []struct {
    					CarrierClassCode string      `json:"carrier_class_code"`
    					CarrierClassName string      `json:"carrier_class_name"`
    					CarrierCode      string      `json:"carrier_code"`
    					CarrierName      string      `json:"carrier_name"`
    					ClassName        string      `json:"class_name"`
    					DateShipped      string      `json:"date_shipped"`
    					ID               string      `json:"id"`
    					MappedShipcode   interface{} `json:"mapped_shipcode"`
    					ShipItems        []struct {
    						ChannelOrderRefnum string `json:"channel_order_refnum"`
    						ChannelRefnum      string `json:"channel_refnum"`
    						ID                 string `json:"id"`
    						ItemGiftwrapAmt    int    `json:"item_giftwrap_amt"`
    						ItemGiftwrapTaxAmt int    `json:"item_giftwrap_tax_amt"`
    						ItemProductAmt     string `json:"item_product_amt"`
    						ItemRecyclingAmt   int    `json:"item_recycling_amt"`
    						ItemShippingAmt    string `json:"item_shipping_amt"`
    						ItemShippingTaxAmt int    `json:"item_shipping_tax_amt"`
    						ItemSubtotalAmt    int    `json:"item_subtotal_amt"`
    						ItemTaxAmt         string `json:"item_tax_amt"`
    						Quantity           int    `json:"quantity"`
    						SkuNum             string `json:"sku_num"`
    					} `json:"ship_items"`
    					TrackingNumber interface{} `json:"tracking_number"`
    					Weight         string      `json:"weight"`
    				} `json:"packages"`
    			} `json:"shipments"`
    			UnshippedItemsRef []interface{} `json:"unshipped_items_ref"`
    		} `json:"order"`
    		Shipment struct {
    			ID       string `json:"id"`
    			Packages []struct {
    				CarrierClassCode string      `json:"carrier_class_code"`
    				CarrierClassName string      `json:"carrier_class_name"`
    				CarrierCode      string      `json:"carrier_code"`
    				CarrierName      string      `json:"carrier_name"`
    				ClassName        string      `json:"class_name"`
    				DateShipped      string      `json:"date_shipped"`
    				ID               string      `json:"id"`
    				MappedShipcode   interface{} `json:"mapped_shipcode"`
    				ShipItems        []struct {
    					ChannelOrderRefnum string `json:"channel_order_refnum"`
    					ChannelRefnum      string `json:"channel_refnum"`
    					ID                 string `json:"id"`
    					ItemGiftwrapAmt    int    `json:"item_giftwrap_amt"`
    					ItemGiftwrapTaxAmt int    `json:"item_giftwrap_tax_amt"`
    					ItemProductAmt     string `json:"item_product_amt"`
    					ItemRecyclingAmt   int    `json:"item_recycling_amt"`
    					ItemShippingAmt    string `json:"item_shipping_amt"`
    					ItemShippingTaxAmt int    `json:"item_shipping_tax_amt"`
    					ItemSubtotalAmt    int    `json:"item_subtotal_amt"`
    					ItemTaxAmt         string `json:"item_tax_amt"`
    					Quantity           int    `json:"quantity"`
    					SkuNum             string `json:"sku_num"`
    				} `json:"ship_items"`
    				TrackingNumber interface{} `json:"tracking_number"`
    				Weight         string      `json:"weight"`
    			} `json:"packages"`
    		} `json:"shipment"`
    	} `json:"data"`
    	Headers struct {
    		ClientID int    `json:"client_id"`
    		Ticket   string `json:"ticket"`
    	} `json:"headers"`
    	Version int `json:"version"`
}

type OrderShipmentSubmitV1Output struct {
    Action      string `json:"action"`
	ChannelInfo struct {
		ID int `json:"id"`
	} `json:"channel_info"`
	ChannelOrderRefnum   string `json:"channel_order_refnum"`
	ClientID             int    `json:"client_id"`
	IntegrationAuthToken string `json:"integration_auth_token"`
	RetailopsOrderID     int    `json:"retailops_order_id"`
	Shipment             struct {
		Packages []Package `json:"packages"` //Package and PackageItem structs declared in callback_items_returned.go
		RetailopsShipmentID int `json:"retailops_shipment_id"`
	} `json:"shipment"`
	Version int `json:"version"`
}

func OrderShipmentSubmitV1(msg *scamp.Message, client *scamp.Client) {
    scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
    var input OrderShipmentSubmitV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %s ", input)
    }

    //TODO: need to munge actual input data to output format for sdk
    var output OrderShipmentSubmitV1Output

    respMsg := scamp.NewResponseMessage()
    respMsg.WriteJson(output)
    respMsg.SetRequestId(msg.RequestId)

    _,err = client.Send(respMsg)
    if err != nil {
      return
    }
}
