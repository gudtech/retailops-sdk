using System;
using System.Collections.Generic;
using dotnet_example_api.Models;

namespace dotnet_example_api.Repositories
{
    public class ChannelRepository : IChannelRepository
    {
        public ChannelRepository()
        {
            // TODO: define constructor
            //  i.e. Add(new TodoItem { Name = "Item1" });
        }
        
        // This method returns a generic "canned" response, 
        // which consitss of a single example event array, for the 
        // purposes of demonstration and testing
        private Event[] _GetStandardResponse()
        {
            List<Event> responseEvents = new List<Event>();
            
            Event responseEvent = new Event(){
                status          = "error",
                error_code      = "ERR1234",
                error_message   = "Example error message",
                diagnostic_data = new Array[0],
                associations    = new string[2]{" \"type\": \"sku\" ", "\"identity\":\"S1234\""}
            };
            
            responseEvents.Add(responseEvent);
            
            return responseEvents.ToArray();; 
        }

        public ConfigResponse catalog_get_config()
        {
            ConfigResponse response = new ConfigResponse();
            response.sku_fanout = "all_skus_for_product";
            return response;
        }
        
        public Event[] catalog_push()
        {   
            //return a canned response
            return _GetStandardResponse();  
        } 
        
        public Event[] inventory_push()
        {
           //return a canned response            
           return _GetStandardResponse();
        } 
          
        public ChannelResponse order_pull()
        {
           //TODO return detailed response
           //
           //
           //            
           return new ChannelResponse();
        }
        
        public Event[] order_acknowledge()
        {
           //return a canned response            
           return _GetStandardResponse();
        }
        
        public Event[] order_update()
        {
           //return a canned response            
           return _GetStandardResponse();
        } 
        
        public Event[] order_cancel()
        {
           //return a canned response            
           return _GetStandardResponse();
        }
        
        public Event[] order_shipment_submit()
        {
           //return a canned response            
           return _GetStandardResponse();
        }   
        
        public Event[] order_complete()
        {
           //return a canned response            
           return _GetStandardResponse();
        } 
        
        public Event[] order_settle_payment()
        {
           //return a canned response            
           return _GetStandardResponse();
        }
        
        public Event[] order_returned()
        {
           //return a canned response            
           return _GetStandardResponse();
        }    
          
    }
}