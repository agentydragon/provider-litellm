# provider-litellm

**Crossplane provider for [LiteLLM](https://docs.litellm.ai/).** Manage your LiteLLM
proxy — models, credentials, virtual keys, teams, organizations, budgets, guardrails,
MCP servers and more — as declarative Kubernetes resources.

[![Release](https://img.shields.io/github/v/tag/TheCodingSheikh/provider-litellm?label=release&sort=semver)](https://github.com/TheCodingSheikh/provider-litellm/releases)
[![Upbound Marketplace](https://img.shields.io/badge/marketplace-provider--litellm-blue)](https://marketplace.upbound.io/providers/thecodingsheikh/provider-litellm)
[![CI](https://github.com/TheCodingSheikh/provider-litellm/actions/workflows/ci.yml/badge.svg)](https://github.com/TheCodingSheikh/provider-litellm/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue)](LICENSE)
[![Crossplane](https://img.shields.io/badge/crossplane-v1%20%7C%20v2-informational)](https://crossplane.io/)

`provider-litellm` turns the [LiteLLM](https://docs.litellm.ai/) proxy management API
into Kubernetes CRDs. Instead of clicking through the LiteLLM Admin UI or POSTing to
`/model/new` from a script, you write YAML, `kubectl apply` it, and Crossplane keeps
LiteLLM converged on that state — the same GitOps loop you already run for the rest of
your platform. It is built with [Upjet](https://github.com/crossplane/upjet) on top of
the [`ncecere/litellm`](https://registry.terraform.io/providers/ncecere/litellm/latest)
Terraform provider.

Both API groups are served, so it works on either Crossplane generation:

| Crossplane | API group | Scope |
| --- | --- | --- |
| v1 | `litellm.crossplane.io` | Cluster-scoped managed resources |
| v2 | `litellm.m.crossplane.io` | Namespaced managed resources |

## Why manage LiteLLM with Crossplane

- **GitOps for the AI gateway.** Argo CD or Flux reconciles your LiteLLM models, teams
  and budgets from Git, with the same review and rollback story as everything else in
  the cluster. No drift between what is in the repo and what is on the proxy.
- **Self-service through Compositions.** Bundle a `Team`, a `Budget` and a `Key` into a
  single XR so application teams can request LLM access without touching the LiteLLM
  admin console or holding the master key.
- **Kubernetes-native secrets.** Upstream provider API keys are read from Kubernetes
  Secrets rather than inlined in the spec, and values LiteLLM generates are published
  as connection details for other resources to consume.
- **Continuous drift correction.** If somebody edits a model in the Admin UI, the
  provider notices and puts it back.
- **No Terraform state to babysit.** The reconciler is the state; the cluster is the
  source of truth.

## Quick start

### 1. Install Crossplane

```console
helm install crossplane --namespace crossplane-system --create-namespace \
  crossplane-stable/crossplane
```

### 2. Install the provider

```console
kubectl apply -f https://raw.githubusercontent.com/TheCodingSheikh/provider-litellm/main/examples/install.yaml
```

Or inline:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-litellm
spec:
  package: xpkg.upbound.io/thecodingsheikh/provider-litellm:v2.0.1
```

The package is published on the
[Upbound Marketplace](https://marketplace.upbound.io/providers/thecodingsheikh/provider-litellm);
check [releases](https://github.com/TheCodingSheikh/provider-litellm/releases) for newer
versions.

### 3. Point it at your LiteLLM proxy

Credentials are a JSON document in a Secret:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: litellm-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_base": "https://litellm.example.com",
      "api_key": "sk-1234567890"
    }
---
apiVersion: litellm.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: litellm-creds
      namespace: crossplane-system
      key: credentials
```

`api_base` and `api_key` are required. `insecure_skip_verify` and `litellm_changed_by`
are optional. On Crossplane v2 use `litellm.m.crossplane.io/v1beta1` with either
`ProviderConfig` (namespaced) or `ClusterProviderConfig`.

### 4. Declare a model

```yaml
apiVersion: litellm.crossplane.io/v1alpha1
kind: Model
metadata:
  name: gpt-4o-mini
spec:
  forProvider:
    modelName: gpt-4o-mini-proxy
    customLlmProvider: openai
    baseModel: gpt-4o-mini
    mode: chat
    inputCostPerMillionTokens: 0.15
    outputCostPerMillionTokens: 0.6
    litellmCredentialNameRef:
      name: openai-credential
```

```console
kubectl get models.litellm.crossplane.io
NAME           READY   SYNCED   EXTERNAL-NAME
gpt-4o-mini    True    True     8f2c1e6a-...
```

Resources reference each other by name, so a `Team` can point at an `Organization` and a
`Model` at a `Credential` without you hardcoding the UUIDs LiteLLM generates:

```yaml
apiVersion: litellm.crossplane.io/v1alpha1
kind: Team
metadata:
  name: platform-team
spec:
  forProvider:
    teamAlias: platform
    maxBudget: 50
    organizationIdRef:
      name: acme
```

More examples for every resource live in [`examples/`](examples/) (cluster-scoped) and
[`examples/namespaced/`](examples/namespaced/).

## Managed resources

24 LiteLLM resources are exposed as CRDs, in both the cluster-scoped and namespaced API
groups.

| Kind | What it manages | Status |
| --- | --- | --- |
| `Model` | An LLM deployment on the proxy: upstream provider, base model, mode, per-token pricing | Verified |
| `Credential` | A reusable upstream provider credential that models reference by name | Verified |
| `Team` | A team, its budget, and the models it may call | Verified |
| `TeamMember` | Membership of a user in a team, with a role | Verified |
| `TeamMemberAdd` | Bulk membership management for a team | Verified |
| `TeamBlock` | Blocks a team from issuing requests | Verified |
| `Organization` | An organization grouping teams and budgets | Verified |
| `OrganizationMember` | Membership of a user in an organization | Verified |
| `User` | An end user, their budget and model access | Verified |
| `Budget` | A reusable budget: max spend, reset duration, TPM/RPM limits | Verified |
| `Tag` | A request tag used for routing, tracking and spend attribution | Verified |
| `AccessGroup` | A named group of models used for access control | Verified |
| `Guardrail` | A guardrail applied to requests passing through the proxy | Verified |
| `Prompt` | A managed prompt template | Verified |
| `Fallback` | A fallback chain for a model (general, context window, content policy) | Verified |
| `VectorStore` | A vector store registered with the proxy | Verified |
| `JWTKeyMapping` | Maps a JWT claim to an existing virtual key | Verified |
| `Key` | A virtual API key with its own budget and model scope | Verified |
| `KeyBlock` | Blocks a virtual key | Verified |
| `MCPServer` | An MCP server registered on the gateway | Verified |
| `UnifiedAccessGroup` | A unified access group across models and MCP servers | Verified |
| `SearchTool` | A web search tool available to the proxy | Verified |
| `Agent` | A managed agent | Verified |
| `Project` | A project grouping resources | Requires LiteLLM Enterprise |

Full schema reference for every field is on the
[Upbound Marketplace listing](https://marketplace.upbound.io/providers/thecodingsheikh/provider-litellm).

### Known limitations

All 24 resources are configured and generated. 23 were verified `Ready` end to end
against a live LiteLLM proxy in **both** API groups -- cluster-scoped
(`litellm.crossplane.io`) and namespaced (`litellm.m.crossplane.io`) -- covering
create, observe and delete, with every configured cross-resource reference
resolving. One cannot work here:

- **`Project`** — LiteLLM answers `/project/new` with
  `403 Project management is an enterprise feature`. It works only against a
  LiteLLM Enterprise deployment.

Two upstream LiteLLM server bugs are also worth knowing about, since neither is
something this provider can work around:

- Setting `budgetId` on an `Organization` makes LiteLLM's own `/organization/update`
  fail with a 500, because it passes `budget_id` twice into `BudgetNewRequest`. The
  organization can be created but never updated. The reference is available; the
  shipped example leaves it unset.
- Assigning a `Model` to a team returns a 403 on non-enterprise deployments, so
  `Model.teamId` is available as a reference but not exercised by the examples.

### OpenTofu

The provider embeds **OpenTofu**, not Terraform. Terraform is BSL licensed from
1.6 onward and this image redistributes the binary, so shipping it is not
permitted; OpenTofu is the MPL-2.0 fork and is CLI compatible. It is installed as
`terraform` because that is the executable name upjet invokes. Using OpenTofu also
provides write-only attribute support, which `JWTKeyMapping` and `Key.keyWo`
require and which Terraform 1.5.7 could not offer.

## Versioning

The provider version tracks the upstream `ncecere/litellm` Terraform provider version it
wraps: `v2.0.1` of this provider is built against `ncecere/litellm` v2.0.1.

## FAQ

**How is this different from the LiteLLM Terraform provider?**
Same idea, different control plane. The Terraform provider needs a plan/apply run and a
state file; this reconciles continuously from inside Kubernetes, corrects drift, and
composes with the rest of Crossplane. If your platform is already Kubernetes and Argo CD,
this fits without adding a Terraform runner.

**Does it work with Crossplane v2 namespaced resources?**
Yes. Every resource is served in both `litellm.crossplane.io` (cluster-scoped, Crossplane
v1 style) and `litellm.m.crossplane.io` (namespaced, Crossplane v2).

**Does it work with Argo CD or Flux?**
Yes. Managed resources are plain Kubernetes objects, so any GitOps controller can apply
them. Crossplane does the reconciliation against LiteLLM.

**Does it manage the LiteLLM deployment itself?**
No. It configures a running LiteLLM proxy through its management API. Deploy the proxy
with the [LiteLLM Helm chart](https://docs.litellm.ai/docs/proxy/deploy) or the official
Terraform modules, then point a `ProviderConfig` at it.

**Which LiteLLM versions are supported?**
Verified against LiteLLM 1.99.0 using OpenTofu 1.12.6 and ncecere/litellm v2.1.0. Anything exposing the same management API should work.

**Can I self-host the package?**
Yes. `make build` produces an xpkg you can push to any OCI registry, and
`spec.package` accepts any reference.

## Developing

Run the code-generation pipeline:

```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

Adding or changing a resource means an entry in `config/external_name.go`, a configurator
in `config/litellm/config.go`, then `make generate`.

## Report a bug

For filing bugs, suggesting improvements, or requesting new features, please open an
[issue](https://github.com/TheCodingSheikh/provider-litellm/issues).

## License

Apache 2.0. See [LICENSE](LICENSE).
