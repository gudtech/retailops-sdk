package main

import (
  // "encoding/json"
  "github.com/gudTECH/scamp-go/scamp"
)

func main() {
  scamp.Initialize("/etc/GTSOA/soa.conf")
  scamp.Info.Printf("dialing")
  client,err := scamp.Dial("127.0.0.1:6000")
  if err != nil {
      scamp.Info.Printf("err: ", err)
  }
  msg := scamp.NewRequestMessage()
  // msg.SetAction("SDK.invpush_transmit")
  // msg.SetAction("SDK.capture_channel_payments")
  msg.SetAction("SDK.writeback")

  msg.SetRequestId(1 /*reqId*/)
  scamp.Info.Printf("reqId: %d", 1/* reqId */)

  var req = []byte (`{
  "headers": {
    "client_id": 1,
    "ticket": "1,1,0,1456437061,315576060,111,WEIhLHAyXHpjZA27OdENwAn1_fHx8fGA-ekng7lhkAvc27Uhnxgd4PZx4VnR_SJ-K85M_5dTAChTXgI3RsmGvfTbaOZ1_U-YJw3G0w1UVWFZ2EC83wjO6bmp91VZdR0tT_b2R1kK4qO1QTJrBk53ZyIuidsOa13lihh8VMgAvSDqCnTwxV2NVV7oN4v-h_tQtpUvklfbW1bnULR3bbaDvoOlb1CVQ_3BdNdo1MaAh-JxrRjf7MkzcHQYs3dN0GuaBZ1KBHvLdrLmGerNYv2p6AMC-fu8YeuukUU3Q6RL9AtF5AA6TPhfwfBM5r05B7QZiSEGySF65FCcfQFT_6lMxQ"
  },
  "version": 1,
  "action": "MagentoExtension.order.writeback",
  "data": {
    "rmas": [],
    "line_items": [
      {
        "apportioned_ship_amt": "5",
        "estimated_extended_cost": "0.00",
        "sku": "132",
        "quantity": "2",
        "estimated_ship_date": 1460151590,
        "direct_ship_amt": "5",
        "corr": "7397",
        "removed": null,
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
        "channel_id": 12,
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
      "grand_total": "65",
      "unshipped_items_ref": [
          {
              "item_giftwrap_amt": 0,
              "quantity": 1,
              "item_product_amt": "1",
              "item_recycling_amt": 0,
              "item_shipping_tax_amt": 0,
              "item_subtotal_amt": 1,
              "sku": "299",
              "item_tax_amt": "0.07",
              "id": "7392",
              "item_shipping_amt": "0.25",
              "channel_order_refnum": "496",
              "item_giftwrap_tax_amt": 0,
              "channel_id": "12",
              "item_credit_amt": 0,
              "reason": "ShortShip",
              "credit_item_refnum": "142628",
              "channel_item_refnum": "141416"
          }
      ],
      "payment_series_id": "2573",
      "from_counterparty_rate": 1,
      "ship_service_name": "Will Call",
      "channel_refnum": "100000085",
      "shipments": [
          {
    		"packages": [{
    			"class_name": "Standard",
    			"carrier_code": "WILLCALL",
    			"carrier_name": "WillCall",
    			"ship_items": [
                    {
        				"item_giftwrap_amt": 0,
        				"quantity": 1,
        				"item_product_amt": "1",
        				"item_recycling_amt": 0,
        				"channel_refnum": "171617",
        				"item_shipping_tax_amt": 0,
        				"item_subtotal_amt": 1,
        				"sku_num": "299",
        				"item_tax_amt": "0.07",
        				"id": "7392",
        				"item_shipping_amt": "0.25",
        				"channel_order_refnum": "496",
        				"item_giftwrap_tax_amt": 0
        			},
                    {
        				"item_giftwrap_amt": 2,
        				"quantity": 5,
        				"item_product_amt": "5",
        				"item_recycling_amt": 0,
        				"channel_refnum": "171615",
        				"item_shipping_tax_amt": 0,
        				"item_subtotal_amt": 1,
        				"sku_num": "299",
        				"item_tax_amt": "0.07",
        				"id": "7393",
        				"item_shipping_amt": "0.50",
        				"channel_order_refnum": "498",
        				"item_giftwrap_tax_amt": 0
        			}
                ],
    			"tracking_number": "ZXH171771919",
    			"mapped_shipcode": null,
    			"date_shipped": "2016-04-08T20:25:29Z",
    			"carrier_class_code": "WILLCALL",
    			"weight": "1",
    			"id": "369",
    			"carrier_class_name": "WillCall Standard"
    		}],
    		"id": "1002804"
    	  }
      ],
      "id": "4898",
      "payment_status": {
        "success": 1,
        "unsettled_external": 0,
        "unsettled_deferred": 0,
        "by_account": [
          {
            "channel_id": 12,
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
        "email_return": 0,
        "inv_suspended_instock": 0,
        "unset_other_media": 0,
        "import_order_attrs": "",
        "base_uri": "http://172.16.4.130/magento1921",
        "express_configurable_super_links": 0,
        "unset_other_attributes": 0,
        "push_cancel": 0,
        "order_writeback": 1,
        "inv_suspended_mode": null,
        "email_invoice": 0,
        "email_tracking": 0
      },
      "id": "12"
    }
  }
}`)

  msg.Write(req)
  respChan,err := client.Send(msg)
  if err != nil {
    scamp.Error.Printf("could not send message: `%s`", err)
  }
  scamp.Info.Printf("waiting")
  respMsg := <-respChan
  scamp.Info.Printf("response: %s", string(respMsg.Bytes()))
  // var resp common.VerifyResponse
  // err = json.NewDecoder(bytes.NewReader(respMsg.Bytes())).Decode(&resp)
  // if err != nil {
  //   scamp.Error.Printf("could not decode response: %s", err.Error())
  //   os.Exit(1)
  // }
  //
  // if resp.Status != "success" {
  //   scamp.Error.Printf("failed!")
  //   os.Exit(1)
  // }

  scamp.Info.Printf("success: response received")
  // }
}
