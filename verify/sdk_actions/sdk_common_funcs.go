package sdk_actions

import (
    "bytes"
    "strings"
)

func BuildShipments(input []ROPInputShipment) ([]Shipment) {
    shipmentArray := make([]Shipment, len(input), (cap(input)+1)*2) //NOTE: do we really need to double?
    for i := range input {
            var tempShipment Shipment
            tempShipment.RetailopsShipmentID = input[i].ID

            tempShipment.Packages = make([]Package, len(input[i].Packages), (cap(input[i].Packages)+1)*2)
            for j := range input[i].Packages {
                var tempPackage Package
                tempPackage.CarrierClassName = input[i].Packages[j].CarrierClassName
                tempPackage.CarrierName = input[i].Packages[j].CarrierName
                tempPackage.ChannelShipCode = input[i].Packages[j].MappedShipcode //is MappedShipcode correct?
                tempPackage.DateShipped = input[i].Packages[j].DateShipped
                tempPackage.RetailopsPackageID = input[i].Packages[j].ID
                tempPackage.TrackingNumber = input[i].Packages[j].TrackingNumber
                tempPackage.WeightKg = input[i].Packages[j].Weight

                tempPackage.PackageItems = make([]PackageItem, len(input[i].Packages[j].ShipItems), (cap(input[i].Packages[j].ShipItems)+1)*2)
                for k := range input[i].Packages[j].ShipItems {
                    var tempPackageItem PackageItem
                    tempPackageItem.ChannelItemRefnum = input[i].Packages[j].ShipItems[k].ChannelRefnum
                    tempPackageItem.Quantity = input[i].Packages[j].ShipItems[k].Quantity
                    tempPackageItem.RetailopsOrderItemID = input[i].Packages[j].ShipItems[k].ChannelOrderRefnum //is this correct?
                    tempPackageItem.RetailopsShipmentItemID = input[i].Packages[j].ShipItems[k].ID
                    tempPackageItem.Sku = input[i].Packages[j].ShipItems[k].SkuNum
                    tempPackageItem.Quantity = input[i].Packages[j].ShipItems[k].Quantity
                    tempPackage.PackageItems[k] = tempPackageItem
                }
                tempShipment.Packages[i] = tempPackage
            }
            shipmentArray[i] = tempShipment
    }
    return shipmentArray
}

func BuildUnshippedItems(input []ROPInputUnshippedItem) ([]UnshippedItem) {
    unshippedItemsArray := make([]UnshippedItem, len(input), (cap(input)+1)*2)
    for i := range input {
        var tempItem UnshippedItem
        tempItem.ChannelItemRefnum = input[i].ChannelItemRefnum
        // tempItem.EffectiveExtendedPrice = input.Data.Order.UnshippedItemsRef[i]. //TODO: what field do we map????
        // tempItem.EffectiveUnitPrice = input.Data.Order.UnshippedItemsRef[i].  //TODO: what field do we map????
        // tempItem.OrderedQuantity = input.Data.Order.UnshippedItemsRef[i]. //TODO: what field do we map????
        tempItem.Sku = input[i].SkuNum
        tempItem.UnshippedQuantity = input[i].Quantity
        unshippedItemsArray[i] = tempItem
    }
    return unshippedItemsArray
}

func BuildURI(baseURI string, action string) (string) {
    var buffer bytes.Buffer
    buffer.WriteString(baseURI)
    //test if baseURI ends with a slash and append action accordingly
    if strings.HasSuffix(baseURI, "/") {
        buffer.WriteString(action)
    }else{
        buffer.WriteString("/")
        buffer.WriteString(action)
    }
    return buffer.String()
}
