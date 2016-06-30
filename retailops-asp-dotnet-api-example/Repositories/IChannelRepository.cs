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
        bool checkToken(string token);

        ConfigResponse catalog_get_config();
        
        EventResponse catalog_push();
        
        EventResponse inventory_push();
        
        OrderPullResponse order_pull();
        
        EventResponse order_acknowledge();
        
        EventResponse order_update();
        
        EventResponseWithStatus order_cancel();
        
        EventResponse order_shipment_submit();
        
        EventResponseWithStatus order_complete();
        
        EventResponse order_settle_payment();
        
        EventResponse order_returned();
    }
}