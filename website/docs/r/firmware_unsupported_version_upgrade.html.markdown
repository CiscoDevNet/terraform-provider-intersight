---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_unsupported_version_upgrade"
description: |-
  This represents an operation managed object used for upgrading equipment that cannot be discovered due to unsupported firmware. Currently, it only supports blade upgrades.
---

# Resource: intersight_firmware_unsupported_version_upgrade
This represents an operation managed object used for upgrading equipment that cannot be discovered due to unsupported firmware. Currently, it only supports blade upgrades.
## Argument Reference
The following arguments are supported:
* `checksum`:(HashMap) - The checksum of the downloaded file as calculated by the download plugin after successfully downloading a file. 
This complex property has following sub-properties:
  + `hash_algorithm`:(string) The hash algorithm used to calculate the checksum.* `crc` - A CRC hash as definded by RFC 3385. Generated with the IEEE polynomial.* `sha256` - A SHA256 hash as defined by RFC 4634. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `device`:(HashMap) -(Computed) A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `distributable`:(HashMap) - A reference to a firmwareDistributable resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `download_error`:(string) Any error encountered. Set to empty when download is in progress or completed. 
* `download_progress`:(int) The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. 
* `download_retries`:(int) The number of retries the plugin attempted before succeeding or failing the download. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `physical_identity`:(HashMap) - A reference to a equipmentPhysicalIdentity resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `upgrade_status`:(string) Workflow status of firmware upgrade.* `None` - Upgrade status is none when upgrade is in progress.* `Completed` - Upgrade completed successfully.* `Failed` - Upgrade status is failed when upgrade has failed. 


## Import
`intersight_firmware_unsupported_version_upgrade` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_firmware_unsupported_version_upgrade.example 1234567890987654321abcde
``` 