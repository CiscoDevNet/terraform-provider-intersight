# FirmwareUpgradeStatusRelationship

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
**DownloadError** | Pointer to **string** | The error message from the endpoint during the download. | [optional] 
**DownloadMessage** | Pointer to **string** | The message from the endpoint during the download. | [optional] 
**DownloadPercentage** | Pointer to **int64** | The percentage of the image downloaded in the endpoint. | [optional] 
**DownloadProgress** | Pointer to **int64** | The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. | [optional] 
**DownloadRetries** | Pointer to **int64** | The number of retries the plugin attempted before succeeding or failing the download. | [optional] 
**DownloadStage** | Pointer to **string** | The image download stages. Example:downloading, flashing. | [optional] 
**EpPowerStatus** | Pointer to **string** | The server power status after the upgrade request is submitted in the endpoint. * &#x60;none&#x60; - Server power status is none. * &#x60;powered on&#x60; - Server power status is powered on. * &#x60;powered off&#x60; - Server power status is powered off. | [optional] [default to "none"]
**OverallError** | Pointer to **string** | The reason for the operation failure. | [optional] 
**OverallPercentage** | Pointer to **int64** | The overall percentage of the operation. | [optional] 
**Overallstatus** | Pointer to **string** | The overall status of the operation. * &#x60;none&#x60; - Upgrade stage is no upgrade stage. * &#x60;started&#x60; - Upgrade stage is started. * &#x60;prepare initiating&#x60; - Upgrade configuration is being prepared. * &#x60;prepare initiated&#x60; - Upgrade configuration is initiated. * &#x60;prepared&#x60; - Upgrade configuration is prepared. * &#x60;download initiating&#x60; - Upgrade stage is download initiating. * &#x60;download initiated&#x60; - Upgrade stage is download initiated. * &#x60;downloading&#x60; - Upgrade stage is downloading. * &#x60;downloaded&#x60; - Upgrade stage is downloaded. * &#x60;upgrade initiating on fabric A&#x60; - Upgrade stage is in upgrade initiating when upgrade is being started in endopint. * &#x60;upgrade initiated on fabric A&#x60; - Upgrade stage is in upgrade initiated when the upgrade has started in endpoint. * &#x60;upgrading fabric A&#x60; - Upgrade stage is in upgrading when the upgrade requires reboot to complete. * &#x60;rebooting fabric A&#x60; - Upgrade is in rebooting when the endpoint is being rebooted. * &#x60;upgraded fabric A&#x60; - Upgrade stage is in upgraded when the corresponding endpoint has completed. * &#x60;upgrade initiating on fabric B&#x60; - Upgrade stage is in upgrade initiating when upgrade is being started in endopint. * &#x60;upgrade initiated on fabric B&#x60; - Upgrade stage is in upgrade initiated when upgrade has started in endpoint. * &#x60;upgrading fabric B&#x60; - Upgrade stage is in upgrading when the upgrade requires reboot to complete. * &#x60;rebooting fabric B&#x60; - Upgrade is in rebooting when the endpoint is being rebooted. * &#x60;upgraded fabric B&#x60; - Upgrade stage is in upgraded when the corresponding endpoint has completed. * &#x60;upgrade initiating&#x60; - Upgrade stage is upgrade initiating. * &#x60;upgrade initiated&#x60; - Upgrade stage is upgrade initiated. * &#x60;upgrading&#x60; - Upgrade stage is upgrading. * &#x60;oob images staging&#x60; - Out-of-band component images staging. * &#x60;oob images staged&#x60; - Out-of-band component images staged. * &#x60;rebooting&#x60; - Upgrade is rebooting the endpoint. * &#x60;upgraded&#x60; - Upgrade stage is upgraded. * &#x60;success&#x60; - Upgrade stage is success. * &#x60;failed&#x60; - Upgrade stage is upgrade failed. * &#x60;terminated&#x60; - Upgrade stage is terminated. * &#x60;pending&#x60; - Upgrade stage is pending. * &#x60;ReadyForCache&#x60; - The image is ready to be cached into the Intersight Appliance. * &#x60;Caching&#x60; - The image will be cached into Intersight Appliance or an endpoint cache. * &#x60;Cached&#x60; - The image has been cached into the Intersight Appliance or endpoint cache. * &#x60;CachingFailed&#x60; - The image caching into the Intersight Appliance failed or endpoint cache. | [optional] [default to "none"]
**PendingType** | Pointer to **string** | Pending reason for the upgrade waiting. * &#x60;none&#x60; - Upgrade pending reason is none. * &#x60;pending for next reboot&#x60; - Upgrade pending reason is pending for next reboot. | [optional] [default to "none"]
**Sha256checksum** | Pointer to **string** | The sha256checksum of the downloaded file as calculated by the download plugin after successfully downloading a file. | [optional] 
**Upgrade** | Pointer to [**FirmwareUpgradeBaseRelationship**](firmware.UpgradeBase.Relationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeStatusRelationship

`func NewFirmwareUpgradeStatusRelationship(classId string, objectType string, ) *FirmwareUpgradeStatusRelationship`

NewFirmwareUpgradeStatusRelationship instantiates a new FirmwareUpgradeStatusRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeStatusRelationshipWithDefaults

`func NewFirmwareUpgradeStatusRelationshipWithDefaults() *FirmwareUpgradeStatusRelationship`

NewFirmwareUpgradeStatusRelationshipWithDefaults instantiates a new FirmwareUpgradeStatusRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *FirmwareUpgradeStatusRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *FirmwareUpgradeStatusRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *FirmwareUpgradeStatusRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *FirmwareUpgradeStatusRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *FirmwareUpgradeStatusRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeStatusRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeStatusRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *FirmwareUpgradeStatusRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FirmwareUpgradeStatusRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FirmwareUpgradeStatusRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FirmwareUpgradeStatusRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *FirmwareUpgradeStatusRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *FirmwareUpgradeStatusRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *FirmwareUpgradeStatusRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *FirmwareUpgradeStatusRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *FirmwareUpgradeStatusRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *FirmwareUpgradeStatusRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *FirmwareUpgradeStatusRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *FirmwareUpgradeStatusRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *FirmwareUpgradeStatusRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *FirmwareUpgradeStatusRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *FirmwareUpgradeStatusRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *FirmwareUpgradeStatusRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *FirmwareUpgradeStatusRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeStatusRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeStatusRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *FirmwareUpgradeStatusRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *FirmwareUpgradeStatusRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *FirmwareUpgradeStatusRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *FirmwareUpgradeStatusRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *FirmwareUpgradeStatusRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *FirmwareUpgradeStatusRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *FirmwareUpgradeStatusRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *FirmwareUpgradeStatusRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *FirmwareUpgradeStatusRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FirmwareUpgradeStatusRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FirmwareUpgradeStatusRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FirmwareUpgradeStatusRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *FirmwareUpgradeStatusRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *FirmwareUpgradeStatusRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *FirmwareUpgradeStatusRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *FirmwareUpgradeStatusRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *FirmwareUpgradeStatusRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *FirmwareUpgradeStatusRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *FirmwareUpgradeStatusRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *FirmwareUpgradeStatusRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *FirmwareUpgradeStatusRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *FirmwareUpgradeStatusRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *FirmwareUpgradeStatusRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FirmwareUpgradeStatusRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FirmwareUpgradeStatusRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FirmwareUpgradeStatusRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *FirmwareUpgradeStatusRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *FirmwareUpgradeStatusRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *FirmwareUpgradeStatusRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *FirmwareUpgradeStatusRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *FirmwareUpgradeStatusRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *FirmwareUpgradeStatusRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *FirmwareUpgradeStatusRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *FirmwareUpgradeStatusRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *FirmwareUpgradeStatusRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *FirmwareUpgradeStatusRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *FirmwareUpgradeStatusRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *FirmwareUpgradeStatusRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDownloadError

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadError() string`

GetDownloadError returns the DownloadError field if non-nil, zero value otherwise.

### GetDownloadErrorOk

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadErrorOk() (*string, bool)`

GetDownloadErrorOk returns a tuple with the DownloadError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadError

`func (o *FirmwareUpgradeStatusRelationship) SetDownloadError(v string)`

SetDownloadError sets DownloadError field to given value.

### HasDownloadError

`func (o *FirmwareUpgradeStatusRelationship) HasDownloadError() bool`

HasDownloadError returns a boolean if a field has been set.

### GetDownloadMessage

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadMessage() string`

GetDownloadMessage returns the DownloadMessage field if non-nil, zero value otherwise.

### GetDownloadMessageOk

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadMessageOk() (*string, bool)`

GetDownloadMessageOk returns a tuple with the DownloadMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMessage

`func (o *FirmwareUpgradeStatusRelationship) SetDownloadMessage(v string)`

SetDownloadMessage sets DownloadMessage field to given value.

### HasDownloadMessage

`func (o *FirmwareUpgradeStatusRelationship) HasDownloadMessage() bool`

HasDownloadMessage returns a boolean if a field has been set.

### GetDownloadPercentage

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadPercentage() int64`

GetDownloadPercentage returns the DownloadPercentage field if non-nil, zero value otherwise.

### GetDownloadPercentageOk

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadPercentageOk() (*int64, bool)`

GetDownloadPercentageOk returns a tuple with the DownloadPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentage

`func (o *FirmwareUpgradeStatusRelationship) SetDownloadPercentage(v int64)`

SetDownloadPercentage sets DownloadPercentage field to given value.

### HasDownloadPercentage

`func (o *FirmwareUpgradeStatusRelationship) HasDownloadPercentage() bool`

HasDownloadPercentage returns a boolean if a field has been set.

### GetDownloadProgress

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadProgress() int64`

GetDownloadProgress returns the DownloadProgress field if non-nil, zero value otherwise.

### GetDownloadProgressOk

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadProgressOk() (*int64, bool)`

GetDownloadProgressOk returns a tuple with the DownloadProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadProgress

`func (o *FirmwareUpgradeStatusRelationship) SetDownloadProgress(v int64)`

SetDownloadProgress sets DownloadProgress field to given value.

### HasDownloadProgress

`func (o *FirmwareUpgradeStatusRelationship) HasDownloadProgress() bool`

HasDownloadProgress returns a boolean if a field has been set.

### GetDownloadRetries

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadRetries() int64`

GetDownloadRetries returns the DownloadRetries field if non-nil, zero value otherwise.

### GetDownloadRetriesOk

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadRetriesOk() (*int64, bool)`

GetDownloadRetriesOk returns a tuple with the DownloadRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadRetries

`func (o *FirmwareUpgradeStatusRelationship) SetDownloadRetries(v int64)`

SetDownloadRetries sets DownloadRetries field to given value.

### HasDownloadRetries

`func (o *FirmwareUpgradeStatusRelationship) HasDownloadRetries() bool`

HasDownloadRetries returns a boolean if a field has been set.

### GetDownloadStage

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadStage() string`

GetDownloadStage returns the DownloadStage field if non-nil, zero value otherwise.

### GetDownloadStageOk

`func (o *FirmwareUpgradeStatusRelationship) GetDownloadStageOk() (*string, bool)`

GetDownloadStageOk returns a tuple with the DownloadStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStage

`func (o *FirmwareUpgradeStatusRelationship) SetDownloadStage(v string)`

SetDownloadStage sets DownloadStage field to given value.

### HasDownloadStage

`func (o *FirmwareUpgradeStatusRelationship) HasDownloadStage() bool`

HasDownloadStage returns a boolean if a field has been set.

### GetEpPowerStatus

`func (o *FirmwareUpgradeStatusRelationship) GetEpPowerStatus() string`

GetEpPowerStatus returns the EpPowerStatus field if non-nil, zero value otherwise.

### GetEpPowerStatusOk

`func (o *FirmwareUpgradeStatusRelationship) GetEpPowerStatusOk() (*string, bool)`

GetEpPowerStatusOk returns a tuple with the EpPowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpPowerStatus

`func (o *FirmwareUpgradeStatusRelationship) SetEpPowerStatus(v string)`

SetEpPowerStatus sets EpPowerStatus field to given value.

### HasEpPowerStatus

`func (o *FirmwareUpgradeStatusRelationship) HasEpPowerStatus() bool`

HasEpPowerStatus returns a boolean if a field has been set.

### GetOverallError

`func (o *FirmwareUpgradeStatusRelationship) GetOverallError() string`

GetOverallError returns the OverallError field if non-nil, zero value otherwise.

### GetOverallErrorOk

`func (o *FirmwareUpgradeStatusRelationship) GetOverallErrorOk() (*string, bool)`

GetOverallErrorOk returns a tuple with the OverallError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallError

`func (o *FirmwareUpgradeStatusRelationship) SetOverallError(v string)`

SetOverallError sets OverallError field to given value.

### HasOverallError

`func (o *FirmwareUpgradeStatusRelationship) HasOverallError() bool`

HasOverallError returns a boolean if a field has been set.

### GetOverallPercentage

`func (o *FirmwareUpgradeStatusRelationship) GetOverallPercentage() int64`

GetOverallPercentage returns the OverallPercentage field if non-nil, zero value otherwise.

### GetOverallPercentageOk

`func (o *FirmwareUpgradeStatusRelationship) GetOverallPercentageOk() (*int64, bool)`

GetOverallPercentageOk returns a tuple with the OverallPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallPercentage

`func (o *FirmwareUpgradeStatusRelationship) SetOverallPercentage(v int64)`

SetOverallPercentage sets OverallPercentage field to given value.

### HasOverallPercentage

`func (o *FirmwareUpgradeStatusRelationship) HasOverallPercentage() bool`

HasOverallPercentage returns a boolean if a field has been set.

### GetOverallstatus

`func (o *FirmwareUpgradeStatusRelationship) GetOverallstatus() string`

GetOverallstatus returns the Overallstatus field if non-nil, zero value otherwise.

### GetOverallstatusOk

`func (o *FirmwareUpgradeStatusRelationship) GetOverallstatusOk() (*string, bool)`

GetOverallstatusOk returns a tuple with the Overallstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallstatus

`func (o *FirmwareUpgradeStatusRelationship) SetOverallstatus(v string)`

SetOverallstatus sets Overallstatus field to given value.

### HasOverallstatus

`func (o *FirmwareUpgradeStatusRelationship) HasOverallstatus() bool`

HasOverallstatus returns a boolean if a field has been set.

### GetPendingType

`func (o *FirmwareUpgradeStatusRelationship) GetPendingType() string`

GetPendingType returns the PendingType field if non-nil, zero value otherwise.

### GetPendingTypeOk

`func (o *FirmwareUpgradeStatusRelationship) GetPendingTypeOk() (*string, bool)`

GetPendingTypeOk returns a tuple with the PendingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingType

`func (o *FirmwareUpgradeStatusRelationship) SetPendingType(v string)`

SetPendingType sets PendingType field to given value.

### HasPendingType

`func (o *FirmwareUpgradeStatusRelationship) HasPendingType() bool`

HasPendingType returns a boolean if a field has been set.

### GetSha256checksum

`func (o *FirmwareUpgradeStatusRelationship) GetSha256checksum() string`

GetSha256checksum returns the Sha256checksum field if non-nil, zero value otherwise.

### GetSha256checksumOk

`func (o *FirmwareUpgradeStatusRelationship) GetSha256checksumOk() (*string, bool)`

GetSha256checksumOk returns a tuple with the Sha256checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256checksum

`func (o *FirmwareUpgradeStatusRelationship) SetSha256checksum(v string)`

SetSha256checksum sets Sha256checksum field to given value.

### HasSha256checksum

`func (o *FirmwareUpgradeStatusRelationship) HasSha256checksum() bool`

HasSha256checksum returns a boolean if a field has been set.

### GetUpgrade

`func (o *FirmwareUpgradeStatusRelationship) GetUpgrade() FirmwareUpgradeBaseRelationship`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *FirmwareUpgradeStatusRelationship) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *FirmwareUpgradeStatusRelationship) SetUpgrade(v FirmwareUpgradeBaseRelationship)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *FirmwareUpgradeStatusRelationship) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetWorkflow

`func (o *FirmwareUpgradeStatusRelationship) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *FirmwareUpgradeStatusRelationship) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *FirmwareUpgradeStatusRelationship) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *FirmwareUpgradeStatusRelationship) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


