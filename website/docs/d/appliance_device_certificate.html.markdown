
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_certificate"
sidebar_current: "docs-intersight-data-source-appliance-device-certificate"
description: |-
DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.
---

# Data Source: intersight_appliance._device_certificate
DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `ca_certificate`:(string) The base64 encoded certificate in PEM format. 
* `ca_certificate_expiry_time`:(string) The expiry datetime of new ca certificate which need to be applied on device connector. 
* `certificate_renewal_expiry_time`:(string) The date time allocated till cert renewal will be executed. This time used here will be based on certrenewal plan. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `configuration_mo_id`:(string) The operation configuration MOId. 
* `end_time`:(string) End date of the certificate renewal. 
* `last_success_poll_time`:(string) The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `start_time`:(string) Start date of the certificate renewal. 
* `status`:(string) The status of ca certificate renewal. 
