/*
*   This example model defines a channel entity, which for this example contains all
*   entities related to our channel, (events, skus, payments, etc...) depending on the design of your 
*   actual web-hook implementation, these may not be neede or my be seperated into different classes altogether
*/

namespace dotnet_example_api.Models
{
    public class Channel
    {
        // public string Key { get; set; }
        // public string Name { get; set; }
        // public bool IsComplete { get; set; }
        
        //TODO: define channel model properties
    }
    
    public class Event
    {
        
    }
    
    public class Sku
    {
        
    }
    
    public class ChannelPayment
    {
        
    }
    
    public class ChannelResponse
    {
        public string Fanout {get; set;}
        public Event Event {get; set;}
    }
    
    //TODO: add clases as needed
}