---
subcategory: "metrics"
layout: "intersight"
page_title: "Intersight: intersight_metrics_resource_configuration"
description: |-
        ResourceConfigurations define metric collection settings for a specific managed resource within an account (for example, a server). They enable or disable metrics at the individual resource level and are always tied to the resource they control.
        #### Purpose
        Allow per-resource enablement of metric collection so administrators can selectively collect metrics from specific endpoints while respecting account limits and policies.
        #### Key Concepts
        - **Resource-scoped toggle:** The `enabled` flag governs whether metrics are collected for that particular resource.
        - **Inheritance from the resource:** Permissions inherit from the referenced resource to ensure only authorized users can change settings.
        - **Domain context:** Tracks the owning domain/target for the resource to support grouping and operational context.
        - **Identity by relationship:** Each ResourceConfiguration is uniquely identified by the resource it applies to.

---

# Data Source: intersight_metrics_resource_configuration
ResourceConfigurations define metric collection settings for a specific managed resource within an account (for example, a server). They enable or disable metrics at the individual resource level and are always tied to the resource they control.
#### Purpose
Allow per-resource enablement of metric collection so administrators can selectively collect metrics from specific endpoints while respecting account limits and policies.
#### Key Concepts
- **Resource-scoped toggle:** The `enabled` flag governs whether metrics are collected for that particular resource.
- **Inheritance from the resource:** Permissions inherit from the referenced resource to ensure only authorized users can change settings.
- **Domain context:** Tracks the owning domain/target for the resource to support grouping and operational context.
- **Identity by relationship:** Each ResourceConfiguration is uniquely identified by the resource it applies to.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_metrics_resource_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Metric collection is enabled for this resource, when enabled all available metrics are collected from the resource into Intersight. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
