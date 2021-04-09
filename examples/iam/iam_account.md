### Resource Creation

```hcl
resource "intersight_iam_account" "iam_account1" {
  name = "iam_account1"
  resource_limits {
    moid        = var.iam_resource_limits
    object_type = "iam.ResourceLimits"
  }
  security_holder {
    moid        = var.iam_security_holder
    object_type = "iam.SecurityHolder"
  }
  session_limits {
    moid        = var.iam_session_limits
    object_type = "iam.SessionLimits"
  }
}
```