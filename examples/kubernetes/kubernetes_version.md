### Resource Creation

```hcl
resource "intersight_kubernetes_version" "kubernetes_version1" {
  boot_iso {
    moid        = var.software_solution_distributable
    object_type = "software.SolutionDistributable"
  }
  catalog {
    moid        = "605e73ff7a6f722d304bd6c8"
    object_type = "kubernetes.Catalog"
  }
  kubernetes_version = "v1.19.5"
  name               = "iks-20210324-2-v1.19.5"
  ova_image_template {
    moid        = var.software_solution_distributable
    object_type = "software.SolutionDistributable"
  }
  parent {
    moid        = var.kubernetes_catalog
    object_type = "kubernetes.Catalog"
  }
  qcow2_node_template {
    moid        = var.software_solution_distributable
    object_type = "software.SolutionDistributable"
  }
}
```