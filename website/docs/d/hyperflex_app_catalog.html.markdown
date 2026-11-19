---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_app_catalog"
description: |-
        AppCatalog is a system-owned, singleton catalog used by the HyperFlex cluster configuration service to manage application settings and supported configuration data. It is created during service initialization and is maintained by the system (for example, via packaged catalog content and future external catalog sources), while being readable in account context for operational and configuration workflows.
        #### Purpose
        Serve as the authoritative catalog for HyperFlex configuration-service application settings and supported software/version data, enabling cluster profile workflows and UIs to reference a consistent set of supported HyperFlex Data Platform versions and software distribution entries.
        #### Key Concepts
        - **Singleton, system-managed catalog**: `owner: system` and restricted CREATE/UPDATE/DELETE (System Administrator only) indicate the platform controls lifecycle and content updates.
        - **Account-context consumption**: Broad READ access enables users and tooling to query supported settings without being able to mutate the catalog.
        - **Catalog versioning**: `version` tracks the catalog revision used by the HyperFlex configuration service.
        - **Supported software versions**: The catalog stores supported HyperFlex Data Platform versions and enables compatibility-driven selection.
        - **Supported distribution entries**: `softwareDistributions` lists supported HyperFlex software distribution entries used by the service for validation and workflows.
        - **Hierarchical catalog structure**: Version and software distribution records establish AppCatalog as the container for these supported-data records.

---

# Data Source: intersight_hyperflex_app_catalog
AppCatalog is a system-owned, singleton catalog used by the HyperFlex cluster configuration service to manage application settings and supported configuration data. It is created during service initialization and is maintained by the system (for example, via packaged catalog content and future external catalog sources), while being readable in account context for operational and configuration workflows.
#### Purpose
Serve as the authoritative catalog for HyperFlex configuration-service application settings and supported software/version data, enabling cluster profile workflows and UIs to reference a consistent set of supported HyperFlex Data Platform versions and software distribution entries.
#### Key Concepts
- **Singleton, system-managed catalog**: `owner: system` and restricted CREATE/UPDATE/DELETE (System Administrator only) indicate the platform controls lifecycle and content updates.
- **Account-context consumption**: Broad READ access enables users and tooling to query supported settings without being able to mutate the catalog.
- **Catalog versioning**: `version` tracks the catalog revision used by the HyperFlex configuration service.
- **Supported software versions**: The catalog stores supported HyperFlex Data Platform versions and enables compatibility-driven selection.
- **Supported distribution entries**: `softwareDistributions` lists supported HyperFlex software distribution entries used by the service for validation and workflows.
- **Hierarchical catalog structure**: Version and software distribution records establish AppCatalog as the container for these supported-data records.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_app_catalog.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) The catalog version used in HyperFlex cluster configuration service. 
 
