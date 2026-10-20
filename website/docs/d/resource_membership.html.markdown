---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_membership"
description: |-
        Memberships (ResourceMembership) represent the membership context of a resource—specifically which groups and organizations the resource is associated with, and how permissions and roles map to that resource through those organizations. This object is used to understand and (in controlled scenarios) refresh how access control applies to a resource.
        #### Purpose
        Provide a consolidated, queryable view of a resource’s organizational/group membership and the resulting permission-to-role mappings, enabling administrators and support personnel to audit and troubleshoot access to that resource.
        #### Key Concepts
        - **Resource-centric access view**: Focuses on a single resource and summarizes its associations to groups and organizations, and how those associations translate into effective roles.
        - **Permission/role mapping output**: `groupPermissionRoles` provides the mapping of groups and organizations to permissions and roles (read-only), making RBAC evaluation visible for the resource.
        - **Controlled lifecycle and updates**: READ is broadly available to org viewing roles; UPDATE is restricted to Account/System Administrators; CREATE is restricted to System Administrators.

---

# Data Source: intersight_resource_membership
Memberships (ResourceMembership) represent the membership context of a resource—specifically which groups and organizations the resource is associated with, and how permissions and roles map to that resource through those organizations. This object is used to understand and (in controlled scenarios) refresh how access control applies to a resource.
#### Purpose
Provide a consolidated, queryable view of a resource’s organizational/group membership and the resulting permission-to-role mappings, enabling administrators and support personnel to audit and troubleshoot access to that resource.
#### Key Concepts
- **Resource-centric access view**: Focuses on a single resource and summarizes its associations to groups and organizations, and how those associations translate into effective roles.
- **Permission/role mapping output**: `groupPermissionRoles` provides the mapping of groups and organizations to permissions and roles (read-only), making RBAC evaluation visible for the resource.
- **Controlled lifecycle and updates**: READ is broadly available to org viewing roles; UPDATE is restricted to Account/System Administrators; CREATE is restricted to System Administrators.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_membership.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reevaluate`:(bool) Set Reevaluate to true to reevaluate the membership of a resource. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_app`:(string) Name of the Service owning the resource. 
 
