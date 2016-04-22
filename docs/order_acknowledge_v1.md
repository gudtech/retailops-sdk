## <a name="resource-order_acknowledge_v1">order_acknowledge</a>

Stability: `draft`

order_acknowledge method RetailOPS webhook API

### order_acknowledge

Order acknowledge method.

```
POST /order_acknowledge
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"order_acknowledge"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:client_id** | *integer* | RetailOPS client id | `497` |
| **data:order:acks** | *array* | array of order IDs | `[496]` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/order_acknowledge \
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
