// Generated from Databricks Terraform provider schema. DO NOT EDIT.

package schema

type ResourceExternalLocation struct {
	Comment        string `json:"comment,omitempty"`
	CredentialName string `json:"credential_name"`
	ForceDestroy   bool   `json:"force_destroy,omitempty"`
	Id             string `json:"id,omitempty"`
	MetastoreId    string `json:"metastore_id,omitempty"`
	Name           string `json:"name"`
	Owner          string `json:"owner,omitempty"`
	ReadOnly       bool   `json:"read_only,omitempty"`
	SkipValidation bool   `json:"skip_validation,omitempty"`
	Url            string `json:"url"`
}
