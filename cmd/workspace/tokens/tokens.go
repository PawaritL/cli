// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "tokens",
	Short: `The Token API allows you to create, list, and revoke tokens that can be used to authenticate and access Databricks REST APIs.`,
	Long: `The Token API allows you to create, list, and revoke tokens that can be used
  to authenticate and access Databricks REST APIs.`,
}

// start create command

var createReq settings.CreateTokenRequest

func init() {
	Cmd.AddCommand(createCmd)
	// TODO: short flags

	createCmd.Flags().StringVar(&createReq.Comment, "comment", createReq.Comment, `Optional description to attach to the token.`)
	createCmd.Flags().Int64Var(&createReq.LifetimeSeconds, "lifetime-seconds", createReq.LifetimeSeconds, `The lifetime of the token, in seconds.`)

}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: `Create a user token.`,
	Long: `Create a user token.
  
  Creates and returns a token for a user. If this call is made through token
  authentication, it creates a token with the same client ID as the
  authenticated token. If the user's token quota is exceeded, this call returns
  an error **QUOTA_EXCEEDED**.`,

	Annotations: map[string]string{},
	Args:        cobra.ExactArgs(0),
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		response, err := w.Tokens.Create(ctx, createReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
}

// start delete command

var deleteReq settings.RevokeTokenRequest

func init() {
	Cmd.AddCommand(deleteCmd)
	// TODO: short flags

}

var deleteCmd = &cobra.Command{
	Use:   "delete TOKEN_ID",
	Short: `Revoke token.`,
	Long: `Revoke token.
  
  Revokes an access token.
  
  If a token with the specified ID is not valid, this call returns an error
  **RESOURCE_DOES_NOT_EXIST**.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if len(args) == 0 {
			names, err := w.Tokens.PublicTokenInfoCommentToTokenIdMap(ctx)
			if err != nil {
				return err
			}
			id, err := cmdio.Select(ctx, names, "The ID of the token to be revoked")
			if err != nil {
				return err
			}
			args = append(args, id)
		}
		if len(args) != 1 {
			return fmt.Errorf("expected to have the id of the token to be revoked")
		}
		deleteReq.TokenId = args[0]

		err = w.Tokens.Delete(ctx, deleteReq)
		if err != nil {
			return err
		}
		return nil
	},
}

// start list command

func init() {
	Cmd.AddCommand(listCmd)

}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: `List tokens.`,
	Long: `List tokens.
  
  Lists all the valid tokens for a user-workspace pair.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		response, err := w.Tokens.ListAll(ctx)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
}

// end service Tokens
