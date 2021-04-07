### Resource Creation

```hcl
resource "intersight_hyperflex_server_model" "hyperflex_server_model1" {
  app_catalog {
    object_type = "hyperflex.AppCatalog"
    moid        = "hyperflex_app_catalog"
  }
  parent {
    object_type = "hyperflex.AppCatalog"
    moid        = var.hyperflex_app_catalog
  }
  server_model_entries = [
    {
      constraint = {
        hxdp_version    = "^[3-9]\\.[0-9]"
        hypervisor_type = "ESXi"
        mgmt_platform   = "FI"
        object_type     = "hyperflex.AppSettingConstraint"
        server_model    = ""
      }
      name        = "SERVER_MODEL"
      object_type = "hyperflex.ServerModelEntry"
      value       = "^HX"
    },
    {
      constraint = {
        hxdp_version    = ".*"
        hypervisor_type = "ESXi"
        mgmt_platform   = "EDGE"
        object_type     = "hyperflex.AppSettingConstraint"
        server_model    = ""
      }
      name        = "SERVER_MODEL"
      object_type = "hyperflex.ServerModelEntry"
      value       = "HX.*M(4|5).*"
    }
  ]
}
```