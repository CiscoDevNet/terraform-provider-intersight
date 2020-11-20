
---
layout: "intersight"
page_title: "Intersight: intersight_iam_qualifier"
sidebar_current: "docs-intersight-data-source-iam-qualifier"
description: |-
The qualifier defines how a user qualifies to be part of a user group.
---

# Data Source: intersight_iam._qualifier
The qualifier defines how a user qualifies to be part of a user group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
