---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_definition"
description: |-
  An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
---

# Data Source: intersight_tam_advisory_definition
An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advisory_id`:(string) Cisco generated identifier for the published security advisory. 
* `date_published`:(string) Date when the security advisory was first published by Cisco. 
* `date_updated`:(string) Date when the security advisory was last updated by Cisco. 
* `description`:(string) Brief description of the advisory details. 
* `external_url`:(string) A link to an external URL describing security Advisory in more details. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user defined name for the Intersight Advisory. 
* `recommendation`:(string) Recommended action to resolve the security advisory. 
* `state`:(string) Current state of the advisory.* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation. 
* `type`:(string) The type (field notice, security advisory etc.) of Intersight advisory.* `securityAdvisory` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).* `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). 
* `nr_version`:(string) Cisco assigned advisory version after latest revision. 
* `workaround`:(string) Workarounds available for the advisory. 
