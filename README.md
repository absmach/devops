## DevOps

Scripts for managing the Magistrala IoT platform.

## Prerequisites

Make sure [Helm](https://helm.sh) is installed. If not, follow the [Helm documentation](https://helm.sh/docs) to get started.

## Usage Instructions

### Step 1: Add the Magistrala DevOps Repository

After installing Helm, add the Magistrala DevOps Helm repository by running:

```bash
helm repo add <name> https://absmach.github.io/devops/
```

- Replace `<name>` with a repository alias of your choice. This alias will be used in future Helm commands.  
  **Example**:

  ```bash
  helm repo add magistrala-devops https://absmach.github.io/devops/
  ```

### Step 2: Add External Helm Repositories for Dependencies

Magistrala charts rely on several external charts. To ensure these dependencies can be resolved, add the required Helm repositories:

```bash
helm repo add nats https://nats-io.github.io/k8s/helm/charts/
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add hashicorp https://helm.releases.hashicorp.com
```

This step is important to ensure all chart dependencies are properly fetched.

### Step 3: Update Helm Repositories

After adding the required repositories, update your local repository cache to fetch the latest chart versions:

```bash
helm repo update
```

This ensures that your Helm client is aware of the latest available versions of the charts in these repositories.

### Step 4: Update Chart Dependencies

If your Helm chart has dependencies specified in `Chart.yaml`, you need to resolve and download them by running:

```bash
helm dependency update charts/magistrala
```

This downloads the necessary dependencies into the `charts/magistrala/charts/` directory so that the Helm chart can install correctly.

### Step 5: Install the Magistrala Chart

To install the Magistrala chart, run:

```bash
helm install my-magistrala magistrala-devops/magistrala --version 1.0.6
```

- `my-magistrala`: Choose a release name for the installation.
- `magistrala-devops/magistrala`: Refers to the chart in the added repository.
- `--version 1.0.6`: Installs version 1.0.6 of the chart.

To customize the installation, you can:

- Use the `--values` flag to provide a custom values file.
- Use the `--set` flag for inline configuration changes.

### Step 6: Search for Charts

To search for charts by keyword in the repository:

```bash
helm search repo [keyword]
```

### Step 7: Uninstall the Magistrala Chart

When you no longer need the chart, uninstall it by running:

```bash
helm uninstall my-magistrala
```

- Replace `my-magistrala` with the release name you used during installation. This command removes all resources created by the chart and clears the release history.

For further details, refer to the [Magistrala Kubernetes documentation](https://docs.magistrala.abstractmachines.fr/kubernetes/).

## License

This project is licensed under the [Apache-2.0](LICENSE).

---