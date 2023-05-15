---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_cloud_target"
description: |-
        Cloud target is a collection of cloud provider configuration information for targets e.g., AWS_S3 or Azure_Cloud.

---

# Data Source: intersight_storage_net_app_cloud_target
Cloud target is a collection of cloud provider configuration information for targets e.g., AWS_S3 or Azure_Cloud.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_cloud_target.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_key`:(string) Access key ID for AWS_S3 and other S3 compatible provider types. 
* `account_moid`:(string) The Account ID for this managed object. 
* `authentication_type`:(string) Authentication used to access the target. 
* `azure_account`:(string) Azure cloud account name. 
* `cap_url`:(string) Specifies a full URL of the request to a CAP server for retrieving temporary credentials (access-key, secret-password, and session token) for accessing the object store. 
* `certificate_validation_enabled`:(bool) Is SSL/TLS certificate validation enabled? The default value is true. This can only be modified for SGWS, IBM_COS, and ONTAP_S3 provider types. 
* `container`:(string) Data bucket/container name. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ipspace`:(string) IPspace to use in order to reach the cloud target. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the cloud target. 
* `owner`:(string) Owner of the target. Allowed values are FabricPool or SnapMirror.* `FabricPool` - NetApp FabricPool Cloud Target owner.* `SnapMirror` - NetApp SnapMirror Cloud Target owner. 
* `port`:(int) Port number of the object store that ONTAP uses when establishing a connection. 
* `provider_type`:(string) Type of cloud provider, e.g., AWS_S3 or ONTAP_S3.* `ONTAP_S3` - Cloud target provider type ONTAP_S3.* `AliCloud` - Cloud target provider type AliCloud.* `AWS_S3` - Cloud target provider type AWS S3.* `Azure_Cloud` - Cloud target provider type Azure_Cloud.* `GoogleCloud` - Cloud target provider type GoogleCloud.* `IBM_COS` - Cloud target provider type IBM_COS.* `SGWS` - Cloud target provider type SGWS. 
* `server`:(string) Fully qualified domain name of the object store server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snap_mirror_use`:(string) Use of the cloud target by SnapMirror.* `data` - Data is stored in the SnapMirror target.* `metadata` - Metadata is stored in the SnapMirror target. 
* `ssl_enabled`:(bool) SSL/HTTPS enabled or not. 
* `used`:(int) The amount of cloud space used by all the aggregates attached to the target, in bytes. 
* `uuid`:(string) Uuid of the cloud target. 
 
