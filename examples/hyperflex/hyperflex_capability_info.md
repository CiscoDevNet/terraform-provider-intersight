### Resource Creation

```hcl
resource "intersight_hyperflex_capability_info" "hyperflex_capability_info1" {
  app_catalog {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
  capability_constraints {
    constraint_name  = "targetHxdpVersion"
    constraint_value = "^4\\.0.*$|^4\\.5.*$"
    object_type      = "hcl.Constraint"
    class_id         = "hcl.Constraint"
  }
  capability_constraints {
    constraint_name  = "serverPlatform"
    constraint_value = "M4|M5"
    object_type      = "hcl.Constraint"
    class_id         = "hcl.Constraint"
  }
  capability_constraints {
    constraint_name  = "mgmtPlatform"
    constraint_value = "FI|EDGE"
    object_type      = "hcl.Constraint"
    class_id         = "hcl.Constraint"
  }
  name = "minUcsVersion"
  parent {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
}
```