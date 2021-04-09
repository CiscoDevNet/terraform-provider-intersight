---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_io_card_descriptor"
description: |-
  Descriptor that uniquely identifies an IO card module.
---

# Data Source: intersight_capability_io_card_descriptor
Descriptor that uniquely identifies an IO card module.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_io_card_descriptor.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_hif_ports`:(int) Number of hif ports per blade for the iocard module. 
* `revision`:(string) Revision for the iocard module. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uif_connectivity`:(string) Connectivity information between UIF Uplink ports and IOM ports.* `inline` - UIF uplink ports and IOM ports are connected inline.* `cross-connected` - UIF uplink ports and IOM ports are cross-connected, a case in washington chassis. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 