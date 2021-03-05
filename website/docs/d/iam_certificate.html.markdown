---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate"
description: |-
  Holds a certificate, signed by a CAcert.
---

# Data Source: intersight_iam_certificate
Holds a certificate, signed by a CAcert.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_certificate.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `status`:(string) Status of the certificate.* `PendingValidation` - The certificate has not been validated.* `Valid` - The certificate is valid.* `Invalid` - Ther certificate is invalid. 
 