using Microsoft.AspNet.Mvc;
using dotnet_example_api.Repositories;
using dotnet_example_api.Models;
using Microsoft.Extensions.Logging;
using System;

namespace dotnet_example_api.Controllers
{
    [Route("api/[controller]/")]
    public class ChannelController : Controller
    {
        [FromServices]
        public IChannelRepository ChannelRepo { get; set; } 
        
        private readonly ILogger<ChannelController> _logger;
        
        public ChannelController(ILogger<ChannelController> logger)
        {
            _logger = logger;
        }


        // The api actions defined below do not parse the incoming JSON included
        // in the request body. In a production application you would parse this JSON
        // and pass the data to your repository where the business logic resides.
        
        [HttpPost("catalog_get_config_v1")]
        public IActionResult catalog_get_config_v1([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            // Console.WriteLine(requestJSON);
            Console.WriteLine("\n");

            if( string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            ConfigResponse response = ChannelRepo.catalog_get_config();
            
            Console.WriteLine("\response: {0}", response);
            
            // the repository methods return a response object if 
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
        
        [HttpPost("inventory_push_v1")]
        public IActionResult inventory_push([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponse response = ChannelRepo.inventory_push();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("catalog_push_v1")]
        public IActionResult catalog_push([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponse response = ChannelRepo.catalog_push();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }

        [HttpPost("order_pull_v1")]
        public IActionResult order_pull([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            OrderPullResponse response = ChannelRepo.order_pull();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_acknowledge_v1")]
        public IActionResult order_acknowledge([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponse response = ChannelRepo.order_acknowledge();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_update_v1")]
        public IActionResult order_update([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponse response = ChannelRepo.order_update();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_cancel_v1")]
        public IActionResult order_cancel([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponseWithStatus response = ChannelRepo.order_cancel();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_shipment_submit_v1")]
        public IActionResult order_shipment_submit([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponse response = ChannelRepo.order_shipment_submit();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_complete_v1")]
        public IActionResult order_complete([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponseWithStatus response = ChannelRepo.order_complete();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_settle_payment_v1")]
        public IActionResult order_settle_payment([FromBody] PostJSON requestJSON)
        {
            Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }

            EventResponse response = ChannelRepo.order_settle_payment();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
        
        [HttpPost("order_returned_v1")]
        public IActionResult order_returned([FromBody] PostJSON requestJSON)
        {
             Console.WriteLine("\nRequestJSON: {0}", requestJSON.integration_auth_token);
            Console.WriteLine("\n");

            if (string.IsNullOrEmpty(requestJSON.integration_auth_token) || requestJSON.integration_auth_token != "RETAILOPS_SDK" )
            {
                return this.HttpUnauthorized();
            }
            
            EventResponse response = ChannelRepo.order_returned();

            if (response == null)
            {
                return HttpNotFound();
            }
            
            return new ObjectResult(response);
        }
    }
}