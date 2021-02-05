---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_adapter_policy"
description: |-
  Set of iSCSI properties that govern the host-side behavior of the adapter.
---

# Data Source: intersight_vnic_iscsi_adapter_policy
Set of iSCSI properties that govern the host-side behavior of the adapter.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `connection_time_out`:(int) The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable. 
* `description`:(string) Description of the policy. 
* `dhcp_timeout`:(int) The number of seconds to wait before the initiator assumes that the DHCP server is unavailable. 
* `lun_busy_retry_count`:(int) The number of times to retry the connection in case of a failure during iSCSI LUN discovery. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
