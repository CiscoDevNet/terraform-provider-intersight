---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_auto_rma_policy"
description: |-
        AutoRmaPolicies control whether RMA-related data collection is enabled for a given registered device. The policy provides a simple on/off switch that determines if the system should collect the diagnostic data needed to support automated or streamlined return-material-authorization (RMA) workflows.
        #### Purpose
        Allow controlled enablement of RMA data collection on a per-device basis, so administrators can turn collection on when needed (or keep it disabled) while maintaining a clear association to the device that the policy governs.
        #### Key Concepts
        - **Per-device applicability**: The policy identity is the `registeredDevice`, indicating there is effectively one policy context per device registration.
        - **Data collection toggle**: `enable` is the primary control; when `true`, RMA data collection is enabled.
        - **Device lifecycle coupling**: `onpeerdelete: cascade` ties the policy to the registered device—if the device registration is removed, the associated policy is removed as well.
        - **Governed administration**: Broad READ visibility is provided across common operator roles, while CREATE/UPDATE is restricted to `System Administrator`.
        - **Inventory/alarm context linkage**: The `registeredDevice` relationship description indicates the policy is associated with the device reporting inventory on which alarms may be raised, supporting operational traceability.

---

# Data Source: intersight_appliance_auto_rma_policy
AutoRmaPolicies control whether RMA-related data collection is enabled for a given registered device. The policy provides a simple on/off switch that determines if the system should collect the diagnostic data needed to support automated or streamlined return-material-authorization (RMA) workflows.
#### Purpose
Allow controlled enablement of RMA data collection on a per-device basis, so administrators can turn collection on when needed (or keep it disabled) while maintaining a clear association to the device that the policy governs.
#### Key Concepts
- **Per-device applicability**: The policy identity is the `registeredDevice`, indicating there is effectively one policy context per device registration.
- **Data collection toggle**: `enable` is the primary control; when `true`, RMA data collection is enabled.
- **Device lifecycle coupling**: `onpeerdelete: cascade` ties the policy to the registered device—if the device registration is removed, the associated policy is removed as well.
- **Governed administration**: Broad READ visibility is provided across common operator roles, while CREATE/UPDATE is restricted to `System Administrator`.
- **Inventory/alarm context linkage**: The `registeredDevice` relationship description indicates the policy is associated with the device reporting inventory on which alarms may be raised, supporting operational traceability.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_auto_rma_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable`:(bool) Status of the data collection mode. If the value is 'true', then data collection is enabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
