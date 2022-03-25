---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_certificate"
description: |-
        DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.

---

# Data Source: intersight_appliance_device_certificate
DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_device_certificate.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `ca_certificate`:(string) The base64 encoded certificate in PEM format. 
* `ca_certificate_expiry_time`:(string) The expiry datetime of new ca certificate which need to be applied on device connector. 
* `certificate_renewal_expiry_time`:(string) The date time allocated till cert renewal will be executed. This time used here will be based on certrenewal plan. 
* `configuration_mo_id`:(string) The operation configuration MOId. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) End date of the certificate renewal. 
* `last_success_poll_time`:(string) The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date of the certificate renewal. 
* `status`:(string) The status of ca certificate renewal. 
 
