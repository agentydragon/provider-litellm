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

	// Composite identifier the provider derives from two required arguments:
	// `terraform import litellm_team_member.example <team_id>:<user_id>`.
	// It must be templated rather than read back from state: the upstream
	// resource never writes id into the Terraform state, so
	// IdentifierFromProvider leaves the external name empty forever.
	"litellm_team_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.team_id }}:{{ .parameters.user_id }}"),

	// Both of these are keyed by the team they act on, not by an identifier of
	// their own: `terraform import litellm_team_block.example <team-id>`.
	// ParameterAsIdentifier("team_id") would be the literal reading, but it
	// omits team_id from the spec, which would force users to hand-write the
	// team's generated UUID into the external-name annotation and give up
	// teamIdRef. The ID equals team_id either way, so taking it from state
	// keeps the reference usable.
	"litellm_team_block": config.IdentifierFromProvider,
	// team_member_add shares litellm_team_member's defect of never writing id
	// into state, so its identifier has to be templated from the team instead.
	"litellm_team_member_add": config.TemplatedStringAsIdentifier("", "{{ .parameters.team_id }}"),

	// The ID is a SHA256 hash of the key token LiteLLM issues on creation, so
	// it cannot be supplied up front.
	"litellm_key": config.IdentifierFromProvider,

	// Keyed by the key token it blocks, which is a sensitive argument and so is
	// surfaced as a SecretRef rather than as the external name.
	"litellm_key_block": config.IdentifierFromProvider,

	// organization_id is a ForceNew argument that LiteLLM only generates when
	// it is omitted, and the resource ID is the same value.
	"litellm_organization": config.ParameterAsIdentifier("organization_id"),

	// Composite identifier, same shape as litellm_team_member:
	// `terraform import litellm_organization_member.example <organization_id>:<user_id>`.
	"litellm_organization_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.organization_id }}:{{ .parameters.user_id }}"),

	// user_id behaves exactly as organization_id above.
	"litellm_user": config.ParameterAsIdentifier("user_id"),

	// server_id is computed only; LiteLLM always generates it.
	"litellm_mcp_server": config.IdentifierFromProvider,

	// Named resources whose import identifier is the name argument itself.
	"litellm_access_group": config.ParameterAsIdentifier("access_group"),
	"litellm_prompt":       config.ParameterAsIdentifier("prompt_id"),
	"litellm_tag":          config.NameAsIdentifier,

	// budget_id is honoured when supplied, so it can carry the external name.
	"litellm_budget": config.ParameterAsIdentifier("budget_id"),

	// guardrail_id looks the same but is not: LiteLLM always allocates its own
	// UUID and ignores the supplied value, so feeding the external name into it
	// makes the next plan a ForceNew that trips upjet's prevent_destroy guard.
	"litellm_guardrail": config.IdentifierFromProvider,

	// Composite of the model and the fallback type:
	// `terraform import litellm_fallback.example "gpt-3.5-turbo:general"`.
	"litellm_fallback": config.TemplatedStringAsIdentifier("", "{{ .parameters.model }}:{{ .parameters.fallback_type }}"),

	// Provider-generated identifiers with no settable argument.
	"litellm_agent":                config.IdentifierFromProvider,
	"litellm_project":              config.IdentifierFromProvider,
	"litellm_search_tool":          config.IdentifierFromProvider,
	"litellm_unified_access_group": config.IdentifierFromProvider,
	"litellm_vector_store":         config.IdentifierFromProvider,
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
