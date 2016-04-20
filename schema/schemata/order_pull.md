## <a name="resource-channel">Channel</a>

Stability: `prototype`

channel properties

### Channel capture_channel_payments

Capture channel payments request

```
POST /payment
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | api action name | `"Spree.order.capture_channel_payments"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean? | `0` |
| **data:channel:params:email_return** | *integer* | boolean? | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean? | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean? | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean? | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:push_cancel** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean? | `0` |
| **data:order:channel_payment:authed** | *integer* |  | `0` |
| **data:order:channel_payment:available** | *integer* |  | `35` |
| **data:order:channel_payment:captured** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_external** | *integer* | boolean? | `1` |
| **data:order:channel_payment:channel_id** | *integer* | channel ID | `12` |
| **data:order:channel_payment:charged** | *integer* |  | `35` |
| **data:order:channel_payment:credited** | *integer* |  | `0` |
| **data:order:channel_payment:id** | *string* |  | `"6"` |
| **data:order:channel_payment:method** | *string* |  | `"channel"` |
| **data:order:channel_payment:module** | *string* |  | `"Channel"` |
| **data:order:channel_payment:settled** | *integer* |  | `35` |
| **data:order:channel_payment:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:channel_payment:unsettled** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_external** | *integer* |  | `0` |
| **data:order:channel_payment:voided** | *integer* |  | `0` |
| **data:order:channel_refnum** | *string* | channel reference number | `"496"` |
| **data:order:from_counterparty_rate** | *integer* |  | `1` |
| **data:order:grand_total** | *string* | order grandtotal | `"35"` |
| **data:order:id** | *string* | order ID | `"4897"` |
| **data:order:payment_series_id** | *string* | payment series ID | `"2572"` |
| **data:order:payment_status:authed** | *integer* |  | `0` |
| **data:order:payment_status:available** | *integer* |  | `35` |
| **data:order:payment_status:by_account** | *array* | payment status by account | `[{"success":1,"unsettled_external":0,"unsettled_deferred":0,"channel_id":12,"authed":0,"captures_are_external":1,"unsettled":0,"settled":35,"charged":35,"captured":0,"captures_are_deferred":0,"voided":0,"id":"6","method":"channel","credited":0,"module":"Channel","available":35}]` |
| **data:order:payment_status:captured** | *integer* |  | `0` |
| **data:order:payment_status:charged** | *integer* |  | `35` |
| **data:order:payment_status:credited** | *integer* |  | `0` |
| **data:order:payment_status:settled** | *integer* |  | `35` |
| **data:order:payment_status:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:payment_status:unsettled** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_deferred** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_external** | *integer* |  | `0` |
| **data:order:ship_service_name** | *string* | name of shipping service | `"Will Call"` |
| **data:order:shipments/id** | *string* | shipment ID | `"100000084"` |
| **data:order:shipments/packages** | *array* | array of packages included in this shipment | `[{"class_name":"Standard","carrier_code":"WILLCALL","carrier_name":"WillCall","ship_items":[null],"tracking_number":"ZX29827782929","mapped_shipcode":null,"date_shipped":"2016-04-08T21:13:11Z","carrier_class_code":"WILLCALL","weight":"1","id":"370","carrier_class_name":"WillCall Standard"}]` |
| **data:order:unshipped_items_ref** | *array* |  | `["496"]` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/payment \
  -d '{
  "headers": {
    "client_id": "497",
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "Spree.order.capture_channel_payments",
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
        "id": "6",
        "method": "channel",
        "credited": 0,
        "module": "Channel",
        "available": 35
      },
      "grand_total": "35",
      "unshipped_items_ref": [
        "496"
      ],
      "payment_series_id": "2572",
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": "496",
      "shipments": [
        {
          "id": "100000084",
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
              "weight": "1",
              "id": "370",
              "carrier_class_name": "WillCall Standard"
            }
          ]
        }
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
            "id": "6",
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
        "id": "66",
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


## <a name="resource-event">Event</a>

Stability: `prototype`

event structure returned in responses


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
| **channel_date_created** | *integer* | date created expressed unix time | `1460142547` |
| **channel_refnum** | *string* | channel reference number | `"496"` |
| **customer:email_address** | *string* | Customer email address | `"john@gudtech.com"` |
| **customer:first_name** | *string* | Firest name | `"John"` |
| **customer:last_name** | *string* | last name | `"Smith"` |
| **customer:phone_number** | *string* | customer phone | `"5555555555"` |
| **discount_amt** | *integer* | amount discounted | `0` |
| **gift_message** | *string* | gift message | `"Happy Birthday"` |
| **grand_total** | *string* | order grandtotal | `"35"` |
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

### Order order_fetch

Order fetch request.

```
POST /orders
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.orderpull.order_fetch"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:StoreID** | *string* | Store ID | `"yhst-18909142938879050075142"` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean? | `0` |
| **data:channel:params:email_return** | *integer* | boolean? | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean? | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean? | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean? | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:next_order_refnum** | *integer* | next order reference number | `496` |
| **data:channel:params:order_ack_status_id** | *string* | order acknowledgement status id | `"32"` |
| **data:channel:params:order_fulfilled_status_id** | *string* | order fulfilled status id | `"34"` |
| **data:channel:params:order_in_filfillment_status_id** | *string* | order in fulfillment status id | `"33"` |
| **data:channel:params:push_cancel** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean? | `0` |
| **data:client_id** | *integer* | RetailOPS client id | `1` |
| **data:max_page_size** | *integer* | maximum number of records to include in paged response | `50` |
| **data:order:channel_refnum** | *string* | channel reference number | `"496"` |
| **data:page_state** | *integer* |  | `null` |
| **data:single** | *integer* | requesting single order? (boolean?) | `0` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "headers": {
    "client_id": "497",
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "testco.orderpull.order_fetch",
  "data": {
    "single": 0,
    "order": {
      "channel_refnum": "496"
    },
    "client_id": 1,
    "channel": {
      "id": 21,
      "params": {
        "StoreID": "yhst-18909142938879050075142",
        "next_order_refnum": 496,
        "order_ack_status_id": "32",
        "order_fulfilled_status_id": "34",
        "order_in_filfillment_status_id": "33",
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

### Order order_ack

Order Acknowledgement request.

```
POST /orders/order_ack
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.orderpull.order_fetch"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:StoreID** | *string* | Store ID | `"yhst-18909142938879050075142"` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean? | `0` |
| **data:channel:params:email_return** | *integer* | boolean? | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean? | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean? | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean? | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:next_order_refnum** | *integer* | next order reference number | `496` |
| **data:channel:params:order_ack_status_id** | *string* | order acknowledgement status id | `"32"` |
| **data:channel:params:order_fulfilled_status_id** | *string* | order fulfilled status id | `"34"` |
| **data:channel:params:order_in_filfillment_status_id** | *string* | order in fulfillment status id | `"33"` |
| **data:channel:params:push_cancel** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean? | `0` |
| **data:client_id** | *integer* | RetailOPS client id | `1` |
| **data:order:acks** | *array* | array of order IDs | `["496"]` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders/order_ack \
  -d '{
  "headers": {
    "client_id": "497",
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "testco.orderpull.order_fetch",
  "data": {
    "order": {
      "acks": [
        "496"
      ]
    },
    "client_id": 1,
    "channel": {
      "id": 21,
      "params": {
        "StoreID": "yhst-18909142938879050075142",
        "next_order_refnum": 496,
        "order_ack_status_id": "32",
        "order_fulfilled_status_id": "34",
        "order_in_filfillment_status_id": "33",
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
        "id": "66",
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

### Order order_complete

Order Complete request.

```
POST /orders/complete
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.orderpull.order_fetch"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean? | `0` |
| **data:channel:params:email_return** | *integer* | boolean? | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean? | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean? | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean? | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:push_cancel** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean? | `0` |
| **data:order:channel_payment:authed** | *integer* |  | `0` |
| **data:order:channel_payment:available** | *integer* |  | `35` |
| **data:order:channel_payment:captured** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_external** | *integer* | boolean? | `1` |
| **data:order:channel_payment:channel_id** | *integer* | channel ID | `12` |
| **data:order:channel_payment:charged** | *integer* |  | `35` |
| **data:order:channel_payment:credited** | *integer* |  | `0` |
| **data:order:channel_payment:id** | *string* |  | `"6"` |
| **data:order:channel_payment:method** | *string* |  | `"channel"` |
| **data:order:channel_payment:module** | *string* |  | `"Channel"` |
| **data:order:channel_payment:settled** | *integer* |  | `35` |
| **data:order:channel_payment:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:channel_payment:unsettled** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_external** | *integer* |  | `0` |
| **data:order:channel_payment:voided** | *integer* |  | `0` |
| **data:order:channel_refnum** | *string* | channel reference number | `"496"` |
| **data:order:from_counterparty_rate** | *integer* |  | `1` |
| **data:order:grand_total** | *string* | order grandtotal | `"35"` |
| **data:order:id** | *string* | order ID | `"4897"` |
| **data:order:payment_series_id** | *string* | payment series ID | `"2572"` |
| **data:order:payment_status:authed** | *integer* |  | `0` |
| **data:order:payment_status:available** | *integer* |  | `35` |
| **data:order:payment_status:by_account** | *array* | payment status by account | `[{"success":1,"unsettled_external":0,"unsettled_deferred":0,"channel_id":12,"authed":0,"captures_are_external":1,"unsettled":0,"settled":35,"charged":35,"captured":0,"captures_are_deferred":0,"voided":0,"id":"6","method":"channel","credited":0,"module":"Channel","available":35}]` |
| **data:order:payment_status:captured** | *integer* |  | `0` |
| **data:order:payment_status:charged** | *integer* |  | `35` |
| **data:order:payment_status:credited** | *integer* |  | `0` |
| **data:order:payment_status:settled** | *integer* |  | `35` |
| **data:order:payment_status:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:payment_status:unsettled** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_deferred** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_external** | *integer* |  | `0` |
| **data:order:ship_service_name** | *string* | name of shipping service | `"Will Call"` |
| **data:order:shipments/id** | *string* | shipment ID | `"100000084"` |
| **data:order:shipments/packages** | *array* | array of packages included in this shipment | `[{"class_name":"Standard","carrier_code":"WILLCALL","carrier_name":"WillCall","ship_items":[null],"tracking_number":"ZX29827782929","mapped_shipcode":null,"date_shipped":"2016-04-08T21:13:11Z","carrier_class_code":"WILLCALL","weight":"1","id":"370","carrier_class_name":"WillCall Standard"}]` |
| **data:order:unshipped_items_ref** | *array* |  | `["496"]` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders/complete \
  -d '{
  "headers": {
    "client_id": "497",
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "testco.orderpull.order_fetch",
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
        "id": "6",
        "method": "channel",
        "credited": 0,
        "module": "Channel",
        "available": 35
      },
      "grand_total": "35",
      "unshipped_items_ref": [
        "496"
      ],
      "payment_series_id": "2572",
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": "496",
      "shipments": [
        {
          "id": "100000084",
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
              "weight": "1",
              "id": "370",
              "carrier_class_name": "WillCall Standard"
            }
          ]
        }
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
            "id": "6",
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
        "id": "66",
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

### Order order_cancel

Order Cancel request.

```
POST /orders/cancel
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.orderpull.order_fetch"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean? | `0` |
| **data:channel:params:email_return** | *integer* | boolean? | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean? | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean? | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean? | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:push_cancel** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean? | `0` |
| **data:order:channel_payment:authed** | *integer* |  | `0` |
| **data:order:channel_payment:available** | *integer* |  | `35` |
| **data:order:channel_payment:captured** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_external** | *integer* | boolean? | `1` |
| **data:order:channel_payment:channel_id** | *integer* | channel ID | `12` |
| **data:order:channel_payment:charged** | *integer* |  | `35` |
| **data:order:channel_payment:credited** | *integer* |  | `0` |
| **data:order:channel_payment:id** | *string* |  | `"6"` |
| **data:order:channel_payment:method** | *string* |  | `"channel"` |
| **data:order:channel_payment:module** | *string* |  | `"Channel"` |
| **data:order:channel_payment:settled** | *integer* |  | `35` |
| **data:order:channel_payment:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:channel_payment:unsettled** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_external** | *integer* |  | `0` |
| **data:order:channel_payment:voided** | *integer* |  | `0` |
| **data:order:channel_refnum** | *string* | channel reference number | `"496"` |
| **data:order:from_counterparty_rate** | *integer* |  | `1` |
| **data:order:grand_total** | *string* | order grandtotal | `"35"` |
| **data:order:id** | *string* | order ID | `"4897"` |
| **data:order:payment_series_id** | *string* | payment series ID | `"2572"` |
| **data:order:payment_status:authed** | *integer* |  | `0` |
| **data:order:payment_status:available** | *integer* |  | `35` |
| **data:order:payment_status:by_account** | *array* | payment status by account | `[{"success":1,"unsettled_external":0,"unsettled_deferred":0,"channel_id":12,"authed":0,"captures_are_external":1,"unsettled":0,"settled":35,"charged":35,"captured":0,"captures_are_deferred":0,"voided":0,"id":"6","method":"channel","credited":0,"module":"Channel","available":35}]` |
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
| **data:order:unshipped_items_ref** | *array* |  | `["496"]` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders/cancel \
  -d '{
  "headers": {
    "client_id": "497",
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "testco.orderpull.order_fetch",
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
        "id": "6",
        "method": "channel",
        "credited": 0,
        "module": "Channel",
        "available": 35
      },
      "grand_total": "35",
      "unshipped_items_ref": [
        "496"
      ],
      "payment_series_id": "2572",
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": "496",
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
            "id": "6",
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
        "id": "66",
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

### Order order_writeback

Order writeback request.

```
POST /orders/writeback
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.orderpull.order_fetch"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean? | `0` |
| **data:channel:params:email_return** | *integer* | boolean? | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean? | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean? | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean? | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:push_cancel** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean? | `0` |
| **data:line_items/apportioned_ship_amt** | *string* | ??? | `"5"` |
| **data:line_items/corr** | *string* | ???? | `"7397"` |
| **data:line_items/direct_ship_amt** | *string* | number of items direct-shipped(?) | `"5"` |
| **data:line_items/estimated_cost** | *string* | estimated cost of shipment(?) | `"0"` |
| **data:line_items/estimated_extended_cost** | *string* | ?? | `"0.00"` |
| **data:line_items/estimated_ship_date** | *integer* | estimated ship date in unix time | `1460151590` |
| **data:line_items/estimated_unit_cost** | *string* | estimated unit cost | `"0"` |
| **data:line_items/quantity** | *string* | quantity | `"2"` |
| **data:line_items/removed** | *integer* | TODO: check type, supplied example was 'NULL' | `0` |
| **data:line_items/sku** | *string* | sku number | `"132"` |
| **data:line_items/unit_price** | *string* | unit price | `"30"` |
| **data:order:channel_payment:authed** | *integer* |  | `0` |
| **data:order:channel_payment:available** | *integer* |  | `35` |
| **data:order:channel_payment:captured** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_external** | *integer* | boolean? | `1` |
| **data:order:channel_payment:channel_id** | *integer* | channel ID | `12` |
| **data:order:channel_payment:charged** | *integer* |  | `35` |
| **data:order:channel_payment:credited** | *integer* |  | `0` |
| **data:order:channel_payment:id** | *string* |  | `"6"` |
| **data:order:channel_payment:method** | *string* |  | `"channel"` |
| **data:order:channel_payment:module** | *string* |  | `"Channel"` |
| **data:order:channel_payment:settled** | *integer* |  | `35` |
| **data:order:channel_payment:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:channel_payment:unsettled** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_external** | *integer* |  | `0` |
| **data:order:channel_payment:voided** | *integer* |  | `0` |
| **data:order:channel_refnum** | *string* | channel reference number | `"496"` |
| **data:order:from_counterparty_rate** | *integer* |  | `1` |
| **data:order:grand_total** | *string* | order grandtotal | `"35"` |
| **data:order:id** | *string* | order ID | `"4897"` |
| **data:order:payment_series_id** | *string* | payment series ID | `"2572"` |
| **data:order:payment_status:authed** | *integer* |  | `0` |
| **data:order:payment_status:available** | *integer* |  | `35` |
| **data:order:payment_status:by_account** | *array* |  | `[{"success":1,"unsettled_external":0,"unsettled_deferred":0,"channel_id":12,"authed":0,"captures_are_external":1,"unsettled":0,"settled":35,"charged":35,"captured":0,"captures_are_deferred":0,"voided":0,"id":"6","method":"channel","credited":0,"module":"Channel","available":35}]` |
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
| **data:order:unshipped_items_ref** | *array* |  | `["496"]` |
| **data:order_info:direct_tax_amt** | *string* |  | `"0"` |
| **data:order_info:discount_amt** | *string* |  | `"0"` |
| **data:order_info:shipping_amt** | *string* |  | `"5"` |
| **data:order_info:tax_amt** | *string* |  | `"0"` |
| **data:rmas** | *array* |  | `[null]` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders/writeback \
  -d '{
  "headers": {
    "client_id": "497",
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "testco.orderpull.order_fetch",
  "data": {
    "rmas": [
      null
    ],
    "line_items": [
      {
        "apportioned_ship_amt": "5",
        "estimated_extended_cost": "0.00",
        "sku": "132",
        "quantity": "2",
        "estimated_ship_date": 1460151590,
        "direct_ship_amt": "5",
        "corr": "7397",
        "removed": 0,
        "estimated_cost": "0",
        "estimated_unit_cost": "0",
        "unit_price": "30"
      }
    ],
    "order_info": {
      "shipping_amt": "5",
      "discount_amt": "0",
      "tax_amt": "0",
      "direct_tax_amt": "0"
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
        "id": "6",
        "method": "channel",
        "credited": 0,
        "module": "Channel",
        "available": 35
      },
      "grand_total": "35",
      "unshipped_items_ref": [
        "496"
      ],
      "payment_series_id": "2572",
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": "496",
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
            "id": "6",
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
  "line_items": [
    {
      "corr": "7397",
      "ship_amt": "5",
      "shiptax_amt": 0,
      "unit_tax": 0,
      "quantity": 1,
      "discount_amt": 0,
      "refnum": "496"
    }
  ],
  "values": {
    "shipping_amt": "0.25",
    "tax_amt": "0.07",
    "discount_amt": 0,
    "tax_pct": 0
  }
}
```

### Order shipment_submit

Shipment Submit request.

```
POST /orders/shipment_submit
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | api action name | `"Aabaco.order.shipment_submit"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:StoreID** | *string* | Store ID | `"yhst-18909142938879050075142"` |
| **data:channel:params:next_order_refnum** | *integer* | next order reference number | `496` |
| **data:channel:params:order_ack_status_id** | *string* | order acknowledgement status id | `"32"` |
| **data:channel:params:order_fulfilled_status_id** | *string* | order fulfilled status id | `"34"` |
| **data:channel:params:order_in_filfillment_status_id** | *string* | order in fulfillment status id | `"33"` |
| **data:order:channel_payment:authed** | *integer* |  | `0` |
| **data:order:channel_payment:available** | *integer* |  | `35` |
| **data:order:channel_payment:captured** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:captures_are_external** | *integer* | boolean? | `1` |
| **data:order:channel_payment:channel_id** | *integer* | channel ID | `12` |
| **data:order:channel_payment:charged** | *integer* |  | `35` |
| **data:order:channel_payment:credited** | *integer* |  | `0` |
| **data:order:channel_payment:id** | *string* |  | `"6"` |
| **data:order:channel_payment:method** | *string* |  | `"channel"` |
| **data:order:channel_payment:module** | *string* |  | `"Channel"` |
| **data:order:channel_payment:settled** | *integer* |  | `35` |
| **data:order:channel_payment:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:channel_payment:unsettled** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_deferred** | *integer* |  | `0` |
| **data:order:channel_payment:unsettled_external** | *integer* |  | `0` |
| **data:order:channel_payment:voided** | *integer* |  | `0` |
| **data:order:channel_refnum** | *string* | channel reference number | `"496"` |
| **data:order:from_counterparty_rate** | *integer* |  | `1` |
| **data:order:grand_total** | *string* | order grandtotal | `"35"` |
| **data:order:id** | *string* | order ID | `"4897"` |
| **data:order:payment_series_id** | *string* | payment series ID | `"2572"` |
| **data:order:payment_status:authed** | *integer* |  | `0` |
| **data:order:payment_status:available** | *integer* |  | `35` |
| **data:order:payment_status:by_account** | *array* | payment status by account | `[{"success":1,"unsettled_external":0,"unsettled_deferred":0,"channel_id":12,"authed":0,"captures_are_external":1,"unsettled":0,"settled":35,"charged":35,"captured":0,"captures_are_deferred":0,"voided":0,"id":"6","method":"channel","credited":0,"module":"Channel","available":35}]` |
| **data:order:payment_status:captured** | *integer* |  | `0` |
| **data:order:payment_status:charged** | *integer* |  | `35` |
| **data:order:payment_status:credited** | *integer* |  | `0` |
| **data:order:payment_status:settled** | *integer* |  | `35` |
| **data:order:payment_status:success** | *integer* | boolean value indicating success of payment | `1` |
| **data:order:payment_status:unsettled** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_deferred** | *integer* |  | `0` |
| **data:order:payment_status:unsettled_external** | *integer* |  | `0` |
| **data:order:ship_service_name** | *string* | name of shipping service | `"Will Call"` |
| **data:order:shipments/id** | *string* | shipment ID | `"100000084"` |
| **data:order:shipments/packages** | *array* | array of packages included in this shipment | `[{"class_name":"Standard","carrier_code":"WILLCALL","carrier_name":"WillCall","ship_items":[null],"tracking_number":"ZX29827782929","mapped_shipcode":null,"date_shipped":"2016-04-08T21:13:11Z","carrier_class_code":"WILLCALL","weight":"1","id":"370","carrier_class_name":"WillCall Standard"}]` |
| **data:order:unshipped_items_ref** | *array* |  | `["496"]` |
| **data:shipment:id** | *string* | Shipment ID | `"10987"` |
| **data:shipment:packages** | *array* | Array of packages in this shipment | `[{"class_name":"Standard","carrier_code":"WILLCALL","carrier_name":"WillCall","ship_items":[null],"tracking_number":"ZX29827782929","mapped_shipcode":null,"date_shipped":"2016-04-08T21:13:11Z","carrier_class_code":"WILLCALL","weight":"1","id":"370","carrier_class_name":"WillCall Standard"}]` |
| **headers:client_id** | *integer* | RetailOPS client id | `"497"` |
| **headers:ticket** | *string* | RetailOPS authorization ticket | `"1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"` |
| **version** | *integer* | RetailOPS api action version | `1` |


#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders/shipment_submit \
  -d '{
  "headers": {
    "client_id": "497",
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "Aabaco.order.shipment_submit",
  "data": {
    "shipment": {
      "id": "10987",
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
          "weight": "1",
          "id": "370",
          "carrier_class_name": "WillCall Standard"
        }
      ]
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
        "id": "6",
        "method": "channel",
        "credited": 0,
        "module": "Channel",
        "available": 35
      },
      "grand_total": "35",
      "unshipped_items_ref": [
        "496"
      ],
      "payment_series_id": "2572",
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": "496",
      "shipments": [
        {
          "id": "100000084",
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
              "weight": "1",
              "id": "370",
              "carrier_class_name": "WillCall Standard"
            }
          ]
        }
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
            "id": "6",
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
      "params": {
        "StoreID": "yhst-18909142938879050075142",
        "next_order_refnum": 496,
        "order_ack_status_id": "32",
        "order_fulfilled_status_id": "34",
        "order_in_filfillment_status_id": "33"
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
        "id": "66",
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


## <a name="resource-channel_integration">order_pull</a>

Stability: `prototype`

order_pull method RetailOPS webhook API

### order_pull order_pull

Order fetch request.

```
POST /orders
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **action** | *string* | RetailOPS api action name | `"testco.orderpull.order_fetch"` |
| **data:channel:id** | *integer* |  | `21` |
| **data:channel:params:StoreID** | *string* | Store ID | `"yhst-18909142938879050075142"` |
| **data:channel:params:base_uri** | *string* | uri | `"http://172.16.4.130/magento1921"` |
| **data:channel:params:email_invoice** | *integer* | boolean? | `0` |
| **data:channel:params:email_return** | *integer* | boolean? | `0` |
| **data:channel:params:email_tracking** | *integer* | boolean? | `0` |
| **data:channel:params:express_configurable_super_links** | *integer* | boolean? | `0` |
| **data:channel:params:import_order_attrs** | *string* |  | `""` |
| **data:channel:params:inv_suspended_instock** | *integer* | boolean? | `0` |
| **data:channel:params:inv_suspended_mode** | *integer* |  | `null` |
| **data:channel:params:next_order_refnum** | *integer* | next order reference number | `496` |
| **data:channel:params:order_ack_status_id** | *string* | order acknowledgement status id | `"32"` |
| **data:channel:params:order_fulfilled_status_id** | *string* | order fulfilled status id | `"34"` |
| **data:channel:params:order_in_filfillment_status_id** | *string* | order in fulfillment status id | `"33"` |
| **data:channel:params:push_cancel** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_attributes** | *integer* | boolean? | `0` |
| **data:channel:params:unset_other_media** | *integer* | boolean? | `0` |
| **data:client_id** | *integer* | RetailOPS client id | `1` |
| **data:max_page_size** | *integer* | maximum number of records to include in paged response | `50` |
| **data:order:channel_refnum** | *string* | channel reference number | `"496"` |
| **data:page_state** | *integer* |  | `null` |
| **data:single** | *integer* | requesting single order? (boolean?) | `0` |
| **version** | *integer* | RetailOPS api action version | `1` |



#### Curl Example

```bash
$ curl -n -X POST https://yoursite.com/orders \
  -d '{
  "version": 1,
  "action": "testco.orderpull.order_fetch",
  "data": {
    "single": 0,
    "order": {
      "channel_refnum": "496"
    },
    "client_id": 1,
    "channel": {
      "id": 21,
      "params": {
        "StoreID": "yhst-18909142938879050075142",
        "next_order_refnum": 496,
        "order_ack_status_id": "32",
        "order_fulfilled_status_id": "34",
        "order_in_filfillment_status_id": "33",
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


