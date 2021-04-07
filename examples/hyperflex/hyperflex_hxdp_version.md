### Resource Creation

```hcl
resource "intersight_hyperflex_hxdp_version" "hyperflex_hxdp_version1" {
    app_catalog {
        object_type = "hyperflex.AppCatalog"
        moid = "hyperflex_app_catalog"
    }
    parent {
      object_type = "hyperflex.AppCatalog"
      moid = var.hyperflex_app_catalog
    }
    Version = "4.0(2e)"
}
  
```