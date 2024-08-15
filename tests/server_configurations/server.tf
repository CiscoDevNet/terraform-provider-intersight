resource "intersight_server_profile" "tf_server_server_configurations" {
  name = "tf_server_server_configurations"
  action = "No-op"
  tags {
    key = "server"
    value = "demo"
  }
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
  tags {
    key = "source"
    value = "terraform"
  }
  policy_bucket {
    moid = intersight_adapter_config_policy.tf_adapter_config.moid
    object_type = "adapter.ConfigPolicy"
  }
  policy_bucket {
    moid = intersight_boot_precision_policy.tf_boot_precision_server_configurations.moid
    object_type = "boot.PrecisionPolicy"
  }
  policy_bucket {
    moid = intersight_ipmioverlan_policy.tf_ipmi.moid
    object_type = "ipmioverlan.Policy"
  }
  policy_bucket {
    moid = intersight_kvm_policy.tf_kvm.moid
    object_type = "kvm.Policy"
  }
  policy_bucket {
    moid = intersight_iam_ldap_policy.tf_ldap2.moid
    object_type = "iam.LdapPolicy"
  }
  policy_bucket {
    moid = intersight_networkconfig_policy.tf_network_config2.moid
    object_type = "networkconfig.Policy"
  }
  policy_bucket {
    moid = intersight_ntp_policy.tf_ntp2.moid
    object_type = "ntp.Policy"
  }
  policy_bucket {
    moid = intersight_sdcard_policy.tf_sdcard.moid
    object_type = "sdcard.Policy"
  }
  policy_bucket {
    moid = intersight_smtp_policy.tf_smtp.moid
    object_type = "smtp.Policy"
  }
  policy_bucket {
    moid = intersight_snmp_policy.tf_snmp.moid
    object_type = "snmp.Policy"
  }
  policy_bucket {
    moid = intersight_sol_policy.tf_sol.moid
    object_type = "sol.Policy"
  }
  policy_bucket {
    moid = intersight_syslog_policy.tf_syslog.moid
    object_type = "syslog.Policy"
  }
  policy_bucket {
    moid = intersight_vmedia_policy.tf_vmedia.moid
    object_type = "vmedia.Policy"
  }
  # policy_bucket {
  #   moid = intersight_vnic_lan_connectivity_policy.tf_vnic_lan.moid
  #   object_type = "vnic.LanConnectivityPolicy"
  # }
  # policy_bucket {
  #   moid = intersight_vnic_san_connectivity_policy.tf_vnic_san.moid
  #   object_type = "vnic.SanConnectivityPolicy"
  # }
}