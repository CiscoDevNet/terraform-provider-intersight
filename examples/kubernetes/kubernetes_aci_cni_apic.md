### Resource Creation

```hcl
resource "intersight_kubernetes_aci_cni_apic" "kubernetes_aci_cni_apic1" {
    num_aci_cni_profiles = 3
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
    registered_device {
        moid        = intersight_registered_device.device1.id
        object_type = "asset.DeviceRegistrations"
    }
}
```