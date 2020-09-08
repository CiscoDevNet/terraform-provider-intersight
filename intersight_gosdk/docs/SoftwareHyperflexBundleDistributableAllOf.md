# SoftwareHyperflexBundleDistributableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 
**Images** | Pointer to [**[]SoftwareHyperflexDistributableRelationship**](software.HyperflexDistributable.Relationship.md) | An array of relationships to softwareHyperflexDistributable resources. | [optional] [readonly] 

## Methods

### NewSoftwareHyperflexBundleDistributableAllOf

`func NewSoftwareHyperflexBundleDistributableAllOf() *SoftwareHyperflexBundleDistributableAllOf`

NewSoftwareHyperflexBundleDistributableAllOf instantiates a new SoftwareHyperflexBundleDistributableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareHyperflexBundleDistributableAllOfWithDefaults

`func NewSoftwareHyperflexBundleDistributableAllOfWithDefaults() *SoftwareHyperflexBundleDistributableAllOf`

NewSoftwareHyperflexBundleDistributableAllOfWithDefaults instantiates a new SoftwareHyperflexBundleDistributableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *SoftwareHyperflexBundleDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareHyperflexBundleDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareHyperflexBundleDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareHyperflexBundleDistributableAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetImages

`func (o *SoftwareHyperflexBundleDistributableAllOf) GetImages() []SoftwareHyperflexDistributableRelationship`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SoftwareHyperflexBundleDistributableAllOf) GetImagesOk() (*[]SoftwareHyperflexDistributableRelationship, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SoftwareHyperflexBundleDistributableAllOf) SetImages(v []SoftwareHyperflexDistributableRelationship)`

SetImages sets Images field to given value.

### HasImages

`func (o *SoftwareHyperflexBundleDistributableAllOf) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *SoftwareHyperflexBundleDistributableAllOf) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *SoftwareHyperflexBundleDistributableAllOf) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


