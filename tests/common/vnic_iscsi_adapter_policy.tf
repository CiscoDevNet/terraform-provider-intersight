
resource "intersight_vnic_iscsi_adapter_policy" "vnic_iscsi_adapter_policy" {
  name                 = "vnic_iscsi_adapter_policy1"
  description          = "vnic iscsi adapter policy"
  dhcp_timeout         = 60
  lun_busy_retry_count = 15
  connection_time_out  = 15
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
