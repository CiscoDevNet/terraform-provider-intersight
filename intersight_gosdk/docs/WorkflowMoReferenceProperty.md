# WorkflowMoReferenceProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayAttributes** | Pointer to **[]string** |  | [optional] 
**Selector** | Pointer to **string** | Field to hold an Intersight API along with an optional filter to narrow down the search options. | [optional] 
**ValueAttribute** | Pointer to **string** | A property from the Intersight object, value of which can be used as value for referenced input definition. | [optional] 

## Methods

### NewWorkflowMoReferenceProperty

`func NewWorkflowMoReferenceProperty() *WorkflowMoReferenceProperty`

NewWorkflowMoReferenceProperty instantiates a new WorkflowMoReferenceProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoReferencePropertyWithDefaults

`func NewWorkflowMoReferencePropertyWithDefaults() *WorkflowMoReferenceProperty`

NewWorkflowMoReferencePropertyWithDefaults instantiates a new WorkflowMoReferenceProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayAttributes

`func (o *WorkflowMoReferenceProperty) GetDisplayAttributes() []string`

GetDisplayAttributes returns the DisplayAttributes field if non-nil, zero value otherwise.

### GetDisplayAttributesOk

`func (o *WorkflowMoReferenceProperty) GetDisplayAttributesOk() (*[]string, bool)`

GetDisplayAttributesOk returns a tuple with the DisplayAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAttributes

`func (o *WorkflowMoReferenceProperty) SetDisplayAttributes(v []string)`

SetDisplayAttributes sets DisplayAttributes field to given value.

### HasDisplayAttributes

`func (o *WorkflowMoReferenceProperty) HasDisplayAttributes() bool`

HasDisplayAttributes returns a boolean if a field has been set.

### GetSelector

`func (o *WorkflowMoReferenceProperty) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *WorkflowMoReferenceProperty) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *WorkflowMoReferenceProperty) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *WorkflowMoReferenceProperty) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetValueAttribute

`func (o *WorkflowMoReferenceProperty) GetValueAttribute() string`

GetValueAttribute returns the ValueAttribute field if non-nil, zero value otherwise.

### GetValueAttributeOk

`func (o *WorkflowMoReferenceProperty) GetValueAttributeOk() (*string, bool)`

GetValueAttributeOk returns a tuple with the ValueAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAttribute

`func (o *WorkflowMoReferenceProperty) SetValueAttribute(v string)`

SetValueAttribute sets ValueAttribute field to given value.

### HasValueAttribute

`func (o *WorkflowMoReferenceProperty) HasValueAttribute() bool`

HasValueAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


