// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package model_registry

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "model-registry",
	Short: `MLflow Model Registry is a centralized model repository and a UI and set of APIs that enable you to manage the full lifecycle of MLflow Models.`,
	Long: `MLflow Model Registry is a centralized model repository and a UI and set of
  APIs that enable you to manage the full lifecycle of MLflow Models.`,
	Annotations: map[string]string{
		"package": "ml",
	},
}

// start approve-transition-request command

var approveTransitionRequestReq ml.ApproveTransitionRequest
var approveTransitionRequestJson flags.JsonFlag

func init() {
	Cmd.AddCommand(approveTransitionRequestCmd)
	// TODO: short flags
	approveTransitionRequestCmd.Flags().Var(&approveTransitionRequestJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	approveTransitionRequestCmd.Flags().StringVar(&approveTransitionRequestReq.Comment, "comment", approveTransitionRequestReq.Comment, `User-provided comment on the action.`)

}

var approveTransitionRequestCmd = &cobra.Command{
	Use:   "approve-transition-request NAME VERSION STAGE ARCHIVE_EXISTING_VERSIONS",
	Short: `Approve transition request.`,
	Long: `Approve transition request.
  
  Approves a model version stage transition request.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(4)
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
			err = approveTransitionRequestJson.Unmarshal(&approveTransitionRequestReq)
			if err != nil {
				return err
			}
		} else {
			approveTransitionRequestReq.Name = args[0]
			approveTransitionRequestReq.Version = args[1]
			_, err = fmt.Sscan(args[2], &approveTransitionRequestReq.Stage)
			if err != nil {
				return fmt.Errorf("invalid STAGE: %s", args[2])
			}
			_, err = fmt.Sscan(args[3], &approveTransitionRequestReq.ArchiveExistingVersions)
			if err != nil {
				return fmt.Errorf("invalid ARCHIVE_EXISTING_VERSIONS: %s", args[3])
			}
		}

		response, err := w.ModelRegistry.ApproveTransitionRequest(ctx, approveTransitionRequestReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start create-comment command

var createCommentReq ml.CreateComment
var createCommentJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createCommentCmd)
	// TODO: short flags
	createCommentCmd.Flags().Var(&createCommentJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var createCommentCmd = &cobra.Command{
	Use:   "create-comment NAME VERSION COMMENT",
	Short: `Post a comment.`,
	Long: `Post a comment.
  
  Posts a comment on a model version. A comment can be submitted either by a
  user or programmatically to display relevant information about the model. For
  example, test results or deployment errors.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(3)
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
			err = createCommentJson.Unmarshal(&createCommentReq)
			if err != nil {
				return err
			}
		} else {
			createCommentReq.Name = args[0]
			createCommentReq.Version = args[1]
			createCommentReq.Comment = args[2]
		}

		response, err := w.ModelRegistry.CreateComment(ctx, createCommentReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start create-model command

var createModelReq ml.CreateModelRequest
var createModelJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createModelCmd)
	// TODO: short flags
	createModelCmd.Flags().Var(&createModelJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	createModelCmd.Flags().StringVar(&createModelReq.Description, "description", createModelReq.Description, `Optional description for registered model.`)
	// TODO: array: tags

}

var createModelCmd = &cobra.Command{
	Use:   "create-model NAME",
	Short: `Create a model.`,
	Long: `Create a model.
  
  Creates a new registered model with the name specified in the request body.
  
  Throws RESOURCE_ALREADY_EXISTS if a registered model with the given name
  exists.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = createModelJson.Unmarshal(&createModelReq)
			if err != nil {
				return err
			}
		} else {
			createModelReq.Name = args[0]
		}

		response, err := w.ModelRegistry.CreateModel(ctx, createModelReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start create-model-version command

var createModelVersionReq ml.CreateModelVersionRequest
var createModelVersionJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createModelVersionCmd)
	// TODO: short flags
	createModelVersionCmd.Flags().Var(&createModelVersionJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	createModelVersionCmd.Flags().StringVar(&createModelVersionReq.Description, "description", createModelVersionReq.Description, `Optional description for model version.`)
	createModelVersionCmd.Flags().StringVar(&createModelVersionReq.RunId, "run-id", createModelVersionReq.RunId, `MLflow run ID for correlation, if source was generated by an experiment run in MLflow tracking server.`)
	createModelVersionCmd.Flags().StringVar(&createModelVersionReq.RunLink, "run-link", createModelVersionReq.RunLink, `MLflow run link - this is the exact link of the run that generated this model version, potentially hosted at another instance of MLflow.`)
	// TODO: array: tags

}

var createModelVersionCmd = &cobra.Command{
	Use:   "create-model-version NAME SOURCE",
	Short: `Create a model version.`,
	Long: `Create a model version.
  
  Creates a model version.`,

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
			err = createModelVersionJson.Unmarshal(&createModelVersionReq)
			if err != nil {
				return err
			}
		} else {
			createModelVersionReq.Name = args[0]
			createModelVersionReq.Source = args[1]
		}

		response, err := w.ModelRegistry.CreateModelVersion(ctx, createModelVersionReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start create-transition-request command

var createTransitionRequestReq ml.CreateTransitionRequest
var createTransitionRequestJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createTransitionRequestCmd)
	// TODO: short flags
	createTransitionRequestCmd.Flags().Var(&createTransitionRequestJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	createTransitionRequestCmd.Flags().StringVar(&createTransitionRequestReq.Comment, "comment", createTransitionRequestReq.Comment, `User-provided comment on the action.`)

}

var createTransitionRequestCmd = &cobra.Command{
	Use:   "create-transition-request NAME VERSION STAGE",
	Short: `Make a transition request.`,
	Long: `Make a transition request.
  
  Creates a model version stage transition request.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(3)
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
			err = createTransitionRequestJson.Unmarshal(&createTransitionRequestReq)
			if err != nil {
				return err
			}
		} else {
			createTransitionRequestReq.Name = args[0]
			createTransitionRequestReq.Version = args[1]
			_, err = fmt.Sscan(args[2], &createTransitionRequestReq.Stage)
			if err != nil {
				return fmt.Errorf("invalid STAGE: %s", args[2])
			}
		}

		response, err := w.ModelRegistry.CreateTransitionRequest(ctx, createTransitionRequestReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start create-webhook command

var createWebhookReq ml.CreateRegistryWebhook
var createWebhookJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createWebhookCmd)
	// TODO: short flags
	createWebhookCmd.Flags().Var(&createWebhookJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	createWebhookCmd.Flags().StringVar(&createWebhookReq.Description, "description", createWebhookReq.Description, `User-specified description for the webhook.`)
	// TODO: complex arg: http_url_spec
	// TODO: complex arg: job_spec
	createWebhookCmd.Flags().StringVar(&createWebhookReq.ModelName, "model-name", createWebhookReq.ModelName, `Name of the model whose events would trigger this webhook.`)
	createWebhookCmd.Flags().Var(&createWebhookReq.Status, "status", `This describes an enum.`)

}

var createWebhookCmd = &cobra.Command{
	Use:   "create-webhook",
	Short: `Create a webhook.`,
	Long: `Create a webhook.
  
  **NOTE**: This endpoint is in Public Preview.
  
  Creates a registry webhook.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = createWebhookJson.Unmarshal(&createWebhookReq)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("please provide command input in JSON format by specifying the --json flag")
		}

		response, err := w.ModelRegistry.CreateWebhook(ctx, createWebhookReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete-comment command

var deleteCommentReq ml.DeleteCommentRequest
var deleteCommentJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteCommentCmd)
	// TODO: short flags
	deleteCommentCmd.Flags().Var(&deleteCommentJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var deleteCommentCmd = &cobra.Command{
	Use:   "delete-comment ID",
	Short: `Delete a comment.`,
	Long: `Delete a comment.
  
  Deletes a comment on a model version.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = deleteCommentJson.Unmarshal(&deleteCommentReq)
			if err != nil {
				return err
			}
		} else {
			deleteCommentReq.Id = args[0]
		}

		err = w.ModelRegistry.DeleteComment(ctx, deleteCommentReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete-model command

var deleteModelReq ml.DeleteModelRequest
var deleteModelJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteModelCmd)
	// TODO: short flags
	deleteModelCmd.Flags().Var(&deleteModelJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var deleteModelCmd = &cobra.Command{
	Use:   "delete-model NAME",
	Short: `Delete a model.`,
	Long: `Delete a model.
  
  Deletes a registered model.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = deleteModelJson.Unmarshal(&deleteModelReq)
			if err != nil {
				return err
			}
		} else {
			deleteModelReq.Name = args[0]
		}

		err = w.ModelRegistry.DeleteModel(ctx, deleteModelReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete-model-tag command

var deleteModelTagReq ml.DeleteModelTagRequest
var deleteModelTagJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteModelTagCmd)
	// TODO: short flags
	deleteModelTagCmd.Flags().Var(&deleteModelTagJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var deleteModelTagCmd = &cobra.Command{
	Use:   "delete-model-tag NAME KEY",
	Short: `Delete a model tag.`,
	Long: `Delete a model tag.
  
  Deletes the tag for a registered model.`,

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
			err = deleteModelTagJson.Unmarshal(&deleteModelTagReq)
			if err != nil {
				return err
			}
		} else {
			deleteModelTagReq.Name = args[0]
			deleteModelTagReq.Key = args[1]
		}

		err = w.ModelRegistry.DeleteModelTag(ctx, deleteModelTagReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete-model-version command

var deleteModelVersionReq ml.DeleteModelVersionRequest
var deleteModelVersionJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteModelVersionCmd)
	// TODO: short flags
	deleteModelVersionCmd.Flags().Var(&deleteModelVersionJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var deleteModelVersionCmd = &cobra.Command{
	Use:   "delete-model-version NAME VERSION",
	Short: `Delete a model version.`,
	Long: `Delete a model version.
  
  Deletes a model version.`,

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
			err = deleteModelVersionJson.Unmarshal(&deleteModelVersionReq)
			if err != nil {
				return err
			}
		} else {
			deleteModelVersionReq.Name = args[0]
			deleteModelVersionReq.Version = args[1]
		}

		err = w.ModelRegistry.DeleteModelVersion(ctx, deleteModelVersionReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete-model-version-tag command

var deleteModelVersionTagReq ml.DeleteModelVersionTagRequest
var deleteModelVersionTagJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteModelVersionTagCmd)
	// TODO: short flags
	deleteModelVersionTagCmd.Flags().Var(&deleteModelVersionTagJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var deleteModelVersionTagCmd = &cobra.Command{
	Use:   "delete-model-version-tag NAME VERSION KEY",
	Short: `Delete a model version tag.`,
	Long: `Delete a model version tag.
  
  Deletes a model version tag.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(3)
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
			err = deleteModelVersionTagJson.Unmarshal(&deleteModelVersionTagReq)
			if err != nil {
				return err
			}
		} else {
			deleteModelVersionTagReq.Name = args[0]
			deleteModelVersionTagReq.Version = args[1]
			deleteModelVersionTagReq.Key = args[2]
		}

		err = w.ModelRegistry.DeleteModelVersionTag(ctx, deleteModelVersionTagReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete-transition-request command

var deleteTransitionRequestReq ml.DeleteTransitionRequestRequest
var deleteTransitionRequestJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteTransitionRequestCmd)
	// TODO: short flags
	deleteTransitionRequestCmd.Flags().Var(&deleteTransitionRequestJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	deleteTransitionRequestCmd.Flags().StringVar(&deleteTransitionRequestReq.Comment, "comment", deleteTransitionRequestReq.Comment, `User-provided comment on the action.`)

}

var deleteTransitionRequestCmd = &cobra.Command{
	Use:   "delete-transition-request NAME VERSION STAGE CREATOR",
	Short: `Delete a transition request.`,
	Long: `Delete a transition request.
  
  Cancels a model version stage transition request.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(4)
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
			err = deleteTransitionRequestJson.Unmarshal(&deleteTransitionRequestReq)
			if err != nil {
				return err
			}
		} else {
			deleteTransitionRequestReq.Name = args[0]
			deleteTransitionRequestReq.Version = args[1]
			_, err = fmt.Sscan(args[2], &deleteTransitionRequestReq.Stage)
			if err != nil {
				return fmt.Errorf("invalid STAGE: %s", args[2])
			}
			deleteTransitionRequestReq.Creator = args[3]
		}

		err = w.ModelRegistry.DeleteTransitionRequest(ctx, deleteTransitionRequestReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start delete-webhook command

var deleteWebhookReq ml.DeleteWebhookRequest
var deleteWebhookJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteWebhookCmd)
	// TODO: short flags
	deleteWebhookCmd.Flags().Var(&deleteWebhookJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	deleteWebhookCmd.Flags().StringVar(&deleteWebhookReq.Id, "id", deleteWebhookReq.Id, `Webhook ID required to delete a registry webhook.`)

}

var deleteWebhookCmd = &cobra.Command{
	Use:   "delete-webhook",
	Short: `Delete a webhook.`,
	Long: `Delete a webhook.
  
  **NOTE:** This endpoint is in Public Preview.
  
  Deletes a registry webhook.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(0)
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
			err = deleteWebhookJson.Unmarshal(&deleteWebhookReq)
			if err != nil {
				return err
			}
		} else {
		}

		err = w.ModelRegistry.DeleteWebhook(ctx, deleteWebhookReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start get-latest-versions command

var getLatestVersionsReq ml.GetLatestVersionsRequest
var getLatestVersionsJson flags.JsonFlag

func init() {
	Cmd.AddCommand(getLatestVersionsCmd)
	// TODO: short flags
	getLatestVersionsCmd.Flags().Var(&getLatestVersionsJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	// TODO: array: stages

}

var getLatestVersionsCmd = &cobra.Command{
	Use:   "get-latest-versions NAME",
	Short: `Get the latest version.`,
	Long: `Get the latest version.
  
  Gets the latest version of a registered model.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = getLatestVersionsJson.Unmarshal(&getLatestVersionsReq)
			if err != nil {
				return err
			}
		} else {
			getLatestVersionsReq.Name = args[0]
		}

		response, err := w.ModelRegistry.GetLatestVersionsAll(ctx, getLatestVersionsReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start get-model command

var getModelReq ml.GetModelRequest
var getModelJson flags.JsonFlag

func init() {
	Cmd.AddCommand(getModelCmd)
	// TODO: short flags
	getModelCmd.Flags().Var(&getModelJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var getModelCmd = &cobra.Command{
	Use:   "get-model NAME",
	Short: `Get model.`,
	Long: `Get model.
  
  Get the details of a model. This is a Databricks workspace version of the
  [MLflow endpoint] that also returns the model's Databricks workspace ID and
  the permission level of the requesting user on the model.
  
  [MLflow endpoint]: https://www.mlflow.org/docs/latest/rest-api.html#get-registeredmodel`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = getModelJson.Unmarshal(&getModelReq)
			if err != nil {
				return err
			}
		} else {
			getModelReq.Name = args[0]
		}

		response, err := w.ModelRegistry.GetModel(ctx, getModelReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start get-model-version command

var getModelVersionReq ml.GetModelVersionRequest
var getModelVersionJson flags.JsonFlag

func init() {
	Cmd.AddCommand(getModelVersionCmd)
	// TODO: short flags
	getModelVersionCmd.Flags().Var(&getModelVersionJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var getModelVersionCmd = &cobra.Command{
	Use:   "get-model-version NAME VERSION",
	Short: `Get a model version.`,
	Long: `Get a model version.
  
  Get a model version.`,

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
			err = getModelVersionJson.Unmarshal(&getModelVersionReq)
			if err != nil {
				return err
			}
		} else {
			getModelVersionReq.Name = args[0]
			getModelVersionReq.Version = args[1]
		}

		response, err := w.ModelRegistry.GetModelVersion(ctx, getModelVersionReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start get-model-version-download-uri command

var getModelVersionDownloadUriReq ml.GetModelVersionDownloadUriRequest
var getModelVersionDownloadUriJson flags.JsonFlag

func init() {
	Cmd.AddCommand(getModelVersionDownloadUriCmd)
	// TODO: short flags
	getModelVersionDownloadUriCmd.Flags().Var(&getModelVersionDownloadUriJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var getModelVersionDownloadUriCmd = &cobra.Command{
	Use:   "get-model-version-download-uri NAME VERSION",
	Short: `Get a model version URI.`,
	Long: `Get a model version URI.
  
  Gets a URI to download the model version.`,

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
			err = getModelVersionDownloadUriJson.Unmarshal(&getModelVersionDownloadUriReq)
			if err != nil {
				return err
			}
		} else {
			getModelVersionDownloadUriReq.Name = args[0]
			getModelVersionDownloadUriReq.Version = args[1]
		}

		response, err := w.ModelRegistry.GetModelVersionDownloadUri(ctx, getModelVersionDownloadUriReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start list-models command

var listModelsReq ml.ListModelsRequest
var listModelsJson flags.JsonFlag

func init() {
	Cmd.AddCommand(listModelsCmd)
	// TODO: short flags
	listModelsCmd.Flags().Var(&listModelsJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	listModelsCmd.Flags().IntVar(&listModelsReq.MaxResults, "max-results", listModelsReq.MaxResults, `Maximum number of registered models desired.`)
	listModelsCmd.Flags().StringVar(&listModelsReq.PageToken, "page-token", listModelsReq.PageToken, `Pagination token to go to the next page based on a previous query.`)

}

var listModelsCmd = &cobra.Command{
	Use:   "list-models",
	Short: `List models.`,
	Long: `List models.
  
  Lists all available registered models, up to the limit specified in
  __max_results__.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(0)
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
			err = listModelsJson.Unmarshal(&listModelsReq)
			if err != nil {
				return err
			}
		} else {
		}

		response, err := w.ModelRegistry.ListModelsAll(ctx, listModelsReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start list-transition-requests command

var listTransitionRequestsReq ml.ListTransitionRequestsRequest
var listTransitionRequestsJson flags.JsonFlag

func init() {
	Cmd.AddCommand(listTransitionRequestsCmd)
	// TODO: short flags
	listTransitionRequestsCmd.Flags().Var(&listTransitionRequestsJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var listTransitionRequestsCmd = &cobra.Command{
	Use:   "list-transition-requests NAME VERSION",
	Short: `List transition requests.`,
	Long: `List transition requests.
  
  Gets a list of all open stage transition requests for the model version.`,

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
			err = listTransitionRequestsJson.Unmarshal(&listTransitionRequestsReq)
			if err != nil {
				return err
			}
		} else {
			listTransitionRequestsReq.Name = args[0]
			listTransitionRequestsReq.Version = args[1]
		}

		response, err := w.ModelRegistry.ListTransitionRequestsAll(ctx, listTransitionRequestsReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start list-webhooks command

var listWebhooksReq ml.ListWebhooksRequest
var listWebhooksJson flags.JsonFlag

func init() {
	Cmd.AddCommand(listWebhooksCmd)
	// TODO: short flags
	listWebhooksCmd.Flags().Var(&listWebhooksJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	// TODO: array: events
	listWebhooksCmd.Flags().StringVar(&listWebhooksReq.ModelName, "model-name", listWebhooksReq.ModelName, `If not specified, all webhooks associated with the specified events are listed, regardless of their associated model.`)
	listWebhooksCmd.Flags().StringVar(&listWebhooksReq.PageToken, "page-token", listWebhooksReq.PageToken, `Token indicating the page of artifact results to fetch.`)

}

var listWebhooksCmd = &cobra.Command{
	Use:   "list-webhooks",
	Short: `List registry webhooks.`,
	Long: `List registry webhooks.
  
  **NOTE:** This endpoint is in Public Preview.
  
  Lists all registry webhooks.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(0)
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
			err = listWebhooksJson.Unmarshal(&listWebhooksReq)
			if err != nil {
				return err
			}
		} else {
		}

		response, err := w.ModelRegistry.ListWebhooksAll(ctx, listWebhooksReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start reject-transition-request command

var rejectTransitionRequestReq ml.RejectTransitionRequest
var rejectTransitionRequestJson flags.JsonFlag

func init() {
	Cmd.AddCommand(rejectTransitionRequestCmd)
	// TODO: short flags
	rejectTransitionRequestCmd.Flags().Var(&rejectTransitionRequestJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	rejectTransitionRequestCmd.Flags().StringVar(&rejectTransitionRequestReq.Comment, "comment", rejectTransitionRequestReq.Comment, `User-provided comment on the action.`)

}

var rejectTransitionRequestCmd = &cobra.Command{
	Use:   "reject-transition-request NAME VERSION STAGE",
	Short: `Reject a transition request.`,
	Long: `Reject a transition request.
  
  Rejects a model version stage transition request.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(3)
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
			err = rejectTransitionRequestJson.Unmarshal(&rejectTransitionRequestReq)
			if err != nil {
				return err
			}
		} else {
			rejectTransitionRequestReq.Name = args[0]
			rejectTransitionRequestReq.Version = args[1]
			_, err = fmt.Sscan(args[2], &rejectTransitionRequestReq.Stage)
			if err != nil {
				return fmt.Errorf("invalid STAGE: %s", args[2])
			}
		}

		response, err := w.ModelRegistry.RejectTransitionRequest(ctx, rejectTransitionRequestReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start rename-model command

var renameModelReq ml.RenameModelRequest
var renameModelJson flags.JsonFlag

func init() {
	Cmd.AddCommand(renameModelCmd)
	// TODO: short flags
	renameModelCmd.Flags().Var(&renameModelJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	renameModelCmd.Flags().StringVar(&renameModelReq.NewName, "new-name", renameModelReq.NewName, `If provided, updates the name for this registered_model.`)

}

var renameModelCmd = &cobra.Command{
	Use:   "rename-model NAME",
	Short: `Rename a model.`,
	Long: `Rename a model.
  
  Renames a registered model.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = renameModelJson.Unmarshal(&renameModelReq)
			if err != nil {
				return err
			}
		} else {
			renameModelReq.Name = args[0]
		}

		response, err := w.ModelRegistry.RenameModel(ctx, renameModelReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start search-model-versions command

var searchModelVersionsReq ml.SearchModelVersionsRequest
var searchModelVersionsJson flags.JsonFlag

func init() {
	Cmd.AddCommand(searchModelVersionsCmd)
	// TODO: short flags
	searchModelVersionsCmd.Flags().Var(&searchModelVersionsJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	searchModelVersionsCmd.Flags().StringVar(&searchModelVersionsReq.Filter, "filter", searchModelVersionsReq.Filter, `String filter condition, like "name='my-model-name'".`)
	searchModelVersionsCmd.Flags().IntVar(&searchModelVersionsReq.MaxResults, "max-results", searchModelVersionsReq.MaxResults, `Maximum number of models desired.`)
	// TODO: array: order_by
	searchModelVersionsCmd.Flags().StringVar(&searchModelVersionsReq.PageToken, "page-token", searchModelVersionsReq.PageToken, `Pagination token to go to next page based on previous search query.`)

}

var searchModelVersionsCmd = &cobra.Command{
	Use:   "search-model-versions",
	Short: `Searches model versions.`,
	Long: `Searches model versions.
  
  Searches for specific model versions based on the supplied __filter__.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(0)
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
			err = searchModelVersionsJson.Unmarshal(&searchModelVersionsReq)
			if err != nil {
				return err
			}
		} else {
		}

		response, err := w.ModelRegistry.SearchModelVersionsAll(ctx, searchModelVersionsReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start search-models command

var searchModelsReq ml.SearchModelsRequest
var searchModelsJson flags.JsonFlag

func init() {
	Cmd.AddCommand(searchModelsCmd)
	// TODO: short flags
	searchModelsCmd.Flags().Var(&searchModelsJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	searchModelsCmd.Flags().StringVar(&searchModelsReq.Filter, "filter", searchModelsReq.Filter, `String filter condition, like "name LIKE 'my-model-name'".`)
	searchModelsCmd.Flags().IntVar(&searchModelsReq.MaxResults, "max-results", searchModelsReq.MaxResults, `Maximum number of models desired.`)
	// TODO: array: order_by
	searchModelsCmd.Flags().StringVar(&searchModelsReq.PageToken, "page-token", searchModelsReq.PageToken, `Pagination token to go to the next page based on a previous search query.`)

}

var searchModelsCmd = &cobra.Command{
	Use:   "search-models",
	Short: `Search models.`,
	Long: `Search models.
  
  Search for registered models based on the specified __filter__.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(0)
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
			err = searchModelsJson.Unmarshal(&searchModelsReq)
			if err != nil {
				return err
			}
		} else {
		}

		response, err := w.ModelRegistry.SearchModelsAll(ctx, searchModelsReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start set-model-tag command

var setModelTagReq ml.SetModelTagRequest
var setModelTagJson flags.JsonFlag

func init() {
	Cmd.AddCommand(setModelTagCmd)
	// TODO: short flags
	setModelTagCmd.Flags().Var(&setModelTagJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var setModelTagCmd = &cobra.Command{
	Use:   "set-model-tag NAME KEY VALUE",
	Short: `Set a tag.`,
	Long: `Set a tag.
  
  Sets a tag on a registered model.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(3)
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
			err = setModelTagJson.Unmarshal(&setModelTagReq)
			if err != nil {
				return err
			}
		} else {
			setModelTagReq.Name = args[0]
			setModelTagReq.Key = args[1]
			setModelTagReq.Value = args[2]
		}

		err = w.ModelRegistry.SetModelTag(ctx, setModelTagReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start set-model-version-tag command

var setModelVersionTagReq ml.SetModelVersionTagRequest
var setModelVersionTagJson flags.JsonFlag

func init() {
	Cmd.AddCommand(setModelVersionTagCmd)
	// TODO: short flags
	setModelVersionTagCmd.Flags().Var(&setModelVersionTagJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var setModelVersionTagCmd = &cobra.Command{
	Use:   "set-model-version-tag NAME VERSION KEY VALUE",
	Short: `Set a version tag.`,
	Long: `Set a version tag.
  
  Sets a model version tag.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(4)
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
			err = setModelVersionTagJson.Unmarshal(&setModelVersionTagReq)
			if err != nil {
				return err
			}
		} else {
			setModelVersionTagReq.Name = args[0]
			setModelVersionTagReq.Version = args[1]
			setModelVersionTagReq.Key = args[2]
			setModelVersionTagReq.Value = args[3]
		}

		err = w.ModelRegistry.SetModelVersionTag(ctx, setModelVersionTagReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start test-registry-webhook command

var testRegistryWebhookReq ml.TestRegistryWebhookRequest
var testRegistryWebhookJson flags.JsonFlag

func init() {
	Cmd.AddCommand(testRegistryWebhookCmd)
	// TODO: short flags
	testRegistryWebhookCmd.Flags().Var(&testRegistryWebhookJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	testRegistryWebhookCmd.Flags().Var(&testRegistryWebhookReq.Event, "event", `If event is specified, the test trigger uses the specified event.`)

}

var testRegistryWebhookCmd = &cobra.Command{
	Use:   "test-registry-webhook ID",
	Short: `Test a webhook.`,
	Long: `Test a webhook.
  
  **NOTE:** This endpoint is in Public Preview.
  
  Tests a registry webhook.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = testRegistryWebhookJson.Unmarshal(&testRegistryWebhookReq)
			if err != nil {
				return err
			}
		} else {
			testRegistryWebhookReq.Id = args[0]
		}

		response, err := w.ModelRegistry.TestRegistryWebhook(ctx, testRegistryWebhookReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start transition-stage command

var transitionStageReq ml.TransitionModelVersionStageDatabricks
var transitionStageJson flags.JsonFlag

func init() {
	Cmd.AddCommand(transitionStageCmd)
	// TODO: short flags
	transitionStageCmd.Flags().Var(&transitionStageJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	transitionStageCmd.Flags().StringVar(&transitionStageReq.Comment, "comment", transitionStageReq.Comment, `User-provided comment on the action.`)

}

var transitionStageCmd = &cobra.Command{
	Use:   "transition-stage NAME VERSION STAGE ARCHIVE_EXISTING_VERSIONS",
	Short: `Transition a stage.`,
	Long: `Transition a stage.
  
  Transition a model version's stage. This is a Databricks workspace version of
  the [MLflow endpoint] that also accepts a comment associated with the
  transition to be recorded.",
  
  [MLflow endpoint]: https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(4)
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
			err = transitionStageJson.Unmarshal(&transitionStageReq)
			if err != nil {
				return err
			}
		} else {
			transitionStageReq.Name = args[0]
			transitionStageReq.Version = args[1]
			_, err = fmt.Sscan(args[2], &transitionStageReq.Stage)
			if err != nil {
				return fmt.Errorf("invalid STAGE: %s", args[2])
			}
			_, err = fmt.Sscan(args[3], &transitionStageReq.ArchiveExistingVersions)
			if err != nil {
				return fmt.Errorf("invalid ARCHIVE_EXISTING_VERSIONS: %s", args[3])
			}
		}

		response, err := w.ModelRegistry.TransitionStage(ctx, transitionStageReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start update-comment command

var updateCommentReq ml.UpdateComment
var updateCommentJson flags.JsonFlag

func init() {
	Cmd.AddCommand(updateCommentCmd)
	// TODO: short flags
	updateCommentCmd.Flags().Var(&updateCommentJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var updateCommentCmd = &cobra.Command{
	Use:   "update-comment ID COMMENT",
	Short: `Update a comment.`,
	Long: `Update a comment.
  
  Post an edit to a comment on a model version.`,

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
			err = updateCommentJson.Unmarshal(&updateCommentReq)
			if err != nil {
				return err
			}
		} else {
			updateCommentReq.Id = args[0]
			updateCommentReq.Comment = args[1]
		}

		response, err := w.ModelRegistry.UpdateComment(ctx, updateCommentReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start update-model command

var updateModelReq ml.UpdateModelRequest
var updateModelJson flags.JsonFlag

func init() {
	Cmd.AddCommand(updateModelCmd)
	// TODO: short flags
	updateModelCmd.Flags().Var(&updateModelJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	updateModelCmd.Flags().StringVar(&updateModelReq.Description, "description", updateModelReq.Description, `If provided, updates the description for this registered_model.`)

}

var updateModelCmd = &cobra.Command{
	Use:   "update-model NAME",
	Short: `Update model.`,
	Long: `Update model.
  
  Updates a registered model.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = updateModelJson.Unmarshal(&updateModelReq)
			if err != nil {
				return err
			}
		} else {
			updateModelReq.Name = args[0]
		}

		err = w.ModelRegistry.UpdateModel(ctx, updateModelReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start update-model-version command

var updateModelVersionReq ml.UpdateModelVersionRequest
var updateModelVersionJson flags.JsonFlag

func init() {
	Cmd.AddCommand(updateModelVersionCmd)
	// TODO: short flags
	updateModelVersionCmd.Flags().Var(&updateModelVersionJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	updateModelVersionCmd.Flags().StringVar(&updateModelVersionReq.Description, "description", updateModelVersionReq.Description, `If provided, updates the description for this registered_model.`)

}

var updateModelVersionCmd = &cobra.Command{
	Use:   "update-model-version NAME VERSION",
	Short: `Update model version.`,
	Long: `Update model version.
  
  Updates the model version.`,

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
			err = updateModelVersionJson.Unmarshal(&updateModelVersionReq)
			if err != nil {
				return err
			}
		} else {
			updateModelVersionReq.Name = args[0]
			updateModelVersionReq.Version = args[1]
		}

		err = w.ModelRegistry.UpdateModelVersion(ctx, updateModelVersionReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start update-webhook command

var updateWebhookReq ml.UpdateRegistryWebhook
var updateWebhookJson flags.JsonFlag

func init() {
	Cmd.AddCommand(updateWebhookCmd)
	// TODO: short flags
	updateWebhookCmd.Flags().Var(&updateWebhookJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	updateWebhookCmd.Flags().StringVar(&updateWebhookReq.Description, "description", updateWebhookReq.Description, `User-specified description for the webhook.`)
	// TODO: array: events
	// TODO: complex arg: http_url_spec
	// TODO: complex arg: job_spec
	updateWebhookCmd.Flags().Var(&updateWebhookReq.Status, "status", `This describes an enum.`)

}

var updateWebhookCmd = &cobra.Command{
	Use:   "update-webhook ID",
	Short: `Update a webhook.`,
	Long: `Update a webhook.
  
  **NOTE:** This endpoint is in Public Preview.
  
  Updates a registry webhook.`,

	Annotations: map[string]string{},
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
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
			err = updateWebhookJson.Unmarshal(&updateWebhookReq)
			if err != nil {
				return err
			}
		} else {
			updateWebhookReq.Id = args[0]
		}

		err = w.ModelRegistry.UpdateWebhook(ctx, updateWebhookReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// end service ModelRegistry
