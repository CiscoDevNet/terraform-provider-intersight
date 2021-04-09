---
subcategory: "techsupportmanagement"
layout: "intersight"
page_title: "Intersight: intersight_techsupportmanagement_download"
description: |-
  Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.
---

# Data Source: intersight_techsupportmanagement_download
Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_techsupportmanagement_download.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 