---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_apic_snmp_community_details"
description: |-
        The ApicSnmpCommunityDetails object tracks SNMP community configurations in APIC environments, supporting network monitoring and policy management.
        #### Purpose
        ApicSnmpCommunityDetails documents community names, distinguished names, and site associations, aiding operational integration.
        #### Key Concepts
        - **Community Configuration**: Captures SNMP community attributes.
        - **Monitoring Integration**: Supports network policy enforcement.
        - **Operational Analytics**: Enhances SNMP management.

---

# Data Source: intersight_niatelemetry_apic_snmp_community_details
The ApicSnmpCommunityDetails object tracks SNMP community configurations in APIC environments, supporting network monitoring and policy management.
#### Purpose
ApicSnmpCommunityDetails documents community names, distinguished names, and site associations, aiding operational integration.
#### Key Concepts
- **Community Configuration**: Captures SNMP community attributes.
- **Monitoring Integration**: Supports network policy enforcement.
- **Operational Analytics**: Enhances SNMP management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_apic_snmp_community_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Dn of the SNMP community in APIC. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of SNMP community in APIC. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
 
