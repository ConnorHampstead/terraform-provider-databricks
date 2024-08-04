package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceModels() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}, w *databricks.WorkspaceClient) error {
		models, err := w.ModelRegistry.ListModelsAll(ctx, ml.ListModelsRequest{})
		if err != nil {
			return err
		}
		for _, v := range models {
			data.Ids = append(data.Ids, v.Name)
		}
		return nil
	})
}
