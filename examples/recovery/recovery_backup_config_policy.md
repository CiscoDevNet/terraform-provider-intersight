### Resource Creation

```hcl
resource "intersight_recovery_backup_config_policy" "recovery_backup_config_policy1" {
  name             = "recovery_backup_config_policy1"
  description      = "recovery_backup_config_policy"
  file_name_prefix = "file_name"
  location_type    = ""
  password         = "ChangeMe"
  path             = "path/to/backup"
  protocol         = "SCP"
  retention_count  = 120
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  user_name = "backup_server_name"
}

 variable "organization" {
   type = string
   description = "value for organization"
 }
```