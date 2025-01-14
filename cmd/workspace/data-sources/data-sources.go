// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_sources

import (
	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "data-sources",
	Short: `This API is provided to assist you in making new query objects.`,
	Long: `This API is provided to assist you in making new query objects. When creating
  a query object, you may optionally specify a data_source_id for the SQL
  warehouse against which it will run. If you don't already know the
  data_source_id for your desired SQL warehouse, this API will help you find
  it.
  
  This API does not support searches. It returns the full list of SQL warehouses
  in your workspace. We advise you to use any text editor, REST client, or
  grep to search the response from this API for the name of your SQL warehouse
  as it appears in Databricks SQL.`,
	Annotations: map[string]string{
		"package": "sql",
	},
}

// start list command

func init() {
	Cmd.AddCommand(listCmd)

}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: `Get a list of SQL warehouses.`,
	Long: `Get a list of SQL warehouses.
  
  Retrieves a full list of SQL warehouses available in this workspace. All
  fields that appear in this API response are enumerated for clarity. However,
  you need only a SQL warehouse's id to create new queries against it.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		response, err := w.DataSources.List(ctx)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// end service DataSources
