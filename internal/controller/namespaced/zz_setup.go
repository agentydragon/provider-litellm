// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accessgroup "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/accessgroup"
	agent "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/agent"
	budget "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/budget"
	credential "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/credential"
	fallback "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/fallback"
	guardrail "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/guardrail"
	key "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/key"
	keyblock "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/keyblock"
	mcpserver "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/mcpserver"
	model "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/model"
	organization "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/organization"
	organizationmember "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/organizationmember"
	project "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/project"
	prompt "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/prompt"
	searchtool "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/searchtool"
	tag "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/tag"
	team "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/team"
	teamblock "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/teamblock"
	teammember "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/teammember"
	teammemberadd "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/teammemberadd"
	unifiedaccessgroup "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/unifiedaccessgroup"
	user "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/user"
	vectorstore "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/litellm/vectorstore"
	providerconfig "github.com/thecodingsheikh/provider-litellm/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessgroup.Setup,
		agent.Setup,
		budget.Setup,
		credential.Setup,
		fallback.Setup,
		guardrail.Setup,
		key.Setup,
		keyblock.Setup,
		mcpserver.Setup,
		model.Setup,
		organization.Setup,
		organizationmember.Setup,
		project.Setup,
		prompt.Setup,
		searchtool.Setup,
		tag.Setup,
		team.Setup,
		teamblock.Setup,
		teammember.Setup,
		teammemberadd.Setup,
		unifiedaccessgroup.Setup,
		user.Setup,
		vectorstore.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessgroup.SetupGated,
		agent.SetupGated,
		budget.SetupGated,
		credential.SetupGated,
		fallback.SetupGated,
		guardrail.SetupGated,
		key.SetupGated,
		keyblock.SetupGated,
		mcpserver.SetupGated,
		model.SetupGated,
		organization.SetupGated,
		organizationmember.SetupGated,
		project.SetupGated,
		prompt.SetupGated,
		searchtool.SetupGated,
		tag.SetupGated,
		team.SetupGated,
		teamblock.SetupGated,
		teammember.SetupGated,
		teammemberadd.SetupGated,
		unifiedaccessgroup.SetupGated,
		user.SetupGated,
		vectorstore.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		accessgroup.SetupWebhookWithManager,
		agent.SetupWebhookWithManager,
		budget.SetupWebhookWithManager,
		credential.SetupWebhookWithManager,
		fallback.SetupWebhookWithManager,
		guardrail.SetupWebhookWithManager,
		key.SetupWebhookWithManager,
		keyblock.SetupWebhookWithManager,
		mcpserver.SetupWebhookWithManager,
		model.SetupWebhookWithManager,
		organization.SetupWebhookWithManager,
		organizationmember.SetupWebhookWithManager,
		project.SetupWebhookWithManager,
		prompt.SetupWebhookWithManager,
		searchtool.SetupWebhookWithManager,
		tag.SetupWebhookWithManager,
		team.SetupWebhookWithManager,
		teamblock.SetupWebhookWithManager,
		teammember.SetupWebhookWithManager,
		teammemberadd.SetupWebhookWithManager,
		unifiedaccessgroup.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		vectorstore.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
