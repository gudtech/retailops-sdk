Verify Serice
---

GT-SOA service invoked for exercising an integration's production deployment.

Playing by the rules for deploying service
---

 1. `docker-entrypoint.sh` generates keypair if missing (hard-coded with service cert name, e.g. `sdk_service.key`)
 3. fragments of the soa.conf file must be written to soa.conf.d and ansible is used to generate the collated soa.conf
  1. put the path to the key/crt in to a fragment
  2. assemble the soa.conf from all fragments (no way to reuse backplane ansible functionality
 4. deploy the service contianer
 5. fingerprint and disseminate fingerprint
  1. 