## <a name="resource-order_acknowledge_v1">order_acknowledge</a>

Stability: `prototype`

order_acknowledge method RetailOPS webhook API

### order_acknowledge order_acknowledge

Order acknowlege method.

```
POST /orders
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"order_acknowledge"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:StoreID** | *string* | Store ID | `"yhst-18909142938879050075142"` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean | `0` |
| **data:channel:params:email_return** | *integer* | boolean | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:next_order_refnum** | *integer* | next order reference number | `496` |
| **data:channel:params:order_ack_status_id** | *integer* | order acknowledgement status id | `32` |
| **data:channel:params:order_fulfilled_status_id** | *integer* | order fulfilled status id | `34` |
| **data:channel:params:order_in_filfillment_status_id** | *integer* | order in fulfillment status id | `33` |
| **data:channel:params:push_cancel** | *integer* | boolean | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean | `0` |
| **data:client_id** | *integer* | RetailOPS client id | `497` |
| **data:order:acks** | *array* | array of order IDs | `[496]` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "version": 1,
  "action": "order_acknowledge",
  "data": {
    "order": {
      "acks": [
        496
      ]
    },
    "client_id": 497,
    "channel": {
      "id": 21,
      "params": {
        "StoreID": "yhst-18909142938879050075142",
        "next_order_refnum": 496,
        "order_ack_status_id": 32,
        "order_fulfilled_status_id": 34,
        "order_in_filfillment_status_id": 33,
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


