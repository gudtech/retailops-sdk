## <a name="resource-order_cancel_v1">order_cancel</a>

Stability: `draft`

order_cancel method RetailOPS webhook API version 1

### order_cancel order_cancel

Order cancel method.

```
POST /order_cancel
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"order_cancel"` |
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
| **data:order:id** | *string* | order ID | `"4897"` |
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
| **data:order:shipments** | *array* |  | `[null]` |
| **data:order:unshipped_items_ref** | *array* |  | `[496]` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/order_cancel \
  -d '{
  "version": 1,
  "action": "order_cancel",
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
        null
      ],
      "id": "4897",
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
      "diagnostic_data": [
        null
      ],
      "associations": [
        {
          "type": "sku",
          "identity": "S1234"
        }
      ]
    }
  ]
}
```


