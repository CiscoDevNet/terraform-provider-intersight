
---
layout: "intersight"
page_title: "Intersight: intersight_kvm_policy"
sidebar_current: "docs-intersight-data-source-kvm-policy"
description: |-
Policy to configure KVM Launch settings.
---

# Data Source: intersight_kvm._policy
Policy to configure KVM Launch settings.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `enable_local_server_video`:(bool) If enabled, displays KVM session on any monitor attached to the server. 
* `enable_video_encryption`:(bool) If enabled, encrypts all video information sent through KVM. 
* `enabled`:(bool) State of the vKVM service on the endpoint. 
* `maximum_sessions`:(int) The maximum number of concurrent KVM sessions allowed. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `remote_port`:(int) The port used for KVM communication. 
