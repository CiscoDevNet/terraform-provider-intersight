
resource "intersight_hyperflex_sys_config_policy" "hyperflex_sys_config_policy1" {
  dns_servers     = ["8.8.8.8"]
  ntp_servers     = ["10.10.58.51"]
  timezone        = "Asia/Calcutta"
  dns_domain_name = "cisco.com"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  name = "hyperflex_sys_config_policy1"
}
