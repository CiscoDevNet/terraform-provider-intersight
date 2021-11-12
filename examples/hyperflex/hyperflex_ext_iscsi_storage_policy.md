### Resource Creation

```hcl
resource "intersight_hyperflex_ext_iscsi_storage_policy" "hyperflex_ext_iscsi_storage_policy1" {
  description = "hyperflex ext iscsi storage policy"
  name        = "hyperflex_ext_iscsi_storage_policy1"
  admin_state = true
  exta_traffic {
    name        = "exta_traffic1"
    vlan_id     = 100
    object_type = "replication.NamedVsan"
  }
  extb_traffic {
    name        = "extb_traffic1"
    vlan_id     = 200
    object_type = "replication.NamedVsan"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```