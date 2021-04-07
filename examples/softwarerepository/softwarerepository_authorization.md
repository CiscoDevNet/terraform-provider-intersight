### Resource Creation

```hcl
resource "intersight_softwarerepository_authorization" "softwarerepository_authorization" {
  password        = "ChangeMe"
  repository_type = "IntersightCloud"
  user_id         = "user_1"
  account {
    object_type = "iam.Account"
    moid        = intersight_account_iam.iam1.id
  }
}
```