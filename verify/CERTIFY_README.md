### Certifying Your RetailOps SDK Channel Integration

After you have completely developed your channel integration, and *have successfully used the verifier tool to
test that your integration is operating correctly*, you are ready to attempt certification
with RetailOps.

> _Do not attempt to certify your integration prior to it being completed, or if your service does
> not successfully pass local testing with the verifier tool. Successful testing with the verifier
> tool is a required prerequisite for certification. Additionally, do not attempt to certify the
> example application provided with the RetailOps SDK_

Run the following command to certify your RetailOps Channel Integration:

   ```
   $ ./verify.go -api-key $API_KEY -base-url http://api.your-server.com -integration-name 'YOUR-INTEGRATION-NAME-HERE' certify
   
   remote certification was a success
   ```
