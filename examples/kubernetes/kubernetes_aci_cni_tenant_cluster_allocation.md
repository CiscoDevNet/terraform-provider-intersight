### Resource Creation

```hcl
resource "intersight_kubernetes_aci_cni_tenant_cluster_allocation" "kubernetes_aci_cni_tenant_cluster_allocation1" {
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```