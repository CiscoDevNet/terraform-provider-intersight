### Resource Creation

```hcl
resource "intersight_iam_session_limits" "iam_session_limits1" {
  idle_time_out    = 18000
  maximum_limit    = 128
  per_user_limit   = 32
  session_time_out = 31536000
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
  parent {
    moid        = var.iam_permission
    object_type = "iam.Permission"
  }
}
```