---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_catalog"
description: |-
        Catalog is an organization-owned container for capability information associated with managed systems. It serves as the top-level grouping construct for organizing capability metadata into logical sections, and is intended to be maintained by DevOps users operating with the appropriate catalog administration privileges.
        #### Purpose
        Provide a centrally managed “capability repository” that can be queried by consumers (read-only roles) and maintained by authorized administrators, enabling consistent organization and lifecycle management of capability data.
        #### Key Concepts
        - **Container for capability data**: The Catalog is the root object that groups capability information for managed systems.
        - **Section-based organization**: The `sections` relationship holds the set of Section objects defined within the catalog, enabling hierarchical structuring of capabilities.
        - **Controlled administration**: Updates are restricted to the `CapabilityCatalog Administrator` role, reflecting DevOps-managed curation.
        - **Stable identity**: `name` is create-only and forms the object identity, ensuring the catalog can be referenced consistently over time.
        - **Cascade lifecycle for contents**: Deleting a catalog cascades deletion to its `sections`, keeping catalog content consistent with the container lifecycle.

---

# Data Source: intersight_capability_catalog
Catalog is an organization-owned container for capability information associated with managed systems. It serves as the top-level grouping construct for organizing capability metadata into logical sections, and is intended to be maintained by DevOps users operating with the appropriate catalog administration privileges.
#### Purpose
Provide a centrally managed “capability repository” that can be queried by consumers (read-only roles) and maintained by authorized administrators, enabling consistent organization and lifecycle management of capability data.
#### Key Concepts
- **Container for capability data**: The Catalog is the root object that groups capability information for managed systems.
- **Section-based organization**: The `sections` relationship holds the set of Section objects defined within the catalog, enabling hierarchical structuring of capabilities.
- **Controlled administration**: Updates are restricted to the `CapabilityCatalog Administrator` role, reflecting DevOps-managed curation.
- **Stable identity**: `name` is create-only and forms the object identity, ensuring the catalog can be referenced consistently over time.
- **Cascade lifecycle for contents**: Deleting a catalog cascades deletion to its `sections`, keeping catalog content consistent with the container lifecycle.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_catalog.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the catalog. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
