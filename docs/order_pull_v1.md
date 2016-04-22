## <a name="resource-order_pull">order_pull</a>

Stability: `draft`

order_pull method RetailOPS webhook API, version 1

### order_pull

Order fetch request.

```
POST /orders
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.order_pull"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:client_id** | *integer* | RetailOPS client id | `"497"` |
| **data:max_page_size** | *integer* | maximum number of records to include in paged response | `50` |
| **data:order:channel_refnum** | *string* | channel reference number for order | `"496"` |
| **data:page_state** | *integer* |  | `null` |
| **data:single** | *integer* | requesting single order? | `0` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "version": 1,
  "action": "testco.order_pull",
  "data": {
    "single": 0,
    "order": {
      "channel_refnum": "496"
    },
    "client_id": "497",
    "channel": {
      "id": 21
    },
    "page_state": null,
    "max_page_size": 50
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
  "next_page_state": 0,
  "next_order_refnum": 496,
  "orders": [
    {
      "shipping_amt": 0.25,
      "calc_mode": "order",
      "channel_date_created": 1460142547,
      "payment": [
        {
          "params": {
            "channel_refnum": "496",
            "payment_type": "Visa"
          },
          "amount": "1.32",
          "type": "charge"
        }
      ],
      "tax_amt": "0.07",
      "bill_addr": {
        "state_match": "CA",
        "country_match": "US United States",
        "last_name": "Smith",
        "address2": "suite 100",
        "city": "San Diego",
        "postal_code": "92101",
        "address1": "123 Main St",
        "company": "gudTECH",
        "first_name": "John"
      },
      "gift_message": "Happy Birthday",
      "ship_addr": {
        "state_match": "CA",
        "country_match": "US United States",
        "last_name": "Smith",
        "address2": "suite 100",
        "city": "San Diego",
        "postal_code": "92101",
        "address1": "123 Main St",
        "company": "gudTECH",
        "first_name": "John"
      },
      "channel_refnum": "496",
      "customer": {
        "email_address": "john@gudtech.com",
        "phone_number": "5555555555",
        "first_name": "John",
        "last_name": "Smith"
      },
      "discount_amt": 0,
      "shipcode": "Ground (5-7 days)",
      "ip_address": "68.7.2.222",
      "attributes": {
      },
      "items": [
        {
          "channel_refnum": "496",
          "sku": "299",
          "unit_tax": 0,
          "quantity": 1,
          "sku_title": "test",
          "unit_price": "1"
        }
      ]
    }
  ]
}
```
