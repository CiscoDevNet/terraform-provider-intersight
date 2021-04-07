### Resource Creation

```hcl
resource "intersight_asset_target" "asset_target1" {
  account {
    object_type = "iam.Account"
    moid        = intersight_account_iam.iam1.id
  }
  assist = None
  connections = [
    {
      credential  = None
      object_type = "asset.IntersightDeviceConnectorConnection"
    }
  ]
  moid = var.asset_target1
  name = "C240-WZP21330QS5"
  registered_device {
    moid        = intersight_registered_device.device1.id
    object_type = "asset.DeviceRegistrations"
  }
  status      = NotConnected
  target_type = "IMCM5"
}
```