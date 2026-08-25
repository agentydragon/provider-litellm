# Provider LiteLLM

`provider-litellm` is a [Crossplane](https://crossplane.io/) provider that is
built using [Upjet](https://github.com/crossplane/upjet) code generation tools
and exposes XRM-conformant managed resources for the
[LiteLLM](https://docs.litellm.ai/) proxy API.

It wraps the [`ncecere/litellm`](https://registry.terraform.io/providers/ncecere/litellm/latest)
Terraform provider, letting you manage LiteLLM models, teams, keys, credentials
and organizations as Kubernetes resources.

## Getting Started

Install the provider:

```console
kubectl apply -f examples/install.yaml
```

That installs `xpkg.upbound.io/thecodingsheikh/provider-litellm:v2.0.1`. Check
the [releases](https://github.com/thecodingsheikh/provider-litellm/releases) for
newer versions.

The package version tracks the upstream `ncecere/litellm` Terraform provider
version it wraps, so `v2.0.1` of this provider is built against
`ncecere/litellm` v2.0.1.

Then create a `ProviderConfig` pointing at a Secret holding your LiteLLM
endpoint and API key:

```console
kubectl apply -f examples/cluster/providerconfig/
```

## Developing

Run code-generation pipeline:

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

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/thecodingsheikh/provider-litellm/issues).
