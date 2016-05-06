/*
*   This example model defines a entities to populate the response JSON for testing. In a production application you 
*   which for this example contains all
*   entities related to our channel, (events, skus, payments, etc...) depending on the design of your 
*   actual web-hook implementation, these may not be neede or my be seperated into different classes altogether
*/
using System;

namespace dotnet_example_api.Models
{ 
    public class Event
    {
        public string event_type {get; set;}
        public string code {get; set;}
        public string message {get; set;}
        public string diagnostic_data {get; set;}
        public Association[] associations {get; set;}            
    }
    
    public class EventResponse
    {
        public Event[] events {get; set;}
    }
    
    public class EventResponseWithStatus
    {
        public string status {get; set;}
        public Event[] events {get; set;}
    }
    
    public class Association
    {
        public string identifier_type {get; set;}
        public string identifier {get; set;}
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
        public string first_name {get; set;}
        public string last_name {get; set;}
        public string company {get; set;}
        public string address1 {get; set;}
        public string address2 {get; set;}
        public string city {get; set;}
        public string state_match {get; set;}
        public string country_match {get; set;}
        public string postal_code {get; set;}
        
    }
    
    public class ChannelPayment
    {
        public double amount {get; set;}
        public string type {get; set;}
        public string channel_refnum {get; set;}
        public string payment_type {get; set;}
        public object payment_params {get; set;}
    }
    
    
    public class OrderItem
    {
        public string channel_refnum {get; set;}
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
        public string channel_refnum {get; set;}
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
    
    public class OrderPullResponse
    {
        public string next_page_token {get; set;}
        public Order[] orders {get; set;}
    }

    //TODO: add clases as needed
}