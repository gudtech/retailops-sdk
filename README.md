[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK

Using the RetailOps SDK, you can create, test, and certify custom integrations for use with RetailOps.

- [Channel Integrations](#channel-integrations)
- [Shipper Integrations (coming soon)](#shipper-integrations)
- [Tax Calculation Integration (coming soon)](#tax-calculation-integrations)

[![Build Status](https://travis-ci.org/gudTECH/retailops-sdk.svg?branch=web-hook-design)](https://travis-ci.org/gudTECH/retailops-sdk)

### Installation Instructions

The SDK includes

 * Example ASP.NET web server with all callbacks implemented
 * Verify Service command line tool for exercising callbacks

All parts of the SDK package are available for Windows, Linux, and OSX. Each release includes specific instructions for your platform.

[Read windows verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.windows.md)

[Read linux verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.linux.md)

[Read darwin (Mac OSX) verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.darwin.md)

## Interactive Documentation
- [RetailOps SDK](http://gudtech.github.io/retailops-sdk/v1/channel)
- [RetailOps Tax Calculation SDK](http://gudtech.github.io/retailops-sdk/v1/tax)

## Channel Integrations

A channel is a source of orders.
RetailOps, serves as the system-of-record for catalog/product data.
It pushes this catalog data out to "Channels" (Storefronts, marketplaces, etc) which in turn take customer orders for said products.
RetailOps then pulls this order information, fulfills said orders, and pushes status / edit information back to the channel in question.

### Channel Interactions:

- [catalog_get_config](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_catalog_get_config) - Configuration information about Catalog Push
- [catalog_push](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_catalog_push) - New product information and product updates
- [inventory_push](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_inventory_push) - Inventory updates
- [order_pull](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_pull_v1) - Fetch new orders from the channel which are ready
- [order_acknowledge](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_acknowledge) - Mark specific fetched orders has having been picked up
- [order_update](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_update) - Update channel order information to reflect order items which have been updated, added, or removed in RetailOps
- [order_cancel](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_cancel) - Mark an order as canceled in the channel
- [order_shipment_submit](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_shipment_submit) - Convey shipping status and tracking information to the channel
- [order_complete](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_complete) - Mark an order as fully completed in the channel
- [order_settle_payment](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_settle_payment) - Cause the channel to collect/capture payment, or otherwise verify payment has been collected
- [order_returned](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_returned) - Notify the channel that a return has been processed against the order

### Tax Calculation Integrations
- [catalog_get_config](http://gudtech.github.io/retailops-sdk/v1/tax#!/Tax/post_calculate_order_v1) - Tax Calculation for order

## Shipper Integrations

*Coming soon*
