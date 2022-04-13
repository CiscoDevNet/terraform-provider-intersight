### Resource Creation

```hcl
resource "intersight_iam_permission" "iam_permission1" {
  account {
    moid        = var.iam_account
    object_type = "iam.Account"
  }
  #description = "As a Device Technician you can claim a device view the device details license status a list of\nthe claimed devices and generate API keys. You cannot perform any other management or\nadministrative task in this role.\n"
  name = "Device Technician"
  parent {
    moid        = var.iam_account
    object_type = "iam.Account"
  }
  roles {
    moid        = var.iam_role
    object_type = "iam.Role"
  }
  users {
    moid        = var.iam_user
    object_type = "iam.User"
  }
}

variable "iam_role" {
  type        = string
  description = "value for iam role"
}

variable "iam_account" {
  type        = string
  description = "value for iam account"
}

variable "iam_user" {
  type        = string
  description = "value for iam user"
}
```