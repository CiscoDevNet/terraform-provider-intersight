# WorkflowTargetPropertyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorAttribute** | Pointer to **string** | A singleton value which will contain the path to connector object from the selected object. | [optional] 
**ConstraintAttributes** | Pointer to **[]string** |  | [optional] 
**DisplayAttributes** | Pointer to **[]string** |  | [optional] 
**Selector** | Pointer to **string** | Field to hold an Intersight API along with an optional filter to narrow down the search options for target device. | [optional] 
**SupportedObjects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkflowTargetPropertyAllOf

`func NewWorkflowTargetPropertyAllOf() *WorkflowTargetPropertyAllOf`

NewWorkflowTargetPropertyAllOf instantiates a new WorkflowTargetPropertyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetPropertyAllOfWithDefaults

`func NewWorkflowTargetPropertyAllOfWithDefaults() *WorkflowTargetPropertyAllOf`

NewWorkflowTargetPropertyAllOfWithDefaults instantiates a new WorkflowTargetPropertyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorAttribute

`func (o *WorkflowTargetPropertyAllOf) GetConnectorAttribute() string`

GetConnectorAttribute returns the ConnectorAttribute field if non-nil, zero value otherwise.

### GetConnectorAttributeOk

`func (o *WorkflowTargetPropertyAllOf) GetConnectorAttributeOk() (*string, bool)`

GetConnectorAttributeOk returns a tuple with the ConnectorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorAttribute

`func (o *WorkflowTargetPropertyAllOf) SetConnectorAttribute(v string)`

SetConnectorAttribute sets ConnectorAttribute field to given value.

### HasConnectorAttribute

`func (o *WorkflowTargetPropertyAllOf) HasConnectorAttribute() bool`

HasConnectorAttribute returns a boolean if a field has been set.

### GetConstraintAttributes

`func (o *WorkflowTargetPropertyAllOf) GetConstraintAttributes() []string`

GetConstraintAttributes returns the ConstraintAttributes field if non-nil, zero value otherwise.

### GetConstraintAttributesOk

`func (o *WorkflowTargetPropertyAllOf) GetConstraintAttributesOk() (*[]string, bool)`

GetConstraintAttributesOk returns a tuple with the ConstraintAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintAttributes

`func (o *WorkflowTargetPropertyAllOf) SetConstraintAttributes(v []string)`

SetConstraintAttributes sets ConstraintAttributes field to given value.

### HasConstraintAttributes

`func (o *WorkflowTargetPropertyAllOf) HasConstraintAttributes() bool`

HasConstraintAttributes returns a boolean if a field has been set.

### GetDisplayAttributes

`func (o *WorkflowTargetPropertyAllOf) GetDisplayAttributes() []string`

GetDisplayAttributes returns the DisplayAttributes field if non-nil, zero value otherwise.

### GetDisplayAttributesOk

`func (o *WorkflowTargetPropertyAllOf) GetDisplayAttributesOk() (*[]string, bool)`

GetDisplayAttributesOk returns a tuple with the DisplayAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAttributes

`func (o *WorkflowTargetPropertyAllOf) SetDisplayAttributes(v []string)`

SetDisplayAttributes sets DisplayAttributes field to given value.

### HasDisplayAttributes

`func (o *WorkflowTargetPropertyAllOf) HasDisplayAttributes() bool`

HasDisplayAttributes returns a boolean if a field has been set.

### GetSelector

`func (o *WorkflowTargetPropertyAllOf) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *WorkflowTargetPropertyAllOf) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *WorkflowTargetPropertyAllOf) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *WorkflowTargetPropertyAllOf) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSupportedObjects

`func (o *WorkflowTargetPropertyAllOf) GetSupportedObjects() []string`

GetSupportedObjects returns the SupportedObjects field if non-nil, zero value otherwise.

### GetSupportedObjectsOk

`func (o *WorkflowTargetPropertyAllOf) GetSupportedObjectsOk() (*[]string, bool)`

GetSupportedObjectsOk returns a tuple with the SupportedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedObjects

`func (o *WorkflowTargetPropertyAllOf) SetSupportedObjects(v []string)`

SetSupportedObjects sets SupportedObjects field to given value.

### HasSupportedObjects

`func (o *WorkflowTargetPropertyAllOf) HasSupportedObjects() bool`

HasSupportedObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


