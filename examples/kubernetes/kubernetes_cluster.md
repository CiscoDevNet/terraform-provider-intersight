### Resource Creation

```hcl
resource "intersight_kubernetes_cluster" "kubernetes_cluster1" {
  connection_status = "Connected"
  name              = "kubernetes_cluster1"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  registered_devices = [{
    moid                  = var.registered_device
    object_type           = "asset.DeviceRegistrations"
    additional_properties = ""
    class_id              = "asset.DeviceRegistrations"
    selector              = null
  }]
}

variable "registered_device" {
  type        = string
  description = "Moid of registered_device"
}
```