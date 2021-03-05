---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_tunnel"
description: |-
  Tunneled Virtual KVM access to the vKVM console of a server.
This must be specified while creating the vKVM session to gain tunneled access.
---

# Data Source: intersight_kvm_tunnel
Tunneled Virtual KVM access to the vKVM console of a server.
This must be specified while creating the vKVM session to gain tunneled access.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_tunnel.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `client_url`:(string) The multiplexer URL for the client to connect on. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 