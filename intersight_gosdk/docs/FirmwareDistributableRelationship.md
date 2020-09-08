# FirmwareDistributableRelationship

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
**Description** | Pointer to **string** | User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. | [optional] 
**DownloadCount** | Pointer to **int64** | The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**ImportAction** | Pointer to **string** | The action to be performed on the imported file. If &#39;PreCache&#39; is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If &#39;Evict&#39; is set, the cached file will be removed. Applicable in Intersight appliance deployment. If &#39;GeneratePreSignedUploadUrl&#39; is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If &#39;CompleteImportProcess&#39; is set, the ImportState is marked as &#39;Imported&#39;. Applicable for local machine source. If &#39;Cancel&#39; is set, the ImportState is marked as &#39;Failed&#39;. Applicable for local machine source. * &#x60;None&#x60; - No action should be taken on the imported file. * &#x60;GeneratePreSignedUploadUrl&#x60; - Generate pre signed URL of file for importing into the repository. * &#x60;GeneratePreSignedDownloadUrl&#x60; - Generate pre signed URL of file in the repository to download. * &#x60;CompleteImportProcess&#x60; - Mark that the import process of the file into the repository is complete. * &#x60;MarkImportFailed&#x60; - Mark to indicate that the import process of the file into the repository failed. * &#x60;PreCache&#x60; - Cache the file into the Intersight Appliance. * &#x60;Cancel&#x60; - The cancel import process for the file into the repository. * &#x60;Extract&#x60; - The action to extract the file in the external repository. * &#x60;Evict&#x60; - Evict the cached file from the Intersight Appliance. | [optional] [default to "None"]
**ImportState** | Pointer to **string** | The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process. * &#x60;ReadyForImport&#x60; - The image is ready to be imported into the repository. * &#x60;Importing&#x60; - The image is being imported into the repository. * &#x60;Imported&#x60; - The image has been extracted and imported into the repository. * &#x60;PendingExtraction&#x60; - Indicates that the image has been imported but not extracted in the repository. * &#x60;Extracting&#x60; - Indicates that the image is being extracted into the repository. * &#x60;Extracted&#x60; - Indicates that the image has been extracted into the repository. * &#x60;Failed&#x60; - The image import from an external source to the repository has failed. * &#x60;MetaOnly&#x60; - The image is present in an external repository. * &#x60;ReadyForCache&#x60; - The image is ready to be cached into the Intersight Appliance. * &#x60;Caching&#x60; - Indicates that the image is being cached into the Intersight Appliance or endpoint cache. * &#x60;Cached&#x60; - Indicates that the image has been cached into the Intersight Appliance or endpoint cache. * &#x60;CachingFailed&#x60; - Indicates that the image caching into the Intersight Appliance failed or endpoint cache. * &#x60;Corrupted&#x60; - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached. * &#x60;Evicted&#x60; - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. | [optional] [readonly] [default to "ReadyForImport"]
**ImportedTime** | Pointer to [**time.Time**](time.Time.md) | The time at which this image or file was imported/cached into the repositry. if the &#39;ImportState&#39; is &#39;Imported&#39;, the time at which this image or file was imported. if the &#39;ImportState&#39; is &#39;Cached&#39;, the time at which this image or file was cached. | [optional] [readonly] 
**LastAccessTime** | Pointer to [**time.Time**](time.Time.md) | The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**Md5sum** | Pointer to **string** | The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**Name** | Pointer to **string** | The name of the file. It is populated as part of the image import operation. | [optional] 
**ReleaseDate** | Pointer to [**time.Time**](time.Time.md) | The date on which the file was released or distributed by its vendor. | [optional] 
**Sha512sum** | Pointer to **string** | The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**Size** | Pointer to **int64** | The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**SoftwareAdvisoryUrl** | Pointer to **string** | The software advisory, if any, provided by the vendor for this file. | [optional] 
**Source** | Pointer to [**SoftwarerepositoryFileServer**](softwarerepository.FileServer.md) |  | [optional] 
**Version** | Pointer to **string** | Vendor provided version for the file. | [optional] 
**BundleType** | Pointer to **string** | The bundle type of the image, as published on cisco.com. | [optional] [readonly] 
**ComponentMeta** | Pointer to [**[]FirmwareComponentMeta**](firmware.ComponentMeta.md) |  | [optional] 
**Guid** | Pointer to **string** | The unique identifier for an image in a Cisco repository. | [optional] [readonly] 
**Mdfid** | Pointer to **string** | The mdfid of the image provided by cisco.com. | [optional] 
**Model** | Pointer to **string** | The endpoint model for which this firmware image is applicable. | [optional] 
**PlatformType** | Pointer to **string** | The platform type of the image. | [optional] [readonly] 
**RecommendedBuild** | Pointer to **string** | The build which is recommended by Cisco. | [optional] 
**ReleaseNotesUrl** | Pointer to **string** | The url for the release notes of this image. | [optional] 
**SoftwareTypeId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | The vendor or publisher of this file. | [optional] 
**DistributableMetas** | Pointer to [**[]FirmwareDistributableMetaRelationship**](firmware.DistributableMeta.Relationship.md) | An array of relationships to firmwareDistributableMeta resources. | [optional] 
**Release** | Pointer to [**SoftwarerepositoryReleaseRelationship**](softwarerepository.Release.Relationship.md) |  | [optional] 
**FileLocation** | Pointer to **string** | The file location of the distributable. | [optional] 
**ImageCategory** | Pointer to **string** | The category into which the distributable falls into according to the supported platform series. For e.g.; C-Series/B-Series/Infrastructure. | [optional] 
**Origin** | Pointer to **string** | The source of the distributable. If it has been created by the user or system. * &#x60;System&#x60; - The distributable has been created by the System. * &#x60;User&#x60; - The distributable has been created by the User. | [optional] [default to "System"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareDistributableRelationship

`func NewFirmwareDistributableRelationship(classId string, objectType string, ) *FirmwareDistributableRelationship`

NewFirmwareDistributableRelationship instantiates a new FirmwareDistributableRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDistributableRelationshipWithDefaults

`func NewFirmwareDistributableRelationshipWithDefaults() *FirmwareDistributableRelationship`

NewFirmwareDistributableRelationshipWithDefaults instantiates a new FirmwareDistributableRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *FirmwareDistributableRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *FirmwareDistributableRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *FirmwareDistributableRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *FirmwareDistributableRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *FirmwareDistributableRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareDistributableRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareDistributableRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *FirmwareDistributableRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FirmwareDistributableRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FirmwareDistributableRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FirmwareDistributableRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *FirmwareDistributableRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *FirmwareDistributableRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *FirmwareDistributableRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *FirmwareDistributableRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *FirmwareDistributableRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *FirmwareDistributableRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *FirmwareDistributableRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *FirmwareDistributableRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *FirmwareDistributableRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *FirmwareDistributableRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *FirmwareDistributableRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *FirmwareDistributableRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *FirmwareDistributableRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareDistributableRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareDistributableRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *FirmwareDistributableRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *FirmwareDistributableRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *FirmwareDistributableRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *FirmwareDistributableRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *FirmwareDistributableRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *FirmwareDistributableRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *FirmwareDistributableRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *FirmwareDistributableRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *FirmwareDistributableRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FirmwareDistributableRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FirmwareDistributableRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FirmwareDistributableRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *FirmwareDistributableRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *FirmwareDistributableRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *FirmwareDistributableRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *FirmwareDistributableRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *FirmwareDistributableRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *FirmwareDistributableRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *FirmwareDistributableRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *FirmwareDistributableRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *FirmwareDistributableRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *FirmwareDistributableRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *FirmwareDistributableRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FirmwareDistributableRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FirmwareDistributableRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FirmwareDistributableRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *FirmwareDistributableRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *FirmwareDistributableRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *FirmwareDistributableRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *FirmwareDistributableRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *FirmwareDistributableRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *FirmwareDistributableRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *FirmwareDistributableRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *FirmwareDistributableRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *FirmwareDistributableRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *FirmwareDistributableRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *FirmwareDistributableRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *FirmwareDistributableRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *FirmwareDistributableRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirmwareDistributableRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirmwareDistributableRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirmwareDistributableRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadCount

`func (o *FirmwareDistributableRelationship) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *FirmwareDistributableRelationship) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *FirmwareDistributableRelationship) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *FirmwareDistributableRelationship) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetImportAction

`func (o *FirmwareDistributableRelationship) GetImportAction() string`

GetImportAction returns the ImportAction field if non-nil, zero value otherwise.

### GetImportActionOk

`func (o *FirmwareDistributableRelationship) GetImportActionOk() (*string, bool)`

GetImportActionOk returns a tuple with the ImportAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAction

`func (o *FirmwareDistributableRelationship) SetImportAction(v string)`

SetImportAction sets ImportAction field to given value.

### HasImportAction

`func (o *FirmwareDistributableRelationship) HasImportAction() bool`

HasImportAction returns a boolean if a field has been set.

### GetImportState

`func (o *FirmwareDistributableRelationship) GetImportState() string`

GetImportState returns the ImportState field if non-nil, zero value otherwise.

### GetImportStateOk

`func (o *FirmwareDistributableRelationship) GetImportStateOk() (*string, bool)`

GetImportStateOk returns a tuple with the ImportState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportState

`func (o *FirmwareDistributableRelationship) SetImportState(v string)`

SetImportState sets ImportState field to given value.

### HasImportState

`func (o *FirmwareDistributableRelationship) HasImportState() bool`

HasImportState returns a boolean if a field has been set.

### GetImportedTime

`func (o *FirmwareDistributableRelationship) GetImportedTime() time.Time`

GetImportedTime returns the ImportedTime field if non-nil, zero value otherwise.

### GetImportedTimeOk

`func (o *FirmwareDistributableRelationship) GetImportedTimeOk() (*time.Time, bool)`

GetImportedTimeOk returns a tuple with the ImportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedTime

`func (o *FirmwareDistributableRelationship) SetImportedTime(v time.Time)`

SetImportedTime sets ImportedTime field to given value.

### HasImportedTime

`func (o *FirmwareDistributableRelationship) HasImportedTime() bool`

HasImportedTime returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *FirmwareDistributableRelationship) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *FirmwareDistributableRelationship) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *FirmwareDistributableRelationship) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *FirmwareDistributableRelationship) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetMd5sum

`func (o *FirmwareDistributableRelationship) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *FirmwareDistributableRelationship) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *FirmwareDistributableRelationship) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *FirmwareDistributableRelationship) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetName

`func (o *FirmwareDistributableRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareDistributableRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareDistributableRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirmwareDistributableRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *FirmwareDistributableRelationship) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *FirmwareDistributableRelationship) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *FirmwareDistributableRelationship) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *FirmwareDistributableRelationship) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSha512sum

`func (o *FirmwareDistributableRelationship) GetSha512sum() string`

GetSha512sum returns the Sha512sum field if non-nil, zero value otherwise.

### GetSha512sumOk

`func (o *FirmwareDistributableRelationship) GetSha512sumOk() (*string, bool)`

GetSha512sumOk returns a tuple with the Sha512sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512sum

`func (o *FirmwareDistributableRelationship) SetSha512sum(v string)`

SetSha512sum sets Sha512sum field to given value.

### HasSha512sum

`func (o *FirmwareDistributableRelationship) HasSha512sum() bool`

HasSha512sum returns a boolean if a field has been set.

### GetSize

`func (o *FirmwareDistributableRelationship) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirmwareDistributableRelationship) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirmwareDistributableRelationship) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FirmwareDistributableRelationship) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSoftwareAdvisoryUrl

`func (o *FirmwareDistributableRelationship) GetSoftwareAdvisoryUrl() string`

GetSoftwareAdvisoryUrl returns the SoftwareAdvisoryUrl field if non-nil, zero value otherwise.

### GetSoftwareAdvisoryUrlOk

`func (o *FirmwareDistributableRelationship) GetSoftwareAdvisoryUrlOk() (*string, bool)`

GetSoftwareAdvisoryUrlOk returns a tuple with the SoftwareAdvisoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAdvisoryUrl

`func (o *FirmwareDistributableRelationship) SetSoftwareAdvisoryUrl(v string)`

SetSoftwareAdvisoryUrl sets SoftwareAdvisoryUrl field to given value.

### HasSoftwareAdvisoryUrl

`func (o *FirmwareDistributableRelationship) HasSoftwareAdvisoryUrl() bool`

HasSoftwareAdvisoryUrl returns a boolean if a field has been set.

### GetSource

`func (o *FirmwareDistributableRelationship) GetSource() SoftwarerepositoryFileServer`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FirmwareDistributableRelationship) GetSourceOk() (*SoftwarerepositoryFileServer, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FirmwareDistributableRelationship) SetSource(v SoftwarerepositoryFileServer)`

SetSource sets Source field to given value.

### HasSource

`func (o *FirmwareDistributableRelationship) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *FirmwareDistributableRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FirmwareDistributableRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FirmwareDistributableRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FirmwareDistributableRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBundleType

`func (o *FirmwareDistributableRelationship) GetBundleType() string`

GetBundleType returns the BundleType field if non-nil, zero value otherwise.

### GetBundleTypeOk

`func (o *FirmwareDistributableRelationship) GetBundleTypeOk() (*string, bool)`

GetBundleTypeOk returns a tuple with the BundleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleType

`func (o *FirmwareDistributableRelationship) SetBundleType(v string)`

SetBundleType sets BundleType field to given value.

### HasBundleType

`func (o *FirmwareDistributableRelationship) HasBundleType() bool`

HasBundleType returns a boolean if a field has been set.

### GetComponentMeta

`func (o *FirmwareDistributableRelationship) GetComponentMeta() []FirmwareComponentMeta`

GetComponentMeta returns the ComponentMeta field if non-nil, zero value otherwise.

### GetComponentMetaOk

`func (o *FirmwareDistributableRelationship) GetComponentMetaOk() (*[]FirmwareComponentMeta, bool)`

GetComponentMetaOk returns a tuple with the ComponentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMeta

`func (o *FirmwareDistributableRelationship) SetComponentMeta(v []FirmwareComponentMeta)`

SetComponentMeta sets ComponentMeta field to given value.

### HasComponentMeta

`func (o *FirmwareDistributableRelationship) HasComponentMeta() bool`

HasComponentMeta returns a boolean if a field has been set.

### GetGuid

`func (o *FirmwareDistributableRelationship) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *FirmwareDistributableRelationship) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *FirmwareDistributableRelationship) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *FirmwareDistributableRelationship) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetMdfid

`func (o *FirmwareDistributableRelationship) GetMdfid() string`

GetMdfid returns the Mdfid field if non-nil, zero value otherwise.

### GetMdfidOk

`func (o *FirmwareDistributableRelationship) GetMdfidOk() (*string, bool)`

GetMdfidOk returns a tuple with the Mdfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfid

`func (o *FirmwareDistributableRelationship) SetMdfid(v string)`

SetMdfid sets Mdfid field to given value.

### HasMdfid

`func (o *FirmwareDistributableRelationship) HasMdfid() bool`

HasMdfid returns a boolean if a field has been set.

### GetModel

`func (o *FirmwareDistributableRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FirmwareDistributableRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FirmwareDistributableRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FirmwareDistributableRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPlatformType

`func (o *FirmwareDistributableRelationship) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *FirmwareDistributableRelationship) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *FirmwareDistributableRelationship) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *FirmwareDistributableRelationship) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRecommendedBuild

`func (o *FirmwareDistributableRelationship) GetRecommendedBuild() string`

GetRecommendedBuild returns the RecommendedBuild field if non-nil, zero value otherwise.

### GetRecommendedBuildOk

`func (o *FirmwareDistributableRelationship) GetRecommendedBuildOk() (*string, bool)`

GetRecommendedBuildOk returns a tuple with the RecommendedBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedBuild

`func (o *FirmwareDistributableRelationship) SetRecommendedBuild(v string)`

SetRecommendedBuild sets RecommendedBuild field to given value.

### HasRecommendedBuild

`func (o *FirmwareDistributableRelationship) HasRecommendedBuild() bool`

HasRecommendedBuild returns a boolean if a field has been set.

### GetReleaseNotesUrl

`func (o *FirmwareDistributableRelationship) GetReleaseNotesUrl() string`

GetReleaseNotesUrl returns the ReleaseNotesUrl field if non-nil, zero value otherwise.

### GetReleaseNotesUrlOk

`func (o *FirmwareDistributableRelationship) GetReleaseNotesUrlOk() (*string, bool)`

GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotesUrl

`func (o *FirmwareDistributableRelationship) SetReleaseNotesUrl(v string)`

SetReleaseNotesUrl sets ReleaseNotesUrl field to given value.

### HasReleaseNotesUrl

`func (o *FirmwareDistributableRelationship) HasReleaseNotesUrl() bool`

HasReleaseNotesUrl returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *FirmwareDistributableRelationship) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *FirmwareDistributableRelationship) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *FirmwareDistributableRelationship) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *FirmwareDistributableRelationship) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *FirmwareDistributableRelationship) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *FirmwareDistributableRelationship) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *FirmwareDistributableRelationship) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *FirmwareDistributableRelationship) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### GetVendor

`func (o *FirmwareDistributableRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FirmwareDistributableRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FirmwareDistributableRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *FirmwareDistributableRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDistributableMetas

`func (o *FirmwareDistributableRelationship) GetDistributableMetas() []FirmwareDistributableMetaRelationship`

GetDistributableMetas returns the DistributableMetas field if non-nil, zero value otherwise.

### GetDistributableMetasOk

`func (o *FirmwareDistributableRelationship) GetDistributableMetasOk() (*[]FirmwareDistributableMetaRelationship, bool)`

GetDistributableMetasOk returns a tuple with the DistributableMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributableMetas

`func (o *FirmwareDistributableRelationship) SetDistributableMetas(v []FirmwareDistributableMetaRelationship)`

SetDistributableMetas sets DistributableMetas field to given value.

### HasDistributableMetas

`func (o *FirmwareDistributableRelationship) HasDistributableMetas() bool`

HasDistributableMetas returns a boolean if a field has been set.

### SetDistributableMetasNil

`func (o *FirmwareDistributableRelationship) SetDistributableMetasNil(b bool)`

 SetDistributableMetasNil sets the value for DistributableMetas to be an explicit nil

### UnsetDistributableMetas
`func (o *FirmwareDistributableRelationship) UnsetDistributableMetas()`

UnsetDistributableMetas ensures that no value is present for DistributableMetas, not even an explicit nil
### GetRelease

`func (o *FirmwareDistributableRelationship) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *FirmwareDistributableRelationship) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *FirmwareDistributableRelationship) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *FirmwareDistributableRelationship) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetFileLocation

`func (o *FirmwareDistributableRelationship) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *FirmwareDistributableRelationship) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *FirmwareDistributableRelationship) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *FirmwareDistributableRelationship) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetImageCategory

`func (o *FirmwareDistributableRelationship) GetImageCategory() string`

GetImageCategory returns the ImageCategory field if non-nil, zero value otherwise.

### GetImageCategoryOk

`func (o *FirmwareDistributableRelationship) GetImageCategoryOk() (*string, bool)`

GetImageCategoryOk returns a tuple with the ImageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCategory

`func (o *FirmwareDistributableRelationship) SetImageCategory(v string)`

SetImageCategory sets ImageCategory field to given value.

### HasImageCategory

`func (o *FirmwareDistributableRelationship) HasImageCategory() bool`

HasImageCategory returns a boolean if a field has been set.

### GetOrigin

`func (o *FirmwareDistributableRelationship) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *FirmwareDistributableRelationship) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *FirmwareDistributableRelationship) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *FirmwareDistributableRelationship) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetCatalog

`func (o *FirmwareDistributableRelationship) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FirmwareDistributableRelationship) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FirmwareDistributableRelationship) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FirmwareDistributableRelationship) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


