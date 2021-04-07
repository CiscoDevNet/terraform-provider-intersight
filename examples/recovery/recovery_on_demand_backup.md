### Resource Creation

```hcl
resource "intersight_recovery_on_demand_backup" "recovery_on_demand_backup1" {
    name = "recovery_on_demand_backup1"
    description = "recovery_on_demand_backup"
    file_name_prefix = "file_name"
    location_type = "Network Storage"
    retention_count = 120
    path = "172.29.109.234/var/backups/"
    protocol = "SCP"
    config_context {
        object_type = "policy.ConfigContext"
        control_action = "deploy"
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
    backup_config {
        object_type = "recovery.BackupProfile"
        moid = var.recovery_backup_profile
    }
    schedule_config {
        object_type = "recovery.ScheduleConfigPolicies"
        moid = var.recovery_schedule_config_policy
    }
}
```