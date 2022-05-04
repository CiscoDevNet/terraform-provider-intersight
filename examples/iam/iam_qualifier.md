### Resource Creation

```hcl
resource "intersight_iam_qualifier" "iam_qualifier1" {
  usergroup {
    object_type = "iam.UserGroup"
    idp {
      object_type          = "iam.Idp"
      domain_name          = "cisco.com"
      enable_single_logout = true
      name                 = "Cisco"
      account {
        object_type = "iam.Account"
        moid        = var.iam_account
      }
      type = "saml"
    }
    permissions = [
      {
        moid        = var.iam_permission
        object_type = "iam.Permission"
      }
    ]
  }
}

variable "iam_permission" {
  type        = string
  description = "value for iam permission"
}

variable "iam_account" {
  type        = string
  description = "value for iam_account"
}
```