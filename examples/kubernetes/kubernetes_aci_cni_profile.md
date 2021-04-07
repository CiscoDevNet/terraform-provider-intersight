### Resource Creation

```hcl
resource "intersight_kubernetes_aci_cni_profile" "kubernetes_aci_cni_profile1" {
    description = "kubernetes aci cni profile"
    name = "kubernetes_aci_cni_profile1"
    type = "instance"
    node_svc_subnet_start = "10.0.0.0/8"
    pod_subnet_start = "10.0.0.0/8"
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