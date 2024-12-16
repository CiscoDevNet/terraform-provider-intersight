### Resource Creation

```hcl
resource "intersight_fabric_switch_profile" "fabric_switch_profile1" {
  name        = "fabric_switch_profile1"
  description = "demo fabric switch profile"
  nr_type        = "instance"
  action      = "No-op"
  switch_cluster_profile {
    moid        = intersight_fabric_switch_cluster_profile.fabric_switch_cluster_profile1.moid
    object_type = "fabric.SwitchClusterProfile"
  }
}
```
