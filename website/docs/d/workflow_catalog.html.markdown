---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_catalog"
description: |-
        The Catalog object serves as a container for organizing definitions related to orchestration, including both workflow, task definitions and more.
        #### Purpose
        It provides a structured way to manage and isolate orchestration assets. Each user account has a local catalog for private definitions, while Cisco provides shared catalogs containing validated workflows and tasks for use within the account context.
        #### Key Concepts
        - **Asset Organization:** Groups workflow and task definitions for better discovery and management.
        - **Access Context:** Supports both private account-level catalogs and shared, Cisco-managed catalogs.
        - **Namespace Management:** Ensures unique identification of orchestration assets within the organizational context.

---

# Data Source: intersight_workflow_catalog
The Catalog object serves as a container for organizing definitions related to orchestration, including both workflow, task definitions and more.
#### Purpose
It provides a structured way to manage and isolate orchestration assets. Each user account has a local catalog for private definitions, while Cisco provides shared catalogs containing validated workflows and tasks for use within the account context.
#### Key Concepts
- **Asset Organization:** Groups workflow and task definitions for better discovery and management.
- **Access Context:** Supports both private account-level catalogs and shared, Cisco-managed catalogs.
- **Namespace Management:** Ensures unique identification of orchestration assets within the organizational context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_catalog.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the catalog. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
