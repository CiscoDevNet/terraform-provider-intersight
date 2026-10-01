---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_geo_location"
description: |-
        GeoLocations represent named geographic locations that can be associated with managed targets. A GeoLocation stores user-defined location identity and optional physical address and coordinate data, enabling location-aware inventory, filtering, and reporting across an account.
        #### Purpose
        Provide a reusable location object that administrators can create and manage to tag/associate managed resources with real-world sites, supporting operational organization, reporting, and location-based views.
        #### Key Concepts
        - **Location identity and metadata**: `name` provides a required, user-friendly identifier; `address` captures a structured street address; `coordinates` stores latitude/longitude.
        - **Account ownership and lifecycle**: The `account` relationship ties the location to an owning `iam.Account` and cascades on account deletion.
        - **Shared org resource**: `sharedorgresource: true` indicates the location can be shared/used across organizations within the account context (supporting cross-org consumption).
        - **Global visibility with governed management**: READ is allowed under global privilege sets, while CREATE/UPDATE/DELETE are restricted to roles with Geo Location management privileges.
        - **Stable identity**: The object is identified by `name`, making it easy to reference consistently.

---

# Data Source: intersight_asset_geo_location
GeoLocations represent named geographic locations that can be associated with managed targets. A GeoLocation stores user-defined location identity and optional physical address and coordinate data, enabling location-aware inventory, filtering, and reporting across an account.
#### Purpose
Provide a reusable location object that administrators can create and manage to tag/associate managed resources with real-world sites, supporting operational organization, reporting, and location-based views.
#### Key Concepts
- **Location identity and metadata**: `name` provides a required, user-friendly identifier; `address` captures a structured street address; `coordinates` stores latitude/longitude.
- **Account ownership and lifecycle**: The `account` relationship ties the location to an owning `iam.Account` and cascades on account deletion.
- **Shared org resource**: `sharedorgresource: true` indicates the location can be shared/used across organizations within the account context (supporting cross-org consumption).
- **Global visibility with governed management**: READ is allowed under global privilege sets, while CREATE/UPDATE/DELETE are restricted to roles with Geo Location management privileges.
- **Stable identity**: The object is identified by `name`, making it easy to reference consistently.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_geo_location.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user provided name for the location. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
