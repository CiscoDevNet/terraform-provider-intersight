### Resource Creation

```hcl
resource "intersight_firmware_eula" "firmware_eula1" {
  account {
    object_type = "iam.Account"
    moid        = var.iam_account_moid
  }
}

variable "iam_account_moid" {
  type        = string
  description = "Moid of iam_account"
}
```