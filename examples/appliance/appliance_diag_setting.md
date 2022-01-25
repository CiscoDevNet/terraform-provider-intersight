### Resource Creation

```hcl
resource "intersight_appliance_diag_setting" "appliance_diag_setting1" {
  password = "ChangeMe"
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
}

variable "account" {
  type        = string
  description = "Moid of iam.Account Mo"
}
```