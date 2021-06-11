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
  mappings = [{
    additional_properties   = ""
    authentication_protocol = "none"
    class_id                = "vmedia.Mapping"
    device_type             = "cdd"
    file_location           = "infra-chx.auslab.cisco.com/software/linux/ubuntu-18.04.5-server-amd64.iso"
    host_name               = "infra-chx.auslab.cisco.com"
    is_password_set         = false
    mount_options           = "RO"
    mount_protocol          = "nfs"
    object_type             = "vmedia.Mapping"
    password                = ""
    remote_file             = "ubuntu-18.04.5-server-amd64.iso"
    remote_path             = "/iso/software/linux"
    sanitized_file_location = "infra-chx.auslab.cisco.com/software/linux/ubuntu-18.04.5-server-amd64.iso"
    username                = ""
    volume_name             = "IMC_DVD"
  }]
}
```