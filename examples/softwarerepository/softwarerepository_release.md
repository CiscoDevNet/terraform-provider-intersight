### Resource Creation

```hcl
resource "intersight_softwarerepository_release" "softwarerepository_release1" {
  nr_type = "FabricSwitch"
   catalog {
     object_type = "softwarerepository.Catalog"
     moid        = var.softwarerepository_catalog
   }
}

 variable "softwarerepository_catalog" {
   type = string
   description = "value for moid"
 }
```
