resource "intersight_deviceconnector_policy" "dcp1" {
  name            = "device_con1"
  description     = "test policy"
  lockout_enabled = true
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

/*
SAMPLE PAYLOAD
-----------------
DeviceconnectorPolicyApi: {
  "Name": "AUTO_POLICY_CLO_CRR",
  "Description": "CLO Testing",
  "Tags": [{"Key": "clo", "Value": "True"}],
  "LockoutEnabled": True,
}
*/