---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_advisory"
description: |-
        The Advisory object represents a security advisory reported by Nexus Dashboard for a specific fabric. It brings together the advisory context with affected nodes, current security status, recommendation guidance, and compensating-control information so API customers can understand how the advisory applies within the fabric.
        #### Purpose
        Advisory provides a fabric-scoped view of a reported security issue and its impact. API customers use it to assess which nodes are affected, understand the reported protection state, and review guidance and compensating controls associated with the advisory.
        #### Key Concepts
        - **Fabric Context**: Each advisory belongs to the Nexus Dashboard fabric for which the security issue is reported.
        - **Affected Nodes**: Node impact connects the advisory to the parts of the fabric that may require evaluation or protection.
        - **Security Status**: The reported status indicates whether the advisory is active, protected, partially protected, monitored, or cleared.
        - **Recommendations and Compensating Controls**: Recommendation guidance describes responses to consider, while compensating-control information conveys protection applied to affected nodes.

---

# Data Source: intersight_niatelemetry_advisory
The Advisory object represents a security advisory reported by Nexus Dashboard for a specific fabric. It brings together the advisory context with affected nodes, current security status, recommendation guidance, and compensating-control information so API customers can understand how the advisory applies within the fabric.
#### Purpose
Advisory provides a fabric-scoped view of a reported security issue and its impact. API customers use it to assess which nodes are affected, understand the reported protection state, and review guidance and compensating controls associated with the advisory.
#### Key Concepts
- **Fabric Context**: Each advisory belongs to the Nexus Dashboard fabric for which the security issue is reported.
- **Affected Nodes**: Node impact connects the advisory to the parts of the fabric that may require evaluation or protection.
- **Security Status**: The reported status indicates whether the advisory is active, protected, partially protected, monitored, or cleared.
- **Recommendations and Compensating Controls**: Recommendation guidance describes responses to consider, while compensating-control information conveys protection applied to affected nodes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_advisory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `advisory_id`:(string) Unique identifier assigned to the fabric advisory. 
* `advisory_score`:(int) Numeric score associated with the fabric advisory. 
* `advisory_string`:(string) Human-readable description of the fabric advisory. 
* `compensating_control`:(string) Compensating control status reported for the fabric advisory.* `available` - A compensating control is available for deployment.* `notAvailable` - No compensating control is available for the advisory.* `protected` - The advisory is protected by a deployed compensating control.* `monitoring` - The compensating control is monitoring relevant activity.* `partiallyApplied` - The compensating control is applied to only some affected nodes.* `deploymentFailed` - Deployment of the compensating control did not complete successfully.* `disabled` - The compensating control is currently disabled.* `queued` - The compensating control is queued for deployment.* `deploying` - The compensating control deployment is currently in progress.* `outOfSync` - The deployed compensating control is not synchronized with its intended configuration.* `redeployRequired` - The compensating control must be redeployed to restore the intended protection. 
* `create_time`:(string) The time when this managed object was created. 
* `cvss`:(float) CVSS base score associated with the fabric advisory. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `entity_name`:(string) Entity name associated with the fabric advisory. 
* `fabric_name`:(string) Fabric name associated with the advisory. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `security_advisory_id`:(string) Security advisory identifier associated with the fabric advisory. 
* `security_status`:(string) Security status reported for the fabric advisory.* `active` - The advisory is active and requires evaluation.* `protected` - The advisory is protected by a compensating control.* `partiallyProtected` - The advisory is partially protected by a compensating control.* `monitoring` - The advisory is being monitored for relevant activity.* `cleared` - The advisory has been cleared and no longer requires action. 
* `severity`:(string) Severity level of the advisory.* `critical` - The advisory has a critical severity level.* `major` - The advisory has a major severity level.* `minor` - The advisory has a minor severity level.* `warning` - The advisory has a warning severity level. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
