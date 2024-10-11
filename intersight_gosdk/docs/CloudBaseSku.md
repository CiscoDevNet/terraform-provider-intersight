# CloudBaseSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**BasePrice** | Pointer to **float32** | The cost per hour for running a Linux instance for the SKU. Cost can be higher for Windows or other licensed OS. | [optional] 
**Currency** | Pointer to **string** | The currency code used for the price. For e.g. USD or EUR. * &#x60;USD&#x60; - The currency code for United states dollar. * &#x60;EUR&#x60; - The currency code for European Union. | [optional] [default to "USD"]
**CustomAttributes** | Pointer to [**[]CloudCustomAttributes**](CloudCustomAttributes.md) |  | [optional] 
**Description** | Pointer to **string** | Any additional description for the instance type. | [optional] 
**IsActive** | Pointer to **bool** | Flag to indicate if this SKU is active or not. | [optional] [default to true]
**IsAutoDiscovered** | Pointer to **bool** | Flag to indicate if SKU is discovered during inventory collection. | [optional] 
**Name** | Pointer to **string** | The display name for instance type. | [optional] 
**PlatformType** | Pointer to **string** | The platformType for this SKU. * &#x60;&#x60; - An unrecognized platform type. * &#x60;APIC&#x60; - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * &#x60;CAPIC&#x60; - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * &#x60;DCNM&#x60; - A Cisco Data Center Network Manager (DCNM) instance. * &#x60;UCSFI&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * &#x60;IMC&#x60; - A standalone Cisco UCS rack server (Deprecated). * &#x60;IMCM4&#x60; - A standalone Cisco UCS C-Series or S-Series M4 server. * &#x60;IMCM5&#x60; - A standalone Cisco UCS C-Series or S-Series M5 server. * &#x60;IMCRack&#x60; - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * &#x60;UCSIOM&#x60; - A Cisco UCS Blade Chassis I/O Module (IOM). * &#x60;HX&#x60; - A Cisco HyperFlex (HX) cluster. * &#x60;UCSD&#x60; - A Cisco UCS Director (UCSD) instance. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance instance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist instance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * &#x60;NexusDevice&#x60; - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * &#x60;ACISwitch&#x60; - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * &#x60;NexusSwitch&#x60; - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * &#x60;MDSSwitch&#x60; - A Cisco MDS Switch that is managed using the embedded Device Connector. * &#x60;MDSDevice&#x60; - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * &#x60;UCSC890&#x60; - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * &#x60;NetAppOntap&#x60; - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVmax&#x60; - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVplex&#x60; - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;VmwareVcenter&#x60; - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * &#x60;AppDynamics&#x60; - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * &#x60;Dynatrace&#x60; - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;NewRelic&#x60; - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;Umbrella&#x60; - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;MySqlServer&#x60; - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;OracleDatabaseServer&#x60; - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;IBMWebSphereApplicationServer&#x60; - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;JavaVirtualMachine&#x60; - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * &#x60;AmazonWebService&#x60; - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatform&#x60; - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatformBilling&#x60; - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureBilling&#x60; - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;DellCompellent&#x60; - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;HPE3Par&#x60; - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * &#x60;NutanixPrismCentral&#x60; - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * &#x60;HPEOneView&#x60; - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * &#x60;GenericTarget&#x60; - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * &#x60;IMCBlade&#x60; - A Cisco UCS blade server managed by Cisco Intersight. * &#x60;TerraformCloud&#x60; - A Terraform Cloud Business Tier account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * &#x60;CustomTarget&#x60; - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * &#x60;AnsibleEndpoint&#x60; - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * &#x60;HTTPEndpoint&#x60; - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * &#x60;SSHEndpoint&#x60; - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * &#x60;CiscoDNAC&#x60; - A Cisco Digital Network Architecture (DNA) Center appliance. * &#x60;CiscoFMC&#x60; - A Cisco Secure Firewall Management Center. * &#x60;ViptelaCloud&#x60; - A Cisco Viptela SD-WAN Cloud. * &#x60;MerakiCloud&#x60; - A Cisco Meraki Organization. * &#x60;CiscoISE&#x60; - A Cisco Identity Services Engine (ISE) target. | [optional] [default to ""]
**RegionId** | Pointer to **string** | The region to which this Sku belongs to. | [optional] 
**ServiceCategory** | Pointer to **string** | Indicates if this sku belongs to Compute, Storage, Database or Network category. * &#x60;Compute&#x60; - Compute service offered by cloud provider. * &#x60;Storage&#x60; - Storage service offered by cloud provider. * &#x60;Database&#x60; - Database service offered by cloud provider. * &#x60;Network&#x60; - Network service offered by cloud provider. | [optional] [default to "Compute"]
**ServiceFamily** | Pointer to **string** | Property to identify the family of service that the sku belongs to. | [optional] 
**ServiceName** | Pointer to **string** | Any display name for the ServiceCategory if available. | [optional] 
**Unit** | Pointer to **string** | The pricing associated with this Sku. Can be hourly, per minute. | [optional] 
**Region** | Pointer to [**NullableCloudRegionsRelationship**](CloudRegionsRelationship.md) |  | [optional] 
**Target** | Pointer to [**NullableAssetTargetRelationship**](AssetTargetRelationship.md) |  | [optional] 

## Methods

### NewCloudBaseSku

`func NewCloudBaseSku(classId string, objectType string, ) *CloudBaseSku`

NewCloudBaseSku instantiates a new CloudBaseSku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseSkuWithDefaults

`func NewCloudBaseSkuWithDefaults() *CloudBaseSku`

NewCloudBaseSkuWithDefaults instantiates a new CloudBaseSku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseSku) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseSku) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseSku) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseSku) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseSku) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseSku) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBasePrice

`func (o *CloudBaseSku) GetBasePrice() float32`

GetBasePrice returns the BasePrice field if non-nil, zero value otherwise.

### GetBasePriceOk

`func (o *CloudBaseSku) GetBasePriceOk() (*float32, bool)`

GetBasePriceOk returns a tuple with the BasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrice

`func (o *CloudBaseSku) SetBasePrice(v float32)`

SetBasePrice sets BasePrice field to given value.

### HasBasePrice

`func (o *CloudBaseSku) HasBasePrice() bool`

HasBasePrice returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudBaseSku) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudBaseSku) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudBaseSku) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudBaseSku) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *CloudBaseSku) GetCustomAttributes() []CloudCustomAttributes`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *CloudBaseSku) GetCustomAttributesOk() (*[]CloudCustomAttributes, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *CloudBaseSku) SetCustomAttributes(v []CloudCustomAttributes)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *CloudBaseSku) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *CloudBaseSku) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *CloudBaseSku) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetDescription

`func (o *CloudBaseSku) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudBaseSku) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudBaseSku) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudBaseSku) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *CloudBaseSku) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CloudBaseSku) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CloudBaseSku) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CloudBaseSku) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsAutoDiscovered

`func (o *CloudBaseSku) GetIsAutoDiscovered() bool`

GetIsAutoDiscovered returns the IsAutoDiscovered field if non-nil, zero value otherwise.

### GetIsAutoDiscoveredOk

`func (o *CloudBaseSku) GetIsAutoDiscoveredOk() (*bool, bool)`

GetIsAutoDiscoveredOk returns a tuple with the IsAutoDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDiscovered

`func (o *CloudBaseSku) SetIsAutoDiscovered(v bool)`

SetIsAutoDiscovered sets IsAutoDiscovered field to given value.

### HasIsAutoDiscovered

`func (o *CloudBaseSku) HasIsAutoDiscovered() bool`

HasIsAutoDiscovered returns a boolean if a field has been set.

### GetName

`func (o *CloudBaseSku) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBaseSku) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBaseSku) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBaseSku) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformType

`func (o *CloudBaseSku) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *CloudBaseSku) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *CloudBaseSku) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *CloudBaseSku) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudBaseSku) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudBaseSku) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudBaseSku) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *CloudBaseSku) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetServiceCategory

`func (o *CloudBaseSku) GetServiceCategory() string`

GetServiceCategory returns the ServiceCategory field if non-nil, zero value otherwise.

### GetServiceCategoryOk

`func (o *CloudBaseSku) GetServiceCategoryOk() (*string, bool)`

GetServiceCategoryOk returns a tuple with the ServiceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCategory

`func (o *CloudBaseSku) SetServiceCategory(v string)`

SetServiceCategory sets ServiceCategory field to given value.

### HasServiceCategory

`func (o *CloudBaseSku) HasServiceCategory() bool`

HasServiceCategory returns a boolean if a field has been set.

### GetServiceFamily

`func (o *CloudBaseSku) GetServiceFamily() string`

GetServiceFamily returns the ServiceFamily field if non-nil, zero value otherwise.

### GetServiceFamilyOk

`func (o *CloudBaseSku) GetServiceFamilyOk() (*string, bool)`

GetServiceFamilyOk returns a tuple with the ServiceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFamily

`func (o *CloudBaseSku) SetServiceFamily(v string)`

SetServiceFamily sets ServiceFamily field to given value.

### HasServiceFamily

`func (o *CloudBaseSku) HasServiceFamily() bool`

HasServiceFamily returns a boolean if a field has been set.

### GetServiceName

`func (o *CloudBaseSku) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CloudBaseSku) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CloudBaseSku) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *CloudBaseSku) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetUnit

`func (o *CloudBaseSku) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CloudBaseSku) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CloudBaseSku) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CloudBaseSku) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetRegion

`func (o *CloudBaseSku) GetRegion() CloudRegionsRelationship`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudBaseSku) GetRegionOk() (*CloudRegionsRelationship, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudBaseSku) SetRegion(v CloudRegionsRelationship)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudBaseSku) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CloudBaseSku) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CloudBaseSku) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetTarget

`func (o *CloudBaseSku) GetTarget() AssetTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudBaseSku) GetTargetOk() (*AssetTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudBaseSku) SetTarget(v AssetTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudBaseSku) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CloudBaseSku) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CloudBaseSku) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


