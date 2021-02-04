resource "intersight_ipmioverlan_policy" "ipmi1" {
  name           = "ipmi1"
  description    = "demo ipmi policy"
  enabled        = true
  privilege      = "admin"
  encryption_key = var.encryption_key
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}

/*
SAMPLE PAYLOAD
-----------------
IpmioverlanPolicyApi: {
  "Name": "AUTO_IPMI_POLICY_ADMIN_CRR",
  "Description": "IPMI Admin Testing",
  "Tags": [{"Key": "ipmiadmin", "Value": "555"}],
  "Enabled": True,
  "Privilege": "admin",
  "EncryptionKey": "12345678",
  "Profiles" : [{
    "Moid":"1234567890",
    "ObjectType":"server.Profile"
  }]
}
*/