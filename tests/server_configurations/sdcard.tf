resource "intersight_sdcard_policy" "tf_sdcard" {
  name = "tf_sdcard"
  description = "test policy"
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
  partitions {
    nr_type = "OS"
    object_type = "sdcard.Partition"

    virtual_drives {
      enable = true
      object_type = "sdcard.OperatingSystem"
      additional_properties = jsonencode({
        Name = "Hypervisor"
      })
    }
  }
}
