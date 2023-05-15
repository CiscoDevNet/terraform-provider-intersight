---
subcategory: "partnerintegration"
layout: "intersight"
page_title: "Intersight: intersight_partnerintegration_file"
description: |-
        A partner integration artifact which will be files containing development code uploaded by our partners to get the build for dc/microservice.

---

# Data Source: intersight_partnerintegration_file
A partner integration artifact which will be files containing development code uploaded by our partners to get the build for dc/microservice.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_partnerintegration_file.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `download_count`:(int) The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `feature_source`:(string) The name of the feature to which the uploaded file belongs.* `System` - This indicates system initiated file uploads.* `OpenAPIImport` - This indicates an OpenAPI file upload.* `PartnerIntegrationImport` - This indicates a Partner-Integration Appliance user file uploads. 
* `file_path`:(string) Path of the file being uploaded. 
* `file_type`:(string) Type of the file being uploaded.* `None` - Invalid file type for partnerIntegration appliance.* `Model` - Model file of Generic Device.* `Etl` - ETL file of Generic Device.* `Ui` - UI file of Generic Device.* `DeviceConnector` - Generic Device Connector file. 
* `import_action`:(string) The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.* `None` - No action should be taken on the imported file.* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.* `PreCache` - Cache the file into the Intersight Appliance.* `Cancel` - The cancel import process for the file into the repository.* `Extract` - The action to extract the file in the external repository.* `Evict` - Evict the cached file from the Intersight Appliance. 
* `import_state`:(string) The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.* `ReadyForImport` - The image is ready to be imported into the repository.* `Importing` - The image is being imported into the repository.* `Imported` - The image has been extracted and imported into the repository.* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.* `Extracting` - Indicates that the image is being extracted into the repository.* `Extracted` - Indicates that the image has been extracted into the repository.* `Failed` - The image import from an external source to the repository has failed.* `MetaOnly` - The image is present in an external repository.* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.* `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image. 
* `imported_time`:(string) The time at which this image or file was imported/cached into the repositry. if the 'ImportState' is 'Imported', the time at which this image or file was imported. if the 'ImportState' is 'Cached', the time at which this image or file was cached. 
* `last_access_time`:(string) The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `md5e_tag`:(string) The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly. 
* `md5sum`:(string) The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the file. It is populated as part of the image import operation. 
* `release_date`:(string) The date on which the file was released or distributed by its vendor. 
* `sha512sum`:(string) The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `software_advisory_url`:(string) The software advisory, if any, provided by the vendor for this file. 
* `nr_version`:(string) Vendor provided version for the file. 
* `workspace_name`:(string) The partner integration workspace to use to upload the File. 
 
