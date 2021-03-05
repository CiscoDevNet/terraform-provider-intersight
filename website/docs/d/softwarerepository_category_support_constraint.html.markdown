---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_support_constraint"
description: |-
  Defines constraints for models which are supported from certain minimum image version.
---

# Data Source: intersight_softwarerepository_category_support_constraint
Defines constraints for models which are supported from certain minimum image version.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_category_support_constraint.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `constraint_id`:(string) Identifier for this managed object. 
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `min_version`:(string) Minimum image version from where the models can be supported. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `parse_from_image_name`:(bool) Fields which tells if the constraint is based on image name parsing. 
 