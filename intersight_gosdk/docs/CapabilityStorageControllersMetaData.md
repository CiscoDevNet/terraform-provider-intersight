# CapabilityStorageControllersMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.StorageControllersMetaData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.StorageControllersMetaData"]
**ControllerPid** | Pointer to **string** | The controller Pid to which the metadata applies. | [optional] [readonly] 
**ControllerSeries** | Pointer to **string** | Product family of the controller. | [optional] [readonly] 
**ControllerType** | Pointer to **string** | Type of the controller. Expected values are Hba, RAID etc. * &#x60;Unknown&#x60; - Storage controller type is not known. * &#x60;Hba&#x60; - The Hba type of storage controller. * &#x60;Raid&#x60; - The Raid type of storage controller. | [optional] [readonly] [default to "Unknown"]
**ControllerVendor** | Pointer to **string** | Manufacturer of the controller. | [optional] [readonly] 
**IsHybridDriveSlotsSupported** | Pointer to **bool** | Controller supports HybridDriveslots. | [optional] [readonly] 
**SupportedStorageServerActions** | Pointer to **[]string** |  | [optional] 
**TargetType** | Pointer to **string** | The target type to which the metadata applies. * &#x60;&#x60; - An unrecognized platform type. * &#x60;APIC&#x60; - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * &#x60;CAPIC&#x60; - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * &#x60;DCNM&#x60; - A Cisco Data Center Network Manager (DCNM) instance. * &#x60;UCSFI&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * &#x60;IMC&#x60; - A standalone Cisco UCS rack server (Deprecated). * &#x60;IMCM4&#x60; - A standalone Cisco UCS C-Series or S-Series M4 server. * &#x60;IMCM5&#x60; - A standalone Cisco UCS C-Series or S-Series M5 server. * &#x60;IMCRack&#x60; - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * &#x60;UCSIOM&#x60; - A Cisco UCS Blade Chassis I/O Module (IOM). * &#x60;UCSPCIeNODE&#x60; - A Cisco UCS PCIe Node in a blade form factor. * &#x60;UCSXFM&#x60; - A Cisco UCS Fabric Extender Module (XFM). * &#x60;HX&#x60; - A Cisco HyperFlex (HX) cluster. * &#x60;UCSD&#x60; - A Cisco UCS Director (UCSD) instance. * &#x60;UCSXECMC&#x60; - A Cisco UCSXE Chassis Management Controller. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance instance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist instance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * &#x60;NexusDevice&#x60; - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * &#x60;ACISwitch&#x60; - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * &#x60;NexusSwitch&#x60; - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * &#x60;MDSSwitch&#x60; - A Cisco MDS Switch that is managed using the embedded Device Connector. * &#x60;MDSDevice&#x60; - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * &#x60;UCSC885&#x60; - A standalone Cisco UCS C885 Server. * &#x60;CAI845A&#x60; - A standalone Cisco AI 845A Server. * &#x60;UCSC890&#x60; - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * &#x60;NetAppOntap&#x60; - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVmax&#x60; - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVplex&#x60; - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;VmwareVcenter&#x60; - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * &#x60;AppDynamics&#x60; - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * &#x60;Dynatrace&#x60; - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;NewRelic&#x60; - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;Umbrella&#x60; - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;MySqlServer&#x60; - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;OracleDatabaseServer&#x60; - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;IBMWebSphereApplicationServer&#x60; - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;JavaVirtualMachine&#x60; - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * &#x60;AmazonWebService&#x60; - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatform&#x60; - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatformBilling&#x60; - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureBilling&#x60; - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;DellCompellent&#x60; - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;HPE3Par&#x60; - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * &#x60;NutanixPrismCentral&#x60; - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * &#x60;HPEOneView&#x60; - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * &#x60;GenericTarget&#x60; - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * &#x60;IMCBlade&#x60; - A Cisco UCS blade server managed by Cisco Intersight. * &#x60;TerraformCloud&#x60; - A Terraform Cloud Business Tier account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * &#x60;CustomTarget&#x60; - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * &#x60;AnsibleEndpoint&#x60; - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * &#x60;HTTPEndpoint&#x60; - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * &#x60;SSHEndpoint&#x60; - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * &#x60;CiscoDNAC&#x60; - A Cisco Digital Network Architecture (DNA) Center appliance. * &#x60;CiscoFMC&#x60; - A Cisco Secure Firewall Management Center. * &#x60;ViptelaCloud&#x60; - A Cisco Viptela SD-WAN Cloud. * &#x60;MerakiCloud&#x60; - A Cisco Meraki Organization. * &#x60;CiscoISE&#x60; - A Cisco Identity Services Engine (ISE) target. | [optional] [readonly] [default to ""]
**Vendor** | Pointer to **string** | Manufacturer of the server. | [optional] [readonly] 

## Methods

### NewCapabilityStorageControllersMetaData

`func NewCapabilityStorageControllersMetaData(classId string, objectType string, ) *CapabilityStorageControllersMetaData`

NewCapabilityStorageControllersMetaData instantiates a new CapabilityStorageControllersMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityStorageControllersMetaDataWithDefaults

`func NewCapabilityStorageControllersMetaDataWithDefaults() *CapabilityStorageControllersMetaData`

NewCapabilityStorageControllersMetaDataWithDefaults instantiates a new CapabilityStorageControllersMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityStorageControllersMetaData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityStorageControllersMetaData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityStorageControllersMetaData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityStorageControllersMetaData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityStorageControllersMetaData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityStorageControllersMetaData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerPid

`func (o *CapabilityStorageControllersMetaData) GetControllerPid() string`

GetControllerPid returns the ControllerPid field if non-nil, zero value otherwise.

### GetControllerPidOk

`func (o *CapabilityStorageControllersMetaData) GetControllerPidOk() (*string, bool)`

GetControllerPidOk returns a tuple with the ControllerPid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerPid

`func (o *CapabilityStorageControllersMetaData) SetControllerPid(v string)`

SetControllerPid sets ControllerPid field to given value.

### HasControllerPid

`func (o *CapabilityStorageControllersMetaData) HasControllerPid() bool`

HasControllerPid returns a boolean if a field has been set.

### GetControllerSeries

`func (o *CapabilityStorageControllersMetaData) GetControllerSeries() string`

GetControllerSeries returns the ControllerSeries field if non-nil, zero value otherwise.

### GetControllerSeriesOk

`func (o *CapabilityStorageControllersMetaData) GetControllerSeriesOk() (*string, bool)`

GetControllerSeriesOk returns a tuple with the ControllerSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerSeries

`func (o *CapabilityStorageControllersMetaData) SetControllerSeries(v string)`

SetControllerSeries sets ControllerSeries field to given value.

### HasControllerSeries

`func (o *CapabilityStorageControllersMetaData) HasControllerSeries() bool`

HasControllerSeries returns a boolean if a field has been set.

### GetControllerType

`func (o *CapabilityStorageControllersMetaData) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *CapabilityStorageControllersMetaData) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *CapabilityStorageControllersMetaData) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *CapabilityStorageControllersMetaData) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### GetControllerVendor

`func (o *CapabilityStorageControllersMetaData) GetControllerVendor() string`

GetControllerVendor returns the ControllerVendor field if non-nil, zero value otherwise.

### GetControllerVendorOk

`func (o *CapabilityStorageControllersMetaData) GetControllerVendorOk() (*string, bool)`

GetControllerVendorOk returns a tuple with the ControllerVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVendor

`func (o *CapabilityStorageControllersMetaData) SetControllerVendor(v string)`

SetControllerVendor sets ControllerVendor field to given value.

### HasControllerVendor

`func (o *CapabilityStorageControllersMetaData) HasControllerVendor() bool`

HasControllerVendor returns a boolean if a field has been set.

### GetIsHybridDriveSlotsSupported

`func (o *CapabilityStorageControllersMetaData) GetIsHybridDriveSlotsSupported() bool`

GetIsHybridDriveSlotsSupported returns the IsHybridDriveSlotsSupported field if non-nil, zero value otherwise.

### GetIsHybridDriveSlotsSupportedOk

`func (o *CapabilityStorageControllersMetaData) GetIsHybridDriveSlotsSupportedOk() (*bool, bool)`

GetIsHybridDriveSlotsSupportedOk returns a tuple with the IsHybridDriveSlotsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHybridDriveSlotsSupported

`func (o *CapabilityStorageControllersMetaData) SetIsHybridDriveSlotsSupported(v bool)`

SetIsHybridDriveSlotsSupported sets IsHybridDriveSlotsSupported field to given value.

### HasIsHybridDriveSlotsSupported

`func (o *CapabilityStorageControllersMetaData) HasIsHybridDriveSlotsSupported() bool`

HasIsHybridDriveSlotsSupported returns a boolean if a field has been set.

### GetSupportedStorageServerActions

`func (o *CapabilityStorageControllersMetaData) GetSupportedStorageServerActions() []string`

GetSupportedStorageServerActions returns the SupportedStorageServerActions field if non-nil, zero value otherwise.

### GetSupportedStorageServerActionsOk

`func (o *CapabilityStorageControllersMetaData) GetSupportedStorageServerActionsOk() (*[]string, bool)`

GetSupportedStorageServerActionsOk returns a tuple with the SupportedStorageServerActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedStorageServerActions

`func (o *CapabilityStorageControllersMetaData) SetSupportedStorageServerActions(v []string)`

SetSupportedStorageServerActions sets SupportedStorageServerActions field to given value.

### HasSupportedStorageServerActions

`func (o *CapabilityStorageControllersMetaData) HasSupportedStorageServerActions() bool`

HasSupportedStorageServerActions returns a boolean if a field has been set.

### SetSupportedStorageServerActionsNil

`func (o *CapabilityStorageControllersMetaData) SetSupportedStorageServerActionsNil(b bool)`

 SetSupportedStorageServerActionsNil sets the value for SupportedStorageServerActions to be an explicit nil

### UnsetSupportedStorageServerActions
`func (o *CapabilityStorageControllersMetaData) UnsetSupportedStorageServerActions()`

UnsetSupportedStorageServerActions ensures that no value is present for SupportedStorageServerActions, not even an explicit nil
### GetTargetType

`func (o *CapabilityStorageControllersMetaData) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CapabilityStorageControllersMetaData) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CapabilityStorageControllersMetaData) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CapabilityStorageControllersMetaData) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetVendor

`func (o *CapabilityStorageControllersMetaData) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityStorageControllersMetaData) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityStorageControllersMetaData) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityStorageControllersMetaData) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


