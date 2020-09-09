# WorkflowDefaultValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Override** | Pointer to **bool** | Override the default value provided for the data type. When true, allow the user to enter value for the data type. | [optional] 
**Value** | Pointer to **interface{}** | Default value for the data type. If default value was provided and the input was required the default value will be used as the input. | [optional] 

## Methods

### NewWorkflowDefaultValue

`func NewWorkflowDefaultValue() *WorkflowDefaultValue`

NewWorkflowDefaultValue instantiates a new WorkflowDefaultValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDefaultValueWithDefaults

`func NewWorkflowDefaultValueWithDefaults() *WorkflowDefaultValue`

NewWorkflowDefaultValueWithDefaults instantiates a new WorkflowDefaultValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverride

`func (o *WorkflowDefaultValue) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *WorkflowDefaultValue) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *WorkflowDefaultValue) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *WorkflowDefaultValue) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowDefaultValue) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowDefaultValue) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowDefaultValue) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowDefaultValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WorkflowDefaultValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WorkflowDefaultValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


