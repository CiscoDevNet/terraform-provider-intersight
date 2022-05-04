### Resource Creation

```hcl
resource "intersight_softwarerepository_operating_system_file" "osf11" {
  nr_version  = "ESXi 6.7 U2"
  description = "ESXi6.7U2 without answers"
  name        = "ESXi6.7 w/o cifs 21"

  nr_source {
    additional_properties = jsonencode({
      FileLocation = "10.10.10.1/Public/iso/esx67u2.iso"
      RemoteIp     = "10.10.10.1"
      RemoteShare  = "/Public/iso/"
      RemoteFile   = "esx67u2.iso"
      Username     = "user"
      Password     = "ChangeMe"
      MountOption  = "sec=ntlm"
    })
    object_type = "softwarerepository.CifsServer"
  }
  vendor = "VMware"
  catalog {
    moid        = var.catalog
    object_type = "softwarerepository.Catalog"
  }
}

variable "catalog" {
  type        = string
  description = "value for moid"
}
```