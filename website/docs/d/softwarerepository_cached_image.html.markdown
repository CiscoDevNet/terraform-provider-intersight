---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_cached_image"
description: |-
  The image cached in the customer's datacenter.
---

# Data Source: intersight_softwarerepository_cached_image
The image cached in the customer's datacenter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_cached_image.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in endpoint. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.* `None` - No action should be taken on the imported file.* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.* `PreCache` - Cache the file into the Intersight Appliance.* `Cancel` - The cancel import process for the file into the repository.* `Extract` - The action to extract the file in the external repository.* `Evict` - Evict the cached file from the Intersight Appliance. 
* `cache_state`:(string) The state  of this file in the endpoint The importState is updated during the cache operation and as part of the cache monitoring process.* `ReadyForImport` - The image is ready to be imported into the repository.* `Importing` - The image is being imported into the repository.* `Imported` - The image has been extracted and imported into the repository.* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.* `Extracting` - Indicates that the image is being extracted into the repository.* `Extracted` - Indicates that the image has been extracted into the repository.* `Failed` - The image import from an external source to the repository has failed.* `MetaOnly` - The image is present in an external repository.* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. 
* `cached_time`:(string) The time when the image or file was cached into the FI storage. 
* `download_error`:(string) Any error encountered. Set to empty when download is in progress or completed. 
* `download_progress`:(int) The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. 
* `download_retries`:(int) The number of retries the plugin attempted before succeeding or failing the download. 
* `md5sum`:(string) The MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `original_sha512sum`:(string) The actual sha512sum of the cached image. 
* `path`:(string) The absolute path of the imported file in the endpoint. 
* `used_count`:(int) The number of times this file has been used to copy or upgrade or install actions. Used by the cache monitoring process to determine the files to be evicted from the cache. 
 