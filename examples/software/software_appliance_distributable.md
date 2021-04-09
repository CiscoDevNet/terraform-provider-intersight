### Resource Creation

```hcl
resource "intersight_software_appliance_distributable" "software_appliance_distributable1" {
  name        = "software_appliance_distributable1"
  description = "software_appliance_distributable"
  sha_512_sum = "<sha_512_checksum>"
  size        = 8485032509
  version     = "1.0.9-214"
  component_meta = [{ object_type = "firmware.ComponentMeta"
    component_label = "BIOS"
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
    object_type = "firmware.DistributableMeta"
  }]
  release_url      = "< url for release notes of this image >"
  supported_models = ["C240-M5"]
}
```