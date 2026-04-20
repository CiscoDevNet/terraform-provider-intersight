---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_application"
description: |-
        Application object has been deprecated. The Application object within the Cisco APIC framework represents the logical constructs that are associated with network policies and services.
        #### Purpose
        An Application in APIC serves as a representation of network services or software, tying together the necessary configurations and policies to ensure seamless operation within the network environment.
        #### Key Concepts
        - **Logical Representation:** Encapsulates the configurations and policies needed for application management in the network.
        - **Policy Integration:** Ensures that network policies are aligned with application requirements for optimal performance.
        - **Service Management:** Supports the operational aspects of applications, including deployment and maintenance.

---

# Data Source: intersight_apic_application
Application object has been deprecated. The Application object within the Cisco APIC framework represents the logical constructs that are associated with network policies and services.
#### Purpose
An Application in APIC serves as a representation of network services or software, tying together the necessary configurations and policies to ensure seamless operation within the network environment.
#### Key Concepts
- **Logical Representation:** Encapsulates the configurations and policies needed for application management in the network.
- **Policy Integration:** Ensures that network policies are aligned with application requirements for optimal performance.
- **Service Management:** Supports the operational aspects of applications, including deployment and maintenance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_application.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Application name of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
