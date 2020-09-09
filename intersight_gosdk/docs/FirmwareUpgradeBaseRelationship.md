# FirmwareUpgradeBaseRelationship

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
**DirectDownload** | Pointer to [**FirmwareDirectDownload**](firmware.DirectDownload.md) |  | [optional] 
**FileServer** | Pointer to [**SoftwarerepositoryFileServer**](softwarerepository.FileServer.md) |  | [optional] 
**NetworkShare** | Pointer to [**FirmwareNetworkShare**](firmware.NetworkShare.md) |  | [optional] 
**SkipEstimateImpact** | Pointer to **bool** | User has the option to skip the estimate impact calculation. | [optional] 
**Status** | Pointer to **string** | Status of the upgrade operation. * &#x60;NONE&#x60; - Upgrade status is not populated. * &#x60;IN_PROGRESS&#x60; - The upgrade is in progress. * &#x60;SUCCESSFUL&#x60; - The upgrade successfully completed. * &#x60;FAILED&#x60; - The upgrade shows failed status. * &#x60;TERMINATED&#x60; - The upgrade has been terminated. | [optional] [default to "NONE"]
**UpgradeType** | Pointer to **string** | Desired upgrade mode to choose either direct download based upgrade or network share upgrade. * &#x60;direct_upgrade&#x60; - Upgrade mode is direct download. * &#x60;network_upgrade&#x60; - Upgrade mode is network upgrade. | [optional] [default to "direct_upgrade"]
**Distributable** | Pointer to [**FirmwareDistributableRelationship**](firmware.Distributable.Relationship.md) |  | [optional] 
**Release** | Pointer to [**SoftwarerepositoryReleaseRelationship**](softwarerepository.Release.Relationship.md) |  | [optional] 
**UpgradeImpact** | Pointer to [**FirmwareUpgradeImpactStatusRelationship**](firmware.UpgradeImpactStatus.Relationship.md) |  | [optional] 
**UpgradeStatus** | Pointer to [**FirmwareUpgradeStatusRelationship**](firmware.UpgradeStatus.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeBaseRelationship

`func NewFirmwareUpgradeBaseRelationship(classId string, objectType string, ) *FirmwareUpgradeBaseRelationship`

NewFirmwareUpgradeBaseRelationship instantiates a new FirmwareUpgradeBaseRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeBaseRelationshipWithDefaults

`func NewFirmwareUpgradeBaseRelationshipWithDefaults() *FirmwareUpgradeBaseRelationship`

NewFirmwareUpgradeBaseRelationshipWithDefaults instantiates a new FirmwareUpgradeBaseRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *FirmwareUpgradeBaseRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *FirmwareUpgradeBaseRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *FirmwareUpgradeBaseRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *FirmwareUpgradeBaseRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *FirmwareUpgradeBaseRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeBaseRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeBaseRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *FirmwareUpgradeBaseRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FirmwareUpgradeBaseRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FirmwareUpgradeBaseRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FirmwareUpgradeBaseRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *FirmwareUpgradeBaseRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *FirmwareUpgradeBaseRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *FirmwareUpgradeBaseRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *FirmwareUpgradeBaseRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *FirmwareUpgradeBaseRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *FirmwareUpgradeBaseRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *FirmwareUpgradeBaseRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *FirmwareUpgradeBaseRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *FirmwareUpgradeBaseRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *FirmwareUpgradeBaseRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *FirmwareUpgradeBaseRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *FirmwareUpgradeBaseRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *FirmwareUpgradeBaseRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeBaseRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeBaseRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *FirmwareUpgradeBaseRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *FirmwareUpgradeBaseRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *FirmwareUpgradeBaseRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *FirmwareUpgradeBaseRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *FirmwareUpgradeBaseRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *FirmwareUpgradeBaseRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *FirmwareUpgradeBaseRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *FirmwareUpgradeBaseRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *FirmwareUpgradeBaseRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FirmwareUpgradeBaseRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FirmwareUpgradeBaseRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FirmwareUpgradeBaseRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *FirmwareUpgradeBaseRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *FirmwareUpgradeBaseRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *FirmwareUpgradeBaseRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *FirmwareUpgradeBaseRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *FirmwareUpgradeBaseRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *FirmwareUpgradeBaseRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *FirmwareUpgradeBaseRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *FirmwareUpgradeBaseRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *FirmwareUpgradeBaseRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *FirmwareUpgradeBaseRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *FirmwareUpgradeBaseRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FirmwareUpgradeBaseRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FirmwareUpgradeBaseRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FirmwareUpgradeBaseRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *FirmwareUpgradeBaseRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *FirmwareUpgradeBaseRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *FirmwareUpgradeBaseRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *FirmwareUpgradeBaseRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *FirmwareUpgradeBaseRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *FirmwareUpgradeBaseRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *FirmwareUpgradeBaseRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *FirmwareUpgradeBaseRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *FirmwareUpgradeBaseRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *FirmwareUpgradeBaseRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *FirmwareUpgradeBaseRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *FirmwareUpgradeBaseRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDirectDownload

`func (o *FirmwareUpgradeBaseRelationship) GetDirectDownload() FirmwareDirectDownload`

GetDirectDownload returns the DirectDownload field if non-nil, zero value otherwise.

### GetDirectDownloadOk

`func (o *FirmwareUpgradeBaseRelationship) GetDirectDownloadOk() (*FirmwareDirectDownload, bool)`

GetDirectDownloadOk returns a tuple with the DirectDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDownload

`func (o *FirmwareUpgradeBaseRelationship) SetDirectDownload(v FirmwareDirectDownload)`

SetDirectDownload sets DirectDownload field to given value.

### HasDirectDownload

`func (o *FirmwareUpgradeBaseRelationship) HasDirectDownload() bool`

HasDirectDownload returns a boolean if a field has been set.

### GetFileServer

`func (o *FirmwareUpgradeBaseRelationship) GetFileServer() SoftwarerepositoryFileServer`

GetFileServer returns the FileServer field if non-nil, zero value otherwise.

### GetFileServerOk

`func (o *FirmwareUpgradeBaseRelationship) GetFileServerOk() (*SoftwarerepositoryFileServer, bool)`

GetFileServerOk returns a tuple with the FileServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServer

`func (o *FirmwareUpgradeBaseRelationship) SetFileServer(v SoftwarerepositoryFileServer)`

SetFileServer sets FileServer field to given value.

### HasFileServer

`func (o *FirmwareUpgradeBaseRelationship) HasFileServer() bool`

HasFileServer returns a boolean if a field has been set.

### GetNetworkShare

`func (o *FirmwareUpgradeBaseRelationship) GetNetworkShare() FirmwareNetworkShare`

GetNetworkShare returns the NetworkShare field if non-nil, zero value otherwise.

### GetNetworkShareOk

`func (o *FirmwareUpgradeBaseRelationship) GetNetworkShareOk() (*FirmwareNetworkShare, bool)`

GetNetworkShareOk returns a tuple with the NetworkShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkShare

`func (o *FirmwareUpgradeBaseRelationship) SetNetworkShare(v FirmwareNetworkShare)`

SetNetworkShare sets NetworkShare field to given value.

### HasNetworkShare

`func (o *FirmwareUpgradeBaseRelationship) HasNetworkShare() bool`

HasNetworkShare returns a boolean if a field has been set.

### GetSkipEstimateImpact

`func (o *FirmwareUpgradeBaseRelationship) GetSkipEstimateImpact() bool`

GetSkipEstimateImpact returns the SkipEstimateImpact field if non-nil, zero value otherwise.

### GetSkipEstimateImpactOk

`func (o *FirmwareUpgradeBaseRelationship) GetSkipEstimateImpactOk() (*bool, bool)`

GetSkipEstimateImpactOk returns a tuple with the SkipEstimateImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEstimateImpact

`func (o *FirmwareUpgradeBaseRelationship) SetSkipEstimateImpact(v bool)`

SetSkipEstimateImpact sets SkipEstimateImpact field to given value.

### HasSkipEstimateImpact

`func (o *FirmwareUpgradeBaseRelationship) HasSkipEstimateImpact() bool`

HasSkipEstimateImpact returns a boolean if a field has been set.

### GetStatus

`func (o *FirmwareUpgradeBaseRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirmwareUpgradeBaseRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirmwareUpgradeBaseRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirmwareUpgradeBaseRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradeType

`func (o *FirmwareUpgradeBaseRelationship) GetUpgradeType() string`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *FirmwareUpgradeBaseRelationship) GetUpgradeTypeOk() (*string, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *FirmwareUpgradeBaseRelationship) SetUpgradeType(v string)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *FirmwareUpgradeBaseRelationship) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetDistributable

`func (o *FirmwareUpgradeBaseRelationship) GetDistributable() FirmwareDistributableRelationship`

GetDistributable returns the Distributable field if non-nil, zero value otherwise.

### GetDistributableOk

`func (o *FirmwareUpgradeBaseRelationship) GetDistributableOk() (*FirmwareDistributableRelationship, bool)`

GetDistributableOk returns a tuple with the Distributable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributable

`func (o *FirmwareUpgradeBaseRelationship) SetDistributable(v FirmwareDistributableRelationship)`

SetDistributable sets Distributable field to given value.

### HasDistributable

`func (o *FirmwareUpgradeBaseRelationship) HasDistributable() bool`

HasDistributable returns a boolean if a field has been set.

### GetRelease

`func (o *FirmwareUpgradeBaseRelationship) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *FirmwareUpgradeBaseRelationship) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *FirmwareUpgradeBaseRelationship) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *FirmwareUpgradeBaseRelationship) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetUpgradeImpact

`func (o *FirmwareUpgradeBaseRelationship) GetUpgradeImpact() FirmwareUpgradeImpactStatusRelationship`

GetUpgradeImpact returns the UpgradeImpact field if non-nil, zero value otherwise.

### GetUpgradeImpactOk

`func (o *FirmwareUpgradeBaseRelationship) GetUpgradeImpactOk() (*FirmwareUpgradeImpactStatusRelationship, bool)`

GetUpgradeImpactOk returns a tuple with the UpgradeImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeImpact

`func (o *FirmwareUpgradeBaseRelationship) SetUpgradeImpact(v FirmwareUpgradeImpactStatusRelationship)`

SetUpgradeImpact sets UpgradeImpact field to given value.

### HasUpgradeImpact

`func (o *FirmwareUpgradeBaseRelationship) HasUpgradeImpact() bool`

HasUpgradeImpact returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *FirmwareUpgradeBaseRelationship) GetUpgradeStatus() FirmwareUpgradeStatusRelationship`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *FirmwareUpgradeBaseRelationship) GetUpgradeStatusOk() (*FirmwareUpgradeStatusRelationship, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *FirmwareUpgradeBaseRelationship) SetUpgradeStatus(v FirmwareUpgradeStatusRelationship)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *FirmwareUpgradeBaseRelationship) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


