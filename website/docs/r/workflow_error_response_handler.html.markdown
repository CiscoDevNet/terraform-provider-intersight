---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_error_response_handler"
description: |-
  Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.
Error Response Handler allows to create a generic error response specification
which can be used by multiple Batch API. The parameters provided in the Error
Response Handler may be used to parse error responses from an API request, if
the response specification provided for the API request does not define
error parameters.
---

# Resource: intersight_workflow_error_response_handler
Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.
Error Response Handler allows to create a generic error response specification
which can be used by multiple Batch API. The parameters provided in the Error
Response Handler may be used to parse error responses from an API request, if
the response specification provided for the API request does not define
error parameters.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `catalog`:(Array with Maximum of one item) - A reference to a workflowCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) A detailed description about the error response handler. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the error response handler. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `parameters`:(Array)
This complex property has following sub-properties:
  + `accept_single_value`:(bool) The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only. 
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `complex_type`:(string) The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property. 
  + `item_type`:(string) The type of the collection item in case this is a collection parameter.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
  + `name`:(string) The name of the parameter. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `path`:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. 
  + `type`:(string) The type of the parameter. Accepted values are simple, complex,collection.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
* `platform_type`:(string) The platform type for which the error response handler is defined.* `` - The device reported an empty or unrecognized platform type.* `APIC` - An Application Policy Infrastructure Controller cluster.* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.* `IMC` - A standalone UCS Server Integrated Management Controller.* `IMCM4` - A standalone UCS M4 Server.* `IMCM5` - A standalone UCS M5 server.* `UCSIOM` - An UCS Chassis IO module.* `HX` - A HyperFlex storage controller.* `HyperFlexAP` - A HyperFlex Application Platform.* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.* `IntersightAssist` - A Cisco Intersight Assist.* `PureStorageFlashArray` - A Pure Storage FlashArray device.* `NetAppOntap` - A NetApp ONTAP storage system.* `EmcScaleIo` - An EMC ScaleIO storage system.* `EmcVmax` - An EMC VMAX storage system.* `EmcVplex` - An EMC VPLEX storage system.* `EmcXtremIo` - An EMC XtremIO storage system.* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.* `AppDynamics` - An AppDynamics controller that monitors applications.* `Dynatrace` - A Dynatrace controller that monitors applications.* `MicrosoftSqlServer` - A Microsoft SQL database server.* `Kubernetes` - A Kubernetes cluster that runs containerized applications.* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.* `IMCBlade` - An Intersight managed UCS Blade Server.* `CustomTarget` - An external endpoint added as Target that can be accessed through its REST API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.* `CiscoCatalyst` - A Cisco Catalyst networking switch device. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `types`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `name`:(string) The unique name of this complex type within the grammar specification. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `parameters`:(Array)
This complex property has following sub-properties:
    + `accept_single_value`:(bool) The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only. 
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
    + `complex_type`:(string) The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property. 
    + `item_type`:(string) The type of the collection item in case this is a collection parameter.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
    + `name`:(string) The name of the parameter. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
    + `path`:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. 
    + `type`:(string) The type of the parameter. Accepted values are simple, complex,collection.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 


## Import
`intersight_workflow_error_response_handler` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_error_response_handler.example 1234567890987654321abcde
```