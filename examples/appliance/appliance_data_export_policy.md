### Resource Creation

```hcl
resource "intersight_appliance_data_export_policy" "appliance_data_export_policy1" {
  name = "appliance_data_export1"
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