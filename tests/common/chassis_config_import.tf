
resource "intersight_chassis_config_import" "chassis_config_import1" {
  description   = "chassis configuration import"
  policy_prefix = "chassis_policy"
  profile_name  = "server_profile_import1"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
