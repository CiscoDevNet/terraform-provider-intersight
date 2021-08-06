# AssetTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.Target"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.Target"]
**ClaimedByUserName** | Pointer to **string** | The name or email id of the user who claimed the target. | [optional] [readonly] 
**Connections** | Pointer to [**[]AssetConnection**](AssetConnection.md) |  | [optional] 
**ConnectorVersion** | Pointer to **string** | The Device Connector version for target types which are managed by via embedded Device Connector. | [optional] [readonly] 
**ExternalIpAddress** | Pointer to **string** | ExternalIpAddress is applicable for targets which are managed via an Intersight Device Connector. The value is the IP Address of the target as seen from Intersight. It is either the IP Address of the managed target&#39;s interface which has a route to the internet or a NAT IP Addresss when the target is deployed in a private network. | [optional] [readonly] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**ManagementLocation** | Pointer to **string** | The location from which Intersight manages the target. * &#x60;Unknown&#x60; - The management mechanism is not detected. Unknown is used as a default by the implementation and is not an allowed user input. * &#x60;Intersight&#x60; - Management of a target is performed directly from Intersight. Network connections are established from Intersight to a management interface of the Target and authenticated using user provided credentials. * &#x60;IntersightAssist&#x60; - Management of a target is performed via a selected Intersight Assist. Network connections are established from the Intersight Assist to a management interface of the Target. * &#x60;DeviceConnector&#x60; - An Intersight Device Connector running within the Target establishes a connection to Intersight and management of the target is performed via this connection. | [optional] [default to "Unknown"]
**Name** | Pointer to **string** | A user provided name for the managed target. | [optional] 
**ProductId** | Pointer to **[]string** |  | [optional] 
**ReadOnly** | Pointer to **bool** | For targets which are managed by an embedded Intersight Device Connector, this field indicates that an administrator of the target has disabled management operations of the Device Connector and only monitoring is permitted. | [optional] [readonly] 
**Services** | Pointer to [**[]AssetService**](AssetService.md) |  | [optional] 
**Status** | Pointer to **string** | Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target. * &#x60;&#x60; - The target details have been persisted but Intersight has not yet attempted to connect to the target. * &#x60;Connected&#x60; - Intersight is able to establish a connection to the target and initiate management activities. * &#x60;NotConnected&#x60; - Intersight is unable to establish a connection to the target. * &#x60;ClaimInProgress&#x60; - Claim of the target is in progress. A connection to the target has not been fully established. * &#x60;Unclaimed&#x60; - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * &#x60;Claimed&#x60; - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect. | [optional] [readonly] [default to ""]
**StatusErrorReason** | Pointer to **string** | StatusErrorReason provides additional context for the Status. | [optional] [readonly] 
**TargetId** | Pointer to **[]string** |  | [optional] 
**TargetType** | Pointer to **string** | The type of the managed target. For example a UCS Server or VMware Vcenter target. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;IMCRack&#x60; - A standalone UCS M6 and above server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;UCSC890&#x60; - A standalone Cisco UCSC890 server. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * &#x60;ReadHatOpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;AmazonWebService&#x60; - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * &#x60;AmazonWebServiceBilling&#x60; - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * &#x60;DellCompellent&#x60; - A Dell Compellent storage system. * &#x60;HPE3Par&#x60; - A HPE 3PAR storage system. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * &#x60;HPEOneView&#x60; - A HPE Oneview management system that manages compute, storage, and networking. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. * &#x60;TerraformCloud&#x60; - A Terraform Cloud account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * &#x60;CustomTarget&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;HTTPEndpoint&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. | [optional] [default to ""]
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Assist** | Pointer to [**AssetTargetRelationship**](asset.Target.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetTarget

`func NewAssetTarget(classId string, objectType string, ) *AssetTarget`

NewAssetTarget instantiates a new AssetTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTargetWithDefaults

`func NewAssetTargetWithDefaults() *AssetTarget`

NewAssetTargetWithDefaults instantiates a new AssetTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetTarget) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetTarget) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetTarget) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetTarget) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetTarget) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetTarget) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClaimedByUserName

`func (o *AssetTarget) GetClaimedByUserName() string`

GetClaimedByUserName returns the ClaimedByUserName field if non-nil, zero value otherwise.

### GetClaimedByUserNameOk

`func (o *AssetTarget) GetClaimedByUserNameOk() (*string, bool)`

GetClaimedByUserNameOk returns a tuple with the ClaimedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedByUserName

`func (o *AssetTarget) SetClaimedByUserName(v string)`

SetClaimedByUserName sets ClaimedByUserName field to given value.

### HasClaimedByUserName

`func (o *AssetTarget) HasClaimedByUserName() bool`

HasClaimedByUserName returns a boolean if a field has been set.

### GetConnections

`func (o *AssetTarget) GetConnections() []AssetConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *AssetTarget) GetConnectionsOk() (*[]AssetConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *AssetTarget) SetConnections(v []AssetConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *AssetTarget) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *AssetTarget) SetConnectionsNil(b bool)`

 SetConnectionsNil sets the value for Connections to be an explicit nil

### UnsetConnections
`func (o *AssetTarget) UnsetConnections()`

UnsetConnections ensures that no value is present for Connections, not even an explicit nil
### GetConnectorVersion

`func (o *AssetTarget) GetConnectorVersion() string`

GetConnectorVersion returns the ConnectorVersion field if non-nil, zero value otherwise.

### GetConnectorVersionOk

`func (o *AssetTarget) GetConnectorVersionOk() (*string, bool)`

GetConnectorVersionOk returns a tuple with the ConnectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorVersion

`func (o *AssetTarget) SetConnectorVersion(v string)`

SetConnectorVersion sets ConnectorVersion field to given value.

### HasConnectorVersion

`func (o *AssetTarget) HasConnectorVersion() bool`

HasConnectorVersion returns a boolean if a field has been set.

### GetExternalIpAddress

`func (o *AssetTarget) GetExternalIpAddress() string`

GetExternalIpAddress returns the ExternalIpAddress field if non-nil, zero value otherwise.

### GetExternalIpAddressOk

`func (o *AssetTarget) GetExternalIpAddressOk() (*string, bool)`

GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddress

`func (o *AssetTarget) SetExternalIpAddress(v string)`

SetExternalIpAddress sets ExternalIpAddress field to given value.

### HasExternalIpAddress

`func (o *AssetTarget) HasExternalIpAddress() bool`

HasExternalIpAddress returns a boolean if a field has been set.

### GetIpAddress

`func (o *AssetTarget) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AssetTarget) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AssetTarget) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AssetTarget) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *AssetTarget) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *AssetTarget) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetManagementLocation

`func (o *AssetTarget) GetManagementLocation() string`

GetManagementLocation returns the ManagementLocation field if non-nil, zero value otherwise.

### GetManagementLocationOk

`func (o *AssetTarget) GetManagementLocationOk() (*string, bool)`

GetManagementLocationOk returns a tuple with the ManagementLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementLocation

`func (o *AssetTarget) SetManagementLocation(v string)`

SetManagementLocation sets ManagementLocation field to given value.

### HasManagementLocation

`func (o *AssetTarget) HasManagementLocation() bool`

HasManagementLocation returns a boolean if a field has been set.

### GetName

`func (o *AssetTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductId

`func (o *AssetTarget) GetProductId() []string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AssetTarget) GetProductIdOk() (*[]string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AssetTarget) SetProductId(v []string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AssetTarget) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *AssetTarget) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *AssetTarget) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetReadOnly

`func (o *AssetTarget) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AssetTarget) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AssetTarget) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AssetTarget) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetServices

`func (o *AssetTarget) GetServices() []AssetService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AssetTarget) GetServicesOk() (*[]AssetService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AssetTarget) SetServices(v []AssetService)`

SetServices sets Services field to given value.

### HasServices

`func (o *AssetTarget) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *AssetTarget) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *AssetTarget) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetStatus

`func (o *AssetTarget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetTarget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetTarget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetTarget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusErrorReason

`func (o *AssetTarget) GetStatusErrorReason() string`

GetStatusErrorReason returns the StatusErrorReason field if non-nil, zero value otherwise.

### GetStatusErrorReasonOk

`func (o *AssetTarget) GetStatusErrorReasonOk() (*string, bool)`

GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusErrorReason

`func (o *AssetTarget) SetStatusErrorReason(v string)`

SetStatusErrorReason sets StatusErrorReason field to given value.

### HasStatusErrorReason

`func (o *AssetTarget) HasStatusErrorReason() bool`

HasStatusErrorReason returns a boolean if a field has been set.

### GetTargetId

`func (o *AssetTarget) GetTargetId() []string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AssetTarget) GetTargetIdOk() (*[]string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AssetTarget) SetTargetId(v []string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AssetTarget) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *AssetTarget) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *AssetTarget) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetType

`func (o *AssetTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AssetTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AssetTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *AssetTarget) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetAccount

`func (o *AssetTarget) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetTarget) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetTarget) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetTarget) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssist

`func (o *AssetTarget) GetAssist() AssetTargetRelationship`

GetAssist returns the Assist field if non-nil, zero value otherwise.

### GetAssistOk

`func (o *AssetTarget) GetAssistOk() (*AssetTargetRelationship, bool)`

GetAssistOk returns a tuple with the Assist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssist

`func (o *AssetTarget) SetAssist(v AssetTargetRelationship)`

SetAssist sets Assist field to given value.

### HasAssist

`func (o *AssetTarget) HasAssist() bool`

HasAssist returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetTarget) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetTarget) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetTarget) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetTarget) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


