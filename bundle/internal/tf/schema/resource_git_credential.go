// Generated from Databricks Terraform provider schema. DO NOT EDIT.

package schema

type ResourceGitCredential struct {
	Force               bool   `json:"force,omitempty"`
	GitProvider         string `json:"git_provider"`
	GitUsername         string `json:"git_username,omitempty"`
	Id                  string `json:"id,omitempty"`
	PersonalAccessToken string `json:"personal_access_token,omitempty"`
}
