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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_advisory_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `advisory_id`:(string) Cisco generated identifier for the published security/field-notice/end-of-life advisory. 
* `create_time`:(string) The time when this managed object was created. 
* `date_published`:(string) Date when the security/field-notice/end-of-life advisory was first published by Cisco. 
* `date_updated`:(string) Date when the security/field-notice/end-of-life advisory was last updated by Cisco. 
* `description`:(string) Brief description of the advisory details. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `execute_on_pod`:(string) Orion pod on which this advisory should process.* `tier1` - Advisory processing will be taken care in first advisory driver of multinode cluster.* `tier2` - Advisory processing will be taken care in second advisory driver of multinode cluster. 
* `external_url`:(string) A link to an external URL describing security Advisory in more details. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user defined name for the Intersight Advisory. 
* `recommendation`:(string) Recommended action to resolve the security advisory. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the advisory.* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation. 
* `type`:(string) The type (field notice, security advisory, end-of-life milestone advisory etc.) of Intersight advisory.* `securityAdvisory` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).* `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).* `eolAdvisory` - Represents product End of Life (EOL) type (https://www.cisco.com/c/en/us/products/eos-eol-policy.html). 
* `nr_version`:(string) Cisco assigned advisory/field-notice/end-of-life version after latest revision. 
* `workaround`:(string) Workarounds available for the advisory. 
 
