resource "intersight_recovery_schedule_config_policy" "recovery_schedule_config_policy1" {
  name        = "recovery_schedule_config_policy1_test"
  description = "recovery_schedule_config_policy"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  schedule {
    object_type    = "recovery.BackupSchedule"
    frequency_unit = "Daily"
    hours          = "8"
    execution_time = "Thu, 26 Oct 2028 15:04:05 UTC"
  }
}
