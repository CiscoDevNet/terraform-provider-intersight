### Resource Creation

```hcl
resource "intersight_tam_advisory_info" "tam_advisory_info" {
  state = "active"
  account {
    object_type = "iam.Account"
    moid        = intersight_account_iam.iam1.id
  }
}
```