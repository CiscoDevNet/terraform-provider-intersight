---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_inc_customer_op"
description: |-
        The IncCustomerOp object is central to managing operations related to the Intersight Nexus Cloud (INC) licensing, offering functionalities for trial management and administrative control.
        #### Purpose
        IncCustomerOp facilitates management of Nexus Cloud licenses, providing tools for trial enablement and tier configuration. It ensures customers can access and manage cloud services efficiently.
        #### Key Concepts
        - **Nexus Cloud Licensing:** Centralizes management of cloud service licenses, allowing customers to optimize their cloud configurations.
        - **Trial Management:** Supports enabling and terminating trial periods, providing flexibility in evaluating cloud services.
        - **Tier Configuration:** Offers settings for active license tiers, ensuring customers can tailor cloud services to their operational needs.
        - **Administrative Activation:** Empowers users to activate license entitlements, ensuring access to cloud service capabilities.

---

# Data Source: intersight_license_inc_customer_op
The IncCustomerOp object is central to managing operations related to the Intersight Nexus Cloud (INC) licensing, offering functionalities for trial management and administrative control.
#### Purpose    
IncCustomerOp facilitates management of Nexus Cloud licenses, providing tools for trial enablement and tier configuration. It ensures customers can access and manage cloud services efficiently.
#### Key Concepts  
- **Nexus Cloud Licensing:** Centralizes management of cloud service licenses, allowing customers to optimize their cloud configurations.
- **Trial Management:** Supports enabling and terminating trial periods, providing flexibility in evaluating cloud services.
- **Tier Configuration:** Offers settings for active license tiers, ensuring customers can tailor cloud services to their operational needs.
- **Administrative Activation:** Empowers users to activate license entitlements, ensuring access to cloud service capabilities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_inc_customer_op.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `active_admin`:(bool) The Nexus Cloud license administrative state.Set this property to 'true' to activate the Nexus Cloud license entitlements. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_trial`:(bool) Enable trial for Nexus Cloud licensing. 
* `evaluation_period`:(int) The default Trial or Grace period the customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `terminate_trial`:(bool) Terminate trial mode for Nexus Cloud. 
 
