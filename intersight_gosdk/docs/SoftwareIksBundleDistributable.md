# SoftwareIksBundleDistributable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.IksBundleDistributable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.IksBundleDistributable"]
**Catalog** | Pointer to [**NullableSoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 
**Images** | Pointer to [**[]SoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) | An array of relationships to softwareSolutionDistributable resources. | [optional] [readonly] 

## Methods

### NewSoftwareIksBundleDistributable

`func NewSoftwareIksBundleDistributable(classId string, objectType string, ) *SoftwareIksBundleDistributable`

NewSoftwareIksBundleDistributable instantiates a new SoftwareIksBundleDistributable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareIksBundleDistributableWithDefaults

`func NewSoftwareIksBundleDistributableWithDefaults() *SoftwareIksBundleDistributable`

NewSoftwareIksBundleDistributableWithDefaults instantiates a new SoftwareIksBundleDistributable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareIksBundleDistributable) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareIksBundleDistributable) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareIksBundleDistributable) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareIksBundleDistributable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareIksBundleDistributable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareIksBundleDistributable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalog

`func (o *SoftwareIksBundleDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareIksBundleDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareIksBundleDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareIksBundleDistributable) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *SoftwareIksBundleDistributable) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *SoftwareIksBundleDistributable) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetImages

`func (o *SoftwareIksBundleDistributable) GetImages() []SoftwareSolutionDistributableRelationship`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SoftwareIksBundleDistributable) GetImagesOk() (*[]SoftwareSolutionDistributableRelationship, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SoftwareIksBundleDistributable) SetImages(v []SoftwareSolutionDistributableRelationship)`

SetImages sets Images field to given value.

### HasImages

`func (o *SoftwareIksBundleDistributable) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *SoftwareIksBundleDistributable) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *SoftwareIksBundleDistributable) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


