### Resource Creation

```hcl
resource "intersight_kubernetes_virtual_machine_node_profile" "kubernetes_virtual_machine_node_profile1" {
  description = "kubernetes virtual machine node profile"
  name        = "kubernetes_virtual_machine_node_profile1"
  nr_type        = "instance"
  action      = "Deploy"
  config_context {
    control_action = "Deploy"
    error_state    = "Pre-config-error"
    object_type    = "policy.ConfigContext"
  }
  cloud_provider = "noProvider"
}
```
