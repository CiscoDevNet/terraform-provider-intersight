### Resource Creation

```hcl
resource "intersight_softwarerepository_category_mapper" "softwarerepository_category_mapper1" {
  name      = "softwarerepository_category_mapper1"
  category  = "Integrated Services Routers"
  file_type = "OperatingSystemFile"
  source    = "IntersightCloud"
}
```