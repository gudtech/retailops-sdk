[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK

* DRAFT - THIS DOCUMENTATION IS PENDING FINALIZATION *

Using the RetailOps SDK, you can create, test, and certify custom integrations for use with RetailOps.

- [Channel Integrations](#channel-integrations)
- [Shipper Integrations (coming soon)](#shipper-integrations)
- [Payment Processor Integrations (coming soon)](#payment-processor-integrations)

[![Build Status](https://travis-ci.org/gudTECH/retailops-sdk.svg?branch=web-hook-design)](https://travis-ci.org/gudTECH/retailops-sdk)

## Interactive Documentation
[RetailOps SDK](http://senseijack.github.io/retailops-sdk)

## Channel Integrations

A channel is a source of orders.
RetailOps, serves as the system-of-record for catalog/product data.
It pushes this catalog data out to "Channels" (Storefronts, marketplaces, etc) which in turn take customer orders for said products.
RetailOps then pulls this order information, fulfills said orders, and pushes status / edit information back to the channel in question.

RetailOps is the initiator of all these requests.
In order to integrate, you must create web service endpoints which implement the following contracts for each desired interaction.
Not all interactions are required to be implemented. Any which are not implemented will be ignored by RetailOps, and instead handled according to the defaults.

### Channel Interactions:

- [catalog_get_config](http://gudtech.github.io/retailops-sdk/#!/default/post_catalog_get_config) - Configuration information about Catalog Push
- [catalog_push](http://gudtech.github.io/retailops-sdk/#!/default/post_catalog_push) - New product information and product updates
- [inventory_push](http://gudtech.github.io/retailops-sdk/#!/default/post_inventory_push) - Inventory updates
- [order_pull](http://gudtech.github.io/retailops-sdk/#!/default/post_order_pull) - Fetch new orders from the channel which are ready
- [order_acknowledge](http://gudtech.github.io/retailops-sdk/#!/default/post_order_acknowledge) - Mark specific fetched orders has having been picked up
- [order_update](http://gudtech.github.io/retailops-sdk/#!/default/post_order_update) - Update channel order information to reflect order items which have been updated, added, or removed in RetailOps
- [order_cancel](http://gudtech.github.io/retailops-sdk/#!/default/post_order_cancel) - Mark an order as canceled in the channel
- [order_shipment_submit](http://gudtech.github.io/retailops-sdk/#!/default/post_shipment_submit) - Convey shipping status and tracking information to the channel
- [order_complete](http://gudtech.github.io/retailops-sdk/#!/default/post_order_complete) - Mark an order as fully completed in the channel
- [order_settle_payment](http://gudtech.github.io/retailops-sdk/#!/default/post_order_settle_payment) - Cause the channel to collect/capture payment, or otherwise verify payment has been collected
- [order_returned](http://gudtech.github.io/retailops-sdk/#!/default/post_order_returned) - Notify the channel that a return has been processed against the order


## Shipper Integrations

*Coming soon*

## Payment Processor Integrations

*Coming soon*
