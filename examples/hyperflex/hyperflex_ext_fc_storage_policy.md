### Resource Creation

```hcl
resource "intersight_hyperflex_ext_fc_storage_policy" "hyperflex_ext_fc_storage_policy1" {
  description = "hyperflex ext fc storage policy"
  name        = "hyperflex_ext_fc_storage_policy1"
  admin_state = true
  exta_traffic {
    name        = "exta_traffic1"
    vsan_id     = 100
    object_type = "replication.NamedVsan"
  }
  extb_traffic {
    name        = "extb_traffic1"
    vsan_id     = 200
    object_type = "replication.NamedVsan"
  }
  wwxn_prefix_range {
    end_addr    = "20:00:00:25:B5:B5"
    start_addr  = "20:00:00:25:A5:10"
    object_type = "hyperflex.WwxnPrefixRange"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```