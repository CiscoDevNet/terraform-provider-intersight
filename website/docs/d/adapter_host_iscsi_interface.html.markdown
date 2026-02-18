---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_host_iscsi_interface"
description: |-
        The adapter.HostIscsiInterface object represents a virtual iSCSI interface as it is presented to the host operating system. It is used to model and manage iSCSI-based storage connectivity for a server.
        #### Purpose
        The main function of the HostIscsiInterface object is to provide an inventory of iSCSI interfaces on a server's adapter. It captures the essential configuration and status details required for iSCSI network communication, such as its name, administrative state, and MAC address.
        #### Key Concepts
        - **iSCSI Interface Model:** Represents a logical iSCSI initiator interface on the adapter.
        - **Identity and State:** Tracks the interface's unique name, administrative state (adminState), and operational state (operState).
        - **Host Visibility:** The hostVisible property indicates whether the interface is exposed to the host operating system.
        - **Hierarchical Context:** Is a child of an adapter.Unit object, linking it directly to the physical adapter it resides on.

---

# Data Source: intersight_adapter_host_iscsi_interface
The adapter.HostIscsiInterface object represents a virtual iSCSI interface as it is presented to the host operating system. It is used to model and manage iSCSI-based storage connectivity for a server.
#### Purpose
The main function of the HostIscsiInterface object is to provide an inventory of iSCSI interfaces on a server's adapter. It captures the essential configuration and status details required for iSCSI network communication, such as its name, administrative state, and MAC address.
#### Key Concepts
- **iSCSI Interface Model:** Represents a logical iSCSI initiator interface on the adapter.
- **Identity and State:** Tracks the interface's unique name, administrative state (adminState), and operational state (operState).
- **Host Visibility:** The hostVisible property indicates whether the interface is exposed to the host operating system.
- **Hierarchical Context:** Is a child of an adapter.Unit object, linking it directly to the physical adapter it resides on.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_host_iscsi_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin Configured State of Host ISCSI Interface. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ep_dn`:(string) The Endpoint Config Dn of the Host ISCSI Interface. 
* `host_iscsi_interface_id`:(int) Identifier of the Host ISCSI Interface. 
* `host_visible`:(string) The visibility of the Host to the endpoint. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mac_address`:(string) MAC address of Host ISCSI Interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Host ISCSI Interface. 
* `oper_state`:(string) Operational State of Host ISCSI Interface. 
* `operability`:(string) Operability status of Host ISCSI Interface. 
* `peer_dn`:(string) PeerPort Dn of Host ISCSI Interface. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
