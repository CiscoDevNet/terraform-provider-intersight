### Resource Creation

```hcl
resource "intersight_appliance_device_claim" "appliance_device_claim1" {
  hostname      = "10.106.233.221"
  platform_type = "IMC"
  request_id    = "718"
  username      = "admin"
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