
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_operating_system_file"
sidebar_current: "docs-intersight-resource-softwarerepository-operating-system-file"
description: |-
  An operating system image that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository.
---

# Resource: intersight_softwarerepository._operating_system_file
An operating system image that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `catalog`:(Array with Maximum of one item) - A reference to a softwarerepositoryCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. 
* `download_count`:(int)(Computed) The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `import_action`:(string) The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.* `None` - No action should be taken on the imported file.* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.* `PreCache` - Cache the file into the Intersight Appliance.* `Cancel` - The cancel import process for the file into the repository.* `Extract` - The action to extract the file in the external repository.* `Evict` - Evict the cached file from the Intersight Appliance. 
* `import_state`:(string)(Computed) The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.* `ReadyForImport` - The image is ready to be imported into the repository.* `Importing` - The image is being imported into the repository.* `Imported` - The image has been extracted and imported into the repository.* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.* `Extracting` - Indicates that the image is being extracted into the repository.* `Extracted` - Indicates that the image has been extracted into the repository.* `Failed` - The image import from an external source to the repository has failed.* `MetaOnly` - The image is present in an external repository.* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. 
* `md5sum`:(string) The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the file. It is populated as part of the image import operation. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `sha512sum`:(string) The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `size`:(int) The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `software_advisory_url`:(string) The software advisory, if any, provided by the vendor for this file. 
* `source`:(Array with Maximum of one item) - Location of the file in an external repository. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vendor`:(string) The vendor or publisher of this file. 
* `version`:(string) Vendor provided version for the file. 
