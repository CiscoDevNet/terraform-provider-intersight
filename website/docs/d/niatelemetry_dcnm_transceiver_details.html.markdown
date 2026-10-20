---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_dcnm_transceiver_details"
description: |-
        The DcnmTransceiverDetails object manages inventory information for transceivers in DCNM environments, supporting hardware tracking and operational analysis.
        #### Purpose
        DcnmTransceiverDetails documents transceiver serial numbers, product IDs, vendor details, and part numbers, aiding hardware management and troubleshooting.
        #### Key Concepts
        - **Transceiver Inventory**: Captures deployment and attributes.
        - **Component Management**: Supports hardware lifecycle and maintenance.
        - **Operational Analytics**: Enhances inventory visibility.

---

# Data Source: intersight_niatelemetry_dcnm_transceiver_details
The DcnmTransceiverDetails object manages inventory information for transceivers in DCNM environments, supporting hardware tracking and operational analysis.
#### Purpose
DcnmTransceiverDetails documents transceiver serial numbers, product IDs, vendor details, and part numbers, aiding hardware management and troubleshooting. 
#### Key Concepts
- **Transceiver Inventory**: Captures deployment and attributes.
- **Component Management**: Supports hardware lifecycle and maintenance.
- **Operational Analytics**: Enhances inventory visibility.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_dcnm_transceiver_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `part_number`:(string) Part number of the transceiver in the fabric inventory. 
* `product_id`:(string) Product Id of the transceiver in the fabric inventory. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `serial_number`:(string) Serial number of the transceiver in the fabric inventory. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor_id`:(string) Vendor Id of the transceiver in the fabric inventory. 
 
