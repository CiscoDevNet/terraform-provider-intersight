### Resource Creation

```hcl
resource "intersight_kubernetes_cluster" "kubernetes_cluster1" {
    connection_status = "Connected"
    name = "kubernetes_cluster1"
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
    registered_device [{
        moid        = intersight_registered_device.device1.id
        object_type = "asset.DeviceRegistrations"
    }]
}
```