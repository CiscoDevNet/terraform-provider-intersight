### Resource Creation

```hcl
resource "intersight_software_hyperflex_bundle_distributable" "software_hyperflex_bundle_distributable1" {
  name      = "software_hyperflex_bundle_distributable1"
  sha512sum = "<sha_512_checksum>"
  size      = 7471044747
  component_meta = [{
    additional_properties = ""
    class_id              = "firmware.ComponentMeta"
    moid                  = ""
    selector              = ""
    object_type           = "firmware.ComponentMeta"
    component_label       = "BIOS"
    component_type        = ""
    description           = ""
    disruption            = ""
    image_path            = ""
    is_oob_supported      = false
    model                 = ""
    oob_manageability     = null
    packed_version        = ""
    redfish_url           = ""
    vendor                = ""
  }]
  model         = ""
  mdfid         = ""
  platform_type = ""
   release_date {
     object_type = "softwarerepository.Release"
     moid        = var.release
   }
   catalog {
     object_type = "softwarerepository.Catalog"
     moid        = var.catalog
   }
  vendor = "Cisco"
  distributable_metas = [{
    object_type           = "firmware.DistributableMeta"
    additional_properties = ""
    class_id              = "firmware.DistributableMeta"
    moid                  = ""
    selector              = ""
  }]
  supported_models = ["C240-M5"]
}

 variable "release" {
   type = string
   description = " value for release"
 }

 variable "catalog" {
   type =  string
   description = "value for catalog"
 }
```