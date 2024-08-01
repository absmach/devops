#!/usr/bin/bash
# Copyright (c) Abstract Machines
# SPDX-License-Identifier: Apache-2.0

scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $scriptdir

readDotEnv() {
    set -o allexport
    source $scriptdir/.env
    set +o allexport
}


write_env() {
    if [ -e "data/secrets" ]; then
        sed -i "s,MG_VAULT_UNSEAL_KEY_1=.*,MG_VAULT_UNSEAL_KEY_1=$(awk -F ": " '$1 == "Unseal Key 1" {print $2}' data/secrets)," .env
        sed -i "s,MG_VAULT_UNSEAL_KEY_2=.*,MG_VAULT_UNSEAL_KEY_2=$(awk -F ": " '$1 == "Unseal Key 2" {print $2}' data/secrets)," .env
        sed -i "s,MG_VAULT_UNSEAL_KEY_3=.*,MG_VAULT_UNSEAL_KEY_3=$(awk -F ": " '$1 == "Unseal Key 3" {print $2}' data/secrets)," .env
        sed -i "s,MG_VAULT_TOKEN=.*,MG_VAULT_TOKEN=$(awk -F ": " '$1 == "Initial Root Token" {print $2}' data/secrets)," .env
        echo "Vault environment varaibles are set successfully in docker/.env"
    else
        echo "Error: Source file 'data/secrets' not found."
    fi
}



source vault_cmd.sh

readDotEnv
mkdir -p data

# Check Vault initialization status
vault operator init -status -address=$MG_VAULT_ADDR
INIT_STATUS=$?
set -euo pipefail

# Check if Vault is not initialized (exit status 2)
if [ $INIT_STATUS -eq 2 ]; then
  echo "Vault is not initialized. Initializing now..."

  # Initialize Vault and store secrets
  vault operator init -address=$MG_VAULT_ADDR 2>&1 | tee >(sed -r 's/\x1b\[[0-9;]*m//g' > data/secrets)

  echo "Vault initialization complete. Secrets stored in data/secrets."
elif [ $INIT_STATUS -eq 0 ]; then
  echo "Vault is already initialized."
else
  echo "An error occurred while checking Vault initialization status. Exit status: $INIT_STATUS"
fi

readDotEnv
write_env

readDotEnv
vault operator unseal -address=${MG_VAULT_ADDR} ${MG_VAULT_UNSEAL_KEY_1}
vault operator unseal -address=${MG_VAULT_ADDR} ${MG_VAULT_UNSEAL_KEY_2}
vault operator unseal -address=${MG_VAULT_ADDR} ${MG_VAULT_UNSEAL_KEY_3}


./vault_set_pki.sh
./vault_create_approle.sh
