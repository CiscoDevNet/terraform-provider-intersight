---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host"
description: |-
        PureHosts represent initiator endpoints (typically servers) that connect to a PureStorage array and are granted access to storage resources.
        #### Purpose
        Model host identity and connectivity so volumes can be mapped/authorized to the correct compute consumers.
        #### Key Concepts
        - **Access endpoint:** A host is the principal to which storage access is granted.
        - **Volume mapping context:**  Hosts are used to define which volumes are presented to which initiators (often through related LUN/mapping objects).

---

# Data Source: intersight_storage_pure_host
PureHosts represent initiator endpoints (typically servers) that connect to a PureStorage array and are granted access to storage resources.
#### Purpose
Model host identity and connectivity so volumes can be mapped/authorized to the correct compute consumers.
#### Key Concepts
- **Access endpoint:** A host is the principal to which storage access is granted.
- **Volume mapping context:**  Hosts are used to define which volumes are presented to which initiators (often through related LUN/mapping objects).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_host.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the host. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_group_name`:(string) Name of host group where the host belongs to. Empty if host is not part of any HostGroup. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host in storage array. 
* `os_type`:(string) Operating system running on the host. 
* `realm_name`:(string) A realm is the core multi-tenancy component on a Pure Flash Array, providing a self-contained, virtual storage environment with dedicated policies and quotas for secure data isolation and predictable performance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
