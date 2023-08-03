---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_target"
description: |-
        Target represents an entity which can be managed by Intersight. This includes physical entities like UCS and HyperFlex servers and software entities like VMware vCenter and Microsoft Azure cloud accounts.

---

# Resource: intersight_asset_target
Target represents an entity which can be managed by Intersight. This includes physical entities like UCS and HyperFlex servers and software entities like VMware vCenter and Microsoft Azure cloud accounts.
## Usage Example
### Resource Creation

```hcl
resource "intersight_asset_target" "asset_target1" {
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
  connections {
    object_type = "asset.IntersightDeviceConnectorConnection"
    class_id    = "asset.IntersightDeviceConnectorConnection"
  }
  moid = var.asset_target1
  name = "C240-WZP21330QS5"
  registered_device {
    moid        = var.registered_device
    object_type = "asset.DeviceRegistrations"
  }
  target_type = "IMCM5"
}

variable "registered_device" {
  type        = string
  description = "Moid of registered device"
}

variable "asset_target1" {
  type        = string
  description = "Moid of asset.Target Mo"
}

variable "account" {
  type        = string
  description = "Moid of iam.Account Mo"
}
```

## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `alarm_summary`:(HashMap) -(ReadOnly) AlarmSummary is a health score calculation and aggregation of alarm counts per severity of alarms which Intersight's ability to manage the Target. 
This complex property has following sub-properties:
  + `critical`:(int)(ReadOnly) The count of active alarms that have severity type Critical. 
  + `health`:(string)(ReadOnly) Health of the managed end point. The highest severity computed from alarmSummary property is set as the health.* `Healthy` - The Enum value represents that the entity is healthy.* `Warning` - The Enum value Warning represents that the entity has one or more active warnings on it.* `Critical` - The Enum value Critical represents that the entity is in a critical state. 
  + `info`:(int)(ReadOnly) The count of active alarms that have severity type Info. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `warning`:(int)(ReadOnly) The count of active alarms that have severity type Warning. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `assist`:(HashMap) - A reference to a assetTarget resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `claimed_by_user_name`:(string)(ReadOnly) The name or email id of the user who claimed the target. 
* `connections`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [asset.CloudConnection](#assetCloudConnection)
[asset.HttpConnection](#assetHttpConnection)
[asset.IntersightDeviceConnectorConnection](#assetIntersightDeviceConnectorConnection)
[asset.ScopedTargetConnection](#assetScopedTargetConnection)
[asset.SshConnection](#assetSshConnection)
  + `credential`:(HashMap) - The credential to be used to authenticate with the managed target. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `connector_version`:(string)(ReadOnly) The Device Connector version for target types which are managed by via embedded Device Connector. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `external_ip_address`:(string)(ReadOnly) ExternalIpAddress is applicable for targets which are managed via an Intersight Device Connector. The value is the IP Address of the target as seen from Intersight. It is either the IP Address of the managed target's interface which has a route to the internet or a NAT IP Address when the target is deployed in a private network. 
* `ip_address`:
                (Array of schema.TypeString) -
* `management_location`:(string) The location from which Intersight manages the target.* `Unknown` - The management mechanism is not detected. Unknown is used as a default by the implementation and is not an allowed user input.* `Intersight` - Management of a target is performed directly from Intersight. Network connections are established from Intersight to a management interface of the Target and authenticated using user provided credentials.* `IntersightAssist` - Management of a target is performed via a selected Intersight Assist. Network connections are established from the Intersight Assist to a management interface of the Target.* `DeviceConnector` - An Intersight Device Connector running within the Target establishes a connection to Intersight and management of the target is performed via this connection. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user provided name for the managed target. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
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
* `product_id`:
                (Array of schema.TypeString) -
* `read_only`:(bool)(ReadOnly) For targets which are managed by an embedded Intersight Device Connector, this field indicates that an administrator of the target has disabled management operations of the Device Connector and only monitoring is permitted. 
* `registered_device`:(HashMap) -(ReadOnly) A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `services`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [asset.OrchestrationService](#assetOrchestrationService)
[asset.TerraformIntegrationService](#assetTerraformIntegrationService)
[asset.VirtualizationService](#assetVirtualizationService)
[asset.WorkloadOptimizerService](#assetWorkloadOptimizerService)
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `options`:(HashMap) - Captures configuration that is specific to a target type for a specific service. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string)(ReadOnly) Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target.* `` - The target details have been persisted but Intersight has not yet attempted to connect to the target.* `Connected` - Intersight is able to establish a connection to the target and initiate management activities.* `NotConnected` - Intersight is unable to establish a connection to the target.* `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established.* `UnclaimInProgress` - Unclaim of the target is in progress. Intersight is able to connect to the target and all management operations are supported.* `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.* `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect. 
* `status_error_reason`:(string)(ReadOnly) StatusErrorReason provides additional context for the Status. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target_id`:
                (Array of schema.TypeString) -
* `target_type`:(string) The type of the managed target. For example a UCS Server or VMware Vcenter target.* `` - An unrecognized platform type.* `APIC` - A Cisco Application Policy Infrastructure Controller (APIC) cluster.* `CAPIC` - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance.* `DCNM` - A Cisco Data Center Network Manager (DCNM) instance.* `UCSFI` - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM).* `UCSFIISM` - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight.* `IMC` - A standalone Cisco UCS rack server (Deprecated).* `IMCM4` - A standalone Cisco UCS C-Series or S-Series M4 server.* `IMCM5` - A standalone Cisco UCS C-Series or S-Series M5 server.* `IMCRack` - A standalone Cisco UCS C-Series or S-Series M6 or newer server.* `UCSIOM` - A Cisco UCS Blade Chassis I/O Module (IOM).* `HX` - A Cisco HyperFlex (HX) cluster.* `UCSD` - A Cisco UCS Director (UCSD) instance.* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance instance.* `IntersightAssist` - A Cisco Intersight Assist instance.* `PureStorageFlashArray` - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device.* `NexusDevice` - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist.* `ACISwitch` - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric.* `NexusSwitch` - A standalone Cisco Nexus Network Switch with the embedded Device Connector.* `MDSSwitch` - A Cisco MDS Switch that is managed using the embedded Device Connector.* `MDSDevice` - A Cisco MDS Switch that is managed using Cisco Intersight Assist.* `UCSC890` - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist.* `RedfishServer` - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers.* `NetAppOntap` - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist.* `EmcScaleIo` - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `EmcVmax` - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `EmcVplex` - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `EmcXtremIo` - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `VmwareVcenter` - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor.* `MicrosoftHyperV` - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor.* `AppDynamics` - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller.* `Dynatrace` - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server.* `NewRelic` - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server.* `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks.* `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.* `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.* `MicrosoftSqlServer` - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database.* `MySqlServer` - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database.* `OracleDatabaseServer` - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database.* `IBMWebSphereApplicationServer` - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server.* `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE.* `ApacheTomcatServer` - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server.* `JavaVirtualMachine` - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application.* `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software.* `Kubernetes` - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster.* `AmazonWebService` - An Amazon Web Service cloud account.  Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud.* `AmazonWebServiceBilling` - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket.  Cisco Intersight Workload Optimizer features are supported on this cloud.* `GoogleCloudPlatform` - A Google Cloud Platform service account with access to one or more projects.  Cisco Intersight Workload Optimizer features are supported on this cloud.* `GoogleCloudPlatformBilling` - A Google Cloud Platform service account used to retrieve billing information from BigQuery.  Cisco Intersight Workload Optimizer features are supported on this cloud.* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal account with access to Azure subscriptions.  Cisco Intersight Workload Optimizer features are supported on this cloud.* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud.* `MicrosoftAzureBilling` - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud.* `DellCompellent` - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `HPE3Par` - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device.* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.* `NutanixAcropolis` - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster.* `HPEOneView` - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist.* `GenericTarget` - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment.* `IMCBlade` - A Cisco UCS blade server managed by Cisco Intersight.* `TerraformCloud` - A Terraform Cloud Business Tier account.* `TerraformAgent` - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace.* `CustomTarget` - CustomTarget is deprecated.  Use HTTPEndpoint type to claim HTTP endpoints.* `AnsibleEndpoint` - An external endpoint that is added as a target  which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows.* `HTTPEndpoint` - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows  directly or using Cisco Intersight Assist.  Authentication Schemes supported are Basic and Bearer Token.* `SSHEndpoint` - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist.* `CiscoCatalyst` - A Cisco Catalyst networking switch device.* `PowerShellEndpoint` - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. 
* `trust_point`:(HashMap) - A reference to a iamTrustPoint resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `vendor`:(string)(ReadOnly) The vendor of the managed target. 
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
* `workflow_info`:(HashMap) -(ReadOnly) A reference to a workflowWorkflowInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 

### Custom keywords
These are
* `wait_for_completion`:(bool) This model object can trigger workflows. Use this option to wait for all running workflows to reach a complete state. Default value is True i.e. wait.

## Import
`intersight_asset_target` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_asset_target.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [asset.CloudConnection](#argument-reference)
CloudConnection provides the necessary details for Intersight to connect to and authenticate with a target at a well-known service address. The service address is inferred based upon the target type. For example Amazon Web Services.

### [asset.HttpConnection](#argument-reference)
HttpConnection provides the necessary details for Intersight to connect to and authenticate with a managed target over an HTTP connection.
* `certificate_authority`:(string) The certificate authority of the target. If set and connection is secure the connection will be validate using the servers identity with this certificate. If not set no validation will be done of the identity. 
* `is_secure`:(bool) Indicates whether a connection to the target should be established using HTTPS. 
* `management_address`:(string) The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target. 
* `port`:(int) The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection. 

### [asset.IntersightDeviceConnectorConnection](#argument-reference)
Target is connected to Intersight using a Device Connector.

### [asset.ScopedTargetConnection](#argument-reference)
ScopedTargetConnection provides the necessary details for Intersight to connect to and authenticate with a managed scoped target such as msSQL, mySQL and Oracle Database etc.
* `full_validation`:(bool) When this flag is set to true, every IWO entity in the scope targets will be checked and discovery of the scope target will be regarded as a failure when anyone of these entities cannot be connected and validated. 
* `is_secure`:(bool) Indicates whether a connection to the target should be established using TLS. 
* `port`:(int) The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection. 
* `scope`:(string) The group id of IWO entities upon which the discover of a scoped target is performed. For example, a database target may be scoped to the group of virtual machines upon which the database application is running. Scope value is group id created for all those virtual machines in this scope. 

### [asset.SshConnection](#argument-reference)
SshConnection provides the necessary details for Intersight to connect and authenticate with a managed target over an SSH connection.
* `management_address`:(string) The DNS hostname or IPv4 Address to be used to connect to the managed target. 
* `port`:(int) The port number to be used to connect to the managed target. Valid values are 1 - 65535. If not provided, a default port of 22 is used to establish the SSH connection to the given target. 
  
### [asset.OrchestrationService](#argument-reference)
OrchestrationService provides the necessary configuration details to enable Intersight Orchestration on the selected managed target. Subject to licensing.

### [asset.TerraformIntegrationService](#argument-reference)
TerraformIntegrationService provides the necessary configuration details to enable Intersight Cloud Region on the selected Terraform Cloud.

### [asset.VirtualizationService](#argument-reference)
The necessary configuration details to enable  Intersight Virtualization features on the selected managed target.

### [asset.WorkloadOptimizerService](#argument-reference)
WorkloadOptimizerService provides the necessary configuration details to enable Intersight Workflow Optimizer on the selected managed target. Subject to licensing.
  
