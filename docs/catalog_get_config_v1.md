## <a name="resource-catalog_get_config_v1">catalog_get_config</a>

Stability: `prototype`

catalog_get_config method RetailOPS webhook API version 1

### catalog_get_config catalog_get_config

Catalog get config method.

```
POST /orders
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"catalog_get_config"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "version": 1,
  "action": "catalog_get_config",
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
  "fanout": "relative"
}
```


