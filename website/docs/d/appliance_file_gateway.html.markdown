---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_file_gateway"
description: |-
        FileGateway is a non-persistent model for accessing files from the Intersight.
        Intersight Appliances queries the FileGateway API, supplying a filename and version,
        to get the signed URL from the Intersight. Intersight Appliance services uses the
        signed URLs to download files and store them in the local image cache. Intersight
        will not store any record of the file access in the initial revision. This model is
        a pure pass through proxy for the cloud storage service.

---

# Data Source: intersight_appliance_file_gateway
FileGateway is a non-persistent model for accessing files from the Intersight.
Intersight Appliances queries the FileGateway API, supplying a filename and version,
to get the signed URL from the Intersight. Intersight Appliance services uses the
signed URLs to download files and store them in the local image cache. Intersight
will not store any record of the file access in the initial revision. This model is
a pure pass through proxy for the cloud storage service.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_file_gateway.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bucket_name`:(string) Bucket name in the cloud storage service where the file is stored. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_size`:(int) Size of the file in bytes. FileSize maybe zero if the storage service does not report file size. 
* `file_time`:(string) File timestamp as reported by the cloud storage service. 
* `file_type`:(string) User defined file type supplied by the caller. 
* `filename`:(string) Full name of the file as it exists in the cloud storage service. If the file is in a subfolder, then the filename must contain the full path. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presigned_url`:(string) Pre-signed URL obtained from the cloud storage service. 
* `server_cert`:(string) SSL certificate of the cloud storage service. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `validity_period`:(int) Signed URL's validity period in minutes. 
* `nr_version`:(string) File version as reported by the cloud storage service. 
 
