---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_compliance"
description: |-
        A compliance instance associated with a cluster reported by a Prism Central.

---

# Data Source: intersight_hci_compliance
A compliance instance associated with a cluster reported by a Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_compliance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_ext_id`:(string) The identifier of a license, usually in LIC-xxxx format. 
* `compliance_state`:(string) The compliance state of the cluster. The compliance state is determined based on thenonComplianceCount and the license enforcement policy.* `NotEnforced` - The license compliance state for a Nutanix cluster is neither in=compliance nor not in-compliance.Typically, when a license is not installed on a cluster, or the license enformancement is explicitlyturned off, the cluster is in this state.* `InCompliance` - The license compliance state for a Nutanix cluster is in compliance.* `NotInCompliance` - The license compliance state for a Nutanix cluster is not in compliance. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `non_compliance_count`:(int) Total number of services that are not in-compliance. A count of 0 does not necessarilymean in compliance. Depending on 'complianceState', the cluster could be either bein compliance or the license check is not enforced.A synthesized property for the services property for ease of querying. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) The type of the cluster which could be NUTANIX, VMWARE, or NON_NUTANIX. 
 
