---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_template"
description: |-
        Collection of information of templates.

---

# Data Source: intersight_dnac_template
Collection of information of templates.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `composite`:(string) Value indicates the template is composite. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `project_id`:(string) Identity of the project template. 
* `project_name`:(string) Name of the project template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `template_id`:(string) Unique identity of the template. 
* `template_name`:(string) Name assigned to the template. 
 
