# v2.0.1

First release of `provider-litellm`, a Crossplane provider for the
[LiteLLM](https://docs.litellm.ai/) proxy API, generated with
[Upjet](https://github.com/crossplane/upjet) from the
[`ncecere/litellm`](https://registry.terraform.io/providers/ncecere/litellm/latest)
Terraform provider. The package version tracks the upstream provider version it
wraps.

## Managed resources

All 23 upstream resources are exposed, in both the cluster-scoped
(`litellm.crossplane.io`) and namespaced Crossplane v2
(`litellm.m.crossplane.io`) API groups:

`AccessGroup`, `Agent`, `Budget`, `Credential`, `Fallback`, `Guardrail`, `Key`,
`KeyBlock`, `MCPServer`, `Model`, `Organization`, `OrganizationMember`,
`Project`, `Prompt`, `SearchTool`, `Tag`, `Team`, `TeamBlock`, `TeamMember`,
`TeamMemberAdd`, `UnifiedAccessGroup`, `User`, `VectorStore`.

Cross-resource references are wired throughout, so a `Team` can point at an
`Organization`, a `Key` at a `Team`, a `Model` at a `Credential` and so on
without hard-coding generated identifiers.

## Verification

Sixteen resources were verified end to end against a live LiteLLM 1.99.0
instance, covering create, observe and delete, with every configured reference
resolving.

Seven are currently blocked by defects in the upstream Terraform provider rather
than by this package. They are shipped so they begin working as soon as upstream
is fixed:

- `Key`, `Project` and `Agent` — the upstream read path errors on an empty
  identifier instead of reporting the resource as absent, so creation never
  starts.
- `MCPServer` and `UnifiedAccessGroup` — the upstream read path expects a JSON
  object where LiteLLM 1.99.0 returns an array.
- `SearchTool` — the endpoint the provider calls returns 405 on LiteLLM 1.99.0.
- `KeyBlock` — needs a real key token, so it is blocked behind `Key`.

## Configuration

Point a `ProviderConfig` at a Secret holding your LiteLLM endpoint and API key:

```json
{
  "api_base": "https://litellm.example.com",
  "api_key": "sk-1234567890"
}
```

`insecure_skip_verify` and `litellm_changed_by` are also accepted.
