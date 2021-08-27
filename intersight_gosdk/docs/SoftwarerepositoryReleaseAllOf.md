# SoftwarerepositoryReleaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.Release"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.Release"]
**ReleaseDate** | Pointer to **time.Time** | The date when the file was released or distributed by its vendor. | [optional] 
**ReleaseNotesUrl** | Pointer to **string** | The URL for the release notes of this image. | [optional] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware. * &#x60;FabricSwitch&#x60; - The images in a release that correspond to Fabric Interconnect switches. * &#x60;ComputeSystem&#x60; - The images in a release that correspond to servers. | [optional] [default to "FabricSwitch"]
**Version** | Pointer to **string** | Cisco provided release version. | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryReleaseAllOf

`func NewSoftwarerepositoryReleaseAllOf(classId string, objectType string, ) *SoftwarerepositoryReleaseAllOf`

NewSoftwarerepositoryReleaseAllOf instantiates a new SoftwarerepositoryReleaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryReleaseAllOfWithDefaults

`func NewSoftwarerepositoryReleaseAllOfWithDefaults() *SoftwarerepositoryReleaseAllOf`

NewSoftwarerepositoryReleaseAllOfWithDefaults instantiates a new SoftwarerepositoryReleaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryReleaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryReleaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryReleaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryReleaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryReleaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryReleaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReleaseDate

`func (o *SoftwarerepositoryReleaseAllOf) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwarerepositoryReleaseAllOf) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwarerepositoryReleaseAllOf) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwarerepositoryReleaseAllOf) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleaseNotesUrl

`func (o *SoftwarerepositoryReleaseAllOf) GetReleaseNotesUrl() string`

GetReleaseNotesUrl returns the ReleaseNotesUrl field if non-nil, zero value otherwise.

### GetReleaseNotesUrlOk

`func (o *SoftwarerepositoryReleaseAllOf) GetReleaseNotesUrlOk() (*string, bool)`

GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotesUrl

`func (o *SoftwarerepositoryReleaseAllOf) SetReleaseNotesUrl(v string)`

SetReleaseNotesUrl sets ReleaseNotesUrl field to given value.

### HasReleaseNotesUrl

`func (o *SoftwarerepositoryReleaseAllOf) HasReleaseNotesUrl() bool`

HasReleaseNotesUrl returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryReleaseAllOf) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryReleaseAllOf) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryReleaseAllOf) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryReleaseAllOf) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *SoftwarerepositoryReleaseAllOf) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *SoftwarerepositoryReleaseAllOf) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil
### GetType

`func (o *SoftwarerepositoryReleaseAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SoftwarerepositoryReleaseAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SoftwarerepositoryReleaseAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SoftwarerepositoryReleaseAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwarerepositoryReleaseAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryReleaseAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryReleaseAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryReleaseAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwarerepositoryReleaseAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwarerepositoryReleaseAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwarerepositoryReleaseAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwarerepositoryReleaseAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


