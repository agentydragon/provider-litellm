// Package litellm contains the Crossplane resource configuration for the
// resources exposed by the ncecere/litellm Terraform provider.
package litellm

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
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
// as the models lists on Team, Key and User) are deliberately left
// unreferenced: LiteLLM matches those against model_name, not the Model's
// external name.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("litellm_model", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Model"
		r.References["litellm_credential_name"] = ujconfig.Reference{
			TerraformName: "litellm_credential",
		}
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: "litellm_team",
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
			TerraformName: "litellm_organization",
		}
	})

	p.AddResourceConfigurator("litellm_team_member", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "TeamMember"
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: "litellm_team",
		}
		r.References["user_id"] = ujconfig.Reference{
			TerraformName: "litellm_user",
		}
	})

	p.AddResourceConfigurator("litellm_key", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Key"
		r.References["team_id"] = ujconfig.Reference{
			TerraformName: "litellm_team",
		}
		r.References["organization_id"] = ujconfig.Reference{
			TerraformName: "litellm_organization",
		}
		r.References["user_id"] = ujconfig.Reference{
			TerraformName: "litellm_user",
		}
	})

	p.AddResourceConfigurator("litellm_organization", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "Organization"
	})

	p.AddResourceConfigurator("litellm_user", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "User"
		// teams is a list of team IDs, which is what a Team's external name holds.
		r.References["teams"] = ujconfig.Reference{
			TerraformName: "litellm_team",
		}
	})

	p.AddResourceConfigurator("litellm_mcp_server", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "MCPServer"
	})
}
