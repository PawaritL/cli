package phases

import (
	"github.com/databricks/bricks/bundle"
	"github.com/databricks/bricks/bundle/artifacts"
	"github.com/databricks/bricks/bundle/deploy/files"
	"github.com/databricks/bricks/bundle/deploy/lock"
	"github.com/databricks/bricks/bundle/deploy/terraform"
)

// The deploy phase deploys artifacts and resources.
func Deploy() bundle.Mutator {
	return newPhase(
		"deploy",
		[]bundle.Mutator{
			lock.Acquire(),
			files.Upload(),
			artifacts.UploadAll(),
			terraform.Interpolate(),
			terraform.Write(),
			terraform.StatePull(),
			terraform.Apply(),
			terraform.StatePush(),
			lock.Release(),
		},
	)
}
