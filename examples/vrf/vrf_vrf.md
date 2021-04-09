### Resource Creation

```hcl
resource "intersight_vrf_vrf" "vrf_vrf1" {
  name        = "vrf_vrf1"
  description = "virtual routing and forwarding"
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
}
```