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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `ca_certificate`:(string) The base64 encoded certificate in PEM format. 
* `ca_certificate_expiry_time`:(string) The expiry datetime of new ca certificate which need to be applied on device connector. 
* `certificate_renewal_expiry_time`:(string) The date time allocated till cert renewal will be executed. This time used here will be based on certrenewal plan. 
* `configuration_mo_id`:(string) The operation configuration MOId. 
* `end_time`:(string) End date of the certificate renewal. 
* `last_success_poll_time`:(string) The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `start_time`:(string) Start date of the certificate renewal. 
* `status`:(string) The status of ca certificate renewal. 
