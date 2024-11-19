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

For detailed documentation on the available scripts and their usage, visit the [Vault Addon Documentation](https://github.com/absmach/magistrala/tree/main/docker/addons/vault#readme).