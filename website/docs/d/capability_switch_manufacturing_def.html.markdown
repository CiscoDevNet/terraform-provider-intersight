---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_manufacturing_def"
description: |-
        The SwitchManufacturingDef object captures manufacturing and descriptive metadata for a switch/fabric-interconnect platform.
        #### Purpose
        This standardizes product identity and descriptive fields (caption, product name, part number, etc.) so catalog-driven UI and workflows can consistently represent the platform.
        #### Key Concepts
        - **Product identity standardization:** Centralizes canonical names and part identifiers.
        - **Catalog-driven UX:** Enables consistent display and reporting across device families.
        - **Platform traceability:** Helps correlate discovered switch inventory to known product definitions.
        - **Separation of concerns:** Keeps identity/description separate from operational configuration state.

---

# Data Source: intersight_capability_switch_manufacturing_def
The SwitchManufacturingDef object captures manufacturing and descriptive metadata for a switch/fabric-interconnect platform.
#### Purpose
This standardizes product identity and descriptive fields (caption, product name, part number, etc.) so catalog-driven UI and workflows can consistently represent the platform.
#### Key Concepts
- **Product identity standardization:** Centralizes canonical names and part identifiers.
- **Catalog-driven UX:** Enables consistent display and reporting across device families.
- **Platform traceability:** Helps correlate discovered switch inventory to known product definitions.
- **Separation of concerns:** Keeps identity/description separate from operational configuration state.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_switch_manufacturing_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `caption`:(string) Caption for Switch/Fabric-Interconnect. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description for Switch/Fabric-Interconnect. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `part_number`:(string) Part Number for Switch/Fabric-Interconnect. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `UCS-FI-6536` - The standard 5th generation UCS Fabric Interconnect with 36 ports.* `UCSX-S9108-100G` - Cisco UCS Fabric Interconnect 9108 100G with 8 ports.* `UCS-FI-6664` - The standard 6th generation UCS Fabric Interconnect with 64 ports.* `UCS-FI-6652` - The standard 6th generation UCS Fabric Interconnect with 52 ports.* `UCSXE-ECMC-10G` - Cisco UCS XE ECMC 10G with 2 ports.* `UCSXE-ECMC-G1` - Cisco UCS XE ECMC G1 with 2 ports.* `unknown` - Unknown device type, usage is TBD. 
* `product_name`:(string) Product Name for Switch/Fabric-Interconnect. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
 
