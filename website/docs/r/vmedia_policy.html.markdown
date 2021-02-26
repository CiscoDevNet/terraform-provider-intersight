---
subcategory: "vmedia"
layout: "intersight"
page_title: "Intersight: intersight_vmedia_policy"
description: |-
  Policy to configure virtual media settings on endpoint.
---

# Resource: intersight_vmedia_policy
Policy to configure virtual media settings on endpoint.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of the Virtual Media service on the endpoint. 
* `encryption`:(bool) If enabled, allows encryption of all Virtual Media communications. 
* `low_power_usb`:(bool) If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. 
* `mappings`:(Array)
This complex property has following sub-properties:
  + `authentication_protocol`:(string) Type of Authentication protocol when CIFS is used for communication with the remote server.* `none` - No authentication is used.* `ntlm` - NT LAN Manager (NTLM) security protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2.* `ntlmi` - NTLMi security protocol. Use this option only when you enable Digital Signing in the CIFS Windows server.* `ntlmv2` - NTLMv2 security protocol. Use this option only with Samba Linux.* `ntlmv2i` - NTLMv2i security protocol. Use this option only with Samba Linux.* `ntlmssp` - NT LAN Manager Security Support Provider (NTLMSSP) protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2.* `ntlmsspi` - NTLMSSPi protocol. Use this option only when you enable Digital Signing in the CIFS Windows server. 
  + `device_type`:(string) Type of remote Virtual Media device.* `cdd` - Uses compact disc drive as the virtual media mount device.* `hdd` - Uses hard disk drive as the virtual media mount device. 
  + `file_location`:(string) Remote location of image. Preferred format is 'hostname/filePath/fileName'. 
  + `host_name`:(string) IP address or hostname of the remote server. 
  + `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
  + `mount_options`:(string) Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\ For NFS, supported options are ro, rw, nolock, noexec, soft, port=VALUE, timeo=VALUE, retry=VALUE.\ For CIFS, supported options are soft, nounix, noserverino, guest.\ For CIFS version < 3.0, vers=VALUE is mandatory. e.g. vers=2.0\ For HTTP/HTTPS, the only supported option is noauto. 
  + `mount_protocol`:(string) Protocol to use to communicate with the remote server.* `nfs` - NFS protocol for vmedia mount.* `cifs` - CIFS protocol for vmedia mount.* `http` - HTTP protocol for vmedia mount.* `https` - HTTPS protocol for vmedia mount. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `password`:(string) Password associated with the username. 
  + `remote_file`:(string) The remote file location path for the virtual media mapping. Accepted formats are:HDD for CIFS/NFS: hostname-or-IP/filePath/fileName.img.CDD for CIFS/NFS: hostname-or-IP/filePath/fileName.iso.HDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.img.CDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.iso. 
  + `remote_path`:(string) URL path to the location of the image on the remote server. The preferred format is '/path'. 
  + `sanitized_file_location`:(string)(Computed) File Location in standard format 'hostname/filePath/fileName'. This field should be used to calculate config drift. User input format may vary while inventory will return data in format in compliance with mount option for the mount. Both will be converged to this standard format for comparison. 
  + `username`:(string) Username to log in to the remote server. 
  + `volume_name`:(string) Identity of the image for Virtual Media mapping. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 

## Usage Example
### Resource Creation
```hcl
resource "intersight_vmedia_policy" "vmedia1" {
  name          = "vmedia1"
  description   = "test policy"
  enabled       = true
  encryption    = true
  low_power_usb = true

  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}
```

## Import
`intersight_vmedia_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vmedia_policy.example 1234567890987654321abcde
``` 