# WorkflowDecisionTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | The condition to evaluate for this decision task. The condition can be a workflow or task variable or an expression based on the input parameters. Example value for condition if its Workflow/task variable is -  \&quot;${task1.output.var1} or ${workflow.input.var2}\&quot; which evaluates to a value matching any of the decision case values. Example value for condition if its an expression is - \&quot;if ( $.element ! &#x3D; null &amp;&amp; $.element &gt; 0 ) &#39;true&#39;; else &#39;false&#39;; \&quot; which evaluates to &#39;true&#39; or &#39;false&#39; and will match one of the decision case values. Here \&quot;element\&quot; is a variable in decisiontask&#39;s inputParameters JSON formatted map. You can also use javascript like functions indexOf, toUpperCase in the expression which will be evaluated by the expression evaluator. | [optional] 
**DecisionCases** | Pointer to [**[]WorkflowDecisionCase**](workflow.DecisionCase.md) |  | [optional] 
**DefaultTask** | Pointer to **string** | The default next Task to execute if the decision cannot be evaluated to any of the DecisionCases. | [optional] 
**InputParameters** | Pointer to **interface{}** | JSON formatted map that defines the input given to the decision task. The inputs are used as variables in the condition property of decision task. The input variables can be static values like \&quot;hello\&quot; , \&quot;24\&quot;, \&quot;true\&quot; OR previous task outputs/workflow inputs like \&quot;${task2.output.var1}}\&quot;. The input variables are referrenced as $.inputVariableName in the condition property. | [optional] 

## Methods

### NewWorkflowDecisionTask

`func NewWorkflowDecisionTask() *WorkflowDecisionTask`

NewWorkflowDecisionTask instantiates a new WorkflowDecisionTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDecisionTaskWithDefaults

`func NewWorkflowDecisionTaskWithDefaults() *WorkflowDecisionTask`

NewWorkflowDecisionTaskWithDefaults instantiates a new WorkflowDecisionTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


