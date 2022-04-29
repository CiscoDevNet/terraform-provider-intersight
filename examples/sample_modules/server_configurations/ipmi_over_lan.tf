resource "intersight_ipmioverlan_policy" "tf_ipmi" {
  name           = "tf_ipmi"
  description    = "demo ipmi policy"
  enabled        = true
  privilege      = "admin"
  encryption_key = var.encryption_key
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
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