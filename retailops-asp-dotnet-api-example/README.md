[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### Example RetailOPS Channel API Application

* DRAFT - THIS DOCUMENTATION IS PENDING FINALIZATION *

This application is provided as an example to illustrate implementation of a channel integration web-hook
developed with ASP .NET Web API. This is not a full production application and should not be used without modification
as it is missing the following:
- Authentication - this example does not implement any sort of authentication
- Data Layer - the example does not connect to any data source (database, channel, etc...) 
- JSON validation - JSON received in requests is currently ignored. A production application needs to parse and use incoming JSON
    
#### Repository Design Pattern
This example is designed using a variation of the repository pattern. This seperates the business logic from the controller by using a repository class
for the data-access and business logic, which is accessed via an interface class. This makes unit-testing easier, and facilitates
changes in the application (such as changing the data source) while constraining most code changes to the repository layer. For example, connecting 
to a database would only require code changes in the repository, as well as some minor changes in the application config.
