# WorkflowSolutionActionDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SolutionActionDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SolutionActionDefinition"]
**ActionType** | Pointer to **string** | Type of actionDefinition which decides on how to trigger the action. * &#x60;External&#x60; - External actions definition can be triggered by enduser to perform actions on the solution. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again. * &#x60;Internal&#x60; - Internal action definition is used to trigger periodic actions on the solution instance. * &#x60;Repetitive&#x60; - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created solution. | [optional] [default to "External"]
**AllowedInstanceStates** | Pointer to **[]string** |  | [optional] 
**CoreWorkflows** | Pointer to [**[]WorkflowActionWorkflowDefinition**](WorkflowActionWorkflowDefinition.md) |  | [optional] 
**Description** | Pointer to **string** | The description for this action which provides information on what are the pre-requisites to use this action on the solution and what features are supported by this action. | [optional] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Name** | Pointer to **string** | The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a solution definition. | [optional] 
**OutputParameters** | Pointer to **interface{}** | The output mappings from workflows in the action definition to the solution output definition. Any output from core or post-core workflow can be mapped to solution output definition. The output can be referred using the name of the workflow definition and the output name in the following format &#39;${&lt;ActionWorkflowDefinition.Name&gt;.output.&lt;outputName&gt;&#39;. | [optional] 
**Periodicity** | Pointer to **int64** | Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity. | [optional] 
**PostCoreWorkflows** | Pointer to [**[]WorkflowActionWorkflowDefinition**](WorkflowActionWorkflowDefinition.md) |  | [optional] 
**PreCoreWorkflows** | Pointer to [**[]WorkflowActionWorkflowDefinition**](WorkflowActionWorkflowDefinition.md) |  | [optional] 
**StopWorkflows** | Pointer to [**[]WorkflowActionWorkflowDefinition**](WorkflowActionWorkflowDefinition.md) |  | [optional] 
**UpgradedMoid** | Pointer to **string** | Stores the upgraded Moid for help during future lookups. | [optional] [readonly] 
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](WorkflowValidationInformation.md) |  | [optional] 
**ValidationWorkflows** | Pointer to [**[]WorkflowActionWorkflowDefinition**](WorkflowActionWorkflowDefinition.md) |  | [optional] 
**SolutionDefinition** | Pointer to [**WorkflowSolutionDefinitionRelationship**](WorkflowSolutionDefinitionRelationship.md) |  | [optional] 
**WorkflowDefinition** | Pointer to [**WorkflowWorkflowDefinitionRelationship**](WorkflowWorkflowDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowSolutionActionDefinitionAllOf

`func NewWorkflowSolutionActionDefinitionAllOf(classId string, objectType string, ) *WorkflowSolutionActionDefinitionAllOf`

NewWorkflowSolutionActionDefinitionAllOf instantiates a new WorkflowSolutionActionDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionActionDefinitionAllOfWithDefaults

`func NewWorkflowSolutionActionDefinitionAllOfWithDefaults() *WorkflowSolutionActionDefinitionAllOf`

NewWorkflowSolutionActionDefinitionAllOfWithDefaults instantiates a new WorkflowSolutionActionDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSolutionActionDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSolutionActionDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSolutionActionDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSolutionActionDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionType

`func (o *WorkflowSolutionActionDefinitionAllOf) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WorkflowSolutionActionDefinitionAllOf) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WorkflowSolutionActionDefinitionAllOf) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetAllowedInstanceStates

`func (o *WorkflowSolutionActionDefinitionAllOf) GetAllowedInstanceStates() []string`

GetAllowedInstanceStates returns the AllowedInstanceStates field if non-nil, zero value otherwise.

### GetAllowedInstanceStatesOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetAllowedInstanceStatesOk() (*[]string, bool)`

GetAllowedInstanceStatesOk returns a tuple with the AllowedInstanceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInstanceStates

`func (o *WorkflowSolutionActionDefinitionAllOf) SetAllowedInstanceStates(v []string)`

SetAllowedInstanceStates sets AllowedInstanceStates field to given value.

### HasAllowedInstanceStates

`func (o *WorkflowSolutionActionDefinitionAllOf) HasAllowedInstanceStates() bool`

HasAllowedInstanceStates returns a boolean if a field has been set.

### SetAllowedInstanceStatesNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetAllowedInstanceStatesNil(b bool)`

 SetAllowedInstanceStatesNil sets the value for AllowedInstanceStates to be an explicit nil

### UnsetAllowedInstanceStates
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetAllowedInstanceStates()`

UnsetAllowedInstanceStates ensures that no value is present for AllowedInstanceStates, not even an explicit nil
### GetCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) GetCoreWorkflows() []WorkflowActionWorkflowDefinition`

GetCoreWorkflows returns the CoreWorkflows field if non-nil, zero value otherwise.

### GetCoreWorkflowsOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetCoreWorkflowsOk() (*[]WorkflowActionWorkflowDefinition, bool)`

GetCoreWorkflowsOk returns a tuple with the CoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) SetCoreWorkflows(v []WorkflowActionWorkflowDefinition)`

SetCoreWorkflows sets CoreWorkflows field to given value.

### HasCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) HasCoreWorkflows() bool`

HasCoreWorkflows returns a boolean if a field has been set.

### SetCoreWorkflowsNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetCoreWorkflowsNil(b bool)`

 SetCoreWorkflowsNil sets the value for CoreWorkflows to be an explicit nil

### UnsetCoreWorkflows
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetCoreWorkflows()`

UnsetCoreWorkflows ensures that no value is present for CoreWorkflows, not even an explicit nil
### GetDescription

`func (o *WorkflowSolutionActionDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSolutionActionDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSolutionActionDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetLabel

`func (o *WorkflowSolutionActionDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowSolutionActionDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowSolutionActionDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSolutionActionDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSolutionActionDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSolutionActionDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputParameters

`func (o *WorkflowSolutionActionDefinitionAllOf) GetOutputParameters() interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetOutputParametersOk() (*interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *WorkflowSolutionActionDefinitionAllOf) SetOutputParameters(v interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *WorkflowSolutionActionDefinitionAllOf) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetPeriodicity

`func (o *WorkflowSolutionActionDefinitionAllOf) GetPeriodicity() int64`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetPeriodicityOk() (*int64, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *WorkflowSolutionActionDefinitionAllOf) SetPeriodicity(v int64)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *WorkflowSolutionActionDefinitionAllOf) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### GetPostCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) GetPostCoreWorkflows() []WorkflowActionWorkflowDefinition`

GetPostCoreWorkflows returns the PostCoreWorkflows field if non-nil, zero value otherwise.

### GetPostCoreWorkflowsOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetPostCoreWorkflowsOk() (*[]WorkflowActionWorkflowDefinition, bool)`

GetPostCoreWorkflowsOk returns a tuple with the PostCoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) SetPostCoreWorkflows(v []WorkflowActionWorkflowDefinition)`

SetPostCoreWorkflows sets PostCoreWorkflows field to given value.

### HasPostCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) HasPostCoreWorkflows() bool`

HasPostCoreWorkflows returns a boolean if a field has been set.

### SetPostCoreWorkflowsNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetPostCoreWorkflowsNil(b bool)`

 SetPostCoreWorkflowsNil sets the value for PostCoreWorkflows to be an explicit nil

### UnsetPostCoreWorkflows
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetPostCoreWorkflows()`

UnsetPostCoreWorkflows ensures that no value is present for PostCoreWorkflows, not even an explicit nil
### GetPreCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) GetPreCoreWorkflows() []WorkflowActionWorkflowDefinition`

GetPreCoreWorkflows returns the PreCoreWorkflows field if non-nil, zero value otherwise.

### GetPreCoreWorkflowsOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetPreCoreWorkflowsOk() (*[]WorkflowActionWorkflowDefinition, bool)`

GetPreCoreWorkflowsOk returns a tuple with the PreCoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) SetPreCoreWorkflows(v []WorkflowActionWorkflowDefinition)`

SetPreCoreWorkflows sets PreCoreWorkflows field to given value.

### HasPreCoreWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) HasPreCoreWorkflows() bool`

HasPreCoreWorkflows returns a boolean if a field has been set.

### SetPreCoreWorkflowsNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetPreCoreWorkflowsNil(b bool)`

 SetPreCoreWorkflowsNil sets the value for PreCoreWorkflows to be an explicit nil

### UnsetPreCoreWorkflows
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetPreCoreWorkflows()`

UnsetPreCoreWorkflows ensures that no value is present for PreCoreWorkflows, not even an explicit nil
### GetStopWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) GetStopWorkflows() []WorkflowActionWorkflowDefinition`

GetStopWorkflows returns the StopWorkflows field if non-nil, zero value otherwise.

### GetStopWorkflowsOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetStopWorkflowsOk() (*[]WorkflowActionWorkflowDefinition, bool)`

GetStopWorkflowsOk returns a tuple with the StopWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) SetStopWorkflows(v []WorkflowActionWorkflowDefinition)`

SetStopWorkflows sets StopWorkflows field to given value.

### HasStopWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) HasStopWorkflows() bool`

HasStopWorkflows returns a boolean if a field has been set.

### SetStopWorkflowsNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetStopWorkflowsNil(b bool)`

 SetStopWorkflowsNil sets the value for StopWorkflows to be an explicit nil

### UnsetStopWorkflows
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetStopWorkflows()`

UnsetStopWorkflows ensures that no value is present for StopWorkflows, not even an explicit nil
### GetUpgradedMoid

`func (o *WorkflowSolutionActionDefinitionAllOf) GetUpgradedMoid() string`

GetUpgradedMoid returns the UpgradedMoid field if non-nil, zero value otherwise.

### GetUpgradedMoidOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetUpgradedMoidOk() (*string, bool)`

GetUpgradedMoidOk returns a tuple with the UpgradedMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedMoid

`func (o *WorkflowSolutionActionDefinitionAllOf) SetUpgradedMoid(v string)`

SetUpgradedMoid sets UpgradedMoid field to given value.

### HasUpgradedMoid

`func (o *WorkflowSolutionActionDefinitionAllOf) HasUpgradedMoid() bool`

HasUpgradedMoid returns a boolean if a field has been set.

### GetValidationInformation

`func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkflowSolutionActionDefinitionAllOf) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkflowSolutionActionDefinitionAllOf) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetValidationWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationWorkflows() []WorkflowActionWorkflowDefinition`

GetValidationWorkflows returns the ValidationWorkflows field if non-nil, zero value otherwise.

### GetValidationWorkflowsOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationWorkflowsOk() (*[]WorkflowActionWorkflowDefinition, bool)`

GetValidationWorkflowsOk returns a tuple with the ValidationWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) SetValidationWorkflows(v []WorkflowActionWorkflowDefinition)`

SetValidationWorkflows sets ValidationWorkflows field to given value.

### HasValidationWorkflows

`func (o *WorkflowSolutionActionDefinitionAllOf) HasValidationWorkflows() bool`

HasValidationWorkflows returns a boolean if a field has been set.

### SetValidationWorkflowsNil

`func (o *WorkflowSolutionActionDefinitionAllOf) SetValidationWorkflowsNil(b bool)`

 SetValidationWorkflowsNil sets the value for ValidationWorkflows to be an explicit nil

### UnsetValidationWorkflows
`func (o *WorkflowSolutionActionDefinitionAllOf) UnsetValidationWorkflows()`

UnsetValidationWorkflows ensures that no value is present for ValidationWorkflows, not even an explicit nil
### GetSolutionDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) GetSolutionDefinition() WorkflowSolutionDefinitionRelationship`

GetSolutionDefinition returns the SolutionDefinition field if non-nil, zero value otherwise.

### GetSolutionDefinitionOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetSolutionDefinitionOk() (*WorkflowSolutionDefinitionRelationship, bool)`

GetSolutionDefinitionOk returns a tuple with the SolutionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) SetSolutionDefinition(v WorkflowSolutionDefinitionRelationship)`

SetSolutionDefinition sets SolutionDefinition field to given value.

### HasSolutionDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) HasSolutionDefinition() bool`

HasSolutionDefinition returns a boolean if a field has been set.

### GetWorkflowDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowSolutionActionDefinitionAllOf) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowSolutionActionDefinitionAllOf) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


