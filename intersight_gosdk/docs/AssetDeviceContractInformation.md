# AssetDeviceContractInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceContractInformation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceContractInformation"]
**Contract** | Pointer to [**NullableAssetContractInformation**](AssetContractInformation.md) |  | [optional] 
**ContractStatus** | Pointer to **string** | Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered. * &#x60;Unknown&#x60; - The device&#39;s contract status cannot be determined. * &#x60;Not Covered&#x60; - The Cisco device does not have a valid support contract. * &#x60;Active&#x60; - The Cisco device is covered under a active support contract. * &#x60;Expiring Soon&#x60; - The contract for this Cisco device is going to expire in the next 30 days. | [optional] [readonly] [default to "Unknown"]
**ContractStatusReason** | Pointer to **string** | Reason for contract status. In case of Not Covered, reason is either Terminated or Expired. * &#x60;&#x60; - There is no reason for the specified contract status. * &#x60;Line Item Expired&#x60; - The Cisco device does not have a valid support contract, it has expired. * &#x60;Line Item Terminated&#x60; - The Cisco device does not have a valid support contract, it has been terminated. | [optional] [readonly] [default to ""]
**ContractUnavailableRetryCount** | Pointer to **int64** | If Contract details doesn&#39;t exists in CCW API response, we will retry couple of times based on threshold value Retry count will track the number of retry happens due to contract details unavailability. | [optional] [readonly] 
**ContractUpdatedTime** | Pointer to **time.Time** | Date and time indicating when the contract data is last fetched from Cisco&#39;s Contract API successfully. | [optional] [readonly] 
**CoveredProductLineEndDate** | Pointer to **string** | End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API. | [optional] [readonly] 
**DeviceId** | Pointer to **string** | Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future. * &#x60;None&#x60; - A default value to catch cases where device type is not correctly detected. * &#x60;CiscoUcsServer&#x60; - A device of type server. It includes Cisco IMC and UCS Managed servers. * &#x60;CiscoUcsFI&#x60; - A device of type Fabric Interconnect. It includes the various types of Cisco Fabric Interconnects supported by Cisco Intersight. * &#x60;CiscoUcsChassis&#x60; - A device of type Chassis. It includes various UCS chassis supported by Cisco Intersight. * &#x60;CiscoNexusSwitch&#x60; - A device of type Nexus switch. It includes various Nexus switches supported by Cisco Intersight. * &#x60;CiscoMDSSwitch&#x60; - A device of type MDS switch. It includes various MDS switches supported by Cisco Intersight. | [optional] [readonly] [default to "None"]
**EndCustomer** | Pointer to [**NullableAssetCustomerInformation**](AssetCustomerInformation.md) |  | [optional] 
**EndUserGlobalUltimate** | Pointer to [**NullableAssetGlobalUltimate**](AssetGlobalUltimate.md) |  | [optional] 
**IsValid** | Pointer to **bool** | Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs. | [optional] [readonly] 
**ItemType** | Pointer to **string** | Item type of this specific Cisco device. example \&quot;Chassis\&quot;. | [optional] [readonly] 
**LastDateOfSupport** | Pointer to **time.Time** | The last date of hardware support for this device. | [optional] [readonly] 
**MaintenancePurchaseOrderNumber** | Pointer to **string** | Maintenance purchase order number for the Cisco device. | [optional] [readonly] 
**MaintenanceSalesOrderNumber** | Pointer to **string** | Maintenance sales order number for the Cisco device. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type of the Cisco device. * &#x60;&#x60; - An unrecognized platform type. * &#x60;APIC&#x60; - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * &#x60;CAPIC&#x60; - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * &#x60;DCNM&#x60; - A Cisco Data Center Network Manager (DCNM) instance. * &#x60;UCSFI&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * &#x60;IMC&#x60; - A standalone Cisco UCS rack server (Deprecated). * &#x60;IMCM4&#x60; - A standalone Cisco UCS C-Series or S-Series M4 server. * &#x60;IMCM5&#x60; - A standalone Cisco UCS C-Series or S-Series M5 server. * &#x60;IMCRack&#x60; - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * &#x60;UCSIOM&#x60; - A Cisco UCS Blade Chassis I/O Module (IOM). * &#x60;HX&#x60; - A Cisco HyperFlex (HX) cluster. * &#x60;UCSD&#x60; - A Cisco UCS Director (UCSD) instance. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance instance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist instance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * &#x60;NexusDevice&#x60; - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * &#x60;ACISwitch&#x60; - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * &#x60;NexusSwitch&#x60; - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * &#x60;MDSSwitch&#x60; - A Cisco MDS Switch that is managed using the embedded Device Connector. * &#x60;MDSDevice&#x60; - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * &#x60;UCSC890&#x60; - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * &#x60;NetAppOntap&#x60; - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVmax&#x60; - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVplex&#x60; - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;VmwareVcenter&#x60; - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * &#x60;AppDynamics&#x60; - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * &#x60;Dynatrace&#x60; - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;NewRelic&#x60; - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;Umbrella&#x60; - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;MySqlServer&#x60; - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;OracleDatabaseServer&#x60; - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;IBMWebSphereApplicationServer&#x60; - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;JavaVirtualMachine&#x60; - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * &#x60;AmazonWebService&#x60; - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatform&#x60; - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatformBilling&#x60; - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureBilling&#x60; - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;DellCompellent&#x60; - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;HPE3Par&#x60; - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * &#x60;HPEOneView&#x60; - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * &#x60;GenericTarget&#x60; - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * &#x60;IMCBlade&#x60; - A Cisco UCS blade server managed by Cisco Intersight. * &#x60;TerraformCloud&#x60; - A Terraform Cloud Business Tier account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * &#x60;CustomTarget&#x60; - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * &#x60;AnsibleEndpoint&#x60; - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * &#x60;HTTPEndpoint&#x60; - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * &#x60;SSHEndpoint&#x60; - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * &#x60;CiscoDNAC&#x60; - A Cisco Digital Network Architecture (DNA) Center appliance. * &#x60;CiscoFMC&#x60; - A Cisco Secure Firewall Management Center. * &#x60;ViptelaCloud&#x60; - A Cisco Viptela SD-WAN Cloud. * &#x60;MerakiCloud&#x60; - A Cisco Meraki Organization. * &#x60;CiscoISE&#x60; - A Cisco Identity Services Engine (ISE) target. | [optional] [readonly] [default to ""]
**Product** | Pointer to [**NullableAssetProductInformation**](AssetProductInformation.md) |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number for the Cisco device. It is a unique number assigned for every purchase. | [optional] [readonly] 
**ResellerGlobalUltimate** | Pointer to [**NullableAssetGlobalUltimate**](AssetGlobalUltimate.md) |  | [optional] 
**SalesOrderNumber** | Pointer to **string** | Sales order number for the Cisco device. It is a unique number assigned for every sale. | [optional] [readonly] 
**ServiceDescription** | Pointer to **string** | The type of service contract that covers the Cisco device. | [optional] [readonly] 
**ServiceEndDate** | Pointer to **time.Time** | End date for the Cisco service contract that covers this Cisco device. | [optional] [readonly] 
**ServiceLevel** | Pointer to **string** | The type of service contract that covers the Cisco device. | [optional] [readonly] 
**ServiceSku** | Pointer to **string** | The SKU of the service contract that covers the Cisco device. | [optional] [readonly] 
**ServiceStartDate** | Pointer to **time.Time** | Start date for the Cisco service contract that covers this Cisco device. | [optional] [readonly] 
**StateContract** | Pointer to **string** | Internal property used for triggering and tracking actions for contract information. * &#x60;Update&#x60; - Sn2Info/Contract information needs to be updated. * &#x60;OK&#x60; - Sn2Info/Contract information was fetched succcessfuly and updated. * &#x60;Failed&#x60; - Sn2Info/Contract information was not available  or failed while fetching. * &#x60;Retry&#x60; - Sn2Info/Contract information update failed and will be retried later. | [optional] [default to "Update"]
**WarrantyEndDate** | Pointer to **string** | End date for the warranty that covers the Cisco device. | [optional] [readonly] 
**WarrantyType** | Pointer to **string** | Type of warranty that covers the Cisco device. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Source** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewAssetDeviceContractInformation

`func NewAssetDeviceContractInformation(classId string, objectType string, ) *AssetDeviceContractInformation`

NewAssetDeviceContractInformation instantiates a new AssetDeviceContractInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceContractInformationWithDefaults

`func NewAssetDeviceContractInformationWithDefaults() *AssetDeviceContractInformation`

NewAssetDeviceContractInformationWithDefaults instantiates a new AssetDeviceContractInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceContractInformation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceContractInformation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceContractInformation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceContractInformation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceContractInformation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceContractInformation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContract

`func (o *AssetDeviceContractInformation) GetContract() AssetContractInformation`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AssetDeviceContractInformation) GetContractOk() (*AssetContractInformation, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AssetDeviceContractInformation) SetContract(v AssetContractInformation)`

SetContract sets Contract field to given value.

### HasContract

`func (o *AssetDeviceContractInformation) HasContract() bool`

HasContract returns a boolean if a field has been set.

### SetContractNil

`func (o *AssetDeviceContractInformation) SetContractNil(b bool)`

 SetContractNil sets the value for Contract to be an explicit nil

### UnsetContract
`func (o *AssetDeviceContractInformation) UnsetContract()`

UnsetContract ensures that no value is present for Contract, not even an explicit nil
### GetContractStatus

`func (o *AssetDeviceContractInformation) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *AssetDeviceContractInformation) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *AssetDeviceContractInformation) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.

### HasContractStatus

`func (o *AssetDeviceContractInformation) HasContractStatus() bool`

HasContractStatus returns a boolean if a field has been set.

### GetContractStatusReason

`func (o *AssetDeviceContractInformation) GetContractStatusReason() string`

GetContractStatusReason returns the ContractStatusReason field if non-nil, zero value otherwise.

### GetContractStatusReasonOk

`func (o *AssetDeviceContractInformation) GetContractStatusReasonOk() (*string, bool)`

GetContractStatusReasonOk returns a tuple with the ContractStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatusReason

`func (o *AssetDeviceContractInformation) SetContractStatusReason(v string)`

SetContractStatusReason sets ContractStatusReason field to given value.

### HasContractStatusReason

`func (o *AssetDeviceContractInformation) HasContractStatusReason() bool`

HasContractStatusReason returns a boolean if a field has been set.

### GetContractUnavailableRetryCount

`func (o *AssetDeviceContractInformation) GetContractUnavailableRetryCount() int64`

GetContractUnavailableRetryCount returns the ContractUnavailableRetryCount field if non-nil, zero value otherwise.

### GetContractUnavailableRetryCountOk

`func (o *AssetDeviceContractInformation) GetContractUnavailableRetryCountOk() (*int64, bool)`

GetContractUnavailableRetryCountOk returns a tuple with the ContractUnavailableRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractUnavailableRetryCount

`func (o *AssetDeviceContractInformation) SetContractUnavailableRetryCount(v int64)`

SetContractUnavailableRetryCount sets ContractUnavailableRetryCount field to given value.

### HasContractUnavailableRetryCount

`func (o *AssetDeviceContractInformation) HasContractUnavailableRetryCount() bool`

HasContractUnavailableRetryCount returns a boolean if a field has been set.

### GetContractUpdatedTime

`func (o *AssetDeviceContractInformation) GetContractUpdatedTime() time.Time`

GetContractUpdatedTime returns the ContractUpdatedTime field if non-nil, zero value otherwise.

### GetContractUpdatedTimeOk

`func (o *AssetDeviceContractInformation) GetContractUpdatedTimeOk() (*time.Time, bool)`

GetContractUpdatedTimeOk returns a tuple with the ContractUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractUpdatedTime

`func (o *AssetDeviceContractInformation) SetContractUpdatedTime(v time.Time)`

SetContractUpdatedTime sets ContractUpdatedTime field to given value.

### HasContractUpdatedTime

`func (o *AssetDeviceContractInformation) HasContractUpdatedTime() bool`

HasContractUpdatedTime returns a boolean if a field has been set.

### GetCoveredProductLineEndDate

`func (o *AssetDeviceContractInformation) GetCoveredProductLineEndDate() string`

GetCoveredProductLineEndDate returns the CoveredProductLineEndDate field if non-nil, zero value otherwise.

### GetCoveredProductLineEndDateOk

`func (o *AssetDeviceContractInformation) GetCoveredProductLineEndDateOk() (*string, bool)`

GetCoveredProductLineEndDateOk returns a tuple with the CoveredProductLineEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveredProductLineEndDate

`func (o *AssetDeviceContractInformation) SetCoveredProductLineEndDate(v string)`

SetCoveredProductLineEndDate sets CoveredProductLineEndDate field to given value.

### HasCoveredProductLineEndDate

`func (o *AssetDeviceContractInformation) HasCoveredProductLineEndDate() bool`

HasCoveredProductLineEndDate returns a boolean if a field has been set.

### GetDeviceId

`func (o *AssetDeviceContractInformation) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetDeviceContractInformation) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetDeviceContractInformation) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetDeviceContractInformation) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *AssetDeviceContractInformation) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AssetDeviceContractInformation) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AssetDeviceContractInformation) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *AssetDeviceContractInformation) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetEndCustomer

`func (o *AssetDeviceContractInformation) GetEndCustomer() AssetCustomerInformation`

GetEndCustomer returns the EndCustomer field if non-nil, zero value otherwise.

### GetEndCustomerOk

`func (o *AssetDeviceContractInformation) GetEndCustomerOk() (*AssetCustomerInformation, bool)`

GetEndCustomerOk returns a tuple with the EndCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomer

`func (o *AssetDeviceContractInformation) SetEndCustomer(v AssetCustomerInformation)`

SetEndCustomer sets EndCustomer field to given value.

### HasEndCustomer

`func (o *AssetDeviceContractInformation) HasEndCustomer() bool`

HasEndCustomer returns a boolean if a field has been set.

### SetEndCustomerNil

`func (o *AssetDeviceContractInformation) SetEndCustomerNil(b bool)`

 SetEndCustomerNil sets the value for EndCustomer to be an explicit nil

### UnsetEndCustomer
`func (o *AssetDeviceContractInformation) UnsetEndCustomer()`

UnsetEndCustomer ensures that no value is present for EndCustomer, not even an explicit nil
### GetEndUserGlobalUltimate

`func (o *AssetDeviceContractInformation) GetEndUserGlobalUltimate() AssetGlobalUltimate`

GetEndUserGlobalUltimate returns the EndUserGlobalUltimate field if non-nil, zero value otherwise.

### GetEndUserGlobalUltimateOk

`func (o *AssetDeviceContractInformation) GetEndUserGlobalUltimateOk() (*AssetGlobalUltimate, bool)`

GetEndUserGlobalUltimateOk returns a tuple with the EndUserGlobalUltimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserGlobalUltimate

`func (o *AssetDeviceContractInformation) SetEndUserGlobalUltimate(v AssetGlobalUltimate)`

SetEndUserGlobalUltimate sets EndUserGlobalUltimate field to given value.

### HasEndUserGlobalUltimate

`func (o *AssetDeviceContractInformation) HasEndUserGlobalUltimate() bool`

HasEndUserGlobalUltimate returns a boolean if a field has been set.

### SetEndUserGlobalUltimateNil

`func (o *AssetDeviceContractInformation) SetEndUserGlobalUltimateNil(b bool)`

 SetEndUserGlobalUltimateNil sets the value for EndUserGlobalUltimate to be an explicit nil

### UnsetEndUserGlobalUltimate
`func (o *AssetDeviceContractInformation) UnsetEndUserGlobalUltimate()`

UnsetEndUserGlobalUltimate ensures that no value is present for EndUserGlobalUltimate, not even an explicit nil
### GetIsValid

`func (o *AssetDeviceContractInformation) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AssetDeviceContractInformation) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AssetDeviceContractInformation) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AssetDeviceContractInformation) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetItemType

`func (o *AssetDeviceContractInformation) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *AssetDeviceContractInformation) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *AssetDeviceContractInformation) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *AssetDeviceContractInformation) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetLastDateOfSupport

`func (o *AssetDeviceContractInformation) GetLastDateOfSupport() time.Time`

GetLastDateOfSupport returns the LastDateOfSupport field if non-nil, zero value otherwise.

### GetLastDateOfSupportOk

`func (o *AssetDeviceContractInformation) GetLastDateOfSupportOk() (*time.Time, bool)`

GetLastDateOfSupportOk returns a tuple with the LastDateOfSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateOfSupport

`func (o *AssetDeviceContractInformation) SetLastDateOfSupport(v time.Time)`

SetLastDateOfSupport sets LastDateOfSupport field to given value.

### HasLastDateOfSupport

`func (o *AssetDeviceContractInformation) HasLastDateOfSupport() bool`

HasLastDateOfSupport returns a boolean if a field has been set.

### GetMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractInformation) GetMaintenancePurchaseOrderNumber() string`

GetMaintenancePurchaseOrderNumber returns the MaintenancePurchaseOrderNumber field if non-nil, zero value otherwise.

### GetMaintenancePurchaseOrderNumberOk

`func (o *AssetDeviceContractInformation) GetMaintenancePurchaseOrderNumberOk() (*string, bool)`

GetMaintenancePurchaseOrderNumberOk returns a tuple with the MaintenancePurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractInformation) SetMaintenancePurchaseOrderNumber(v string)`

SetMaintenancePurchaseOrderNumber sets MaintenancePurchaseOrderNumber field to given value.

### HasMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractInformation) HasMaintenancePurchaseOrderNumber() bool`

HasMaintenancePurchaseOrderNumber returns a boolean if a field has been set.

### GetMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractInformation) GetMaintenanceSalesOrderNumber() string`

GetMaintenanceSalesOrderNumber returns the MaintenanceSalesOrderNumber field if non-nil, zero value otherwise.

### GetMaintenanceSalesOrderNumberOk

`func (o *AssetDeviceContractInformation) GetMaintenanceSalesOrderNumberOk() (*string, bool)`

GetMaintenanceSalesOrderNumberOk returns a tuple with the MaintenanceSalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractInformation) SetMaintenanceSalesOrderNumber(v string)`

SetMaintenanceSalesOrderNumber sets MaintenanceSalesOrderNumber field to given value.

### HasMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractInformation) HasMaintenanceSalesOrderNumber() bool`

HasMaintenanceSalesOrderNumber returns a boolean if a field has been set.

### GetPlatformType

`func (o *AssetDeviceContractInformation) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AssetDeviceContractInformation) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AssetDeviceContractInformation) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AssetDeviceContractInformation) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetProduct

`func (o *AssetDeviceContractInformation) GetProduct() AssetProductInformation`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AssetDeviceContractInformation) GetProductOk() (*AssetProductInformation, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AssetDeviceContractInformation) SetProduct(v AssetProductInformation)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AssetDeviceContractInformation) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *AssetDeviceContractInformation) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *AssetDeviceContractInformation) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetPurchaseOrderNumber

`func (o *AssetDeviceContractInformation) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *AssetDeviceContractInformation) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *AssetDeviceContractInformation) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *AssetDeviceContractInformation) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetResellerGlobalUltimate

`func (o *AssetDeviceContractInformation) GetResellerGlobalUltimate() AssetGlobalUltimate`

GetResellerGlobalUltimate returns the ResellerGlobalUltimate field if non-nil, zero value otherwise.

### GetResellerGlobalUltimateOk

`func (o *AssetDeviceContractInformation) GetResellerGlobalUltimateOk() (*AssetGlobalUltimate, bool)`

GetResellerGlobalUltimateOk returns a tuple with the ResellerGlobalUltimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerGlobalUltimate

`func (o *AssetDeviceContractInformation) SetResellerGlobalUltimate(v AssetGlobalUltimate)`

SetResellerGlobalUltimate sets ResellerGlobalUltimate field to given value.

### HasResellerGlobalUltimate

`func (o *AssetDeviceContractInformation) HasResellerGlobalUltimate() bool`

HasResellerGlobalUltimate returns a boolean if a field has been set.

### SetResellerGlobalUltimateNil

`func (o *AssetDeviceContractInformation) SetResellerGlobalUltimateNil(b bool)`

 SetResellerGlobalUltimateNil sets the value for ResellerGlobalUltimate to be an explicit nil

### UnsetResellerGlobalUltimate
`func (o *AssetDeviceContractInformation) UnsetResellerGlobalUltimate()`

UnsetResellerGlobalUltimate ensures that no value is present for ResellerGlobalUltimate, not even an explicit nil
### GetSalesOrderNumber

`func (o *AssetDeviceContractInformation) GetSalesOrderNumber() string`

GetSalesOrderNumber returns the SalesOrderNumber field if non-nil, zero value otherwise.

### GetSalesOrderNumberOk

`func (o *AssetDeviceContractInformation) GetSalesOrderNumberOk() (*string, bool)`

GetSalesOrderNumberOk returns a tuple with the SalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrderNumber

`func (o *AssetDeviceContractInformation) SetSalesOrderNumber(v string)`

SetSalesOrderNumber sets SalesOrderNumber field to given value.

### HasSalesOrderNumber

`func (o *AssetDeviceContractInformation) HasSalesOrderNumber() bool`

HasSalesOrderNumber returns a boolean if a field has been set.

### GetServiceDescription

`func (o *AssetDeviceContractInformation) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *AssetDeviceContractInformation) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *AssetDeviceContractInformation) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.

### HasServiceDescription

`func (o *AssetDeviceContractInformation) HasServiceDescription() bool`

HasServiceDescription returns a boolean if a field has been set.

### GetServiceEndDate

`func (o *AssetDeviceContractInformation) GetServiceEndDate() time.Time`

GetServiceEndDate returns the ServiceEndDate field if non-nil, zero value otherwise.

### GetServiceEndDateOk

`func (o *AssetDeviceContractInformation) GetServiceEndDateOk() (*time.Time, bool)`

GetServiceEndDateOk returns a tuple with the ServiceEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndDate

`func (o *AssetDeviceContractInformation) SetServiceEndDate(v time.Time)`

SetServiceEndDate sets ServiceEndDate field to given value.

### HasServiceEndDate

`func (o *AssetDeviceContractInformation) HasServiceEndDate() bool`

HasServiceEndDate returns a boolean if a field has been set.

### GetServiceLevel

`func (o *AssetDeviceContractInformation) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *AssetDeviceContractInformation) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *AssetDeviceContractInformation) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *AssetDeviceContractInformation) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetServiceSku

`func (o *AssetDeviceContractInformation) GetServiceSku() string`

GetServiceSku returns the ServiceSku field if non-nil, zero value otherwise.

### GetServiceSkuOk

`func (o *AssetDeviceContractInformation) GetServiceSkuOk() (*string, bool)`

GetServiceSkuOk returns a tuple with the ServiceSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSku

`func (o *AssetDeviceContractInformation) SetServiceSku(v string)`

SetServiceSku sets ServiceSku field to given value.

### HasServiceSku

`func (o *AssetDeviceContractInformation) HasServiceSku() bool`

HasServiceSku returns a boolean if a field has been set.

### GetServiceStartDate

`func (o *AssetDeviceContractInformation) GetServiceStartDate() time.Time`

GetServiceStartDate returns the ServiceStartDate field if non-nil, zero value otherwise.

### GetServiceStartDateOk

`func (o *AssetDeviceContractInformation) GetServiceStartDateOk() (*time.Time, bool)`

GetServiceStartDateOk returns a tuple with the ServiceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStartDate

`func (o *AssetDeviceContractInformation) SetServiceStartDate(v time.Time)`

SetServiceStartDate sets ServiceStartDate field to given value.

### HasServiceStartDate

`func (o *AssetDeviceContractInformation) HasServiceStartDate() bool`

HasServiceStartDate returns a boolean if a field has been set.

### GetStateContract

`func (o *AssetDeviceContractInformation) GetStateContract() string`

GetStateContract returns the StateContract field if non-nil, zero value otherwise.

### GetStateContractOk

`func (o *AssetDeviceContractInformation) GetStateContractOk() (*string, bool)`

GetStateContractOk returns a tuple with the StateContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateContract

`func (o *AssetDeviceContractInformation) SetStateContract(v string)`

SetStateContract sets StateContract field to given value.

### HasStateContract

`func (o *AssetDeviceContractInformation) HasStateContract() bool`

HasStateContract returns a boolean if a field has been set.

### GetWarrantyEndDate

`func (o *AssetDeviceContractInformation) GetWarrantyEndDate() string`

GetWarrantyEndDate returns the WarrantyEndDate field if non-nil, zero value otherwise.

### GetWarrantyEndDateOk

`func (o *AssetDeviceContractInformation) GetWarrantyEndDateOk() (*string, bool)`

GetWarrantyEndDateOk returns a tuple with the WarrantyEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEndDate

`func (o *AssetDeviceContractInformation) SetWarrantyEndDate(v string)`

SetWarrantyEndDate sets WarrantyEndDate field to given value.

### HasWarrantyEndDate

`func (o *AssetDeviceContractInformation) HasWarrantyEndDate() bool`

HasWarrantyEndDate returns a boolean if a field has been set.

### GetWarrantyType

`func (o *AssetDeviceContractInformation) GetWarrantyType() string`

GetWarrantyType returns the WarrantyType field if non-nil, zero value otherwise.

### GetWarrantyTypeOk

`func (o *AssetDeviceContractInformation) GetWarrantyTypeOk() (*string, bool)`

GetWarrantyTypeOk returns a tuple with the WarrantyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyType

`func (o *AssetDeviceContractInformation) SetWarrantyType(v string)`

SetWarrantyType sets WarrantyType field to given value.

### HasWarrantyType

`func (o *AssetDeviceContractInformation) HasWarrantyType() bool`

HasWarrantyType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetDeviceContractInformation) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetDeviceContractInformation) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetDeviceContractInformation) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetDeviceContractInformation) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *AssetDeviceContractInformation) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *AssetDeviceContractInformation) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSource

`func (o *AssetDeviceContractInformation) GetSource() MoBaseMoRelationship`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AssetDeviceContractInformation) GetSourceOk() (*MoBaseMoRelationship, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AssetDeviceContractInformation) SetSource(v MoBaseMoRelationship)`

SetSource sets Source field to given value.

### HasSource

`func (o *AssetDeviceContractInformation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AssetDeviceContractInformation) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AssetDeviceContractInformation) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


