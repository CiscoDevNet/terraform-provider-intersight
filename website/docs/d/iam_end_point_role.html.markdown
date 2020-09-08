
---
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_role"
sidebar_current: "docs-intersight-data-source-iam-end-point-role"
description: |-
The role defined in the end point which can be assigned to a user.
---

# Data Source: intersight_iam._end_point_role
The role defined in the end point which can be assigned to a user.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the end point role. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `role_type`:(string) User specified tags to roles like as ep-admin or ep-readonly. 
* `type`:(string) The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.* `` - The device reported an empty or unrecognized platform type.* `APIC` - An Application Policy Infrastructure Controller cluster.* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.* `IMC` - A standalone UCS Server Integrated Management Controller.* `IMCM4` - A standalone UCS M4 Server.* `IMCM5` - A standalone UCS M5 server.* `UCSIOM` - An UCS Chassis IO module.* `HX` - A HyperFlex storage controller.* `HyperFlexAP` - A HyperFlex Application Platform.* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.* `IntersightAppliance` - Intersight on-premise appliance.* `PureStorageFlashArray` - A Pure Storage FlashArray device.* `NetAppOntap` - A NetApp ONTAP storage system.* `EmcScaleIo` - An EMC ScaleIO storage system.* `EmcVmax` - An EMC VMAX storage system.* `EmcVplex` - An EMC VPLEX storage system.* `EmcXtremIo` - An EMC XtremIO storage system.* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.* `AppDynamics` - An AppDynamics controller that monitors applications.* `Dynatrace` - A Dynatrace controller that monitors applications.* `MicrosoftSqlServer` - A Microsoft SQL database server.* `Kubernetes` - A Kubernetes cluster that runs containerized applications.* `MicrosoftAzure` - A Microsoft Azure target.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `IMCBlade` - An Intersight managed UCS Blade Server. 
