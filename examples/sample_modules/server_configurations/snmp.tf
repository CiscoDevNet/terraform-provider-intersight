resource "intersight_snmp_policy" "snmp1" {
  name                    = "snmp1"
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
  }
  snmp_traps {
    destination = "10.10.10.1"
    enabled     = false
    port        = 660
    type        = "Trap"
    user        = "demouser"
    nr_version     = "V3"
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

/*
SAMPLE PAYLOAD
-----------------
SnmpPolicyApi: {
    "Name": "AUTO_SNMP_POLICY_CRR",
    "Enabled": True,
    "Description": "SNMP Test policy",
    "Tags": [{"Key": "snmp", "Value": "snmp_policy"}],
    "SnmpPort": 1983,
    "AccessCommunityString": "dummy123",
    "CommunityAccess": "Disabled",
    "TrapCommunity": "TrapCommunity",
    "SysContact": "demouser",
    "SysLocation": "Karnataka",
    "EngineId": "vvb",
    "SnmpUsers": [{"Name": "demouser", "SecurityLevel": "AuthPriv",
                   "AuthType": "SHA", "AuthPassword": "changeMe",
                   "PrivacyType": "AES", "PrivacyPassword": "changeMe"}],
    "SnmpTraps": [{"Enabled": True, "Version": "V3",
                   "Type": "Trap", "User": "demouser",
                   "Destination": "10.10.10.1", "Port": 660}]
}
*/