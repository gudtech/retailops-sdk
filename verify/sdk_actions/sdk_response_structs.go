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
