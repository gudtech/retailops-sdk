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
	Events []Event `json:"events"`
}

type Event struct {
    Associations []Association `json:"associations"`
    Code           string `json:"code"`
    DiagnosticData string `json:"diagnostic_data"`
    EventType      string `json:"event_type"`
    Message        string `json:"message"`
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

type OrderPullSDKResponseV1 struct {
	NextPageToken string `json:"next_page_token"`
    // NextPageState string `json:"next_page_state"`
    NextOrderRefnum int `json:"next_order_refnum"` // TODO: Update swagger: new field X
	Orders        []OrderSDK `json:"orders"`
}

type OrderSDK struct {
    Attributes []OrderAttributeSDK `json:"attributes"`
    CalcMode           string `json:"calc_mode"` //TODO: Update Swagger: new field
    BillingAddress struct {
        Address1     string `json:"address1"`
        Address2     string `json:"address2"`
        City         string `json:"city"`
        Company      string `json:"company"`
        CountryMatch string `json:"country_match"`
        FirstName    string `json:"first_name"`
        LastName     string `json:"last_name"`
        PostalCode   string `json:"postal_code"`
        StateMatch   string `json:"state_match"`
    } `json:"billing_address"`
    ChannelDateCreated string `json:"channel_date_created"`
    ChannelOrderRefnum string `json:"channel_order_refnum"`
    CurrencyCode       string `json:"currency_code"`
    CurrencyValues     struct {
        DiscountAmt float64 `json:"discount_amt"`
        ShippingAmt float64 `json:"shipping_amt"`
        TaxAmt      float64 `json:"tax_amt"`
    } `json:"currency_values"`
    CustomerInfo struct {
        EmailAddress string `json:"email_address"`
        FullName     string `json:"full_name"`
        PhoneNumber  string `json:"phone_number"`
    } `json:"customer_info"`
    GiftMessage string `json:"gift_message"`
    IPAddress   string `json:"ip_address"`
    OrderItems  []OrderItemSDK `json:"order_items"`
    PaymentTransactions []PaymentTransactionSDK `json:"payment_transactions"`
    ShipServiceCode string `json:"ship_service_code"`
    ShippingAddress struct {
        Address1     string `json:"address1"`
        Address2     string `json:"address2"`
        City         string `json:"city"`
        Company      string `json:"company"`
        CountryMatch string `json:"country_match"`
        FirstName    string `json:"first_name"`
        LastName     string `json:"last_name"`
        PostalCode   string `json:"postal_code"`
        StateMatch   string `json:"state_match"`
    } `json:"shipping_address"`
}

type OrderAttributeSDK struct {
    AttributeType string `json:"attribute_type"`
    Handle        string `json:"handle"`
    Value         string `json:"value"` //TODO: Update Swagger: new field X
}

type OrderItemSDK struct {
    ChannelItemRefnum string `json:"channel_item_refnum"`
    CurrencyValues    struct {
        DiscountAmt  float64 `json:"discount_amt"`
        DiscountPct  float64 `json:"discount_pct"`
        RecyclingAmt float64 `json:"recycling_amt"`
        ShipAmt      float64 `json:"ship_amt"`
        ShiptaxAmt   float64 `json:"shiptax_amt"`
        UnitPrice    float64 `json:"unit_price"`
        UnitTax      float64 `json:"unit_tax"`
        UnitTaxPct   float64 `json:"unit_tax_pct"`
        VatPct       float64 `json:"vat_pct"`
    } `json:"currency_values"`
    ItemType       string `json:"item_type"`
    Quantity       int    `json:"quantity"`
    Sku            string `json:"sku"`
    SkuDescription string `json:"sku_description"`
}

type PaymentTransactionSDK struct {
    Amount                float64 `json:"amount"`
    PaymentProcessingType string  `json:"payment_processing_type"`
    PaymentType           string  `json:"payment_type"`
    TransactionType       string  `json:"transaction_type"`
}

/*
* structs returned to RetailOps perl service
*/

type OrderPullResponseV1  struct {
	NextOrderRefnum int `json:"next_order_refnum"`
	NextPageState   int `json:"next_page_state"`
    NextPageToken   string `json:"next_page_token"` //NOTE: currently ignored by ROP
	Orders          []ROPOrder `json:"orders"`
}

type ROPOrder struct {
    Attributes string `json:"attributes"`
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
    ChannelDateCreated int64  `json:"channel_date_created"`
    ChannelRefnum      string `json:"channel_refnum"`
    Customer           struct {
        EmailAddress string `json:"email_address"`
        FirstName    string `json:"first_name"`
        LastName     string `json:"last_name"`
        PhoneNumber  string `json:"phone_number"`
    } `json:"customer"`
    DiscountAmt int         `json:"discount_amt"`
    GiftMessage string      `json:"gift_message"`
    IPAddress   string      `json:"ip_address"`
    Items       []ROPOrderItem `json:"items"`
    Payment []ROPOrderPayment `json:"payment"`
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
}

type ROPOrderPayment struct {
    Amount string `json:"amount"`
    Params struct {
        ChannelRefnum string `json:"channel_refnum"`
        PaymentType   string `json:"payment_type"`
    } `json:"params"`
    Type string `json:"type"`
}

type ROPOrderItem struct {
    ChannelRefnum string `json:"channel_refnum"`
    Quantity      int    `json:"quantity"`
    Sku           string `json:"sku"`
    SkuTitle      string `json:"sku_title"`
    UnitPrice     string `json:"unit_price"`
    UnitTax       float64 `json:"unit_tax"`
}
