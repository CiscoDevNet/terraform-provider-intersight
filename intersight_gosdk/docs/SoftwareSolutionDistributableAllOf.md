# SoftwareSolutionDistributableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.SolutionDistributable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.SolutionDistributable"]
**FilePath** | Pointer to **string** | The path of the file in S3/minio bucket. | [optional] [readonly] 
**SolutionName** | Pointer to **string** | The name of the solution in which the image belongs. | [optional] 
**SubType** | Pointer to **string** | The type of the file like OS image, Script etc. * &#x60;osimage&#x60; - The solution OS image for deployment. * &#x60;script&#x60; - The Python script for the solution VM configuration and deployment. | [optional] [default to "osimage"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewSoftwareSolutionDistributableAllOf

`func NewSoftwareSolutionDistributableAllOf(classId string, objectType string, ) *SoftwareSolutionDistributableAllOf`

NewSoftwareSolutionDistributableAllOf instantiates a new SoftwareSolutionDistributableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareSolutionDistributableAllOfWithDefaults

`func NewSoftwareSolutionDistributableAllOfWithDefaults() *SoftwareSolutionDistributableAllOf`

NewSoftwareSolutionDistributableAllOfWithDefaults instantiates a new SoftwareSolutionDistributableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareSolutionDistributableAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareSolutionDistributableAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareSolutionDistributableAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareSolutionDistributableAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareSolutionDistributableAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareSolutionDistributableAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilePath

`func (o *SoftwareSolutionDistributableAllOf) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SoftwareSolutionDistributableAllOf) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SoftwareSolutionDistributableAllOf) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SoftwareSolutionDistributableAllOf) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetSolutionName

`func (o *SoftwareSolutionDistributableAllOf) GetSolutionName() string`

GetSolutionName returns the SolutionName field if non-nil, zero value otherwise.

### GetSolutionNameOk

`func (o *SoftwareSolutionDistributableAllOf) GetSolutionNameOk() (*string, bool)`

GetSolutionNameOk returns a tuple with the SolutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionName

`func (o *SoftwareSolutionDistributableAllOf) SetSolutionName(v string)`

SetSolutionName sets SolutionName field to given value.

### HasSolutionName

`func (o *SoftwareSolutionDistributableAllOf) HasSolutionName() bool`

HasSolutionName returns a boolean if a field has been set.

### GetSubType

`func (o *SoftwareSolutionDistributableAllOf) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *SoftwareSolutionDistributableAllOf) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *SoftwareSolutionDistributableAllOf) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *SoftwareSolutionDistributableAllOf) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwareSolutionDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareSolutionDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareSolutionDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareSolutionDistributableAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


