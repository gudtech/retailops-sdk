[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK
----

Self-verify your RetailOps integration with this SDK.

 1. Follow the prerequisite steps below
 2. Download the `Verify Service` release for your operating system here: [RetailOps SDK Releases Page](https://github.com/gudTECH/retailops-sdk/releases)
 3. Unzip downloaded file
 4. Use terminal and enter unzipped directory (e.g., `verify_darwin_v0.0.6`)
 5. Start the example server application

    ```
    cd retailops-asp-dotnet-api-example
    dnu restore
    dnx web
    ```

    The example web server should now be running on http://0.0.0.0:5000 .

 6. From another terminal run the `verify` utility:

    ```
    $ ./verify -schema-path schema/schemata/
    11 TESTS TO BE GENERATED
    [[ TRUNCATED OUTPUT ]]
    ```

 7. When ready for certification with RetailOps

    ```
    $ ./verify -api-key $API_KEY -base-url http://www.example.com/api/channel certify
    remote certification was a success
    ```

OSX Prerequisites
---

  Please follow the official instructions found here: https://github.com/dotnet/coreclr/blob/master/Documentation/install/get-dotnetcore-dnx-osx.md
