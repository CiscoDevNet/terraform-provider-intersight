# WorkflowParameterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ParameterSet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ParameterSet"]
**Condition** | Pointer to **string** | The condition to be evaluated. * &#x60;eq&#x60; - Checks if the values of the two parameters are equal. * &#x60;ne&#x60; - Checks if the values of the two parameters are not equal. * &#x60;contains&#x60; - Checks if the second parameter string value is a substring of the first parameter string value. * &#x60;matchesPattern&#x60; - Checks if a string matches a regular expression. | [optional] [default to "eq"]
**ControlParameter** | Pointer to **string** | Name of the controlling entity, whose value will be used for evaluating the parameter set. | [optional] 
**EnableParameters** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name for the parameter set.  Limited to 64 alphanumeric characters (upper and lower case), and special characters &#39;-&#39; and &#39;_&#39;. | [optional] 
**Value** | Pointer to **string** | The controlling parameter will be evaluated against this &#39;value&#39;. | [optional] 

## Methods

### NewWorkflowParameterSet

`func NewWorkflowParameterSet(classId string, objectType string, ) *WorkflowParameterSet`

NewWorkflowParameterSet instantiates a new WorkflowParameterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowParameterSetWithDefaults

`func NewWorkflowParameterSetWithDefaults() *WorkflowParameterSet`

NewWorkflowParameterSetWithDefaults instantiates a new WorkflowParameterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowParameterSet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowParameterSet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowParameterSet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowParameterSet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowParameterSet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowParameterSet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *WorkflowParameterSet) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WorkflowParameterSet) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WorkflowParameterSet) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WorkflowParameterSet) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetControlParameter

`func (o *WorkflowParameterSet) GetControlParameter() string`

GetControlParameter returns the ControlParameter field if non-nil, zero value otherwise.

### GetControlParameterOk

`func (o *WorkflowParameterSet) GetControlParameterOk() (*string, bool)`

GetControlParameterOk returns a tuple with the ControlParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlParameter

`func (o *WorkflowParameterSet) SetControlParameter(v string)`

SetControlParameter sets ControlParameter field to given value.

### HasControlParameter

`func (o *WorkflowParameterSet) HasControlParameter() bool`

HasControlParameter returns a boolean if a field has been set.

### GetEnableParameters

`func (o *WorkflowParameterSet) GetEnableParameters() []string`

GetEnableParameters returns the EnableParameters field if non-nil, zero value otherwise.

### GetEnableParametersOk

`func (o *WorkflowParameterSet) GetEnableParametersOk() (*[]string, bool)`

GetEnableParametersOk returns a tuple with the EnableParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParameters

`func (o *WorkflowParameterSet) SetEnableParameters(v []string)`

SetEnableParameters sets EnableParameters field to given value.

### HasEnableParameters

`func (o *WorkflowParameterSet) HasEnableParameters() bool`

HasEnableParameters returns a boolean if a field has been set.

### SetEnableParametersNil

`func (o *WorkflowParameterSet) SetEnableParametersNil(b bool)`

 SetEnableParametersNil sets the value for EnableParameters to be an explicit nil

### UnsetEnableParameters
`func (o *WorkflowParameterSet) UnsetEnableParameters()`

UnsetEnableParameters ensures that no value is present for EnableParameters, not even an explicit nil
### GetName

`func (o *WorkflowParameterSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowParameterSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowParameterSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowParameterSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowParameterSet) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowParameterSet) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowParameterSet) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowParameterSet) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


