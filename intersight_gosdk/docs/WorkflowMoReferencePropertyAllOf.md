# WorkflowMoReferencePropertyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MoReferenceProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MoReferenceProperty"]
**DisplayAttributes** | Pointer to **[]string** |  | [optional] 
**Selector** | Pointer to **string** | Field to hold an Intersight API along with an optional filter to narrow down the search options. | [optional] 
**ValueAttribute** | Pointer to **string** | A property from the Intersight object, value of which can be used as value for referenced input definition. | [optional] 

## Methods

### NewWorkflowMoReferencePropertyAllOf

`func NewWorkflowMoReferencePropertyAllOf(classId string, objectType string, ) *WorkflowMoReferencePropertyAllOf`

NewWorkflowMoReferencePropertyAllOf instantiates a new WorkflowMoReferencePropertyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoReferencePropertyAllOfWithDefaults

`func NewWorkflowMoReferencePropertyAllOfWithDefaults() *WorkflowMoReferencePropertyAllOf`

NewWorkflowMoReferencePropertyAllOfWithDefaults instantiates a new WorkflowMoReferencePropertyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoReferencePropertyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoReferencePropertyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoReferencePropertyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoReferencePropertyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoReferencePropertyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoReferencePropertyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplayAttributes

`func (o *WorkflowMoReferencePropertyAllOf) GetDisplayAttributes() []string`

GetDisplayAttributes returns the DisplayAttributes field if non-nil, zero value otherwise.

### GetDisplayAttributesOk

`func (o *WorkflowMoReferencePropertyAllOf) GetDisplayAttributesOk() (*[]string, bool)`

GetDisplayAttributesOk returns a tuple with the DisplayAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAttributes

`func (o *WorkflowMoReferencePropertyAllOf) SetDisplayAttributes(v []string)`

SetDisplayAttributes sets DisplayAttributes field to given value.

### HasDisplayAttributes

`func (o *WorkflowMoReferencePropertyAllOf) HasDisplayAttributes() bool`

HasDisplayAttributes returns a boolean if a field has been set.

### SetDisplayAttributesNil

`func (o *WorkflowMoReferencePropertyAllOf) SetDisplayAttributesNil(b bool)`

 SetDisplayAttributesNil sets the value for DisplayAttributes to be an explicit nil

### UnsetDisplayAttributes
`func (o *WorkflowMoReferencePropertyAllOf) UnsetDisplayAttributes()`

UnsetDisplayAttributes ensures that no value is present for DisplayAttributes, not even an explicit nil
### GetSelector

`func (o *WorkflowMoReferencePropertyAllOf) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *WorkflowMoReferencePropertyAllOf) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *WorkflowMoReferencePropertyAllOf) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *WorkflowMoReferencePropertyAllOf) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetValueAttribute

`func (o *WorkflowMoReferencePropertyAllOf) GetValueAttribute() string`

GetValueAttribute returns the ValueAttribute field if non-nil, zero value otherwise.

### GetValueAttributeOk

`func (o *WorkflowMoReferencePropertyAllOf) GetValueAttributeOk() (*string, bool)`

GetValueAttributeOk returns a tuple with the ValueAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAttribute

`func (o *WorkflowMoReferencePropertyAllOf) SetValueAttribute(v string)`

SetValueAttribute sets ValueAttribute field to given value.

### HasValueAttribute

`func (o *WorkflowMoReferencePropertyAllOf) HasValueAttribute() bool`

HasValueAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


