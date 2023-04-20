package config

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRootMarshalUnmarshal(t *testing.T) {
	// Marshal empty
	buf, err := json.Marshal(&Root{})
	require.NoError(t, err)

	// Unmarshal empty
	var root Root
	err = json.Unmarshal(buf, &root)
	require.NoError(t, err)

	// Compare
	assert.True(t, reflect.DeepEqual(Root{}, root))
}

func TestRootLoad(t *testing.T) {
	root := &Root{}
	err := root.Load("../tests/basic/bundle.yml")
	require.NoError(t, err)
	assert.Equal(t, "basic", root.Bundle.Name)
}

func TestRootMergeStruct(t *testing.T) {
	root := &Root{
		Path: "path",
		Workspace: Workspace{
			Host:    "foo",
			Profile: "profile",
		},
	}
	other := &Root{
		Path: "path",
		Workspace: Workspace{
			Host: "bar",
		},
	}
	assert.NoError(t, root.Merge(other))
	assert.Equal(t, "bar", root.Workspace.Host)
	assert.Equal(t, "profile", root.Workspace.Profile)
}

func TestRootMergeMap(t *testing.T) {
	root := &Root{
		Path: "path",
		Environments: map[string]*Environment{
			"development": {
				Workspace: &Workspace{
					Host:    "foo",
					Profile: "profile",
				},
			},
		},
	}
	other := &Root{
		Path: "path",
		Environments: map[string]*Environment{
			"development": {
				Workspace: &Workspace{
					Host: "bar",
				},
			},
		},
	}
	assert.NoError(t, root.Merge(other))
	assert.Equal(t, &Workspace{Host: "bar", Profile: "profile"}, root.Environments["development"].Workspace)
}

func TestDuplicateIdOnLoadReturnsError(t *testing.T) {
	root := &Root{}
	err := root.Load("./testdata/duplicate_resource_names_in_root/bundle.yml")
	assert.ErrorContains(t, err, "multiple resources named foo (job at ./testdata/duplicate_resource_names_in_root/bundle.yml, pipeline at ./testdata/duplicate_resource_names_in_root/bundle.yml)")
}

func TestDuplicateIdOnMergeReturnsError(t *testing.T) {
	root := &Root{}
	err := root.Load("./testdata/duplicate_resource_name_in_subconfiguration/bundle.yml")
	require.NoError(t, err)

	other := &Root{}
	err = other.Load("./testdata/duplicate_resource_name_in_subconfiguration/resources.yml")
	require.NoError(t, err)

	err = root.Merge(other)
	assert.ErrorContains(t, err, "multiple resources named foo (job at ./testdata/duplicate_resource_name_in_subconfiguration/bundle.yml, pipeline at ./testdata/duplicate_resource_name_in_subconfiguration/resources.yml)")
}
