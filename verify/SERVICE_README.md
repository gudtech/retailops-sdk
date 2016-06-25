Verify Serice
---

GT-SOA service invoked for exercising an integration's production deployment.

Playing by the rules for deploying service
---

 1. `docker-entrypoint.sh` generates keypair if missing (hard-coded with service cert name, e.g. `sdk_service.key`)
 3. fragments of the soa.conf file must be written to soa.conf.d and ansible is used to generate the collated soa.conf
  1. put the path to the key/crt in to a fragment
  2. assemble the soa.conf from all fragments (no way to reuse backplane ansible functionality
 4. deploy the service container
 5. fingerprint and disseminate fingerprint
 6. install the actions in the authz table as public:

```
docker exec -it main bash -c 'echo "PUBLIC actions+ integration.channel.certify" | perl ~/gt-core/bin/priv/load'
```

 7. generate the permission for registration

```
docker exec -it main bash -c 'echo "integration.register actions+ integration.channel.register" | perl ~/gt-core/bin/priv/load'
```

 8. generate the long-lived token for the container to use when calling registration

```
docker exec -it auth bash -c 'mkdir -p /etc/GT/endorsements && perl /opt/gt/gt-auth-service/script/maketicket --timeout 10y --priv integration.register --user "RO Automated Agent" > /tmp/gtuser-sdk_service.ticket'
```

 9. Need to add a foreign key to contstans for the `implementer_user_id` field:

```
ALTER TABLE constants.channel_definition ADD COLUMN implementer_user_id int(10) unsigned;
```