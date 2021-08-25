resource "intersight_kvm_policy" "tf_kvm" {
  name                      = "tf_kvm"
  description               = "demo kvm policy"
  enabled                   = true
  maximum_sessions          = 3
  remote_port               = 2069
  enable_video_encryption   = true
  enable_local_server_video = true
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
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