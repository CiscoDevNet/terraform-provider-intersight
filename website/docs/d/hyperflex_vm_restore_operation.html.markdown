---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_vm_restore_operation"
description: |-
        The VmRestoreOperation object provides mechanisms for restoring Virtual Machines within the HyperFlex infrastructure. It is crucial for disaster recovery scenarios, enabling quick and efficient restoration of Virtual Machines to a predefined state.
        #### Purpose
        VmRestoreOperation is designed to manage the restoration processes, ensuring that Virtual Machines can be rapidly recovered in case of data loss or corruption. It supports both planned and unplanned restoration events, offering flexibility and control.
        #### Key Concepts
        - **Disaster Recovery:** Central to the object’s functionality, enabling organizations to restore Virtual Machines efficiently in response to failures.
        - **Automated Processes:** Supports automation of restoration tasks, reducing manual intervention and potential for errors.
        - **Flexibility:** Offers options for restoring Virtual Machines to different states or configurations, catering to varied recovery needs.
        - **Access Control:** Ensures secure execution of restore operations, with permissions designed to limit access to authorized users only.

---

# Data Source: intersight_hyperflex_vm_restore_operation
The VmRestoreOperation object provides mechanisms for restoring Virtual Machines within the HyperFlex infrastructure. It is crucial for disaster recovery scenarios, enabling quick and efficient restoration of Virtual Machines to a predefined state.
#### Purpose
VmRestoreOperation is designed to manage the restoration processes, ensuring that Virtual Machines can be rapidly recovered in case of data loss or corruption. It supports both planned and unplanned restoration events, offering flexibility and control.
#### Key Concepts
- **Disaster Recovery:** Central to the object’s functionality, enabling organizations to restore Virtual Machines efficiently in response to failures.
- **Automated Processes:** Supports automation of restoration tasks, reducing manual intervention and potential for errors.
- **Flexibility:** Offers options for restoring Virtual Machines to different states or configurations, catering to varied recovery needs.
- **Access Control:** Ensures secure execution of restore operations, with permissions designed to limit access to authorized users only.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_vm_restore_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `new_name`:(string) New name for the Virtual Machine after recovery. 
* `power_on`:(bool) Power on the Virtual Machine after recovery. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(int) Start time for the replication. 
 
