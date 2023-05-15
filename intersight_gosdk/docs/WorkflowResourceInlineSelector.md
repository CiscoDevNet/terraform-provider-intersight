# WorkflowResourceInlineSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ResourceInlineSelector"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ResourceInlineSelector"]
**Selector** | Pointer to **string** | The odata based filter URL to select the resources. | [optional] 

## Methods

### NewWorkflowResourceInlineSelector

`func NewWorkflowResourceInlineSelector(classId string, objectType string, ) *WorkflowResourceInlineSelector`

NewWorkflowResourceInlineSelector instantiates a new WorkflowResourceInlineSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowResourceInlineSelectorWithDefaults

`func NewWorkflowResourceInlineSelectorWithDefaults() *WorkflowResourceInlineSelector`

NewWorkflowResourceInlineSelectorWithDefaults instantiates a new WorkflowResourceInlineSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowResourceInlineSelector) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowResourceInlineSelector) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowResourceInlineSelector) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowResourceInlineSelector) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowResourceInlineSelector) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowResourceInlineSelector) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSelector

`func (o *WorkflowResourceInlineSelector) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *WorkflowResourceInlineSelector) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *WorkflowResourceInlineSelector) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *WorkflowResourceInlineSelector) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


