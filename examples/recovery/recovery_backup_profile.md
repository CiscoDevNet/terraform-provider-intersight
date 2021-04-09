### Resource Creation

```hcl
resource "intersight_recovery_backup_profile" "recovery_backup_profile1" {
  name        = "recovery_backup_profile1"
  description = "recovery_backup_profile"
  type        = "instance"
  config_context {
    object_type    = "policy.ConfigContext"
    control_action = "deploy"
  }
  enabled = true
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  backup_config {
    object_type = "recovery.BackupProfile"
    moid        = var.recovery_backup_profile
  }
  schedule_config {
    object_type = "recovery.ScheduleConfigPolicies"
    moid        = var.recovery_schedule_config_policy
  }
}
```