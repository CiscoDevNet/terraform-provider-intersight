### Resource Creation

```hcl
resource "intersight_hyperflex_node_profile" "hyperflex_node_profile1" {
  description = "This is hyperflex_node_profile1"
  assigned_server = {
    moid = var.hyperflex_node_profile
  }
  cluster_profile = {
    moid = var.hyperflex_cluster_profile
  }
  name = "ucsback-10G-3nodehx-cluster-NP1"
}
```