## How to Install and Configure `vault` with `certs`

### Prerequisites:

1. **Kubernetes Configuration**: Ensure your `KUBECONFIG` is set up to point to the Kubernetes cluster where you want to deploy `vault`. This can typically be done by running:
   ```bash
   export KUBECONFIG=/path/to/your/kubeconfig
   ```
   This command tells your local machine which Kubernetes cluster to interact with.

### Step 1: Install `vault` using Helm

1. **Navigate to the `magistrala` Helm chart directory**:

   ```bash
   cd charts/magistrala
   ```

2. **Install `vault`**:
   ```bash
   helm upgrade magistrala . -n mg --set vault.enabled=true
   ```
   This command uses Helm to upgrade (or install) the `magistrala` release in the `mg` namespace with `vault` enabled.

### Step 2: Initialize `vault`

1. **Navigate to the `vault` Scripts Directory**:

   If you are currently in the `charts/magistrala` directory, go up two levels to the root and then to the `vault` scripts directory by running:

   ```bash
   cd ../../scripts/vault
   ```

   If you are at the root of the repository, navigate to the `vault` scripts directory directly by running:

   ```bash
   cd scripts/vault
   ```

2. **Run the `vault_init.sh` script**:
   ```bash
   ./vault_init.sh
   ```
   This script initializes `vault` by setting up necessary configurations, such as unsealing the vault and applying initial policies. This is a crucial step to get `vault` ready for use.

### Step 3: Enable the `certs` Service and Apply Configuration

1. **Load Environment Variables**:

   ```bash
   source .env
   ```

   This command loads environment variables from the `.env` file into your current shell session. These variables are required for the next step to configure the `certs` service.

2. **Navigate back to the `magistrala` Helm chart directory**:

   ```bash
   cd ../../charts/magistrala
   ```

3. **Upgrade the `magistrala` installation with `certs` enabled**:
   ```bash
   helm upgrade magistrala --create-namespace -n mg . \
       --set certs.vault.url=$MG_VAULT_ADDR \
       --set certs.vault.approleRoleid=$MG_VAULT_THINGS_CERTS_ISSUER_ROLEID \
       --set certs.vault.approleSecret=$MG_VAULT_THINGS_CERTS_ISSUER_SECRET \
       --set certs.vault.namespace=$MG_VAULT_NAMESPACE
   ```
