package mlflow

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestModelsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registered-models/list?",
				Response: ml.ListModelsResponse{
					RegisteredModels: []ml.Model{
						{
							Name: "b",
						},
						{
							Name: "a",
						},
					},
				},
			},
		},
		Resource:    DataSourceModels(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{"a", "b"},
	})
}

func TestModelsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceModels(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
