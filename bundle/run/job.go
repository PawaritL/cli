package run

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/databricks/bricks/bundle"
	"github.com/databricks/bricks/bundle/config/resources"
	"github.com/databricks/bricks/libs/log"
	"github.com/databricks/bricks/libs/progress"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

// JobOptions defines options for running a job.
type JobOptions struct {
	dbtCommands       []string
	jarParams         []string
	notebookParams    map[string]string
	pipelineParams    map[string]string
	pythonNamedParams map[string]string
	pythonParams      []string
	sparkSubmitParams []string
	sqlParams         map[string]string
}

func (o *JobOptions) Define(fs *flag.FlagSet) {
	fs.StringSliceVar(&o.dbtCommands, "dbt-commands", nil, "A list of commands to execute for jobs with DBT tasks.")
	fs.StringSliceVar(&o.jarParams, "jar-params", nil, "A list of parameters for jobs with Spark JAR tasks.")
	fs.StringToStringVar(&o.notebookParams, "notebook-params", nil, "A map from keys to values for jobs with notebook tasks.")
	fs.StringToStringVar(&o.pipelineParams, "pipeline-params", nil, "A map from keys to values for jobs with pipeline tasks.")
	fs.StringToStringVar(&o.pythonNamedParams, "python-named-params", nil, "A map from keys to values for jobs with Python wheel tasks.")
	fs.StringSliceVar(&o.pythonParams, "python-params", nil, "A list of parameters for jobs with Python tasks.")
	fs.StringSliceVar(&o.sparkSubmitParams, "spark-submit-params", nil, "A list of parameters for jobs with Spark submit tasks.")
	fs.StringToStringVar(&o.sqlParams, "sql-params", nil, "A map from keys to values for jobs with SQL tasks.")
}

func (o *JobOptions) validatePipelineParams() (*jobs.PipelineParams, error) {
	if len(o.pipelineParams) == 0 {
		return nil, nil
	}

	var defaultErr = fmt.Errorf("job run argument --pipeline-params only supports `full_refresh=<bool>`")
	v, ok := o.pipelineParams["full_refresh"]
	if !ok {
		return nil, defaultErr
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return nil, defaultErr
	}

	pipelineParams := &jobs.PipelineParams{
		FullRefresh: b,
	}

	return pipelineParams, nil
}

func (o *JobOptions) toPayload(jobID int64) (*jobs.RunNow, error) {
	pipelineParams, err := o.validatePipelineParams()
	if err != nil {
		return nil, err
	}

	payload := &jobs.RunNow{
		JobId: jobID,

		DbtCommands:       o.dbtCommands,
		JarParams:         o.jarParams,
		NotebookParams:    o.notebookParams,
		PipelineParams:    pipelineParams,
		PythonNamedParams: o.pythonNamedParams,
		PythonParams:      o.pythonParams,
		SparkSubmitParams: o.sparkSubmitParams,
		SqlParams:         o.sqlParams,
	}

	return payload, nil
}

// Default timeout for waiting for a job run to complete.
var jobRunTimeout time.Duration = 2 * time.Hour

type jobRunner struct {
	key

	bundle *bundle.Bundle
	job    *resources.Job
}

func isFailed(task jobs.RunTask) bool {
	return task.State.LifeCycleState == jobs.RunLifeCycleStateInternalError ||
		(task.State.LifeCycleState == jobs.RunLifeCycleStateTerminated &&
			task.State.ResultState == jobs.RunResultStateFailed)
}

func isSuccess(task jobs.RunTask) bool {
	return task.State.LifeCycleState == jobs.RunLifeCycleStateTerminated &&
		task.State.ResultState == jobs.RunResultStateSuccess
}

func (r *jobRunner) logFailedTasks(ctx context.Context, runId int64) {
	w := r.bundle.WorkspaceClient()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	run, err := w.Jobs.GetRun(ctx, jobs.GetRun{
		RunId: runId,
	})
	if err != nil {
		log.Errorf(ctx, "failed to log job run. Error: %s", err)
		return
	}
	if run.State.ResultState == jobs.RunResultStateSuccess {
		return
	}
	for _, task := range run.Tasks {
		if isSuccess(task) {
			log.Infof(ctx, "task %s completed successfully", green(task.TaskKey))
		} else if isFailed(task) {
			taskInfo, err := w.Jobs.GetRunOutput(ctx, jobs.GetRunOutput{
				RunId: task.RunId,
			})
			if err != nil {
				log.Errorf(ctx, "task %s failed. Unable to fetch error trace: %s", red(task.TaskKey), err)
				continue
			}
			log.Errorf(ctx, "Task %s failed!\nError:\n%s\nTrace:\n%s",
				red(task.TaskKey), taskInfo.Error, taskInfo.ErrorTrace)
		} else {
			log.Infof(ctx, "task %s is in state %s",
				yellow(task.TaskKey), task.State.LifeCycleState)
		}
	}
}

func pullRunIdCallback(runId *int64) func(info *retries.Info[jobs.Run]) {
	return func(info *retries.Info[jobs.Run]) {
		i := info.Info
		if i == nil {
			return
		}

		if *runId == 0 {
			*runId = i.RunId
		}
	}
}

func logDebugCallback(ctx context.Context, runId *int64) func(info *retries.Info[jobs.Run]) {
	var prevState *jobs.RunState
	return func(info *retries.Info[jobs.Run]) {
		i := info.Info
		if i == nil {
			return
		}

		state := i.State
		if state == nil {
			return
		}

		// Log the job run URL as soon as it is available.
		if prevState == nil {
			log.Infof(ctx, "Run available at %s", info.Info.RunPageUrl)
		}
		if prevState == nil || prevState.LifeCycleState != state.LifeCycleState {
			log.Infof(ctx, "Run status: %s", info.Info.State.LifeCycleState)
			prevState = state
		}
	}
}

func logProgressCallback(ctx context.Context, progressLogger *progress.Logger) func(info *retries.Info[jobs.Run]) {
	var prevState *jobs.RunState
	return func(info *retries.Info[jobs.Run]) {
		i := info.Info
		if i == nil {
			return
		}

		state := i.State
		if state == nil {
			return
		}

		if prevState != nil && prevState.LifeCycleState == state.LifeCycleState &&
			prevState.ResultState == state.ResultState {
			return
		} else {
			prevState = state
		}

		event := &JobProgressEvent{
			Timestamp:  time.Now(),
			JobId:      i.JobId,
			RunId:      i.RunId,
			RunName:    i.RunName,
			State:      *i.State,
			RunPageURL: i.RunPageUrl,
		}

		// log progress events to stderr
		progressLogger.Log(event)

		// log progress events in using the default logger
		log.Infof(ctx, event.String())
	}
}

func (r *jobRunner) Run(ctx context.Context, opts *Options) (RunOutput, error) {
	jobID, err := strconv.ParseInt(r.job.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("job ID is not an integer: %s", r.job.ID)
	}

	runId := new(int64)

	// construct request payload from cmd line flags args
	req, err := opts.Job.toPayload(jobID)
	if err != nil {
		return nil, err
	}

	// Include resource key in logger.
	ctx = log.NewContext(ctx, log.GetLogger(ctx).With("resource", r.Key()))

	w := r.bundle.WorkspaceClient()

	// gets the run id from inside Jobs.RunNowAndWait
	pullRunId := pullRunIdCallback(runId)

	// callback to log status updates to the universal log destination.
	// Called on every poll request
	logDebug := logDebugCallback(ctx, runId)

	// callback to log progress events. Called on every poll request
	progressLogger, ok := progress.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no progress logger found")
	}
	logProgress := logProgressCallback(ctx, progressLogger)

	run, err := w.Jobs.RunNowAndWait(ctx, *req,
		retries.Timeout[jobs.Run](jobRunTimeout), pullRunId, logDebug, logProgress)
	if err != nil && runId != nil {
		r.logFailedTasks(ctx, *runId)
	}
	if err != nil {
		return nil, err
	}
	if run.State.LifeCycleState == jobs.RunLifeCycleStateSkipped {
		log.Infof(ctx, "Run was skipped!")
		return nil, fmt.Errorf("run skipped: %s", run.State.StateMessage)
	}

	switch run.State.ResultState {
	// The run was canceled at user request.
	case jobs.RunResultStateCanceled:
		log.Infof(ctx, "Run was cancelled!")
		return nil, fmt.Errorf("run canceled: %s", run.State.StateMessage)

	// The task completed with an error.
	case jobs.RunResultStateFailed:
		log.Infof(ctx, "Run has failed!")
		return nil, fmt.Errorf("run failed: %s", run.State.StateMessage)

	// The task completed successfully.
	case jobs.RunResultStateSuccess:
		log.Infof(ctx, "Run has completed successfully!")
		return getJobOutput(ctx, r.bundle.WorkspaceClient(), *runId)

	// The run was stopped after reaching the timeout.
	case jobs.RunResultStateTimedout:
		log.Infof(ctx, "Run has timed out!")
		return nil, fmt.Errorf("run timed out: %s", run.State.StateMessage)
	}

	return nil, err
}
