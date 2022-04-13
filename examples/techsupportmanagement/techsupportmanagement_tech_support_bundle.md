### Resource Creation

```hcl
resource "intersight_techsupportmanagement_tech_support_bundle" "techsupportmanagement_tech_support_bundle" {
  platform_type = "UCSD"
  target_resource {
    class_id    = "techsupportmanagement.TechSupportBundle"
    object_type = "techsupportmanagement.TechSupportBundle"
  }
  device_registration {
    moid        = var.asset_deviceRegistration
    object_type = "asset.DeviceRegistration"
  }
}

variable "asset_deviceRegistration" {
  type        = string
  description = "value for moid"
}
```