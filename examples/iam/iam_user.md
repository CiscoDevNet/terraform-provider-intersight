### Resource Creation

```hcl
resource "intersight_iam_user" "iam_user1" {
  email = "email@example.com"
  idpreference {
    moid        = var.iam_idp_reference
    object_type = "iam.IdpReference"
  }
  parent {
    moid        = var.iam_idp_reference
    object_type = "iam.IdpReference"
  }
  permissions {
    moid        = var.iam_idp_reference
    object_type = "iam.Permission"
  }
  user_id_or_email = "email@example.com"
}


variable "iam_idp_reference" {
  type        = string
  description = "value for iam idp ref"
}
```