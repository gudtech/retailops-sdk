package sdk_actions

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

//API reponse struct
type OrderUpdateV1Response struct {
	Events []struct {
		Associations []struct {
			Identifier     string `json:"identifier"`
			IdentifierType string `json:"identifier_type"`
		} `json:"associations"`
		Code           string `json:"code"`
		DiagnosticData string `json:"diagnostic_data"`
		EventType      string `json:"event_type"`
		Message        string `json:"message"`
	} `json:"events"`
}

//for testing only
type OrderUpdateV1InvalidResponse struct {
	Events []struct {
		Associations []struct {
			IdentifierTEST     string `json:"identifier2"`
			IdentifierTypeTEST string `json:"identifier_type2"`
		} `json:"associations"`
		Code           string `json:"code"`
		DiagnosticData string `json:"diagnostic_data"`
		EventType      string `json:"event_type"`
		Message        string `json:"message"`
	} `json:"events"`
}
