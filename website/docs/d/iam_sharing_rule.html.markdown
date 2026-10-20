---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_sharing_rule"
description: |-
        SharingRules define cross-resource sharing relationships where objects created under a source permission resource (such as an Organization) can be referenced by a destination permission resource. In effect, a SharingRule enables controlled “object re-use” across organizational boundaries by establishing a sanctioned share from source to destination.
        #### Purpose
        Enable governed sharing of resources across supported permission resources (for example, between organizations) so teams can reference and reuse objects created in a shared source context from another destination context.
        #### Key Concepts
        - **Reference-based sharing outcome**: Sharing allows objects created in the source context to be referenced by objects in the destination context, without duplicating those objects.
        - **Account-scoped governance**: The rule is tied to an `account`, ensuring sharing is managed within an account boundary.
        - **Controlled administration**: CREATE/UPDATE/DELETE require organization-management privileges; READ is available to org view/manage and common admin/support roles.

---

# Data Source: intersight_iam_sharing_rule
SharingRules define cross-resource sharing relationships where objects created under a source permission resource (such as an Organization) can be referenced by a destination permission resource. In effect, a SharingRule enables controlled “object re-use” across organizational boundaries by establishing a sanctioned share from source to destination.
#### Purpose
Enable governed sharing of resources across supported permission resources (for example, between organizations) so teams can reference and reuse objects created in a shared source context from another destination context.
#### Key Concepts
- **Reference-based sharing outcome**: Sharing allows objects created in the source context to be referenced by objects in the destination context, without duplicating those objects.
- **Account-scoped governance**: The rule is tied to an `account`, ensuring sharing is managed within an account boundary.
- **Controlled administration**: CREATE/UPDATE/DELETE require organization-management privileges; READ is available to org view/manage and common admin/support roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_sharing_rule.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
