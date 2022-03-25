# WorkflowServiceItemActionDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemActionDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemActionDefinition"]
**ActionType** | Pointer to **string** | Type of actionDefinition which decides on how to trigger the action. * &#x60;External&#x60; - External actions definition can be triggered by enduser to perform actions on the service item. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again. * &#x60;Internal&#x60; - Internal action definition is used to trigger periodic actions on the service item instance. * &#x60;Repetitive&#x60; - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created service item. | [optional] [default to "External"]
**AllowedInstanceStates** | Pointer to **[]string** |  | [optional] 
**CoreWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**Description** | Pointer to **string** | The description for this action which provides information on what are the pre-requisites to use this action on the service item and what features are supported by this action. | [optional] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Name** | Pointer to **string** | The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a service item definition. | [optional] 
**OutputParameters** | Pointer to **interface{}** | The output mappings from workflows in the action definition to the service item output definition. Any output from core or post-core workflow can be mapped to service item output definition. The output can be referred using the name of the workflow definition and the output name in the following format &#39;${&lt;ServiceItemActionWorkflowDefinition.Name&gt;.output.&lt;outputName&gt;&#39;. | [optional] 
**Periodicity** | Pointer to **int64** | Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity. | [optional] 
**PostCoreWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**PreCoreWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**StopWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](WorkflowValidationInformation.md) |  | [optional] 
**ValidationWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**ServiceItemDefinition** | Pointer to [**WorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 
**WorkflowDefinition** | Pointer to [**WorkflowWorkflowDefinitionRelationship**](WorkflowWorkflowDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemActionDefinitionAllOf

`func NewWorkflowServiceItemActionDefinitionAllOf(classId string, objectType string, ) *WorkflowServiceItemActionDefinitionAllOf`

NewWorkflowServiceItemActionDefinitionAllOf instantiates a new WorkflowServiceItemActionDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemActionDefinitionAllOfWithDefaults

`func NewWorkflowServiceItemActionDefinitionAllOfWithDefaults() *WorkflowServiceItemActionDefinitionAllOf`

NewWorkflowServiceItemActionDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemActionDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionType

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetAllowedInstanceStates

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetAllowedInstanceStates() []string`

GetAllowedInstanceStates returns the AllowedInstanceStates field if non-nil, zero value otherwise.

### GetAllowedInstanceStatesOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetAllowedInstanceStatesOk() (*[]string, bool)`

GetAllowedInstanceStatesOk returns a tuple with the AllowedInstanceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInstanceStates

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetAllowedInstanceStates(v []string)`

SetAllowedInstanceStates sets AllowedInstanceStates field to given value.

### HasAllowedInstanceStates

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasAllowedInstanceStates() bool`

HasAllowedInstanceStates returns a boolean if a field has been set.

### SetAllowedInstanceStatesNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetAllowedInstanceStatesNil(b bool)`

 SetAllowedInstanceStatesNil sets the value for AllowedInstanceStates to be an explicit nil

### UnsetAllowedInstanceStates
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetAllowedInstanceStates()`

UnsetAllowedInstanceStates ensures that no value is present for AllowedInstanceStates, not even an explicit nil
### GetCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetCoreWorkflows returns the CoreWorkflows field if non-nil, zero value otherwise.

### GetCoreWorkflowsOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetCoreWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetCoreWorkflowsOk returns a tuple with the CoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetCoreWorkflows sets CoreWorkflows field to given value.

### HasCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasCoreWorkflows() bool`

HasCoreWorkflows returns a boolean if a field has been set.

### SetCoreWorkflowsNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetCoreWorkflowsNil(b bool)`

 SetCoreWorkflowsNil sets the value for CoreWorkflows to be an explicit nil

### UnsetCoreWorkflows
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetCoreWorkflows()`

UnsetCoreWorkflows ensures that no value is present for CoreWorkflows, not even an explicit nil
### GetDescription

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetLabel

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputParameters

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetOutputParameters() interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetOutputParametersOk() (*interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetOutputParameters(v interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetPeriodicity

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetPeriodicity() int64`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetPeriodicityOk() (*int64, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetPeriodicity(v int64)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### GetPostCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetPostCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetPostCoreWorkflows returns the PostCoreWorkflows field if non-nil, zero value otherwise.

### GetPostCoreWorkflowsOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetPostCoreWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetPostCoreWorkflowsOk returns a tuple with the PostCoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetPostCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetPostCoreWorkflows sets PostCoreWorkflows field to given value.

### HasPostCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasPostCoreWorkflows() bool`

HasPostCoreWorkflows returns a boolean if a field has been set.

### SetPostCoreWorkflowsNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetPostCoreWorkflowsNil(b bool)`

 SetPostCoreWorkflowsNil sets the value for PostCoreWorkflows to be an explicit nil

### UnsetPostCoreWorkflows
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetPostCoreWorkflows()`

UnsetPostCoreWorkflows ensures that no value is present for PostCoreWorkflows, not even an explicit nil
### GetPreCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetPreCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetPreCoreWorkflows returns the PreCoreWorkflows field if non-nil, zero value otherwise.

### GetPreCoreWorkflowsOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetPreCoreWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetPreCoreWorkflowsOk returns a tuple with the PreCoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetPreCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetPreCoreWorkflows sets PreCoreWorkflows field to given value.

### HasPreCoreWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasPreCoreWorkflows() bool`

HasPreCoreWorkflows returns a boolean if a field has been set.

### SetPreCoreWorkflowsNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetPreCoreWorkflowsNil(b bool)`

 SetPreCoreWorkflowsNil sets the value for PreCoreWorkflows to be an explicit nil

### UnsetPreCoreWorkflows
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetPreCoreWorkflows()`

UnsetPreCoreWorkflows ensures that no value is present for PreCoreWorkflows, not even an explicit nil
### GetStopWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetStopWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetStopWorkflows returns the StopWorkflows field if non-nil, zero value otherwise.

### GetStopWorkflowsOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetStopWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetStopWorkflowsOk returns a tuple with the StopWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetStopWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetStopWorkflows sets StopWorkflows field to given value.

### HasStopWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasStopWorkflows() bool`

HasStopWorkflows returns a boolean if a field has been set.

### SetStopWorkflowsNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetStopWorkflowsNil(b bool)`

 SetStopWorkflowsNil sets the value for StopWorkflows to be an explicit nil

### UnsetStopWorkflows
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetStopWorkflows()`

UnsetStopWorkflows ensures that no value is present for StopWorkflows, not even an explicit nil
### GetValidationInformation

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetValidationWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetValidationWorkflows returns the ValidationWorkflows field if non-nil, zero value otherwise.

### GetValidationWorkflowsOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetValidationWorkflowsOk returns a tuple with the ValidationWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetValidationWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetValidationWorkflows sets ValidationWorkflows field to given value.

### HasValidationWorkflows

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasValidationWorkflows() bool`

HasValidationWorkflows returns a boolean if a field has been set.

### SetValidationWorkflowsNil

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetValidationWorkflowsNil(b bool)`

 SetValidationWorkflowsNil sets the value for ValidationWorkflows to be an explicit nil

### UnsetValidationWorkflows
`func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetValidationWorkflows()`

UnsetValidationWorkflows ensures that no value is present for ValidationWorkflows, not even an explicit nil
### GetServiceItemDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.

### GetWorkflowDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowServiceItemActionDefinitionAllOf) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowServiceItemActionDefinitionAllOf) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


