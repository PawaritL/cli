package artifacts

import (
	"context"
	"fmt"

	"github.com/databricks/cli/bundle"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// all is an internal proxy for producing a list of mutators for all artifacts.
// It is used to produce the [BuildAll] and [UploadAll] mutators.
type all struct {
	name string
	fn   func(name string) (bundle.Mutator, error)
}

func (m *all) Name() string {
	return fmt.Sprintf("artifacts.%sAll", m.name)
}

func (m *all) Apply(ctx context.Context, b *bundle.Bundle) ([]bundle.Mutator, error) {
	var out []bundle.Mutator

	// Iterate with stable ordering.
	keys := maps.Keys(b.Config.Artifacts)
	slices.Sort(keys)

	for _, name := range keys {
		m, err := m.fn(name)
		if err != nil {
			return nil, err
		}
		if m != nil {
			out = append(out, m)
		}
	}

	return out, nil
}
