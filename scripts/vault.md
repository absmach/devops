# Install and configure `vault` with `certs`
Make sure you configured your `KUBECONFIG` to point to destination cluster.

Install vault

```
helm install vault hashicorp/vault -n mg
```

Initialize vault
```bash
kubectl exec -it vault-0 -n mg -- vault operator init  -key-shares=1 -key-threshold=1
```

Take a not for unseal keys and root token, by default on `init` operation you will get 1 keys and you need 1 to unseal
```bash
kubectl exec vault-0 -n mg -- vault operator unseal <VAULT_UNSEAL_KEY>
```

Edit `.env` and set to `MG_VAULT_TOKEN` to value of root token, additonaly, to setup `mTLS` properly `MG_VAULT_CA_CN` must match host that `Magistrala` is deployed to. 

Execute `/vault-set-pki.sh`

Now upgrade installation of magistrala enabling certs service and setting proper values
```bash
 helm upgrade magistrala  --create-namespace -n mg . \
                        ...
                        --set certs.enabled=true \
                        --set certs.signVaultToken=s.8by6kA75cKciQBQvvkCu21m \
                        --set certs.signVaultHost=http://vault:8200 \
                        --set certs.signVaultPKIPath=pki_int \
                        --set certs.signVaultRole=magistrala
```
