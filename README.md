[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

## RetailOps SDK

The Retailops SDK provides the tools necessary to enable customers to create their own RetailOps SDK channel integrations. The SDK process is not entirely self-service, prior to beginning an SDK integration customers must contact RetailOps tech support so that we can add their developers to our Slack channel for SDK support. 

The SDK is different from the RetailOps API. When integrating with the API, you call the RetailOps servers.
When integrating with the SDK, the RetailOps servers call your servers. [Click here for more details about the RetailOps API](http://help.retailops.com/hc/en-us/articles/206283535-Getting-Started-with-the-RetailOps-API)

The SDK provides a detailed schema which will describe the interactions between our respective servers.
Herein, the term "SDK" will be used to refer to RetailOps calling your servers.

#### Why should RetailOps call my servers?

With most of the interactions covered by the SDK, RetailOps is the active system, pushing relevant data outward.
In a minority of the cases, RetailOps will poll for information which must be retrieved. Because a majority of the interactions are Pushed outward from RetailOps, this approach is generally simpler and more responsive.

On Rare occasions, it may be desirable to combine the RetailOps SDK with a few RetailOps API calls.
These scenarios are not presently discussed in this documentation, but may be discussed on the [RetailOps community forums](http://help.retailops.com/hc/en-us/community/topics), or via the RetailOps Partner Network.

Using the RetailOps SDK, you can create, test, and certify custom integrations for use with RetailOps.

- [Channel Integrations](#channel-integrations)

### Installation Instructions

The SDK includes

 * Verify Service command line tool for exercising callbacks 

The SDK command line tool is available for Windows, Linux, and OSX. Each release includes specific instructions for your platform.

[Read windows verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.windows.md)

[Read linux verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.linux.md)

[Read darwin (Mac OSX) verifier instructions here](https://github.com/gudTECH/retailops-sdk/blob/master/verify/README.darwin.md)

## Interactive Documentation
- [RetailOps Channel SDK](http://gudtech.github.io/retailops-sdk/)

## Channel Integrations

A channel is a source of orders.
RetailOps, serves as the system-of-record for catalog/product data.
It pushes this catalog data out to "Channels" (Storefronts, marketplaces, etc) which in turn take customer orders for said products.
RetailOps then pulls this order information, fulfills said orders, and pushes status / edit information back to the channel in question.

### Channel Interactions:

- [catalog_get_config](http://gudtech.github.io/retailops-sdk/#/default/post_catalog_get_config_v1) - Configuration information about Catalog Push 
- [catalog_push](http://gudtech.github.io/retailops-sdk/#/default/post_catalog_push_v1) - New product information and product updates
- [inventory_push](http://gudtech.github.io/retailops-sdk/#/default/post_inventory_push_v1) - Inventory updates
- [order_pull](http://gudtech.github.io/retailops-sdk/#/default/post_order_pull_v1) - Fetch new orders from the channel which are ready
- [order_acknowledge](http://gudtech.github.io/retailops-sdk/#/default/post_order_acknowledge_v1) - Mark specific fetched orders has having been picked up
- [order_update](http://gudtech.github.io/retailops-sdk/#/default/post_order_update_v1) - Update channel order information to reflect order items which have been updated, added, or removed in RetailOps
- [order_cancel](http://gudtech.github.io/retailops-sdk/#/default/post_order_cancel_v1) - Mark an order as canceled in the channel
- [order_shipment_submit](http://gudtech.github.io/retailops-sdk/#/default/post_order_shipment_submit_v1) - Convey shipping status and tracking information to the channel
- [order_complete](http://gudtech.github.io/retailops-sdk/#/default/post_order_complete_v1) - Mark an order as fully completed in the channel
- [order_settle_payment](http://gudtech.github.io/retailops-sdk/#/default/post_order_settle_payment_v1) - Cause the channel to collect/capture payment, or otherwise verify payment has been collected (coming soon)
- [order_returned](http://gudtech.github.io/retailops-sdk/#/default/post_order_returned_v1) - Notify the channel that a return has been processed against the order
