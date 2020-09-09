# AssetClusterMemberRelationship

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
**Leadership** | Pointer to **string** | The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection. * &#x60;Unknown&#x60; - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight. * &#x60;Primary&#x60; - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary in a given cluster, while the underlying platform may be active-active only one connector will assume the primary role. * &#x60;Secondary&#x60; - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. e.g. file upload will be enabled, but requests to the underlying platform management will be disabled. | [optional] [readonly] [default to "Unknown"]
**LockedLeader** | Pointer to **bool** | Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true. | [optional] 
**MemberIdentity** | Pointer to **string** | The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time. | [optional] [readonly] 
**ParentClusterMemberIdentity** | Pointer to **string** | The member idenity of the cluster member through which this device is connected if applicable. | [optional] [readonly] 
**Sudi** | Pointer to [**AssetSudiInfo**](asset.SudiInfo.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetClusterMemberRelationship

`func NewAssetClusterMemberRelationship(classId string, objectType string, ) *AssetClusterMemberRelationship`

NewAssetClusterMemberRelationship instantiates a new AssetClusterMemberRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetClusterMemberRelationshipWithDefaults

`func NewAssetClusterMemberRelationshipWithDefaults() *AssetClusterMemberRelationship`

NewAssetClusterMemberRelationshipWithDefaults instantiates a new AssetClusterMemberRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *AssetClusterMemberRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *AssetClusterMemberRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *AssetClusterMemberRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *AssetClusterMemberRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *AssetClusterMemberRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetClusterMemberRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetClusterMemberRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *AssetClusterMemberRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AssetClusterMemberRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AssetClusterMemberRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AssetClusterMemberRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *AssetClusterMemberRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *AssetClusterMemberRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *AssetClusterMemberRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *AssetClusterMemberRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *AssetClusterMemberRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *AssetClusterMemberRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *AssetClusterMemberRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *AssetClusterMemberRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *AssetClusterMemberRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *AssetClusterMemberRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *AssetClusterMemberRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *AssetClusterMemberRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *AssetClusterMemberRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetClusterMemberRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetClusterMemberRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *AssetClusterMemberRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AssetClusterMemberRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AssetClusterMemberRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *AssetClusterMemberRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *AssetClusterMemberRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *AssetClusterMemberRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *AssetClusterMemberRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *AssetClusterMemberRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *AssetClusterMemberRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetClusterMemberRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetClusterMemberRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetClusterMemberRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *AssetClusterMemberRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *AssetClusterMemberRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *AssetClusterMemberRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *AssetClusterMemberRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *AssetClusterMemberRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *AssetClusterMemberRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *AssetClusterMemberRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *AssetClusterMemberRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *AssetClusterMemberRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *AssetClusterMemberRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *AssetClusterMemberRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AssetClusterMemberRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AssetClusterMemberRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AssetClusterMemberRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *AssetClusterMemberRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *AssetClusterMemberRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *AssetClusterMemberRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *AssetClusterMemberRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *AssetClusterMemberRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *AssetClusterMemberRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *AssetClusterMemberRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *AssetClusterMemberRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *AssetClusterMemberRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *AssetClusterMemberRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *AssetClusterMemberRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *AssetClusterMemberRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetApiVersion

`func (o *AssetClusterMemberRelationship) GetApiVersion() int64`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AssetClusterMemberRelationship) GetApiVersionOk() (*int64, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AssetClusterMemberRelationship) SetApiVersion(v int64)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AssetClusterMemberRelationship) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAppPartitionNumber

`func (o *AssetClusterMemberRelationship) GetAppPartitionNumber() int64`

GetAppPartitionNumber returns the AppPartitionNumber field if non-nil, zero value otherwise.

### GetAppPartitionNumberOk

`func (o *AssetClusterMemberRelationship) GetAppPartitionNumberOk() (*int64, bool)`

GetAppPartitionNumberOk returns a tuple with the AppPartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPartitionNumber

`func (o *AssetClusterMemberRelationship) SetAppPartitionNumber(v int64)`

SetAppPartitionNumber sets AppPartitionNumber field to given value.

### HasAppPartitionNumber

`func (o *AssetClusterMemberRelationship) HasAppPartitionNumber() bool`

HasAppPartitionNumber returns a boolean if a field has been set.

### GetConnectionId

`func (o *AssetClusterMemberRelationship) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AssetClusterMemberRelationship) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AssetClusterMemberRelationship) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AssetClusterMemberRelationship) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionReason

`func (o *AssetClusterMemberRelationship) GetConnectionReason() string`

GetConnectionReason returns the ConnectionReason field if non-nil, zero value otherwise.

### GetConnectionReasonOk

`func (o *AssetClusterMemberRelationship) GetConnectionReasonOk() (*string, bool)`

GetConnectionReasonOk returns a tuple with the ConnectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionReason

`func (o *AssetClusterMemberRelationship) SetConnectionReason(v string)`

SetConnectionReason sets ConnectionReason field to given value.

### HasConnectionReason

`func (o *AssetClusterMemberRelationship) HasConnectionReason() bool`

HasConnectionReason returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AssetClusterMemberRelationship) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AssetClusterMemberRelationship) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AssetClusterMemberRelationship) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AssetClusterMemberRelationship) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetConnectionStatusLastChangeTime

`func (o *AssetClusterMemberRelationship) GetConnectionStatusLastChangeTime() time.Time`

GetConnectionStatusLastChangeTime returns the ConnectionStatusLastChangeTime field if non-nil, zero value otherwise.

### GetConnectionStatusLastChangeTimeOk

`func (o *AssetClusterMemberRelationship) GetConnectionStatusLastChangeTimeOk() (*time.Time, bool)`

GetConnectionStatusLastChangeTimeOk returns a tuple with the ConnectionStatusLastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatusLastChangeTime

`func (o *AssetClusterMemberRelationship) SetConnectionStatusLastChangeTime(v time.Time)`

SetConnectionStatusLastChangeTime sets ConnectionStatusLastChangeTime field to given value.

### HasConnectionStatusLastChangeTime

`func (o *AssetClusterMemberRelationship) HasConnectionStatusLastChangeTime() bool`

HasConnectionStatusLastChangeTime returns a boolean if a field has been set.

### GetConnectorVersion

`func (o *AssetClusterMemberRelationship) GetConnectorVersion() string`

GetConnectorVersion returns the ConnectorVersion field if non-nil, zero value otherwise.

### GetConnectorVersionOk

`func (o *AssetClusterMemberRelationship) GetConnectorVersionOk() (*string, bool)`

GetConnectorVersionOk returns a tuple with the ConnectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorVersion

`func (o *AssetClusterMemberRelationship) SetConnectorVersion(v string)`

SetConnectorVersion sets ConnectorVersion field to given value.

### HasConnectorVersion

`func (o *AssetClusterMemberRelationship) HasConnectorVersion() bool`

HasConnectorVersion returns a boolean if a field has been set.

### GetDeviceExternalIpAddress

`func (o *AssetClusterMemberRelationship) GetDeviceExternalIpAddress() string`

GetDeviceExternalIpAddress returns the DeviceExternalIpAddress field if non-nil, zero value otherwise.

### GetDeviceExternalIpAddressOk

`func (o *AssetClusterMemberRelationship) GetDeviceExternalIpAddressOk() (*string, bool)`

GetDeviceExternalIpAddressOk returns a tuple with the DeviceExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceExternalIpAddress

`func (o *AssetClusterMemberRelationship) SetDeviceExternalIpAddress(v string)`

SetDeviceExternalIpAddress sets DeviceExternalIpAddress field to given value.

### HasDeviceExternalIpAddress

`func (o *AssetClusterMemberRelationship) HasDeviceExternalIpAddress() bool`

HasDeviceExternalIpAddress returns a boolean if a field has been set.

### GetProxyApp

`func (o *AssetClusterMemberRelationship) GetProxyApp() string`

GetProxyApp returns the ProxyApp field if non-nil, zero value otherwise.

### GetProxyAppOk

`func (o *AssetClusterMemberRelationship) GetProxyAppOk() (*string, bool)`

GetProxyAppOk returns a tuple with the ProxyApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyApp

`func (o *AssetClusterMemberRelationship) SetProxyApp(v string)`

SetProxyApp sets ProxyApp field to given value.

### HasProxyApp

`func (o *AssetClusterMemberRelationship) HasProxyApp() bool`

HasProxyApp returns a boolean if a field has been set.

### GetLeadership

`func (o *AssetClusterMemberRelationship) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *AssetClusterMemberRelationship) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *AssetClusterMemberRelationship) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *AssetClusterMemberRelationship) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetLockedLeader

`func (o *AssetClusterMemberRelationship) GetLockedLeader() bool`

GetLockedLeader returns the LockedLeader field if non-nil, zero value otherwise.

### GetLockedLeaderOk

`func (o *AssetClusterMemberRelationship) GetLockedLeaderOk() (*bool, bool)`

GetLockedLeaderOk returns a tuple with the LockedLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedLeader

`func (o *AssetClusterMemberRelationship) SetLockedLeader(v bool)`

SetLockedLeader sets LockedLeader field to given value.

### HasLockedLeader

`func (o *AssetClusterMemberRelationship) HasLockedLeader() bool`

HasLockedLeader returns a boolean if a field has been set.

### GetMemberIdentity

`func (o *AssetClusterMemberRelationship) GetMemberIdentity() string`

GetMemberIdentity returns the MemberIdentity field if non-nil, zero value otherwise.

### GetMemberIdentityOk

`func (o *AssetClusterMemberRelationship) GetMemberIdentityOk() (*string, bool)`

GetMemberIdentityOk returns a tuple with the MemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIdentity

`func (o *AssetClusterMemberRelationship) SetMemberIdentity(v string)`

SetMemberIdentity sets MemberIdentity field to given value.

### HasMemberIdentity

`func (o *AssetClusterMemberRelationship) HasMemberIdentity() bool`

HasMemberIdentity returns a boolean if a field has been set.

### GetParentClusterMemberIdentity

`func (o *AssetClusterMemberRelationship) GetParentClusterMemberIdentity() string`

GetParentClusterMemberIdentity returns the ParentClusterMemberIdentity field if non-nil, zero value otherwise.

### GetParentClusterMemberIdentityOk

`func (o *AssetClusterMemberRelationship) GetParentClusterMemberIdentityOk() (*string, bool)`

GetParentClusterMemberIdentityOk returns a tuple with the ParentClusterMemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentClusterMemberIdentity

`func (o *AssetClusterMemberRelationship) SetParentClusterMemberIdentity(v string)`

SetParentClusterMemberIdentity sets ParentClusterMemberIdentity field to given value.

### HasParentClusterMemberIdentity

`func (o *AssetClusterMemberRelationship) HasParentClusterMemberIdentity() bool`

HasParentClusterMemberIdentity returns a boolean if a field has been set.

### GetSudi

`func (o *AssetClusterMemberRelationship) GetSudi() AssetSudiInfo`

GetSudi returns the Sudi field if non-nil, zero value otherwise.

### GetSudiOk

`func (o *AssetClusterMemberRelationship) GetSudiOk() (*AssetSudiInfo, bool)`

GetSudiOk returns a tuple with the Sudi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudi

`func (o *AssetClusterMemberRelationship) SetSudi(v AssetSudiInfo)`

SetSudi sets Sudi field to given value.

### HasSudi

`func (o *AssetClusterMemberRelationship) HasSudi() bool`

HasSudi returns a boolean if a field has been set.

### GetDevice

`func (o *AssetClusterMemberRelationship) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetClusterMemberRelationship) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetClusterMemberRelationship) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetClusterMemberRelationship) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


