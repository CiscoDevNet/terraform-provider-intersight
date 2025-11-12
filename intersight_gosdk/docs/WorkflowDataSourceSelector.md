# WorkflowDataSourceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DataSourceSelector"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DataSourceSelector"]
**DisplayAttributes** | Pointer to **[]string** |  | [optional] 
**Selector** | Pointer to **string** | This field holds mapping information used to provide suggestions to the user. The mapping should be in the &#39;${workflow.input.property}&#39; format. It supports workflow input mapping for workflows, and for User Actions, it supports workflow inputs, workflow outputs, workflow variables, and outputs from previous tasks. | [optional] 
**ValueAttribute** | Pointer to **string** | A property from the mapped parameter, value of which can be used as value for referenced input definition. | [optional] 

## Methods

### NewWorkflowDataSourceSelector

`func NewWorkflowDataSourceSelector(classId string, objectType string, ) *WorkflowDataSourceSelector`

NewWorkflowDataSourceSelector instantiates a new WorkflowDataSourceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDataSourceSelectorWithDefaults

`func NewWorkflowDataSourceSelectorWithDefaults() *WorkflowDataSourceSelector`

NewWorkflowDataSourceSelectorWithDefaults instantiates a new WorkflowDataSourceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDataSourceSelector) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDataSourceSelector) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDataSourceSelector) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDataSourceSelector) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDataSourceSelector) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDataSourceSelector) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplayAttributes

`func (o *WorkflowDataSourceSelector) GetDisplayAttributes() []string`

GetDisplayAttributes returns the DisplayAttributes field if non-nil, zero value otherwise.

### GetDisplayAttributesOk

`func (o *WorkflowDataSourceSelector) GetDisplayAttributesOk() (*[]string, bool)`

GetDisplayAttributesOk returns a tuple with the DisplayAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAttributes

`func (o *WorkflowDataSourceSelector) SetDisplayAttributes(v []string)`

SetDisplayAttributes sets DisplayAttributes field to given value.

### HasDisplayAttributes

`func (o *WorkflowDataSourceSelector) HasDisplayAttributes() bool`

HasDisplayAttributes returns a boolean if a field has been set.

### SetDisplayAttributesNil

`func (o *WorkflowDataSourceSelector) SetDisplayAttributesNil(b bool)`

 SetDisplayAttributesNil sets the value for DisplayAttributes to be an explicit nil

### UnsetDisplayAttributes
`func (o *WorkflowDataSourceSelector) UnsetDisplayAttributes()`

UnsetDisplayAttributes ensures that no value is present for DisplayAttributes, not even an explicit nil
### GetSelector

`func (o *WorkflowDataSourceSelector) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *WorkflowDataSourceSelector) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *WorkflowDataSourceSelector) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *WorkflowDataSourceSelector) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetValueAttribute

`func (o *WorkflowDataSourceSelector) GetValueAttribute() string`

GetValueAttribute returns the ValueAttribute field if non-nil, zero value otherwise.

### GetValueAttributeOk

`func (o *WorkflowDataSourceSelector) GetValueAttributeOk() (*string, bool)`

GetValueAttributeOk returns a tuple with the ValueAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAttribute

`func (o *WorkflowDataSourceSelector) SetValueAttribute(v string)`

SetValueAttribute sets ValueAttribute field to given value.

### HasValueAttribute

`func (o *WorkflowDataSourceSelector) HasValueAttribute() bool`

HasValueAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


