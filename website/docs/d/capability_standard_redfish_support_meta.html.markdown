---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_standard_redfish_support_meta"
description: |-
        The StandardRedfishSupportMeta object provides internal metadata to verify if a platform supports the standard Redfish SimpleUpdate operation.
        #### Purpose
        It enables the system to determine if a server platform can utilize the Redfish standard for firmware updates, ensuring the correct upgrade protocol is selected.
        #### Key Concepts
        - **Standardization:** Identifies platforms that adhere to Redfish SimpleUpdate standards.
        - **Protocol Selection:** Guides the system in choosing the appropriate firmware upgrade mechanism.

---

# Data Source: intersight_capability_standard_redfish_support_meta
The StandardRedfishSupportMeta object provides internal metadata to verify if a platform supports the standard Redfish SimpleUpdate operation.
#### Purpose
It enables the system to determine if a server platform can utilize the Redfish standard for firmware updates, ensuring the correct upgrade protocol is selected.
#### Key Concepts
- **Standardization:** Identifies platforms that adhere to Redfish SimpleUpdate standards.
- **Protocol Selection:** Guides the system in choosing the appropriate firmware upgrade mechanism.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_standard_redfish_support_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Verbose description regarding this group of platform. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `series_id`:(string) Classification of a set of server platform type. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
