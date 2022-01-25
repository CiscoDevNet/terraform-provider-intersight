### Resource Creation

```hcl
resource "intersight_kvm_session" "kvm_session1" {
  device {
    moid        = var.asset_device_registration
    object_type = "asset.DeviceRegistration"
  }
  server {
    moid        = var.compute_rack_unit
    object_type = "compute.RackUnit"
  }
  tunnel {
    moid        = var.kvm_tunnel
    object_type = "kvm.Tunnel"
  }
}

variable "asset_device_registration" {
  type        = string
  description = "Moid of asset_device_registration"
}

variable "compute_rack_unit" {
  type        = string
  description = "Moid of compute_rack_unit"
}

variable "kvm_tunnel" {
  type        = string
  description = "Moid of kvm_tunnel"
}
```