using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
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
        

        [HttpGet("get_channel_config")]
        public IActionResult get_channel_config([FromBody]string request)
        {
            ChannelResponse response = ChannelRepo.catalog_get_config(1);
            if (response == null)
            {
                return HttpNotFound();//TODO may need to define custome type to return exepected event structure
            }
            return new ObjectResult(response);
        }
        
        [HttpPost("inventory_push")]
        public IActionResult inventory_push([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement inventory_push controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpGet("catalog_push")]
        public IActionResult catalog_push([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement catalog_push controller method\" }";

            return new ObjectResult(todo);
        }

        [HttpPost("order_pull")]
        public IActionResult order_pull([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_pull controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpPost("order_acknowledge")]
        public IActionResult order_acknowledge([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_acknowledge controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpPost("order_update")]
        public IActionResult order_update([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_update controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpPost("order_cancel")]
        public IActionResult order_cancel([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_cancel controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpPost("order_shipment_submit")]
        public IActionResult order_shipment_submit([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_shipment_submit controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpPost("order_complete")]
        public IActionResult order_complete([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_complete controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpPost("order_settle_payment")]
        public IActionResult order_settle_payment([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_settle_payment controller method\" }";

            return new ObjectResult(todo);
        }
        
        [HttpPost("order_returned")]
        public IActionResult order_returned([FromBody]string request)
        {
            string todo = "{ \"TODO\": \"Implement order_update controller method\" }";

            return new ObjectResult(todo);
        }
    }
}