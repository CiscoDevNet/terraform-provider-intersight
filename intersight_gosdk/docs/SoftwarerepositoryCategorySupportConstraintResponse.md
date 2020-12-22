# SoftwarerepositoryCategorySupportConstraintResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | A discriminator value to disambiguate the schema of a HTTP GET response body. | 
**Count** | Pointer to **int32** | The total number of &#39;softwarerepository.CategorySupportConstraint&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]MoTagKeySummary**](MoTagKeySummary.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryCategorySupportConstraintResponse

`func NewSoftwarerepositoryCategorySupportConstraintResponse(objectType string, ) *SoftwarerepositoryCategorySupportConstraintResponse`

NewSoftwarerepositoryCategorySupportConstraintResponse instantiates a new SoftwarerepositoryCategorySupportConstraintResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCategorySupportConstraintResponseWithDefaults

`func NewSoftwarerepositoryCategorySupportConstraintResponseWithDefaults() *SoftwarerepositoryCategorySupportConstraintResponse`

NewSoftwarerepositoryCategorySupportConstraintResponseWithDefaults instantiates a new SoftwarerepositoryCategorySupportConstraintResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) GetResults() []MoTagKeySummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) GetResultsOk() (*[]MoTagKeySummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) SetResults(v []MoTagKeySummary)`

SetResults sets Results field to given value.

### HasResults

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *SoftwarerepositoryCategorySupportConstraintResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SoftwarerepositoryCategorySupportConstraintResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


