/*
 *  This interface defines CRUD operations. Using an interface decouples the repository class from the controller(s) that use it. 
 *  Interfaces also makes it easier to unit test controllers. Unit tests should inject a stub version of the interface, 
 *  this way your tests target the controller logic instead of the data access layer.
 */
using System.Collections.Generic;
using dotnet_example_api.Models;

namespace dotnet_example_api.Repositories
{
    public interface IChannelRepository
    {
        // TODO: define the following api repository methods
        ChannelResponse catalog_get_config(int id);
        //  catalog_push 
        //  inventory_push 
        //  order_pull
        //  order_acknowledge
        //  order_update
        //  order_cancel
        //  order_shipment_submit
        //  order_complete
        //  order_settle_payment
        //  order_returned
    }
}