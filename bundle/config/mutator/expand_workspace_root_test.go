package mutator_test

import (
	"context"
	"testing"

	"github.com/databricks/cli/bundle"
	"github.com/databricks/cli/bundle/config"
	"github.com/databricks/cli/bundle/config/mutator"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpandWorkspaceRoot(t *testing.T) {
	bundle := &bundle.Bundle{
		Config: config.Root{
			Workspace: config.Workspace{
				CurrentUser: &iam.User{
					UserName: "jane@doe.com",
				},
				RootPath: "~/foo",
			},
		},
	}
	_, err := mutator.ExpandWorkspaceRoot().Apply(context.Background(), bundle)
	require.NoError(t, err)
	assert.Equal(t, "/Users/jane@doe.com/foo", bundle.Config.Workspace.RootPath)
}

func TestExpandWorkspaceRootDoesNothing(t *testing.T) {
	bundle := &bundle.Bundle{
		Config: config.Root{
			Workspace: config.Workspace{
				CurrentUser: &iam.User{
					UserName: "jane@doe.com",
				},
				RootPath: "/Users/charly@doe.com/foo",
			},
		},
	}
	_, err := mutator.ExpandWorkspaceRoot().Apply(context.Background(), bundle)
	require.NoError(t, err)
	assert.Equal(t, "/Users/charly@doe.com/foo", bundle.Config.Workspace.RootPath)
}

func TestExpandWorkspaceRootWithoutRoot(t *testing.T) {
	bundle := &bundle.Bundle{
		Config: config.Root{
			Workspace: config.Workspace{
				CurrentUser: &iam.User{
					UserName: "jane@doe.com",
				},
			},
		},
	}
	_, err := mutator.ExpandWorkspaceRoot().Apply(context.Background(), bundle)
	require.Error(t, err)
}

func TestExpandWorkspaceRootWithoutCurrentUser(t *testing.T) {
	bundle := &bundle.Bundle{
		Config: config.Root{
			Workspace: config.Workspace{
				RootPath: "~/foo",
			},
		},
	}
	_, err := mutator.ExpandWorkspaceRoot().Apply(context.Background(), bundle)
	require.Error(t, err)
}
