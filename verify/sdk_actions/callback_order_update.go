package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
  "bytes"
  "strings"
  "net/http"
  "time"
)

func OrderUpdateV1(msg *scamp.Message, client *scamp.Client) {
    // var err error
    var input OrderUpdateV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %+v\n ", err)
        respMsg := scamp.NewResponseMessage()
        respMsg.SetError(err.Error())
        respMsg.Write([]byte(err.Error()))
        respMsg.SetRequestId(msg.RequestId)
        _,err := client.Send(respMsg)
        if err != nil {
          return
        }
    } else {
        var output OrderUpdateV1Output
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
                        tempPackageItem.Quantity = input.Data.Order.Shipments[i].Packages[j].ShipItems[k].Quantity
                        // add packageitem to packageitems
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

        // build uri
        baseURI := input.Data.Channel.Params.BaseURI
        if len(baseURI) == 0 {
            return
        }

        //test if baseURI ends with a slash and append action accordingly
        var buffer bytes.Buffer
        buffer.WriteString(baseURI)
        if strings.HasSuffix(baseURI, "/") {
            buffer.WriteString("order_update_v1")
        }else{
            buffer.WriteString("/order_update_v1")
        }

        channelURI := buffer.String()
        var requestBuffer bytes.Buffer
        err := json.NewEncoder(&requestBuffer).Encode(output)
        if err != nil {
          return
        }

        // send http POST request to API
        var httpClient = &http.Client{
          Timeout: time.Minute * 5, //NOTE: this timeout seems kind of long
        }

        response,err := httpClient.Post(channelURI, "application/json", &requestBuffer)
        defer response.Body.Close()
        if err != nil {
          return
        }

        var apiResp OrderUpdateV1InvalidResponse //OrderUpdateV1Response
        err = json.NewDecoder(response.Body).Decode(&apiResp)
        if err != nil {
          return
        }

        // responseJSON, err := json.Marshal(&apiResp)
        // if err != nil {
        //     scamp.Info.Printf("responseJSON err: %s", err)
        // }

        // testJSON := `"{"test": "test", "test2": "test2"}"`
        validResponse, err := ValidateResponse("../verify/schema/order_update_v1.json", &apiResp )//string(responseJSON)
        if err != nil {
            //TODO see what type of error I should return to ROP when service response fails to validate
            scamp.Info.Printf("There was an error validating the response: %+v\n ", err)
            respMsg := scamp.NewResponseMessage()
            respMsg.SetError(err.Error())
            respMsg.Write([]byte(err.Error()))
            respMsg.SetRequestId(msg.RequestId)
            _,err := client.Send(respMsg)
            if err != nil {
                scamp.Info.Printf("SDK: callback_order_update error %s", err)
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
                scamp.Info.Printf("SDK: callback_order_update error %s", err)
            }
            return
        }

        // TODO: munge API response back into what perl expects and return JSON below

        // encode and return JSON
        respMsg := scamp.NewResponseMessage()
        respMsg.WriteJson(apiResp)
        respMsg.SetRequestId(msg.RequestId)
        _,err = client.Send(respMsg)
        if err != nil {
            scamp.Info.Printf("SDK: callback_order_update error %s", err)
            return
        }
    }
}
