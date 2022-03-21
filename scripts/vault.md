# Install and configure `vault` with `certs`

Install vault

```
helm install vault hashicorp/vault -n mf
```

Initialize vault
```bash
kubectl exec -it vault-0 -n mf -- vault operator init  -key-shares=1 -key-threshold=1
```

Take a not for unseal keys and root token, by default on `init` operation you will get 5 keys and you need 3 to unseal
```bash
kubectl exec vault-0 -n vault -- vault operator unseal <VAULT_UNSEAL_KEY>
```

Edit `.env` and set to `MF_VAULT_TOKEN` to value of root token 

Execute `/vault-set-pki.sh`

Now upgrade installation of mainflux enabling certs service and setting proper values
```bash
 helm upgrade mainflux  --create-namespace -n mf . \
                        ...
                        --set certs.enabled=true \
                        --set certs.signVaultToken=s.8by6kA75cKciQBQvvkCu21m \
                        --set certs.signVaultHost=http://vault:8200 \
                        --set certs.signVaultPKIPath=pki_int \
                        --set certs.signVaultRole=mainflux
```
