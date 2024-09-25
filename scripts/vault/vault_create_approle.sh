#!/usr/bin/bash
# Copyright (c) Abstract Machines
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $scriptdir
echo "Script directory set to: $scriptdir"

SKIP_ENABLE_APP_ROLE=${1:-}

readDotEnv() {
    set -o allexport
    echo "Sourcing environment variables from .env file..."
    source $scriptdir/.env
    set +o allexport
}

# Check if a service is running in the Kubernetes cluster
is_service_running() {
    local service_name="$1"
    local namespace="${2:-default}"  # Default namespace is 'default' if not specified

    echo "Checking if service $service_name is running in namespace $namespace..."
    if kubectl get svc -n "$namespace" | grep -q "^$service_name"; then
        echo "Service $service_name is running."
        return 0
    else
        echo "Service $service_name is not running or not found in the namespace $namespace."
        return 1
    fi
}

source vault_cmd.sh

vaultCreatePolicyFile() {
    echo "Creating policy file from template..."
    envsubst '
    ${MG_VAULT_PKI_INT_PATH}
    ${MG_VAULT_PKI_INT_THINGS_CERTS_ROLE_NAME}
    ' < magistrala_things_certs_issue.template.hcl > magistrala_things_certs_issue.hcl

    if [ -f magistrala_things_certs_issue.hcl ]; then
        echo "Policy file magistrala_things_certs_issue.hcl created successfully."
    else
        echo "Failed to create policy file magistrala_things_certs_issue.hcl."
        exit 1
    fi
}

vaultCreatePolicy() {
    echo "Creating new policy for AppRole"
    if is_service_running "magistrala-vault" "mg"; then
        echo "Proceeding with policy creation..."
        
        echo "Copying policy file to the pod..."
        kubectl cp magistrala_things_certs_issue.hcl mg/magistrala-vault-0:/tmp/magistrala_things_certs_issue.hcl
        
        echo "Policy file copied to pod. Now attempting to create policy in Vault..."
        
        # Run the policy creation inside the pod
        kubectl exec magistrala-vault-0 -n mg -- vault policy write -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} magistrala_things_certs_issue /tmp/magistrala_things_certs_issue.hcl
        
    else
        echo "Service magistrala-vault is not running or not found in the mg namespace."
        exit 1
    fi
}

vaultEnableAppRole() {
   if [ "$SKIP_ENABLE_APP_ROLE" == "--skip-enable-approle" ]; then
        echo "Skipping Enable AppRole"
    else
        echo "Enabling AppRole"
        vault auth enable -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} approle
    fi
}

vaultDeleteRole() {
    echo "Deleting old AppRole"
    vault delete -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} auth/approle/role/magistrala_things_certs_issuer
}

vaultCreateRole() {
    echo "Creating new AppRole"
    vault write -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} auth/approle/role/magistrala_things_certs_issuer \
    token_policies=magistrala_things_certs_issue  secret_id_num_uses=0 \
    secret_id_ttl=0 token_ttl=1h token_max_ttl=3h  token_num_uses=0
}

vaultWriteCustomRoleID(){
    echo "Writing custom role id"
    vault read -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} auth/approle/role/magistrala_things_certs_issuer/role-id
    vault write -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} auth/approle/role/magistrala_things_certs_issuer/role-id role_id=${MG_VAULT_THINGS_CERTS_ISSUER_ROLEID}
}

vaultWriteCustomSecret() {
    echo "Writing custom secret"
    vault write -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} -f auth/approle/role/magistrala_things_certs_issuer/secret-id
    vault write -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} auth/approle/role/magistrala_things_certs_issuer/custom-secret-id secret_id=${MG_VAULT_THINGS_CERTS_ISSUER_SECRET} num_uses=0 ttl=0
}

vaultTestRoleLogin() {
    echo "Testing custom roleid secret by logging in"
    vault write -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} auth/approle/login \
        role_id=${MG_VAULT_THINGS_CERTS_ISSUER_ROLEID} \
        secret_id=${MG_VAULT_THINGS_CERTS_ISSUER_SECRET}

}
if ! command -v jq &> /dev/null
then
    echo "jq command could not be found, please install it and try again."
    exit 1
fi

echo "Reading environment variables..."
readDotEnv

echo "Logging into Vault..."
vault login -namespace=${MG_VAULT_NAMESPACE} -address=${MG_VAULT_ADDR} ${MG_VAULT_TOKEN}

echo "Creating policy file..."
vaultCreatePolicyFile

echo "Creating policy in Vault..."
vaultCreatePolicy

vaultEnableAppRole
vaultDeleteRole
vaultCreateRole
vaultWriteCustomRoleID
vaultWriteCustomSecret
vaultTestRoleLogin

exit 0
