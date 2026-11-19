---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_license"
description: |-
        Licenses represent a HyperFlex cluster’s feature or license identifier and current entitlement state, providing read-only visibility into the applied license type, whether Smart Licensing or license reservation is enabled, and whether the cluster complies with its entitlements, including evaluation and out-of-compliance timing details. Administrators and monitoring tools can use this information to query the cluster’s license type, registration and authorization details, and compliance state, identify entitlement issues, and take corrective action outside the API as needed. The licenseType field identifies the applied license, complianceState reports entitlement compliance, getOutOfComplianceStartAt records when the cluster became or would become out of compliance for audit and escalation, inEvaluation indicates whether it is in an evaluation period, and smartLicensingEnabled and plrEnabled indicate whether Smart Licensing and license reservation (PLR) are enabled. The licenseRegistration and licenseAuthorization fields provide structured registration and authorization details. Each License is linked to a specific non-null, read-only registeredDevice; deleting that device registration cascades to the License object. Indexes on licenseType, complianceState, and cluster support efficient filtering and reporting at scale.

---

# Data Source: intersight_hyperflex_license
Licenses represent a HyperFlex cluster’s feature or license identifier and current entitlement state, providing read-only visibility into the applied license type, whether Smart Licensing or license reservation is enabled, and whether the cluster complies with its entitlements, including evaluation and out-of-compliance timing details. Administrators and monitoring tools can use this information to query the cluster’s license type, registration and authorization details, and compliance state, identify entitlement issues, and take corrective action outside the API as needed. The licenseType field identifies the applied license, complianceState reports entitlement compliance, getOutOfComplianceStartAt records when the cluster became or would become out of compliance for audit and escalation, inEvaluation indicates whether it is in an evaluation period, and smartLicensingEnabled and plrEnabled indicate whether Smart Licensing and license reservation (PLR) are enabled. The licenseRegistration and licenseAuthorization fields provide structured registration and authorization details. Each License is linked to a specific non-null, read-only registeredDevice; deleting that device registration cascades to the License object. Indexes on licenseType, complianceState, and cluster support efficient filtering and reporting at scale.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_license.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `compliance_state`:(string) Is the cluster complaint with the license entitlements? 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `get_out_of_compliance_start_at`:(string) Out of compliance date of the cluster 
* `in_evaluation`:(bool) Is the cluster in evalution period? 
* `license_type`:(string) The type of license applied on the cluster 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `plr_enabled`:(bool) Is reservation enabled for the cluster? 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `smart_licensing_enabled`:(bool) Is Smart Licensing Enabled for this cluster? 
 
