### Resource Creation

```hcl
resource "intersight_techsupportmanagement_tech_support_bundle" "techsupportmanagement_tech_support_bundle" {
  platform_type = "UCSD"
  target_resource = {
    object_type   = "techsupportmanagement.TechSupportBundle"
    platform_type = "UCSD"
  }
  device_registration {
    moid        = var.asset.deviceRegistration
    object_type = "asset.DeviceRegistration"
  }
}
```