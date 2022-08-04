---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_server_descriptor"
description: |-
        Descriptor that uniquely identifies an IMM server.

---

# Data Source: intersight_capability_server_descriptor
Descriptor that uniquely identifies an IMM server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_server_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_ncsi_enabled`:(bool) Indicates whether the CIMC to VIC side-band interface is enabled on the server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_form_factor`:(string) The form factor (blade/rack/etc) of the server.* `unknown` - The form factor of the server is unknown.* `blade` - Blade server form factor.* `rack` - Rack unit server form factor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
