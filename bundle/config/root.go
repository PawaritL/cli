package config

import (
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/imdario/mergo"
)

// FileName is the name of bundle configuration file.
const FileName = "bundle.yml"

type Root struct {
	// Path contains the directory path to the root of the bundle.
	// It is set when loading `bundle.yml`.
	Path string `json:"-" bundle:"readonly"`

	// Bundle contains details about this bundle, such as its name,
	// version of the spec (TODO), default cluster, default warehouse, etc.
	Bundle Bundle `json:"bundle"`

	// Include specifies a list of patterns of file names to load and
	// merge into the this configuration. If not set in `bundle.yml`,
	// it defaults to loading `*.yml` and `*/*.yml`.
	//
	// Also see [mutator.DefineDefaultInclude].
	//
	Include []string `json:"include,omitempty"`

	// Workspace contains details about the workspace to connect to
	// and paths in the workspace tree to use for this bundle.
	Workspace Workspace `json:"workspace"`

	// Artifacts contains a description of all code artifacts in this bundle.
	Artifacts map[string]*Artifact `json:"artifacts,omitempty"`

	// Resources contains a description of all Databricks resources
	// to deploy in this bundle (e.g. jobs, pipelines, etc.).
	Resources Resources `json:"resources,omitempty"`

	// Environments can be used to differentiate settings and resources between
	// bundle deployment environments (e.g. development, staging, production).
	// If not specified, the code below initializes this field with a
	// single default-initialized environment called "default".
	Environments map[string]*Environment `json:"environments,omitempty"`
}

func Load(path string) (*Root, error) {
	var r Root

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// If we were given a directory, assume this is the bundle root.
	if stat.IsDir() {
		path = filepath.Join(path, FileName)
	}

	if err := r.Load(path); err != nil {
		return nil, err
	}

	return &r, nil
}

// SetConfigFilePath configures the path that its configuration
// was loaded from in configuration leafs that require it.
func (r *Root) SetConfigFilePath(path string) {
	r.Resources.SetConfigFilePath(path)
	if r.Environments != nil {
		for _, env := range r.Environments {
			if env.Resources != nil {
				env.Resources.SetConfigFilePath(path)
			}
		}
	}
}

func (r *Root) Load(path string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, r)
	if err != nil {
		return err
	}

	r.Path = filepath.Dir(path)
	r.SetConfigFilePath(path)

	_, err = r.Resources.VerifyUniqueResourceIdentifiers()
	return err
}

func (r *Root) Merge(other *Root) error {
	// TODO: when hooking into merge semantics, disallow setting path on the target instance.
	other.Path = ""

	// Check for safe merge, protecting against duplicate resource identifiers
	err := r.Resources.VerifySafeMerge(&other.Resources)
	if err != nil {
		return err
	}

	// TODO: define and test semantics for merging.
	return mergo.MergeWithOverwrite(r, other)
}

func (r *Root) MergeEnvironment(env *Environment) error {
	var err error

	// Environment may be nil if it's empty.
	if env == nil {
		return nil
	}

	if env.Bundle != nil {
		err = mergo.MergeWithOverwrite(&r.Bundle, env.Bundle)
		if err != nil {
			return err
		}
	}

	if env.Workspace != nil {
		err = mergo.MergeWithOverwrite(&r.Workspace, env.Workspace)
		if err != nil {
			return err
		}
	}

	if env.Artifacts != nil {
		err = mergo.Merge(&r.Artifacts, env.Artifacts, mergo.WithAppendSlice)
		if err != nil {
			return err
		}
	}

	if env.Resources != nil {
		err = mergo.Merge(&r.Resources, env.Resources, mergo.WithAppendSlice)
		if err != nil {
			return err
		}
	}

	return nil
}
