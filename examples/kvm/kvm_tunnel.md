### Resource Creation

```hcl
resource "intersight_kvm_tunnel" "kvm_tunnel1" {
  device {
    moid        = var.asset_device_registration
    object_type = "asset.DeviceRegistration"
  }
  server {
    moid        = var.compute_rack_unit
    object_type = "compute.RackUnit"
  }
}
```