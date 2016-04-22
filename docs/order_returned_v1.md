## <a name="resource-order_returned_v1">order_returned</a>

Stability: `draft`

order_returned method RetailOPS webhook API version 1

### order_returned order_returned

Order returned method.

```
POST /orders
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"order_returned"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:order:channel_payment:authed** | *integer* |  | `0` |
| **data:order:channel_payment:available** | *integer* |  | `35` |
| **data:order:channel_payment:captured** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_external** | *integer* | boolean? | `1` |
| **data:order:channel_payment:channel_id** | *integer* | channel ID | `12` |
| **data:order:channel_payment:charged** | *integer* |  | `35` |
| **data:order:channel_payment:credited** | *integer* |  | `0` |
| **data:order:channel_payment:id** | *integer* |  | `6` |
| **data:order:channel_payment:method** | *string* |  | `"channel"` |
| **data:order:channel_payment:module** | *string* |  | `"Channel"` |
| **data:order:channel_payment:settled** | *integer* |  | `35` |
| **data:order:channel_payment:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:channel_payment:unsettled** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_external** | *integer* |  | `0` |
| **data:order:channel_payment:voided** | *integer* |  | `0` |
| **data:order:channel_refnum** | *integer* | channel reference number for order | `496` |
| **data:order:from_counterparty_rate** | *integer* |  | `1` |
| **data:order:grand_total** | *number* | order grandtotal | `35` |
| **data:order:id** | *integer* | order ID | `4897` |
| **data:order:payment_series_id** | *integer* | payment series ID | `2572` |
| **data:order:payment_status:authed** | *integer* |  | `0` |
| **data:order:payment_status:available** | *integer* |  | `35` |
| **data:order:payment_status:by_account** | *array* | payment status by account | `[{"success":1,"unsettled_external":0,"unsettled_deferred":0,"channel_id":12,"authed":0,"captures_are_external":1,"unsettled":0,"settled":35,"charged":35,"captured":0,"captures_are_deferred":0,"voided":0,"id":6,"method":"channel","credited":0,"module":"Channel","available":35}]` |
| **data:order:payment_status:captured** | *integer* |  | `0` |
| **data:order:payment_status:charged** | *integer* |  | `35` |
| **data:order:payment_status:credited** | *integer* |  | `0` |
| **data:order:payment_status:settled** | *integer* |  | `35` |
| **data:order:payment_status:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:payment_status:unsettled** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_deferred** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_external** | *integer* |  | `0` |
| **data:order:ship_service_name** | *string* | name of shipping service | `"Will Call"` |
| **data:order:shipments/id** | *integer* | shipment ID | `100000084` |
| **data:order:shipments/packages** | *array* | array of packages included in this shipment | `[{"class_name":"Standard","carrier_code":"WILLCALL","carrier_name":"WillCall","ship_items":[null],"tracking_number":"ZX29827782929","mapped_shipcode":null,"date_shipped":"2016-04-08T21:13:11Z","carrier_class_code":"WILLCALL","weight":1,"id":370,"carrier_class_name":"WillCall Standard"}]` |
| **data:order:unshipped_items_ref** | *array* |  | `[496]` |
| **data:return:credit_items_ref** | *array* |  | `[{"sku":132,"item_shipping_tax_amt":30.0,"credit_item_refnum":"return_item 90","item_tax_amt":0.0,"channel_order_refnum":100000084,"item_shipping_amt":0.0,"item_restock_fee_amt":0.0,"channel_id":12,"item_giftwrap_amt":0,"channel_item_refnum":88,"quantity":1,"reason":"CustomerReturn","item_product_amt":30.0,"item_recycling_amt":0,"item_subtotal_amt":30.0,"item_credit_amt":30.0,"item_giftwrap_tax_amt":0.0}]` |
| **data:return:discount_amt** | *number* | amount of applied discount | `0.0` |
| **data:return:id** | *integer* | ID of return | `87` |
| **data:return:items/channel_refnum** | *integer* | channel reference number for order | `496` |
| **data:return:items/order_item_id** | *integer* | order item id | `7395` |
| **data:return:items/quantity** | *integer* | quantity of sku in order | `1` |
| **data:return:items/sku** | *integer* | sku number (id) | `53` |
| **data:return:product_amt** | *number* |  | `30.0` |
| **data:return:refund_action** | *string* | action name of return | `"refund"` |
| **data:return:refund_amt** | *number* | amount refunded | `30.0` |
| **data:return:rma_id** | *string* | ID of RMA | `"null"` |
| **data:return:shipping_amt** | *integer* | amount shipped | `0` |
| **data:return:subtotal_amt** | *number* |  | `30.0` |
| **data:return:tax_amt** | *number* | tax amount on returned items | `0.0` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "version": 1,
  "action": "order_returned",
  "data": {
    "order": {
      "channel_payment": {
        "success": 1,
        "unsettled_external": 0,
        "unsettled_deferred": 0,
        "channel_id": 12,
        "authed": 0,
        "captures_are_external": 1,
        "unsettled": 0,
        "settled": 35,
        "charged": 35,
        "captured": 0,
        "captures_are_deferred": 0,
        "voided": 0,
        "id": 6,
        "method": "channel",
        "credited": 0,
        "module": "Channel",
        "available": 35
      },
      "grand_total": 35,
      "unshipped_items_ref": [
        496
      ],
      "payment_series_id": 2572,
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": 496,
      "shipments": [
        {
          "id": 100000084,
          "packages": [
            {
              "class_name": "Standard",
              "carrier_code": "WILLCALL",
              "carrier_name": "WillCall",
              "ship_items": [
                null
              ],
              "tracking_number": "ZX29827782929",
              "mapped_shipcode": null,
              "date_shipped": "2016-04-08T21:13:11Z",
              "carrier_class_code": "WILLCALL",
              "weight": 1,
              "id": 370,
              "carrier_class_name": "WillCall Standard"
            }
          ]
        }
      ],
      "id": 4897,
      "payment_status": {
        "success": 1,
        "unsettled_external": 0,
        "unsettled_deferred": 0,
        "by_account": [
          {
            "success": 1,
            "unsettled_external": 0,
            "unsettled_deferred": 0,
            "channel_id": 12,
            "authed": 0,
            "captures_are_external": 1,
            "unsettled": 0,
            "settled": 35,
            "charged": 35,
            "captured": 0,
            "captures_are_deferred": 0,
            "voided": 0,
            "id": 6,
            "method": "channel",
            "credited": 0,
            "module": "Channel",
            "available": 35
          }
        ],
        "unsettled": 0,
        "settled": 35,
        "charged": 35,
        "captured": 0,
        "authed": 0,
        "credited": 0,
        "available": 35
      }
    },
    "channel": {
      "id": 21
    },
    "return": {
      "shipping_amt": 0,
      "subtotal_amt": 30.0,
      "product_amt": 30.0,
      "refund_amt": 30.0,
      "tax_amt": 0.0,
      "rma_id": "null",
      "refund_action": "refund",
      "discount_amt": 0.0,
      "credit_items_ref": [
        {
          "sku": 132,
          "item_shipping_tax_amt": 30.0,
          "credit_item_refnum": "return_item 90",
          "item_tax_amt": 0.0,
          "channel_order_refnum": 100000084,
          "item_shipping_amt": 0.0,
          "item_restock_fee_amt": 0.0,
          "channel_id": 12,
          "item_giftwrap_amt": 0,
          "channel_item_refnum": 88,
          "quantity": 1,
          "reason": "CustomerReturn",
          "item_product_amt": 30.0,
          "item_recycling_amt": 0,
          "item_subtotal_amt": 30.0,
          "item_credit_amt": 30.0,
          "item_giftwrap_tax_amt": 0.0
        }
      ],
      "id": 87,
      "items": [
        {
          "channel_refnum": 496,
          "order_item_id": 7395,
          "sku": 53,
          "quantity": 1
        }
      ]
    }
  }
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "events": [
    {
      "status": "error",
      "error_code": "ERR1234",
      "error_message": "Example error message",
      "diagnostic_data": [],
      "associations": [
        {
          "type": "sku",
          "identity": "S1234",
        }
      ],
    }
  ]
}
```
