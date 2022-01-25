### Resource Creation

```hcl
resource "intersight_appliance_device_claim" "appliance_device_claim1" {
  device_id      = "WZP23350KMU"
  hostname       = "10.106.233.221"
  message        = "Endpoint claimed successfully"
  platform_type  = "IMC"
  request_id     = "718"
  security_token = "8707352528DC"
  status         = "completed"
  username       = "admin"
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