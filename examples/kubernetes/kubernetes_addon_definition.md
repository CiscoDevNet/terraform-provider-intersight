### Resource Creation

```hcl
resource "intersight_kubernetes_addon_definition" "kubernetes_addon_definition1" {
  catalog {
    moid        = var.kubernetes_catalog
    object_type = "kubernetes.Catalog"
  }
  chart_url                = "/path/to/chart/url"
  default_install_strategy = "Always"
  default_namespace        = "iks"
  default_upgrade_strategy = "UpgradeOnly"
  description              = "A docker registry"
  digest                   = "e8b9e7e3aeddc98f480af2b47785d4fe41734571a67ffcadc92d0148c5b6305a"
  name                     = "essential-registry"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  parent {
    moid        = var.kubernetes_catalog
    object_type = "kubernetes.Catalog"
  }
  version = "1.8.3-cisco10-helm3"
}
```