---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_exempted_catalog"
description: |-
  Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.
---

# Data Source: intersight_hcl_exempted_catalog
Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `comments`:(string) Reason for the exemption. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique descriptive name of the exemption. 
* `os_vendor`:(string) Vendor of the Operating System. 
* `os_version`:(string) Version of the Operating system. 
* `processor_name`:(string) Name of the processor supported for the server. 
* `product_type`:(string) Type of the product/adapter say GPU for graphic cards.* `` - Default type of the Product.* `Adapter` - Represents network adapter cards.* `StorageController` - Represents storage controllers.* `GPU` - Represents graphics cards. 
* `server_pid`:(string) Three part ID representing the server model as returned by UCSM/CIMC XML APIs. 
* `ucs_version`:(string) Version of the UCS software. 
* `version_type`:(string) Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. 
