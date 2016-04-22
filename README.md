[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK

* DRAFT - THIS DOCUMENTATION IS PENDING FINALIZATION *

Using the RetailOps SDK, you can create, test, and certify custom integrations for use with RetailOps.

- [Channel Integrations](#channel-integrations)
- [Shipper Integrations (coming soon)](#shipper-integrations)
- [Payment Processor Integrations (coming soon)](#payment-processor-integrations)

## Channel Integrations

A channel is a source of orders.
RetailOps, serves as the system-of-record for catalog/product data.
It pushes this catalog data out to "Channels" (Storefronts, marketplaces, etc) which in turn take customer orders for said products.
RetailOps then pulls this order information, fulfills said orders, and pushes status / edit information back to the channel in question.

RetailOps is the initiator of all these requests.
In order to integrate, you must create web service endpoints which implement the following contracts for each desired interaction.
Not all interactions are required to be implemented. Any which are not implemented will be ignored by RetailOps, and instead handled according to the defaults.

### Channel Interactions:

- [catalog_get_config](docs/catalog_get_config_v1.md) - Configuration information about Catalog Push
- [catalog_push](docs/catalog_push_v1.md) - New product information and product updates
- [inventory_push](docs/inventory_push_v1.md) - Inventory updates
- [order_pull](docs/order_pull_v1.md) - Fetch new orders from the channel which are ready
- [order_acknowledge](docs/order_acknowledge_v1.md) - Mark specific fetched orders has having been picked up
- [order_update](docs/order_update_v1.md) - Update channel order information to reflect order items which have been updated, added, or removed in RetailOps
- [order_cancel](docs/order_cancel_v1.md) - Mark an order as canceled in the channel
- [order_shipment_submit](docs/order_shipment_submit_v1.md) - Convey shipping status and tracking information to the channel
- [order_complete](docs/order_complete_v1.md) - Mark an order as fully completed in the channel
- [order_settle_payment](docs/order_settle_payment_v1.md) - Cause the channel to collect/capture payment, or otherwise verify payment has been collected
- [order_returned](docs/order_returned_v1.md) - Notify the channel that a return has been processed against the order


## Shipper Integrations

*Coming soon*

## Payment Processor Integrations

*Coming soon*
