
---
layout: "intersight"
page_title: "Intersight: intersight_kvm_kvm_session"
sidebar_current: "docs-intersight-data-source-kvm-kvm-session"
description: |-
API to launch the KVM Session.
---

# Data Source: intersight_kvm._kvm_session
API to launch the KVM Session.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `kvm_mux_url`:(string) The URL of the KVM MUX to connect to. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
