---
subcategory: "changelog"
layout: "intersight"
page_title: "Intersight: intersight_changelog_item"
description: |-
        # Overview
        The Item object serves as a pivotal component within the API contract changelog system,
        representing individual changes as specified by version attributes.
        It acts as a concrete entity within the changelog package, enabling precise tracking and management of updates across semantic and date versions.
        ## Purpose
        The Item object encapsulates and manages changes within the API contract,
        providing a structured approach to version tracking and documentation.
        It ensures that changes are recorded with detailed versioning information, supporting clarity and consistency across the API lifecycle.
        ## Key Concepts
        - **Version Control** – Supports semantic versioning to track changes systematically, ensuring smooth transitions across different API states.
        - **Access Control** – Privilege sets dictate access and operations, ensuring only authorized users can read, create, update, or delete changelog items.
        - **Catalog Integration** – Closely linked with catalogs, providing a structured environment for storing and retrieving version-related information.
        - **Identity Management** – Uses identity properties to prevent duplication, ensuring updates reflect the latest modifications without creating redundant entries.

---

# Data Source: intersight_changelog_item
# Overview
The Item object serves as a pivotal component within the API contract changelog system,  
representing individual changes as specified by version attributes.  
It acts as a concrete entity within the changelog package, enabling precise tracking and management of updates across semantic and date versions.
## Purpose
The Item object encapsulates and manages changes within the API contract,  
providing a structured approach to version tracking and documentation.  
It ensures that changes are recorded with detailed versioning information, supporting clarity and consistency across the API lifecycle.
## Key Concepts
- **Version Control** – Supports semantic versioning to track changes systematically, ensuring smooth transitions across different API states.
- **Access Control** – Privilege sets dictate access and operations, ensuring only authorized users can read, create, update, or delete changelog items.
- **Catalog Integration** – Closely linked with catalogs, providing a structured environment for storing and retrieving version-related information.
- **Identity Management** – Uses identity properties to prevent duplication, ensuring updates reflect the latest modifications without creating redundant entries.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_changelog_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `date_version`:(string) Date version for the API contract changelog item in rfc3339 format, without fractional seconds. Note, multiple items can exist for a single DateVersion. Example: 2023-12-19T00:00:00Z . 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `entity`:(string) Operation Id of the endpoint for which the changelog item is generated. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `semantic_version`:(string) Semantic version for the API contract changelog item. Note, multiple items can exist for a single SemanticVersion. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `value`:(string) The value associated with the API contract changelog item. 
 
