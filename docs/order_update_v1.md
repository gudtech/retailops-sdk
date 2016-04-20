## <a name="resource-order_update_v1">order_update</a>

Stability: `prototype`

order_update method RetailOPS webhook API version 1

### order_update order_update

Order update method.

```
POST /orders
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"order_update"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean | `0` |
| **data:channel:params:email_return** | *integer* | boolean | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:push_cancel** | *integer* | boolean | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean | `0` |
| **data:line_items/apportioned_ship_amt** | *integer* | ??? | `5` |
| **data:line_items/corr** | *integer* | ???? | `7397` |
| **data:line_items/direct_ship_amt** | *integer* | number of items direct-shipped(?) | `5` |
| **data:line_items/estimated_cost** | *number* | estimated cost of shipment | `0.0` |
| **data:line_items/estimated_extended_cost** | *number* | ?? | `0.0` |
| **data:line_items/estimated_ship_date** | *integer* | estimated ship date in unix time | `1460151590` |
| **data:line_items/estimated_unit_cost** | *number* | estimated unit cost | `0.0` |
| **data:line_items/quantity** | *integer* | quantity | `2` |
| **data:line_items/removed** | *integer* |  | `0` |
| **data:line_items/sku** | *integer* | sku number | `132` |
| **data:line_items/unit_price** | *number* | unit price | `30.0` |
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
| **data:order:grand_total** | *number* | order grandtotal | `35.0` |
| **data:order:id** | *integer* | order ID | `4897` |
| **data:order:payment_series_id** | *integer* | payment series ID | `2572` |
| **data:order:payment_status:authed** | *integer* |  | `0` |
| **data:order:payment_status:available** | *integer* |  | `35` |
| **data:order:payment_status:by_account** | *array* |  | `[{"success":1,"unsettled_external":0,"unsettled_deferred":0,"channel_id":12,"authed":0,"captures_are_external":1,"unsettled":0,"settled":35,"charged":35,"captured":0,"captures_are_deferred":0,"voided":0,"id":6,"method":"channel","credited":0,"module":"Channel","available":35}]` |
| **data:order:payment_status:captured** | *integer* |  | `0` |
| **data:order:payment_status:charged** | *integer* |  | `35` |
| **data:order:payment_status:credited** | *integer* |  | `0` |
| **data:order:payment_status:settled** | *integer* |  | `35` |
| **data:order:payment_status:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:payment_status:unsettled** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_deferred** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_external** | *integer* |  | `0` |
| **data:order:ship_service_name** | *string* | name of shipping service | `"Will Call"` |
| **data:order:shipments** | *array* |  | `[null]` |
| **data:order:unshipped_items_ref** | *array* |  | `[496]` |
| **data:order_info:direct_tax_amt** | *number* |  | `0.0` |
| **data:order_info:discount_amt** | *number* |  | `0.0` |
| **data:order_info:shipping_amt** | *integer* |  | `5` |
| **data:order_info:tax_amt** | *number* |  | `0.0` |
| **data:rmas** | *array* |  | `[null]` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "version": 1,
  "action": "order_update",
  "data": {
    "rmas": [
      null
    ],
    "line_items": [
      {
        "apportioned_ship_amt": 5,
        "estimated_extended_cost": 0.0,
        "sku": 132,
        "quantity": 2,
        "estimated_ship_date": 1460151590,
        "direct_ship_amt": 5,
        "corr": 7397,
        "removed": 0,
        "estimated_cost": 0.0,
        "estimated_unit_cost": 0.0,
        "unit_price": 30.0
      }
    ],
    "order_info": {
      "shipping_amt": 5,
      "discount_amt": 0.0,
      "tax_amt": 0.0,
      "direct_tax_amt": 0.0
    },
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
      "grand_total": 35.0,
      "unshipped_items_ref": [
        496
      ],
      "payment_series_id": 2572,
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": 496,
      "shipments": [
        null
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
      "id": 21,
      "params": {
        "email_return": 0,
        "inv_suspended_instock": 0,
        "unset_other_media": 0,
        "import_order_attrs": "",
        "base_uri": "http://172.16.4.130/magento1921",
        "express_configurable_super_links": 0,
        "unset_other_attributes": 0,
        "push_cancel": 0,
        "inv_suspended_mode": null,
        "email_invoice": 0,
        "email_tracking": 0
      }
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
      "handle": "channel_catpush_fail",
      "secondary": {
        "id": 66,
        "concept": "sku"
      },
      "data": {
        "sub_code": "",
        "status": 42,
        "is_failure": 1,
        "request_url": "https://t14961.sandbox.mozu.com/api/commerce/catalog/admin/products/PP20?responseFields=",
        "data_items": [
          null
        ],
        "additonal": "[{\"name\":\"ParameterName\",\"value\":\"PackagWeight.Unit\"}]",
        "code": "MISSING_OR_INVALID_PARAMETER",
        "message": "Error: ident c1.channel-55-api_auth not found"
      }
    }
  ]
}
```


