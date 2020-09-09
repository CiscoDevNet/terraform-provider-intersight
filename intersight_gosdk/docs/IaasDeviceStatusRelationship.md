# IaasDeviceStatusRelationship

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
**AccountName** | Pointer to **string** | The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD. | [optional] [readonly] 
**AccountType** | Pointer to **string** | The UCSD Infra Account type. | [optional] [readonly] 
**Category** | Pointer to **string** | The UCSD Infra Account category. | [optional] [readonly] 
**ClaimStatus** | Pointer to **string** | Describes if the device is claimed in Intersight or not. * &#x60;Unknown&#x60; - If UCS Director managed account claim status information is unknown. * &#x60;Yes&#x60; - If UCS Director managed account is claimed in Intersight. * &#x60;No&#x60; - If UCS Director managed account is not claimed in Intersight. * &#x60;Not Applicable&#x60; - If UCS Director managed account is not capable of providing claim status information. | [optional] [readonly] [default to "Unknown"]
**ConnectionStatus** | Pointer to **string** | Describes about the connection status between the UCSD and the actual end device. | [optional] [readonly] 
**DeviceModel** | Pointer to **string** | Describes about the device model. | [optional] [readonly] 
**DeviceVendor** | Pointer to **string** | Describes about the device vendor/manufacturer of the device. | [optional] [readonly] 
**DeviceVersion** | Pointer to **string** | Describes about the current firmware version running on the device. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | The IPAddress of the device. | [optional] [readonly] 
**Pod** | Pointer to **string** | Describes about the pod to which this device belongs to in UCSD. | [optional] [readonly] 
**PodType** | Pointer to **string** | Describes about the podType of Pod to which this device belongs to in UCSD. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 

## Methods

### NewIaasDeviceStatusRelationship

`func NewIaasDeviceStatusRelationship(classId string, objectType string, ) *IaasDeviceStatusRelationship`

NewIaasDeviceStatusRelationship instantiates a new IaasDeviceStatusRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasDeviceStatusRelationshipWithDefaults

`func NewIaasDeviceStatusRelationshipWithDefaults() *IaasDeviceStatusRelationship`

NewIaasDeviceStatusRelationshipWithDefaults instantiates a new IaasDeviceStatusRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IaasDeviceStatusRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IaasDeviceStatusRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IaasDeviceStatusRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IaasDeviceStatusRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IaasDeviceStatusRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasDeviceStatusRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasDeviceStatusRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IaasDeviceStatusRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IaasDeviceStatusRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IaasDeviceStatusRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IaasDeviceStatusRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IaasDeviceStatusRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IaasDeviceStatusRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IaasDeviceStatusRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IaasDeviceStatusRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IaasDeviceStatusRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IaasDeviceStatusRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IaasDeviceStatusRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IaasDeviceStatusRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IaasDeviceStatusRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IaasDeviceStatusRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IaasDeviceStatusRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IaasDeviceStatusRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IaasDeviceStatusRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasDeviceStatusRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasDeviceStatusRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IaasDeviceStatusRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IaasDeviceStatusRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IaasDeviceStatusRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IaasDeviceStatusRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IaasDeviceStatusRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IaasDeviceStatusRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IaasDeviceStatusRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IaasDeviceStatusRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IaasDeviceStatusRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IaasDeviceStatusRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IaasDeviceStatusRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IaasDeviceStatusRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IaasDeviceStatusRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IaasDeviceStatusRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IaasDeviceStatusRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IaasDeviceStatusRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IaasDeviceStatusRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IaasDeviceStatusRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IaasDeviceStatusRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IaasDeviceStatusRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IaasDeviceStatusRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IaasDeviceStatusRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IaasDeviceStatusRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IaasDeviceStatusRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IaasDeviceStatusRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IaasDeviceStatusRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IaasDeviceStatusRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IaasDeviceStatusRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IaasDeviceStatusRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IaasDeviceStatusRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IaasDeviceStatusRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IaasDeviceStatusRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IaasDeviceStatusRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IaasDeviceStatusRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IaasDeviceStatusRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IaasDeviceStatusRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IaasDeviceStatusRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IaasDeviceStatusRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAccountName

`func (o *IaasDeviceStatusRelationship) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *IaasDeviceStatusRelationship) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *IaasDeviceStatusRelationship) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *IaasDeviceStatusRelationship) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountType

`func (o *IaasDeviceStatusRelationship) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *IaasDeviceStatusRelationship) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *IaasDeviceStatusRelationship) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *IaasDeviceStatusRelationship) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCategory

`func (o *IaasDeviceStatusRelationship) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IaasDeviceStatusRelationship) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IaasDeviceStatusRelationship) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IaasDeviceStatusRelationship) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClaimStatus

`func (o *IaasDeviceStatusRelationship) GetClaimStatus() string`

GetClaimStatus returns the ClaimStatus field if non-nil, zero value otherwise.

### GetClaimStatusOk

`func (o *IaasDeviceStatusRelationship) GetClaimStatusOk() (*string, bool)`

GetClaimStatusOk returns a tuple with the ClaimStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimStatus

`func (o *IaasDeviceStatusRelationship) SetClaimStatus(v string)`

SetClaimStatus sets ClaimStatus field to given value.

### HasClaimStatus

`func (o *IaasDeviceStatusRelationship) HasClaimStatus() bool`

HasClaimStatus returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *IaasDeviceStatusRelationship) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *IaasDeviceStatusRelationship) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *IaasDeviceStatusRelationship) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *IaasDeviceStatusRelationship) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDeviceModel

`func (o *IaasDeviceStatusRelationship) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *IaasDeviceStatusRelationship) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *IaasDeviceStatusRelationship) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *IaasDeviceStatusRelationship) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *IaasDeviceStatusRelationship) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *IaasDeviceStatusRelationship) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *IaasDeviceStatusRelationship) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *IaasDeviceStatusRelationship) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDeviceVersion

`func (o *IaasDeviceStatusRelationship) GetDeviceVersion() string`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *IaasDeviceStatusRelationship) GetDeviceVersionOk() (*string, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVersion

`func (o *IaasDeviceStatusRelationship) SetDeviceVersion(v string)`

SetDeviceVersion sets DeviceVersion field to given value.

### HasDeviceVersion

`func (o *IaasDeviceStatusRelationship) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### GetIpAddress

`func (o *IaasDeviceStatusRelationship) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IaasDeviceStatusRelationship) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IaasDeviceStatusRelationship) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *IaasDeviceStatusRelationship) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPod

`func (o *IaasDeviceStatusRelationship) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *IaasDeviceStatusRelationship) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *IaasDeviceStatusRelationship) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *IaasDeviceStatusRelationship) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPodType

`func (o *IaasDeviceStatusRelationship) GetPodType() string`

GetPodType returns the PodType field if non-nil, zero value otherwise.

### GetPodTypeOk

`func (o *IaasDeviceStatusRelationship) GetPodTypeOk() (*string, bool)`

GetPodTypeOk returns a tuple with the PodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodType

`func (o *IaasDeviceStatusRelationship) SetPodType(v string)`

SetPodType sets PodType field to given value.

### HasPodType

`func (o *IaasDeviceStatusRelationship) HasPodType() bool`

HasPodType returns a boolean if a field has been set.

### GetGuid

`func (o *IaasDeviceStatusRelationship) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasDeviceStatusRelationship) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasDeviceStatusRelationship) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasDeviceStatusRelationship) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


