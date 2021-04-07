### Resource Creation

```hcl
resource "intersight_kvm_session" "kvm_session1" {
    device {
        moid = var.asset_device_registration
        object_type = 'asset.DeviceRegistration'
    }
    server {
        moid = var.compute_rack_unit
        object_type = 'compute.RackUnit'
    }
    tunnel {
        moid = var.kvm_tunnel
        object_type = 'kvm.Tunnel'
    }
}
```