// Generated from Databricks Terraform provider schema. DO NOT EDIT.

package schema

type ResourceClusterPolicy struct {
	Definition         string `json:"definition"`
	Id                 string `json:"id,omitempty"`
	MaxClustersPerUser int    `json:"max_clusters_per_user,omitempty"`
	Name               string `json:"name"`
	PolicyId           string `json:"policy_id,omitempty"`
}
