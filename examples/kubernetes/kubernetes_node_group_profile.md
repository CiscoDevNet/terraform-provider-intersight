### Resource Creation

```hcl
resource "intersight_kubernetes_node_group_profile" "kubernetes_node_group_profile1" {
  description = "kubernetes node group profile"
  name        = "kubernetes_node_group_profile1"
  nr_type        = "instance"
  action      = "Deploy"
  config_context {
    control_action = "Deploy"
    error_state    = "Pre-config-error"
    object_type    = "policy.ConfigContext"
  }
  node_type = "Worker"
  cluster_profile {
    moid        = var.kubernetes_cluster_profile
    object_type = "kubernetes.ClusterProfile"
  }
}

variable "kubernetes_cluster_profile" {
  type        = string
  description = "Moid of kubernetes.ClusterProfile"
}
```
