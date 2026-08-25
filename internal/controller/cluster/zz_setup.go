// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	credential "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/credential"
	key "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/key"
	mcpserver "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/mcpserver"
	model "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/model"
	organization "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/organization"
	team "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/team"
	teammember "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/teammember"
	user "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/litellm/user"
	providerconfig "github.com/thecodingsheikh/provider-litellm/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		credential.Setup,
		key.Setup,
		mcpserver.Setup,
		model.Setup,
		organization.Setup,
		team.Setup,
		teammember.Setup,
		user.Setup,
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
		credential.SetupGated,
		key.SetupGated,
		mcpserver.SetupGated,
		model.SetupGated,
		organization.SetupGated,
		team.SetupGated,
		teammember.SetupGated,
		user.SetupGated,
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
		credential.SetupWebhookWithManager,
		key.SetupWebhookWithManager,
		mcpserver.SetupWebhookWithManager,
		model.SetupWebhookWithManager,
		organization.SetupWebhookWithManager,
		team.SetupWebhookWithManager,
		teammember.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
