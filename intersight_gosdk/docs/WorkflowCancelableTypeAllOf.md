# WorkflowCancelableTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CancelableType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CancelableType"]
**CancelableStates** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** | When true the workflow can be cancelled. The action can be further restricted by the mode and cancelableStates properties. | [optional] [default to true]
**Mode** | Pointer to **string** | Mode controls how the workflow can be canceled. * &#x60;ApiOnly&#x60; - The workflow can only be canceled via API call. * &#x60;All&#x60; - The workflow can be canceled from API or from the user interface. | [optional] [default to "ApiOnly"]

## Methods

### NewWorkflowCancelableTypeAllOf

`func NewWorkflowCancelableTypeAllOf(classId string, objectType string, ) *WorkflowCancelableTypeAllOf`

NewWorkflowCancelableTypeAllOf instantiates a new WorkflowCancelableTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCancelableTypeAllOfWithDefaults

`func NewWorkflowCancelableTypeAllOfWithDefaults() *WorkflowCancelableTypeAllOf`

NewWorkflowCancelableTypeAllOfWithDefaults instantiates a new WorkflowCancelableTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCancelableTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCancelableTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCancelableTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCancelableTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCancelableTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCancelableTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCancelableStates

`func (o *WorkflowCancelableTypeAllOf) GetCancelableStates() []string`

GetCancelableStates returns the CancelableStates field if non-nil, zero value otherwise.

### GetCancelableStatesOk

`func (o *WorkflowCancelableTypeAllOf) GetCancelableStatesOk() (*[]string, bool)`

GetCancelableStatesOk returns a tuple with the CancelableStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelableStates

`func (o *WorkflowCancelableTypeAllOf) SetCancelableStates(v []string)`

SetCancelableStates sets CancelableStates field to given value.

### HasCancelableStates

`func (o *WorkflowCancelableTypeAllOf) HasCancelableStates() bool`

HasCancelableStates returns a boolean if a field has been set.

### SetCancelableStatesNil

`func (o *WorkflowCancelableTypeAllOf) SetCancelableStatesNil(b bool)`

 SetCancelableStatesNil sets the value for CancelableStates to be an explicit nil

### UnsetCancelableStates
`func (o *WorkflowCancelableTypeAllOf) UnsetCancelableStates()`

UnsetCancelableStates ensures that no value is present for CancelableStates, not even an explicit nil
### GetEnabled

`func (o *WorkflowCancelableTypeAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowCancelableTypeAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowCancelableTypeAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowCancelableTypeAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *WorkflowCancelableTypeAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WorkflowCancelableTypeAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WorkflowCancelableTypeAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WorkflowCancelableTypeAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


