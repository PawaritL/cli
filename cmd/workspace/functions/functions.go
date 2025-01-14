// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package functions

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "functions",
	Short: `Functions implement User-Defined Functions (UDFs) in Unity Catalog.`,
	Long: `Functions implement User-Defined Functions (UDFs) in Unity Catalog.
  
  The function implementation can be any SQL expression or Query, and it can be
  invoked wherever a table reference is allowed in a query. In Unity Catalog, a
  function resides at the same level as a table, so it can be referenced with
  the form __catalog_name__.__schema_name__.__function_name__.`,
	Annotations: map[string]string{
		"package": "catalog",
	},
}

// start create command

var createReq catalog.CreateFunction
var createJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createCmd)
	// TODO: short flags
	createCmd.Flags().Var(&createJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	createCmd.Flags().StringVar(&createReq.Comment, "comment", createReq.Comment, `User-provided free-form text description.`)
	createCmd.Flags().StringVar(&createReq.ExternalLanguage, "external-language", createReq.ExternalLanguage, `External function language.`)
	createCmd.Flags().StringVar(&createReq.ExternalName, "external-name", createReq.ExternalName, `External function name.`)
	// TODO: map via StringToStringVar: properties
	createCmd.Flags().StringVar(&createReq.SqlPath, "sql-path", createReq.SqlPath, `List of schemes whose objects can be referenced without qualification.`)

}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: `Create a function.`,
	Long: `Create a function.
  
  Creates a new function
  
  The user must have the following permissions in order for the function to be
  created: - **USE_CATALOG** on the function's parent catalog - **USE_SCHEMA**
  and **CREATE_FUNCTION** on the function's parent schema`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = createJson.Unmarshal(&createReq)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("please provide command input in JSON format by specifying the --json flag")
		}

		response, err := w.Functions.Create(ctx, createReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete command

var deleteReq catalog.DeleteFunctionRequest
var deleteJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteCmd)
	// TODO: short flags
	deleteCmd.Flags().Var(&deleteJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	deleteCmd.Flags().BoolVar(&deleteReq.Force, "force", deleteReq.Force, `Force deletion even if the function is notempty.`)

}

var deleteCmd = &cobra.Command{
	Use:   "delete NAME",
	Short: `Delete a function.`,
	Long: `Delete a function.
  
  Deletes the function that matches the supplied name. For the deletion to
  succeed, the user must satisfy one of the following conditions: - Is the owner
  of the function's parent catalog - Is the owner of the function's parent
  schema and have the **USE_CATALOG** privilege on its parent catalog - Is the
  owner of the function itself and have both the **USE_CATALOG** privilege on
  its parent catalog and the **USE_SCHEMA** privilege on its parent schema`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = deleteJson.Unmarshal(&deleteReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No NAME argument specified. Loading names for Functions drop-down."
				names, err := w.Functions.FunctionInfoNameToFullNameMap(ctx, catalog.ListFunctionsRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Functions drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "The fully-qualified name of the function (of the form __catalog_name__.__schema_name__.__function__name__)")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have the fully-qualified name of the function (of the form __catalog_name__.__schema_name__.__function__name__)")
			}
			deleteReq.Name = args[0]
		}

		err = w.Functions.Delete(ctx, deleteReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start get command

var getReq catalog.GetFunctionRequest
var getJson flags.JsonFlag

func init() {
	Cmd.AddCommand(getCmd)
	// TODO: short flags
	getCmd.Flags().Var(&getJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var getCmd = &cobra.Command{
	Use:   "get NAME",
	Short: `Get a function.`,
	Long: `Get a function.
  
  Gets a function from within a parent catalog and schema. For the fetch to
  succeed, the user must satisfy one of the following requirements: - Is a
  metastore admin - Is an owner of the function's parent catalog - Have the
  **USE_CATALOG** privilege on the function's parent catalog and be the owner of
  the function - Have the **USE_CATALOG** privilege on the function's parent
  catalog, the **USE_SCHEMA** privilege on the function's parent schema, and the
  **EXECUTE** privilege on the function itself`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = getJson.Unmarshal(&getReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No NAME argument specified. Loading names for Functions drop-down."
				names, err := w.Functions.FunctionInfoNameToFullNameMap(ctx, catalog.ListFunctionsRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Functions drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "The fully-qualified name of the function (of the form __catalog_name__.__schema_name__.__function__name__)")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have the fully-qualified name of the function (of the form __catalog_name__.__schema_name__.__function__name__)")
			}
			getReq.Name = args[0]
		}

		response, err := w.Functions.Get(ctx, getReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start list command

var listReq catalog.ListFunctionsRequest
var listJson flags.JsonFlag

func init() {
	Cmd.AddCommand(listCmd)
	// TODO: short flags
	listCmd.Flags().Var(&listJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var listCmd = &cobra.Command{
	Use:   "list CATALOG_NAME SCHEMA_NAME",
	Short: `List functions.`,
	Long: `List functions.
  
  List functions within the specified parent catalog and schema. If the user is
  a metastore admin, all functions are returned in the output list. Otherwise,
  the user must have the **USE_CATALOG** privilege on the catalog and the
  **USE_SCHEMA** privilege on the schema, and the output list contains only
  functions for which either the user has the **EXECUTE** privilege or the user
  is the owner. There is no guarantee of a specific ordering of the elements in
  the array.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		if cmd.Flags().Changed("json") {
			check = cobra.ExactArgs(0)
		}
		return check(cmd, args)
	},
	PreRunE: root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = listJson.Unmarshal(&listReq)
			if err != nil {
				return err
			}
		} else {
			listReq.CatalogName = args[0]
			listReq.SchemaName = args[1]
		}

		response, err := w.Functions.ListAll(ctx, listReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start update command

var updateReq catalog.UpdateFunction
var updateJson flags.JsonFlag

func init() {
	Cmd.AddCommand(updateCmd)
	// TODO: short flags
	updateCmd.Flags().Var(&updateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	updateCmd.Flags().StringVar(&updateReq.Owner, "owner", updateReq.Owner, `Username of current owner of function.`)

}

var updateCmd = &cobra.Command{
	Use:   "update NAME",
	Short: `Update a function.`,
	Long: `Update a function.
  
  Updates the function that matches the supplied name. Only the owner of the
  function can be updated. If the user is not a metastore admin, the user must
  be a member of the group that is the new function owner. - Is a metastore
  admin - Is the owner of the function's parent catalog - Is the owner of the
  function's parent schema and has the **USE_CATALOG** privilege on its parent
  catalog - Is the owner of the function itself and has the **USE_CATALOG**
  privilege on its parent catalog as well as the **USE_SCHEMA** privilege on the
  function's parent schema.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = updateJson.Unmarshal(&updateReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No NAME argument specified. Loading names for Functions drop-down."
				names, err := w.Functions.FunctionInfoNameToFullNameMap(ctx, catalog.ListFunctionsRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Functions drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "The fully-qualified name of the function (of the form __catalog_name__.__schema_name__.__function__name__)")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have the fully-qualified name of the function (of the form __catalog_name__.__schema_name__.__function__name__)")
			}
			updateReq.Name = args[0]
		}

		response, err := w.Functions.Update(ctx, updateReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// end service Functions
