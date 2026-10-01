---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_running_firmware"
description: |-
        RunningFirmwares represent the firmware currently running on a managed endpoint component. They provide an inventory view of “what version is actually active” for a broad range of hardware and management components (servers, storage controllers/disks, BIOS, PSUs, PCIe switches, graphics cards, and more), enabling firmware visibility and lifecycle correlation across the infrastructure.
        #### Purpose
        Expose the effective (running) firmware state of endpoint components so administrators can audit versions, validate compliance, and support firmware lifecycle operations and troubleshooting.
        #### Key Concepts
        - **Observed running state (not desired)**: Captures the firmware version that is currently active on a component, independent of any upgrade plan or policy intent.
        - **Cross-component applicability**: Inherits permissions from many component contexts (BIOS units, storage controllers/disks, management controller, PSUs, PCI switches, graphics cards, etc.), reflecting that firmware is tracked across diverse hardware types.
        - **Endpoint inventory integration**: Extends `inventory.Base`, aligning it with inventory/discovery pipelines and allowing it to be referenced by other inventory objects.
        - **Efficient correlation by component and ancestry**: Indexed by `Ancestors` and `Component`, supporting queries like “show me running firmware for all components under this server/chassis/device.”
        - **Governed updates**: UPDATE is available only to server/FI management roles, reflecting controlled workflows that may refresh or reconcile firmware state as part of management operations.

---

# Data Source: intersight_firmware_running_firmware
RunningFirmwares represent the firmware currently running on a managed endpoint component. They provide an inventory view of “what version is actually active” for a broad range of hardware and management components (servers, storage controllers/disks, BIOS, PSUs, PCIe switches, graphics cards, and more), enabling firmware visibility and lifecycle correlation across the infrastructure.
#### Purpose
Expose the effective (running) firmware state of endpoint components so administrators can audit versions, validate compliance, and support firmware lifecycle operations and troubleshooting.
#### Key Concepts
- **Observed running state (not desired)**: Captures the firmware version that is currently active on a component, independent of any upgrade plan or policy intent.
- **Cross-component applicability**: Inherits permissions from many component contexts (BIOS units, storage controllers/disks, management controller, PSUs, PCI switches, graphics cards, etc.), reflecting that firmware is tracked across diverse hardware types.
- **Endpoint inventory integration**: Extends `inventory.Base`, aligning it with inventory/discovery pipelines and allowing it to be referenced by other inventory objects.
- **Efficient correlation by component and ancestry**: Indexed by `Ancestors` and `Component`, supporting queries like “show me running firmware for all components under this server/chassis/device.”
- **Governed updates**: UPDATE is available only to server/FI management roles, reflecting controlled workflows that may refresh or reconcile firmware state as part of management operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_running_firmware.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `component`:(string) Kind of the firmware - boot-booloader/system/kernel. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `package_version`:(string) Bundle version which the firmware belongs to. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) The type of the firmware. 
* `nr_version`:(string) The version of the firmware. 
 
