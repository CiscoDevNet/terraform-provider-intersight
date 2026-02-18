---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_iks_customer_op"
description: |-
        The IksCustomerOp object handles customer operations related to the Intersight Kubernetes Service (IKS) licensing, providing mechanisms for trial enablement and tier management.
        #### Purpose
        IksCustomerOp is designed to facilitate the management of Kubernetes service licenses. It enables trial periods, tier settings, and administrative controls, ensuring customers can leverage Kubernetes services effectively.
        #### Key Concepts
        - **Kubernetes Service Management:** Focuses on the licensing aspects of Kubernetes services, ensuring customers can optimize their configurations.
        - **Trial Enablement:** Allows users to activate trial modes, providing opportunities to explore Kubernetes service features.
        - **License Tier Configuration:** Supports active tier settings, allowing customization based on customer requirements and service levels.
        - **Administrative Entitlements:** Provides controls to activate licenses, ensuring customers maintain access to Kubernetes service capabilities.

---

# Data Source: intersight_license_iks_customer_op
The IksCustomerOp object handles customer operations related to the Intersight Kubernetes Service (IKS) licensing, providing mechanisms for trial enablement and tier management.
#### Purpose
IksCustomerOp is designed to facilitate the management of Kubernetes service licenses. It enables trial periods, tier settings, and administrative controls, ensuring customers can leverage Kubernetes services effectively.
#### Key Concepts
- **Kubernetes Service Management:** Focuses on the licensing aspects of Kubernetes services, ensuring customers can optimize their configurations.
- **Trial Enablement:** Allows users to activate trial modes, providing opportunities to explore Kubernetes service features.
- **License Tier Configuration:** Supports active tier settings, allowing customization based on customer requirements and service levels.
- **Administrative Entitlements:** Provides controls to activate licenses, ensuring customers maintain access to Kubernetes service capabilities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_iks_customer_op.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `active_admin`:(bool) The Intersight Kubernetes Service license administrative state.Set this property to 'true' to activate the IKS license entitlements. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_trial`:(bool) Enable trial for IKS licensing. 
* `evaluation_period`:(int) The default Trial or Grace period the customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
