package debug

import (
	"github.com/spf13/cobra"

	parent "github.com/databricks/bricks/cmd/bundle"
)

var debugCmd = &cobra.Command{
	Use: "debug",
}

func init() {
	parent.AddCommand(debugCmd)
}
