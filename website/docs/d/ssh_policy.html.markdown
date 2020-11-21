
---
layout: "intersight"
page_title: "Intersight: intersight_ssh_policy"
sidebar_current: "docs-intersight-data-source-ssh-policy"
description: |-
Secure shell policy on the endpoint.
---

# Data Source: intersight_ssh._policy
Secure shell policy on the endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of SSH service on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `port`:(int) Port used for secure shell access. 
* `timeout`:(int) Number of seconds to wait before the system considers a SSH request to have timed out. 
