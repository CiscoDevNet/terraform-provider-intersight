data "intersight_hyperflex_server_firmware_version_entry" "firmware_version"{
}

resource "intersight_hyperflex_ucsm_config_policy" "hyperflex_ucsm_config_policy1" {
  name        = "hyperflex_ucsm_config_policy1"
  description = "hyperflex ucsm configuration policy"
  kvm_ip_range {
    end_addr   = "10.10.10.100"
    gateway    = "10.10.10.1"
    netmask    = "255.255.255.0"
    start_addr = "10.10.10.10"
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  server_firmware_version = data.intersight_hyperflex_server_firmware_version_entry.firmware_version.results[length(data.intersight_hyperflex_server_firmware_version_entry.firmware_version)-1].nr_version
}
