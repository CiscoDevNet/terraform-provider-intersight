### Resource Creation

```hcl
resource "intersight_iam_user_group" "iam_user_group1" {
  name = "iam_user_group1"
  permissions {
    moid        = var.iam_permission
    object_type = "iam.Permission"
    class_id    = "iam.Permission"
  }
}

variable "iam_permission" {
  type        = string
  description = "value for iam_permission"
}
```