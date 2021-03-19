# WorkflowDefaultValueAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DefaultValue"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DefaultValue"]
**IsValueSet** | Pointer to **bool** | A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses. | [optional] [readonly] 
**Override** | Pointer to **bool** | Override the default value provided for the data type. When true, allow the user to enter value for the data type. | [optional] 
**Value** | Pointer to **interface{}** | Default value for the data type. If default value was provided and the input was required the default value will be used as the input. | [optional] 

## Methods

### NewWorkflowDefaultValueAllOf

`func NewWorkflowDefaultValueAllOf(classId string, objectType string, ) *WorkflowDefaultValueAllOf`

NewWorkflowDefaultValueAllOf instantiates a new WorkflowDefaultValueAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDefaultValueAllOfWithDefaults

`func NewWorkflowDefaultValueAllOfWithDefaults() *WorkflowDefaultValueAllOf`

NewWorkflowDefaultValueAllOfWithDefaults instantiates a new WorkflowDefaultValueAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDefaultValueAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDefaultValueAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDefaultValueAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDefaultValueAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDefaultValueAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDefaultValueAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsValueSet

`func (o *WorkflowDefaultValueAllOf) GetIsValueSet() bool`

GetIsValueSet returns the IsValueSet field if non-nil, zero value otherwise.

### GetIsValueSetOk

`func (o *WorkflowDefaultValueAllOf) GetIsValueSetOk() (*bool, bool)`

GetIsValueSetOk returns a tuple with the IsValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueSet

`func (o *WorkflowDefaultValueAllOf) SetIsValueSet(v bool)`

SetIsValueSet sets IsValueSet field to given value.

### HasIsValueSet

`func (o *WorkflowDefaultValueAllOf) HasIsValueSet() bool`

HasIsValueSet returns a boolean if a field has been set.

### GetOverride

`func (o *WorkflowDefaultValueAllOf) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *WorkflowDefaultValueAllOf) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *WorkflowDefaultValueAllOf) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *WorkflowDefaultValueAllOf) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowDefaultValueAllOf) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowDefaultValueAllOf) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowDefaultValueAllOf) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowDefaultValueAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WorkflowDefaultValueAllOf) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WorkflowDefaultValueAllOf) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


