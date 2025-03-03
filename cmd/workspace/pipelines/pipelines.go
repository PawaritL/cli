// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"fmt"
	"time"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "pipelines",
	Short: `The Delta Live Tables API allows you to create, edit, delete, start, and view details about pipelines.`,
	Long: `The Delta Live Tables API allows you to create, edit, delete, start, and view
  details about pipelines.
  
  Delta Live Tables is a framework for building reliable, maintainable, and
  testable data processing pipelines. You define the transformations to perform
  on your data, and Delta Live Tables manages task orchestration, cluster
  management, monitoring, data quality, and error handling.
  
  Instead of defining your data pipelines using a series of separate Apache
  Spark tasks, Delta Live Tables manages how your data is transformed based on a
  target schema you define for each processing step. You can also enforce data
  quality with Delta Live Tables expectations. Expectations allow you to define
  expected data quality and specify how to handle records that fail those
  expectations.`,
	Annotations: map[string]string{
		"package": "pipelines",
	},
}

// start create command

var createReq pipelines.CreatePipeline
var createJson flags.JsonFlag

func init() {
	Cmd.AddCommand(createCmd)
	// TODO: short flags
	createCmd.Flags().Var(&createJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: `Create a pipeline.`,
	Long: `Create a pipeline.
  
  Creates a new data processing pipeline based on the requested configuration.
  If successful, this method returns the ID of the new pipeline.`,

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
			err = createJson.Unmarshal(&createReq)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("please provide command input in JSON format by specifying the --json flag")
		}

		response, err := w.Pipelines.Create(ctx, createReq)
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

var deleteReq pipelines.DeletePipelineRequest
var deleteJson flags.JsonFlag

func init() {
	Cmd.AddCommand(deleteCmd)
	// TODO: short flags
	deleteCmd.Flags().Var(&deleteJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var deleteCmd = &cobra.Command{
	Use:   "delete PIPELINE_ID",
	Short: `Delete a pipeline.`,
	Long: `Delete a pipeline.
  
  Deletes a pipeline.`,

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
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have ")
			}
			deleteReq.PipelineId = args[0]
		}

		err = w.Pipelines.Delete(ctx, deleteReq)
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

var getReq pipelines.GetPipelineRequest
var getJson flags.JsonFlag
var getSkipWait bool
var getTimeout time.Duration

func init() {
	Cmd.AddCommand(getCmd)

	getCmd.Flags().BoolVar(&getSkipWait, "no-wait", getSkipWait, `do not wait to reach RUNNING state`)
	getCmd.Flags().DurationVar(&getTimeout, "timeout", 20*time.Minute, `maximum amount of time to reach RUNNING state`)
	// TODO: short flags
	getCmd.Flags().Var(&getJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var getCmd = &cobra.Command{
	Use:   "get PIPELINE_ID",
	Short: `Get a pipeline.`,
	Long:  `Get a pipeline.`,

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
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have ")
			}
			getReq.PipelineId = args[0]
		}

		response, err := w.Pipelines.Get(ctx, getReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start get-update command

var getUpdateReq pipelines.GetUpdateRequest
var getUpdateJson flags.JsonFlag

func init() {
	Cmd.AddCommand(getUpdateCmd)
	// TODO: short flags
	getUpdateCmd.Flags().Var(&getUpdateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var getUpdateCmd = &cobra.Command{
	Use:   "get-update PIPELINE_ID UPDATE_ID",
	Short: `Get a pipeline update.`,
	Long: `Get a pipeline update.
  
  Gets an update from an active pipeline.`,

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
			err = getUpdateJson.Unmarshal(&getUpdateReq)
			if err != nil {
				return err
			}
		} else {
			getUpdateReq.PipelineId = args[0]
			getUpdateReq.UpdateId = args[1]
		}

		response, err := w.Pipelines.GetUpdate(ctx, getUpdateReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start list-pipeline-events command

var listPipelineEventsReq pipelines.ListPipelineEventsRequest
var listPipelineEventsJson flags.JsonFlag

func init() {
	Cmd.AddCommand(listPipelineEventsCmd)
	// TODO: short flags
	listPipelineEventsCmd.Flags().Var(&listPipelineEventsJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	listPipelineEventsCmd.Flags().StringVar(&listPipelineEventsReq.Filter, "filter", listPipelineEventsReq.Filter, `Criteria to select a subset of results, expressed using a SQL-like syntax.`)
	listPipelineEventsCmd.Flags().IntVar(&listPipelineEventsReq.MaxResults, "max-results", listPipelineEventsReq.MaxResults, `Max number of entries to return in a single page.`)
	// TODO: array: order_by
	listPipelineEventsCmd.Flags().StringVar(&listPipelineEventsReq.PageToken, "page-token", listPipelineEventsReq.PageToken, `Page token returned by previous call.`)

}

var listPipelineEventsCmd = &cobra.Command{
	Use:   "list-pipeline-events PIPELINE_ID",
	Short: `List pipeline events.`,
	Long: `List pipeline events.
  
  Retrieves events for a pipeline.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = listPipelineEventsJson.Unmarshal(&listPipelineEventsReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have ")
			}
			listPipelineEventsReq.PipelineId = args[0]
		}

		response, err := w.Pipelines.ListPipelineEventsAll(ctx, listPipelineEventsReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start list-pipelines command

var listPipelinesReq pipelines.ListPipelinesRequest
var listPipelinesJson flags.JsonFlag

func init() {
	Cmd.AddCommand(listPipelinesCmd)
	// TODO: short flags
	listPipelinesCmd.Flags().Var(&listPipelinesJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	listPipelinesCmd.Flags().StringVar(&listPipelinesReq.Filter, "filter", listPipelinesReq.Filter, `Select a subset of results based on the specified criteria.`)
	listPipelinesCmd.Flags().IntVar(&listPipelinesReq.MaxResults, "max-results", listPipelinesReq.MaxResults, `The maximum number of entries to return in a single page.`)
	// TODO: array: order_by
	listPipelinesCmd.Flags().StringVar(&listPipelinesReq.PageToken, "page-token", listPipelinesReq.PageToken, `Page token returned by previous call.`)

}

var listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: `List pipelines.`,
	Long: `List pipelines.
  
  Lists pipelines defined in the Delta Live Tables system.`,

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
			err = listPipelinesJson.Unmarshal(&listPipelinesReq)
			if err != nil {
				return err
			}
		} else {
		}

		response, err := w.Pipelines.ListPipelinesAll(ctx, listPipelinesReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start list-updates command

var listUpdatesReq pipelines.ListUpdatesRequest
var listUpdatesJson flags.JsonFlag

func init() {
	Cmd.AddCommand(listUpdatesCmd)
	// TODO: short flags
	listUpdatesCmd.Flags().Var(&listUpdatesJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	listUpdatesCmd.Flags().IntVar(&listUpdatesReq.MaxResults, "max-results", listUpdatesReq.MaxResults, `Max number of entries to return in a single page.`)
	listUpdatesCmd.Flags().StringVar(&listUpdatesReq.PageToken, "page-token", listUpdatesReq.PageToken, `Page token returned by previous call.`)
	listUpdatesCmd.Flags().StringVar(&listUpdatesReq.UntilUpdateId, "until-update-id", listUpdatesReq.UntilUpdateId, `If present, returns updates until and including this update_id.`)

}

var listUpdatesCmd = &cobra.Command{
	Use:   "list-updates PIPELINE_ID",
	Short: `List pipeline updates.`,
	Long: `List pipeline updates.
  
  List updates for an active pipeline.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = listUpdatesJson.Unmarshal(&listUpdatesReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "The pipeline to return updates for")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have the pipeline to return updates for")
			}
			listUpdatesReq.PipelineId = args[0]
		}

		response, err := w.Pipelines.ListUpdates(ctx, listUpdatesReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start reset command

var resetReq pipelines.ResetRequest
var resetJson flags.JsonFlag
var resetSkipWait bool
var resetTimeout time.Duration

func init() {
	Cmd.AddCommand(resetCmd)

	resetCmd.Flags().BoolVar(&resetSkipWait, "no-wait", resetSkipWait, `do not wait to reach RUNNING state`)
	resetCmd.Flags().DurationVar(&resetTimeout, "timeout", 20*time.Minute, `maximum amount of time to reach RUNNING state`)
	// TODO: short flags
	resetCmd.Flags().Var(&resetJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var resetCmd = &cobra.Command{
	Use:   "reset PIPELINE_ID",
	Short: `Reset a pipeline.`,
	Long: `Reset a pipeline.
  
  Resets a pipeline.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = resetJson.Unmarshal(&resetReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have ")
			}
			resetReq.PipelineId = args[0]
		}

		wait, err := w.Pipelines.Reset(ctx, resetReq)
		if err != nil {
			return err
		}
		if resetSkipWait {
			return nil
		}
		spinner := cmdio.Spinner(ctx)
		info, err := wait.OnProgress(func(i *pipelines.GetPipelineResponse) {
			statusMessage := i.Cause
			spinner <- statusMessage
		}).GetWithTimeout(resetTimeout)
		close(spinner)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, info)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start start-update command

var startUpdateReq pipelines.StartUpdate
var startUpdateJson flags.JsonFlag

func init() {
	Cmd.AddCommand(startUpdateCmd)
	// TODO: short flags
	startUpdateCmd.Flags().Var(&startUpdateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	startUpdateCmd.Flags().Var(&startUpdateReq.Cause, "cause", ``)
	startUpdateCmd.Flags().BoolVar(&startUpdateReq.FullRefresh, "full-refresh", startUpdateReq.FullRefresh, `If true, this update will reset all tables before running.`)
	// TODO: array: full_refresh_selection
	// TODO: array: refresh_selection

}

var startUpdateCmd = &cobra.Command{
	Use:   "start-update PIPELINE_ID",
	Short: `Queue a pipeline update.`,
	Long: `Queue a pipeline update.
  
  Starts or queues a pipeline update.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = startUpdateJson.Unmarshal(&startUpdateReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have ")
			}
			startUpdateReq.PipelineId = args[0]
		}

		response, err := w.Pipelines.StartUpdate(ctx, startUpdateReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start stop command

var stopReq pipelines.StopRequest
var stopJson flags.JsonFlag
var stopSkipWait bool
var stopTimeout time.Duration

func init() {
	Cmd.AddCommand(stopCmd)

	stopCmd.Flags().BoolVar(&stopSkipWait, "no-wait", stopSkipWait, `do not wait to reach IDLE state`)
	stopCmd.Flags().DurationVar(&stopTimeout, "timeout", 20*time.Minute, `maximum amount of time to reach IDLE state`)
	// TODO: short flags
	stopCmd.Flags().Var(&stopJson, "json", `either inline JSON string or @path/to/file.json with request body`)

}

var stopCmd = &cobra.Command{
	Use:   "stop PIPELINE_ID",
	Short: `Stop a pipeline.`,
	Long: `Stop a pipeline.
  
  Stops a pipeline.`,

	Annotations: map[string]string{},
	PreRunE:     root.MustWorkspaceClient,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		if cmd.Flags().Changed("json") {
			err = stopJson.Unmarshal(&stopReq)
			if err != nil {
				return err
			}
		} else {
			if len(args) == 0 {
				promptSpinner := cmdio.Spinner(ctx)
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have ")
			}
			stopReq.PipelineId = args[0]
		}

		wait, err := w.Pipelines.Stop(ctx, stopReq)
		if err != nil {
			return err
		}
		if stopSkipWait {
			return nil
		}
		spinner := cmdio.Spinner(ctx)
		info, err := wait.OnProgress(func(i *pipelines.GetPipelineResponse) {
			statusMessage := i.Cause
			spinner <- statusMessage
		}).GetWithTimeout(stopTimeout)
		close(spinner)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, info)
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// start update command

var updateReq pipelines.EditPipeline
var updateJson flags.JsonFlag

func init() {
	Cmd.AddCommand(updateCmd)
	// TODO: short flags
	updateCmd.Flags().Var(&updateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	updateCmd.Flags().BoolVar(&updateReq.AllowDuplicateNames, "allow-duplicate-names", updateReq.AllowDuplicateNames, `If false, deployment will fail if name has changed and conflicts the name of another pipeline.`)
	updateCmd.Flags().StringVar(&updateReq.Catalog, "catalog", updateReq.Catalog, `A catalog in Unity Catalog to publish data from this pipeline to.`)
	updateCmd.Flags().StringVar(&updateReq.Channel, "channel", updateReq.Channel, `DLT Release Channel that specifies which version to use.`)
	// TODO: array: clusters
	// TODO: map via StringToStringVar: configuration
	updateCmd.Flags().BoolVar(&updateReq.Continuous, "continuous", updateReq.Continuous, `Whether the pipeline is continuous or triggered.`)
	updateCmd.Flags().BoolVar(&updateReq.Development, "development", updateReq.Development, `Whether the pipeline is in Development mode.`)
	updateCmd.Flags().StringVar(&updateReq.Edition, "edition", updateReq.Edition, `Pipeline product edition.`)
	updateCmd.Flags().Int64Var(&updateReq.ExpectedLastModified, "expected-last-modified", updateReq.ExpectedLastModified, `If present, the last-modified time of the pipeline settings before the edit.`)
	// TODO: complex arg: filters
	updateCmd.Flags().StringVar(&updateReq.Id, "id", updateReq.Id, `Unique identifier for this pipeline.`)
	// TODO: array: libraries
	updateCmd.Flags().StringVar(&updateReq.Name, "name", updateReq.Name, `Friendly identifier for this pipeline.`)
	updateCmd.Flags().BoolVar(&updateReq.Photon, "photon", updateReq.Photon, `Whether Photon is enabled for this pipeline.`)
	updateCmd.Flags().StringVar(&updateReq.PipelineId, "pipeline-id", updateReq.PipelineId, `Unique identifier for this pipeline.`)
	updateCmd.Flags().BoolVar(&updateReq.Serverless, "serverless", updateReq.Serverless, `Whether serverless compute is enabled for this pipeline.`)
	updateCmd.Flags().StringVar(&updateReq.Storage, "storage", updateReq.Storage, `DBFS root directory for storing checkpoints and tables.`)
	updateCmd.Flags().StringVar(&updateReq.Target, "target", updateReq.Target, `Target schema (database) to add tables in this pipeline to.`)
	// TODO: complex arg: trigger

}

var updateCmd = &cobra.Command{
	Use:   "update PIPELINE_ID",
	Short: `Edit a pipeline.`,
	Long: `Edit a pipeline.
  
  Updates a pipeline with the supplied configuration.`,

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
				promptSpinner <- "No PIPELINE_ID argument specified. Loading names for Pipelines drop-down."
				names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
				close(promptSpinner)
				if err != nil {
					return fmt.Errorf("failed to load names for Pipelines drop-down. Please manually specify required arguments. Original error: %w", err)
				}
				id, err := cmdio.Select(ctx, names, "Unique identifier for this pipeline")
				if err != nil {
					return err
				}
				args = append(args, id)
			}
			if len(args) != 1 {
				return fmt.Errorf("expected to have unique identifier for this pipeline")
			}
			updateReq.PipelineId = args[0]
		}

		err = w.Pipelines.Update(ctx, updateReq)
		if err != nil {
			return err
		}
		return nil
	},
	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	ValidArgsFunction: cobra.NoFileCompletions,
}

// end service Pipelines
