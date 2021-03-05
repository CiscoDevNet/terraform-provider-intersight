---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_trust_point"
description: |-
  To affirm the identity of trusted source.
Allows import of third-party CA certificates in X.509 (CER) format.
It can be a root CA or an trust chain that leads to a root CA.
---

# Data Source: intersight_iam_trust_point
To affirm the identity of trusted source.
Allows import of third-party CA certificates in X.509 (CER) format.
It can be a root CA or an trust chain that leads to a root CA.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_trust_point.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chain`:(string) The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 