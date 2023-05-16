package terraform

import (
	"context"
	"os/exec"
	"testing"

	"github.com/databricks/cli/bundle"
	"github.com/databricks/cli/bundle/config"
	"github.com/stretchr/testify/require"
)

func TestLoadWithNoState(t *testing.T) {
	_, err := exec.LookPath("terraform")
	if err != nil {
		t.Skipf("cannot find terraform binary: %s", err)
	}

	b := &bundle.Bundle{
		Config: config.Root{
			Path: t.TempDir(),
			Bundle: config.Bundle{
				Environment: "whatever",
				Terraform: &config.Terraform{
					ExecPath: "terraform",
				},
			},
		},
	}

	t.Setenv("DATABRICKS_HOST", "https://x")
	t.Setenv("DATABRICKS_TOKEN", "foobar")
	b.WorkspaceClient()

	err = bundle.Apply(context.Background(), b, []bundle.Mutator{
		Initialize(),
		Load(),
	})

	require.ErrorContains(t, err, "Did you forget to run 'databricks bundle deploy'")
}
