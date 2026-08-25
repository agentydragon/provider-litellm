package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// This map doubles as the provider's include list — only the resources listed
// here are code generated. See ExternalNameConfigured.
//
// Each entry mirrors the "## Import" section of the corresponding page under
// docs/resources in the ncecere/terraform-provider-litellm repository, which is
// what determines the shape of the Terraform ID.
var ExternalNameConfigs = map[string]config.ExternalName{
	// LiteLLM generates the model ID on creation and it is distinct from
	// model_name: `terraform import litellm_model.example <model-id>`.
	"litellm_model": config.IdentifierFromProvider,

	// The ID is the credential name, a ForceNew argument:
	// `terraform import litellm_credential.example "<credential-name>"`.
	"litellm_credential": config.ParameterAsIdentifier("credential_name"),

	// LiteLLM generates the team ID; the resource exposes no settable
	// team_id argument: `terraform import litellm_team.example <team-id>`.
	"litellm_team": config.IdentifierFromProvider,

	// Composite identifier built by the provider from two required arguments:
	// `terraform import litellm_team_member.example <team_id>:<user_id>`.
	"litellm_team_member": config.IdentifierFromProvider,

	// The ID is a SHA256 hash of the key token LiteLLM issues on creation, so
	// it cannot be supplied up front.
	"litellm_key": config.IdentifierFromProvider,

	// organization_id is a ForceNew argument that LiteLLM only generates when
	// it is omitted, and the resource ID is the same value.
	"litellm_organization": config.ParameterAsIdentifier("organization_id"),

	// user_id behaves exactly as organization_id above.
	"litellm_user": config.ParameterAsIdentifier("user_id"),

	// server_id is computed only; LiteLLM always generates it.
	"litellm_mcp_server": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
