# ManagementInterfaceRelationship

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
**DeviceMoId** | Pointer to **string** | The database identifier of the registered device of an object. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 
**Gateway** | Pointer to **string** | Default gateway for the interface. | [optional] [readonly] 
**HostName** | Pointer to **string** | Hostname configured for the interface. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the interface. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | IPv4 address of the interface. | [optional] [readonly] 
**Ipv4Gateway** | Pointer to **string** | IPv4 default gateway for the interface. | [optional] [readonly] 
**Ipv4Mask** | Pointer to **string** | IPv4 Netmask for the interface. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | IPv6 address of the interface. | [optional] 
**Ipv6Gateway** | Pointer to **string** | IPv6 default gateway for the interface. | [optional] 
**Ipv6Prefix** | Pointer to **int64** | IPv6 prefix for the interface. | [optional] 
**MacAddress** | Pointer to **string** | MAC address configured for the interface. | [optional] [readonly] 
**Mask** | Pointer to **string** | Netmask for the interface. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Id connected to the interface. | [optional] 
**UemConnStatus** | Pointer to **string** | The event channel connection status for the interface. | [optional] 
**VirtualHostName** | Pointer to **string** | Virtual hostname configured for the interface in case of clustered environment. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**ManagementController** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewManagementInterfaceRelationship

`func NewManagementInterfaceRelationship(classId string, objectType string, ) *ManagementInterfaceRelationship`

NewManagementInterfaceRelationship instantiates a new ManagementInterfaceRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementInterfaceRelationshipWithDefaults

`func NewManagementInterfaceRelationshipWithDefaults() *ManagementInterfaceRelationship`

NewManagementInterfaceRelationshipWithDefaults instantiates a new ManagementInterfaceRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ManagementInterfaceRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ManagementInterfaceRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ManagementInterfaceRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ManagementInterfaceRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ManagementInterfaceRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ManagementInterfaceRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ManagementInterfaceRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ManagementInterfaceRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ManagementInterfaceRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ManagementInterfaceRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ManagementInterfaceRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ManagementInterfaceRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ManagementInterfaceRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ManagementInterfaceRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ManagementInterfaceRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ManagementInterfaceRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ManagementInterfaceRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ManagementInterfaceRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ManagementInterfaceRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ManagementInterfaceRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ManagementInterfaceRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ManagementInterfaceRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ManagementInterfaceRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ManagementInterfaceRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ManagementInterfaceRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ManagementInterfaceRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ManagementInterfaceRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ManagementInterfaceRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ManagementInterfaceRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ManagementInterfaceRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ManagementInterfaceRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ManagementInterfaceRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ManagementInterfaceRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ManagementInterfaceRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ManagementInterfaceRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagementInterfaceRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagementInterfaceRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagementInterfaceRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ManagementInterfaceRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ManagementInterfaceRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ManagementInterfaceRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ManagementInterfaceRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ManagementInterfaceRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ManagementInterfaceRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ManagementInterfaceRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ManagementInterfaceRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ManagementInterfaceRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ManagementInterfaceRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ManagementInterfaceRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ManagementInterfaceRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ManagementInterfaceRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ManagementInterfaceRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ManagementInterfaceRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ManagementInterfaceRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ManagementInterfaceRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ManagementInterfaceRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ManagementInterfaceRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ManagementInterfaceRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ManagementInterfaceRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ManagementInterfaceRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ManagementInterfaceRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ManagementInterfaceRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ManagementInterfaceRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ManagementInterfaceRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *ManagementInterfaceRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ManagementInterfaceRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ManagementInterfaceRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ManagementInterfaceRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *ManagementInterfaceRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ManagementInterfaceRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ManagementInterfaceRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ManagementInterfaceRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *ManagementInterfaceRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *ManagementInterfaceRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *ManagementInterfaceRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *ManagementInterfaceRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetGateway

`func (o *ManagementInterfaceRelationship) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ManagementInterfaceRelationship) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ManagementInterfaceRelationship) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ManagementInterfaceRelationship) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostName

`func (o *ManagementInterfaceRelationship) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ManagementInterfaceRelationship) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ManagementInterfaceRelationship) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ManagementInterfaceRelationship) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpAddress

`func (o *ManagementInterfaceRelationship) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ManagementInterfaceRelationship) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ManagementInterfaceRelationship) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ManagementInterfaceRelationship) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv4Address

`func (o *ManagementInterfaceRelationship) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *ManagementInterfaceRelationship) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *ManagementInterfaceRelationship) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *ManagementInterfaceRelationship) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv4Gateway

`func (o *ManagementInterfaceRelationship) GetIpv4Gateway() string`

GetIpv4Gateway returns the Ipv4Gateway field if non-nil, zero value otherwise.

### GetIpv4GatewayOk

`func (o *ManagementInterfaceRelationship) GetIpv4GatewayOk() (*string, bool)`

GetIpv4GatewayOk returns a tuple with the Ipv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Gateway

`func (o *ManagementInterfaceRelationship) SetIpv4Gateway(v string)`

SetIpv4Gateway sets Ipv4Gateway field to given value.

### HasIpv4Gateway

`func (o *ManagementInterfaceRelationship) HasIpv4Gateway() bool`

HasIpv4Gateway returns a boolean if a field has been set.

### GetIpv4Mask

`func (o *ManagementInterfaceRelationship) GetIpv4Mask() string`

GetIpv4Mask returns the Ipv4Mask field if non-nil, zero value otherwise.

### GetIpv4MaskOk

`func (o *ManagementInterfaceRelationship) GetIpv4MaskOk() (*string, bool)`

GetIpv4MaskOk returns a tuple with the Ipv4Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Mask

`func (o *ManagementInterfaceRelationship) SetIpv4Mask(v string)`

SetIpv4Mask sets Ipv4Mask field to given value.

### HasIpv4Mask

`func (o *ManagementInterfaceRelationship) HasIpv4Mask() bool`

HasIpv4Mask returns a boolean if a field has been set.

### GetIpv6Address

`func (o *ManagementInterfaceRelationship) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *ManagementInterfaceRelationship) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *ManagementInterfaceRelationship) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *ManagementInterfaceRelationship) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetIpv6Gateway

`func (o *ManagementInterfaceRelationship) GetIpv6Gateway() string`

GetIpv6Gateway returns the Ipv6Gateway field if non-nil, zero value otherwise.

### GetIpv6GatewayOk

`func (o *ManagementInterfaceRelationship) GetIpv6GatewayOk() (*string, bool)`

GetIpv6GatewayOk returns a tuple with the Ipv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Gateway

`func (o *ManagementInterfaceRelationship) SetIpv6Gateway(v string)`

SetIpv6Gateway sets Ipv6Gateway field to given value.

### HasIpv6Gateway

`func (o *ManagementInterfaceRelationship) HasIpv6Gateway() bool`

HasIpv6Gateway returns a boolean if a field has been set.

### GetIpv6Prefix

`func (o *ManagementInterfaceRelationship) GetIpv6Prefix() int64`

GetIpv6Prefix returns the Ipv6Prefix field if non-nil, zero value otherwise.

### GetIpv6PrefixOk

`func (o *ManagementInterfaceRelationship) GetIpv6PrefixOk() (*int64, bool)`

GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefix

`func (o *ManagementInterfaceRelationship) SetIpv6Prefix(v int64)`

SetIpv6Prefix sets Ipv6Prefix field to given value.

### HasIpv6Prefix

`func (o *ManagementInterfaceRelationship) HasIpv6Prefix() bool`

HasIpv6Prefix returns a boolean if a field has been set.

### GetMacAddress

`func (o *ManagementInterfaceRelationship) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ManagementInterfaceRelationship) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ManagementInterfaceRelationship) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ManagementInterfaceRelationship) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMask

`func (o *ManagementInterfaceRelationship) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *ManagementInterfaceRelationship) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *ManagementInterfaceRelationship) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *ManagementInterfaceRelationship) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetSwitchId

`func (o *ManagementInterfaceRelationship) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *ManagementInterfaceRelationship) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *ManagementInterfaceRelationship) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *ManagementInterfaceRelationship) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetUemConnStatus

`func (o *ManagementInterfaceRelationship) GetUemConnStatus() string`

GetUemConnStatus returns the UemConnStatus field if non-nil, zero value otherwise.

### GetUemConnStatusOk

`func (o *ManagementInterfaceRelationship) GetUemConnStatusOk() (*string, bool)`

GetUemConnStatusOk returns a tuple with the UemConnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUemConnStatus

`func (o *ManagementInterfaceRelationship) SetUemConnStatus(v string)`

SetUemConnStatus sets UemConnStatus field to given value.

### HasUemConnStatus

`func (o *ManagementInterfaceRelationship) HasUemConnStatus() bool`

HasUemConnStatus returns a boolean if a field has been set.

### GetVirtualHostName

`func (o *ManagementInterfaceRelationship) GetVirtualHostName() string`

GetVirtualHostName returns the VirtualHostName field if non-nil, zero value otherwise.

### GetVirtualHostNameOk

`func (o *ManagementInterfaceRelationship) GetVirtualHostNameOk() (*string, bool)`

GetVirtualHostNameOk returns a tuple with the VirtualHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHostName

`func (o *ManagementInterfaceRelationship) SetVirtualHostName(v string)`

SetVirtualHostName sets VirtualHostName field to given value.

### HasVirtualHostName

`func (o *ManagementInterfaceRelationship) HasVirtualHostName() bool`

HasVirtualHostName returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ManagementInterfaceRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ManagementInterfaceRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ManagementInterfaceRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ManagementInterfaceRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementController

`func (o *ManagementInterfaceRelationship) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *ManagementInterfaceRelationship) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *ManagementInterfaceRelationship) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *ManagementInterfaceRelationship) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ManagementInterfaceRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ManagementInterfaceRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ManagementInterfaceRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ManagementInterfaceRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


