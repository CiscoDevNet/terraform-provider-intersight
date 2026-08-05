---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_domain_manager"
description: |-
        The DomainManager object represents the Nutanix Prism Central instance managing multiple clusters.
        #### Purpose
        The DomainManager object reports basic information about the Nutanix Prism Central instance such as its name and size. It also is a parent object for all other entities managed by Prism Central, including clusters, nodes, VMs, GPUs, and licenses.
        #### Key Concepts
        - **Inventory Collection Limits:** The API limits property defines the maximum number of inventory objects that can be collected based on the current Prism Central resource configuration.
        - **Anchor point for resources in the Prism Central:** All HCI clusters registered with Prism Central, including the Prism Central cluster itself, along with their associated resources—such as nodes, GPUs, VMs, and licenses.

---

# Data Source: intersight_hci_domain_manager
The DomainManager object represents the Nutanix Prism Central instance managing multiple clusters.
#### Purpose 
The DomainManager object reports basic information about the Nutanix Prism Central instance such as its name and size. It also is a parent object for all other entities managed by Prism Central, including clusters, nodes, VMs, GPUs, and licenses.
#### Key Concepts
- **Inventory Collection Limits:** The API limits property defines the maximum number of inventory objects that can be collected based on the current Prism Central resource configuration.
- **Anchor point for resources in the Prism Central:** All HCI clusters registered with Prism Central, including the Prism Central cluster itself, along with their associated resources—such as nodes, GPUs, VMs, and licenses.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_domain_manager.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `api_limits_string`:(string) The string representation of the API limits as a string. It can be used by Alarm. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `lcm_connectivity_type`:(string) The LCM (Life Cycle Manager) connectivity type. Possible values: CONNECTED_SITE (has internet connectivity), DARKSITE_DIRECT_UPLOAD (no external connectivity with direct upload), DARKSITE_WEB_SERVER (no external connectivity with darksite webserver). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the domain manager. 
* `pc_ext_id`:(string) The unique identifier of the domain manager (Prism Central) instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(string) The size of the domain manager such as STARTER, SMALL, LARGE, EXTRALARGE.It determines the resources used by the domain manager. 
 
