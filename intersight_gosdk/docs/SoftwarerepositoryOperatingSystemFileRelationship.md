# SoftwarerepositoryOperatingSystemFileRelationship

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
**Vendor** | Pointer to **string** | The vendor or publisher of this file. | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryOperatingSystemFileRelationship

`func NewSoftwarerepositoryOperatingSystemFileRelationship(classId string, objectType string, ) *SoftwarerepositoryOperatingSystemFileRelationship`

NewSoftwarerepositoryOperatingSystemFileRelationship instantiates a new SoftwarerepositoryOperatingSystemFileRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryOperatingSystemFileRelationshipWithDefaults

`func NewSoftwarerepositoryOperatingSystemFileRelationshipWithDefaults() *SoftwarerepositoryOperatingSystemFileRelationship`

NewSoftwarerepositoryOperatingSystemFileRelationshipWithDefaults instantiates a new SoftwarerepositoryOperatingSystemFileRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *SoftwarerepositoryOperatingSystemFileRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *SoftwarerepositoryOperatingSystemFileRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *SoftwarerepositoryOperatingSystemFileRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadCount

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetImportAction

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetImportAction() string`

GetImportAction returns the ImportAction field if non-nil, zero value otherwise.

### GetImportActionOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetImportActionOk() (*string, bool)`

GetImportActionOk returns a tuple with the ImportAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAction

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetImportAction(v string)`

SetImportAction sets ImportAction field to given value.

### HasImportAction

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasImportAction() bool`

HasImportAction returns a boolean if a field has been set.

### GetImportState

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetImportState() string`

GetImportState returns the ImportState field if non-nil, zero value otherwise.

### GetImportStateOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetImportStateOk() (*string, bool)`

GetImportStateOk returns a tuple with the ImportState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportState

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetImportState(v string)`

SetImportState sets ImportState field to given value.

### HasImportState

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasImportState() bool`

HasImportState returns a boolean if a field has been set.

### GetImportedTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetImportedTime() time.Time`

GetImportedTime returns the ImportedTime field if non-nil, zero value otherwise.

### GetImportedTimeOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetImportedTimeOk() (*time.Time, bool)`

GetImportedTimeOk returns a tuple with the ImportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetImportedTime(v time.Time)`

SetImportedTime sets ImportedTime field to given value.

### HasImportedTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasImportedTime() bool`

HasImportedTime returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetMd5sum

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetName

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSha512sum

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSha512sum() string`

GetSha512sum returns the Sha512sum field if non-nil, zero value otherwise.

### GetSha512sumOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSha512sumOk() (*string, bool)`

GetSha512sumOk returns a tuple with the Sha512sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512sum

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetSha512sum(v string)`

SetSha512sum sets Sha512sum field to given value.

### HasSha512sum

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasSha512sum() bool`

HasSha512sum returns a boolean if a field has been set.

### GetSize

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSoftwareAdvisoryUrl() string`

GetSoftwareAdvisoryUrl returns the SoftwareAdvisoryUrl field if non-nil, zero value otherwise.

### GetSoftwareAdvisoryUrlOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSoftwareAdvisoryUrlOk() (*string, bool)`

GetSoftwareAdvisoryUrlOk returns a tuple with the SoftwareAdvisoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetSoftwareAdvisoryUrl(v string)`

SetSoftwareAdvisoryUrl sets SoftwareAdvisoryUrl field to given value.

### HasSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasSoftwareAdvisoryUrl() bool`

HasSoftwareAdvisoryUrl returns a boolean if a field has been set.

### GetSource

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSource() SoftwarerepositoryFileServer`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetSourceOk() (*SoftwarerepositoryFileServer, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetSource(v SoftwarerepositoryFileServer)`

SetSource sets Source field to given value.

### HasSource

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVendor

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwarerepositoryOperatingSystemFileRelationship) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


