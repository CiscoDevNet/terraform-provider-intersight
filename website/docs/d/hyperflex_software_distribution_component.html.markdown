---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_software_distribution_component"
description: |-
  A HyperFlex Software Distribution Component.
---

# Data Source: intersight_hyperflex_software_distribution_component
A HyperFlex Software Distribution Component.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_software_distribution_component.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bucket_name`:(string) The bucket name where the files are present, if source is external cloud store. 
* `component_id`:(string) The HyperFlex Software Distribution Component Identifier. 
* `component_name`:(string) The HyperFlex Software Distribution Component Name. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_path`:(string) File location on the cloud storage. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) The HyperFlex Software Distribution Component Version. 
 