# TechsupportmanagementTechSupportStatusRelationship

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
**FileName** | Pointer to **string** | The name of the Techsupport bundle file. | [optional] 
**Reason** | Pointer to **string** | Reason for techsupport failure, if any. | [optional] 
**RelayReason** | Pointer to **string** | Reason for status relay failure, if any. | [optional] [readonly] 
**RelayStatus** | Pointer to **string** | Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed. | [optional] [readonly] 
**RequestTs** | Pointer to [**time.Time**](time.Time.md) | The time at which the techsupport request was initiated. | [optional] 
**Status** | Pointer to **string** | Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection. | [optional] 
**TechsupportDownloadUrl** | Pointer to **string** | The Url to download the techsupport file. | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](asset.ClusterMember.Relationship.md) |  | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**OriginResource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**TechSupportRequest** | Pointer to [**TechsupportmanagementTechSupportBundleRelationship**](techsupportmanagement.TechSupportBundle.Relationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementTechSupportStatusRelationship

`func NewTechsupportmanagementTechSupportStatusRelationship(classId string, objectType string, ) *TechsupportmanagementTechSupportStatusRelationship`

NewTechsupportmanagementTechSupportStatusRelationship instantiates a new TechsupportmanagementTechSupportStatusRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportStatusRelationshipWithDefaults

`func NewTechsupportmanagementTechSupportStatusRelationshipWithDefaults() *TechsupportmanagementTechSupportStatusRelationship`

NewTechsupportmanagementTechSupportStatusRelationshipWithDefaults instantiates a new TechsupportmanagementTechSupportStatusRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *TechsupportmanagementTechSupportStatusRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *TechsupportmanagementTechSupportStatusRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *TechsupportmanagementTechSupportStatusRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetFileName

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetReason

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelayReason

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetRelayReason() string`

GetRelayReason returns the RelayReason field if non-nil, zero value otherwise.

### GetRelayReasonOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetRelayReasonOk() (*string, bool)`

GetRelayReasonOk returns a tuple with the RelayReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayReason

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetRelayReason(v string)`

SetRelayReason sets RelayReason field to given value.

### HasRelayReason

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasRelayReason() bool`

HasRelayReason returns a boolean if a field has been set.

### GetRelayStatus

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetRelayStatus() string`

GetRelayStatus returns the RelayStatus field if non-nil, zero value otherwise.

### GetRelayStatusOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetRelayStatusOk() (*string, bool)`

GetRelayStatusOk returns a tuple with the RelayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayStatus

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetRelayStatus(v string)`

SetRelayStatus sets RelayStatus field to given value.

### HasRelayStatus

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasRelayStatus() bool`

HasRelayStatus returns a boolean if a field has been set.

### GetRequestTs

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetRequestTs() time.Time`

GetRequestTs returns the RequestTs field if non-nil, zero value otherwise.

### GetRequestTsOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetRequestTsOk() (*time.Time, bool)`

GetRequestTsOk returns a tuple with the RequestTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTs

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetRequestTs(v time.Time)`

SetRequestTs sets RequestTs field to given value.

### HasRequestTs

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasRequestTs() bool`

HasRequestTs returns a boolean if a field has been set.

### GetStatus

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetTechsupportDownloadUrl() string`

GetTechsupportDownloadUrl returns the TechsupportDownloadUrl field if non-nil, zero value otherwise.

### GetTechsupportDownloadUrlOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetTechsupportDownloadUrlOk() (*string, bool)`

GetTechsupportDownloadUrlOk returns a tuple with the TechsupportDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetTechsupportDownloadUrl(v string)`

SetTechsupportDownloadUrl sets TechsupportDownloadUrl field to given value.

### HasTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasTechsupportDownloadUrl() bool`

HasTechsupportDownloadUrl returns a boolean if a field has been set.

### GetClusterMember

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetOriginResource

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetOriginResource() MoBaseMoRelationship`

GetOriginResource returns the OriginResource field if non-nil, zero value otherwise.

### GetOriginResourceOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetOriginResourceOk() (*MoBaseMoRelationship, bool)`

GetOriginResourceOk returns a tuple with the OriginResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginResource

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetOriginResource(v MoBaseMoRelationship)`

SetOriginResource sets OriginResource field to given value.

### HasOriginResource

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasOriginResource() bool`

HasOriginResource returns a boolean if a field has been set.

### GetTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetTechSupportRequest() TechsupportmanagementTechSupportBundleRelationship`

GetTechSupportRequest returns the TechSupportRequest field if non-nil, zero value otherwise.

### GetTechSupportRequestOk

`func (o *TechsupportmanagementTechSupportStatusRelationship) GetTechSupportRequestOk() (*TechsupportmanagementTechSupportBundleRelationship, bool)`

GetTechSupportRequestOk returns a tuple with the TechSupportRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatusRelationship) SetTechSupportRequest(v TechsupportmanagementTechSupportBundleRelationship)`

SetTechSupportRequest sets TechSupportRequest field to given value.

### HasTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatusRelationship) HasTechSupportRequest() bool`

HasTechSupportRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


