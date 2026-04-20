---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_external_routed_layer_three_domain"
description: |-
        ExternalRoutedLayerThreeDomain object has been deprecated. The ExternalRoutedLayerThreeDomain object in Cisco APIC focuses on managing external Layer 3 domains, facilitating communication outside the local network.
        #### Purpose
        External Routed Layer Three Domains provide the framework for handling traffic destined for external networks, ensuring efficient routing and policy application.
        #### Key Concepts
        - **Communication:** Supports the management of traffic and routing policies for external network domains.
        - **Control:** Provides mechanisms for configuring and managing external route policies.
        - **Policy Integration:** Ensures that external routing policies are aligned with overall network policies.

---

# Data Source: intersight_apic_external_routed_layer_three_domain
ExternalRoutedLayerThreeDomain object has been deprecated. The ExternalRoutedLayerThreeDomain object in Cisco APIC focuses on managing external Layer 3 domains, facilitating communication outside the local network.
#### Purpose
External Routed Layer Three Domains provide the framework for handling traffic destined for external networks, ensuring efficient routing and policy application.
#### Key Concepts
- **Communication:** Supports the management of traffic and routing policies for external network domains.
- **Control:** Provides mechanisms for configuring and managing external route policies.
- **Policy Integration:** Ensures that external routing policies are aligned with overall network policies.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_external_routed_layer_three_domain.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Object name in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
