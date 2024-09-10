## Usage

[Helm](https://helm.sh) must be installed to use Magistrala charts. Please refer to Helm's [documentation](https://helm.sh/docs) to get started if you haven't already.

### Step 1: Adding the Magistrala DevOps Repo

Once Helm has been set up correctly, add the Magistrala DevOps repository by running the following command:

```bash
helm repo add magistrala-devops https://absmach.github.io/devops/
```

This command adds the Magistrala DevOps Helm repository to your local Helm setup, allowing you to access the charts it contains.

### Step 2: Updating the Repository

If you had already added this repo earlier, run the following command to retrieve the latest versions of the packages:

```bash
helm repo update
```

This ensures that you are working with the latest charts from the repository.

### Step 3: Searching for Charts

You can now search the repository for a keyword in charts using the following command:

```bash
helm search repo [keyword] [flags]
```

### Step 4: Installing the Chart

To install the Magistrala chart, run the following command:

```bash
helm install my-magistrala magistrala-devops/magistrala --version 1.0.5
```

- `my-magistrala` is the release name you choose for the installation. You can modify this to suit your needs.
- `magistrala-devops/magistrala` specifies the chart name from the added repository.
- `--version 1.0.5` ensures you install the specific version of the chart.

You can also add additional flags as needed, such as specifying a custom `values.yaml` file, enabling debugging, or overriding values directly in the command.

### Step 5: Uninstalling the Chart

To uninstall the chart when you no longer need it, run the following command:

```bash
helm uninstall my-magistrala [flags]
```

Make sure to replace `my-magistrala` with your actual release name if you changed it during installation. The command removes all of the resources associated with the last release of the chart as well as the release history, freeing it up for future use.

---
