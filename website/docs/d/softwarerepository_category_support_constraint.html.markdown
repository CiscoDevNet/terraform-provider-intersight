---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_support_constraint"
description: |-
        The CategorySupportConstraint object defines constraints for models supported from certain minimum image versions. It plays a vital role in managing software compatibility and upgrade paths.
        #### Purpose
        CategorySupportConstraint ensures that only supported hardware models receive updates or installations, safeguarding against compatibility issues.
        #### Key Concepts
        - **Version Constraints:** Specifies minimum image versions required for model support, preventing incompatibility.
        - **Management:** Helps manage software installation and upgrade processes according to defined constraints.
        - **Compatibility Assurance:** Protects systems by ensuring only compatible models receive updates.
        - **Integration:** Works seamlessly with the Intersight catalog for hardware model verification.

---

# Data Source: intersight_softwarerepository_category_support_constraint
The CategorySupportConstraint object defines constraints for models supported from certain minimum image versions. It plays a vital role in managing software compatibility and upgrade paths.
#### Purpose
CategorySupportConstraint ensures that only supported hardware models receive updates or installations, safeguarding against compatibility issues.
#### Key Concepts
- **Version Constraints:** Specifies minimum image versions required for model support, preventing incompatibility.
- **Management:** Helps manage software installation and upgrade processes according to defined constraints.
- **Compatibility Assurance:** Protects systems by ensuring only compatible models receive updates.
- **Integration:** Works seamlessly with the Intersight catalog for hardware model verification.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_category_support_constraint.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `constraint_id`:(string) Identifier for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `min_version`:(string) Minimum image version from where the models can be supported. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `parse_from_image_name`:(bool) Fields which tells if the constraint is based on image name parsing. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
