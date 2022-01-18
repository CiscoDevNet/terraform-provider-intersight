### Resource Creation

```hcl
resource "intersight_license_iwo_license_count" "license_iwo_license_count1" {
  account_license_data {
    moid        = var.account_license_moid
    object_type = "license.AccountLicenseData"
  }
}

variable "account_license_moid" {
  type        = string
  description = "Moid of account_license_data"
}
```