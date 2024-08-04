---
subcategory: "MLflow"
---
# databricks_mlflow_models Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Retrieves a list of [databricks_mlflow_model](../resources/mlflow_model.md) ids, that were created by Terraform or manually, so that special handling could be applied.

## Example Usage

Listing all mlflow models:

```hcl
data "databricks_mlflow_models" "all" {}

output "all_mlflow_models" {
  value = data.databricks_mlflow_models.all
}
```

## Attribute Reference

This data source exports the following attributes:

* `ids` - set of [databricks_mlflow_model](../resources/mlflow_model.md) ids

## Related Resources

The following resources are used in the same context:

* [databricks_model](../resources/mlflow_model.md) to manage models within MLflow.