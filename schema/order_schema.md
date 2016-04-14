## <a name="resource-order">Order</a>

Stability: `prototype`

order resource for RetailOPS webhook API

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **address:address1** | *string* | Address line one | `"123 Main St"` |
| **address:address2** | *string* | Address second line | `"suite 100"` |
| **address:city** | *string* | City | `"San Diego"` |
| **address:company** | *string* | Company name | `"gudTECH"` |
| **address:country_match** | *string* | Country | `"US United States"` |
| **address:first_name** | *string* | First Name | `"John"` |
| **address:last_name** | *string* | Last Name | `"Smith"` |
| **address:postal_code** | *string* | Postal Code | `"92101"` |
| **address:state_match** | *string* | state | `"CA"` |
| **amount** | *string* | payment amount | `"1.32"` |
| **calc_mode** | *string* | calculation mode | `"order"` |
| **channel:id** | *integer* |  | `21` |
| **channel:params:StoreID** | *string* | Store ID | `"yhst-18909142938879050075142"` |
| **channel:params:next_order_refnum** | *integer* | next order reference number | `496` |
| **channel:params:order_ack_status_id** | *string* | order acknowledgement status id | `"32"` |
| **channel:params:order_fulfilled_status_id** | *string* | order fulfilled status id | `"34"` |
| **channel:params:order_in_filfillment_status_id** | *string* | order in fulfillment status id | `"33"` |
| **channel_date_created** | *integer* | date created expressed unix time | `1460142547` |
| **channel_refnum** | *string* | channel reference number | `"496"` |
| **customer:email_address** | *string* | Customer email address | `"john@gudtech.com"` |
| **customer:first_name** | *string* | Firest name | `"John"` |
| **customer:last_name** | *string* | last name | `"Smith"` |
| **customer:phone_number** | *string* | customer phone | `"5555555555"` |
| **discount_amt** | *integer* | amount discounted | `0` |
| **gift_message** | *string* | gift message | `"Happy Birthday"` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **ip_address** | *string* | ip address used to place order | `"68.7.2.222"` |
| **max_page_size** | *integer* | maximum number of records to include in paged response | `50` |
| **page_state** | *integer* |  | `null` |
| **payment_method** | *string* | payment method | `"charge"` |
| **payment_type** | *string* | payment type | `"Visa"` |
| **quantity** | *integer* | quantity of sku in order | `1` |
| **shipcode** | *string* | shipping code | `"Ground (5-7 days)"` |
| **shipping_amt** | *string* |  | `"0.25"` |
| **sku** | *string* | sku number | `"299"` |
| **sku_title** | *string* | sku title | `"test"` |
| **tax_amt** | *string* | tax amount | `"0.07"` |
| **unit_price** | *string* | unit price for sku | `"1"` |
| **unit_tax** | *integer* | unit tax | `0` |
| **version** | *integer* | RetailOPS api action version | `1` |

### Order Fetch

Order fetch request.

```
GET /orders
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.orderpull.order_fetch"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:StoreID** | *string* | Store ID | `"yhst-18909142938879050075142"` |
| **data:channel:params:next_order_refnum** | *integer* | next order reference number | `496` |
| **data:channel:params:order_ack_status_id** | *string* | order acknowledgement status id | `"32"` |
| **data:channel:params:order_fulfilled_status_id** | *string* | order fulfilled status id | `"34"` |
| **data:channel:params:order_in_filfillment_status_id** | *string* | order in fulfillment status id | `"33"` |
| **data:client_id** | *integer* | RetailOPS client id | `1` |
| **data:max_page_size** | *integer* | maximum number of records to include in paged response | `50` |
| **data:order:channel_refnum** | *integer* | channel order reference number | `""` |
| **data:page_state** | *integer* |  | `null` |
| **data:single** | *integer* | requesting single order? (boolean?) | `0` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n https://yoursite.com/orders
 -G \
  -d headers[client_id]=497 \
  -d headers[ticket]=1%2C1%2C0%2C1456437061%2C315576060%2C111%2CWEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ \
  -d version=1 \
  -d action=testco.orderpull.order_fetch \
  -d data[single]=0 \
  -d data[order][channel_refnum]= \
  -d data[client_id]=1 \
  -d data[channel][id]=21 \
  -d data[channel][params][StoreID]=yhst-18909142938879050075142 \
  -d data[channel][params][next_order_refnum]=496 \
  -d data[channel][params][order_ack_status_id]=32 \
  -d data[channel][params][order_fulfilled_status_id]=34 \
  -d data[channel][params][order_in_filfillment_status_id]=33 \
  -d data[max_page_size]=50
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
      "shipping_amt": "0.25",
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


