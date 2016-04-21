## <a name="resource-inventory_push_v1">inventory_push</a>

Stability: `prototype`

inventory_push method RetailOPS webhook API version 1

### inventory_push inventory_push

Inventory push method.

```
POST /orders
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"inventory_push"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:appKey** | *string* | application key | `"acme_co.retailops.0.1.0.Release"` |
| **data:channel:params:breakdown_inventory** | *integer* |  | `0` |
| **data:channel:params:tenant** | *integer* | tenant ID | `15394` |
| **data:client_id** | *integer* | client ID | `1` |
| **data:inventory:data:concept** | *string* |  | `"sku"` |
| **data:inventory:data:id** | *integer* |  | `21` |
| **data:inventory:data:qty_available** | *integer* | quantity available | `24` |
| **data:inventory:data:qty_breakdown/est_ship** | *string* | estimated ship date | `"2016-04-22T11:02:47-07:00"` |
| **data:inventory:data:qty_breakdown/reserving_orders** | *array* | order ids | `[null]` |
| **data:inventory:data:qty_breakdown/sellable** | *integer* | number of units available to sale | `5` |
| **data:inventory:data:qty_breakdown/unclaimed** | *integer* |  | `0` |
| **data:inventory:data:qty_breakdown/vendor** | *string* | vendor name | `"002_Acme Corp"` |
| **data:inventory:data:sku** | *string* | sku number (id) | `"53"` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "version": 1,
  "action": "inventory_push",
  "data": {
    "inventory": {
      "data": {
        "sku": "53",
        "id": 21,
        "concept": "sku",
        "qty_breakdown": [
          {
            "est_ship": "2016-04-22T11:02:47-07:00",
            "reserving_orders": [
              null
            ],
            "sellable": 5,
            "unclaimed": 0,
            "vendor": "002_Acme Corp"
          }
        ],
        "qty_available": 24
      }
    },
    "client_id": 1,
    "channel": {
      "params": {
        "breakdown_inventory": 0,
        "tenant": 15394,
        "appKey": "acme_co.retailops.0.1.0.Release"
      },
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


