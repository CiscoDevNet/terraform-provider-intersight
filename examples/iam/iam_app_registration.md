### Resource Creation

```hcl
resource "intersight_iam_app_registration" "iam_app_registration1" {
  client_name         = "name1"
  client_type         = "confidential"
  revoke              = true
  renew_client_secret = true
  roles {
    moid        = var.iam_role
    object_type = "iam.Role"
    class_id    = "iam.Role"
  }
  permission {
    moid        = var.iam_permission
    object_type = "iam.Permission"
  }
}

variable "iam_permission" {
  type        = string
  description = "value for iam_permission"
}

variable "iam_role" {
  type        = string
  description = "value for iam_role"
}
```