# Provider LiteLLM

`provider-litellm` is a Crossplane provider built with
[Upjet](https://github.com/crossplane/upjet) on top of the
[`ncecere/litellm`](https://registry.terraform.io/providers/ncecere/litellm/latest)
Terraform provider. It exposes XRM-conformant managed resources for the
[LiteLLM](https://docs.litellm.ai/) proxy API, so a LiteLLM deployment can
be configured declaratively from Kubernetes.

Managed resources cover the LiteLLM control plane: `Model`, `Credential`,
`Team`, `TeamMember`, `Key`, `Organization`, `OrganizationMember`, `User`,
`Budget`, `Tag`, `AccessGroup`, `Guardrail`, `Prompt`, `Fallback`,
`VectorStore`, `MCPServer` and more. Both the cluster-scoped
(`litellm.crossplane.io`) and the namespaced Crossplane v2
(`litellm.m.crossplane.io`) API groups are served.

Configure the provider with a `ProviderConfig` referencing a Secret that
holds the LiteLLM endpoint and API key:

```json
{
  "api_base": "https://litellm.example.com",
  "api_key": "sk-1234567890"
}
```

The optional `insecure_skip_verify` and `litellm_changed_by` keys are also
accepted.
