## <a name="resource-catalog_push_v1">catalog_push</a>

Stability: `draft`

catalog_push method RetailOPS webhook API version 1

### catalog_push catalog_push

Catalog push method.

```
POST /catalog_push
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"catalog_push"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/catalog_push \
  -d '{
  "version": 1,
  "action": "catalog_push",
  "data": {
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


