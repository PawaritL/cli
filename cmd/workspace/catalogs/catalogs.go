// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalogs

import (
	"github.com/databricks/bricks/cmd/root"
	"github.com/databricks/bricks/libs/cmdio"
	"github.com/databricks/bricks/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "catalogs",
	Short: `A catalog is the first layer of Unity Catalog’s three-level namespace.`,
	Long: `A catalog is the first layer of Unity Catalog’s three-level namespace.
  It’s used to organize your data assets. Users can see all catalogs on which
  they have been assigned the USE_CATALOG data permission.
  
  In Unity Catalog, admins and data stewards manage users and their access to
  data centrally across all of the workspaces in a Databricks account. Users in
  different workspaces can share access to the same data, depending on
  privileges granted centrally in Unity Catalog.`,
}

// start create command

var createReq catalog.CreateCatalog
var createJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createCmd)
	// TODO: short flags
	createCmd.Flags().Var(&createJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	createCmd.Flags().StringVar(&createReq.Comment, "comment", createReq.Comment, `User-provided free-form text description.`)
	// TODO: map via StringToStringVar: properties
	createCmd.Flags().StringVar(&createReq.ProviderName, "provider-name", createReq.ProviderName, `The name of delta sharing provider.`)
	createCmd.Flags().StringVar(&createReq.ShareName, "share-name", createReq.ShareName, `The name of the share under the share provider.`)
	createCmd.Flags().StringVar(&createReq.StorageRoot, "storage-root", createReq.StorageRoot, `Storage root URL for managed tables within catalog.`)

}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: `Create a catalog.`,
	Long: `Create a catalog.
  
  Creates a new catalog instance in the parent metastore if the caller is a
  metastore admin or has the **CREATE_CATALOG** privilege.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		err = createJson.Unmarshal(&createReq)
		if err != nil {
			return err
		}
		createReq.Name = args[0]

		response, err := w.Catalogs.Create(ctx, createReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
}

// start delete command

var deleteReq catalog.DeleteCatalogRequest

func init() {
	Cmd.AddCommand(deleteCmd)
	// TODO: short flags

	deleteCmd.Flags().BoolVar(&deleteReq.Force, "force", deleteReq.Force, `Force deletion even if the catalog is not empty.`)

}

var deleteCmd = &cobra.Command{
	Use:   "delete NAME",
	Short: `Delete a catalog.`,
	Long: `Delete a catalog.
  
  Deletes the catalog that matches the supplied name. The caller must be a
  metastore admin or the owner of the catalog.`,

	Annotations: map[string]string{},
	Args:        cobra.ExactArgs(1),
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		deleteReq.Name = args[0]

		err = w.Catalogs.Delete(ctx, deleteReq)
		if err != nil {
			return err
		}
		return nil
	},
}

// start get command

var getReq catalog.GetCatalogRequest

func init() {
	Cmd.AddCommand(getCmd)
	// TODO: short flags

}

var getCmd = &cobra.Command{
	Use:   "get NAME",
	Short: `Get a catalog.`,
	Long: `Get a catalog.
  
  Gets the specified catalog in a metastore. The caller must be a metastore
  admin, the owner of the catalog, or a user that has the **USE_CATALOG**
  privilege set for their account.`,

	Annotations: map[string]string{},
	Args:        cobra.ExactArgs(1),
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		getReq.Name = args[0]

		response, err := w.Catalogs.Get(ctx, getReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
}

// start list command

func init() {
	Cmd.AddCommand(listCmd)

}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: `List catalogs.`,
	Long: `List catalogs.
  
  Gets an array of catalogs in the metastore. If the caller is the metastore
  admin, all catalogs will be retrieved. Otherwise, only catalogs owned by the
  caller (or for which the caller has the **USE_CATALOG** privilege) will be
  retrieved. There is no guarantee of a specific ordering of the elements in the
  array.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		response, err := w.Catalogs.ListAll(ctx)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
}

// start update command

var updateReq catalog.UpdateCatalog
var updateJson flags.JsonFlag

func init() {
	Cmd.AddCommand(updateCmd)
	// TODO: short flags
	updateCmd.Flags().Var(&updateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	updateCmd.Flags().StringVar(&updateReq.Comment, "comment", updateReq.Comment, `User-provided free-form text description.`)
	updateCmd.Flags().StringVar(&updateReq.Name, "name", updateReq.Name, `Name of catalog.`)
	updateCmd.Flags().StringVar(&updateReq.Owner, "owner", updateReq.Owner, `Username of current owner of catalog.`)
	// TODO: map via StringToStringVar: properties

}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: `Update a catalog.`,
	Long: `Update a catalog.
  
  Updates the catalog that matches the supplied name. The caller must be either
  the owner of the catalog, or a metastore admin (when changing the owner field
  of the catalog).`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		err = updateJson.Unmarshal(&updateReq)
		if err != nil {
			return err
		}
		updateReq.Name = args[0]

		response, err := w.Catalogs.Update(ctx, updateReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
}

// end service Catalogs
