using System;
using System.Collections.Generic;
using System.Collections.Concurrent;
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

        // TODO: implement the following api repository methods as 
        // defined in IChannelRepository
        // 
        public ChannelResponse catalog_get_config(int id)
        {
            ChannelResponse response = new ChannelResponse();
            response.Fanout = "all relatives";
            return response;
        } 
        //  catalog_push 
        //  inventory_push 
        //  order_pull
        //  order_acknowledge
        //  order_update
        //  order_cancel
        //  order_shipment_submit
        //  order_complete
        //  order_settle_payment
        //  order_returned
    }
}