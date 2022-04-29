resource "intersight_snmp_policy" "tf_snmp" {
  name                    = "tf_snmp"
  description             = "testing smtp policy"
  enabled                 = true
  snmp_port               = 1983
  access_community_string = "dummy123"
  community_access        = "Disabled"
  trap_community          = "TrapCommunity"
  sys_contact             = "aanimish"
  sys_location            = "Karnataka"
  engine_id               = "vvb"
  snmp_users {
    name             = "demouser"
    privacy_type     = "AES"
    auth_password    = var.auth_password
    privacy_password = var.privacy_password
    security_level   = "AuthPriv"
    auth_type        = "SHA"
    object_type      = "snmp.User"
  }
  snmp_traps {
    destination = "10.10.10.1"
    enabled     = false
    port        = 660
    type        = "Trap"
    user        = "demouser"
    nr_version  = "V3"
    object_type = "snmp.Trap"
  }
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}