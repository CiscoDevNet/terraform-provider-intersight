### Resource Creation

```hcl
resource "intersight_vmedia_policy" "vmedia1" {
  name          = "vmedia1"
  description   = "demo vmedia policy"
  enabled       = true
  encryption    = true
  low_power_usb = true
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  mappings {
    class_id       = "vmedia.Mapping"
    device_type    = "cdd"
    file_location  = "infra-chx.auslab.cisco.com/software/linux/ubuntu-18.04.5-server-amd64.iso"
    host_name      = "infra-chx.auslab.cisco.com"
    mount_options  = "RO"
    mount_protocol = "nfs"
    object_type    = "vmedia.Mapping"
    remote_file    = "ubuntu-18.04.5-server-amd64.iso"
    remote_path    = "/iso/software/linux"
    volume_name    = "IMC_DVD"
  }
}

variable "organization" {
  type        = string
  description = "<value for organization>"
}
```