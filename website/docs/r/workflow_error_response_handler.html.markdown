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
## Usage Example
### Resource Creation

```hcl
resource "intersight_workflow_error_response_handler" "workflow_error_response_handler1" {
  name          = "workflow_error_response_handler1"
  platform_type = "UCSD"
  parameters = [{
    object_type         = "content.TextParameter"
    accept_single_value = true
    name                = "show-pure"
    item_type           = "string"

  }]
  types {
    object_type = "content.ComplexType"
    parameters = [{
      object_type = "content.TextParameter"
      name        = "show-hitachi"
      item_type   = "string"
      type        = "string"
    }]
  }
  catalog {
    object_type = "workflow.Catalog"
    moid        = var.workflow_catalog
  }

}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(Computed) The Account ID for this managed object. 
* `ancestors`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `catalog`:(HashMap) - A reference to a workflowCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(Computed) The time when this managed object was created. 
* `description`:(string) A detailed description about the error response handler. 
* `domain_group_moid`:(string)(Computed) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(Computed) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the error response handler. 
* `owners`:
                (Array of schema.TypeString) -(Computed)
* `parameters`:(Array)
This complex property has following sub-properties:
  + `accept_single_value`:(bool) The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only. 
  + `complex_type`:(string) The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property. 
  + `item_type`:(string) The type of the collection item in case this is a collection parameter.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
  + `name`:(string) The name of the parameter. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `path`:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. 
  + `secure`:(bool) The flag indicates if the extracted value is secure.This flag is applicable for parameters of type String only. 
  + `type`:(string) The type of the parameter. Accepted values are simple, complex,collection.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
* `parent`:(HashMap) -(Computed) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `platform_type`:(string) The platform type for which the error response handler is defined.* `` - The device reported an empty or unrecognized platform type.* `APIC` - An Application Policy Infrastructure Controller cluster.* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.* `IMC` - A standalone UCS Server Integrated Management Controller.* `IMCM4` - A standalone UCS M4 Server.* `IMCM5` - A standalone UCS M5 server.* `IMCRack` - A standalone UCS M6 and above server.* `UCSIOM` - An UCS Chassis IO module.* `HX` - A HyperFlex storage controller.* `HyperFlexAP` - A HyperFlex Application Platform.* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.* `IntersightAssist` - A Cisco Intersight Assist.* `PureStorageFlashArray` - A Pure Storage FlashArray device.* `UCSC890` - A standalone Cisco UCSC890 server.* `NetAppOntap` - A NetApp ONTAP storage system.* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.* `EmcScaleIo` - An EMC ScaleIO storage system.* `EmcVmax` - An EMC VMAX storage system.* `EmcVplex` - An EMC VPLEX storage system.* `EmcXtremIo` - An EMC XtremIO storage system.* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.* `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines.* `AppDynamics` - An AppDynamics controller that monitors applications.* `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation.* `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.* `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.* `MicrosoftSqlServer` - A Microsoft SQL database server.* `Kubernetes` - A Kubernetes cluster that runs containerized applications.* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.* `DellCompellent` - A Dell Compellent storage system.* `HPE3Par` - A HPE 3PAR storage system.* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.* `IMCBlade` - An Intersight managed UCS Blade Server.* `TerraformCloud` - A Terraform Cloud account.* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.* `CiscoCatalyst` - A Cisco Catalyst networking switch device. 
* `shared_scope`:(string)(Computed) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `types`:(Array)
This complex property has following sub-properties:
  + `name`:(string) The unique name of this complex type within the grammar specification. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `parameters`:(Array)
This complex property has following sub-properties:
    + `accept_single_value`:(bool) The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only. 
    + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [content.Parameter](#contentParameter)
[content.TextParameter](#contentTextParameter)
    + `complex_type`:(string) The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property. 
    + `item_type`:(string) The type of the collection item in case this is a collection parameter.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
    + `name`:(string) The name of the parameter. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
    + `path`:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. 
    + `secure`:(bool) The flag indicates if the extracted value is secure.This flag is applicable for parameters of type String only. 
    + `type`:(string) The type of the parameter. Accepted values are simple, complex,collection.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `string` - The parameter value to be extracted is of the string type.* `integer` - The parameter value to be extracted is of the number type.* `float` - The parameter value to be extracted is of the float number type.* `boolean` - The parameter value to be extracted is of the boolean type.* `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
* `version_context`:(HashMap) -(Computed) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(Computed) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(Computed) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(Computed) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(Computed) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_workflow_error_response_handler` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_error_response_handler.example 1234567890987654321abcde
``` 