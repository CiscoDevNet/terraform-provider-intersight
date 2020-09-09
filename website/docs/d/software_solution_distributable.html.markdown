
---
layout: "intersight"
page_title: "Intersight: intersight_software_solution_distributable"
sidebar_current: "docs-intersight-data-source-software-solution-distributable"
description: |-
A solution image distributed by Cisco.
---

# Data Source: intersight_software._solution_distributable
A solution image distributed by Cisco.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bundle_type`:(string) The bundle type of the image, as published on cisco.com. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. 
* `download_count`:(int) The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `file_path`:(string) The path of the file in S3/minio bucket. 
* `guid`:(string) The unique identifier for an image in a Cisco repository. 
* `import_action`:(string) The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.* `None` - No action should be taken on the imported file.* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.* `PreCache` - Cache the file into the Intersight Appliance.* `Cancel` - The cancel import process for the file into the repository.* `Extract` - The action to extract the file in the external repository.* `Evict` - Evict the cached file from the Intersight Appliance. 
* `import_state`:(string) The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.* `ReadyForImport` - The image is ready to be imported into the repository.* `Importing` - The image is being imported into the repository.* `Imported` - The image has been extracted and imported into the repository.* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.* `Extracting` - Indicates that the image is being extracted into the repository.* `Extracted` - Indicates that the image has been extracted into the repository.* `Failed` - The image import from an external source to the repository has failed.* `MetaOnly` - The image is present in an external repository.* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. 
* `md5sum`:(string) The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `mdfid`:(string) The mdfid of the image provided by cisco.com. 
* `model`:(string) The endpoint model for which this firmware image is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the file. It is populated as part of the image import operation. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `platform_type`:(string) The platform type of the image. 
* `recommended_build`:(string) The build which is recommended by Cisco. 
* `release_notes_url`:(string) The url for the release notes of this image. 
* `sha512sum`:(string) The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `size`:(int) The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `software_advisory_url`:(string) The software advisory, if any, provided by the vendor for this file. 
* `software_type_id`:(string) The software type id provided by cisco.com. 
* `solution_name`:(string) The name of the solution in which the image belongs. 
* `sub_type`:(string) The type of the file like OS image, Script etc.* `osimage` - The solution OS image for deployment.* `script` - The Python script for the solution VM configuration and deployment. 
* `vendor`:(string) The vendor or publisher of this file. 
* `version`:(string) Vendor provided version for the file. 
