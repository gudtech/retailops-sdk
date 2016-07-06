[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK
----

Self-verify your RetailOps integration with this SDK.

 1. Follow the [prerequisite steps](#osx-prerequisites) below
 2. Download the `Verify Service` release for your operating system here: [RetailOps SDK Releases Page](https://github.com/gudTECH/retailops-sdk/releases)
 3. Unzip downloaded file
 4. Use terminal and enter unzipped directory (e.g., `verify_darwin_v0.0.6`)
 5. Start the example server application

    ```
    cd retailops-asp-dotnet-api-example
    dnu restore
    dnx web
    ```

    The example web server should now be running on http://0.0.0.0:5000.

    > _Note: the example application is provided for testing purposes only, and is not a production application.
    > Do not attempt to use it in place of writing your own channel integration, it exists to help set up and test
    > the verifier tool, and to provide a suggested starting point for .NET developers. It returns hard-coded JSON
    > responses for testing the verifier tool and does not implement any required integration logic, or authentication._      

 6. From another terminal run the `verify` utility:

    ```
    $ ./verify -schema-path schema/schemata/
    11 TESTS TO BE GENERATED
    [[ TRUNCATED OUTPUT ]]
    ```

 After you have completely developed your channel integration, and have successfully used the verifier tool to
 test that your integration is operating correctly, you are ready to attempt certification
 with RetailOps.

Follow the instructions here: [Certifying Your RetailOps SDK Channel Integration](https://github.com/gudTECH/retailops-sdk/blob/master/verify/CERTIFY_README.md)

Managing Your Integration Auth Key
---

The integration auth key is used to confirm data is coming from RetailOps. This key must be kept secret and must be checked on every request from RetailOps. A bad integration auth key value must result in a HTTP 401. The integration auth key is a long, random string you will generate using the verify tool.

The verify tool helps manage this key. By default it is stored in your `$HOME/.retailops/` folder. When collaborating with other developers you may need to share and install this key.

To generate a new key, required for every service:

```
$ ./verify generate-auth-key
$ ./verify show-auth-key
==== INTEGRATION AUTH KEY BELOW ====
dDkma7W9zA/Nd7r3AUcEI/9MA9KWg+Uc
==== INTEGRATION AUTH KEY DONE ====
```

To install a different auth key, setting it to whatever value you choose:

```
$ ./verify install-auth-key FOOBARBAZ
$ ./verify show-auth-key
==== INTEGRATION AUTH KEY BELOW ====
FOOBARBAZ
==== INTEGRATION AUTH KEY DONE ====
```

> _Note: the integration auth key generated here is only for certification of a production service. The local verification will use the the integration_auth_key "RETAILOPS_SDK" and you should reject all other values._


OSX Prerequisites
---

  Please follow the official instructions found here: https://github.com/dotnet/coreclr/blob/master/Documentation/install/get-dotnetcore-dnx-osx.md
