# WorkloadStateAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.StateAggregation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.StateAggregation"]
**Count** | Pointer to **int64** | The total number of referenced objects included in the aggregation. | [optional] [readonly] 
**State** | Pointer to **string** | The overall aggregated state as a string, summarizing the status of all referenced objects. | [optional] [readonly] 

## Methods

### NewWorkloadStateAggregation

`func NewWorkloadStateAggregation(classId string, objectType string, ) *WorkloadStateAggregation`

NewWorkloadStateAggregation instantiates a new WorkloadStateAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadStateAggregationWithDefaults

`func NewWorkloadStateAggregationWithDefaults() *WorkloadStateAggregation`

NewWorkloadStateAggregationWithDefaults instantiates a new WorkloadStateAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadStateAggregation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadStateAggregation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadStateAggregation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadStateAggregation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadStateAggregation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadStateAggregation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *WorkloadStateAggregation) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkloadStateAggregation) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkloadStateAggregation) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *WorkloadStateAggregation) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetState

`func (o *WorkloadStateAggregation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkloadStateAggregation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkloadStateAggregation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WorkloadStateAggregation) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


