---
subcategory: "certificatemanagement"
layout: "intersight"
page_title: "Intersight: intersight_certificatemanagement_policy"
description: |-
  Certificate Management policy models a reusable certificate and private key configuration that can be applied to multiple servers via profile association.
---

# Data Source: intersight_certificatemanagement_policy
Certificate Management policy models a reusable certificate and private key configuration that can be applied to multiple servers via profile association.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_certificatemanagement_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 