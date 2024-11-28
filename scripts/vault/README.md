# Running Vault Setup Scripts with `--env-file`

To execute a Vault setup script, use the `--env-file` option to provide the path to your `.env` file. Here's the general syntax:

```bash
./<script-name>.sh --env-file <path-to-your-env-file>
```

### Example

To initialize Vault using the provided setup script, run:

```bash
scripts/vault/scripts/vault_init.sh --env-file scripts/vault/.env
```

# Developer Guide: Updating Vault Scripts from Magistrala

This guide provides step-by-step instructions to update the Vault scripts in your local Magistrala DevOps repository, sync them with the `main` branch of the Magistrala repository, and create a pull request (PR) to merge the changes.

## Prerequisites

Make sure you have the following:

- A local clone of the Magistrala DevOps repository.
- Access to the Magistrala GitHub repository (`https://github.com/absmach/magistrala.git`).

## Step 1: Create new branch

Create a new branch from `master` brach

````bash
git checkout -b <your-branch-name>
``

Replace `<your-branch-name>` with the name a new branch name.

## Step 2: Add the Magistrala Remote

If the Magistrala remote is not already added to your local repository, use the following command to add it:

```bash
git remote add -f magistrala https://github.com/absmach/magistrala.git
``

## Step 3: Add Subtree for Vault Scripts

If the subtree for the Vault scripts has not been added, execute the following command:

```bash
git subtree add --prefix=scripts/vault/scripts magistrala main --squash --prefix=docker/addons/vault/scripts
``

- `--prefix=scripts/vault/scripts`: Specifies the target directory in your local repository where the Vault scripts will be added.
- `magistrala main`: Refers to the `main` branch of the Magistrala repository.
- `--squash`: Combines all commits from the Magistrala `main` branch into a single commit when adding the subtree.

## Step 4: Update Vault Scripts to the Latest Version

To update the Vault scripts to the latest version and synchronize with the `main` branch of the Magistrala repository, use the following command:

```bash
git subtree pull --prefix=scripts/vault/scripts magistrala main --squash --prefix=docker/addons/vault/scripts
``

- This command pulls the latest changes from the `main` branch of the Magistrala repository.
- `--squash` creates a single commit for the changes, making the history simpler to manage.

## Step 5: Push Changes and Create a Pull Request

After syncing with the Magistrala `main` branch, push the changes to your working branch:

```bash
git push origin <your-branch-name>
``

Replace `<your-branch-name>` with the name of the branch you're working on.

Once the changes are pushed, go to your GitHub repository and create a pull request (PR) to merge the updates.

## Summary of Commands related to Git subtree

### Add Magistrala Remote
```bash
git remote add -f magistrala https://github.com/absmach/magistrala.git
``

### Add Subtree for Vault Scripts
```bash
git subtree add --prefix=scripts/vault/scripts magistrala main --squash --prefix=docker/addons/vault/scripts
``

### Update Vault Scripts to the Latest
```bash
git subtree pull --prefix=scripts/vault/scripts magistrala main --squash --prefix=docker/addons/vault/scripts
``
````
