// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cmd

import (
	"github.com/databricks/cli/cmd/root"

	alerts "github.com/databricks/cli/cmd/workspace/alerts"
	catalogs "github.com/databricks/cli/cmd/workspace/catalogs"
	cluster_policies "github.com/databricks/cli/cmd/workspace/cluster-policies"
	clusters "github.com/databricks/cli/cmd/workspace/clusters"
	current_user "github.com/databricks/cli/cmd/workspace/current-user"
	dashboards "github.com/databricks/cli/cmd/workspace/dashboards"
	data_sources "github.com/databricks/cli/cmd/workspace/data-sources"
	experiments "github.com/databricks/cli/cmd/workspace/experiments"
	external_locations "github.com/databricks/cli/cmd/workspace/external-locations"
	functions "github.com/databricks/cli/cmd/workspace/functions"
	git_credentials "github.com/databricks/cli/cmd/workspace/git-credentials"
	global_init_scripts "github.com/databricks/cli/cmd/workspace/global-init-scripts"
	grants "github.com/databricks/cli/cmd/workspace/grants"
	groups "github.com/databricks/cli/cmd/workspace/groups"
	instance_pools "github.com/databricks/cli/cmd/workspace/instance-pools"
	instance_profiles "github.com/databricks/cli/cmd/workspace/instance-profiles"
	ip_access_lists "github.com/databricks/cli/cmd/workspace/ip-access-lists"
	jobs "github.com/databricks/cli/cmd/workspace/jobs"
	libraries "github.com/databricks/cli/cmd/workspace/libraries"
	metastores "github.com/databricks/cli/cmd/workspace/metastores"
	model_registry "github.com/databricks/cli/cmd/workspace/model-registry"
	permissions "github.com/databricks/cli/cmd/workspace/permissions"
	pipelines "github.com/databricks/cli/cmd/workspace/pipelines"
	policy_families "github.com/databricks/cli/cmd/workspace/policy-families"
	providers "github.com/databricks/cli/cmd/workspace/providers"
	queries "github.com/databricks/cli/cmd/workspace/queries"
	query_history "github.com/databricks/cli/cmd/workspace/query-history"
	recipient_activation "github.com/databricks/cli/cmd/workspace/recipient-activation"
	recipients "github.com/databricks/cli/cmd/workspace/recipients"
	repos "github.com/databricks/cli/cmd/workspace/repos"
	schemas "github.com/databricks/cli/cmd/workspace/schemas"
	secrets "github.com/databricks/cli/cmd/workspace/secrets"
	service_principals "github.com/databricks/cli/cmd/workspace/service-principals"
	serving_endpoints "github.com/databricks/cli/cmd/workspace/serving-endpoints"
	shares "github.com/databricks/cli/cmd/workspace/shares"
	storage_credentials "github.com/databricks/cli/cmd/workspace/storage-credentials"
	table_constraints "github.com/databricks/cli/cmd/workspace/table-constraints"
	tables "github.com/databricks/cli/cmd/workspace/tables"
	token_management "github.com/databricks/cli/cmd/workspace/token-management"
	tokens "github.com/databricks/cli/cmd/workspace/tokens"
	users "github.com/databricks/cli/cmd/workspace/users"
	volumes "github.com/databricks/cli/cmd/workspace/volumes"
	warehouses "github.com/databricks/cli/cmd/workspace/warehouses"
	workspace "github.com/databricks/cli/cmd/workspace/workspace"
	workspace_conf "github.com/databricks/cli/cmd/workspace/workspace-conf"
)

func init() {

	root.RootCmd.AddCommand(alerts.Cmd)
	root.RootCmd.AddCommand(catalogs.Cmd)
	root.RootCmd.AddCommand(cluster_policies.Cmd)
	root.RootCmd.AddCommand(clusters.Cmd)
	root.RootCmd.AddCommand(current_user.Cmd)
	root.RootCmd.AddCommand(dashboards.Cmd)
	root.RootCmd.AddCommand(data_sources.Cmd)
	root.RootCmd.AddCommand(experiments.Cmd)
	root.RootCmd.AddCommand(external_locations.Cmd)
	root.RootCmd.AddCommand(functions.Cmd)
	root.RootCmd.AddCommand(git_credentials.Cmd)
	root.RootCmd.AddCommand(global_init_scripts.Cmd)
	root.RootCmd.AddCommand(grants.Cmd)
	root.RootCmd.AddCommand(groups.Cmd)
	root.RootCmd.AddCommand(instance_pools.Cmd)
	root.RootCmd.AddCommand(instance_profiles.Cmd)
	root.RootCmd.AddCommand(ip_access_lists.Cmd)
	root.RootCmd.AddCommand(jobs.Cmd)
	root.RootCmd.AddCommand(libraries.Cmd)
	root.RootCmd.AddCommand(metastores.Cmd)
	root.RootCmd.AddCommand(model_registry.Cmd)
	root.RootCmd.AddCommand(permissions.Cmd)
	root.RootCmd.AddCommand(pipelines.Cmd)
	root.RootCmd.AddCommand(policy_families.Cmd)
	root.RootCmd.AddCommand(providers.Cmd)
	root.RootCmd.AddCommand(queries.Cmd)
	root.RootCmd.AddCommand(query_history.Cmd)
	root.RootCmd.AddCommand(recipient_activation.Cmd)
	root.RootCmd.AddCommand(recipients.Cmd)
	root.RootCmd.AddCommand(repos.Cmd)
	root.RootCmd.AddCommand(schemas.Cmd)
	root.RootCmd.AddCommand(secrets.Cmd)
	root.RootCmd.AddCommand(service_principals.Cmd)
	root.RootCmd.AddCommand(serving_endpoints.Cmd)
	root.RootCmd.AddCommand(shares.Cmd)
	root.RootCmd.AddCommand(storage_credentials.Cmd)
	root.RootCmd.AddCommand(table_constraints.Cmd)
	root.RootCmd.AddCommand(tables.Cmd)
	root.RootCmd.AddCommand(token_management.Cmd)
	root.RootCmd.AddCommand(tokens.Cmd)
	root.RootCmd.AddCommand(users.Cmd)
	root.RootCmd.AddCommand(volumes.Cmd)
	root.RootCmd.AddCommand(warehouses.Cmd)
	root.RootCmd.AddCommand(workspace.Cmd)
	root.RootCmd.AddCommand(workspace_conf.Cmd)
}
