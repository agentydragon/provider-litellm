# v2.0.2-rc.1

Prerelease. Moves the wrapped Terraform provider from the released
`ncecere/litellm` v2.0.1 to the upstream `issues` development branch (pinned to
commit `eb28657`), which adds a resource and a large body of fixes, and carries
local patches that make the whole surface usable from Crossplane.

## Added

- **`JWTKeyMapping`** resource, mapping a JWT claim to an existing virtual key.
  Creating one needs Terraform 1.11 write-only attributes, so on this provider it
  can be observed and imported but not created -- see Known limitations.
- New upstream fields throughout, including `Key.routerSettings`,
  `Credential.credentialValuesJson`, `Team.teamId`, `KeyBlock.keyHash` and
  `Organization.softBudget`.

## Fixed

Six resources that could not previously be created now work: `Key`, `KeyBlock`,
`MCPServer`, `UnifiedAccessGroup`, `SearchTool` and `Agent`. Verified end to end
against a live LiteLLM proxy.

The upstream branch reports an absent resource as a hard error rather than as
not-found, which stops Terraform ever creating one, because a generated consumer
observes before it creates. This release carries patches that:

- report a resource whose state has no identifier as absent, instead of issuing a
  degenerate request (an empty query parameter, or a URL collapsing to a trailing
  slash that LiteLLM answers with a redirect the provider will not follow);
- classify `Budget` and `Tag` absence as not-found, including LiteLLM's
  404-inside-a-500 tag response;
- mark `Key.organizationId` computed, since LiteLLM derives it from the key's team
  and returns it even when the configuration leaves it null. Without this the apply
  is rejected as inconsistent and the key is tainted, forcing a replacement that
  `prevent_destroy` then blocks.

## Known limitations

- `Project` requires LiteLLM Enterprise (`403` on `/project/new`).
- `JWTKeyMapping` cannot be created without Terraform 1.11 write-only attributes;
  this provider is pinned to Terraform 1.5.7 because 1.6 and later are BSL licensed.
