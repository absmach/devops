## Usage

[Helm](https://helm.sh) must be installed to use Magistrala charts. Please refer to Helm's [documentation](https://helm.sh/docs) to get started if you haven't already.

### Step 1: Adding the Magistrala DevOps Repo

Once Helm has been set up correctly, add the Magistrala DevOps repository by running the following command:

```bash
helm repo add <name> https://absmach.github.io/devops/
```

`<name>` is the alias or label that you assign to the repository you are adding. It can be any name of your choice, and it will act as a reference for the repository in your subsequent Helm commands.

For example:

```bash
helm repo add magistrala-devops https://absmach.github.io/devops/
```

### Step 2: Installing the Chart

To install the Magistrala chart, run the following command:

```bash
helm install my-magistrala magistrala-devops/magistrala --version 1.0.6
```

* `my-magistrala` is the release name you choose for the installation. You can modify this to suit your needs.
* `magistrala-devops/magistrala` specifies the chart name from the added repository.
* `--version 1.0.6` ensures you install the specific version of the chart.

* To override values in the chart, use either the `--values` flag and pass in a file or use the `--set` flag and pass configuration from the command line.

### Step 3: Updating the Repository

If you had already added this repo earlier, run the following command to retrieve the latest versions of the packages:

```bash
helm repo update
```

### Step 4: Searching for Charts

You can search the repository for a keyword in charts using the following command:

```bash
helm search repo [keyword] [flags]
```

### Step 5: Uninstalling the Chart

To uninstall the chart when you no longer need it, run the following command:

```bash
helm uninstall my-magistrala [flags]
```

Make sure to replace `my-magistrala` with your actual release name if you changed it during installation. The command removes all of the resources associated with the last release of the chart as well as the release history, freeing it up for future use.

---