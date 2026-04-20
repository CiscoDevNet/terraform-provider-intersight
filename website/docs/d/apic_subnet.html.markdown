---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_subnet"
description: |-
        Subnet object has been deprecated. The Subnet object within the Cisco APIC framework is designed to manage the IP subnets associated with bridge domains.
        #### Purpose
        Subnets in APIC play a critical role in defining the IP address ranges within bridge domains, enabling efficient IP management and routing.
        #### Key Concepts
        - **IP Management:** Facilitates the organization and allocation of IP address ranges within the network infrastructure.
        - **Routing Control:** Supports the configuration of routing policies specific to subnet IP ranges.
        - **Network Visibility:** Enhances the visibility and management of IP networks within the bridge domain.

---

# Data Source: intersight_apic_subnet
Subnet object has been deprecated. The Subnet object within the Cisco APIC framework is designed to manage the IP subnets associated with bridge domains.
#### Purpose
Subnets in APIC play a critical role in defining the IP address ranges within bridge domains, enabling efficient IP management and routing.
#### Key Concepts
- **IP Management:** Facilitates the organization and allocation of IP address ranges within the network infrastructure.
- **Routing Control:** Supports the configuration of routing policies specific to subnet IP ranges.
- **Network Visibility:** Enhances the visibility and management of IP networks within the bridge domain.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_subnet.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ip`:(string) IP of an object in Cisco Application Policy Infrastructure Controller (APIC). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Object name in Cisco Application Policy Infrastructure Controller (APIC). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
