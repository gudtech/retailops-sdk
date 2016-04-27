/*
*   This example model defines a entities to populate the response JSON for testing. In a production application you 
*   which for this example contains all
*   entities related to our channel, (events, skus, payments, etc...) depending on the design of your 
*   actual web-hook implementation, these may not be neede or my be seperated into different classes altogether
*/
using System.Collections;
using System;

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
        public string status {get; set;}
        public string error_code {get; set;}
        public string error_message {get; set;}
        public Array diagnostic_data {get; set;}
        public String[] associations {get; set;}            
    }
    
    public class Sku
    {
        
    }
    
    public class ChannelPayment
    {
        
    }
    
    public class ChannelResponse
    {
        public string sku_fanout {get; set;}
        public Event[] events {get; set;}
        public string ection { get; set; }
        public int version { get; set; }
        public string handle { get; set; }
        public int id { get; set; }
        public string concept { get; set; }
        public string sub_code { get; set; }
        public int status { get; set; }
        public int is_failure { get; set; }
        public string requestUrl { get; set; }
        public string code { get; set; }
        public string message { get; set; }
    }

    //TODO: add clases as needed
}