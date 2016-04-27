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
        ChannelResponse catalog_get_config();
        
        ChannelResponse catalog_push();
        
        ChannelResponse inventory_push();
        
        ChannelResponse order_pull();
        
        ChannelResponse order_acknowledge();
        
        ChannelResponse order_update();
        
        ChannelResponse order_cancel();
        
        ChannelResponse order_shipment_submit();
        
        ChannelResponse order_complete();
        
        ChannelResponse order_settle_payment();
        
        ChannelResponse order_returned();
    }
}