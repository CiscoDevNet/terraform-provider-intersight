# WorkflowTaskConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetDataType** | Pointer to **interface{}** | List of property constraints that helps to narrow down task implementations based on target device input. | [optional] 

## Methods

### NewWorkflowTaskConstraints

`func NewWorkflowTaskConstraints() *WorkflowTaskConstraints`

NewWorkflowTaskConstraints instantiates a new WorkflowTaskConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskConstraintsWithDefaults

`func NewWorkflowTaskConstraintsWithDefaults() *WorkflowTaskConstraints`

NewWorkflowTaskConstraintsWithDefaults instantiates a new WorkflowTaskConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetDataType

`func (o *WorkflowTaskConstraints) GetTargetDataType() interface{}`

GetTargetDataType returns the TargetDataType field if non-nil, zero value otherwise.

### GetTargetDataTypeOk

`func (o *WorkflowTaskConstraints) GetTargetDataTypeOk() (*interface{}, bool)`

GetTargetDataTypeOk returns a tuple with the TargetDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataType

`func (o *WorkflowTaskConstraints) SetTargetDataType(v interface{})`

SetTargetDataType sets TargetDataType field to given value.

### HasTargetDataType

`func (o *WorkflowTaskConstraints) HasTargetDataType() bool`

HasTargetDataType returns a boolean if a field has been set.

### SetTargetDataTypeNil

`func (o *WorkflowTaskConstraints) SetTargetDataTypeNil(b bool)`

 SetTargetDataTypeNil sets the value for TargetDataType to be an explicit nil

### UnsetTargetDataType
`func (o *WorkflowTaskConstraints) UnsetTargetDataType()`

UnsetTargetDataType ensures that no value is present for TargetDataType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


