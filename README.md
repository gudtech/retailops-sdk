[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK

Using the RetailOps SDK, you can create, test, and certify custom integrations for use with RetailOps.

- [Channel Integrations](#channel-integrations)
- [Shipper Integrations](#shipping-integrations) (coming soon)
- [Tax Calculation Integrations](#tax-calculation-integrations) (coming soon)

[![Build Status](https://travis-ci.org/gudTECH/retailops-sdk.svg?branch=web-hook-design)](https://travis-ci.org/gudTECH/retailops-sdk)

### Installation Instructions

The SDK includes

 * Example ASP.NET web server with samples of all callbacks implemented
 * Verify Service command line tool for exercising callbacks

All parts of the SDK package are available for Windows, Linux, and OSX. Each release includes specific instructions for your platform.

[Read windows verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.windows.md)

[Read linux verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.linux.md)

[Read darwin (Mac OSX) verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.darwin.md)

## Interactive Documentation
- [RetailOps Channel SDK](http://gudtech.github.io/retailops-sdk/v1/channel)
- [RetailOps Tax Calculation SDK](http://gudtech.github.io/retailops-sdk/v1/tax) (coming soon)
- [RetailOps Shipment Processor SDK](http://gudtech.github.io/retailops-sdk/v1/shipping) (coming soon)

## Channel Integrations

A channel is a source of orders.
RetailOps, serves as the system-of-record for catalog/product data.
It pushes this catalog data out to "Channels" (Storefronts, marketplaces, etc) which in turn take customer orders for said products.
RetailOps then pulls this order information, fulfills said orders, and pushes status / edit information back to the channel in question.

### Channel Interactions:

- [catalog_get_config](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_catalog_get_config_v1) - Configuration information about Catalog Push (coming soon)
- [catalog_push](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_catalog_push_v1) - New product information and product updates (coming soon)
- [inventory_push](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_inventory_push_v1) - Inventory updates
- [order_pull](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_pull_v1) - Fetch new orders from the channel which are ready
- [order_acknowledge](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_acknowledge_v1) - Mark specific fetched orders has having been picked up
- [order_update](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_update_v1) - Update channel order information to reflect order items which have been updated, added, or removed in RetailOps
- [order_cancel](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_cancel_v1) - Mark an order as canceled in the channel
- [order_shipment_submit](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_shipment_submit_v1) - Convey shipping status and tracking information to the channel
- [order_complete](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_complete_v1) - Mark an order as fully completed in the channel
- [order_settle_payment](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_settle_payment_v1) - Cause the channel to collect/capture payment, or otherwise verify payment has been collected (coming soon)
- [order_returned](http://gudtech.github.io/retailops-sdk/v1/channel/#!/default/post_order_returned_v1) - Notify the channel that a return has been processed against the order

## Tax Calculation Integrations 

RetailOps supports integration with third party sales tax management services. RetailOps serves as the system-of-record for
customer and order data and provides necessary data to third party tax services for calculation of sales tax. 

### Tax Calculation Interactions (coming soon)
- [calculate_order](http://gudtech.github.io/retailops-sdk/v1/tax#!/Tax/post_calculate_order_v1) - Tax Calculation for order

## Shipping Integrations

RetailOps supports shipping orders via direct integrations with shipping processors and carriers. RetailOps serves as the system-of-record
for package/shipment data, and provides this data to shippers for the generation of shipping labels and tracking numbers.

### Shipping Interactions (coming soon)
- [get_labels](http://gudtech.github.io/retailops-sdk/v1/shipping/#!/Shipping_Labels/post_get_labels_v1) - Get shipping labels for ROP shipments
- [void_label](http://gudtech.github.io/retailops-sdk/v1/shipping/#!/Shipping_Labels/post_void_label_v1) - void shipping label for ROP shipment
- [get_rate](http://gudtech.github.io/retailops-sdk/v1/shipping/#!/Rate_and_Time/post_get_rate_v1) - Get rate for carrier class
- [get_rate_and_time](http://gudtech.github.io/retailops-sdk/v1/shipping/#!/Rate_and_Time/post_get_rate_and_time_v1) - Get rate and estimated delivery time for carrier class
- [get_time](http://gudtech.github.io/retailops-sdk/v1/shipping/#!/Rate_and_Time/post_get_time_v1) - Get estimated delivery time for carrier class
