### Resource Creation

```hcl
resource "intersight_hyperflex_feature_limit_internal" "hyperflex_feature_limit_internal1" {
  app_catalog {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
  feature_limit_entries = [
    {
      constraint = [{
        hxdp_version          = "^[3-9]\\.[0-9]"
        hypervisor_type       = "ESXi"
        mgmt_platform         = "FI"
        object_type           = "hyperflex.AppSettingConstraint"
        server_model          = "^HX"
        additional_properties = ""
        class_id              = "hyperflex.AppSettingConstraint"
        deployment_type       = "NA"
      }]
      name                  = "MAX_NODE"
      object_type           = "hyperflex.FeatureLimitEntry"
      value                 = "32"
      class_id              = "hyperflex.FeatureLimitEntry"
      additional_properties = ""
    },
    {
      constraint = [{
        hxdp_version          = "^[3-9]\\.[0-9]"
        hypervisor_type       = "ESXi"
        mgmt_platform         = "FI"
        object_type           = "hyperflex.AppSettingConstraint"
        server_model          = "^HX.*L$"
        additional_properties = ""
        class_id              = "hyperflex.AppSettingConstraint"
        deployment_type       = "NA"
      }]
      name                  = "MAX_NODE"
      object_type           = "hyperflex.FeatureLimitEntry"
      value                 = "8"
      class_id              = "hyperflex.FeatureLimitEntry"
      additional_properties = ""
    },
    {
      constraint = [{
        hxdp_version          = "^4\\.[5-9]|^4\\.0\\(2"
        hypervisor_type       = "ESXi"
        mgmt_platform         = "EDGE"
        object_type           = "hyperflex.AppSettingConstraint"
        server_model          = "^HX"
        additional_properties = ""
        class_id              = "hyperflex.AppSettingConstraint"
        deployment_type       = "NA"
      }]
      name                  = "MAX_NODE"
      object_type           = "hyperflex.FeatureLimitEntry"
      value                 = "8"
      class_id              = "hyperflex.FeatureLimitEntry"
      additional_properties = ""
    }
  ]
  parent {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
}
```