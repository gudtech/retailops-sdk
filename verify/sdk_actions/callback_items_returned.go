package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
)
//TODO: both input and output structs need to have internal arrays
//defined as structs
type OrderReturnedV1Input struct {
    Action string `json:"action"`
	Data   struct {
		Channel struct {
			ID     string `json:"id"`
			Params struct {
				BaseURI                       string      `json:"base_uri"`
				EmailInvoice                  int         `json:"email_invoice"`
				EmailReturn                   int         `json:"email_return"`
				EmailTracking                 int         `json:"email_tracking"`
				ExpressConfigurableSuperLinks int         `json:"express_configurable_super_links"`
				ImportOrderAttrs              string      `json:"import_order_attrs"`
				InvSuspendedInstock           int         `json:"inv_suspended_instock"`
				InvSuspendedMode              interface{} `json:"inv_suspended_mode"`
				PushCancel                    int         `json:"push_cancel"`
				UnsetOtherAttributes          int         `json:"unset_other_attributes"`
				UnsetOtherMedia               int         `json:"unset_other_media"`
			} `json:"params"`
		} `json:"channel"`
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
			GrandTotal           string `json:"grand_total"`
			ID                   string `json:"id"`
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
						ItemShippingAmt    int    `json:"item_shipping_amt"`
						ItemShippingTaxAmt int    `json:"item_shipping_tax_amt"`
						ItemSubtotalAmt    int    `json:"item_subtotal_amt"`
						ItemTaxAmt         int    `json:"item_tax_amt"`
						Quantity           int    `json:"quantity"`
						SkuNum             string `json:"sku_num"`
					} `json:"ship_items"`
					TrackingNumber interface{} `json:"tracking_number"`
					Weight         string      `json:"weight"`
				} `json:"packages"`
			} `json:"shipments"`
			UnshippedItemsRef []interface{} `json:"unshipped_items_ref"`
		} `json:"order"`
		Return struct {
			CreditItemsRef []struct {
				ChannelID          string `json:"channel_id"`
				ChannelItemRefnum  string `json:"channel_item_refnum"`
				ChannelOrderRefnum string `json:"channel_order_refnum"`
				CreditItemRefnum   string `json:"credit_item_refnum"`
				ItemCreditAmt      string `json:"item_credit_amt"`
				ItemGiftwrapAmt    string `json:"item_giftwrap_amt"`
				ItemGiftwrapTaxAmt string `json:"item_giftwrap_tax_amt"`
				ItemProductAmt     string `json:"item_product_amt"`
				ItemRecyclingAmt   string `json:"item_recycling_amt"`
				ItemRestockFeeAmt  string `json:"item_restock_fee_amt"`
				ItemShippingAmt    string `json:"item_shipping_amt"`
				ItemShippingTaxAmt string `json:"item_shipping_tax_amt"`
				ItemSubtotalAmt    string `json:"item_subtotal_amt"`
				ItemTaxAmt         string `json:"item_tax_amt"`
				Quantity           string `json:"quantity"`
				Reason             string `json:"reason"`
				Sku                string `json:"sku"`
			} `json:"credit_items_ref"`
			DiscountAmt string `json:"discount_amt"`
			ID          string `json:"id"`
			Items       []struct {
				ChannelRefnum string `json:"channel_refnum"`
				OrderItemID   string `json:"order_item_id"`
				Quantity      string `json:"quantity"`
				Sku           string `json:"sku"`
			} `json:"items"`
			ProductAmt   string      `json:"product_amt"`
			RefundAction string      `json:"refund_action"`
			RefundAmt    string      `json:"refund_amt"`
			RmaID        interface{} `json:"rma_id"`
			ShippingAmt  string      `json:"shipping_amt"`
			SubtotalAmt  string      `json:"subtotal_amt"`
			TaxAmt       string      `json:"tax_amt"`
		} `json:"return"`
	} `json:"data"`
	Headers struct {
		ClientID int    `json:"client_id"`
		Ticket   string `json:"ticket"`
	} `json:"headers"`
	Version int `json:"version"`
}

//output structs
type OrderReturnedV1Output struct {
    Action      string `json:"action"`
	ChannelInfo struct {
		ID int `json:"id"`
	} `json:"channel_info"`
	ClientID             int    `json:"client_id"`
	IntegrationAuthToken string `json:"integration_auth_token"`
	Order                struct {
		ChannelOrderRefnum string `json:"channel_order_refnum"`
		GrandTotal         int    `json:"grand_total"`
		RetailopsOrderID   int    `json:"retailops_order_id"`
		Shipments          []Shipment `json:"shipments"`
		UnshippedItems []UnshippedItem `json:"unshipped_items"`
	} `json:"order"`
	Return struct {
		DiscountAmt int `json:"discount_amt"`
		Items       []ReturnItem `json:"items"`
		ProductAmt        int    `json:"product_amt"`
		RefundAction      string `json:"refund_action"`
		RefundAmt         int    `json:"refund_amt"`
		RetailopsReturnID int    `json:"retailops_return_id"`
		RetailopsRmaID    string `json:"retailops_rma_id"`
		ShippingAmt       int    `json:"shipping_amt"`
		SubtotalAmt       int    `json:"subtotal_amt"`
		TaxAmt            int    `json:"tax_amt"`
	} `json:"return"`
	Version int `json:"version"`
}

func OrderReturnedV1(msg *scamp.Message, client *scamp.Client) {
    scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
    var input OrderReturnedV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %s ", input)
    }

    //TODO: need to munge actual input data to output format for sdk
    var output OrderReturnedV1Output

    respMsg := scamp.NewResponseMessage()
    respMsg.WriteJson(output)
    respMsg.SetRequestId(msg.RequestId)

    _,err = client.Send(respMsg)
    if err != nil {
      return
    }
}
