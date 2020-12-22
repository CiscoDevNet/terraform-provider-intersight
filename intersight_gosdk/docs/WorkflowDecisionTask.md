# WorkflowDecisionTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DecisionTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DecisionTask"]
**Condition** | Pointer to **string** | The condition to evaluate for this decision task. The condition can be a workflow or task variable or an JavaScript expression. Example value for condition could be a variable like \&quot;${task1.output.var1} or ${workflow.input.var2}\&quot; which evaluates to a value matching any of the decision case values. Example value for condition if it&#39;s an expression - \&quot;if ( ${task1.output.var1} ! &#x3D; null &amp;&amp; ${task1.output.var1} &gt; 0 ) &#39;true&#39;; else &#39;false&#39;; \&quot; which evaluates to &#39;true&#39; or &#39;false&#39; and will match one of the decision case values. You can also use JavaScript functions like indexOf, toUpperCase in the expression which will be evaluated by the expression evaluator. | [optional] 
**DecisionCases** | Pointer to [**[]WorkflowDecisionCase**](WorkflowDecisionCase.md) |  | [optional] 
**DefaultTask** | Pointer to **string** | The default next Task to execute if the decision cannot be evaluated to any of the DecisionCases. | [optional] 
**InputParameters** | Pointer to **interface{}** | This field is deprecated. Decision case conditions can be added using the workflow input or task output variables in the Condition field. Refer to Condition field for more details. | [optional] 

## Methods

### NewWorkflowDecisionTask

`func NewWorkflowDecisionTask(classId string, objectType string, ) *WorkflowDecisionTask`

NewWorkflowDecisionTask instantiates a new WorkflowDecisionTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDecisionTaskWithDefaults

`func NewWorkflowDecisionTaskWithDefaults() *WorkflowDecisionTask`

NewWorkflowDecisionTaskWithDefaults instantiates a new WorkflowDecisionTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDecisionTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDecisionTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDecisionTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDecisionTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDecisionTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDecisionTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *WorkflowDecisionTask) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WorkflowDecisionTask) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WorkflowDecisionTask) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WorkflowDecisionTask) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDecisionCases

`func (o *WorkflowDecisionTask) GetDecisionCases() []WorkflowDecisionCase`

GetDecisionCases returns the DecisionCases field if non-nil, zero value otherwise.

### GetDecisionCasesOk

`func (o *WorkflowDecisionTask) GetDecisionCasesOk() (*[]WorkflowDecisionCase, bool)`

GetDecisionCasesOk returns a tuple with the DecisionCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionCases

`func (o *WorkflowDecisionTask) SetDecisionCases(v []WorkflowDecisionCase)`

SetDecisionCases sets DecisionCases field to given value.

### HasDecisionCases

`func (o *WorkflowDecisionTask) HasDecisionCases() bool`

HasDecisionCases returns a boolean if a field has been set.

### SetDecisionCasesNil

`func (o *WorkflowDecisionTask) SetDecisionCasesNil(b bool)`

 SetDecisionCasesNil sets the value for DecisionCases to be an explicit nil

### UnsetDecisionCases
`func (o *WorkflowDecisionTask) UnsetDecisionCases()`

UnsetDecisionCases ensures that no value is present for DecisionCases, not even an explicit nil
### GetDefaultTask

`func (o *WorkflowDecisionTask) GetDefaultTask() string`

GetDefaultTask returns the DefaultTask field if non-nil, zero value otherwise.

### GetDefaultTaskOk

`func (o *WorkflowDecisionTask) GetDefaultTaskOk() (*string, bool)`

GetDefaultTaskOk returns a tuple with the DefaultTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTask

`func (o *WorkflowDecisionTask) SetDefaultTask(v string)`

SetDefaultTask sets DefaultTask field to given value.

### HasDefaultTask

`func (o *WorkflowDecisionTask) HasDefaultTask() bool`

HasDefaultTask returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowDecisionTask) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowDecisionTask) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowDecisionTask) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowDecisionTask) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowDecisionTask) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowDecisionTask) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


