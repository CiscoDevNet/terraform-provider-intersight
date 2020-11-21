# WorkflowDecisionCaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DecisionCase"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DecisionCase"]
**Description** | Pointer to **string** | Description of this decision case. | [optional] 
**NextTask** | Pointer to **string** | The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node. | [optional] 
**Value** | Pointer to **string** | Value for the decision case. | [optional] 

## Methods

### NewWorkflowDecisionCaseAllOf

`func NewWorkflowDecisionCaseAllOf(classId string, objectType string, ) *WorkflowDecisionCaseAllOf`

NewWorkflowDecisionCaseAllOf instantiates a new WorkflowDecisionCaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDecisionCaseAllOfWithDefaults

`func NewWorkflowDecisionCaseAllOfWithDefaults() *WorkflowDecisionCaseAllOf`

NewWorkflowDecisionCaseAllOfWithDefaults instantiates a new WorkflowDecisionCaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDecisionCaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDecisionCaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDecisionCaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDecisionCaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDecisionCaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDecisionCaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowDecisionCaseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowDecisionCaseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowDecisionCaseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowDecisionCaseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNextTask

`func (o *WorkflowDecisionCaseAllOf) GetNextTask() string`

GetNextTask returns the NextTask field if non-nil, zero value otherwise.

### GetNextTaskOk

`func (o *WorkflowDecisionCaseAllOf) GetNextTaskOk() (*string, bool)`

GetNextTaskOk returns a tuple with the NextTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTask

`func (o *WorkflowDecisionCaseAllOf) SetNextTask(v string)`

SetNextTask sets NextTask field to given value.

### HasNextTask

`func (o *WorkflowDecisionCaseAllOf) HasNextTask() bool`

HasNextTask returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowDecisionCaseAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowDecisionCaseAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowDecisionCaseAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowDecisionCaseAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


