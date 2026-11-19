---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_capability_info"
description: |-
        CapabilityInfos capture HyperFlex capability metadata—feature sets and feature limits—expressed as a set of constraints and an associated value. The constraints define when a capability applies (evaluated together with an AND semantics), and the resulting value indicates either whether a feature is supported for a given component/version combination or what limit/threshold should be used.
        #### Purpose
        Enable HyperFlex workflows (such as upgrades and compatibility checks) to make data-driven runtime decisions about whether an intended operation is supported and, when supported, what minimum versions or limit values should be enforced for the detected cluster/component set.
        #### Key Concepts
        - **Constraint-driven capability evaluation**: `capabilityConstraints` define the applicability rules; all constraints must match for the capability to be considered applicable.
        - **Capability value as outcome/limit**: `value` represents the effective value when constraints match—commonly minimum version requirements (e.g., `minUcsVersion`, `minHxdpVersion`) or feature limits for specific component sets.
        - **Runtime decision support**: The model is intended to support “go/no-go” and “choose correct limit” logic during operations like upgrades, validations, or feature enablement.
        - **Catalog-governed access**: Inherits permissions from `appCatalog`, indicating capabilities are managed in the context of an application catalog and follow that catalog’s access boundaries.
        - **Discoverability by name**: Indexed by `Name` (`byName`) to enable quick lookups of a capability/feature set by identifier.
        - **Controlled lifecycle**: READ is broadly available to HyperFlex operators; CREATE/UPDATE requires elevated/admin/profile-management privileges; DELETE is restricted to System Administrator.

---

# Data Source: intersight_hyperflex_capability_info
CapabilityInfos capture HyperFlex capability metadata—feature sets and feature limits—expressed as a set of constraints and an associated value. The constraints define when a capability applies (evaluated together with an AND semantics), and the resulting value indicates either whether a feature is supported for a given component/version combination or what limit/threshold should be used.
#### Purpose
Enable HyperFlex workflows (such as upgrades and compatibility checks) to make data-driven runtime decisions about whether an intended operation is supported and, when supported, what minimum versions or limit values should be enforced for the detected cluster/component set.
#### Key Concepts
- **Constraint-driven capability evaluation**: `capabilityConstraints` define the applicability rules; all constraints must match for the capability to be considered applicable.
- **Capability value as outcome/limit**: `value` represents the effective value when constraints match—commonly minimum version requirements (e.g., `minUcsVersion`, `minHxdpVersion`) or feature limits for specific component sets.
- **Runtime decision support**: The model is intended to support “go/no-go” and “choose correct limit” logic during operations like upgrades, validations, or feature enablement.
- **Catalog-governed access**: Inherits permissions from `appCatalog`, indicating capabilities are managed in the context of an application catalog and follow that catalog’s access boundaries.
- **Discoverability by name**: Indexed by `Name` (`byName`) to enable quick lookups of a capability/feature set by identifier.
- **Controlled lifecycle**: READ is broadly available to HyperFlex operators; CREATE/UPDATE requires elevated/admin/profile-management privileges; DELETE is restricted to System Administrator.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_capability_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the capability or feature set consisting of a collection of constraint rules and value. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `value`:(string) Capability Value which is valid only iff all specified constraints match. 
 
