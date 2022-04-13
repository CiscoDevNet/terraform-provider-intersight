### Resource Creation

```hcl
resource "intersight_software_hyperflex_distributable" "software_hyperflex_distributable1" {
  name      = "software_hyperflex_distributable1"
  sha512sum = "<sha_512_checksum>"
  size      = 7471044747
  component_meta {
    class_id         = "firmware.ComponentMeta"
    object_type      = "firmware.ComponentMeta"
    component_label  = "BIOS"
    is_oob_supported = false
  }
  model = ""
  mdfid = ""
  release {
    object_type = "softwarerepository.Release"
    moid        = var.release
  }
  catalog {
    object_type = "softwarerepository.Catalog"
    moid        = var.catalog
  }
  vendor = "Cisco"
  distributable_metas {
    object_type = "firmware.DistributableMeta"
    class_id    = "firmware.DistributableMeta"
  }
  supported_models = ["HyperFlex Data Platform"]
}

variable "release" {
  type        = string
  description = " value for release"
}

variable "catalog" {
  type        = string
  description = "value for catalog"
}
```