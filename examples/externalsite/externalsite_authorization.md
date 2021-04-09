### Resource Creation

```hcl
resource "intersight_externalsite_authorization" "externalsite_authorization1" {
  password        = "ChangeMe"
  repository_type = "cisco"
  user_id         = "user1"
  account {
    object_type = "iam.Account"
    moid        = intersight_account_iam.iam1.id
  }
}
```