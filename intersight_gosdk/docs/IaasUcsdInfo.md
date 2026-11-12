# IaasUcsdInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.UcsdInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.UcsdInfo"]
**DeviceConnectionStatus** | Pointer to **string** | The UCS Director device connection status. Possible values are Connected, Disconnected, Unknown. * &#x60;&#x60; - The target details have been persisted but Intersight has not yet attempted to connect to the target. * &#x60;Connected&#x60; - Intersight is able to establish a connection to the target and initiate management activities. * &#x60;NotConnected&#x60; - Intersight is unable to establish a connection to the target. * &#x60;ClaimInProgress&#x60; - Claim of the target is in progress. A connection to the target has not been fully established. * &#x60;UnclaimInProgress&#x60; - Unclaim of the target is in progress. Intersight is able to connect to the target and all management operations are supported. * &#x60;Unclaimed&#x60; - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * &#x60;Claimed&#x60; - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect. * &#x60;PreClaimed&#x60; - Target has been pre-claimed and is pending claim completion. Pre-claimed targets are pre-provisioned before endpoint is available to connect to Intersight, once target connects the claim can be completed by providing the security token to the pre-claim API. | [optional] [readonly] [default to ""]
**DeviceId** | Pointer to **string** | Moid of the UCS Director device connector&#39;s asset.DeviceRegistration. | [optional] [readonly] 
**Guid** | Pointer to **string** | Unique ID of UCS Director getting registerd with Intersight. | [optional] [readonly] 
**HostName** | Pointer to **string** | The UCS Director hostname for management. | [optional] [readonly] 
**Ip** | Pointer to **string** | The UCS Director IP address for management. | [optional] [readonly] 
**LastBackup** | Pointer to **time.Time** | Last successful backup created for this UCS Director appliance if backup is configured. | [optional] [readonly] 
**Model** | Pointer to **string** | Model of the UCS Director appliance. | [optional] [readonly] 
**NodeType** | Pointer to **string** | NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node. | [optional] [readonly] 
**Pid** | Pointer to **string** | Product ID (PID) of the UCS Director appliance. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | UCS Director platform or deployment type reported by the appliance. * &#x60;&#x60; - An unrecognized platform type. * &#x60;APIC&#x60; - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * &#x60;CAPIC&#x60; - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * &#x60;DCNM&#x60; - A Cisco Data Center Network Manager (DCNM) instance. * &#x60;UCSFI&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * &#x60;IMC&#x60; - A standalone Cisco UCS rack server (Deprecated). * &#x60;IMCM4&#x60; - A standalone Cisco UCS C-Series or S-Series M4 server. * &#x60;IMCM5&#x60; - A standalone Cisco UCS C-Series or S-Series M5 server. * &#x60;IMCRack&#x60; - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * &#x60;UCSIOM&#x60; - A Cisco UCS Blade Chassis I/O Module (IOM). * &#x60;UCSPCIeNODE&#x60; - A Cisco UCS PCIe Node in a blade form factor. * &#x60;UCSXFM&#x60; - A Cisco UCS Fabric Extender Module (XFM). * &#x60;HX&#x60; - A Cisco HyperFlex (HX) cluster. * &#x60;UCSD&#x60; - A Cisco UCS Director (UCSD) instance. * &#x60;UCSXECMC&#x60; - A Cisco UCSXE Chassis Management Controller. * &#x60;UCSXESecureRouter&#x60; - A secure networking router node designed to operate as part of a UCSXE chassis. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance instance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist instance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * &#x60;PureStorageFlashBladeSystem&#x60; - A Pure Storage FlashBlade target that is managed using Cisco Intersight Assist. Storage management feature is supported on this device. * &#x60;NexusDevice&#x60; - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * &#x60;ACISwitch&#x60; - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * &#x60;NexusSwitch&#x60; - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * &#x60;MDSSwitch&#x60; - A Cisco MDS Switch that is managed using the embedded Device Connector. * &#x60;MDSDevice&#x60; - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * &#x60;UCSC885&#x60; - A standalone Cisco UCS C885 Server. * &#x60;CAI845A&#x60; - A standalone Cisco AI 845A Server. * &#x60;UCSC890&#x60; - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * &#x60;NetAppOntap&#x60; - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVmax&#x60; - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVplex&#x60; - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;VmwareVcenter&#x60; - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * &#x60;AppDynamics&#x60; - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * &#x60;Dynatrace&#x60; - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;NewRelic&#x60; - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;Umbrella&#x60; - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;MySqlServer&#x60; - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;OracleDatabaseServer&#x60; - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;IBMWebSphereApplicationServer&#x60; - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;JavaVirtualMachine&#x60; - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * &#x60;AmazonWebService&#x60; - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatform&#x60; - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatformBilling&#x60; - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureBilling&#x60; - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;DellCompellent&#x60; - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;HPE3Par&#x60; - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * &#x60;NutanixPrismCentral&#x60; - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * &#x60;HPEOneView&#x60; - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * &#x60;GenericTarget&#x60; - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * &#x60;IMCBlade&#x60; - A Cisco UCS blade server managed by Cisco Intersight. * &#x60;TerraformCloud&#x60; - A Terraform Cloud Business Tier account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * &#x60;CustomTarget&#x60; - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * &#x60;AnsibleEndpoint&#x60; - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * &#x60;HTTPEndpoint&#x60; - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * &#x60;SSHEndpoint&#x60; - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * &#x60;CiscoDNAC&#x60; - A Cisco Digital Network Architecture (DNA) Center appliance. * &#x60;CiscoFMC&#x60; - A Cisco Secure Firewall Management Center. * &#x60;ViptelaCloud&#x60; - ViptelaCloud is deprecated. Use CatalystSDWANManager type to claim Catalyst SDWAN Manager which can manage Catalyst SDWAN devices. * &#x60;CatalystSDWANManager&#x60; - A Cisco Catalyst SD-WAN Manager. * &#x60;MerakiCloud&#x60; - A Cisco Meraki Organization. * &#x60;CiscoISE&#x60; - A Cisco Identity Services Engine (ISE) target. | [optional] [readonly] [default to ""]
**ProductName** | Pointer to **string** | The UCS Director product name. | [optional] [readonly] 
**ProductVendor** | Pointer to **string** | The UCS Director product vendor. | [optional] [readonly] 
**ProductVersion** | Pointer to **string** | The UCS Director product/platform version. | [optional] [readonly] 
**Serial** | Pointer to **string** | Serial number of the UCS Director appliance. | [optional] [readonly] 
**Status** | Pointer to **string** | The UCS Director status. Possible values are Active, Inactive, Unknown. | [optional] [readonly] 
**ConnectorPack** | Pointer to [**[]IaasConnectorPackRelationship**](IaasConnectorPackRelationship.md) | An array of relationships to iaasConnectorPack resources. | [optional] [readonly] 
**CustomTaskInfo** | Pointer to [**[]IaasCustomTaskInfoRelationship**](IaasCustomTaskInfoRelationship.md) | An array of relationships to iaasCustomTaskInfo resources. | [optional] [readonly] 
**DeviceStatus** | Pointer to [**[]IaasDeviceStatusRelationship**](IaasDeviceStatusRelationship.md) | An array of relationships to iaasDeviceStatus resources. | [optional] [readonly] 
**LicenseInfo** | Pointer to [**NullableIaasLicenseInfoRelationship**](IaasLicenseInfoRelationship.md) |  | [optional] 
**MostRunTasks** | Pointer to [**[]IaasMostRunTasksRelationship**](IaasMostRunTasksRelationship.md) | An array of relationships to iaasMostRunTasks resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SystemTaskInfo** | Pointer to [**[]IaasSystemTaskInfoRelationship**](IaasSystemTaskInfoRelationship.md) | An array of relationships to iaasSystemTaskInfo resources. | [optional] [readonly] 
**UcsdManagedInfra** | Pointer to [**NullableIaasUcsdManagedInfraRelationship**](IaasUcsdManagedInfraRelationship.md) |  | [optional] 

## Methods

### NewIaasUcsdInfo

`func NewIaasUcsdInfo(classId string, objectType string, ) *IaasUcsdInfo`

NewIaasUcsdInfo instantiates a new IaasUcsdInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasUcsdInfoWithDefaults

`func NewIaasUcsdInfoWithDefaults() *IaasUcsdInfo`

NewIaasUcsdInfoWithDefaults instantiates a new IaasUcsdInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasUcsdInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasUcsdInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasUcsdInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasUcsdInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasUcsdInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasUcsdInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceConnectionStatus

`func (o *IaasUcsdInfo) GetDeviceConnectionStatus() string`

GetDeviceConnectionStatus returns the DeviceConnectionStatus field if non-nil, zero value otherwise.

### GetDeviceConnectionStatusOk

`func (o *IaasUcsdInfo) GetDeviceConnectionStatusOk() (*string, bool)`

GetDeviceConnectionStatusOk returns a tuple with the DeviceConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConnectionStatus

`func (o *IaasUcsdInfo) SetDeviceConnectionStatus(v string)`

SetDeviceConnectionStatus sets DeviceConnectionStatus field to given value.

### HasDeviceConnectionStatus

`func (o *IaasUcsdInfo) HasDeviceConnectionStatus() bool`

HasDeviceConnectionStatus returns a boolean if a field has been set.

### GetDeviceId

`func (o *IaasUcsdInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IaasUcsdInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IaasUcsdInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IaasUcsdInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGuid

`func (o *IaasUcsdInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasUcsdInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasUcsdInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasUcsdInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHostName

`func (o *IaasUcsdInfo) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *IaasUcsdInfo) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *IaasUcsdInfo) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *IaasUcsdInfo) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIp

`func (o *IaasUcsdInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IaasUcsdInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IaasUcsdInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IaasUcsdInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastBackup

`func (o *IaasUcsdInfo) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *IaasUcsdInfo) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *IaasUcsdInfo) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.

### HasLastBackup

`func (o *IaasUcsdInfo) HasLastBackup() bool`

HasLastBackup returns a boolean if a field has been set.

### GetModel

`func (o *IaasUcsdInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IaasUcsdInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IaasUcsdInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IaasUcsdInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNodeType

`func (o *IaasUcsdInfo) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *IaasUcsdInfo) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *IaasUcsdInfo) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *IaasUcsdInfo) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetPid

`func (o *IaasUcsdInfo) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *IaasUcsdInfo) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *IaasUcsdInfo) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *IaasUcsdInfo) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformType

`func (o *IaasUcsdInfo) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *IaasUcsdInfo) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *IaasUcsdInfo) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *IaasUcsdInfo) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetProductName

`func (o *IaasUcsdInfo) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *IaasUcsdInfo) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *IaasUcsdInfo) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *IaasUcsdInfo) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductVendor

`func (o *IaasUcsdInfo) GetProductVendor() string`

GetProductVendor returns the ProductVendor field if non-nil, zero value otherwise.

### GetProductVendorOk

`func (o *IaasUcsdInfo) GetProductVendorOk() (*string, bool)`

GetProductVendorOk returns a tuple with the ProductVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVendor

`func (o *IaasUcsdInfo) SetProductVendor(v string)`

SetProductVendor sets ProductVendor field to given value.

### HasProductVendor

`func (o *IaasUcsdInfo) HasProductVendor() bool`

HasProductVendor returns a boolean if a field has been set.

### GetProductVersion

`func (o *IaasUcsdInfo) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *IaasUcsdInfo) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *IaasUcsdInfo) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *IaasUcsdInfo) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetSerial

`func (o *IaasUcsdInfo) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *IaasUcsdInfo) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *IaasUcsdInfo) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *IaasUcsdInfo) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *IaasUcsdInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IaasUcsdInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IaasUcsdInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IaasUcsdInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnectorPack

`func (o *IaasUcsdInfo) GetConnectorPack() []IaasConnectorPackRelationship`

GetConnectorPack returns the ConnectorPack field if non-nil, zero value otherwise.

### GetConnectorPackOk

`func (o *IaasUcsdInfo) GetConnectorPackOk() (*[]IaasConnectorPackRelationship, bool)`

GetConnectorPackOk returns a tuple with the ConnectorPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPack

`func (o *IaasUcsdInfo) SetConnectorPack(v []IaasConnectorPackRelationship)`

SetConnectorPack sets ConnectorPack field to given value.

### HasConnectorPack

`func (o *IaasUcsdInfo) HasConnectorPack() bool`

HasConnectorPack returns a boolean if a field has been set.

### SetConnectorPackNil

`func (o *IaasUcsdInfo) SetConnectorPackNil(b bool)`

 SetConnectorPackNil sets the value for ConnectorPack to be an explicit nil

### UnsetConnectorPack
`func (o *IaasUcsdInfo) UnsetConnectorPack()`

UnsetConnectorPack ensures that no value is present for ConnectorPack, not even an explicit nil
### GetCustomTaskInfo

`func (o *IaasUcsdInfo) GetCustomTaskInfo() []IaasCustomTaskInfoRelationship`

GetCustomTaskInfo returns the CustomTaskInfo field if non-nil, zero value otherwise.

### GetCustomTaskInfoOk

`func (o *IaasUcsdInfo) GetCustomTaskInfoOk() (*[]IaasCustomTaskInfoRelationship, bool)`

GetCustomTaskInfoOk returns a tuple with the CustomTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskInfo

`func (o *IaasUcsdInfo) SetCustomTaskInfo(v []IaasCustomTaskInfoRelationship)`

SetCustomTaskInfo sets CustomTaskInfo field to given value.

### HasCustomTaskInfo

`func (o *IaasUcsdInfo) HasCustomTaskInfo() bool`

HasCustomTaskInfo returns a boolean if a field has been set.

### SetCustomTaskInfoNil

`func (o *IaasUcsdInfo) SetCustomTaskInfoNil(b bool)`

 SetCustomTaskInfoNil sets the value for CustomTaskInfo to be an explicit nil

### UnsetCustomTaskInfo
`func (o *IaasUcsdInfo) UnsetCustomTaskInfo()`

UnsetCustomTaskInfo ensures that no value is present for CustomTaskInfo, not even an explicit nil
### GetDeviceStatus

`func (o *IaasUcsdInfo) GetDeviceStatus() []IaasDeviceStatusRelationship`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *IaasUcsdInfo) GetDeviceStatusOk() (*[]IaasDeviceStatusRelationship, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *IaasUcsdInfo) SetDeviceStatus(v []IaasDeviceStatusRelationship)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *IaasUcsdInfo) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### SetDeviceStatusNil

`func (o *IaasUcsdInfo) SetDeviceStatusNil(b bool)`

 SetDeviceStatusNil sets the value for DeviceStatus to be an explicit nil

### UnsetDeviceStatus
`func (o *IaasUcsdInfo) UnsetDeviceStatus()`

UnsetDeviceStatus ensures that no value is present for DeviceStatus, not even an explicit nil
### GetLicenseInfo

`func (o *IaasUcsdInfo) GetLicenseInfo() IaasLicenseInfoRelationship`

GetLicenseInfo returns the LicenseInfo field if non-nil, zero value otherwise.

### GetLicenseInfoOk

`func (o *IaasUcsdInfo) GetLicenseInfoOk() (*IaasLicenseInfoRelationship, bool)`

GetLicenseInfoOk returns a tuple with the LicenseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInfo

`func (o *IaasUcsdInfo) SetLicenseInfo(v IaasLicenseInfoRelationship)`

SetLicenseInfo sets LicenseInfo field to given value.

### HasLicenseInfo

`func (o *IaasUcsdInfo) HasLicenseInfo() bool`

HasLicenseInfo returns a boolean if a field has been set.

### SetLicenseInfoNil

`func (o *IaasUcsdInfo) SetLicenseInfoNil(b bool)`

 SetLicenseInfoNil sets the value for LicenseInfo to be an explicit nil

### UnsetLicenseInfo
`func (o *IaasUcsdInfo) UnsetLicenseInfo()`

UnsetLicenseInfo ensures that no value is present for LicenseInfo, not even an explicit nil
### GetMostRunTasks

`func (o *IaasUcsdInfo) GetMostRunTasks() []IaasMostRunTasksRelationship`

GetMostRunTasks returns the MostRunTasks field if non-nil, zero value otherwise.

### GetMostRunTasksOk

`func (o *IaasUcsdInfo) GetMostRunTasksOk() (*[]IaasMostRunTasksRelationship, bool)`

GetMostRunTasksOk returns a tuple with the MostRunTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRunTasks

`func (o *IaasUcsdInfo) SetMostRunTasks(v []IaasMostRunTasksRelationship)`

SetMostRunTasks sets MostRunTasks field to given value.

### HasMostRunTasks

`func (o *IaasUcsdInfo) HasMostRunTasks() bool`

HasMostRunTasks returns a boolean if a field has been set.

### SetMostRunTasksNil

`func (o *IaasUcsdInfo) SetMostRunTasksNil(b bool)`

 SetMostRunTasksNil sets the value for MostRunTasks to be an explicit nil

### UnsetMostRunTasks
`func (o *IaasUcsdInfo) UnsetMostRunTasks()`

UnsetMostRunTasks ensures that no value is present for MostRunTasks, not even an explicit nil
### GetRegisteredDevice

`func (o *IaasUcsdInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasUcsdInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasUcsdInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasUcsdInfo) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *IaasUcsdInfo) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *IaasUcsdInfo) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSystemTaskInfo

`func (o *IaasUcsdInfo) GetSystemTaskInfo() []IaasSystemTaskInfoRelationship`

GetSystemTaskInfo returns the SystemTaskInfo field if non-nil, zero value otherwise.

### GetSystemTaskInfoOk

`func (o *IaasUcsdInfo) GetSystemTaskInfoOk() (*[]IaasSystemTaskInfoRelationship, bool)`

GetSystemTaskInfoOk returns a tuple with the SystemTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTaskInfo

`func (o *IaasUcsdInfo) SetSystemTaskInfo(v []IaasSystemTaskInfoRelationship)`

SetSystemTaskInfo sets SystemTaskInfo field to given value.

### HasSystemTaskInfo

`func (o *IaasUcsdInfo) HasSystemTaskInfo() bool`

HasSystemTaskInfo returns a boolean if a field has been set.

### SetSystemTaskInfoNil

`func (o *IaasUcsdInfo) SetSystemTaskInfoNil(b bool)`

 SetSystemTaskInfoNil sets the value for SystemTaskInfo to be an explicit nil

### UnsetSystemTaskInfo
`func (o *IaasUcsdInfo) UnsetSystemTaskInfo()`

UnsetSystemTaskInfo ensures that no value is present for SystemTaskInfo, not even an explicit nil
### GetUcsdManagedInfra

`func (o *IaasUcsdInfo) GetUcsdManagedInfra() IaasUcsdManagedInfraRelationship`

GetUcsdManagedInfra returns the UcsdManagedInfra field if non-nil, zero value otherwise.

### GetUcsdManagedInfraOk

`func (o *IaasUcsdInfo) GetUcsdManagedInfraOk() (*IaasUcsdManagedInfraRelationship, bool)`

GetUcsdManagedInfraOk returns a tuple with the UcsdManagedInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsdManagedInfra

`func (o *IaasUcsdInfo) SetUcsdManagedInfra(v IaasUcsdManagedInfraRelationship)`

SetUcsdManagedInfra sets UcsdManagedInfra field to given value.

### HasUcsdManagedInfra

`func (o *IaasUcsdInfo) HasUcsdManagedInfra() bool`

HasUcsdManagedInfra returns a boolean if a field has been set.

### SetUcsdManagedInfraNil

`func (o *IaasUcsdInfo) SetUcsdManagedInfraNil(b bool)`

 SetUcsdManagedInfraNil sets the value for UcsdManagedInfra to be an explicit nil

### UnsetUcsdManagedInfra
`func (o *IaasUcsdInfo) UnsetUcsdManagedInfra()`

UnsetUcsdManagedInfra ensures that no value is present for UcsdManagedInfra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


