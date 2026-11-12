---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_manifest"
description: |-
        ### Overview
        The **Manifest** object records the versioned set of workload related objects required in the shared catalog.
        It represents the cumulative state of packaged blueprints, default policies, and templates as they are imported into the catalog.
        ### Purpose
        A Manifest provides durable progress tracking for importing packaged workload content.
        It enables the system to resume pending imports, process manifest versions in order, and avoid reimporting catalog objects that are already available.
        ### Key Concepts
        * **Ordered import**: Manifest versions are processed in order to populate the shared catalog with packaged workload content.
        * **Cumulative state**: Each version builds on previously imported content and retains catalog-object status across import cycles.
        * **Import tracking**: The manifest records whether its content is imported or pending, allowing interrupted imports to resume safely.

---

# Data Source: intersight_workload_manifest
### Overview
The **Manifest** object records the versioned set of workload related objects required in the shared catalog.
It represents the cumulative state of packaged blueprints, default policies, and templates as they are imported into the catalog.
### Purpose
A Manifest provides durable progress tracking for importing packaged workload content.
It enables the system to resume pending imports, process manifest versions in order, and avoid reimporting catalog objects that are already available.
### Key Concepts
* **Ordered import**: Manifest versions are processed in order to populate the shared catalog with packaged workload content.
* **Cumulative state**: Each version builds on previously imported content and retains catalog-object status across import cycles.
* **Import tracking**: The manifest records whether its content is imported or pending, allowing interrupted imports to resume safely.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workload_manifest.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `import_status`:(string) The import status for the current manifest. If all items in the manifest items list are imported, the current manifest will be marked as imported.* `ImportPending` - The item from the manifest file is pending an import into the shared catalog.* `Imported` - The item from the manifest file has been imported in the shared catalog. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) The version number for the manifest. 
 
