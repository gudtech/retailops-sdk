package sdk_actions

type OrderUpdateV1Input struct {
    Action string `json:"action"`
	Data   struct {
		Channel struct {
			ID     int `json:"id,string"`
			Params struct {
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
            Definition struct {
              	Handle string `json:"handle"`
              	Params struct {
              			Interactions []ChannelInteraction `json:"interactions"`
              	} `json:"params"`
             } `json:"definition"`
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
			Shipments         []ROPInputShipment `json:"shipments"`
			UnshippedItemsRef []ROPInputUnshippedItem `json:"unshipped_items_ref"`
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
