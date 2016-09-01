package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
  "bytes"
  "net/http"
  "time"
  "fmt"
  "strings"
  "strconv"
)
//NOTE: We may need to convert Channel. Params to take an interface because params vary widely between channels
type OrderPullV1Input struct {
	Channel struct {
		ID     int `json:"id"`
		Params struct {
			StoreID                    string `json:"StoreID"`
			NextOrderRefnum            int    `json:"next_order_refnum,"`
			OrderAckStatusID           string `json:"order_ack_status_id"`
			OrderFulfilledStatusID     string `json:"order_fulfilled_status_id"`
			OrderInFilfillmentStatusID string `json:"order_in_filfillment_status_id"`
		} `json:"params"`
        Definition struct {
          	Handle string `json:"handle"`
          	Params struct {
          			Interactions []ChannelInteraction `json:"interactions"`
          	} `json:"params"`
         } `json:"definition"`
	} `json:"channel"`
	ClientID    int `json:"client_id,string"`
	MaxPageSize int `json:"max_page_size"`
	Order       struct {
		ChannelRefnum int `json:"channel_refnum"`
	} `json:"order"`
	// PageState interface{} `json:"page_state"`
    PageState string `json:"page_state"`
	Single    int         `json:"single"`
}

// type OrderPullV1Input struct {
//     Action string `json:"action"`
// 	Data   struct {
// 		Channel struct {
// 			ID     int `json:"id"`
// 			Params struct {
//                 BaseURI                    string `json:"base_uri"`//added was not in original perl request
// 				StoreID                    string `json:"StoreID"`
// 				NextOrderRefnum            int    `json:"next_order_refnum,"`
// 				OrderAckStatusID           string `json:"order_ack_status_id"`
// 				OrderFulfilledStatusID     string `json:"order_fulfilled_status_id"`
// 				OrderInFilfillmentStatusID string `json:"order_in_filfillment_status_id"`
// 			} `json:"params"`
//             Definition struct {
//               	Handle string `json:"handle"`
//               	Params struct {
//               			Interactions []ChannelInteraction `json:"interactions"`
//               	} `json:"params"`
//              } `json:"definition"`
// 		} `json:"channel"`
// 		ClientID    int `json:"client_id"`
// 		MaxPageSize int `json:"max_page_size"`
// 		Order       struct {
// 			ChannelRefnum int `json:"channel_refnum"`
// 		} `json:"order"`
// 		// PageState interface{} `json:"page_state"`
//         PageState string `json:"page_state"`
// 		Single    int         `json:"single"`
// 	} `json:"data"`
// 	Headers struct {
// 		ClientID int    `json:"client_id"`
// 		Ticket   string `json:"ticket"`
// 	} `json:"headers"`
// 	Version int `json:"version"`
// }

type OrderPullV1Output struct {
    Action      string `json:"action"`
	ChannelInfo struct {
		ID int `json:"id"`
	} `json:"channel_info"`
	ClientID             int    `json:"client_id"`
	IntegrationAuthToken string `json:"integration_auth_token"`
	PageToken            string `json:"page_token"`
	SpecificOrders       []SpecificOrder `json:"specific_orders"`
	Version int64 `json:"version"`
}

type SpecificOrder struct {
    ChannelRefnum string `json:"channel_refnum"`
}

func OrderPullV1(msg *scamp.Message, client *scamp.Client) {
    var input OrderPullV1Input
    // scamp.Info.Printf("incoming: %s", string(msg.Bytes()))

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %+v\n ", err)
        respMsg := scamp.NewResponseMessage()
        respMsg.SetError(err.Error())
        respMsg.Write([]byte(err.Error()))
        respMsg.SetRequestId(msg.RequestId)
        _,err := client.Send(respMsg)
        if err != nil {
            scamp.Info.Printf("SDK: callback_order_update error %s", err)
        }
        return
    } else {
        handle := input.Channel.Definition.Handle
        if len(handle) == 0 {
            scamp.Info.Printf("SDK: handle cannot be blank")
            return
        }

        lastIndex := strings.LastIndex(handle, "_")
        integration_auth_token := handle[lastIndex + 1:len(handle)]
        if len(integration_auth_token) == 0 {
            scamp.Info.Printf("SDK: integration_auth_token could not be extracted from handle")
            return //to do return formatted scamp error message
        }

        var output OrderPullV1Output
        //TODO: need to munge actual input data to output format for sdk
        output.Action = msg.Action
        output.ChannelInfo.ID = input.Channel.ID
        output.ClientID = input.ClientID
        output.IntegrationAuthToken = integration_auth_token
        output.Version = msg.Version
        output.PageToken = input.PageState


        var endPointURI string
        var version int
        interactions := input.Channel.Definition.Params.Interactions
        for i := range interactions {
            if interactions[i].Action == "order_pull" {
                endPointURI = interactions[i].EndpointURL
                version = interactions[i].Version
            }
        }

        if len(endPointURI) == 0 || version <= 0 {
            scamp.Info.Printf("endpoint or version is blank")
            return
        }
        channelURI := BuildURI(endPointURI, version )

        var requestBuffer bytes.Buffer
        err := json.NewEncoder(&requestBuffer).Encode(output)
        if err != nil {
            return
        }

        var httpClient = &http.Client{
            Timeout: time.Minute * 5,
        }

        scamp.Info.Printf("Making API call to: %s", channelURI)
        scamp.Info.Printf("Token: %s", integration_auth_token)
        response,err := httpClient.Post(channelURI, "application/json", &requestBuffer)
        if err != nil {
            scamp.Info.Printf("err: %s\n", err)
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(err.Error())
            respMsg.Write([]byte(err.Error()))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: order_pull error %s", err)
            }
            return
        }
        defer response.Body.Close()

        if response.StatusCode != 200 {
            // scamp.Info.Printf("Request: %s", reqBody)
            var errMsg string
            errMsg = fmt.Sprintf("%s", response.Status)
            scamp.Info.Printf("%s", response.Status)
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(errMsg)
            respMsg.Write([]byte(response.Status))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: order_pull error %s", err)
            }
            return
        }

        var apiResp OrderPullSDKResponseV1
        err = json.NewDecoder(response.Body).Decode(&apiResp)
        if err != nil {
            scamp.Info.Printf("Could not unmarshal SDK response: %s", err)
            return
        }

        validResponse, err := ValidateResponse("./src/github.com/gudtech/retailops-sdk/verify/schema/order_pull_v1.json", &apiResp )
        if err != nil {
            scamp.Info.Printf("There was an error validating the response: %+v\n ", err)
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(err.Error())
            respMsg.Write([]byte(err.Error()))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: order_pull error %s", err)
            }
            return
        }

        if !validResponse {
            validationMsg := "API response was invalid"
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(validationMsg)
            respMsg.Write([]byte(validationMsg))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: order_pull error %s", err)
            }
            return
        }

        var orderFetchResponse OrderPullResponseV1
        orderFetchResponse.NextOrderRefnum = 0 //this is not currently defined in API response!
        orderFetchResponse.NextPageState = 1 //temp passing "1"
        orderFetchResponse.NextPageToken = apiResp.NextPageToken

        orderArray := make([]ROPOrder, len(apiResp.Orders), (cap(apiResp.Orders)+1)*2)

        for i := range apiResp.Orders {
            var tempOrder ROPOrder
            tempOrder.Attributes = ""

            // NOTE send to ROP as single object with unique fields
            // attributes => {
            //     <attribute_type>_<handle> => <value>,
            //     <attribute_type>_<handle> => <value>,
            //      ...
            // }
            
            tempOrder.CalcMode = ""  //NOTE: see #RO_SDK slack channel for discussion

            //convert date to Unix timestamp
            t, err := time.Parse(time.RFC3339Nano, apiResp.Orders[i].ChannelDateCreated)
        	if err != nil {
                scamp.Info.Printf("Could not parse ChannelDateCreated: %s", err)
        		return
        	}
        	unixTime := t.Unix()
        	//fmt.Println(unixTime)
            tempOrder.ChannelDateCreated = unixTime

            tempOrder.ChannelRefnum = apiResp.Orders[i].ChannelOrderRefnum

            tempOrder.Customer.EmailAddress = apiResp.Orders[i].CustomerInfo.EmailAddress
            tempOrder.Customer.PhoneNumber = apiResp.Orders[i].CustomerInfo.PhoneNumber
            //NOTE: ROP is concatinating first + last so just send full in first and empty string in last
            tempOrder.Customer.FirstName = apiResp.Orders[i].CustomerInfo.FullName
            tempOrder.Customer.LastName = ""

            //TODO: convert DiscountAmt to float in both structs
            tempOrder.DiscountAmt = apiResp.Orders[i].CurrencyValues.DiscountAmt
            tempOrder.GiftMessage = apiResp.Orders[i].GiftMessage
            tempOrder.IPAddress = apiResp.Orders[i].IPAddress
            tempOrder.Shipcode = apiResp.Orders[i].ShipServiceCode

            tempOrder.ShippingAmt = strconv.FormatFloat(apiResp.Orders[i].CurrencyValues.ShippingAmt, 'f', -1, 32)
            tempOrder.TaxAmt = strconv.FormatFloat(apiResp.Orders[i].CurrencyValues.TaxAmt, 'f', -1, 32)

            //order items array
            orderItemsArray := make([]ROPOrderItem, len(apiResp.Orders[i].OrderItems), (cap(apiResp.Orders[i].OrderItems)+1)*2)
            for j := range apiResp.Orders[i].OrderItems {
                var tempItem ROPOrderItem
                tempItem.ChannelRefnum = apiResp.Orders[i].OrderItems[j].ChannelItemRefnum
                tempItem.Quantity = apiResp.Orders[i].OrderItems[j].Quantity
                tempItem.Sku = apiResp.Orders[i].OrderItems[j].Sku
                tempItem.SkuTitle = apiResp.Orders[i].OrderItems[j].SkuDescription
                tempItem.UnitPrice = strconv.FormatFloat(apiResp.Orders[i].OrderItems[j].CurrencyValues.UnitPrice, 'f', -1, 32)
                tempItem.UnitTax = apiResp.Orders[i].OrderItems[j].CurrencyValues.UnitTax
                orderItemsArray[j] = tempItem
            }
            tempOrder.Items = orderItemsArray

            //payments array
            paymentArray := make([]ROPOrderPayment, len(apiResp.Orders[i].PaymentTransactions), (cap(apiResp.Orders[i].PaymentTransactions)+1)*2)
            for k := range apiResp.Orders[i].PaymentTransactions {
                var tempPayment ROPOrderPayment
                tempPayment.Amount = strconv.FormatFloat(apiResp.Orders[i].PaymentTransactions[k].Amount, 'f', -1, 32)
                tempPayment.Type = apiResp.Orders[i].PaymentTransactions[k].TransactionType
                tempPayment.Params.ChannelRefnum = apiResp.Orders[i].ChannelOrderRefnum
                tempPayment.Params.PaymentType = apiResp.Orders[i].PaymentTransactions[k].PaymentType
                paymentArray[k] = tempPayment
            }
            tempOrder.Payment = paymentArray

            //ship_addr
            tempOrder.ShipAddr.Address1 = apiResp.Orders[i].ShippingAddress.Address1
            tempOrder.ShipAddr.Address2 = apiResp.Orders[i].ShippingAddress.Address2
            tempOrder.ShipAddr.City = apiResp.Orders[i].ShippingAddress.City
            tempOrder.ShipAddr.Company = apiResp.Orders[i].ShippingAddress.Company
            tempOrder.ShipAddr.CountryMatch = apiResp.Orders[i].ShippingAddress.CountryMatch
            tempOrder.ShipAddr.FirstName = apiResp.Orders[i].ShippingAddress.FirstName
            tempOrder.ShipAddr.LastName = apiResp.Orders[i].ShippingAddress.LastName
            tempOrder.ShipAddr.PostalCode = apiResp.Orders[i].ShippingAddress.PostalCode
            tempOrder.ShipAddr.StateMatch = apiResp.Orders[i].ShippingAddress.StateMatch

            //billing_address
            tempOrder.BillAddr.Address1 = apiResp.Orders[i].BillingAddress.Address1
            tempOrder.BillAddr.Address2 = apiResp.Orders[i].BillingAddress.Address2
            tempOrder.BillAddr.City = apiResp.Orders[i].BillingAddress.City
            tempOrder.BillAddr.Company = apiResp.Orders[i].BillingAddress.Company
            tempOrder.BillAddr.CountryMatch = apiResp.Orders[i].BillingAddress.CountryMatch
            tempOrder.BillAddr.FirstName = apiResp.Orders[i].BillingAddress.FirstName
            tempOrder.BillAddr.LastName = apiResp.Orders[i].BillingAddress.LastName
            tempOrder.BillAddr.PostalCode = apiResp.Orders[i].BillingAddress.PostalCode
            tempOrder.BillAddr.StateMatch = apiResp.Orders[i].BillingAddress.StateMatch

            orderArray[i] = tempOrder
        }

        orderFetchResponse.Orders = orderArray
        // munge data

        respMsg := scamp.NewResponseMessage()
        respMsg.WriteJson(orderFetchResponse)
        respMsg.SetRequestId(msg.RequestId)

        _,err = client.Send(respMsg)
        if err != nil {
          return
        }
    }
}
