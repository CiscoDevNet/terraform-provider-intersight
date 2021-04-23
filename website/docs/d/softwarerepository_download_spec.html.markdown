---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_download_spec"
description: |-
  The URL, certificate and other associated information required to download an image listed in an Intersight catalog.
---

# Data Source: intersight_softwarerepository_download_spec
The URL, certificate and other associated information required to download an image listed in an Intersight catalog.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_download_spec.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auth_token`:(string) The OAuth2 token that will be used during image download by the endpoint to authenticate with file server. 
* `certificate`:(string) The certificate of file server that will be used by the endpoint to validate the server before starting the file download. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filename`:(string) The name of the firmware image. 
* `md5sum`:(string) MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image, post download. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) The size (in bytes) of the firmware image. 
* `url`:(string) The URL of this file in file server. The endpoint uses this URL to download the file from the file server. 
 