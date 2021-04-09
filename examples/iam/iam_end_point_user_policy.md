### Resource Creation

```hcl
resource "intersight_iam_end_point_user_policy" "user_policy1" {
  name        = "user_policy1"
  description = "test policy"

  password_properties {
    enforce_strong_password  = true
    enable_password_expiry   = true
    password_expiry_duration = 50
    password_history         = 5
    notification_period      = 1
    grace_period             = 2
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```