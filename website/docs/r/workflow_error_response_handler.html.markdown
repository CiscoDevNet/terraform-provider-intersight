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
  parameters {
    class_id            = "content.TextParameter"
    secure              = false
    accept_single_value = true
    object_type         = "content.TextParameter"
    name                = "show-pure"
    item_type           = "string"
    type                = "string"
  }
  types {
    object_type = "content.ComplexType"
    parameters {
      class_id            = "content.TextParameter"
      secure              = false
      accept_single_value = true
      object_type         = "content.TextParameter"
      name                = "show-hitachi"
      item_type           = "string"
      type                = "string"
    }
  }
  catalog {
    object_type = "workflow.Catalog"
    moid        = var.workflow_catalog
  }

}

variable "workflow_catalog" {
  type        = string
  description = "moid for workflow catalog"
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `catalog`:(HashMap) - A reference to a workflowCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) A detailed description about the error response handler. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the error response handler. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
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
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `platform_type`:(string) The platform type for which the error response handler is defined.* `` - An unrecognized platform type.* `APIC` - A Cisco Application Policy Infrastructure Controller (APIC) cluster.* `CAPIC` - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance.* `DCNM` - A Cisco Data Center Network Manager (DCNM) instance.* `UCSFI` - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM).* `UCSFIISM` - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight.* `IMC` - A standalone Cisco UCS rack server (Deprecated).* `IMCM4` - A standalone Cisco UCS C-Series or S-Series M4 server.* `IMCM5` - A standalone Cisco UCS C-Series or S-Series M5 server.* `IMCRack` - A standalone Cisco UCS C-Series or S-Series M6 or newer server.* `UCSIOM` - A Cisco UCS Blade Chassis I/O Module (IOM).* `HX` - A Cisco HyperFlex (HX) cluster.* `UCSD` - A Cisco UCS Director (UCSD) instance.* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance instance.* `IntersightAssist` - A Cisco Intersight Assist instance.* `PureStorageFlashArray` - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device.* `NexusDevice` - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist.* `ACISwitch` - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric.* `NexusSwitch` - A standalone Cisco Nexus Network Switch with the embedded Device Connector.* `MDSSwitch` - A Cisco MDS Switch that is managed using the embedded Device Connector.* `MDSDevice` - A Cisco MDS Switch that is managed using Cisco Intersight Assist.* `UCSC890` - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist.* `RedfishServer` - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers.* `NetAppOntap` - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist.* `EmcScaleIo` - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `EmcVmax` - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `EmcVplex` - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `EmcXtremIo` - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `VmwareVcenter` - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor.* `MicrosoftHyperV` - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor.* `AppDynamics` - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller.* `Dynatrace` - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server.* `NewRelic` - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server.* `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks.* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.* `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.* `MicrosoftSqlServer` - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database.* `MySqlServer` - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database.* `OracleDatabaseServer` - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database.* `IBMWebSphereApplicationServer` - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server.* `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE.* `ApacheTomcatServer` - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server.* `JavaVirtualMachine` - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application.* `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software.* `Kubernetes` - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster.* `AmazonWebService` - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud.* `AmazonWebServiceBilling` - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud.* `GoogleCloudPlatform` - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud.* `GoogleCloudPlatformBilling` - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud.* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud.* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud.* `MicrosoftAzureBilling` - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud.* `DellCompellent` - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `HPE3Par` - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.* `NutanixAcropolis` - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster.* `HPEOneView` - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist.* `GenericTarget` - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment.* `IMCBlade` - A Cisco UCS blade server managed by Cisco Intersight.* `TerraformCloud` - A Terraform Cloud Business Tier account.* `TerraformAgent` - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace.* `CustomTarget` - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints.* `AnsibleEndpoint` - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows.* `HTTPEndpoint` - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token.* `SSHEndpoint` - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist.* `CiscoCatalyst` - A Cisco Catalyst networking switch device.* `PowerShellEndpoint` - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist.* `CiscoDNAC` - A Cisco Digital Network Architecture (DNA) Center appliance.* `CiscoFMC` - A Cisco Secure Firewall Management Center.* `ViptelaCloud` - A Cisco Viptela SD-WAN Cloud.* `MerakiCloud` - A Cisco Meraki Organization.* `CiscoISE` - A Cisco Identity Services Engine (ISE) target. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
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
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_workflow_error_response_handler` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_error_response_handler.example 1234567890987654321abcde
``` 
