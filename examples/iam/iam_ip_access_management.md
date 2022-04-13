### Resource Creation

```hcl
resource "intersight_iam_ip_access_management" "iam_ip_access_management1" {
  enable = true
  holder {
    moid        = var.iam_security_holder
    object_type = "iam.SecurityHolder"
  }
  parent {
    moid        = var.iam_system
    object_type = "iam.System"
  }
}

variable "iam_security_holder" {
  type        = string
  description = "value for moid"
}

variable "iam_system" {
  type        = string
  description = " value for moid"
}
```