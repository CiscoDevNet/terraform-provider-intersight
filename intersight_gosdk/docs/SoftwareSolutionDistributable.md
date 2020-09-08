# SoftwareSolutionDistributable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | Pointer to **string** | The path of the file in S3/minio bucket. | [optional] [readonly] 
**SolutionName** | Pointer to **string** | The name of the solution in which the image belongs. | [optional] 
**SubType** | Pointer to **string** | The type of the file like OS image, Script etc. * &#x60;osimage&#x60; - The solution OS image for deployment. * &#x60;script&#x60; - The Python script for the solution VM configuration and deployment. | [optional] [default to "osimage"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewSoftwareSolutionDistributable

`func NewSoftwareSolutionDistributable() *SoftwareSolutionDistributable`

NewSoftwareSolutionDistributable instantiates a new SoftwareSolutionDistributable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareSolutionDistributableWithDefaults

`func NewSoftwareSolutionDistributableWithDefaults() *SoftwareSolutionDistributable`

NewSoftwareSolutionDistributableWithDefaults instantiates a new SoftwareSolutionDistributable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *SoftwareSolutionDistributable) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SoftwareSolutionDistributable) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SoftwareSolutionDistributable) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SoftwareSolutionDistributable) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetSolutionName

`func (o *SoftwareSolutionDistributable) GetSolutionName() string`

GetSolutionName returns the SolutionName field if non-nil, zero value otherwise.

### GetSolutionNameOk

`func (o *SoftwareSolutionDistributable) GetSolutionNameOk() (*string, bool)`

GetSolutionNameOk returns a tuple with the SolutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionName

`func (o *SoftwareSolutionDistributable) SetSolutionName(v string)`

SetSolutionName sets SolutionName field to given value.

### HasSolutionName

`func (o *SoftwareSolutionDistributable) HasSolutionName() bool`

HasSolutionName returns a boolean if a field has been set.

### GetSubType

`func (o *SoftwareSolutionDistributable) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *SoftwareSolutionDistributable) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *SoftwareSolutionDistributable) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *SoftwareSolutionDistributable) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwareSolutionDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareSolutionDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareSolutionDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareSolutionDistributable) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


