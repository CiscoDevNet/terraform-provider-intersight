---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_export_policy"
description: |-
  NetApp export policies enable client access control to volumes. Clients that match specific IP addresses and/or specific authentication types are granted access.
---

# Data Source: intersight_storage_net_app_export_policy
NetApp export policies enable client access control to volumes. Clients that match specific IP addresses and/or specific authentication types are granted access.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the NFS export in storage array. 
* `policy_id`:(int) ID for the Export Policy. 
* `uuid`:(string) The uuid of this NFS export. 
