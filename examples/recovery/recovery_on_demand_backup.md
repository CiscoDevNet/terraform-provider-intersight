### Resource Creation

```hcl
resource "intersight_recovery_on_demand_backup" "recovery_on_demand_backup1" {
  name             = "recovery_on_demand_backup1"
  description      = "recovery_on_demand_backup"
  file_name_prefix = "file_name"
  location_type    = "Network Storage"
  retention_count  = 120
  path             = "path/to/backup"
  protocol         = "SCP"

  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  config_result {
    object_type = "recovery.BackupProfile"
    moid        = var.recovery_backup_profile
  }
}

variable "recovery_backup_profile" {
  type        = string
  description = "value for recovery backup profile"
}

variable "organization" {
  type        = string
  description = "value for organization"
}
```