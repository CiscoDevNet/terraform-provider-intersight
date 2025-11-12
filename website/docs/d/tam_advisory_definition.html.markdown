---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_definition"
description: |-
        ### Overview
        The AdvisoryDefinition object is a critical component in Intersight. It encompasses various types of advisories, including Field Notices, End of Life (EOL), End of Sale (EOS), and End of Support advisories, each with specific recommendations for addressing them.
        #### Purpose
        AdvisoryDefinition provides a comprehensive framework for defining potential advisories that may impact managed objects in the datacenter. It delivers insights and guidance for handling Field Notices, EOL, EOS, and similar advisories effectively.
        #### Key Concepts
        - **Diverse Advisory Types** - Supports Field Notices, EOL, EOS, and other advisory types, each with tailored recommendations.
        - **Actionable Guidance** - Provides recommendations and potential workarounds to mitigate identified issues.
        - **Role-Based Access** - Access to create, update, or delete advisories is controlled through user privilege sets.

---

# Data Source: intersight_tam_advisory_definition
### Overview
The AdvisoryDefinition object is a critical component in Intersight. It encompasses various types of advisories, including Field Notices, End of Life (EOL), End of Sale (EOS), and End of Support advisories, each with specific recommendations for addressing them.
#### Purpose
AdvisoryDefinition provides a comprehensive framework for defining potential advisories that may impact managed objects in the datacenter. It delivers insights and guidance for handling Field Notices, EOL, EOS, and similar advisories effectively.
#### Key Concepts
- **Diverse Advisory Types** - Supports Field Notices, EOL, EOS, and other advisory types, each with tailored recommendations.
- **Actionable Guidance** - Provides recommendations and potential workarounds to mitigate identified issues.
- **Role-Based Access** - Access to create, update, or delete advisories is controlled through user privilege sets.
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
* `execute_on_pod`:(string) Orion pod on which this advisory should process.* `tier1` - Advisory processing will be taken care by batch processing.* `tier2` - Advisory processing will be taken care by stream processing. 
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
 
