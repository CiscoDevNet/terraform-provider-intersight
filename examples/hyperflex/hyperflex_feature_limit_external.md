### Resource Creation

```hcl
resource "intersight_hyperflex_feature_limit_external" "hyperflex_feature_limit_external1" {
    app_catalog {
        object_type = "hyperflex.AppCatalog"
        moid = "hyperflex_app_catalog"
    }
    feature_limit_entries [
      {
        constraint {
          hxdp_version = "^[3-9]\\.[0-9]"
          hypervisor_type = "ESXi"
          mgmt_platform = "FI"
          object_type = "hyperflex.AppSettingConstraint"
          server_model = "^HX"
        }
        name = "MAX_EXT_NODE"
        object_type = "hyperflex.FeatureLimitEntry"
        value = "34"
      }
    ]
    parent {
        object_type = "hyperflex.AppCatalog"
        moid = var.hyperflex_app_catalog
    }
}
```