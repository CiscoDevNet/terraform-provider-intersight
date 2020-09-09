# SoftwarerepositoryRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseDate** | Pointer to [**time.Time**](time.Time.md) | The date when the file was released or distributed by its vendor. | [optional] 
**ReleaseNotesUrl** | Pointer to **string** | The URL for the release notes of this image. | [optional] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | The platform type for which the images are released. This can be a fabric interconnect or compute server hardware. * &#x60;FabricSwitch&#x60; - The images in a release that correspond to Fabric Interconnect switches. * &#x60;ComputeSystem&#x60; - The images in a release that correspond to servers. | [optional] [default to "FabricSwitch"]
**Version** | Pointer to **string** | Cisco provided release version. | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryRelease

`func NewSoftwarerepositoryRelease() *SoftwarerepositoryRelease`

NewSoftwarerepositoryRelease instantiates a new SoftwarerepositoryRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryReleaseWithDefaults

`func NewSoftwarerepositoryReleaseWithDefaults() *SoftwarerepositoryRelease`

NewSoftwarerepositoryReleaseWithDefaults instantiates a new SoftwarerepositoryRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseDate

`func (o *SoftwarerepositoryRelease) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwarerepositoryRelease) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwarerepositoryRelease) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwarerepositoryRelease) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleaseNotesUrl

`func (o *SoftwarerepositoryRelease) GetReleaseNotesUrl() string`

GetReleaseNotesUrl returns the ReleaseNotesUrl field if non-nil, zero value otherwise.

### GetReleaseNotesUrlOk

`func (o *SoftwarerepositoryRelease) GetReleaseNotesUrlOk() (*string, bool)`

GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotesUrl

`func (o *SoftwarerepositoryRelease) SetReleaseNotesUrl(v string)`

SetReleaseNotesUrl sets ReleaseNotesUrl field to given value.

### HasReleaseNotesUrl

`func (o *SoftwarerepositoryRelease) HasReleaseNotesUrl() bool`

HasReleaseNotesUrl returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryRelease) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryRelease) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryRelease) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryRelease) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### GetType

`func (o *SoftwarerepositoryRelease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SoftwarerepositoryRelease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SoftwarerepositoryRelease) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SoftwarerepositoryRelease) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwarerepositoryRelease) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryRelease) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryRelease) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryRelease) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwarerepositoryRelease) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwarerepositoryRelease) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwarerepositoryRelease) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwarerepositoryRelease) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


