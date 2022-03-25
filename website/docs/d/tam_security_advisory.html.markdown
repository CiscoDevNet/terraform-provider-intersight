---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_security_advisory"
description: |-
        Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.

---

# Data Source: intersight_tam_security_advisory
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_security_advisory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `advisory_id`:(string) Cisco generated identifier for the published security advisory. 
* `base_score`:(float) CVSS version 3 base score for the security Advisory. 
* `create_time`:(string) The time when this managed object was created. 
* `date_published`:(string) Date when the security advisory was first published by Cisco. 
* `date_updated`:(string) Date when the security advisory was last updated by Cisco. 
* `description`:(string) Brief description of the advisory details. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `environmental_score`:(float) CVSS version 3 environmental score for the security Advisory. 
* `external_url`:(string) A link to an external URL describing security Advisory in more details. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user defined name for the Intersight Advisory. 
* `recommendation`:(string) Recommended action to resolve the security advisory. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the advisory.* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation. 
* `status`:(string) Cisco assigned status of the published advisory based on whether the investigation is complete or on-going.* `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available.* `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability. 
* `temporal_score`:(float) CVSS version 3 temporal score for the security Advisory. 
* `nr_version`:(string) Cisco assigned advisory version after latest revision. 
* `workaround`:(string) Workarounds available for the advisory. 
 
