# WorkflowServiceItemHealthCheckDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemHealthCheckDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemHealthCheckDefinition"]
**Category** | Pointer to **string** | Category that the health check belongs to. | [optional] 
**CommonCauseAndResolution** | Pointer to **string** | Static information detailing the common cause for the health check failure. It also gives possible remediation actions that can be taken to remedy the health check failure. | [optional] 
**Description** | Pointer to **string** | Description of the health check definition. | [optional] 
**ExecutionMode** | Pointer to **string** | Execution mode of the health check on service item. * &#x60;OnDemand&#x60; - Execute the health check on-demand. * &#x60;Periodic&#x60; - Execute the health check on a periodic basis. | [optional] [default to "OnDemand"]
**HealthCheckWorkflow** | Pointer to [**NullableWorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**Label** | Pointer to **string** | Label for the health check definition that is displayed on UI. | [optional] 
**Name** | Pointer to **string** | Name of the health check definition. | [optional] 
**ServiceItemDefinition** | Pointer to [**WorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemHealthCheckDefinition

`func NewWorkflowServiceItemHealthCheckDefinition(classId string, objectType string, ) *WorkflowServiceItemHealthCheckDefinition`

NewWorkflowServiceItemHealthCheckDefinition instantiates a new WorkflowServiceItemHealthCheckDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemHealthCheckDefinitionWithDefaults

`func NewWorkflowServiceItemHealthCheckDefinitionWithDefaults() *WorkflowServiceItemHealthCheckDefinition`

NewWorkflowServiceItemHealthCheckDefinitionWithDefaults instantiates a new WorkflowServiceItemHealthCheckDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemHealthCheckDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemHealthCheckDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemHealthCheckDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemHealthCheckDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *WorkflowServiceItemHealthCheckDefinition) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WorkflowServiceItemHealthCheckDefinition) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WorkflowServiceItemHealthCheckDefinition) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCommonCauseAndResolution

`func (o *WorkflowServiceItemHealthCheckDefinition) GetCommonCauseAndResolution() string`

GetCommonCauseAndResolution returns the CommonCauseAndResolution field if non-nil, zero value otherwise.

### GetCommonCauseAndResolutionOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetCommonCauseAndResolutionOk() (*string, bool)`

GetCommonCauseAndResolutionOk returns a tuple with the CommonCauseAndResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonCauseAndResolution

`func (o *WorkflowServiceItemHealthCheckDefinition) SetCommonCauseAndResolution(v string)`

SetCommonCauseAndResolution sets CommonCauseAndResolution field to given value.

### HasCommonCauseAndResolution

`func (o *WorkflowServiceItemHealthCheckDefinition) HasCommonCauseAndResolution() bool`

HasCommonCauseAndResolution returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowServiceItemHealthCheckDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemHealthCheckDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemHealthCheckDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *WorkflowServiceItemHealthCheckDefinition) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *WorkflowServiceItemHealthCheckDefinition) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *WorkflowServiceItemHealthCheckDefinition) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetHealthCheckWorkflow

`func (o *WorkflowServiceItemHealthCheckDefinition) GetHealthCheckWorkflow() WorkflowServiceItemActionWorkflowDefinition`

GetHealthCheckWorkflow returns the HealthCheckWorkflow field if non-nil, zero value otherwise.

### GetHealthCheckWorkflowOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetHealthCheckWorkflowOk() (*WorkflowServiceItemActionWorkflowDefinition, bool)`

GetHealthCheckWorkflowOk returns a tuple with the HealthCheckWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckWorkflow

`func (o *WorkflowServiceItemHealthCheckDefinition) SetHealthCheckWorkflow(v WorkflowServiceItemActionWorkflowDefinition)`

SetHealthCheckWorkflow sets HealthCheckWorkflow field to given value.

### HasHealthCheckWorkflow

`func (o *WorkflowServiceItemHealthCheckDefinition) HasHealthCheckWorkflow() bool`

HasHealthCheckWorkflow returns a boolean if a field has been set.

### SetHealthCheckWorkflowNil

`func (o *WorkflowServiceItemHealthCheckDefinition) SetHealthCheckWorkflowNil(b bool)`

 SetHealthCheckWorkflowNil sets the value for HealthCheckWorkflow to be an explicit nil

### UnsetHealthCheckWorkflow
`func (o *WorkflowServiceItemHealthCheckDefinition) UnsetHealthCheckWorkflow()`

UnsetHealthCheckWorkflow ensures that no value is present for HealthCheckWorkflow, not even an explicit nil
### GetLabel

`func (o *WorkflowServiceItemHealthCheckDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemHealthCheckDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemHealthCheckDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemHealthCheckDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemHealthCheckDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemHealthCheckDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceItemDefinition

`func (o *WorkflowServiceItemHealthCheckDefinition) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemHealthCheckDefinition) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemHealthCheckDefinition) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemHealthCheckDefinition) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


