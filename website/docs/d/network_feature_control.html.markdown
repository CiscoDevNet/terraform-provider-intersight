---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_feature_control"
description: |-
        FeatureControls represent switch feature inventory entries that expose which features are available on a network element and their current administrative and operational states. Each FeatureControl record identifies a feature, how many instances exist, and provides a status message for additional detail.
        #### Purpose
        Provide read-only visibility into switch feature availability and state so administrators can understand which switch capabilities are enabled/active and diagnose feature readiness issues.
        #### Key Concepts
        - **Feature inventory:** Each record represents a specific feature available on the switch (identified by `name`).
        - **State reporting:** Captures both **admin state** (configured/desired enablement) and **operational state** (runtime status).
        - **Instance awareness:** `instance` indicates the number of instances of the feature present/active on the device.
        - **Operational diagnostics:** `statusMsg` provides detail/context for the current admin/operational state.
        - **Device association and scoping:** Tied to a specific `registeredDevice` and inherits permissions from the `networkElement` context for consistent RBAC.

---

# Data Source: intersight_network_feature_control
FeatureControls represent switch feature inventory entries that expose which features are available on a network element and their current administrative and operational states. Each FeatureControl record identifies a feature, how many instances exist, and provides a status message for additional detail.
#### Purpose
Provide read-only visibility into switch feature availability and state so administrators can understand which switch capabilities are enabled/active and diagnose feature readiness issues.
#### Key Concepts
- **Feature inventory:** Each record represents a specific feature available on the switch (identified by `name`).
- **State reporting:** Captures both **admin state** (configured/desired enablement) and **operational state** (runtime status).
- **Instance awareness:** `instance` indicates the number of instances of the feature present/active on the device.
- **Operational diagnostics:** `statusMsg` provides detail/context for the current admin/operational state.
- **Device association and scoping:** Tied to a specific `registeredDevice` and inherits permissions from the `networkElement` context for consistent RBAC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_feature_control.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) The admin state of the feature. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `instance`:(int) The number of instances of the feature. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name to identify the feature. 
* `operational_state`:(string) The operational state of the feature. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status_msg`:(string) The status message to capture admin state detailed information. 
 
