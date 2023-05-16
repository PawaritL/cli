package metastores

import "github.com/databricks/cli/libs/cmdio"

func init() {
	listCmd.Annotations["template"] = cmdio.Heredoc(`
	{{white "ID"}}	{{white "Name"}}	{{white "Region"}}
	{{range .}}{{.MetastoreId|green}}	{{.Name|white}}	{{.Region}}
	{{end}}`)
}
