---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_qos_policy"
description: |-
        The FcQosPolicy object specifies Quality of Service (QoS) for Fibre Channel traffic on virtual HBAs, enabling administrators to control bandwidth, priority, and class of service for storage communications.
        #### Purpose
        FcQosPolicy ensures storage traffic receives appropriate prioritization and resource allocation, optimizing performance and reliability for storage-intensive workloads.
        #### Key Concepts
        - **Storage QoS Control:** Sets limits and priorities for Fibre Channel data flows.
        - **Integration with vHBAs:** Applied to virtual HBAs for consistent QoS enforcement.

---

# Data Source: intersight_vnic_fc_qos_policy
The FcQosPolicy object specifies Quality of Service (QoS) for Fibre Channel traffic on virtual HBAs, enabling administrators to control bandwidth, priority, and class of service for storage communications.
#### Purpose
FcQosPolicy ensures storage traffic receives appropriate prioritization and resource allocation, optimizing performance and reliability for storage-intensive workloads.
#### Key Concepts
- **Storage QoS Control:** Sets limits and priorities for Fibre Channel data flows.
- **Integration with vHBAs:** Applied to virtual HBAs for consistent QoS enforcement.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_fc_qos_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `burst`:(int) The burst traffic, in bytes, allowed on the vHBA. 
* `cos`:(int) Class of Service to be associated to the traffic on the virtual interface. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `max_data_field_size`:(int) The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `priority`:(string) The priortity matching the System QoS specified in the fabric profile.* `Best Effort` - QoS Priority for Best-effort traffic.* `FC` - QoS Priority for FC traffic.* `Platinum` - QoS Priority for Platinum traffic.* `Gold` - QoS Priority for Gold traffic.* `Silver` - QoS Priority for Silver traffic.* `Bronze` - QoS Priority for Bronze traffic. 
* `rate_limit`:(int) The value in Mbps to use for limiting the data rate on the virtual interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
