### Resource Creation

```hcl
resource "intersight_hyperflex_hxdp_version" "hyperflex_hxdp_version1" {
  app_catalog {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
  parent {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
}
```