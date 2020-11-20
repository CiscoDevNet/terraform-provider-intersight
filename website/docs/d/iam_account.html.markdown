
---
layout: "intersight"
page_title: "Intersight: intersight_iam_account"
sidebar_current: "docs-intersight-data-source-iam-account"
description: |-
The Intersight Account used to access Intersight.
---

# Data Source: intersight_iam._account
The Intersight Account used to access Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Intersight account. By default, name is same as the MoID of the account. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `status`:(string) Status of the account. To activate the Intersight account, claim a device to the account. 
