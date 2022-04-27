### Resource Creation

```hcl
resource "intersight_hyperflex_ucsm_config_policy" "hyperflex_ucsm_config_policy1" {
  name        = "hyperflex_ucsm_config_policy1"
  description = "hyperflex ucsm configuration policy"
  kvm_ip_range {
    end_addr    = "10.10.10.100"
    gateway     = "10.10.10.1"
    netmask     = "255.255.255.0"
    start_addr  = "10.10.10.10"
    object_type = "hyperflex.IpAddrRange"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  server_firmware_version = "3.5(2h)"
}
```