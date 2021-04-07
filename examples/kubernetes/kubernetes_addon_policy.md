### Resource Creation

```hcl
resource "intersight_kubernetes_addon_policy" "kubernetes_addon_policy1" {
    description = "kubernetes addon policy"
    name = "kubernetes_addon_policy1"
    addon_configuration {
        install_strategy = "Always"
        release_name = "helm-release-3"
        upgrade_strategy = "UpgradeOnly"
    }
    addon_definintion {
        chart_url = "/opt/ccp/charts/docker-registry.tgz"
        default_install_strategy = "Always"
        default_namespace = "iks"
        default_upgrade_strategy = "UpgradeOnly"
        description = "A docker registry"
        digest = "e8b9e7e3aeddc98f480af2b47785d4fe41734571a67ffcadc92d0148c5b6305a"
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```