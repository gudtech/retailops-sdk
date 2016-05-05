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
        private EventResponse _GetStandardResponse()
        {
            List<Event> responseEvents = new List<Event>();
            
            Association assoc = new Association(){
                identifier_type = "order_id",
                identifier      = "S1234"
            };
            
            Event responseEvent = new Event(){
                event_type      = "warning",
                code            = "1234",
                message         = "Example warning message",
                diagnostic_data = new Array[0],
                associations    = new []{assoc}
            };
            
            responseEvents.Add(responseEvent);
            
            EventResponse events = new EventResponse(){
                events = responseEvents.ToArray()
            };
            
            return events; 
        }

        public ConfigResponse catalog_get_config()
        {
            ConfigResponse response = new ConfigResponse();
            response.sku_fanout = "related_skus";
            return response;
        }
        
        public EventResponse catalog_push()
        {   
            //return a canned response
            return _GetStandardResponse();  
        } 
        
        public EventResponse inventory_push()
        {
           //return a canned response            
           return _GetStandardResponse();
        } 
          
        public OrderPullResponse order_pull()
        {
           List<Order> orders = new List<Order>();
           List<ChannelPayment> payment = new List<ChannelPayment>();
           List<OrderItem> order_items = new List<OrderItem>();
           
           OrderItem orderItem = new OrderItem() {
               channel_refnum   = 496,
               sku              = 299,
               unit_tax         = 0,
               quantity         = 1,
               sku_title        = "test product",
               unit_price       = 1.00
           };
           
           order_items.Add(orderItem);
           
           Address address = new Address(){
               first_name       = "John",
               last_name        = "Smith",
               company          = "gudTECH",
               address1         = "600 B Street, Suite 2120",
               address2         = "",
               city             = "San Diego",
               state_match      = "CA",
               country_match    = "USA",
               postal_code      = "92101"
           };
           
           ChannelPayment newPayment = new ChannelPayment(){
             amount             = 1.32,
             type               = "charge",
             payment_params     = new {
                 channel_refnum = 496,
                 payment_type   = "Visa"
             }  
           };
           
           payment.Add(newPayment);
           
           Customer customer = new Customer(){
               first_name       = "John",
               last_name        = "Smith",
               email_address    = "john.smith@gmail.com"
           };
           
           Order newOrder = new Order(){
             shipping_amt           = 0.25,
             calc_mode              = "order",
             channel_date_created   = 1460142547,
             payment                = payment.ToArray(),
             tax_amt                = 0.07,
             bill_addr              = address,
             ship_addr              = address,
             gift_message           = "Happy Birthday",
             channel_refnum         = 496,
             customer               = customer,
             discount_amt           = 0,
             shipcode               = "Ground (5-7 days)",
             ip_address             = "192.168.1.187",
             attributes             = new string[0],
             items                  = order_items.ToArray()
           }; 
           
           orders.Add(newOrder);
           
           OrderPullResponse response = new OrderPullResponse(){
               next_page_state = 0,
               next_order_refnum = 496,
               orders = orders.ToArray()
           };
                      
           return response;
        }
        
        public EventResponse order_acknowledge()
        {
           //return a canned response            
           return _GetStandardResponse();
        }
        
        public EventResponse order_update()
        {
           //return a canned response            
           return _GetStandardResponse();
        } 
        
        public EventResponse order_cancel()
        {
           //return a canned response            
           return _GetStandardResponse();
        }
        
        public EventResponse order_shipment_submit()
        {
           //return a canned response            
           return _GetStandardResponse();
        }   
        
        public EventResponse order_complete()
        {
           //return a canned response            
           return _GetStandardResponse();
        } 
        
        public EventResponse order_settle_payment()
        {
           //return a canned response            
           return _GetStandardResponse();
        }
        
        public EventResponse order_returned()
        {
           //return a canned response            
           return _GetStandardResponse();
        }    
          
    }
}