using Microsoft.AspNet.Mvc;
using dotnet_example_api.Repositories;
using dotnet_example_api.Models;

namespace dotnet_example_api.Controllers
{
    [Route("api/[controller]/")]
    public class ChannelController : Controller
    {
        [FromServices]
        public IChannelRepository ChannelRepo { get; set; } //TODO: add Channel Model
        
        // The api actions defined below do not parse the incoming JSON included
        // in the request body. In a production application you would parse this JSON
        // and pass the data to your repository where the business logic resides.
        
        [HttpPost("get_channel_config")]
        public IActionResult get_channel_config([FromBody]string request)
        {
            ConfigResponse response = ChannelRepo.catalog_get_config();
            
            // the repository methods return a ChannelResponse object if 
            // they are successful, if not they return a null object and 
            // we return an HTTP not found error 
            if (response == null)
            {
                //in a production application you would define a custom 
                //erorr handler that would return a retailOPS formatted error 
                //with defined event structure
                return HttpNotFound();
            }
            return new ObjectResult(response);
        }
        
        [HttpPost("inventory_push")]
        public IActionResult inventory_push([FromBody]string request)
        {
            Event[] response = ChannelRepo.inventory_push();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("catalog_push")]
        public IActionResult catalog_push([FromBody]string request)
        {
            Event[] response = ChannelRepo.catalog_push();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }

        [HttpPost("order_pull")]
        public IActionResult order_pull([FromBody]string request)
        {
            ChannelResponse response = ChannelRepo.order_pull();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_acknowledge")]
        public IActionResult order_acknowledge([FromBody]string request)
        {
            Event[] response = ChannelRepo.order_acknowledge();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_update")]
        public IActionResult order_update([FromBody]string request)
        {
            Event[] response = ChannelRepo.order_update();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_cancel")]
        public IActionResult order_cancel([FromBody]string request)
        {
            Event[] response = ChannelRepo.order_cancel();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_shipment_submit")]
        public IActionResult order_shipment_submit([FromBody]string request)
        {
            Event[] response = ChannelRepo.order_shipment_submit();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_complete")]
        public IActionResult order_complete([FromBody]string request)
        {
            Event[] response = ChannelRepo.order_complete();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_settle_payment")]
        public IActionResult order_settle_payment([FromBody]string request)
        {
            Event[] response = ChannelRepo.order_settle_payment();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_returned")]
        public IActionResult order_returned([FromBody]string request)
        {
            Event[] response = ChannelRepo.order_returned();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
    }
}