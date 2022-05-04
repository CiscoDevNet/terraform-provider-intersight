### Resource Creation

```hcl
resource "intersight_recovery_schedule_config_policy" "recovery_schedule_config_policy1" {
  name        = "recovery_schedule_config_policy1"
  description = "recovery_schedule_config_policy"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  backup_profiles {
    object_type = "recovery.BackupProfile"
    moid        = var.recovery_backup_profile
  }
  schedule {
    object_type    = "recovery.BackupSchedule"
    frequency_unit = "Daily"
    hours          = "8"
  }
}

variable "organization" {
  type        = string
  description = "value for organization"
}

variable "recovery_backup_profile" {
  type        = string
  description = "value recovery backup profile"
}
```