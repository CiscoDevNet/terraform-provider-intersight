
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host_group"
sidebar_current: "docs-intersight-data-source-storage-pure-host-group"
description: |-
A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.
---

# Data Source: intersight_storage._pure_host_group
A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Short description about the host group. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host group in storage array. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
