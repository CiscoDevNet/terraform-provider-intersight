---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_blueprint"
description: |-
        A blueprint detailed description.

---

# Data Source: intersight_workload_blueprint
A blueprint detailed description.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workload_blueprint.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_version`:(bool) The flag to indicate that this is the default version of the blueprint. 
* `description`:(string) The description for this blueprint which provides information on what are the pre-requisites to deploy the blueprint and what features are supported on the blueprint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_meta`:(bool) When set to false the blueprint is created for use within internal services. 
* `label`:(string) A user friendly short name to identify the blueprint. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this blueprint. You can have multiple versions of the blueprint with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), or an underscore (_). 
* `platform_type`:(string) The Intersight platforms supported by this blueprint.* `None` - Default value is none, platform is not mentioned.* `UnifiedEdge` - The blueprint is created for Unified Edge platform. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(int) The version of the blueprint to support multiple versions. 
 
