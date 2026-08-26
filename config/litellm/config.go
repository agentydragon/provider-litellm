// Package litellm contains the Crossplane resource configuration for the
// resources exposed by the ncecere/litellm Terraform provider.
package litellm

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Terraform names of the resources referenced from more than one place.
const (
	resourceTeam         = "litellm_team"
	resourceUser         = "litellm_user"
	resourceOrganization = "litellm_organization"
	resourceBudget       = "litellm_budget"
)

// Configure configures the LiteLLM resources.
//
// Every managed resource is served from the provider's root API group, so
// ShortGroup is left empty throughout. Kinds are set explicitly because upjet
// derives them from the shape of the Terraform name, which is inconsistent
// here: litellm_team_member would default to group "team", kind "Member" while
// litellm_model would default to group "litellm", kind "Model".
//
// References point at the external name of the referenced resource, which is
// why they line up with the external name configuration in
// config/external_name.go. Fields holding model *names* rather than IDs (such
// as the models lists on Team, Key and User, or Fallback's model) are
// deliberately left unreferenced: LiteLLM matches those against model_name,
// not against the Model's external name. Identifier arguments are likewise
// unreferenced, because ParameterAsIdentifier omits them from the spec.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("litellm_model", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Model"
		r.References["litellm_credential_name"] = ujconfig.Reference{
			TerraformName: "litellm_credential",
		}
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
	})

	p.AddResourceConfigurator("litellm_credential", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Credential"
		r.References["model_id"] = ujconfig.Reference{
			TerraformName: "litellm_model",
		}
	})

	p.AddResourceConfigurator("litellm_team", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Team"
		r.References["organization_id"] = ujconfig.Reference{
			TerraformName: resourceOrganization,
		}
	})

	p.AddResourceConfigurator("litellm_team_member", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "TeamMember"
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
		r.References["user_id"] = ujconfig.Reference{
			TerraformName: resourceUser,
		}
	})

	p.AddResourceConfigurator("litellm_team_member_add", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "TeamMemberAdd"
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
	})

	p.AddResourceConfigurator("litellm_team_block", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "TeamBlock"
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
	})

	p.AddResourceConfigurator("litellm_key", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Key"
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
		r.References["organization_id"] = ujconfig.Reference{
			TerraformName: resourceOrganization,
		}
		r.References["user_id"] = ujconfig.Reference{
			TerraformName: resourceUser,
		}
		r.References["budget_id"] = ujconfig.Reference{
			TerraformName: resourceBudget,
		}
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "litellm_project",
		}
	})

	p.AddResourceConfigurator("litellm_key_block", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "KeyBlock"
	})

	p.AddResourceConfigurator("litellm_organization", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Organization"
		r.References["budget_id"] = ujconfig.Reference{
			TerraformName: resourceBudget,
		}
	})

	p.AddResourceConfigurator("litellm_organization_member", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "OrganizationMember"
		r.References["organization_id"] = ujconfig.Reference{
			TerraformName: resourceOrganization,
		}
		r.References["user_id"] = ujconfig.Reference{
			TerraformName: resourceUser,
		}
	})

	p.AddResourceConfigurator("litellm_user", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "User"
		// teams is a list of team IDs, which is what a Team's external name holds.
		r.References["teams"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
	})

	p.AddResourceConfigurator("litellm_project", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Project"
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
	})

	p.AddResourceConfigurator("litellm_mcp_server", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "MCPServer"
	})

	p.AddResourceConfigurator("litellm_access_group", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "AccessGroup"
	})

	p.AddResourceConfigurator("litellm_unified_access_group", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "UnifiedAccessGroup"
		r.References["assigned_team_ids"] = ujconfig.Reference{
			TerraformName: resourceTeam,
		}
		r.References["access_mcp_server_ids"] = ujconfig.Reference{
			TerraformName: "litellm_mcp_server",
		}
		r.References["access_agent_ids"] = ujconfig.Reference{
			TerraformName: "litellm_agent",
		}
	})

	p.AddResourceConfigurator("litellm_agent", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Agent"
		// agent_card is the only nested block on this provider that itself
		// contains further blocks, and upjet classifies it as observation only,
		// so it lands in status.atProvider and cannot be set. The upstream
		// provider meanwhile rejects any plan that does not configure
		// agent_card.name and agent_card.url, which leaves the resource
		// impossible to create. Put the block back into the spec.
		if s := r.TerraformResource.Schema["agent_card"]; s != nil {
			s.Computed = false
			s.Optional = true
		}
		// Both blocks use Terraform's single nesting mode, so the provider reads
		// them as objects. upjet's default singleton-list representation makes
		// it reject the state with `AttributeName("agent_card"): invalid JSON,
		// expected "{", got "["`, so convert them to embedded objects.
		r.AddSingletonListConversion("agent_card", "agentCard")
		r.AddSingletonListConversion("object_permission", "objectPermission")
	})

	p.AddResourceConfigurator("litellm_budget", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Budget"
	})

	p.AddResourceConfigurator("litellm_fallback", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Fallback"
	})

	p.AddResourceConfigurator("litellm_guardrail", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Guardrail"
	})

	p.AddResourceConfigurator("litellm_prompt", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Prompt"
	})

	p.AddResourceConfigurator("litellm_search_tool", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "SearchTool"
	})

	p.AddResourceConfigurator("litellm_tag", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Tag"
	})

	p.AddResourceConfigurator("litellm_jwt_key_mapping", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "JWTKeyMapping"
	})

	p.AddResourceConfigurator("litellm_vector_store", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "VectorStore"
	})
}
