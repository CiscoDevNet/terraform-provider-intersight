### Resource Creation

```hcl
resource "intersight_iam_user_group" "iam_user_group1" {
  name = "iam_user_group1"
  idp {
    moid = var.idp_moid
    object_type = "iam.Idp"
  }
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

variable "idp_moid" {
  type        = string
  description = "value for idp"
}
```
