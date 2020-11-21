# WorkflowTargetProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TargetProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TargetProperty"]
**ConnectorAttribute** | Pointer to **string** | A singleton value which will contain the path to connector object from the selected object. | [optional] 
**ConstraintAttributes** | Pointer to **[]string** |  | [optional] 
**DisplayAttributes** | Pointer to **[]string** |  | [optional] 
**Selector** | Pointer to **string** | Field to hold an Intersight API along with an optional filter to narrow down the search options for target device. | [optional] 
**SupportedObjects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkflowTargetProperty

`func NewWorkflowTargetProperty(classId string, objectType string, ) *WorkflowTargetProperty`

NewWorkflowTargetProperty instantiates a new WorkflowTargetProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetPropertyWithDefaults

`func NewWorkflowTargetPropertyWithDefaults() *WorkflowTargetProperty`

NewWorkflowTargetPropertyWithDefaults instantiates a new WorkflowTargetProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTargetProperty) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTargetProperty) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTargetProperty) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTargetProperty) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTargetProperty) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTargetProperty) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectorAttribute

`func (o *WorkflowTargetProperty) GetConnectorAttribute() string`

GetConnectorAttribute returns the ConnectorAttribute field if non-nil, zero value otherwise.

### GetConnectorAttributeOk

`func (o *WorkflowTargetProperty) GetConnectorAttributeOk() (*string, bool)`

GetConnectorAttributeOk returns a tuple with the ConnectorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorAttribute

`func (o *WorkflowTargetProperty) SetConnectorAttribute(v string)`

SetConnectorAttribute sets ConnectorAttribute field to given value.

### HasConnectorAttribute

`func (o *WorkflowTargetProperty) HasConnectorAttribute() bool`

HasConnectorAttribute returns a boolean if a field has been set.

### GetConstraintAttributes

`func (o *WorkflowTargetProperty) GetConstraintAttributes() []string`

GetConstraintAttributes returns the ConstraintAttributes field if non-nil, zero value otherwise.

### GetConstraintAttributesOk

`func (o *WorkflowTargetProperty) GetConstraintAttributesOk() (*[]string, bool)`

GetConstraintAttributesOk returns a tuple with the ConstraintAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintAttributes

`func (o *WorkflowTargetProperty) SetConstraintAttributes(v []string)`

SetConstraintAttributes sets ConstraintAttributes field to given value.

### HasConstraintAttributes

`func (o *WorkflowTargetProperty) HasConstraintAttributes() bool`

HasConstraintAttributes returns a boolean if a field has been set.

### SetConstraintAttributesNil

`func (o *WorkflowTargetProperty) SetConstraintAttributesNil(b bool)`

 SetConstraintAttributesNil sets the value for ConstraintAttributes to be an explicit nil

### UnsetConstraintAttributes
`func (o *WorkflowTargetProperty) UnsetConstraintAttributes()`

UnsetConstraintAttributes ensures that no value is present for ConstraintAttributes, not even an explicit nil
### GetDisplayAttributes

`func (o *WorkflowTargetProperty) GetDisplayAttributes() []string`

GetDisplayAttributes returns the DisplayAttributes field if non-nil, zero value otherwise.

### GetDisplayAttributesOk

`func (o *WorkflowTargetProperty) GetDisplayAttributesOk() (*[]string, bool)`

GetDisplayAttributesOk returns a tuple with the DisplayAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAttributes

`func (o *WorkflowTargetProperty) SetDisplayAttributes(v []string)`

SetDisplayAttributes sets DisplayAttributes field to given value.

### HasDisplayAttributes

`func (o *WorkflowTargetProperty) HasDisplayAttributes() bool`

HasDisplayAttributes returns a boolean if a field has been set.

### SetDisplayAttributesNil

`func (o *WorkflowTargetProperty) SetDisplayAttributesNil(b bool)`

 SetDisplayAttributesNil sets the value for DisplayAttributes to be an explicit nil

### UnsetDisplayAttributes
`func (o *WorkflowTargetProperty) UnsetDisplayAttributes()`

UnsetDisplayAttributes ensures that no value is present for DisplayAttributes, not even an explicit nil
### GetSelector

`func (o *WorkflowTargetProperty) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *WorkflowTargetProperty) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *WorkflowTargetProperty) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *WorkflowTargetProperty) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSupportedObjects

`func (o *WorkflowTargetProperty) GetSupportedObjects() []string`

GetSupportedObjects returns the SupportedObjects field if non-nil, zero value otherwise.

### GetSupportedObjectsOk

`func (o *WorkflowTargetProperty) GetSupportedObjectsOk() (*[]string, bool)`

GetSupportedObjectsOk returns a tuple with the SupportedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedObjects

`func (o *WorkflowTargetProperty) SetSupportedObjects(v []string)`

SetSupportedObjects sets SupportedObjects field to given value.

### HasSupportedObjects

`func (o *WorkflowTargetProperty) HasSupportedObjects() bool`

HasSupportedObjects returns a boolean if a field has been set.

### SetSupportedObjectsNil

`func (o *WorkflowTargetProperty) SetSupportedObjectsNil(b bool)`

 SetSupportedObjectsNil sets the value for SupportedObjects to be an explicit nil

### UnsetSupportedObjects
`func (o *WorkflowTargetProperty) UnsetSupportedObjects()`

UnsetSupportedObjects ensures that no value is present for SupportedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


