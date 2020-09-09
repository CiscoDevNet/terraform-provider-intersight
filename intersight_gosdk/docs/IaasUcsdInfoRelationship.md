# IaasUcsdInfoRelationship

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
**DeviceId** | Pointer to **string** | Moid of the UCS Director device connector&#39;s asset.DeviceRegistration. | [optional] [readonly] 
**Guid** | Pointer to **string** | Unique ID of UCS Director getting registerd with Intersight. | [optional] [readonly] 
**HostName** | Pointer to **string** | The UCS Director hostname for management. | [optional] [readonly] 
**Ip** | Pointer to **string** | The UCS Director IP address for management. | [optional] [readonly] 
**LastBackup** | Pointer to [**time.Time**](time.Time.md) | Last successful backup created for this UCS Director appliance if backup is configured. | [optional] [readonly] 
**NodeType** | Pointer to **string** | NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node. | [optional] [readonly] 
**ProductName** | Pointer to **string** | The UCS Director product name. | [optional] [readonly] 
**ProductVendor** | Pointer to **string** | The UCS Director product vendor. | [optional] [readonly] 
**ProductVersion** | Pointer to **string** | The UCS Director product/platform version. | [optional] [readonly] 
**Status** | Pointer to **string** | The UCS Director status. Possible values are Active, Inactive, Unknown. | [optional] [readonly] 
**ConnectorPack** | Pointer to [**[]IaasConnectorPackRelationship**](iaas.ConnectorPack.Relationship.md) | An array of relationships to iaasConnectorPack resources. | [optional] [readonly] 
**DeviceStatus** | Pointer to [**[]IaasDeviceStatusRelationship**](iaas.DeviceStatus.Relationship.md) | An array of relationships to iaasDeviceStatus resources. | [optional] [readonly] 
**LicenseInfo** | Pointer to [**IaasLicenseInfoRelationship**](iaas.LicenseInfo.Relationship.md) |  | [optional] 
**MostRunTasks** | Pointer to [**[]IaasMostRunTasksRelationship**](iaas.MostRunTasks.Relationship.md) | An array of relationships to iaasMostRunTasks resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**UcsdManagedInfra** | Pointer to [**IaasUcsdManagedInfraRelationship**](iaas.UcsdManagedInfra.Relationship.md) |  | [optional] 

## Methods

### NewIaasUcsdInfoRelationship

`func NewIaasUcsdInfoRelationship(classId string, objectType string, ) *IaasUcsdInfoRelationship`

NewIaasUcsdInfoRelationship instantiates a new IaasUcsdInfoRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasUcsdInfoRelationshipWithDefaults

`func NewIaasUcsdInfoRelationshipWithDefaults() *IaasUcsdInfoRelationship`

NewIaasUcsdInfoRelationshipWithDefaults instantiates a new IaasUcsdInfoRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IaasUcsdInfoRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IaasUcsdInfoRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IaasUcsdInfoRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IaasUcsdInfoRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IaasUcsdInfoRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasUcsdInfoRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasUcsdInfoRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IaasUcsdInfoRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IaasUcsdInfoRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IaasUcsdInfoRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IaasUcsdInfoRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IaasUcsdInfoRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IaasUcsdInfoRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IaasUcsdInfoRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IaasUcsdInfoRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IaasUcsdInfoRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IaasUcsdInfoRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IaasUcsdInfoRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IaasUcsdInfoRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IaasUcsdInfoRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IaasUcsdInfoRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IaasUcsdInfoRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IaasUcsdInfoRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IaasUcsdInfoRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasUcsdInfoRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasUcsdInfoRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IaasUcsdInfoRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IaasUcsdInfoRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IaasUcsdInfoRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IaasUcsdInfoRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IaasUcsdInfoRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IaasUcsdInfoRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IaasUcsdInfoRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IaasUcsdInfoRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IaasUcsdInfoRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IaasUcsdInfoRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IaasUcsdInfoRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IaasUcsdInfoRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IaasUcsdInfoRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IaasUcsdInfoRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IaasUcsdInfoRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IaasUcsdInfoRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IaasUcsdInfoRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IaasUcsdInfoRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IaasUcsdInfoRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IaasUcsdInfoRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IaasUcsdInfoRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IaasUcsdInfoRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IaasUcsdInfoRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IaasUcsdInfoRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IaasUcsdInfoRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IaasUcsdInfoRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IaasUcsdInfoRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IaasUcsdInfoRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IaasUcsdInfoRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IaasUcsdInfoRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IaasUcsdInfoRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IaasUcsdInfoRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IaasUcsdInfoRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IaasUcsdInfoRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IaasUcsdInfoRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IaasUcsdInfoRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IaasUcsdInfoRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IaasUcsdInfoRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceId

`func (o *IaasUcsdInfoRelationship) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IaasUcsdInfoRelationship) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IaasUcsdInfoRelationship) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IaasUcsdInfoRelationship) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGuid

`func (o *IaasUcsdInfoRelationship) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasUcsdInfoRelationship) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasUcsdInfoRelationship) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasUcsdInfoRelationship) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHostName

`func (o *IaasUcsdInfoRelationship) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *IaasUcsdInfoRelationship) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *IaasUcsdInfoRelationship) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *IaasUcsdInfoRelationship) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIp

`func (o *IaasUcsdInfoRelationship) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IaasUcsdInfoRelationship) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IaasUcsdInfoRelationship) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IaasUcsdInfoRelationship) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastBackup

`func (o *IaasUcsdInfoRelationship) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *IaasUcsdInfoRelationship) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *IaasUcsdInfoRelationship) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.

### HasLastBackup

`func (o *IaasUcsdInfoRelationship) HasLastBackup() bool`

HasLastBackup returns a boolean if a field has been set.

### GetNodeType

`func (o *IaasUcsdInfoRelationship) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *IaasUcsdInfoRelationship) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *IaasUcsdInfoRelationship) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *IaasUcsdInfoRelationship) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetProductName

`func (o *IaasUcsdInfoRelationship) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *IaasUcsdInfoRelationship) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *IaasUcsdInfoRelationship) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *IaasUcsdInfoRelationship) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductVendor

`func (o *IaasUcsdInfoRelationship) GetProductVendor() string`

GetProductVendor returns the ProductVendor field if non-nil, zero value otherwise.

### GetProductVendorOk

`func (o *IaasUcsdInfoRelationship) GetProductVendorOk() (*string, bool)`

GetProductVendorOk returns a tuple with the ProductVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVendor

`func (o *IaasUcsdInfoRelationship) SetProductVendor(v string)`

SetProductVendor sets ProductVendor field to given value.

### HasProductVendor

`func (o *IaasUcsdInfoRelationship) HasProductVendor() bool`

HasProductVendor returns a boolean if a field has been set.

### GetProductVersion

`func (o *IaasUcsdInfoRelationship) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *IaasUcsdInfoRelationship) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *IaasUcsdInfoRelationship) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *IaasUcsdInfoRelationship) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetStatus

`func (o *IaasUcsdInfoRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IaasUcsdInfoRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IaasUcsdInfoRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IaasUcsdInfoRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnectorPack

`func (o *IaasUcsdInfoRelationship) GetConnectorPack() []IaasConnectorPackRelationship`

GetConnectorPack returns the ConnectorPack field if non-nil, zero value otherwise.

### GetConnectorPackOk

`func (o *IaasUcsdInfoRelationship) GetConnectorPackOk() (*[]IaasConnectorPackRelationship, bool)`

GetConnectorPackOk returns a tuple with the ConnectorPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPack

`func (o *IaasUcsdInfoRelationship) SetConnectorPack(v []IaasConnectorPackRelationship)`

SetConnectorPack sets ConnectorPack field to given value.

### HasConnectorPack

`func (o *IaasUcsdInfoRelationship) HasConnectorPack() bool`

HasConnectorPack returns a boolean if a field has been set.

### SetConnectorPackNil

`func (o *IaasUcsdInfoRelationship) SetConnectorPackNil(b bool)`

 SetConnectorPackNil sets the value for ConnectorPack to be an explicit nil

### UnsetConnectorPack
`func (o *IaasUcsdInfoRelationship) UnsetConnectorPack()`

UnsetConnectorPack ensures that no value is present for ConnectorPack, not even an explicit nil
### GetDeviceStatus

`func (o *IaasUcsdInfoRelationship) GetDeviceStatus() []IaasDeviceStatusRelationship`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *IaasUcsdInfoRelationship) GetDeviceStatusOk() (*[]IaasDeviceStatusRelationship, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *IaasUcsdInfoRelationship) SetDeviceStatus(v []IaasDeviceStatusRelationship)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *IaasUcsdInfoRelationship) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### SetDeviceStatusNil

`func (o *IaasUcsdInfoRelationship) SetDeviceStatusNil(b bool)`

 SetDeviceStatusNil sets the value for DeviceStatus to be an explicit nil

### UnsetDeviceStatus
`func (o *IaasUcsdInfoRelationship) UnsetDeviceStatus()`

UnsetDeviceStatus ensures that no value is present for DeviceStatus, not even an explicit nil
### GetLicenseInfo

`func (o *IaasUcsdInfoRelationship) GetLicenseInfo() IaasLicenseInfoRelationship`

GetLicenseInfo returns the LicenseInfo field if non-nil, zero value otherwise.

### GetLicenseInfoOk

`func (o *IaasUcsdInfoRelationship) GetLicenseInfoOk() (*IaasLicenseInfoRelationship, bool)`

GetLicenseInfoOk returns a tuple with the LicenseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInfo

`func (o *IaasUcsdInfoRelationship) SetLicenseInfo(v IaasLicenseInfoRelationship)`

SetLicenseInfo sets LicenseInfo field to given value.

### HasLicenseInfo

`func (o *IaasUcsdInfoRelationship) HasLicenseInfo() bool`

HasLicenseInfo returns a boolean if a field has been set.

### GetMostRunTasks

`func (o *IaasUcsdInfoRelationship) GetMostRunTasks() []IaasMostRunTasksRelationship`

GetMostRunTasks returns the MostRunTasks field if non-nil, zero value otherwise.

### GetMostRunTasksOk

`func (o *IaasUcsdInfoRelationship) GetMostRunTasksOk() (*[]IaasMostRunTasksRelationship, bool)`

GetMostRunTasksOk returns a tuple with the MostRunTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRunTasks

`func (o *IaasUcsdInfoRelationship) SetMostRunTasks(v []IaasMostRunTasksRelationship)`

SetMostRunTasks sets MostRunTasks field to given value.

### HasMostRunTasks

`func (o *IaasUcsdInfoRelationship) HasMostRunTasks() bool`

HasMostRunTasks returns a boolean if a field has been set.

### SetMostRunTasksNil

`func (o *IaasUcsdInfoRelationship) SetMostRunTasksNil(b bool)`

 SetMostRunTasksNil sets the value for MostRunTasks to be an explicit nil

### UnsetMostRunTasks
`func (o *IaasUcsdInfoRelationship) UnsetMostRunTasks()`

UnsetMostRunTasks ensures that no value is present for MostRunTasks, not even an explicit nil
### GetRegisteredDevice

`func (o *IaasUcsdInfoRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasUcsdInfoRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasUcsdInfoRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasUcsdInfoRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetUcsdManagedInfra

`func (o *IaasUcsdInfoRelationship) GetUcsdManagedInfra() IaasUcsdManagedInfraRelationship`

GetUcsdManagedInfra returns the UcsdManagedInfra field if non-nil, zero value otherwise.

### GetUcsdManagedInfraOk

`func (o *IaasUcsdInfoRelationship) GetUcsdManagedInfraOk() (*IaasUcsdManagedInfraRelationship, bool)`

GetUcsdManagedInfraOk returns a tuple with the UcsdManagedInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsdManagedInfra

`func (o *IaasUcsdInfoRelationship) SetUcsdManagedInfra(v IaasUcsdManagedInfraRelationship)`

SetUcsdManagedInfra sets UcsdManagedInfra field to given value.

### HasUcsdManagedInfra

`func (o *IaasUcsdInfoRelationship) HasUcsdManagedInfra() bool`

HasUcsdManagedInfra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


