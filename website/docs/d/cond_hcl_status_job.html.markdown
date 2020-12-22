---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_hcl_status_job"
description: |-
  An HCLStatusJob is used to batch mo inventory notifications and process the evaluation of HCLStatus. When we receive a notification for an inventory MO, we will create a HCLStatusJob and inserted into the DB if it doesn't already exist. Then based on a timer we process the jobs in the DB and clear them.
---

# Data Source: intersight_cond_hcl_status_job
An HCLStatusJob is used to batch mo inventory notifications and process the evaluation of HCLStatus. When we receive a notification for an inventory MO, we will create a HCLStatusJob and inserted into the DB if it doesn't already exist. Then based on a timer we process the jobs in the DB and clear them.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
