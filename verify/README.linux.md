[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK
----

Self-verify your RetailOps integration with this SDK.

 1. Follow the prerequisite steps below
 2. Download the `Verify Service` release for your operating system here: [RetailOps SDK Releases Page](https://github.com/gudTECH/retailops-sdk/releases)
 3. Unzip downloaded file
 4. Use terminal and enter unzipped directory (e.g., `verify_linux_v0.0.6`)
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


Linux Prerequisites
---

Install the .NET runtime for Linux:

```
sudo apt-get install unzip curl libcurl3-gnutls libcurl3-gnutls
curl -sSL https://raw.githubusercontent.com/aspnet/Home/dev/dnvminstall.sh | DNX_BRANCH=dev sh && source ~/.dnx/dnvm/dnvm.sh
sudo apt-get install libunwind8 gettext libssl-dev libcurl4-openssl-dev zlib1g libicu-dev uuid-dev
dnvm upgrade -r coreclr
```

Install libUV for the .NET web server
```
sudo apt-get install build-essential automake libtool
curl -LO https://github.com/libuv/libuv/archive/v1.9.1.tar.gz
tar xzf v1.9.1.tar.gz
cd libuv-1.9.1/
sh autogen.sh
./configure
make
sudo make install
sudo ldconfig
```

Your environment is now setup. You can follow the instructions from the top of this README.
