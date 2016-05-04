[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### Example RetailOPS Channel API Application

* DRAFT - THIS EXAMPLE APPLICATION IS PENDING FINALIZATION *

This application is provided as an example to illustrate implementation of a channel integration web-hook
developed with ASP .NET Web API. This is not a full production application and should not be used without modification
as it is missing the following:
- Authentication - this example does not implement any sort of authentication
- Data Layer - the example does not connect to any data source (database, channel, etc...) 
- JSON validation - JSON received in requests is currently ignored. A production application needs to parse and use incoming JSON
    
### Windows Installation
To run the example ASP .NET aplication for testing, you must first install ASP .NET 5. Please follow the instructions provided
[here](http://docs.asp.net/en/latest/getting-started/installing-on-windows.html). 
Once ASP .NET 5 is installed open a command window, and navigate to the directory that contains the applications source files.
Once in the directory run the following commands:
- `dnu restore` this will restore all of the project's required resources  
- `dnu build` to rebuild the executable for your environment
- `dnx web` to start the application on http://localhost:5000. (leave this command prompt open)  

After the ASP .NET application is running, open a second command prompt, and navigate to the `verify` release folder
and run:
- `verify.exe -schema-path ".\schema\schemata"`

If the command is succesful, you'll see something like this:
```Bash
3 REQUESTS TO BE GENERATED
-----------
POST /api/channel/catalog_get_config
{
  "version": 1,
  "action": "catpush_config",
  "data": {}
}
HTTP status code 200
data:
{
  "sku_fanout": "all_skus_for_product"
}
response valid: true
SUCCESS
-----------
POST /api/channel/order_acknowledge
{
  "version": 1,
  "action": "order_acknowledge",
  "data": {}
}
HTTP status code 200
data:
{
  "events": [
    {
      "associations": [
        {
          "identity": "S1234",
          "type": "sku"
        }
      ],
      "diagnostic_data": [],
      "error_code": "ERR1234",
      "error_message": "Example error message",
      "status": "error"
    }
  ]
}
response valid: true
SUCCESS
-----------
```
#### Repository Design Pattern
This example is designed using a variation of the repository pattern. This seperates the business logic from the controller 
by using a repository class for the data-access and business logic, which is accessed via an interface class. 
This loose coupling makes makes unit-testing, and adapting to future requirments much easier.