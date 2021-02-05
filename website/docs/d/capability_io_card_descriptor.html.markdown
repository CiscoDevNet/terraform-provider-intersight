---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_io_card_descriptor"
description: |-
  Descriptor that uniquely identifies an IO card module.
---

# Data Source: intersight_capability_io_card_descriptor
Descriptor that uniquely identifies an IO card module.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Detailed information about the endpoint. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_hif_ports`:(int) Number of hif ports per blade for the iocard module. 
* `revision`:(string) Revision for the iocard module. 
* `uif_connectivity`:(string) Connectivity information between UIF Uplink ports and IOM ports.* `inline` - UIF uplink ports and IOM ports are connected inline.* `cross-connected` - UIF uplink ports and IOM ports are cross-connected, a case in washington chassis. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
