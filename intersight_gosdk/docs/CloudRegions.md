# CloudRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.Regions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.Regions"]
**AlternateNames** | Pointer to **[]string** |  | [optional] 
**DefaultZone** | Pointer to **string** | The default zone for this region. | [optional] 
**Group** | Pointer to **string** | Property to identify regions in same category which shares credentials. For e.g. AWS Govcloud Vs AWS Global Vs AWS China. | [optional] 
**IsActive** | Pointer to **bool** | Flag to indicate of this region is active or not. | [optional] [default to true]
**IsBillingOnly** | Pointer to **bool** | Property to pick regions for orchestration. | [optional] 
**Name** | Pointer to **string** | The display name of the region. | [optional] 
**PlatformType** | Pointer to **string** | The platform type for this region. For e.g. AmazonWebService. * &#x60;&#x60; - An unrecognized platform type. * &#x60;APIC&#x60; - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * &#x60;CAPIC&#x60; - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * &#x60;DCNM&#x60; - A Cisco Data Center Network Manager (DCNM) instance. * &#x60;UCSFI&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * &#x60;IMC&#x60; - A standalone Cisco UCS rack server (Deprecated). * &#x60;IMCM4&#x60; - A standalone Cisco UCS C-Series or S-Series M4 server. * &#x60;IMCM5&#x60; - A standalone Cisco UCS C-Series or S-Series M5 server. * &#x60;IMCRack&#x60; - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * &#x60;UCSIOM&#x60; - A Cisco UCS Blade Chassis I/O Module (IOM). * &#x60;HX&#x60; - A Cisco HyperFlex (HX) cluster. * &#x60;UCSD&#x60; - A Cisco UCS Director (UCSD) instance. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance instance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist instance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * &#x60;NexusDevice&#x60; - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * &#x60;ACISwitch&#x60; - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * &#x60;NexusSwitch&#x60; - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * &#x60;MDSSwitch&#x60; - A Cisco MDS Switch that is managed using the embedded Device Connector. * &#x60;MDSDevice&#x60; - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * &#x60;UCSC890&#x60; - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * &#x60;NetAppOntap&#x60; - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVmax&#x60; - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVplex&#x60; - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;VmwareVcenter&#x60; - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * &#x60;AppDynamics&#x60; - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * &#x60;Dynatrace&#x60; - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;NewRelic&#x60; - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;Umbrella&#x60; - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;MySqlServer&#x60; - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;OracleDatabaseServer&#x60; - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;IBMWebSphereApplicationServer&#x60; - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;JavaVirtualMachine&#x60; - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * &#x60;AmazonWebService&#x60; - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatform&#x60; - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatformBilling&#x60; - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureBilling&#x60; - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;DellCompellent&#x60; - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;HPE3Par&#x60; - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * &#x60;NutanixPrismCentral&#x60; - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * &#x60;HPEOneView&#x60; - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * &#x60;GenericTarget&#x60; - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * &#x60;IMCBlade&#x60; - A Cisco UCS blade server managed by Cisco Intersight. * &#x60;TerraformCloud&#x60; - A Terraform Cloud Business Tier account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * &#x60;CustomTarget&#x60; - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * &#x60;AnsibleEndpoint&#x60; - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * &#x60;HTTPEndpoint&#x60; - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * &#x60;SSHEndpoint&#x60; - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * &#x60;CiscoDNAC&#x60; - A Cisco Digital Network Architecture (DNA) Center appliance. * &#x60;CiscoFMC&#x60; - A Cisco Secure Firewall Management Center. * &#x60;ViptelaCloud&#x60; - A Cisco Viptela SD-WAN Cloud. * &#x60;MerakiCloud&#x60; - A Cisco Meraki Organization. * &#x60;CiscoISE&#x60; - A Cisco Identity Services Engine (ISE) target. | [optional] [default to ""]
**RegionEndPoint** | Pointer to **string** | HTTP endpoint of the region. For example https://ec2.us-east-2.amazonaws.com. | [optional] 
**RegionId** | Pointer to **string** | The region Id which is assigned by the cloud provider. For e.g. us-east-1. | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 
**Target** | Pointer to [**NullableAssetTargetRelationship**](AssetTargetRelationship.md) |  | [optional] 

## Methods

### NewCloudRegions

`func NewCloudRegions(classId string, objectType string, ) *CloudRegions`

NewCloudRegions instantiates a new CloudRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegionsWithDefaults

`func NewCloudRegionsWithDefaults() *CloudRegions`

NewCloudRegionsWithDefaults instantiates a new CloudRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudRegions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudRegions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudRegions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudRegions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudRegions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudRegions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlternateNames

`func (o *CloudRegions) GetAlternateNames() []string`

GetAlternateNames returns the AlternateNames field if non-nil, zero value otherwise.

### GetAlternateNamesOk

`func (o *CloudRegions) GetAlternateNamesOk() (*[]string, bool)`

GetAlternateNamesOk returns a tuple with the AlternateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateNames

`func (o *CloudRegions) SetAlternateNames(v []string)`

SetAlternateNames sets AlternateNames field to given value.

### HasAlternateNames

`func (o *CloudRegions) HasAlternateNames() bool`

HasAlternateNames returns a boolean if a field has been set.

### SetAlternateNamesNil

`func (o *CloudRegions) SetAlternateNamesNil(b bool)`

 SetAlternateNamesNil sets the value for AlternateNames to be an explicit nil

### UnsetAlternateNames
`func (o *CloudRegions) UnsetAlternateNames()`

UnsetAlternateNames ensures that no value is present for AlternateNames, not even an explicit nil
### GetDefaultZone

`func (o *CloudRegions) GetDefaultZone() string`

GetDefaultZone returns the DefaultZone field if non-nil, zero value otherwise.

### GetDefaultZoneOk

`func (o *CloudRegions) GetDefaultZoneOk() (*string, bool)`

GetDefaultZoneOk returns a tuple with the DefaultZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZone

`func (o *CloudRegions) SetDefaultZone(v string)`

SetDefaultZone sets DefaultZone field to given value.

### HasDefaultZone

`func (o *CloudRegions) HasDefaultZone() bool`

HasDefaultZone returns a boolean if a field has been set.

### GetGroup

`func (o *CloudRegions) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CloudRegions) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CloudRegions) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CloudRegions) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIsActive

`func (o *CloudRegions) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CloudRegions) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CloudRegions) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CloudRegions) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsBillingOnly

`func (o *CloudRegions) GetIsBillingOnly() bool`

GetIsBillingOnly returns the IsBillingOnly field if non-nil, zero value otherwise.

### GetIsBillingOnlyOk

`func (o *CloudRegions) GetIsBillingOnlyOk() (*bool, bool)`

GetIsBillingOnlyOk returns a tuple with the IsBillingOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingOnly

`func (o *CloudRegions) SetIsBillingOnly(v bool)`

SetIsBillingOnly sets IsBillingOnly field to given value.

### HasIsBillingOnly

`func (o *CloudRegions) HasIsBillingOnly() bool`

HasIsBillingOnly returns a boolean if a field has been set.

### GetName

`func (o *CloudRegions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRegions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRegions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRegions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformType

`func (o *CloudRegions) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *CloudRegions) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *CloudRegions) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *CloudRegions) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRegionEndPoint

`func (o *CloudRegions) GetRegionEndPoint() string`

GetRegionEndPoint returns the RegionEndPoint field if non-nil, zero value otherwise.

### GetRegionEndPointOk

`func (o *CloudRegions) GetRegionEndPointOk() (*string, bool)`

GetRegionEndPointOk returns a tuple with the RegionEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionEndPoint

`func (o *CloudRegions) SetRegionEndPoint(v string)`

SetRegionEndPoint sets RegionEndPoint field to given value.

### HasRegionEndPoint

`func (o *CloudRegions) HasRegionEndPoint() bool`

HasRegionEndPoint returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudRegions) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudRegions) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudRegions) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *CloudRegions) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetZones

`func (o *CloudRegions) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CloudRegions) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CloudRegions) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CloudRegions) HasZones() bool`

HasZones returns a boolean if a field has been set.

### SetZonesNil

`func (o *CloudRegions) SetZonesNil(b bool)`

 SetZonesNil sets the value for Zones to be an explicit nil

### UnsetZones
`func (o *CloudRegions) UnsetZones()`

UnsetZones ensures that no value is present for Zones, not even an explicit nil
### GetTarget

`func (o *CloudRegions) GetTarget() AssetTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudRegions) GetTargetOk() (*AssetTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudRegions) SetTarget(v AssetTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudRegions) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CloudRegions) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CloudRegions) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


