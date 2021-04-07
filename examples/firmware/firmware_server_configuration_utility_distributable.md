### Resource Creation

```hcl
resource "intersight_firmware_server_configuration_utility_distributable" "scu1" {
  name = "SCU-6.0.4c nfs"
  nr_source {
    object_type = "softwarerepository.CifsServer"
    additional_properties = jsonencode({
      FileLocation "10.10.10.1/Public/iso/scu-604c.iso"
      RemoteIp "10.10.10.1"
      RemoteShare "/Public/iso/"
      RemoteFile "scu-604c.iso"
      Username "user"
      Password "ChangeMe"
      MountOption "sec=ntlm"
    })
  }
  vendor = "Cisco"
  nr_version = "6.0.(4c)"
  supported_models = [
    "C-series"
  ]
  description = "Cisco SCU-6.0(4c)"
  catalog {
    moid = var.catalog
  }
}
```