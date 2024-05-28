# WorkflowUiFormConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.UiFormConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.UiFormConfig"]
**Config** | Pointer to **interface{}** | Schema of the UI config for the set of defined datatypes. | [optional] 
**Name** | Pointer to **string** | Name of the entity inside the meta definition for which a collection of workflow datatype definition is added. Use StartTask as the name for the input definition inside a workflow definition or if workflow definition contains user action task, then it will be the name of the user action task. | [optional] 

## Methods

### NewWorkflowUiFormConfig

`func NewWorkflowUiFormConfig(classId string, objectType string, ) *WorkflowUiFormConfig`

NewWorkflowUiFormConfig instantiates a new WorkflowUiFormConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUiFormConfigWithDefaults

`func NewWorkflowUiFormConfigWithDefaults() *WorkflowUiFormConfig`

NewWorkflowUiFormConfigWithDefaults instantiates a new WorkflowUiFormConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowUiFormConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowUiFormConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowUiFormConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowUiFormConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowUiFormConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowUiFormConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfig

`func (o *WorkflowUiFormConfig) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WorkflowUiFormConfig) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WorkflowUiFormConfig) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WorkflowUiFormConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *WorkflowUiFormConfig) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *WorkflowUiFormConfig) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetName

`func (o *WorkflowUiFormConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowUiFormConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowUiFormConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowUiFormConfig) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


