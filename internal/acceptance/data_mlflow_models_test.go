package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceModels(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_mlflow_models" "this" {
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_mlflow_models.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			ids := r.Primary.Attributes["ids.#"]
			if ids == "" {
				return fmt.Errorf("ids are empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}