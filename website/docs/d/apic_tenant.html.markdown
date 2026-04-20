---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_tenant"
description: |-
        Tenant object has been deprecated. The Tenant object is an integral component of the Cisco Application Policy Infrastructure Controller (APIC) framework, serving as a logical container for policies, configurations, and applications within the network environment.
        #### Purpose
        A Tenant in APIC represents an organizational unit or a customer environment, encapsulating the network policies, applications, and services specific to that unit. It helps in isolating and organizing various network elements, ensuring that policies are applied consistently within the tenant's scope.
        #### Key Concepts
        - **Isolation:** Ensures that network policies and configurations are tenant-specific, preventing cross-tenant interference.
        - **Organizational Structure:** Provides a structured way to manage network resources and policies aligned with organizational or customer needs.
        - **Policy Management:** Facilitates the application and management of network policies specific to the tenant's environment.

---

# Data Source: intersight_apic_tenant
Tenant object has been deprecated. The Tenant object is an integral component of the Cisco Application Policy Infrastructure Controller (APIC) framework, serving as a logical container for policies, configurations, and applications within the network environment.
#### Purpose
A Tenant in APIC represents an organizational unit or a customer environment, encapsulating the network policies, applications, and services specific to that unit. It helps in isolating and organizing various network elements, ensuring that policies are applied consistently within the tenant's scope.
#### Key Concepts
- **Isolation:** Ensures that network policies and configurations are tenant-specific, preventing cross-tenant interference.
- **Organizational Structure:** Provides a structured way to manage network resources and policies aligned with organizational or customer needs.
- **Policy Management:** Facilitates the application and management of network policies specific to the tenant's environment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_tenant.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Tenant description in Cisco Application Policy Infrastructure Controller (APIC). 
* `dn`:(string) Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Tenant name in Cisco Application Policy Infrastructure Controller (APIC). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
