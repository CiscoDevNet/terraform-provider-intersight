# AssetManagedDeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**CommCredential**](comm.Credential.md) |  | [optional] 
**DeviceType** | Pointer to **string** | Type of the Device such as VMware, Pure Storage supported by Intersight Assist. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - Intersight on-premise appliance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft HyperV system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A Dynatrace controller that monitors applications. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;MicrosoftAzure&#x60; - A Microsoft Azure target. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. | [optional] [default to ""]
**IgnoreCert** | Pointer to **bool** | Ignore Certificates with protocol https for connecting to the Managed Device. It is not used for other protocols. | [optional] 
**IsEnabled** | Pointer to **bool** | Device is Enabled/Disabled. | [optional] 
**ManagementAddress** | Pointer to **string** | Management address of the device. It can be IPv4, IPv6 or FQDN. | [optional] 
**Port** | Pointer to **int64** | Port to use for connecting to the Managed Device. Port is optional. If not specified, default port for protocol is used. | [optional] 
**Protocol** | Pointer to **string** | Protocol to use for connecting to the Managed Device. * &#x60;https&#x60; - HTTPS communication is used for connecting to Device. * &#x60;http&#x60; - HTTP communication is used for connecting to Device. | [optional] [default to "https"]
**Status** | Pointer to [**AssetManagedDeviceStatus**](asset.ManagedDeviceStatus.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**DeviceConnectorManager** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewAssetManagedDeviceAllOf

`func NewAssetManagedDeviceAllOf() *AssetManagedDeviceAllOf`

NewAssetManagedDeviceAllOf instantiates a new AssetManagedDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetManagedDeviceAllOfWithDefaults

`func NewAssetManagedDeviceAllOfWithDefaults() *AssetManagedDeviceAllOf`

NewAssetManagedDeviceAllOfWithDefaults instantiates a new AssetManagedDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *AssetManagedDeviceAllOf) GetCredential() CommCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AssetManagedDeviceAllOf) GetCredentialOk() (*CommCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AssetManagedDeviceAllOf) SetCredential(v CommCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AssetManagedDeviceAllOf) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDeviceType

`func (o *AssetManagedDeviceAllOf) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AssetManagedDeviceAllOf) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AssetManagedDeviceAllOf) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *AssetManagedDeviceAllOf) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIgnoreCert

`func (o *AssetManagedDeviceAllOf) GetIgnoreCert() bool`

GetIgnoreCert returns the IgnoreCert field if non-nil, zero value otherwise.

### GetIgnoreCertOk

`func (o *AssetManagedDeviceAllOf) GetIgnoreCertOk() (*bool, bool)`

GetIgnoreCertOk returns a tuple with the IgnoreCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCert

`func (o *AssetManagedDeviceAllOf) SetIgnoreCert(v bool)`

SetIgnoreCert sets IgnoreCert field to given value.

### HasIgnoreCert

`func (o *AssetManagedDeviceAllOf) HasIgnoreCert() bool`

HasIgnoreCert returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AssetManagedDeviceAllOf) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AssetManagedDeviceAllOf) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AssetManagedDeviceAllOf) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AssetManagedDeviceAllOf) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetManagementAddress

`func (o *AssetManagedDeviceAllOf) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *AssetManagedDeviceAllOf) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *AssetManagedDeviceAllOf) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *AssetManagedDeviceAllOf) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPort

`func (o *AssetManagedDeviceAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AssetManagedDeviceAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AssetManagedDeviceAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *AssetManagedDeviceAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *AssetManagedDeviceAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AssetManagedDeviceAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AssetManagedDeviceAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AssetManagedDeviceAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *AssetManagedDeviceAllOf) GetStatus() AssetManagedDeviceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetManagedDeviceAllOf) GetStatusOk() (*AssetManagedDeviceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetManagedDeviceAllOf) SetStatus(v AssetManagedDeviceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetManagedDeviceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *AssetManagedDeviceAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetManagedDeviceAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetManagedDeviceAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetManagedDeviceAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDeviceConnectorManager

`func (o *AssetManagedDeviceAllOf) GetDeviceConnectorManager() AssetDeviceRegistrationRelationship`

GetDeviceConnectorManager returns the DeviceConnectorManager field if non-nil, zero value otherwise.

### GetDeviceConnectorManagerOk

`func (o *AssetManagedDeviceAllOf) GetDeviceConnectorManagerOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceConnectorManagerOk returns a tuple with the DeviceConnectorManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConnectorManager

`func (o *AssetManagedDeviceAllOf) SetDeviceConnectorManager(v AssetDeviceRegistrationRelationship)`

SetDeviceConnectorManager sets DeviceConnectorManager field to given value.

### HasDeviceConnectorManager

`func (o *AssetManagedDeviceAllOf) HasDeviceConnectorManager() bool`

HasDeviceConnectorManager returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetManagedDeviceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetManagedDeviceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetManagedDeviceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetManagedDeviceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *AssetManagedDeviceAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *AssetManagedDeviceAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *AssetManagedDeviceAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *AssetManagedDeviceAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


