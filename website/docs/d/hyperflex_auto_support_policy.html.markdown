---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_auto_support_policy"
description: |-
  A policy specifying the configuration to automatically generate support tickets with Cisco TAC.
---

# Data Source: intersight_hyperflex_auto_support_policy
A policy specifying the configuration to automatically generate support tickets with Cisco TAC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_auto_support_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(bool) Enable or disable Auto-Support. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `service_ticket_receipient`:(string) The recipient email address for support tickets. 
 