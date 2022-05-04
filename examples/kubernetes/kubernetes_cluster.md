### Resource Creation

```hcl
resource "intersight_kubernetes_cluster" "kubernetes_cluster1" {
  connection_status = "Connected"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  registered_devices {
    moid        = var.registered_device
    object_type = "asset.DeviceRegistrations"
    class_id    = "asset.DeviceRegistrations"
  }
}

variable "registered_device" {
  type        = string
  description = "Moid of registered_device"
}
```