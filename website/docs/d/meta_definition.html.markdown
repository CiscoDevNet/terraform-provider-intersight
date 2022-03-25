---
subcategory: "meta"
layout: "intersight"
page_title: "Intersight: intersight_meta_definition"
description: |-
        The meta-data of managed objects and complex types.

---

# Data Source: intersight_meta_definition
The meta-data of managed objects and complex types.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_meta_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_concrete`:(bool) Boolean flag to specify whether the meta class is a concrete class or not. 
* `meta_type`:(string) Indicates whether the meta class is a complex type or managed object.* `ManagedObject` - The meta.Definition object describes a managed object.* `ComplexType` - The meta.Definition object describes a nested complex type within a managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The fully-qualified class name of the Managed Object or complex type. For example, \ compute:Blade\  where the Managed Object is \ Blade\  and the package is 'compute'. 
* `namespace`:(string) The namespace of the meta. 
* `parent_class`:(string) The fully-qualified name of the parent metaclass in the class inheritance hierarchy. 
* `permission_supported`:(bool) Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations. 
* `rbac_resource`:(bool) Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource=true. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set. 
* `rest_path`:(string) Restful URL path for the meta. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) The version of the service that defines the meta-data. 
 
