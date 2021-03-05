---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_policy"
description: |-
  Policy to configure KVM Launch settings.
---

# Data Source: intersight_kvm_policy
Policy to configure KVM Launch settings.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `enable_local_server_video`:(bool) If enabled, displays KVM session on any monitor attached to the server. 
* `enable_video_encryption`:(bool) If enabled, encrypts all video information sent through KVM. 
* `enabled`:(bool) State of the vKVM service on the endpoint. 
* `maximum_sessions`:(int) The maximum number of concurrent KVM sessions allowed. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `remote_port`:(int) The port used for KVM communication. 
 