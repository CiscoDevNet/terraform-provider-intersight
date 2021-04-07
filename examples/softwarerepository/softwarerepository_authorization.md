### Resource Creation

```hcl
resource "intersight_softwarerepository_authorization" "softwarerepository_authorization" {
    password        = "changeMe"
    repository_type = "IntersightCloud"
    UserId          = "user_1"
    account {
        object_type = "iam.Account"
        moid = intersight_account_iam.iam1.id
    }
}
```