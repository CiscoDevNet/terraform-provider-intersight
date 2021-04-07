### Resource Creation

```hcl
resource "intersight_kubernetes_addon_repository" "kubernetes_addon_repository1" {
  insecure_skip_verification = false
  name                       = "kubernetes_addon_repository1"
  username                   = "admin1"
  catalog {
    moid        = var.kubernetes_catalog
    object_type = "kubernetes.Catalog"
  }
  registered_device {
    moid        = intersight_registered_device.device1.id
    object_type = "asset.DeviceRegistrations"
  }
}
```