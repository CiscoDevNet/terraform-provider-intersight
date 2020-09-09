# AssetDeviceRegistrationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | An identifier for the credential used by the device connector to authenticate with the Intersight web socket gateway. | [optional] 
**ClaimedByUserName** | Pointer to **string** | The name of the user who claimed the device for the account. | [optional] [readonly] 
**ClaimedTime** | Pointer to [**time.Time**](time.Time.md) | The date and time at which the device was claimed to this account. | [optional] [readonly] 
**DeviceHostname** | Pointer to **[]string** |  | [optional] 
**DeviceIpAddress** | Pointer to **[]string** |  | [optional] 
**ExecutionMode** | Pointer to **string** | Indicates if the platform is an actual device or an emulated device for testing, demos, etc. Permitted values are [Normal, Emulator, ContainerEmulator]. * &#x60;&#x60; - The device reported an empty or unrecognized executionMode. * &#x60;Normal&#x60; - The device connector is running in normal mode, i.e. it is not a simulation. * &#x60;Emulator&#x60; - The device connector is running in simulation mode inside an emulated device. * &#x60;ContainerEmulator&#x60; - The device connector is running in simulation mode inside a containerized emulated device. | [optional] [default to ""]
**ParentSignature** | Pointer to [**AssetParentConnectionSignature**](asset.ParentConnectionSignature.md) |  | [optional] 
**Pid** | Pointer to **[]string** |  | [optional] 
**PlatformType** | Pointer to **string** | The platform type on which device connector is executing. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - Intersight on-premise appliance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft HyperV system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A Dynatrace controller that monitors applications. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;MicrosoftAzure&#x60; - A Microsoft Azure target. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. | [optional] [default to ""]
**PublicAccessKey** | Pointer to **string** | The device connector&#39;s public key used by Intersight to authenticate a connection from the device connector. The public key is used to verify that the signature a device connector sends on connect has been signed by the connector&#39;s private key stored on the device&#39;s filesystem. Must be a PEM encoded RSA public key string. | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | Flag reported by devices to indicate an administrator of the device has disabled management operations of the device connector and only monitoring is permitted. | [optional] [readonly] 
**Serial** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | The vendor of the managed device. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**ClaimedByUser** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 
**ClusterMembers** | Pointer to [**[]AssetClusterMemberRelationship**](asset.ClusterMember.Relationship.md) | An array of relationships to assetClusterMember resources. | [optional] [readonly] 
**DeviceClaim** | Pointer to [**AssetDeviceClaimRelationship**](asset.DeviceClaim.Relationship.md) |  | [optional] 
**DeviceConfiguration** | Pointer to [**AssetDeviceConfigurationRelationship**](asset.DeviceConfiguration.Relationship.md) |  | [optional] 
**DomainGroup** | Pointer to [**IamDomainGroupRelationship**](iam.DomainGroup.Relationship.md) |  | [optional] 
**ParentConnection** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetDeviceRegistrationAllOf

`func NewAssetDeviceRegistrationAllOf() *AssetDeviceRegistrationAllOf`

NewAssetDeviceRegistrationAllOf instantiates a new AssetDeviceRegistrationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceRegistrationAllOfWithDefaults

`func NewAssetDeviceRegistrationAllOfWithDefaults() *AssetDeviceRegistrationAllOf`

NewAssetDeviceRegistrationAllOfWithDefaults instantiates a new AssetDeviceRegistrationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AssetDeviceRegistrationAllOf) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AssetDeviceRegistrationAllOf) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AssetDeviceRegistrationAllOf) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AssetDeviceRegistrationAllOf) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetClaimedByUserName

`func (o *AssetDeviceRegistrationAllOf) GetClaimedByUserName() string`

GetClaimedByUserName returns the ClaimedByUserName field if non-nil, zero value otherwise.

### GetClaimedByUserNameOk

`func (o *AssetDeviceRegistrationAllOf) GetClaimedByUserNameOk() (*string, bool)`

GetClaimedByUserNameOk returns a tuple with the ClaimedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedByUserName

`func (o *AssetDeviceRegistrationAllOf) SetClaimedByUserName(v string)`

SetClaimedByUserName sets ClaimedByUserName field to given value.

### HasClaimedByUserName

`func (o *AssetDeviceRegistrationAllOf) HasClaimedByUserName() bool`

HasClaimedByUserName returns a boolean if a field has been set.

### GetClaimedTime

`func (o *AssetDeviceRegistrationAllOf) GetClaimedTime() time.Time`

GetClaimedTime returns the ClaimedTime field if non-nil, zero value otherwise.

### GetClaimedTimeOk

`func (o *AssetDeviceRegistrationAllOf) GetClaimedTimeOk() (*time.Time, bool)`

GetClaimedTimeOk returns a tuple with the ClaimedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedTime

`func (o *AssetDeviceRegistrationAllOf) SetClaimedTime(v time.Time)`

SetClaimedTime sets ClaimedTime field to given value.

### HasClaimedTime

`func (o *AssetDeviceRegistrationAllOf) HasClaimedTime() bool`

HasClaimedTime returns a boolean if a field has been set.

### GetDeviceHostname

`func (o *AssetDeviceRegistrationAllOf) GetDeviceHostname() []string`

GetDeviceHostname returns the DeviceHostname field if non-nil, zero value otherwise.

### GetDeviceHostnameOk

`func (o *AssetDeviceRegistrationAllOf) GetDeviceHostnameOk() (*[]string, bool)`

GetDeviceHostnameOk returns a tuple with the DeviceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHostname

`func (o *AssetDeviceRegistrationAllOf) SetDeviceHostname(v []string)`

SetDeviceHostname sets DeviceHostname field to given value.

### HasDeviceHostname

`func (o *AssetDeviceRegistrationAllOf) HasDeviceHostname() bool`

HasDeviceHostname returns a boolean if a field has been set.

### GetDeviceIpAddress

`func (o *AssetDeviceRegistrationAllOf) GetDeviceIpAddress() []string`

GetDeviceIpAddress returns the DeviceIpAddress field if non-nil, zero value otherwise.

### GetDeviceIpAddressOk

`func (o *AssetDeviceRegistrationAllOf) GetDeviceIpAddressOk() (*[]string, bool)`

GetDeviceIpAddressOk returns a tuple with the DeviceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIpAddress

`func (o *AssetDeviceRegistrationAllOf) SetDeviceIpAddress(v []string)`

SetDeviceIpAddress sets DeviceIpAddress field to given value.

### HasDeviceIpAddress

`func (o *AssetDeviceRegistrationAllOf) HasDeviceIpAddress() bool`

HasDeviceIpAddress returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AssetDeviceRegistrationAllOf) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AssetDeviceRegistrationAllOf) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AssetDeviceRegistrationAllOf) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AssetDeviceRegistrationAllOf) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetParentSignature

`func (o *AssetDeviceRegistrationAllOf) GetParentSignature() AssetParentConnectionSignature`

GetParentSignature returns the ParentSignature field if non-nil, zero value otherwise.

### GetParentSignatureOk

`func (o *AssetDeviceRegistrationAllOf) GetParentSignatureOk() (*AssetParentConnectionSignature, bool)`

GetParentSignatureOk returns a tuple with the ParentSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSignature

`func (o *AssetDeviceRegistrationAllOf) SetParentSignature(v AssetParentConnectionSignature)`

SetParentSignature sets ParentSignature field to given value.

### HasParentSignature

`func (o *AssetDeviceRegistrationAllOf) HasParentSignature() bool`

HasParentSignature returns a boolean if a field has been set.

### GetPid

`func (o *AssetDeviceRegistrationAllOf) GetPid() []string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *AssetDeviceRegistrationAllOf) GetPidOk() (*[]string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *AssetDeviceRegistrationAllOf) SetPid(v []string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *AssetDeviceRegistrationAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformType

`func (o *AssetDeviceRegistrationAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AssetDeviceRegistrationAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AssetDeviceRegistrationAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AssetDeviceRegistrationAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPublicAccessKey

`func (o *AssetDeviceRegistrationAllOf) GetPublicAccessKey() string`

GetPublicAccessKey returns the PublicAccessKey field if non-nil, zero value otherwise.

### GetPublicAccessKeyOk

`func (o *AssetDeviceRegistrationAllOf) GetPublicAccessKeyOk() (*string, bool)`

GetPublicAccessKeyOk returns a tuple with the PublicAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessKey

`func (o *AssetDeviceRegistrationAllOf) SetPublicAccessKey(v string)`

SetPublicAccessKey sets PublicAccessKey field to given value.

### HasPublicAccessKey

`func (o *AssetDeviceRegistrationAllOf) HasPublicAccessKey() bool`

HasPublicAccessKey returns a boolean if a field has been set.

### GetReadOnly

`func (o *AssetDeviceRegistrationAllOf) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AssetDeviceRegistrationAllOf) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AssetDeviceRegistrationAllOf) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AssetDeviceRegistrationAllOf) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSerial

`func (o *AssetDeviceRegistrationAllOf) GetSerial() []string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AssetDeviceRegistrationAllOf) GetSerialOk() (*[]string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AssetDeviceRegistrationAllOf) SetSerial(v []string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AssetDeviceRegistrationAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *AssetDeviceRegistrationAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AssetDeviceRegistrationAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AssetDeviceRegistrationAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AssetDeviceRegistrationAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAccount

`func (o *AssetDeviceRegistrationAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetDeviceRegistrationAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetDeviceRegistrationAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetDeviceRegistrationAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClaimedByUser

`func (o *AssetDeviceRegistrationAllOf) GetClaimedByUser() IamUserRelationship`

GetClaimedByUser returns the ClaimedByUser field if non-nil, zero value otherwise.

### GetClaimedByUserOk

`func (o *AssetDeviceRegistrationAllOf) GetClaimedByUserOk() (*IamUserRelationship, bool)`

GetClaimedByUserOk returns a tuple with the ClaimedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedByUser

`func (o *AssetDeviceRegistrationAllOf) SetClaimedByUser(v IamUserRelationship)`

SetClaimedByUser sets ClaimedByUser field to given value.

### HasClaimedByUser

`func (o *AssetDeviceRegistrationAllOf) HasClaimedByUser() bool`

HasClaimedByUser returns a boolean if a field has been set.

### GetClusterMembers

`func (o *AssetDeviceRegistrationAllOf) GetClusterMembers() []AssetClusterMemberRelationship`

GetClusterMembers returns the ClusterMembers field if non-nil, zero value otherwise.

### GetClusterMembersOk

`func (o *AssetDeviceRegistrationAllOf) GetClusterMembersOk() (*[]AssetClusterMemberRelationship, bool)`

GetClusterMembersOk returns a tuple with the ClusterMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMembers

`func (o *AssetDeviceRegistrationAllOf) SetClusterMembers(v []AssetClusterMemberRelationship)`

SetClusterMembers sets ClusterMembers field to given value.

### HasClusterMembers

`func (o *AssetDeviceRegistrationAllOf) HasClusterMembers() bool`

HasClusterMembers returns a boolean if a field has been set.

### SetClusterMembersNil

`func (o *AssetDeviceRegistrationAllOf) SetClusterMembersNil(b bool)`

 SetClusterMembersNil sets the value for ClusterMembers to be an explicit nil

### UnsetClusterMembers
`func (o *AssetDeviceRegistrationAllOf) UnsetClusterMembers()`

UnsetClusterMembers ensures that no value is present for ClusterMembers, not even an explicit nil
### GetDeviceClaim

`func (o *AssetDeviceRegistrationAllOf) GetDeviceClaim() AssetDeviceClaimRelationship`

GetDeviceClaim returns the DeviceClaim field if non-nil, zero value otherwise.

### GetDeviceClaimOk

`func (o *AssetDeviceRegistrationAllOf) GetDeviceClaimOk() (*AssetDeviceClaimRelationship, bool)`

GetDeviceClaimOk returns a tuple with the DeviceClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClaim

`func (o *AssetDeviceRegistrationAllOf) SetDeviceClaim(v AssetDeviceClaimRelationship)`

SetDeviceClaim sets DeviceClaim field to given value.

### HasDeviceClaim

`func (o *AssetDeviceRegistrationAllOf) HasDeviceClaim() bool`

HasDeviceClaim returns a boolean if a field has been set.

### GetDeviceConfiguration

`func (o *AssetDeviceRegistrationAllOf) GetDeviceConfiguration() AssetDeviceConfigurationRelationship`

GetDeviceConfiguration returns the DeviceConfiguration field if non-nil, zero value otherwise.

### GetDeviceConfigurationOk

`func (o *AssetDeviceRegistrationAllOf) GetDeviceConfigurationOk() (*AssetDeviceConfigurationRelationship, bool)`

GetDeviceConfigurationOk returns a tuple with the DeviceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfiguration

`func (o *AssetDeviceRegistrationAllOf) SetDeviceConfiguration(v AssetDeviceConfigurationRelationship)`

SetDeviceConfiguration sets DeviceConfiguration field to given value.

### HasDeviceConfiguration

`func (o *AssetDeviceRegistrationAllOf) HasDeviceConfiguration() bool`

HasDeviceConfiguration returns a boolean if a field has been set.

### GetDomainGroup

`func (o *AssetDeviceRegistrationAllOf) GetDomainGroup() IamDomainGroupRelationship`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *AssetDeviceRegistrationAllOf) GetDomainGroupOk() (*IamDomainGroupRelationship, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *AssetDeviceRegistrationAllOf) SetDomainGroup(v IamDomainGroupRelationship)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *AssetDeviceRegistrationAllOf) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### GetParentConnection

`func (o *AssetDeviceRegistrationAllOf) GetParentConnection() AssetDeviceRegistrationRelationship`

GetParentConnection returns the ParentConnection field if non-nil, zero value otherwise.

### GetParentConnectionOk

`func (o *AssetDeviceRegistrationAllOf) GetParentConnectionOk() (*AssetDeviceRegistrationRelationship, bool)`

GetParentConnectionOk returns a tuple with the ParentConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnection

`func (o *AssetDeviceRegistrationAllOf) SetParentConnection(v AssetDeviceRegistrationRelationship)`

SetParentConnection sets ParentConnection field to given value.

### HasParentConnection

`func (o *AssetDeviceRegistrationAllOf) HasParentConnection() bool`

HasParentConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


