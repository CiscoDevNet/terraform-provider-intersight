
---
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege"
sidebar_current: "docs-intersight-data-source-iam-privilege"
description: |-
Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.
---

# Data Source: intersight_iam._privilege
Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `hostname_prefix`:(string) The hostname prefix of the resource corresponding to this privilege. For example \\'sentry\\' in https://sentry.intersight.com . 
* `method`:(string) The API method on the rest resource corresponding to privilege. For example READ, CREATE, UPDATE etc. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the privilege reported by microservice. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `rest_path`:(string) The REST API path of the resource corresponding to this privilege. For example /v1/iam/Accounts, /v1/iam/Sessions. 
* `url_prefix`:(string) The URL path prefix of the resource corresponding to this privilege. For example /devops/kibana, /devops/grafana etc. 
