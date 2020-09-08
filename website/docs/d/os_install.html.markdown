
---
layout: "intersight"
page_title: "Intersight: intersight_os_install"
sidebar_current: "docs-intersight-data-source-os-install"
description: |-
This MO models the target server, answers and other properties needed for
OS installation. The OS installation can be started in the target server by doing
a POST on this MO.
The requests to this MO starts a OS installation workflow that can be tracked
using workflow engine MO workflow.WorkflowInfo.
---

# Data Source: intersight_os._install
This MO models the target server, answers and other properties needed for
OS installation. The OS installation can be started in the target server by doing
a POST on this MO.
The requests to this MO starts a OS installation workflow that can be tracked
using workflow engine MO workflow.WorkflowInfo.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) User provided description about the OS install configuration. 
* `install_method`:(string) The install method to be used for OS installation - vMedia, iPXE. Only vMedia is supported as of now.* `vMedia` - OS image is mounted as vMedia in target server for OS installation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS install configuration. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
