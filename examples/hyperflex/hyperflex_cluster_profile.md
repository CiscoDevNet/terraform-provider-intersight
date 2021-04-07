### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_profile" "hyperflex_cluster_profile1" {
    storage_data_vlan = {
        name = "hx-storage-data" 
        vlan_id = 27
    }
    mgmt_ip_address = "10.225.68.237"
    mac_address_prefix = "00:25:B5:D5"
    software_version {
      moid = var.hyperflex_software_version_policy
      object_type = "hyperflex.SoftwareVersionPolicy"
    }
    local_credential {
      moid = var.hyperflex_local_credential_policy
      object_type = "hyperflex.LocalCredentialPolicy"
    }
    sys_config {
      moid = var.hyperflex_sys_config_policy
      object_type = "hyperflex.SysConfigPolicy"
    }
    node_config {
      moid = var.hyperflex_node_config_policy
      object_type = "hyperflex.NodeConfigPolicy"
    }
    cluster_network = {
      moid = var.hyperflex_cluster_network_policy
      object_type = "hyperflex.ClusterNetworkPolicy"
    }
    cluster_storage = {
      moid = var.hyperflex_cluster_storage_policy
      object_type = "hyperflex.ClusterStoragePolicy"
    }
    vcenter_config = {
      moid = var.hyperflex_vcenter_config_policy
      object_type = "hyperflex.VcenterConfigPolicy"
    }
    hypervisor_type = "ESXi"
    storage_type = "HyperFlexDp"
    auto_support = {
      moid = var.hyperflex_auto_support_policy
      object_type = "hyperflex.AutoSupportPolicy"
    }
    proxy_setting = {
      moid = var.hyperflex_proxy_setting_policy
      object_type = "hyperflex.ProxySettingPolicy"
    }
    mgmt_platform = "EDGE"
    replication = 3
    description = "This is hyperflex cluster profile"
    tags {
        key = "test" 
        value = "ucsback-10G-3nodehx-cluster-"
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization_organization
    }
    name = "hyperflex_cluster_profile1"
}
```