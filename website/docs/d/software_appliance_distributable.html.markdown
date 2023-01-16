---
subcategory: "software"
layout: "intersight"
page_title: "Intersight: intersight_software_appliance_distributable"
description: |-
        Appliance image distributed by Cisco. This image is required to upgrade the on-premise Intersight Appliance.
        There are two use cases. In Intersight SaaS, the object represents a downloadable image, whereas on the
        Appliance the represents the image that is uploaded by the user and to be used for upgrade.

---

# Data Source: intersight_software_appliance_distributable
Appliance image distributed by Cisco. This image is required to upgrade the on-premise Intersight Appliance.
There are two use cases. In Intersight SaaS, the object represents a downloadable image, whereas on the
Appliance the represents the image that is uploaded by the user and to be used for upgrade.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_software_appliance_distributable.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bundle_type`:(string) The bundle type of the image, as published on cisco.com. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `download_count`:(int) The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `feature_source`:(string) The name of the feature to which the uploaded file belongs.* `System` - This indicates system initiated file uploads.* `OpenAPIImport` - This indicates an OpenAPI file upload. 
* `guid`:(string) The unique identifier for an image in a Cisco repository. 
* `image_type`:(string) The type of image which the distributable falls into according to the component it can upgrade. For e.g.; Standalone server, Intersight managed server, UCS Managed Fabric Interconnect. The field is used in private appliance mode, where image does not have description populated from CCO. 
* `import_action`:(string) The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.* `None` - No action should be taken on the imported file.* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.* `PreCache` - Cache the file into the Intersight Appliance.* `Cancel` - The cancel import process for the file into the repository.* `Extract` - The action to extract the file in the external repository.* `Evict` - Evict the cached file from the Intersight Appliance. 
* `import_state`:(string) The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.* `ReadyForImport` - The image is ready to be imported into the repository.* `Importing` - The image is being imported into the repository.* `Imported` - The image has been extracted and imported into the repository.* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.* `Extracting` - Indicates that the image is being extracted into the repository.* `Extracted` - Indicates that the image has been extracted into the repository.* `Failed` - The image import from an external source to the repository has failed.* `MetaOnly` - The image is present in an external repository.* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.* `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image. 
* `imported_time`:(string) The time at which this image or file was imported/cached into the repositry. if the 'ImportState' is 'Imported', the time at which this image or file was imported. if the 'ImportState' is 'Cached', the time at which this image or file was cached. 
* `last_access_time`:(string) The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `md5e_tag`:(string) The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly. 
* `md5sum`:(string) The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `mdfid`:(string) The mdfid of the image provided by cisco.com. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The endpoint model for which this firmware image is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the file. It is populated as part of the image import operation. 
* `platform_type`:(string) The platform type of the image. 
* `recommended_build`:(string) The build which is recommended by Cisco. 
* `release_date`:(string) The date on which the file was released or distributed by its vendor. 
* `release_notes_url`:(string) The url for the release notes of this image. 
* `sha512sum`:(string) The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `software_advisory_url`:(string) The software advisory, if any, provided by the vendor for this file. 
* `software_type_id`:(string) The software type id provided by cisco.com. 
* `vendor`:(string) The vendor or publisher of this file. 
* `nr_version`:(string) Vendor provided version for the file. 
 
