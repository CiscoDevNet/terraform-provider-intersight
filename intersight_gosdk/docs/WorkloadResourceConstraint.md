# WorkloadResourceConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.ResourceConstraint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.ResourceConstraint"]
**Input** | Pointer to **interface{}** | The input values from the user for the resource definition of the blueprint. | [optional] 

## Methods

### NewWorkloadResourceConstraint

`func NewWorkloadResourceConstraint(classId string, objectType string, ) *WorkloadResourceConstraint`

NewWorkloadResourceConstraint instantiates a new WorkloadResourceConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadResourceConstraintWithDefaults

`func NewWorkloadResourceConstraintWithDefaults() *WorkloadResourceConstraint`

NewWorkloadResourceConstraintWithDefaults instantiates a new WorkloadResourceConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadResourceConstraint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadResourceConstraint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadResourceConstraint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadResourceConstraint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadResourceConstraint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadResourceConstraint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInput

`func (o *WorkloadResourceConstraint) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkloadResourceConstraint) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkloadResourceConstraint) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkloadResourceConstraint) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkloadResourceConstraint) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkloadResourceConstraint) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


