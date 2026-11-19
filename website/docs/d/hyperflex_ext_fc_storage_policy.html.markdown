---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ext_fc_storage_policy"
description: |-
        The ExtFcStoragePolicies object specifies connectivity settings for fabric-attached external Fibre Channel (FC) storage.
        #### Purpose
        It enables the integration of external FC storage arrays with the HyperFlex cluster, defining the VSANs and WWxN ranges required for connectivity.
        #### Key Concepts
        - **External Storage Integration:** Enables fabric-attached FC storage connectivity.
        - **VSAN Configuration:** Defines primary and secondary VSANs for storage traffic.
        - **Identity Management:** Manages WWxN prefix ranges for FC endpoints.

---

# Data Source: intersight_hyperflex_ext_fc_storage_policy
The ExtFcStoragePolicies object specifies connectivity settings for fabric-attached external Fibre Channel (FC) storage.
#### Purpose
It enables the integration of external FC storage arrays with the HyperFlex cluster, defining the VSANs and WWxN ranges required for connectivity.
#### Key Concepts
- **External Storage Integration:** Enables fabric-attached FC storage connectivity.
- **VSAN Configuration:** Defines primary and secondary VSANs for storage traffic.
- **Identity Management:** Manages WWxN prefix ranges for FC endpoints.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_ext_fc_storage_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(bool) Enables or disables external FC storage configuration. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
