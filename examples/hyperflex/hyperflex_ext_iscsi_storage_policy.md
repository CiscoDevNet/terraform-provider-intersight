### Resource Creation

```hcl
resource "intersight_hyperflex_ext_iscsi_storage_policy" "hyperflex_ext_iscsi_storage_policy1" {
  description = "hyperflex ext iscsi storage policy"
  name        = "hyperflex_ext_iscsi_storage_policy1"
  admin_state = true
  exta_traffic {
    name        = "exta_traffic1"
    vsan_id     = 100
    object_type = "replication.NamedVsan"
    moid        = var.replication_named_vsan
  }
  extb_traffic {
    name        = "extb_traffic1"
    vsan_id     = 200
    object_type = "replication.NamedVsan"
    moid        = var.replication_named_vsan
  }
  wwxn_prefix_range {
    end_addr    = "10.10.10.100"
    start_addr  = "10.10.10.10"
    object_type = "hyperflex.WwxnPrefixRange"
    moid        = var.hyperflex_wwxn_prefix_range
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
}
```