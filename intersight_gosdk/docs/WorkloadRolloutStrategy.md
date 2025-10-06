# WorkloadRolloutStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**FailureThreshold** | Pointer to **int64** | Specifies no of errors can be allowed to skip executing next set. | [optional] [readonly] [default to 10]

## Methods

### NewWorkloadRolloutStrategy

`func NewWorkloadRolloutStrategy(classId string, objectType string, ) *WorkloadRolloutStrategy`

NewWorkloadRolloutStrategy instantiates a new WorkloadRolloutStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadRolloutStrategyWithDefaults

`func NewWorkloadRolloutStrategyWithDefaults() *WorkloadRolloutStrategy`

NewWorkloadRolloutStrategyWithDefaults instantiates a new WorkloadRolloutStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadRolloutStrategy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadRolloutStrategy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadRolloutStrategy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadRolloutStrategy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadRolloutStrategy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadRolloutStrategy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailureThreshold

`func (o *WorkloadRolloutStrategy) GetFailureThreshold() int64`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *WorkloadRolloutStrategy) GetFailureThresholdOk() (*int64, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *WorkloadRolloutStrategy) SetFailureThreshold(v int64)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *WorkloadRolloutStrategy) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


