package root

import (
	"context"
	"testing"

	"github.com/databricks/bricks/libs/flags"
	"github.com/databricks/bricks/libs/progress"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitializeErrorOnIncompatibleConfig(t *testing.T) {
	logLevel.Set("info")
	logFile.Set("stderr")
	progressFormat.Set("inplace")
	_, err := initializeProgressLogger(context.Background())
	assert.ErrorContains(t, err, "inplace progress logging cannot be used when log-file is stderr")
}

func TestNoErrorOnDisabledLogLevel(t *testing.T) {
	logLevel.Set("disabled")
	logFile.Set("stderr")
	progressFormat.Set("inplace")
	_, err := initializeProgressLogger(context.Background())
	assert.NoError(t, err)
}

func TestNoErrorOnNonStderrLogFile(t *testing.T) {
	logLevel.Set("info")
	logFile.Set("stdout")
	progressFormat.Set("inplace")
	_, err := initializeProgressLogger(context.Background())
	assert.NoError(t, err)
}

func TestDefaultLoggerModeResolution(t *testing.T) {
	progressFormat = flags.NewProgressLogFormat()
	require.Equal(t, progressFormat, flags.ModeDefault)
	ctx, err := initializeProgressLogger(context.Background())
	require.NoError(t, err)
	logger, ok := progress.FromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, logger.Mode, flags.ModeAppend)
}
