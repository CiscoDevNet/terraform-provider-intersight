### Resource Creation

```hcl
resource "intersight_kvm_tunnel" "kvm_tunnel1" {
  device {
    moid        = var.device_registration
    object_type = "asset.DeviceRegistration"
  }
  server {
    moid        = var.compute_rack_unit_moid
    object_type = "compute.RackUnit"
  }
}

variable "device_registration" {
  type        = string
  description = "Moid of asset_device_registration"
}

variable "compute_rack_unit_moid" {
  type        = string
  description = "Moid of compute_rack_unit"
}
```