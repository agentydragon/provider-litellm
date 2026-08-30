# v2.1.0

Tracks the upstream `ncecere/litellm` **v2.1.0** release, replacing
the pinned development-branch build used by the v2.0.2-rc.1 prerelease.

## Changed

- **Wrapped Terraform provider moves to the released v2.1.0.** The previous
  prerelease built from an unreleased development commit; this one builds from a
  tag, so the provenance of what ships is a published upstream release.
- **36 new upstream fields**, most of them on `MCPServer`: OAuth token exchange
  (`issuer`, `tokenExchangeEndpoint`, `audience`), DCR bridging, BYOK metadata,
  `envVars`, timeouts and concurrency limits. `Key`, `Model`, `Team`,
  `Organization` and `Project` each gain lossless `*Json` siblings for
  heterogeneous metadata alongside their existing `map(string)` attributes.
- **`MCPServer.url` is now sensitive upstream**, so it is read from a Secret
  (`urlSecretRef`) rather than set inline. The shipped example was updated; any
  existing manifest setting `url:` directly must move to `urlSecretRef`.

## Verified

24 of 25 example objects reconcile `Ready` and `Synced` end to end against a live
LiteLLM proxy in **both** API groups, cluster-scoped and namespaced, covering
create, observe and delete, with every cross-resource reference resolving in each
group. `Project` remains unavailable because LiteLLM gates it behind Enterprise.

## Upstream patches

v2.1.0 fixed none of the compatibility problems this provider patches, and added
three more. `hack/patches` now carries eight fixes, each hunk commented with the
failure it prevents:

- a sensitive list of objects (`MCPServer.envVars`) that upjet cannot build a
  type for at all -- sensitivity moved to the nested `value` leaf, which is the
  only leaf holding a secret;
- `Team` retaining creation-recovery state when a create is rejected outright,
  which turned a transient dependency-ordering failure into a permanent one;
- `Prompt` absence being reported as a 400 that upstream declines to treat as
  absence, leaving the resource impossible to create.
