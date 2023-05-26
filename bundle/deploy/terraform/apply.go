package terraform

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/databricks/cli/bundle"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/hashicorp/terraform-exec/tfexec"
)

type apply struct{}

func (w *apply) Name() string {
	return "terraform.Apply"
}

func (w *apply) Apply(ctx context.Context, b *bundle.Bundle) error {
	tf := b.Terraform
	if tf == nil {
		return fmt.Errorf("terraform not initialized")
	}

	cmdio.LogString(ctx, "Starting resource deployment")

	err := tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		return fmt.Errorf("terraform init: %w", err)
	}

	tfVersion, providerVersion, err := tf.Version(ctx, true)
	if err != nil {
		return fmt.Errorf("terraform version: %w", err)
	}

	cmdio.LogString(ctx, fmt.Sprintf("Terraform version: %s", tfVersion))
	for provider, version := range providerVersion {
		cmdio.LogString(ctx, fmt.Sprintf("Terraform provider %s version: %s", provider, version.String()))
	}

	d, _ := Dir(b)
	tf.SetLogProvider("DEBUG")
	tf.SetLogPath(filepath.Join(d, "terraform.log"))

	err = tf.Apply(ctx)
	if err != nil {
		return fmt.Errorf("terraform apply: %w", err)
	}

	cmdio.LogString(ctx, "Resource deployment completed!")
	return nil
}

// Apply returns a [bundle.Mutator] that runs the equivalent of `terraform apply`
// from the bundle's ephemeral working directory for Terraform.
func Apply() bundle.Mutator {
	return &apply{}
}
