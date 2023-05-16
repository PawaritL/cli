package tables

import "github.com/databricks/cli/libs/cmdio"

func init() {
	listCmd.Annotations["template"] = cmdio.Heredoc(`
	{{white "Full Name"}}	{{white "Table Type"}}
	{{range .}}{{.FullName|green}}	{{blue "%s" .TableType}}
	{{end}}`)
}
