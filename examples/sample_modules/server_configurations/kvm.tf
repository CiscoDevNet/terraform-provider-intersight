resource "intersight_kvm_policy" "kvm1" {
  name                      = "kvm1"
  description               = "demo kvm policy"
  enabled                   = true
  maximum_sessions          = 3
  remote_port               = 2069
  enable_video_encryption   = true
  enable_local_server_video = true
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
  tags {
    key = "source"
    value = "terraform"
  }
}

/*
SAMPLE PAYLOAD
-----------------
KvmPolicyApi: {
  "Name": 'AUTO_KVM_POLICY_CRR',
  "Description": "testing KVM policy",
  "Tags": [{"Key": "kvm", "Value": "kvm_policy"}],
  "Enabled": True,
  "MaximumSessions": 4,
  "RemotePort": 1800,
  "EnableVideoEncryption": False,
  "EnableLocalServerVideo": False
}
*/