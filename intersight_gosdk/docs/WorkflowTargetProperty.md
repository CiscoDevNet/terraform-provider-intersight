# WorkflowTargetProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorAttribute** | Pointer to **string** | A singleton value which will contain the path to connector object from the selected object. | [optional] 
**ConstraintAttributes** | Pointer to **[]string** |  | [optional] 
**DisplayAttributes** | Pointer to **[]string** |  | [optional] 
**Selector** | Pointer to **string** | Field to hold an Intersight API along with an optional filter to narrow down the search options for target device. | [optional] 
**SupportedObjects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkflowTargetProperty

`func NewWorkflowTargetProperty() *WorkflowTargetProperty`

NewWorkflowTargetProperty instantiates a new WorkflowTargetProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetPropertyWithDefaults

`func NewWorkflowTargetPropertyWithDefaults() *WorkflowTargetProperty`

NewWorkflowTargetPropertyWithDefaults instantiates a new WorkflowTargetProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


