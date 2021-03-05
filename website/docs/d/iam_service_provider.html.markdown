---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_service_provider"
description: |-
  SAML Service provider related information in Intersight.
---

# Data Source: intersight_iam_service_provider
SAML Service provider related information in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_service_provider.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `entity_id`:(string) Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider. 
* `metadata`:(string) Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Intersight Service Provider. 
 