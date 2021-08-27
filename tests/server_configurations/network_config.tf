resource "intersight_networkconfig_policy" "tf_network_config1" {
  name                     = "tf_network_config1"
  description              = "test policy"
  enable_dynamic_dns       = false
  preferred_ipv6dns_server = "::"
  enable_ipv6              = true
  enable_ipv6dns_from_dhcp = false
  preferred_ipv4dns_server = "10.10.10.1"
  alternate_ipv4dns_server = "10.10.10.1"
  alternate_ipv6dns_server = "::"
  dynamic_dns_domain       = ""
  enable_ipv4dns_from_dhcp = false
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_networkconfig_policy" "tf_network_config2" {
  name                     = "tf_network_config2"
  description              = "test policy"
  enable_dynamic_dns       = true
  enable_ipv4dns_from_dhcp = false
  enable_ipv6              = true
  enable_ipv6dns_from_dhcp = false
  preferred_ipv4dns_server = "10.10.10.1"
  alternate_ipv4dns_server = "10.10.10.1"
  dynamic_dns_domain       = "qw12"
  alternate_ipv6dns_server = "::"
  preferred_ipv6dns_server = "::"
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}