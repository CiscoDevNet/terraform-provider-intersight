### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_backup_policy_deployment" "hyperflex_cluster_backup_policy_deployment1" {
  discovered      = true
  source_detached = true
  target_detached = true
  replication_schedule {
    object_type     = "hyperflex.ReplicationSchedule"
    moid            = var.hyperflex_replication_schedule
    backup_interval = 1440
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
}
```