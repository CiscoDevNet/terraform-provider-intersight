# WorkloadCanaryDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.CanaryDeployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.CanaryDeployment"]
**RolloutPercentages** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewWorkloadCanaryDeployment

`func NewWorkloadCanaryDeployment(classId string, objectType string, ) *WorkloadCanaryDeployment`

NewWorkloadCanaryDeployment instantiates a new WorkloadCanaryDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadCanaryDeploymentWithDefaults

`func NewWorkloadCanaryDeploymentWithDefaults() *WorkloadCanaryDeployment`

NewWorkloadCanaryDeploymentWithDefaults instantiates a new WorkloadCanaryDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadCanaryDeployment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadCanaryDeployment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadCanaryDeployment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadCanaryDeployment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadCanaryDeployment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadCanaryDeployment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRolloutPercentages

`func (o *WorkloadCanaryDeployment) GetRolloutPercentages() []int64`

GetRolloutPercentages returns the RolloutPercentages field if non-nil, zero value otherwise.

### GetRolloutPercentagesOk

`func (o *WorkloadCanaryDeployment) GetRolloutPercentagesOk() (*[]int64, bool)`

GetRolloutPercentagesOk returns a tuple with the RolloutPercentages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentages

`func (o *WorkloadCanaryDeployment) SetRolloutPercentages(v []int64)`

SetRolloutPercentages sets RolloutPercentages field to given value.

### HasRolloutPercentages

`func (o *WorkloadCanaryDeployment) HasRolloutPercentages() bool`

HasRolloutPercentages returns a boolean if a field has been set.

### SetRolloutPercentagesNil

`func (o *WorkloadCanaryDeployment) SetRolloutPercentagesNil(b bool)`

 SetRolloutPercentagesNil sets the value for RolloutPercentages to be an explicit nil

### UnsetRolloutPercentages
`func (o *WorkloadCanaryDeployment) UnsetRolloutPercentages()`

UnsetRolloutPercentages ensures that no value is present for RolloutPercentages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


