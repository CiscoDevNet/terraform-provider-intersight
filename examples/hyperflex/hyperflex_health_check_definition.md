### Resource Creation

```hcl
resource "intersight_hyperflex_health_check_definition" "hyperflex_health_check_definition1" {
  category                          = "Cluster Resiliency Check"
  common_causes                     = "Persistent drive on one of the nodes is in deny-list or unavailable."
  description                       = "Number of Persistent Drive Failures Tolerable"
  internal_name                     = "NumberOfPersistentDriveFailuresTolerable"
  name                              = "Number of Persistent Drive Failures Tolerable"
  resolution                        = "Check the status of the disk in HX Connect."
  script_execution_mode             = "ON_DEMAND"
  script_execution_on_compute_nodes = false
  target_execution_type             = "EXECUTE_ON_LEADER_NODE"
  timeout                           = 0
}
```