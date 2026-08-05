---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_site_ip_pool"
description: |-
        The SiteIpPool object provides a structured approach to managing IP pools associated with network sites, ensuring efficient and organized IP address allocation within the network infrastructure.
        #### Purpose
        A SiteIpPool object serves as a central repository for IP pool configurations tied to specific network sites, aiding in systematic IP address management and allocation.
        #### Key Concepts
        - **IP Pool Configuration:** Allows for detailed specification and management of IP pools, facilitating precise control over IP address distribution.
        - **Hierarchical Organization:** Maintains a structured hierarchy for site association, promoting clarity and order in IP management.
        - **Secure Access:** Utilizes privilege sets to control access and modification rights, ensuring secure interaction with IP pool data.

---

# Data Source: intersight_dnac_site_ip_pool
The SiteIpPool object provides a structured approach to managing IP pools associated with network sites, ensuring efficient and organized IP address allocation within the network infrastructure.
#### Purpose
A SiteIpPool object serves as a central repository for IP pool configurations tied to specific network sites, aiding in systematic IP address management and allocation.
#### Key Concepts
- **IP Pool Configuration:** Allows for detailed specification and management of IP pools, facilitating precise control over IP address distribution.
- **Hierarchical Organization:** Maintains a structured hierarchy for site association, promoting clarity and order in IP management.
- **Secure Access:** Utilizes privilege sets to control access and modification rights, ensuring secure interaction with IP pool data.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_site_ip_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ip_pool_cidr`:(string) Ip Pool Cidr in format e.g. 10.0.0.0/24. 
* `ip_pool_id`:(string) Site ip pool identification. 
* `ip_pool_name`:(string) Name for the site ip pool. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_id`:(string) Site to which ip pool is associated with. 
* `site_name_hierarchy`:(string) Hierarchy of the site names. 
 
