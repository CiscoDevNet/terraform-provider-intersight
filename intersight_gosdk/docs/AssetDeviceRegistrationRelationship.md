# AssetDeviceRegistrationRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**ApiVersion** | Pointer to **int64** | The version of the connector API, describes the capability of the connector&#39;s framework. If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded. | [optional] [readonly] 
**AppPartitionNumber** | Pointer to **int64** | The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter. | [optional] [readonly] 
**ConnectionReason** | Pointer to **string** | If &#39;connectionStatus&#39; is not equal to Connected, connectionReason provides further details about why the device is not connected with Intersight. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | The status of the persistent connection between the device connector and Intersight. * &#x60;&#x60; - The device registered with Intersight but subsequently did not establish a persistent websocket connection. * &#x60;Connected&#x60; - The device&#39;s connection to Intersight has been established and is active. * &#x60;NotConnected&#x60; - The device&#39;s connection to Intersight has been disconnected. * &#x60;ClaimInProgress&#x60; - Claim of the device is in progress. * &#x60;Unclaimed&#x60; - The device was un-claimed from the users account by an Administrator of the device. | [optional] [readonly] [default to ""]
**ConnectionStatusLastChangeTime** | Pointer to [**time.Time**](time.Time.md) | The last time at which the &#39;connectionStatus&#39; property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between Intersight and Device Connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with Intersight. | [optional] [readonly] 
**ConnectorVersion** | Pointer to **string** | The version of the device connector running on the managed device. | [optional] [readonly] 
**DeviceExternalIpAddress** | Pointer to **string** | The IP Address of the managed device as seen from Intersight at the time of registration. This could be the IP address of the managed device&#39;s interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network. | [optional] [readonly] 
**ProxyApp** | Pointer to **string** | The name of the app which will proxy the messages to the device connector. | [optional] [readonly] 
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

### NewAssetDeviceRegistrationRelationship

`func NewAssetDeviceRegistrationRelationship(classId string, objectType string, ) *AssetDeviceRegistrationRelationship`

NewAssetDeviceRegistrationRelationship instantiates a new AssetDeviceRegistrationRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceRegistrationRelationshipWithDefaults

`func NewAssetDeviceRegistrationRelationshipWithDefaults() *AssetDeviceRegistrationRelationship`

NewAssetDeviceRegistrationRelationshipWithDefaults instantiates a new AssetDeviceRegistrationRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *AssetDeviceRegistrationRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *AssetDeviceRegistrationRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *AssetDeviceRegistrationRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *AssetDeviceRegistrationRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *AssetDeviceRegistrationRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceRegistrationRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceRegistrationRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *AssetDeviceRegistrationRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AssetDeviceRegistrationRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AssetDeviceRegistrationRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AssetDeviceRegistrationRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *AssetDeviceRegistrationRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *AssetDeviceRegistrationRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *AssetDeviceRegistrationRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *AssetDeviceRegistrationRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *AssetDeviceRegistrationRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *AssetDeviceRegistrationRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *AssetDeviceRegistrationRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *AssetDeviceRegistrationRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *AssetDeviceRegistrationRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *AssetDeviceRegistrationRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *AssetDeviceRegistrationRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *AssetDeviceRegistrationRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *AssetDeviceRegistrationRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceRegistrationRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceRegistrationRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *AssetDeviceRegistrationRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AssetDeviceRegistrationRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AssetDeviceRegistrationRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *AssetDeviceRegistrationRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *AssetDeviceRegistrationRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *AssetDeviceRegistrationRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *AssetDeviceRegistrationRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *AssetDeviceRegistrationRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *AssetDeviceRegistrationRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetDeviceRegistrationRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetDeviceRegistrationRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetDeviceRegistrationRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *AssetDeviceRegistrationRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *AssetDeviceRegistrationRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *AssetDeviceRegistrationRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *AssetDeviceRegistrationRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *AssetDeviceRegistrationRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *AssetDeviceRegistrationRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *AssetDeviceRegistrationRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *AssetDeviceRegistrationRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *AssetDeviceRegistrationRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *AssetDeviceRegistrationRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *AssetDeviceRegistrationRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AssetDeviceRegistrationRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AssetDeviceRegistrationRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AssetDeviceRegistrationRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *AssetDeviceRegistrationRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *AssetDeviceRegistrationRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *AssetDeviceRegistrationRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *AssetDeviceRegistrationRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *AssetDeviceRegistrationRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *AssetDeviceRegistrationRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *AssetDeviceRegistrationRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *AssetDeviceRegistrationRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *AssetDeviceRegistrationRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *AssetDeviceRegistrationRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *AssetDeviceRegistrationRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *AssetDeviceRegistrationRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetApiVersion

`func (o *AssetDeviceRegistrationRelationship) GetApiVersion() int64`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AssetDeviceRegistrationRelationship) GetApiVersionOk() (*int64, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AssetDeviceRegistrationRelationship) SetApiVersion(v int64)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AssetDeviceRegistrationRelationship) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAppPartitionNumber

`func (o *AssetDeviceRegistrationRelationship) GetAppPartitionNumber() int64`

GetAppPartitionNumber returns the AppPartitionNumber field if non-nil, zero value otherwise.

### GetAppPartitionNumberOk

`func (o *AssetDeviceRegistrationRelationship) GetAppPartitionNumberOk() (*int64, bool)`

GetAppPartitionNumberOk returns a tuple with the AppPartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPartitionNumber

`func (o *AssetDeviceRegistrationRelationship) SetAppPartitionNumber(v int64)`

SetAppPartitionNumber sets AppPartitionNumber field to given value.

### HasAppPartitionNumber

`func (o *AssetDeviceRegistrationRelationship) HasAppPartitionNumber() bool`

HasAppPartitionNumber returns a boolean if a field has been set.

### GetConnectionId

`func (o *AssetDeviceRegistrationRelationship) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AssetDeviceRegistrationRelationship) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AssetDeviceRegistrationRelationship) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AssetDeviceRegistrationRelationship) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionReason

`func (o *AssetDeviceRegistrationRelationship) GetConnectionReason() string`

GetConnectionReason returns the ConnectionReason field if non-nil, zero value otherwise.

### GetConnectionReasonOk

`func (o *AssetDeviceRegistrationRelationship) GetConnectionReasonOk() (*string, bool)`

GetConnectionReasonOk returns a tuple with the ConnectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionReason

`func (o *AssetDeviceRegistrationRelationship) SetConnectionReason(v string)`

SetConnectionReason sets ConnectionReason field to given value.

### HasConnectionReason

`func (o *AssetDeviceRegistrationRelationship) HasConnectionReason() bool`

HasConnectionReason returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AssetDeviceRegistrationRelationship) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AssetDeviceRegistrationRelationship) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AssetDeviceRegistrationRelationship) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AssetDeviceRegistrationRelationship) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetConnectionStatusLastChangeTime

`func (o *AssetDeviceRegistrationRelationship) GetConnectionStatusLastChangeTime() time.Time`

GetConnectionStatusLastChangeTime returns the ConnectionStatusLastChangeTime field if non-nil, zero value otherwise.

### GetConnectionStatusLastChangeTimeOk

`func (o *AssetDeviceRegistrationRelationship) GetConnectionStatusLastChangeTimeOk() (*time.Time, bool)`

GetConnectionStatusLastChangeTimeOk returns a tuple with the ConnectionStatusLastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatusLastChangeTime

`func (o *AssetDeviceRegistrationRelationship) SetConnectionStatusLastChangeTime(v time.Time)`

SetConnectionStatusLastChangeTime sets ConnectionStatusLastChangeTime field to given value.

### HasConnectionStatusLastChangeTime

`func (o *AssetDeviceRegistrationRelationship) HasConnectionStatusLastChangeTime() bool`

HasConnectionStatusLastChangeTime returns a boolean if a field has been set.

### GetConnectorVersion

`func (o *AssetDeviceRegistrationRelationship) GetConnectorVersion() string`

GetConnectorVersion returns the ConnectorVersion field if non-nil, zero value otherwise.

### GetConnectorVersionOk

`func (o *AssetDeviceRegistrationRelationship) GetConnectorVersionOk() (*string, bool)`

GetConnectorVersionOk returns a tuple with the ConnectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorVersion

`func (o *AssetDeviceRegistrationRelationship) SetConnectorVersion(v string)`

SetConnectorVersion sets ConnectorVersion field to given value.

### HasConnectorVersion

`func (o *AssetDeviceRegistrationRelationship) HasConnectorVersion() bool`

HasConnectorVersion returns a boolean if a field has been set.

### GetDeviceExternalIpAddress

`func (o *AssetDeviceRegistrationRelationship) GetDeviceExternalIpAddress() string`

GetDeviceExternalIpAddress returns the DeviceExternalIpAddress field if non-nil, zero value otherwise.

### GetDeviceExternalIpAddressOk

`func (o *AssetDeviceRegistrationRelationship) GetDeviceExternalIpAddressOk() (*string, bool)`

GetDeviceExternalIpAddressOk returns a tuple with the DeviceExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceExternalIpAddress

`func (o *AssetDeviceRegistrationRelationship) SetDeviceExternalIpAddress(v string)`

SetDeviceExternalIpAddress sets DeviceExternalIpAddress field to given value.

### HasDeviceExternalIpAddress

`func (o *AssetDeviceRegistrationRelationship) HasDeviceExternalIpAddress() bool`

HasDeviceExternalIpAddress returns a boolean if a field has been set.

### GetProxyApp

`func (o *AssetDeviceRegistrationRelationship) GetProxyApp() string`

GetProxyApp returns the ProxyApp field if non-nil, zero value otherwise.

### GetProxyAppOk

`func (o *AssetDeviceRegistrationRelationship) GetProxyAppOk() (*string, bool)`

GetProxyAppOk returns a tuple with the ProxyApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyApp

`func (o *AssetDeviceRegistrationRelationship) SetProxyApp(v string)`

SetProxyApp sets ProxyApp field to given value.

### HasProxyApp

`func (o *AssetDeviceRegistrationRelationship) HasProxyApp() bool`

HasProxyApp returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *AssetDeviceRegistrationRelationship) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AssetDeviceRegistrationRelationship) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AssetDeviceRegistrationRelationship) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AssetDeviceRegistrationRelationship) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetClaimedByUserName

`func (o *AssetDeviceRegistrationRelationship) GetClaimedByUserName() string`

GetClaimedByUserName returns the ClaimedByUserName field if non-nil, zero value otherwise.

### GetClaimedByUserNameOk

`func (o *AssetDeviceRegistrationRelationship) GetClaimedByUserNameOk() (*string, bool)`

GetClaimedByUserNameOk returns a tuple with the ClaimedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedByUserName

`func (o *AssetDeviceRegistrationRelationship) SetClaimedByUserName(v string)`

SetClaimedByUserName sets ClaimedByUserName field to given value.

### HasClaimedByUserName

`func (o *AssetDeviceRegistrationRelationship) HasClaimedByUserName() bool`

HasClaimedByUserName returns a boolean if a field has been set.

### GetClaimedTime

`func (o *AssetDeviceRegistrationRelationship) GetClaimedTime() time.Time`

GetClaimedTime returns the ClaimedTime field if non-nil, zero value otherwise.

### GetClaimedTimeOk

`func (o *AssetDeviceRegistrationRelationship) GetClaimedTimeOk() (*time.Time, bool)`

GetClaimedTimeOk returns a tuple with the ClaimedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedTime

`func (o *AssetDeviceRegistrationRelationship) SetClaimedTime(v time.Time)`

SetClaimedTime sets ClaimedTime field to given value.

### HasClaimedTime

`func (o *AssetDeviceRegistrationRelationship) HasClaimedTime() bool`

HasClaimedTime returns a boolean if a field has been set.

### GetDeviceHostname

`func (o *AssetDeviceRegistrationRelationship) GetDeviceHostname() []string`

GetDeviceHostname returns the DeviceHostname field if non-nil, zero value otherwise.

### GetDeviceHostnameOk

`func (o *AssetDeviceRegistrationRelationship) GetDeviceHostnameOk() (*[]string, bool)`

GetDeviceHostnameOk returns a tuple with the DeviceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHostname

`func (o *AssetDeviceRegistrationRelationship) SetDeviceHostname(v []string)`

SetDeviceHostname sets DeviceHostname field to given value.

### HasDeviceHostname

`func (o *AssetDeviceRegistrationRelationship) HasDeviceHostname() bool`

HasDeviceHostname returns a boolean if a field has been set.

### GetDeviceIpAddress

`func (o *AssetDeviceRegistrationRelationship) GetDeviceIpAddress() []string`

GetDeviceIpAddress returns the DeviceIpAddress field if non-nil, zero value otherwise.

### GetDeviceIpAddressOk

`func (o *AssetDeviceRegistrationRelationship) GetDeviceIpAddressOk() (*[]string, bool)`

GetDeviceIpAddressOk returns a tuple with the DeviceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIpAddress

`func (o *AssetDeviceRegistrationRelationship) SetDeviceIpAddress(v []string)`

SetDeviceIpAddress sets DeviceIpAddress field to given value.

### HasDeviceIpAddress

`func (o *AssetDeviceRegistrationRelationship) HasDeviceIpAddress() bool`

HasDeviceIpAddress returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AssetDeviceRegistrationRelationship) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AssetDeviceRegistrationRelationship) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AssetDeviceRegistrationRelationship) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AssetDeviceRegistrationRelationship) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetParentSignature

`func (o *AssetDeviceRegistrationRelationship) GetParentSignature() AssetParentConnectionSignature`

GetParentSignature returns the ParentSignature field if non-nil, zero value otherwise.

### GetParentSignatureOk

`func (o *AssetDeviceRegistrationRelationship) GetParentSignatureOk() (*AssetParentConnectionSignature, bool)`

GetParentSignatureOk returns a tuple with the ParentSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSignature

`func (o *AssetDeviceRegistrationRelationship) SetParentSignature(v AssetParentConnectionSignature)`

SetParentSignature sets ParentSignature field to given value.

### HasParentSignature

`func (o *AssetDeviceRegistrationRelationship) HasParentSignature() bool`

HasParentSignature returns a boolean if a field has been set.

### GetPid

`func (o *AssetDeviceRegistrationRelationship) GetPid() []string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *AssetDeviceRegistrationRelationship) GetPidOk() (*[]string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *AssetDeviceRegistrationRelationship) SetPid(v []string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *AssetDeviceRegistrationRelationship) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformType

`func (o *AssetDeviceRegistrationRelationship) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AssetDeviceRegistrationRelationship) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AssetDeviceRegistrationRelationship) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AssetDeviceRegistrationRelationship) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPublicAccessKey

`func (o *AssetDeviceRegistrationRelationship) GetPublicAccessKey() string`

GetPublicAccessKey returns the PublicAccessKey field if non-nil, zero value otherwise.

### GetPublicAccessKeyOk

`func (o *AssetDeviceRegistrationRelationship) GetPublicAccessKeyOk() (*string, bool)`

GetPublicAccessKeyOk returns a tuple with the PublicAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessKey

`func (o *AssetDeviceRegistrationRelationship) SetPublicAccessKey(v string)`

SetPublicAccessKey sets PublicAccessKey field to given value.

### HasPublicAccessKey

`func (o *AssetDeviceRegistrationRelationship) HasPublicAccessKey() bool`

HasPublicAccessKey returns a boolean if a field has been set.

### GetReadOnly

`func (o *AssetDeviceRegistrationRelationship) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AssetDeviceRegistrationRelationship) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AssetDeviceRegistrationRelationship) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AssetDeviceRegistrationRelationship) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSerial

`func (o *AssetDeviceRegistrationRelationship) GetSerial() []string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AssetDeviceRegistrationRelationship) GetSerialOk() (*[]string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AssetDeviceRegistrationRelationship) SetSerial(v []string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AssetDeviceRegistrationRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *AssetDeviceRegistrationRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AssetDeviceRegistrationRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AssetDeviceRegistrationRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AssetDeviceRegistrationRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAccount

`func (o *AssetDeviceRegistrationRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetDeviceRegistrationRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetDeviceRegistrationRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetDeviceRegistrationRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClaimedByUser

`func (o *AssetDeviceRegistrationRelationship) GetClaimedByUser() IamUserRelationship`

GetClaimedByUser returns the ClaimedByUser field if non-nil, zero value otherwise.

### GetClaimedByUserOk

`func (o *AssetDeviceRegistrationRelationship) GetClaimedByUserOk() (*IamUserRelationship, bool)`

GetClaimedByUserOk returns a tuple with the ClaimedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedByUser

`func (o *AssetDeviceRegistrationRelationship) SetClaimedByUser(v IamUserRelationship)`

SetClaimedByUser sets ClaimedByUser field to given value.

### HasClaimedByUser

`func (o *AssetDeviceRegistrationRelationship) HasClaimedByUser() bool`

HasClaimedByUser returns a boolean if a field has been set.

### GetClusterMembers

`func (o *AssetDeviceRegistrationRelationship) GetClusterMembers() []AssetClusterMemberRelationship`

GetClusterMembers returns the ClusterMembers field if non-nil, zero value otherwise.

### GetClusterMembersOk

`func (o *AssetDeviceRegistrationRelationship) GetClusterMembersOk() (*[]AssetClusterMemberRelationship, bool)`

GetClusterMembersOk returns a tuple with the ClusterMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMembers

`func (o *AssetDeviceRegistrationRelationship) SetClusterMembers(v []AssetClusterMemberRelationship)`

SetClusterMembers sets ClusterMembers field to given value.

### HasClusterMembers

`func (o *AssetDeviceRegistrationRelationship) HasClusterMembers() bool`

HasClusterMembers returns a boolean if a field has been set.

### SetClusterMembersNil

`func (o *AssetDeviceRegistrationRelationship) SetClusterMembersNil(b bool)`

 SetClusterMembersNil sets the value for ClusterMembers to be an explicit nil

### UnsetClusterMembers
`func (o *AssetDeviceRegistrationRelationship) UnsetClusterMembers()`

UnsetClusterMembers ensures that no value is present for ClusterMembers, not even an explicit nil
### GetDeviceClaim

`func (o *AssetDeviceRegistrationRelationship) GetDeviceClaim() AssetDeviceClaimRelationship`

GetDeviceClaim returns the DeviceClaim field if non-nil, zero value otherwise.

### GetDeviceClaimOk

`func (o *AssetDeviceRegistrationRelationship) GetDeviceClaimOk() (*AssetDeviceClaimRelationship, bool)`

GetDeviceClaimOk returns a tuple with the DeviceClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClaim

`func (o *AssetDeviceRegistrationRelationship) SetDeviceClaim(v AssetDeviceClaimRelationship)`

SetDeviceClaim sets DeviceClaim field to given value.

### HasDeviceClaim

`func (o *AssetDeviceRegistrationRelationship) HasDeviceClaim() bool`

HasDeviceClaim returns a boolean if a field has been set.

### GetDeviceConfiguration

`func (o *AssetDeviceRegistrationRelationship) GetDeviceConfiguration() AssetDeviceConfigurationRelationship`

GetDeviceConfiguration returns the DeviceConfiguration field if non-nil, zero value otherwise.

### GetDeviceConfigurationOk

`func (o *AssetDeviceRegistrationRelationship) GetDeviceConfigurationOk() (*AssetDeviceConfigurationRelationship, bool)`

GetDeviceConfigurationOk returns a tuple with the DeviceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfiguration

`func (o *AssetDeviceRegistrationRelationship) SetDeviceConfiguration(v AssetDeviceConfigurationRelationship)`

SetDeviceConfiguration sets DeviceConfiguration field to given value.

### HasDeviceConfiguration

`func (o *AssetDeviceRegistrationRelationship) HasDeviceConfiguration() bool`

HasDeviceConfiguration returns a boolean if a field has been set.

### GetDomainGroup

`func (o *AssetDeviceRegistrationRelationship) GetDomainGroup() IamDomainGroupRelationship`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *AssetDeviceRegistrationRelationship) GetDomainGroupOk() (*IamDomainGroupRelationship, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *AssetDeviceRegistrationRelationship) SetDomainGroup(v IamDomainGroupRelationship)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *AssetDeviceRegistrationRelationship) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### GetParentConnection

`func (o *AssetDeviceRegistrationRelationship) GetParentConnection() AssetDeviceRegistrationRelationship`

GetParentConnection returns the ParentConnection field if non-nil, zero value otherwise.

### GetParentConnectionOk

`func (o *AssetDeviceRegistrationRelationship) GetParentConnectionOk() (*AssetDeviceRegistrationRelationship, bool)`

GetParentConnectionOk returns a tuple with the ParentConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnection

`func (o *AssetDeviceRegistrationRelationship) SetParentConnection(v AssetDeviceRegistrationRelationship)`

SetParentConnection sets ParentConnection field to given value.

### HasParentConnection

`func (o *AssetDeviceRegistrationRelationship) HasParentConnection() bool`

HasParentConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


