### Resource Creation

```hcl
resource "intersight_asset_target" "asset_target1" {
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
  assist = []
  connections = [
    {
      credential            = []
      object_type           = "asset.IntersightDeviceConnectorConnection"
      additional_properties = ""
      class_id              = "asset.IntersightDeviceConnectorConnection"
    }
  ]
  moid = var.asset_target1
  name = "C240-WZP21330QS5"
  registered_device {
    moid        = var.registered_device
    object_type = "asset.DeviceRegistrations"
  }
  status      = "NotConnected"
  target_type = "IMCM5"
}

variable "registered_device" {
  type        = string
  description = "Moid of registered device"
}

variable "asset_target1" {
  type        = string
  description = "Moid of asset.Target Mo"
}

variable "account" {
  type        = string
  description = "Moid of iam.Account Mo"
}
```