# WorkloadDeploymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.DeploymentInput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.DeploymentInput"]
**GenNumber** | Pointer to **int64** | A sequential number that increments whenever the input changes. | [optional] [readonly] 
**Input** | Pointer to [**[]WorkloadDeploymentBlueprintInputType**](WorkloadDeploymentBlueprintInputType.md) |  | [optional] 
**WorkloadDeployment** | Pointer to [**NullableWorkloadWorkloadDeploymentRelationship**](WorkloadWorkloadDeploymentRelationship.md) |  | [optional] 

## Methods

### NewWorkloadDeploymentInput

`func NewWorkloadDeploymentInput(classId string, objectType string, ) *WorkloadDeploymentInput`

NewWorkloadDeploymentInput instantiates a new WorkloadDeploymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDeploymentInputWithDefaults

`func NewWorkloadDeploymentInputWithDefaults() *WorkloadDeploymentInput`

NewWorkloadDeploymentInputWithDefaults instantiates a new WorkloadDeploymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadDeploymentInput) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadDeploymentInput) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadDeploymentInput) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadDeploymentInput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadDeploymentInput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadDeploymentInput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGenNumber

`func (o *WorkloadDeploymentInput) GetGenNumber() int64`

GetGenNumber returns the GenNumber field if non-nil, zero value otherwise.

### GetGenNumberOk

`func (o *WorkloadDeploymentInput) GetGenNumberOk() (*int64, bool)`

GetGenNumberOk returns a tuple with the GenNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenNumber

`func (o *WorkloadDeploymentInput) SetGenNumber(v int64)`

SetGenNumber sets GenNumber field to given value.

### HasGenNumber

`func (o *WorkloadDeploymentInput) HasGenNumber() bool`

HasGenNumber returns a boolean if a field has been set.

### GetInput

`func (o *WorkloadDeploymentInput) GetInput() []WorkloadDeploymentBlueprintInputType`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkloadDeploymentInput) GetInputOk() (*[]WorkloadDeploymentBlueprintInputType, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkloadDeploymentInput) SetInput(v []WorkloadDeploymentBlueprintInputType)`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkloadDeploymentInput) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkloadDeploymentInput) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkloadDeploymentInput) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetWorkloadDeployment

`func (o *WorkloadDeploymentInput) GetWorkloadDeployment() WorkloadWorkloadDeploymentRelationship`

GetWorkloadDeployment returns the WorkloadDeployment field if non-nil, zero value otherwise.

### GetWorkloadDeploymentOk

`func (o *WorkloadDeploymentInput) GetWorkloadDeploymentOk() (*WorkloadWorkloadDeploymentRelationship, bool)`

GetWorkloadDeploymentOk returns a tuple with the WorkloadDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDeployment

`func (o *WorkloadDeploymentInput) SetWorkloadDeployment(v WorkloadWorkloadDeploymentRelationship)`

SetWorkloadDeployment sets WorkloadDeployment field to given value.

### HasWorkloadDeployment

`func (o *WorkloadDeploymentInput) HasWorkloadDeployment() bool`

HasWorkloadDeployment returns a boolean if a field has been set.

### SetWorkloadDeploymentNil

`func (o *WorkloadDeploymentInput) SetWorkloadDeploymentNil(b bool)`

 SetWorkloadDeploymentNil sets the value for WorkloadDeployment to be an explicit nil

### UnsetWorkloadDeployment
`func (o *WorkloadDeploymentInput) UnsetWorkloadDeployment()`

UnsetWorkloadDeployment ensures that no value is present for WorkloadDeployment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


