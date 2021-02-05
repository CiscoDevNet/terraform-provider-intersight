resource "intersight_softwarerepository_operating_system_file" "osf1" {
  nr_version = "ESXi 6.7 U2"
  description = "ESXi6.7U2 without answers"
  name = "ESXi6.7 w/o cifs 21"

  nr_source {
    additional_properties = jsonencode({
      FileLocation = "10.10.10.1/Public/iso/esx67u2.iso"
      RemoteIp = "10.10.10.1"
      RemoteShare = "/Public/iso/"
      RemoteFile = "esx67u2.iso"
      Username = "user"
      Password = "ChangeMe"
      MountOption = "sec=ntlm"
    })
    object_type = "softwarerepository.CifsServer"
  }
  vendor = "VMware"
  catalog {
    moid = var.catalog
  }
}

resource "intersight_softwarerepository_operating_system_file" "osf2" {

  nr_source {
    additional_properties = jsonencode({
      RemoteIp = "10.10.10.1"
      RemoteShare = "/Public/iso/"
      RemoteFile = "esx6u3.iso"
      Username = "user"
      Password = var.os_file_password_1
      MountOption = "sec=ntlm"
      FileLocation = "10.10.10.1/Public/iso/esx6u3.iso"
    })

    object_type = "softwarerepository.CifsServer"
  }
  name = "ESXi6.0-CIFS"
  vendor = "VMware"
  nr_version = "ESXi 6.0 U3"
  description = "ESXi6.0-CIFS"
  catalog {
    moid = var.catalog
  }
}
