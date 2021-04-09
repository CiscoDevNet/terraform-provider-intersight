### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_backup_policy" "hyperflex_cluster_backup_policy1" {
  description = "hyperflex cluster backup policy"
  name        = "hyperflex_cluster_backup_policy1"
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