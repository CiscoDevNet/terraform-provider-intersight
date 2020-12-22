---
subcategory: "techsupportmanagement"
layout: "intersight"
page_title: "Intersight: intersight_techsupportmanagement_tech_support_bundle"
description: |-
  A request to collect techsupport and upload it to Intersight Storage Service. The serial number, PID and/or relationship to the target resource provided by the user is used to determine the device type for techsupport collection. If the serial number, PID and target resource are specified in the request, the values must match. Valid values of device types are network.Element for fabric interconnect, compute.Blade for blade server, compute.RackUnit for rack server, equipment.Chassis for chassis, equipment.IoCard for IO Module, equipment.FEX for fabric extender and adapter.Unit for network adapter. UCSM techsupport is collected for device type network.Element. Chassis techsupport is collected for compute.Blade, equipment.Chassis, equipment.IoCard, and blade adapter.Unit. Server techsupport is collected for compute.RackUnit and rack adapter.Unit. Fabric extender techsupport is collected for device type equipment.FEX. Hyper Flex node level techsupport is collected when the request specifies the platform type (HX) and the device type is Hyperflex.Node.
---

# Resource: intersight_techsupportmanagement_tech_support_bundle
A request to collect techsupport and upload it to Intersight Storage Service. The serial number, PID and/or relationship to the target resource provided by the user is used to determine the device type for techsupport collection. If the serial number, PID and target resource are specified in the request, the values must match. Valid values of device types are network.Element for fabric interconnect, compute.Blade for blade server, compute.RackUnit for rack server, equipment.Chassis for chassis, equipment.IoCard for IO Module, equipment.FEX for fabric extender and adapter.Unit for network adapter. UCSM techsupport is collected for device type network.Element. Chassis techsupport is collected for compute.Blade, equipment.Chassis, equipment.IoCard, and blade adapter.Unit. Server techsupport is collected for compute.RackUnit and rack adapter.Unit. Fabric extender techsupport is collected for device type equipment.FEX. Hyper Flex node level techsupport is collected when the request specifies the platform type (HX) and the device type is Hyperflex.Node.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_identifier`:(string)(Computed) The device identifier used to uniquely identify an individual device. 
* `device_registration`:(Array with Maximum of one item) -(Computed) A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `device_type`:(string)(Computed) The device type obtained from the inventory. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pid`:(string) Product identification of the device. 
* `platform_param`:(Array with Maximum of one item) - A platform specific data payload. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `platform_type`:(string) The platform type of the device.* `` - The device reported an empty or unrecognized platform type.* `APIC` - An Application Policy Infrastructure Controller cluster.* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.* `IMC` - A standalone UCS Server Integrated Management Controller.* `IMCM4` - A standalone UCS M4 Server.* `IMCM5` - A standalone UCS M5 server.* `UCSIOM` - An UCS Chassis IO module.* `HX` - A HyperFlex storage controller.* `HyperFlexAP` - A HyperFlex Application Platform.* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.* `IntersightAssist` - A Cisco Intersight Assist.* `PureStorageFlashArray` - A Pure Storage FlashArray device.* `NetAppOntap` - A NetApp ONTAP storage system.* `EmcScaleIo` - An EMC ScaleIO storage system.* `EmcVmax` - An EMC VMAX storage system.* `EmcVplex` - An EMC VPLEX storage system.* `EmcXtremIo` - An EMC XtremIO storage system.* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.* `AppDynamics` - An AppDynamics controller that monitors applications.* `Dynatrace` - A Dynatrace controller that monitors applications.* `MicrosoftSqlServer` - A Microsoft SQL database server.* `Kubernetes` - A Kubernetes cluster that runs containerized applications.* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.* `IMCBlade` - An Intersight managed UCS Blade Server.* `CustomTarget` - An external endpoint added as Target that can be accessed through its REST API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.* `CiscoCatalyst` - A Cisco Catalyst networking switch device. 
* `serial`:(string) Serial number of the device. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target_resource`:(Array with Maximum of one item) - A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tech_support_status`:(Array with Maximum of one item) -(Computed) A reference to a techsupportmanagementTechSupportStatus resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 


## Import
`intersight_techsupportmanagement_tech_support_bundle` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_techsupportmanagement_tech_support_bundle.example 1234567890987654321abcde
```