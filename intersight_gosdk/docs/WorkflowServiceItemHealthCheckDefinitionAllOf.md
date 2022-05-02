# WorkflowServiceItemHealthCheckDefinitionAllOf

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

### NewWorkflowServiceItemHealthCheckDefinitionAllOf

`func NewWorkflowServiceItemHealthCheckDefinitionAllOf(classId string, objectType string, ) *WorkflowServiceItemHealthCheckDefinitionAllOf`

NewWorkflowServiceItemHealthCheckDefinitionAllOf instantiates a new WorkflowServiceItemHealthCheckDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemHealthCheckDefinitionAllOfWithDefaults

`func NewWorkflowServiceItemHealthCheckDefinitionAllOfWithDefaults() *WorkflowServiceItemHealthCheckDefinitionAllOf`

NewWorkflowServiceItemHealthCheckDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemHealthCheckDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCommonCauseAndResolution

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCommonCauseAndResolution() string`

GetCommonCauseAndResolution returns the CommonCauseAndResolution field if non-nil, zero value otherwise.

### GetCommonCauseAndResolutionOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCommonCauseAndResolutionOk() (*string, bool)`

GetCommonCauseAndResolutionOk returns a tuple with the CommonCauseAndResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonCauseAndResolution

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetCommonCauseAndResolution(v string)`

SetCommonCauseAndResolution sets CommonCauseAndResolution field to given value.

### HasCommonCauseAndResolution

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasCommonCauseAndResolution() bool`

HasCommonCauseAndResolution returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetHealthCheckWorkflow

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetHealthCheckWorkflow() WorkflowServiceItemActionWorkflowDefinition`

GetHealthCheckWorkflow returns the HealthCheckWorkflow field if non-nil, zero value otherwise.

### GetHealthCheckWorkflowOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetHealthCheckWorkflowOk() (*WorkflowServiceItemActionWorkflowDefinition, bool)`

GetHealthCheckWorkflowOk returns a tuple with the HealthCheckWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckWorkflow

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetHealthCheckWorkflow(v WorkflowServiceItemActionWorkflowDefinition)`

SetHealthCheckWorkflow sets HealthCheckWorkflow field to given value.

### HasHealthCheckWorkflow

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasHealthCheckWorkflow() bool`

HasHealthCheckWorkflow returns a boolean if a field has been set.

### SetHealthCheckWorkflowNil

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetHealthCheckWorkflowNil(b bool)`

 SetHealthCheckWorkflowNil sets the value for HealthCheckWorkflow to be an explicit nil

### UnsetHealthCheckWorkflow
`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) UnsetHealthCheckWorkflow()`

UnsetHealthCheckWorkflow ensures that no value is present for HealthCheckWorkflow, not even an explicit nil
### GetLabel

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceItemDefinition

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


