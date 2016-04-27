/*
*   This example model defines a entities to populate the response JSON for testing. In a production application you 
*   which for this example contains all
*   entities related to our channel, (events, skus, payments, etc...) depending on the design of your 
*   actual web-hook implementation, these may not be neede or my be seperated into different classes altogether
*/
using System;

namespace dotnet_example_api.Models
{
    public class Channel
    {
        // public string Key {get; set;}
        // public string Name {get; set;}
        // public bool IsComplete {get; set;}
        
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
    public class Customer
    {
        public string first_name {get; set;}
        public string last_name {get; set;}
        public string email_address {get; set;}
        public string phone_number {get; set;}
        
    }
    
    public class Address
    {
        
    }
    
    public class ChannelPayment
    {
        
    }
    
    public class OrderItem
    {
        public int channel_refnum {get; set;}
        public int sku {get; set;}
        public double unit_tax {get; set;}
        public int quantity {get; set;}
        public string sku_title {get; set;}
        public double unit_price{get; set;}
    }
    
    public class Order
    {
        public double shipping_amt {get; set;}
        public string calc_mode {get; set;}
        public int channel_date_created {get; set;}
        public ChannelPayment[] payment {get; set;}
        public double tax_amt {get; set;}
        public Address bill_addr {get; set;}
        public string gift_message {get; set;}
        public Address ship_addr {get; set;}
        public int channel_refnum {get; set;}
        public Customer customer {get; set;}
        public double discount_amt {get; set;}
        public string shipcode {get; set;}
        public string ip_address {get; set;}
        public object attributes {get; set;}
        public OrderItem[] items {get; set;}
    }
    
    public class ConfigResponse
    {
        public string sku_fanout {get; set;}
    }
    
    public class ChannelResponse
    {
        public string sku_fanout {get; set;}
        public Event[] events {get; set;}
        public string ection {get; set;}
        public int version {get; set;}
        public string handle {get; set;}
        public int id {get; set;}
        public string concept {get; set;}
        public string sub_code {get; set;}
        public int status {get; set;}
        public int is_failure {get; set;}
        public string requestUrl {get; set;}
        public string code {get; set;}
        public string message {get; set;}
        public int next_page_state {get; set;}
        public int next_order_refnum {get; set;}
        public Order[] orders {get; set;}
        
    }

    //TODO: add clases as needed
}