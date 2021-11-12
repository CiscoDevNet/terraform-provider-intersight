### Resource Creation

```hcl
resource "intersight_hyperflex_server_model" "hyperflex_server_model1" {
  app_catalog {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
  parent {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
  server_model_entries = [
    {
      constraint = [{
        hxdp_version          = "^[3-9]\\.[0-9]"
        hypervisor_type       = "ESXi"
        mgmt_platform         = "FI"
        object_type           = "hyperflex.AppSettingConstraint"
        server_model          = ""
        class_id              = "hyperflex.AppSettingConstraint"
        additional_properties = ""
        deployment_type       = "NA"
      }]
      name                  = "SERVER_MODEL"
      object_type           = "hyperflex.ServerModelEntry"
      value                 = "^HX"
      additional_properties = ""
      class_id              = "hyperflex.ServerModelEntry"
    },
    {
      constraint = [{
        hxdp_version          = ".*"
        hypervisor_type       = "ESXi"
        mgmt_platform         = "EDGE"
        object_type           = "hyperflex.AppSettingConstraint"
        server_model          = ""
        class_id              = "hyperflex.AppSettingConstraint"
        additional_properties = ""
        deployment_type       = "NA"
      }]
      name                  = "SERVER_MODEL"
      object_type           = "hyperflex.ServerModelEntry"
      value                 = "HX.*M(4|5).*"
      additional_properties = ""
      class_id              = "hyperflex.ServerModelEntry"
    }
  ]
}
```