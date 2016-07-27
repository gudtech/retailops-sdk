package sdk_actions

type OrderUpdateV1Response struct {
	Events []struct {
		Associations []Association `json:"associations"`
		Code           string `json:"code"`
		DiagnosticData string `json:"diagnostic_data"`
		EventType      string `json:"event_type"`
		Message        string `json:"message"`
	} `json:"events"`
}

type CommonV1Response struct {
	Events []struct {
		Associations []Association `json:"associations"`
		Code           string `json:"code"`
		DiagnosticData string `json:"diagnostic_data"`
		EventType      string `json:"event_type"`
		Message        string `json:"message"`
	} `json:"events"`
}

type Association struct {
    Identifier     string `json:"identifier"`
    IdentifierType string `json:"identifier_type"`
}


//TODO: check if we can combine these structs into common response struct
type InvPushTransmitResponse struct {
    Events []InvPushTransmitResponseEvent `json:"events"`
}

type InvPushTransmitResponseEvent struct {
    Data struct {
        IsFailure int    `json:"is_failure"` // ????
        Message   string `json:"message"` // message
        Status    string `json:"status"` // code
    } `json:"data"`
    Handle    string `json:"handle"` // event_type
    Secondary []InvPushTransmitResponseSecondary `json:"secondary"`
}

type InvPushTransmitResponseSecondary struct { // associations
    Concept string `json:"concept"` // identifier_type
    ID      string `json:"id"` // identifier
}

//may be able to replace InvPushTransmitResponse with the following
type CommonResponseV1 struct{
    Events []CommonResponseEvent `json:"events"`
}

type CommonResponseEvent struct {
    Data struct {
        IsFailure int    `json:"is_failure"` // ????
        Message   string `json:"message"` // message
        Status    string `json:"status"` // code
    } `json:"data"`
    Handle    string `json:"handle"` // event_type
    Secondary []CommonResponseSecondary `json:"secondary"`
}

type CommonResponseSecondary struct { // associations
    Concept string `json:"concept"` // identifier_type
    ID      string `json:"id"` // identifier
}

//struct returned to perl service
type OrderPullResponseV1  struct {
	NextOrderRefnum int `json:"next_order_refnum"`
	NextPageState   int `json:"next_page_state"`
	Orders          []struct {
		Attributes struct{} `json:"attributes"`
		BillAddr   struct {
			Address1     string `json:"address1"`
			Address2     string `json:"address2"`
			City         string `json:"city"`
			Company      string `json:"company"`
			CountryMatch string `json:"country_match"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			PostalCode   string `json:"postal_code"`
			StateMatch   string `json:"state_match"`
		} `json:"bill_addr"`
		CalcMode           string `json:"calc_mode"`
		ChannelDateCreated int    `json:"channel_date_created"`
		ChannelRefnum      string `json:"channel_refnum"`
		Customer           struct {
			EmailAddress string `json:"email_address"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			PhoneNumber  string `json:"phone_number"`
		} `json:"customer"`
		DiscountAmt int         `json:"discount_amt"`
		GiftMessage interface{} `json:"gift_message"`
		IPAddress   string      `json:"ip_address"`
		Items       []struct {
			ChannelRefnum string `json:"channel_refnum"`
			Quantity      int    `json:"quantity"`
			Sku           string `json:"sku"`
			SkuTitle      string `json:"sku_title"`
			UnitPrice     string `json:"unit_price"`
			UnitTax       int    `json:"unit_tax"`
		} `json:"items"`
		Payment []struct {
			Amount string `json:"amount"`
			Params struct {
				ChannelRefnum string `json:"channel_refnum"`
				PaymentType   string `json:"payment_type"`
			} `json:"params"`
			Type string `json:"type"`
		} `json:"payment"`
		ShipAddr struct {
			Address1     string `json:"address1"`
			Address2     string `json:"address2"`
			City         string `json:"city"`
			Company      string `json:"company"`
			CountryMatch string `json:"country_match"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			PostalCode   string `json:"postal_code"`
			StateMatch   string `json:"state_match"`
		} `json:"ship_addr"`
		Shipcode    string `json:"shipcode"`
		ShippingAmt string `json:"shipping_amt"`
		TaxAmt      string `json:"tax_amt"`
	} `json:"orders"`
}
