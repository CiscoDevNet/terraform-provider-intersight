
---
layout: "intersight"
page_title: "Intersight: intersight_ipmioverlan_policy"
sidebar_current: "docs-intersight-data-source-ipmioverlan-policy"
description: |-
Intelligent Platform Management Interface Over LAN Policy.
---

# Data Source: intersight_ipmioverlan._policy
Intelligent Platform Management Interface Over LAN Policy.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of the IPMI Over LAN service on the endpoint. 
* `encryption_key`:(string) The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters. 
* `is_encryption_key_set`:(bool) Indicates whether the value of the 'encryptionKey' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `privilege`:(string) The highest privilege level that can be assigned to an IPMI session on a server.* `admin` - Privilege to perform all actions available through IPMI.* `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks.* `read-only` - Privilege to view information throught IPMI but restriction on making any changes. 
