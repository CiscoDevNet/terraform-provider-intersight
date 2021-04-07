### Resource Creation

```hcl
resource "intersight_kubernetes_addon" "kubernetes_addon1" {
    install_strategy = "Always"
    name = "kubernetes_addon1"
    upgrade_strategy = "UpgradeOnly"
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
    parent {
        moid = var.kubernetes_addon
        object_type = "kubernetes.Addon"
    }
}
```