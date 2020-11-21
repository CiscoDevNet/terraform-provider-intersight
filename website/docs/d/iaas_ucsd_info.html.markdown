
---
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_info"
sidebar_current: "docs-intersight-data-source-iaas-ucsd-info"
description: |-
UCS Director accounts managed by Intersight.
---

# Data Source: intersight_iaas._ucsd_info
UCS Director accounts managed by Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_id`:(string) Moid of the UCS Director device connector's asset.DeviceRegistration. 
* `guid`:(string) Unique ID of UCS Director getting registerd with Intersight. 
* `host_name`:(string) The UCS Director hostname for management. 
* `ip`:(string) The UCS Director IP address for management. 
* `last_backup`:(string) Last successful backup created for this UCS Director appliance if backup is configured. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_type`:(string) NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `product_name`:(string) The UCS Director product name. 
* `product_vendor`:(string) The UCS Director product vendor. 
* `product_version`:(string) The UCS Director product/platform version. 
* `status`:(string) The UCS Director status. Possible values are Active, Inactive, Unknown. 
