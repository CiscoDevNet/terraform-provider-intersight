
---
layout: "intersight"
page_title: "Intersight: intersight_tam_security_advisory"
sidebar_current: "docs-intersight-data-source-tam-security-advisory"
description: |-
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
---

# Data Source: intersight_tam._security_advisory
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advisory_id`:(string) Cisco generated identifier for the published security advisory. 
* `base_score`:(float) CVSS version 3 base score for the security Advisory. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `date_published`:(string) Date when the security advisory was first published by Cisco. 
* `date_updated`:(string) Date when the security advisory was last updated by Cisco. 
* `description`:(string) Brief description of the advisory details. 
* `environmental_score`:(float) CVSS version 3 environmental score for the security Advisory. 
* `external_url`:(string) A link to an external URL describing security Advisory in more details. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user defined name for the Intersight Advisory. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `recommendation`:(string) Recommended action to resolve the security advisory. 
* `state`:(string) Current state of the advisory.* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation. 
* `status`:(string) Cisco assigned status of the published advisory based on whether the investigation is complete or on-going.* `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available.* `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability. 
* `temporal_score`:(float) CVSS version 3 temporal score for the security Advisory. 
* `version`:(string) Cisco assigned advisory version after latest revision. 
* `workaround`:(string) Workarounds available for the advisory. 
