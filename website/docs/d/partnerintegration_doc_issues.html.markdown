---
subcategory: "partnerintegration"
layout: "intersight"
page_title: "Intersight: intersight_partnerintegration_doc_issues"
description: |-
        The DocIssues object captures documentation-related issues identified during the build operation of an inventory collection.
        #### Purpose
        This provides feedback on documentation generation, allowing developers to identify and resolve issues in the generated API documentation or SDKs.
        #### Key Concepts
        - **Issue Tracking:** Lists specific documentation errors or warnings encountered during the build.
        - **Quality Assurance:** Helps ensure that generated documentation is accurate and complete.

---

# Data Source: intersight_partnerintegration_doc_issues
The DocIssues object captures documentation-related issues identified during the build operation of an inventory collection.
#### Purpose
This provides feedback on documentation generation, allowing developers to identify and resolve issues in the generated API documentation or SDKs.
#### Key Concepts
- **Issue Tracking:** Lists specific documentation errors or warnings encountered during the build.
- **Quality Assurance:** Helps ensure that generated documentation is accurate and complete.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_partnerintegration_doc_issues.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
