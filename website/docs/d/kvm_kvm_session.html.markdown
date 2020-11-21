
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `kvm_mux_url`:(string) The URL of the KVM MUX to connect to. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
