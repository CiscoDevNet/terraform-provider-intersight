
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_claim"
sidebar_current: "docs-intersight-resource-appliance-device-claim"
description: |-
  DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.
Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.
---

# Resource: intersight_appliance._device_claim
DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.
Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.
## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) -(Computed) A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_id`:(string)(Computed) Device identifier of the endpoint device. 
* `hostname`:(string) Hostname or IP address of the endpoint device the user wants to claim. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `message`:(string)(Computed) Message set by the device claim process. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `password`:(string) Password to be used to login to the endpoint device. 
* `platform_type`:(string) Platform type of the endpoint device.* `` - The device reported an empty or unrecognized platform type.* `APIC` - An Application Policy Infrastructure Controller cluster.* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.* `IMC` - A standalone UCS Server Integrated Management Controller.* `IMCM4` - A standalone UCS M4 Server.* `IMCM5` - A standalone UCS M5 server.* `UCSIOM` - An UCS Chassis IO module.* `HX` - A HyperFlex storage controller.* `HyperFlexAP` - A HyperFlex Application Platform.* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.* `IntersightAppliance` - Intersight on-premise appliance.* `PureStorageFlashArray` - A Pure Storage FlashArray device.* `NetAppOntap` - A NetApp ONTAP storage system.* `EmcScaleIo` - An EMC ScaleIO storage system.* `EmcVmax` - An EMC VMAX storage system.* `EmcVplex` - An EMC VPLEX storage system.* `EmcXtremIo` - An EMC XtremIO storage system.* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.* `AppDynamics` - An AppDynamics controller that monitors applications.* `Dynatrace` - A Dynatrace controller that monitors applications.* `MicrosoftSqlServer` - A Microsoft SQL database server.* `Kubernetes` - A Kubernetes cluster that runs containerized applications.* `MicrosoftAzure` - A Microsoft Azure target.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `IMCBlade` - An Intersight managed UCS Blade Server. 
* `request_id`:(string) User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set. 
* `security_token`:(string)(Computed) Device security token of the endpoint device. 
* `status`:(string)(Computed) Status of the device claim process.* `started` - Device claim operation has started.* `failed` - Device claim operation has failed.* `completed` - Device claim operation has completed. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `username`:(string) Username to log in to the endpoint device. 
