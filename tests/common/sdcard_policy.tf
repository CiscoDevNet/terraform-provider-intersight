
resource "intersight_sdcard_policy" "sdcard1" {
  name        = "sdcard1"
  description = "demo sd card policy"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  partitions {
    type        = "OS"
    object_type = "sdcard.Partition"

    virtual_drives {
      enable      = true
      object_type = "sdcard.OperatingSystem"
      additional_properties = jsonencode({
        Name = "Hypervisor"
      })
    }
  }
}
