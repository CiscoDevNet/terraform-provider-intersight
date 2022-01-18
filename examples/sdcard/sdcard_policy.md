### Resource Creation

```hcl
resource "intersight_sdcard_policy" "sdcard1" {
  name        = "sdcard1"
  description = "demo sd card policy"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
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

 variable "organization" {
   type = string
   description = "value for organization"
 }
```