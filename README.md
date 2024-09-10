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

### Step 2: Install the Magistrala Chart

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

### Step 3: Update the Repository

If you’ve already added the Magistrala repo, update it to fetch the latest chart versions:

```bash
helm repo update
```

### Step 4: Search for Charts

To search for charts by keyword in the repository:

```bash
helm search repo [keyword]
```

### Step 5: Uninstall the Magistrala Chart

When you no longer need the chart, uninstall it by running:

```bash
helm uninstall my-magistrala
```

- Replace `my-magistrala` with the release name you used during installation. This command removes all resources created by the chart and clears the release history.

For further details, refer to the [Magistrala Kubernetes documentation](https://docs.magistrala.abstractmachines.fr/kubernetes/).

## License

This project is licensed under the [Apache-2.0](LICENSE).

---