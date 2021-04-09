### Resource Creation

```hcl
resource "intersight_software_hcl_meta" "software_hcl_meta1" {
  name         = "software_hcl_meta1"
  description  = "software_hcl_meta"
  release_date = "2021-01-30T000000Z"
  sha_512_sum  = "<sha_512_checksum>"
  size         = 7471044747
  component_meta = [{
    object_type     = "firmware.ComponentMeta",
    component_label = "BIOS"
  }]
  content_type  = "Full"
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