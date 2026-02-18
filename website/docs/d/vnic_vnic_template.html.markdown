---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_vnic_template"
description: |-
        The VnicTemplate object encapsulates a reusable template for vNIC configuration. It standardizes the creation and management of multiple vNICs with consistent settings.
        #### Purpose
        A VnicTemplate simplifies and accelerates the deployment of standardized vNIC configurations across servers. It allows for the derivation of new vNICs from a common template and supports synchronization and property overrides to accommodate evolving requirements.
        #### Key Concepts
        - **Reusability:** Enables the rapid creation of vNICs with predefined settings.
        - **Template-Driven Management:** Supports synchronization of template changes to derived objects.
        - **Override Mechanism:** Provides flexibility to override specific properties as needed.

---

# Data Source: intersight_vnic_vnic_template
The VnicTemplate object encapsulates a reusable template for vNIC configuration. It standardizes the creation and management of multiple vNICs with consistent settings.
#### Purpose
A VnicTemplate simplifies and accelerates the deployment of standardized vNIC configurations across servers. It allows for the derivation of new vNICs from a common template and supports synchronization and property overrides to accommodate evolving requirements.
#### Key Concepts
- **Reusability:** Enables the rapid creation of vNICs with predefined settings.
- **Template-Driven Management:** Supports synchronization of template changes to derived objects.
- **Override Mechanism:** Provides flexibility to override specific properties as needed.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_vnic_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the vNIC template. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_override`:(bool) When enabled, the configuration of the derived instances may override the template configuration. 
* `failover_enabled`:(bool) Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster. 
* `lcp_usage_count`:(int) The count of the Lan Connectivity Policies using vNIC template. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the vNIC template. 
* `peer_vnic_name`:(string) Name of the peer vNIC which belongs to the peer FI. 
* `pin_group_name`:(string) Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) The fabric port to which the vNICs will be associated.* `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.* `A` - Fabric A of the FI cluster.* `B` - Fabric B of the FI cluster. 
* `update_status`:(string) The template sync status with all derived objects.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `usage_count`:(int) The number of objects derived from a Template MO instance. 
 
