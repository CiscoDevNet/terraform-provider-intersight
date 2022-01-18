### Resource Creation

```hcl
resource "intersight_kubernetes_cluster_addon_profile" "iks1" {
  name = "cluster_addon_profile1"

  addons {
    addon_policy {
      moid = intersight_kubernetes_addon_policy.kubernetes_addon_policy1.moid
    }
  }
  addons {
    addon_policy {
      moid = intersight_kubernetes_addon_policy.kubernetes_addon_policy1.moid
    }
  }
  organization {
    moid        = var.organization
    object_type = "organization.Organization"
  }
}
```