# AssetDeviceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceRegistration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceRegistration"]
**AccessKeyId** | Pointer to **string** | An identifier for the credential used by the device connector to authenticate with the Intersight web socket gateway. | [optional] 
**ClaimedByUserName** | Pointer to **string** | The name of the user who claimed the device for the account. | [optional] [readonly] 
**ClaimedTime** | Pointer to **time.Time** | The date and time at which the device was claimed to this account. | [optional] [readonly] 
**DeviceHostname** | Pointer to **[]string** |  | [optional] 
**DeviceIpAddress** | Pointer to **[]string** |  | [optional] 
**ExecutionMode** | Pointer to **string** | Indicates if the platform is an actual device or an emulated device for testing, demos, etc. Permitted values are [Normal, Emulator, ContainerEmulator]. * &#x60;&#x60; - The device reported an empty or unrecognized executionMode. * &#x60;Normal&#x60; - The device connector is running in normal mode, i.e. it is not a simulation. * &#x60;Emulator&#x60; - The device connector is running in simulation mode inside an emulated device. * &#x60;ContainerEmulator&#x60; - The device connector is running in simulation mode inside a containerized emulated device. | [optional] [default to ""]
**ParentSignature** | Pointer to [**NullableAssetClaimSignature**](AssetClaimSignature.md) |  | [optional] 
**Pid** | Pointer to **[]string** |  | [optional] 
**PlatformType** | Pointer to **string** | The platform type on which device connector is executing. * &#x60;&#x60; - An unrecognized platform type. * &#x60;APIC&#x60; - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * &#x60;CAPIC&#x60; - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * &#x60;DCNM&#x60; - A Cisco Data Center Network Manager (DCNM) instance. * &#x60;UCSFI&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * &#x60;IMC&#x60; - A standalone Cisco UCS rack server (Deprecated). * &#x60;IMCM4&#x60; - A standalone Cisco UCS C-Series or S-Series M4 server. * &#x60;IMCM5&#x60; - A standalone Cisco UCS C-Series or S-Series M5 server. * &#x60;IMCRack&#x60; - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * &#x60;UCSIOM&#x60; - A Cisco UCS Blade Chassis I/O Module (IOM). * &#x60;HX&#x60; - A Cisco HyperFlex (HX) cluster. * &#x60;UCSD&#x60; - A Cisco UCS Director (UCSD) instance. * &#x60;UCSXECMC&#x60; - A Cisco UCSXE Chassis Management Controller. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance instance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist instance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * &#x60;NexusDevice&#x60; - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * &#x60;ACISwitch&#x60; - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * &#x60;NexusSwitch&#x60; - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * &#x60;MDSSwitch&#x60; - A Cisco MDS Switch that is managed using the embedded Device Connector. * &#x60;MDSDevice&#x60; - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * &#x60;UCSC885&#x60; - A standalone Cisco UCS C885 Server. * &#x60;CAI845A&#x60; - A standalone Cisco AI 845A Server. * &#x60;UCSC890&#x60; - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * &#x60;NetAppOntap&#x60; - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVmax&#x60; - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVplex&#x60; - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;VmwareVcenter&#x60; - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * &#x60;AppDynamics&#x60; - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * &#x60;Dynatrace&#x60; - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;NewRelic&#x60; - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;Umbrella&#x60; - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;MySqlServer&#x60; - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;OracleDatabaseServer&#x60; - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;IBMWebSphereApplicationServer&#x60; - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;JavaVirtualMachine&#x60; - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * &#x60;AmazonWebService&#x60; - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatform&#x60; - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatformBilling&#x60; - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureBilling&#x60; - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;DellCompellent&#x60; - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;HPE3Par&#x60; - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * &#x60;NutanixPrismCentral&#x60; - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * &#x60;HPEOneView&#x60; - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * &#x60;GenericTarget&#x60; - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * &#x60;IMCBlade&#x60; - A Cisco UCS blade server managed by Cisco Intersight. * &#x60;TerraformCloud&#x60; - A Terraform Cloud Business Tier account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * &#x60;CustomTarget&#x60; - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * &#x60;AnsibleEndpoint&#x60; - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * &#x60;HTTPEndpoint&#x60; - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * &#x60;SSHEndpoint&#x60; - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * &#x60;CiscoDNAC&#x60; - A Cisco Digital Network Architecture (DNA) Center appliance. * &#x60;CiscoFMC&#x60; - A Cisco Secure Firewall Management Center. * &#x60;ViptelaCloud&#x60; - A Cisco Viptela SD-WAN Cloud. * &#x60;MerakiCloud&#x60; - A Cisco Meraki Organization. * &#x60;CiscoISE&#x60; - A Cisco Identity Services Engine (ISE) target. | [optional] [default to ""]
**PublicAccessKey** | Pointer to **string** | The device connector&#39;s public key used by Intersight to authenticate a connection from the device connector. The public key is used to verify that the signature a device connector sends on connect has been signed by the connector&#39;s private key stored on the device&#39;s filesystem. Must be a PEM encoded RSA or Ed22519 public key string. | [optional] [readonly] 
**PublicAccessKeyRotated** | Pointer to **string** | The device connectors rotated public key. The device connector may rotate its key pair in conditions where it has been determined the key may have been compromised or if the key type is changing (e.g. moving from RSA to elliptical curve). The device connector will rotate its key in a two step process in order to ensure Intersight has received and committed the new key before invalidating the previous. Device connectors will first send the new public key on PublicAccessKeyRotated, the public key string will also be signed by the current private key. On receipt of a 200 response from Intersight the device connector commits the new key to its local database and will use it for all future authentication requests. Intersight will then allow both the previous and rotated key periods until the device connects with the rotated key, at which point the previous key is invalidated and removed. | [optional] [readonly] 
**PublicEncryptionKey** | Pointer to **string** | The device connector public key used by Intersight for encryption. The public key is used to encrypt ephemeral aes keys to be used for decrypting sensitive data from Intersight. Must be a PEM encoded RSA public key string. | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | Flag reported by devices to indicate an administrator of the device has disabled management operations of the device connector and only monitoring is permitted. | [optional] [readonly] 
**RotateAccessKey** | Pointer to **bool** | Request the device to rotate its key pair. SRE team may set this field to trigger the device to rotate its key pair in conditions where it has been identified that the device&#39;s key has been compromised. When RotateAccessKey is set to true the device will be forced to re-connect and rotate its key. | [optional] 
**Serial** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | The vendor of the managed device. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**ClaimedByUser** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 
**ClusterMembers** | Pointer to [**[]AssetClusterMemberRelationship**](AssetClusterMemberRelationship.md) | An array of relationships to assetClusterMember resources. | [optional] [readonly] 
**CustomPermissionResources** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DeviceClaim** | Pointer to [**NullableAssetDeviceClaimRelationship**](AssetDeviceClaimRelationship.md) |  | [optional] 
**DeviceConfiguration** | Pointer to [**NullableAssetDeviceConfigurationRelationship**](AssetDeviceConfigurationRelationship.md) |  | [optional] 
**DomainGroup** | Pointer to [**NullableIamDomainGroupRelationship**](IamDomainGroupRelationship.md) |  | [optional] 
**ParentConnection** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Target** | Pointer to [**NullableAssetTargetRelationship**](AssetTargetRelationship.md) |  | [optional] 

## Methods

### NewAssetDeviceRegistration

`func NewAssetDeviceRegistration(classId string, objectType string, ) *AssetDeviceRegistration`

NewAssetDeviceRegistration instantiates a new AssetDeviceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceRegistrationWithDefaults

`func NewAssetDeviceRegistrationWithDefaults() *AssetDeviceRegistration`

NewAssetDeviceRegistrationWithDefaults instantiates a new AssetDeviceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceRegistration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceRegistration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceRegistration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceRegistration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceRegistration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceRegistration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessKeyId

`func (o *AssetDeviceRegistration) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AssetDeviceRegistration) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AssetDeviceRegistration) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AssetDeviceRegistration) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetClaimedByUserName

`func (o *AssetDeviceRegistration) GetClaimedByUserName() string`

GetClaimedByUserName returns the ClaimedByUserName field if non-nil, zero value otherwise.

### GetClaimedByUserNameOk

`func (o *AssetDeviceRegistration) GetClaimedByUserNameOk() (*string, bool)`

GetClaimedByUserNameOk returns a tuple with the ClaimedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedByUserName

`func (o *AssetDeviceRegistration) SetClaimedByUserName(v string)`

SetClaimedByUserName sets ClaimedByUserName field to given value.

### HasClaimedByUserName

`func (o *AssetDeviceRegistration) HasClaimedByUserName() bool`

HasClaimedByUserName returns a boolean if a field has been set.

### GetClaimedTime

`func (o *AssetDeviceRegistration) GetClaimedTime() time.Time`

GetClaimedTime returns the ClaimedTime field if non-nil, zero value otherwise.

### GetClaimedTimeOk

`func (o *AssetDeviceRegistration) GetClaimedTimeOk() (*time.Time, bool)`

GetClaimedTimeOk returns a tuple with the ClaimedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedTime

`func (o *AssetDeviceRegistration) SetClaimedTime(v time.Time)`

SetClaimedTime sets ClaimedTime field to given value.

### HasClaimedTime

`func (o *AssetDeviceRegistration) HasClaimedTime() bool`

HasClaimedTime returns a boolean if a field has been set.

### GetDeviceHostname

`func (o *AssetDeviceRegistration) GetDeviceHostname() []string`

GetDeviceHostname returns the DeviceHostname field if non-nil, zero value otherwise.

### GetDeviceHostnameOk

`func (o *AssetDeviceRegistration) GetDeviceHostnameOk() (*[]string, bool)`

GetDeviceHostnameOk returns a tuple with the DeviceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHostname

`func (o *AssetDeviceRegistration) SetDeviceHostname(v []string)`

SetDeviceHostname sets DeviceHostname field to given value.

### HasDeviceHostname

`func (o *AssetDeviceRegistration) HasDeviceHostname() bool`

HasDeviceHostname returns a boolean if a field has been set.

### SetDeviceHostnameNil

`func (o *AssetDeviceRegistration) SetDeviceHostnameNil(b bool)`

 SetDeviceHostnameNil sets the value for DeviceHostname to be an explicit nil

### UnsetDeviceHostname
`func (o *AssetDeviceRegistration) UnsetDeviceHostname()`

UnsetDeviceHostname ensures that no value is present for DeviceHostname, not even an explicit nil
### GetDeviceIpAddress

`func (o *AssetDeviceRegistration) GetDeviceIpAddress() []string`

GetDeviceIpAddress returns the DeviceIpAddress field if non-nil, zero value otherwise.

### GetDeviceIpAddressOk

`func (o *AssetDeviceRegistration) GetDeviceIpAddressOk() (*[]string, bool)`

GetDeviceIpAddressOk returns a tuple with the DeviceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIpAddress

`func (o *AssetDeviceRegistration) SetDeviceIpAddress(v []string)`

SetDeviceIpAddress sets DeviceIpAddress field to given value.

### HasDeviceIpAddress

`func (o *AssetDeviceRegistration) HasDeviceIpAddress() bool`

HasDeviceIpAddress returns a boolean if a field has been set.

### SetDeviceIpAddressNil

`func (o *AssetDeviceRegistration) SetDeviceIpAddressNil(b bool)`

 SetDeviceIpAddressNil sets the value for DeviceIpAddress to be an explicit nil

### UnsetDeviceIpAddress
`func (o *AssetDeviceRegistration) UnsetDeviceIpAddress()`

UnsetDeviceIpAddress ensures that no value is present for DeviceIpAddress, not even an explicit nil
### GetExecutionMode

`func (o *AssetDeviceRegistration) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AssetDeviceRegistration) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AssetDeviceRegistration) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AssetDeviceRegistration) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetParentSignature

`func (o *AssetDeviceRegistration) GetParentSignature() AssetClaimSignature`

GetParentSignature returns the ParentSignature field if non-nil, zero value otherwise.

### GetParentSignatureOk

`func (o *AssetDeviceRegistration) GetParentSignatureOk() (*AssetClaimSignature, bool)`

GetParentSignatureOk returns a tuple with the ParentSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSignature

`func (o *AssetDeviceRegistration) SetParentSignature(v AssetClaimSignature)`

SetParentSignature sets ParentSignature field to given value.

### HasParentSignature

`func (o *AssetDeviceRegistration) HasParentSignature() bool`

HasParentSignature returns a boolean if a field has been set.

### SetParentSignatureNil

`func (o *AssetDeviceRegistration) SetParentSignatureNil(b bool)`

 SetParentSignatureNil sets the value for ParentSignature to be an explicit nil

### UnsetParentSignature
`func (o *AssetDeviceRegistration) UnsetParentSignature()`

UnsetParentSignature ensures that no value is present for ParentSignature, not even an explicit nil
### GetPid

`func (o *AssetDeviceRegistration) GetPid() []string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *AssetDeviceRegistration) GetPidOk() (*[]string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *AssetDeviceRegistration) SetPid(v []string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *AssetDeviceRegistration) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *AssetDeviceRegistration) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *AssetDeviceRegistration) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetPlatformType

`func (o *AssetDeviceRegistration) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AssetDeviceRegistration) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AssetDeviceRegistration) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AssetDeviceRegistration) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPublicAccessKey

`func (o *AssetDeviceRegistration) GetPublicAccessKey() string`

GetPublicAccessKey returns the PublicAccessKey field if non-nil, zero value otherwise.

### GetPublicAccessKeyOk

`func (o *AssetDeviceRegistration) GetPublicAccessKeyOk() (*string, bool)`

GetPublicAccessKeyOk returns a tuple with the PublicAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessKey

`func (o *AssetDeviceRegistration) SetPublicAccessKey(v string)`

SetPublicAccessKey sets PublicAccessKey field to given value.

### HasPublicAccessKey

`func (o *AssetDeviceRegistration) HasPublicAccessKey() bool`

HasPublicAccessKey returns a boolean if a field has been set.

### GetPublicAccessKeyRotated

`func (o *AssetDeviceRegistration) GetPublicAccessKeyRotated() string`

GetPublicAccessKeyRotated returns the PublicAccessKeyRotated field if non-nil, zero value otherwise.

### GetPublicAccessKeyRotatedOk

`func (o *AssetDeviceRegistration) GetPublicAccessKeyRotatedOk() (*string, bool)`

GetPublicAccessKeyRotatedOk returns a tuple with the PublicAccessKeyRotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessKeyRotated

`func (o *AssetDeviceRegistration) SetPublicAccessKeyRotated(v string)`

SetPublicAccessKeyRotated sets PublicAccessKeyRotated field to given value.

### HasPublicAccessKeyRotated

`func (o *AssetDeviceRegistration) HasPublicAccessKeyRotated() bool`

HasPublicAccessKeyRotated returns a boolean if a field has been set.

### GetPublicEncryptionKey

`func (o *AssetDeviceRegistration) GetPublicEncryptionKey() string`

GetPublicEncryptionKey returns the PublicEncryptionKey field if non-nil, zero value otherwise.

### GetPublicEncryptionKeyOk

`func (o *AssetDeviceRegistration) GetPublicEncryptionKeyOk() (*string, bool)`

GetPublicEncryptionKeyOk returns a tuple with the PublicEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEncryptionKey

`func (o *AssetDeviceRegistration) SetPublicEncryptionKey(v string)`

SetPublicEncryptionKey sets PublicEncryptionKey field to given value.

### HasPublicEncryptionKey

`func (o *AssetDeviceRegistration) HasPublicEncryptionKey() bool`

HasPublicEncryptionKey returns a boolean if a field has been set.

### GetReadOnly

`func (o *AssetDeviceRegistration) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AssetDeviceRegistration) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AssetDeviceRegistration) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AssetDeviceRegistration) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRotateAccessKey

`func (o *AssetDeviceRegistration) GetRotateAccessKey() bool`

GetRotateAccessKey returns the RotateAccessKey field if non-nil, zero value otherwise.

### GetRotateAccessKeyOk

`func (o *AssetDeviceRegistration) GetRotateAccessKeyOk() (*bool, bool)`

GetRotateAccessKeyOk returns a tuple with the RotateAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAccessKey

`func (o *AssetDeviceRegistration) SetRotateAccessKey(v bool)`

SetRotateAccessKey sets RotateAccessKey field to given value.

### HasRotateAccessKey

`func (o *AssetDeviceRegistration) HasRotateAccessKey() bool`

HasRotateAccessKey returns a boolean if a field has been set.

### GetSerial

`func (o *AssetDeviceRegistration) GetSerial() []string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AssetDeviceRegistration) GetSerialOk() (*[]string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AssetDeviceRegistration) SetSerial(v []string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AssetDeviceRegistration) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *AssetDeviceRegistration) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *AssetDeviceRegistration) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetVendor

`func (o *AssetDeviceRegistration) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AssetDeviceRegistration) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AssetDeviceRegistration) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AssetDeviceRegistration) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAccount

`func (o *AssetDeviceRegistration) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetDeviceRegistration) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetDeviceRegistration) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetDeviceRegistration) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *AssetDeviceRegistration) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AssetDeviceRegistration) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetClaimedByUser

`func (o *AssetDeviceRegistration) GetClaimedByUser() IamUserRelationship`

GetClaimedByUser returns the ClaimedByUser field if non-nil, zero value otherwise.

### GetClaimedByUserOk

`func (o *AssetDeviceRegistration) GetClaimedByUserOk() (*IamUserRelationship, bool)`

GetClaimedByUserOk returns a tuple with the ClaimedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedByUser

`func (o *AssetDeviceRegistration) SetClaimedByUser(v IamUserRelationship)`

SetClaimedByUser sets ClaimedByUser field to given value.

### HasClaimedByUser

`func (o *AssetDeviceRegistration) HasClaimedByUser() bool`

HasClaimedByUser returns a boolean if a field has been set.

### SetClaimedByUserNil

`func (o *AssetDeviceRegistration) SetClaimedByUserNil(b bool)`

 SetClaimedByUserNil sets the value for ClaimedByUser to be an explicit nil

### UnsetClaimedByUser
`func (o *AssetDeviceRegistration) UnsetClaimedByUser()`

UnsetClaimedByUser ensures that no value is present for ClaimedByUser, not even an explicit nil
### GetClusterMembers

`func (o *AssetDeviceRegistration) GetClusterMembers() []AssetClusterMemberRelationship`

GetClusterMembers returns the ClusterMembers field if non-nil, zero value otherwise.

### GetClusterMembersOk

`func (o *AssetDeviceRegistration) GetClusterMembersOk() (*[]AssetClusterMemberRelationship, bool)`

GetClusterMembersOk returns a tuple with the ClusterMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMembers

`func (o *AssetDeviceRegistration) SetClusterMembers(v []AssetClusterMemberRelationship)`

SetClusterMembers sets ClusterMembers field to given value.

### HasClusterMembers

`func (o *AssetDeviceRegistration) HasClusterMembers() bool`

HasClusterMembers returns a boolean if a field has been set.

### SetClusterMembersNil

`func (o *AssetDeviceRegistration) SetClusterMembersNil(b bool)`

 SetClusterMembersNil sets the value for ClusterMembers to be an explicit nil

### UnsetClusterMembers
`func (o *AssetDeviceRegistration) UnsetClusterMembers()`

UnsetClusterMembers ensures that no value is present for ClusterMembers, not even an explicit nil
### GetCustomPermissionResources

`func (o *AssetDeviceRegistration) GetCustomPermissionResources() []MoBaseMoRelationship`

GetCustomPermissionResources returns the CustomPermissionResources field if non-nil, zero value otherwise.

### GetCustomPermissionResourcesOk

`func (o *AssetDeviceRegistration) GetCustomPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetCustomPermissionResourcesOk returns a tuple with the CustomPermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPermissionResources

`func (o *AssetDeviceRegistration) SetCustomPermissionResources(v []MoBaseMoRelationship)`

SetCustomPermissionResources sets CustomPermissionResources field to given value.

### HasCustomPermissionResources

`func (o *AssetDeviceRegistration) HasCustomPermissionResources() bool`

HasCustomPermissionResources returns a boolean if a field has been set.

### SetCustomPermissionResourcesNil

`func (o *AssetDeviceRegistration) SetCustomPermissionResourcesNil(b bool)`

 SetCustomPermissionResourcesNil sets the value for CustomPermissionResources to be an explicit nil

### UnsetCustomPermissionResources
`func (o *AssetDeviceRegistration) UnsetCustomPermissionResources()`

UnsetCustomPermissionResources ensures that no value is present for CustomPermissionResources, not even an explicit nil
### GetDeviceClaim

`func (o *AssetDeviceRegistration) GetDeviceClaim() AssetDeviceClaimRelationship`

GetDeviceClaim returns the DeviceClaim field if non-nil, zero value otherwise.

### GetDeviceClaimOk

`func (o *AssetDeviceRegistration) GetDeviceClaimOk() (*AssetDeviceClaimRelationship, bool)`

GetDeviceClaimOk returns a tuple with the DeviceClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClaim

`func (o *AssetDeviceRegistration) SetDeviceClaim(v AssetDeviceClaimRelationship)`

SetDeviceClaim sets DeviceClaim field to given value.

### HasDeviceClaim

`func (o *AssetDeviceRegistration) HasDeviceClaim() bool`

HasDeviceClaim returns a boolean if a field has been set.

### SetDeviceClaimNil

`func (o *AssetDeviceRegistration) SetDeviceClaimNil(b bool)`

 SetDeviceClaimNil sets the value for DeviceClaim to be an explicit nil

### UnsetDeviceClaim
`func (o *AssetDeviceRegistration) UnsetDeviceClaim()`

UnsetDeviceClaim ensures that no value is present for DeviceClaim, not even an explicit nil
### GetDeviceConfiguration

`func (o *AssetDeviceRegistration) GetDeviceConfiguration() AssetDeviceConfigurationRelationship`

GetDeviceConfiguration returns the DeviceConfiguration field if non-nil, zero value otherwise.

### GetDeviceConfigurationOk

`func (o *AssetDeviceRegistration) GetDeviceConfigurationOk() (*AssetDeviceConfigurationRelationship, bool)`

GetDeviceConfigurationOk returns a tuple with the DeviceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfiguration

`func (o *AssetDeviceRegistration) SetDeviceConfiguration(v AssetDeviceConfigurationRelationship)`

SetDeviceConfiguration sets DeviceConfiguration field to given value.

### HasDeviceConfiguration

`func (o *AssetDeviceRegistration) HasDeviceConfiguration() bool`

HasDeviceConfiguration returns a boolean if a field has been set.

### SetDeviceConfigurationNil

`func (o *AssetDeviceRegistration) SetDeviceConfigurationNil(b bool)`

 SetDeviceConfigurationNil sets the value for DeviceConfiguration to be an explicit nil

### UnsetDeviceConfiguration
`func (o *AssetDeviceRegistration) UnsetDeviceConfiguration()`

UnsetDeviceConfiguration ensures that no value is present for DeviceConfiguration, not even an explicit nil
### GetDomainGroup

`func (o *AssetDeviceRegistration) GetDomainGroup() IamDomainGroupRelationship`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *AssetDeviceRegistration) GetDomainGroupOk() (*IamDomainGroupRelationship, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *AssetDeviceRegistration) SetDomainGroup(v IamDomainGroupRelationship)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *AssetDeviceRegistration) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### SetDomainGroupNil

`func (o *AssetDeviceRegistration) SetDomainGroupNil(b bool)`

 SetDomainGroupNil sets the value for DomainGroup to be an explicit nil

### UnsetDomainGroup
`func (o *AssetDeviceRegistration) UnsetDomainGroup()`

UnsetDomainGroup ensures that no value is present for DomainGroup, not even an explicit nil
### GetParentConnection

`func (o *AssetDeviceRegistration) GetParentConnection() AssetDeviceRegistrationRelationship`

GetParentConnection returns the ParentConnection field if non-nil, zero value otherwise.

### GetParentConnectionOk

`func (o *AssetDeviceRegistration) GetParentConnectionOk() (*AssetDeviceRegistrationRelationship, bool)`

GetParentConnectionOk returns a tuple with the ParentConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnection

`func (o *AssetDeviceRegistration) SetParentConnection(v AssetDeviceRegistrationRelationship)`

SetParentConnection sets ParentConnection field to given value.

### HasParentConnection

`func (o *AssetDeviceRegistration) HasParentConnection() bool`

HasParentConnection returns a boolean if a field has been set.

### SetParentConnectionNil

`func (o *AssetDeviceRegistration) SetParentConnectionNil(b bool)`

 SetParentConnectionNil sets the value for ParentConnection to be an explicit nil

### UnsetParentConnection
`func (o *AssetDeviceRegistration) UnsetParentConnection()`

UnsetParentConnection ensures that no value is present for ParentConnection, not even an explicit nil
### GetTarget

`func (o *AssetDeviceRegistration) GetTarget() AssetTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AssetDeviceRegistration) GetTargetOk() (*AssetTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AssetDeviceRegistration) SetTarget(v AssetTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AssetDeviceRegistration) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *AssetDeviceRegistration) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *AssetDeviceRegistration) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


