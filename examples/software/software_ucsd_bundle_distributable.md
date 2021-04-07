### Resource Creation

```hcl
resource "intersight_software_ucsd_bundle_distributable" "software_ucsd_bundle_distributable1" {
  name        = "software_ucsd_bundle_distributable1"
  sha_512_sum = "<sha_512_checksum>"
  size        = 7471044747
  component_meta = [{
    object_type     = "firmware.ComponentMeta"
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

  image_path = ""
  vendor     = "Cisco"
  distributable_metas = [{
    object_type = "firmware.DistributableMeta"
  }]
  release_url      = "< url for release notes of this image >"
  suppotred_models = ["HyperFlex Data Platform"]
  version          = "6.7.4.1"
}
```