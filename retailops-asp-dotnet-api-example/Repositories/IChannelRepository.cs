/*
 *  This interface defines CRUD operations. Using an interface decouples the repository class from the controller(s) that use it. 
 *  Interfaces also makes it easier to unit test controllers. Unit tests should inject a stub version of the interface, 
 *  this way your tests target the controller logic instead of the data access layer.
 */
using dotnet_example_api.Models;

namespace dotnet_example_api.Repositories
{
    public interface IChannelRepository
    {
        ConfigResponse catalog_get_config();
        
        Event[] catalog_push();
        
        Event[] inventory_push();
        
        ChannelResponse order_pull();
        
        Event[] order_acknowledge();
        
        Event[] order_update();
        
        Event[] order_cancel();
        
        Event[] order_shipment_submit();
        
        Event[] order_complete();
        
        Event[] order_settle_payment();
        
        Event[] order_returned();
    }
}