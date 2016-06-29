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

Managing Your Integration Auth Key
---

The integration auth key is used to confirm data is coming from RetailOps. This key must be kept secret and must be checked on every request from RetailOps. This key is a long, random string you will generate using the verify tool.

The verify tool helps manage this key. By default it is stored in your `$HOME/.retailops/` folder. When collaborating with other developers you may need to share this key.

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
$ ./verify.go show-auth-key
==== INTEGRATION AUTH KEY BELOW ====
FOOBARBAZ
==== INTEGRATION AUTH KEY DONE ====
```

OSX Prerequisites
---

  Please follow the official instructions found here: https://github.com/dotnet/coreclr/blob/master/Documentation/install/get-dotnetcore-dnx-osx.md
