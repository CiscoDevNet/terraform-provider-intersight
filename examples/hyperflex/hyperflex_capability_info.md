### Resource Creation

```hcl
resource "intersight_hyperflex_capability_info" "hyperflex_capability_info1" {
  app_catalog {
    object_type = "hyperflex.AppCatalog"
    moid        = "hyperflex_app_catalog"
  }
  capability_constraints = [
    {
      constraint_name  = "targetHxdpVersion"
      constraint_value = "^4\\.0.*$|^4\\.5.*$"
      object_type      = "hcl.Constraint"
    },
    {
      constraint_name  = "serverPlatform"
      constraint_value = "M4|M5"
      object_type      = "hcl.Constraint"
    },
    {
      constraint_name  = "mgmtPlatform"
      constraint_value = "FI|EDGE"
      object_type      = "hcl.Constraint"
    }
  ]
  name = "minUcsVersion"
  parent {
    object_type = "hyperflex.AppCatalog"
    moid        = var.hyperflex_app_catalog
  }
}
```