// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networks

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "networks",
	Short: `These APIs manage network configurations for customer-managed VPCs (optional).`,
	Long: `These APIs manage network configurations for customer-managed VPCs (optional).
  Its ID is used when creating a new workspace if you use customer-managed VPCs.`,
	Annotations: map[string]string{
		"package": "provisioning",
	},
}

// start create command

var createReq provisioning.CreateNetworkRequest
var createJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createCmd)
	// TODO: short flags
	createCmd.Flags().Var(&createJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	// TODO: complex arg: gcp_network_info
	// TODO: array: security_group_ids
	// TODO: array: subnet_ids
	// TODO: complex arg: vpc_endpoints
	createCmd.Flags().StringVar(&createReq.VpcId, "vpc-id", createReq.VpcId, `The ID of the VPC associated with this network.`)

}

var createCmd = &cobra.Command{
	Use:   "create NETWORK_NAME",
	Short: `Create network configuration.`,
	Long: `Create network configuration.
  
  Creates a Databricks network configuration that represents an VPC and its
  resources. The VPC will be used for new Databricks clusters. This requires a
  pre-existing VPC and subnets.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
		if cmd.Flags().Changed("json") {
			check = cobra.ExactArgs(0)
		}
		return check(cmd, args)
	},
	PreRunE: root.MustAccountClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)
		if cmd.Flags().Changed("json") {
			err = createJson.Unmarshal(&createReq)
			if err != nil {
				return err
			}
		} else {
			createReq.NetworkName = args[0]
		}

		response, err := a.Networks.Create(ctx, createReq)
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

var deleteReq provisioning.DeleteNetworkRequest
var deleteJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteCmd)
	// TODO: short flags
	deleteCmd.Flags().Var(&deleteJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var deleteCmd = &cobra.Command{
	Use:   "delete NETWORK_ID",
	Short: `Delete a network configuration.`,
	Long: `Delete a network configuration.
  
  Deletes a Databricks network configuration, which represents a cloud VPC and
  its resources. You cannot delete a network that is associated with a
  workspace.
  
  This operation is available only if your account is on the E2 version of the
  platform.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustAccountClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)
		if cmd.Flags().Changed("json") {
			err = deleteJson.Unmarshal(&deleteReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No NETWORK_ID argument specified. Loading names for Networks drop-down."
				names, err := a.Networks.NetworkNetworkNameToNetworkIdMap(ctx)
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Networks drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "Databricks Account API network configuration ID")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have databricks account api network configuration id")
			}
			deleteReq.NetworkId = args[0]
		}

		err = a.Networks.Delete(ctx, deleteReq)
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

var getReq provisioning.GetNetworkRequest
var getJson flags.JsonFlag

func init() {
	Cmd.AddCommand(getCmd)
	// TODO: short flags
	getCmd.Flags().Var(&getJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var getCmd = &cobra.Command{
	Use:   "get NETWORK_ID",
	Short: `Get a network configuration.`,
	Long: `Get a network configuration.
  
  Gets a Databricks network configuration, which represents a cloud VPC and its
  resources.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustAccountClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)
		if cmd.Flags().Changed("json") {
			err = getJson.Unmarshal(&getReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No NETWORK_ID argument specified. Loading names for Networks drop-down."
				names, err := a.Networks.NetworkNetworkNameToNetworkIdMap(ctx)
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Networks drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "Databricks Account API network configuration ID")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have databricks account api network configuration id")
			}
			getReq.NetworkId = args[0]
		}

		response, err := a.Networks.Get(ctx, getReq)
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

func init() {
	Cmd.AddCommand(listCmd)

}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: `Get all network configurations.`,
	Long: `Get all network configurations.
  
  Gets a list of all Databricks network configurations for an account, specified
  by ID.
  
  This operation is available only if your account is on the E2 version of the
  platform.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustAccountClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)
		response, err := a.Networks.List(ctx)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// end service Networks
