# SoftwareUcsdBundleDistributableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.UcsdBundleDistributable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.UcsdBundleDistributable"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 
**Images** | Pointer to [**[]SoftwareUcsdDistributableRelationship**](SoftwareUcsdDistributableRelationship.md) | An array of relationships to softwareUcsdDistributable resources. | [optional] [readonly] 

## Methods

### NewSoftwareUcsdBundleDistributableAllOf

`func NewSoftwareUcsdBundleDistributableAllOf(classId string, objectType string, ) *SoftwareUcsdBundleDistributableAllOf`

NewSoftwareUcsdBundleDistributableAllOf instantiates a new SoftwareUcsdBundleDistributableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareUcsdBundleDistributableAllOfWithDefaults

`func NewSoftwareUcsdBundleDistributableAllOfWithDefaults() *SoftwareUcsdBundleDistributableAllOf`

NewSoftwareUcsdBundleDistributableAllOfWithDefaults instantiates a new SoftwareUcsdBundleDistributableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareUcsdBundleDistributableAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareUcsdBundleDistributableAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareUcsdBundleDistributableAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareUcsdBundleDistributableAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareUcsdBundleDistributableAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareUcsdBundleDistributableAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalog

`func (o *SoftwareUcsdBundleDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareUcsdBundleDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareUcsdBundleDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareUcsdBundleDistributableAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetImages

`func (o *SoftwareUcsdBundleDistributableAllOf) GetImages() []SoftwareUcsdDistributableRelationship`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SoftwareUcsdBundleDistributableAllOf) GetImagesOk() (*[]SoftwareUcsdDistributableRelationship, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SoftwareUcsdBundleDistributableAllOf) SetImages(v []SoftwareUcsdDistributableRelationship)`

SetImages sets Images field to given value.

### HasImages

`func (o *SoftwareUcsdBundleDistributableAllOf) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *SoftwareUcsdBundleDistributableAllOf) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *SoftwareUcsdBundleDistributableAllOf) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


