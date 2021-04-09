---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nib_metadata"
description: |-
  Contains the latest metadata available for download from server.
---

# Data Source: intersight_niaapi_nib_metadata
Contains the latest metadata available for download from server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_nib_metadata.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `date`:(string) The date when the package was generated. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `metadata_chksum`:(string) Chksum used to check the integrity of the metadata file downloaded. 
* `metadata_filename`:(string) The filename of the metadata package. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(int) The version number of the metadata package. 
 