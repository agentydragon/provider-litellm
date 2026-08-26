# v2.0.2-rc.1

Prerelease. Replaces the embedded Terraform CLI with OpenTofu, which resolves a
licensing problem and unblocks the last resource that could not be created.

## Changed

- **The provider now embeds OpenTofu 1.12.6 instead of Terraform 1.5.7.**
  Terraform is BSL licensed from 1.6 onward and this image redistributes the
  binary, so the provider had been pinned to the last MPL release, 1.5.7, frozen
  since 2023. OpenTofu is the MPL-2.0 fork and is CLI compatible, so upjet drives
  it unchanged; it is installed as `terraform` because that is the executable
  upjet invokes.

## Added

- **`JWTKeyMapping` can now be created**, not merely observed. It needs the
  write-only `key_wo` argument, which requires a CLI supporting write-only
  attributes (OpenTofu 1.11+); Terraform 1.5.7 could not provide that. The same
  applies to `Key.keyWo` and `send_invite_email`.

## Verified

23 of 24 resources reconcile `Ready` and `Synced` end to end against a live
LiteLLM proxy in **both** API groups, cluster-scoped and namespaced, covering
create, observe and delete, with every cross-resource reference resolving in each
group. `Project` remains unavailable because LiteLLM gates it behind Enterprise.

## Also in this prerelease

Tracks the upstream `ncecere/litellm` `issues` development branch (pinned to
commit `eb28657`) rather than the released v2.0.1, which adds
`litellm_jwt_key_mapping` and reworks `Key`, `Agent`, `Credential` and
`MCPServer`. That branch reports an absent resource as a hard error instead of
as not-found, which stops Terraform ever creating one; local patches under
`hack/patches` correct that and several related issues, each hunk commented with
the failure it prevents.
