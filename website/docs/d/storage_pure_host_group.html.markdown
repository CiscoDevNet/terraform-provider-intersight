---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host_group"
description: |-
  A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.
---

# Data Source: intersight_storage_pure_host_group
A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_host_group.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Short description about the host group. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host group in storage array. 
 