### Resource Creation

```hcl
resource "intersight_kubernetes_addon_policy" "kubernetes_addon_policy1" {
  description = "kubernetes addon policy"
  name        = "kubernetes_addon_policy1"
  addon_configuration {
    install_strategy = "Always"
    release_name     = "helm-release-3"
    upgrade_strategy = "UpgradeOnly"
    object_type      = "kubernetes.AddonConfiguration"
  }
  addon_definition {
    moid        = intersight_kubernetes_addon_definition.kubernetes_addon_definition1.moid
    object_type = "kubernetes.AddonDefinition"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```