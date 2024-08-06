# Install and configure `vault` with `certs`
Make sure you configured your `KUBECONFIG` to point to destination cluster.

Install vault

```bash
cd charts/magistrala
helm upgrade magistrala . -n mg  --set vault.enabled=true
```

Initialize vault
```bash
cd ../../scripts/vault
./vault_init.sh
```


Now upgrade installation of magistrala enabling certs service and setting proper values
```bash
source .env
cd ../../charts/magistrala
helm upgrade magistrala  --create-namespace -n mg . \
                    --set certs.vault.url=$MG_VAULT_ADDR \
                    --set certs.vault.approleRoleid=$MG_VAULT_THINGS_CERTS_ISSUER_ROLEID \
                    --set certs.vault.approleSecret=$MG_VAULT_THINGS_CERTS_ISSUER_SECRET \
                    --set certs.vault.namespace=$MG_VAULT_NAMESPACE 
```
