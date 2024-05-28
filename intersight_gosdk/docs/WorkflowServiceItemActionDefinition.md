# WorkflowServiceItemActionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemActionDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemActionDefinition"]
**ActionProperties** | Pointer to [**NullableWorkflowServiceItemActionProperties**](WorkflowServiceItemActionProperties.md) |  | [optional] 
**ActionType** | Pointer to **string** | Type of actionDefinition which decides on how to trigger the action. * &#x60;External&#x60; - External actions definition can be triggered by enduser to perform actions on the service item. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again. * &#x60;Internal&#x60; - Internal action definition is used to trigger periodic actions on the service item instance. * &#x60;Repetitive&#x60; - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created service item. | [optional] [default to "External"]
**AllowedInstanceStates** | Pointer to **[]string** |  | [optional] 
**AttributeParameters** | Pointer to **interface{}** | The mappings from workflows in the action definition to the service item attribute definition. Any output from core or post-core workflow can be mapped to service item attribute definition. The attribute can be referred using the name of the workflow definition and the attribute name in the following format &#39;${&lt;ServiceItemActionWorkflowDefinition.Name&gt;.output.&lt;outputName&gt;&#39;. | [optional] 
**CoreWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**Description** | Pointer to **string** | The description for this action which provides information on what are the pre-requisites to use this action on the service item and what features are supported by this action. | [optional] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Name** | Pointer to **string** | The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a service item definition. | [optional] 
**Periodicity** | Pointer to **int64** | Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity. | [optional] 
**PostCoreWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**PreCoreWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**RestrictOnPrivateAppliance** | Pointer to **bool** | The flag to indicate that action is restricted on a Private Virtual Appliance. | [optional] [default to false]
**StopWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**UserIdOrEmail** | Pointer to **string** | The user identifier who created or updated the service item action definition. | [optional] [readonly] 
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](WorkflowValidationInformation.md) |  | [optional] 
**ValidationWorkflows** | Pointer to [**[]WorkflowServiceItemActionWorkflowDefinition**](WorkflowServiceItemActionWorkflowDefinition.md) |  | [optional] 
**AssociatedRoles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] 
**ServiceItemDefinition** | Pointer to [**NullableWorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 
**WorkflowDefinition** | Pointer to [**NullableWorkflowWorkflowDefinitionRelationship**](WorkflowWorkflowDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemActionDefinition

`func NewWorkflowServiceItemActionDefinition(classId string, objectType string, ) *WorkflowServiceItemActionDefinition`

NewWorkflowServiceItemActionDefinition instantiates a new WorkflowServiceItemActionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemActionDefinitionWithDefaults

`func NewWorkflowServiceItemActionDefinitionWithDefaults() *WorkflowServiceItemActionDefinition`

NewWorkflowServiceItemActionDefinitionWithDefaults instantiates a new WorkflowServiceItemActionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemActionDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemActionDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemActionDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemActionDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemActionDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemActionDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionProperties

`func (o *WorkflowServiceItemActionDefinition) GetActionProperties() WorkflowServiceItemActionProperties`

GetActionProperties returns the ActionProperties field if non-nil, zero value otherwise.

### GetActionPropertiesOk

`func (o *WorkflowServiceItemActionDefinition) GetActionPropertiesOk() (*WorkflowServiceItemActionProperties, bool)`

GetActionPropertiesOk returns a tuple with the ActionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionProperties

`func (o *WorkflowServiceItemActionDefinition) SetActionProperties(v WorkflowServiceItemActionProperties)`

SetActionProperties sets ActionProperties field to given value.

### HasActionProperties

`func (o *WorkflowServiceItemActionDefinition) HasActionProperties() bool`

HasActionProperties returns a boolean if a field has been set.

### SetActionPropertiesNil

`func (o *WorkflowServiceItemActionDefinition) SetActionPropertiesNil(b bool)`

 SetActionPropertiesNil sets the value for ActionProperties to be an explicit nil

### UnsetActionProperties
`func (o *WorkflowServiceItemActionDefinition) UnsetActionProperties()`

UnsetActionProperties ensures that no value is present for ActionProperties, not even an explicit nil
### GetActionType

`func (o *WorkflowServiceItemActionDefinition) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WorkflowServiceItemActionDefinition) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WorkflowServiceItemActionDefinition) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WorkflowServiceItemActionDefinition) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetAllowedInstanceStates

`func (o *WorkflowServiceItemActionDefinition) GetAllowedInstanceStates() []string`

GetAllowedInstanceStates returns the AllowedInstanceStates field if non-nil, zero value otherwise.

### GetAllowedInstanceStatesOk

`func (o *WorkflowServiceItemActionDefinition) GetAllowedInstanceStatesOk() (*[]string, bool)`

GetAllowedInstanceStatesOk returns a tuple with the AllowedInstanceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInstanceStates

`func (o *WorkflowServiceItemActionDefinition) SetAllowedInstanceStates(v []string)`

SetAllowedInstanceStates sets AllowedInstanceStates field to given value.

### HasAllowedInstanceStates

`func (o *WorkflowServiceItemActionDefinition) HasAllowedInstanceStates() bool`

HasAllowedInstanceStates returns a boolean if a field has been set.

### SetAllowedInstanceStatesNil

`func (o *WorkflowServiceItemActionDefinition) SetAllowedInstanceStatesNil(b bool)`

 SetAllowedInstanceStatesNil sets the value for AllowedInstanceStates to be an explicit nil

### UnsetAllowedInstanceStates
`func (o *WorkflowServiceItemActionDefinition) UnsetAllowedInstanceStates()`

UnsetAllowedInstanceStates ensures that no value is present for AllowedInstanceStates, not even an explicit nil
### GetAttributeParameters

`func (o *WorkflowServiceItemActionDefinition) GetAttributeParameters() interface{}`

GetAttributeParameters returns the AttributeParameters field if non-nil, zero value otherwise.

### GetAttributeParametersOk

`func (o *WorkflowServiceItemActionDefinition) GetAttributeParametersOk() (*interface{}, bool)`

GetAttributeParametersOk returns a tuple with the AttributeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeParameters

`func (o *WorkflowServiceItemActionDefinition) SetAttributeParameters(v interface{})`

SetAttributeParameters sets AttributeParameters field to given value.

### HasAttributeParameters

`func (o *WorkflowServiceItemActionDefinition) HasAttributeParameters() bool`

HasAttributeParameters returns a boolean if a field has been set.

### SetAttributeParametersNil

`func (o *WorkflowServiceItemActionDefinition) SetAttributeParametersNil(b bool)`

 SetAttributeParametersNil sets the value for AttributeParameters to be an explicit nil

### UnsetAttributeParameters
`func (o *WorkflowServiceItemActionDefinition) UnsetAttributeParameters()`

UnsetAttributeParameters ensures that no value is present for AttributeParameters, not even an explicit nil
### GetCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) GetCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetCoreWorkflows returns the CoreWorkflows field if non-nil, zero value otherwise.

### GetCoreWorkflowsOk

`func (o *WorkflowServiceItemActionDefinition) GetCoreWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetCoreWorkflowsOk returns a tuple with the CoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) SetCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetCoreWorkflows sets CoreWorkflows field to given value.

### HasCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) HasCoreWorkflows() bool`

HasCoreWorkflows returns a boolean if a field has been set.

### SetCoreWorkflowsNil

`func (o *WorkflowServiceItemActionDefinition) SetCoreWorkflowsNil(b bool)`

 SetCoreWorkflowsNil sets the value for CoreWorkflows to be an explicit nil

### UnsetCoreWorkflows
`func (o *WorkflowServiceItemActionDefinition) UnsetCoreWorkflows()`

UnsetCoreWorkflows ensures that no value is present for CoreWorkflows, not even an explicit nil
### GetDescription

`func (o *WorkflowServiceItemActionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemActionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemActionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemActionDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowServiceItemActionDefinition) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowServiceItemActionDefinition) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowServiceItemActionDefinition) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowServiceItemActionDefinition) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowServiceItemActionDefinition) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowServiceItemActionDefinition) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetLabel

`func (o *WorkflowServiceItemActionDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemActionDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemActionDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemActionDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemActionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemActionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemActionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemActionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeriodicity

`func (o *WorkflowServiceItemActionDefinition) GetPeriodicity() int64`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *WorkflowServiceItemActionDefinition) GetPeriodicityOk() (*int64, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *WorkflowServiceItemActionDefinition) SetPeriodicity(v int64)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *WorkflowServiceItemActionDefinition) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### GetPostCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) GetPostCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetPostCoreWorkflows returns the PostCoreWorkflows field if non-nil, zero value otherwise.

### GetPostCoreWorkflowsOk

`func (o *WorkflowServiceItemActionDefinition) GetPostCoreWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetPostCoreWorkflowsOk returns a tuple with the PostCoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) SetPostCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetPostCoreWorkflows sets PostCoreWorkflows field to given value.

### HasPostCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) HasPostCoreWorkflows() bool`

HasPostCoreWorkflows returns a boolean if a field has been set.

### SetPostCoreWorkflowsNil

`func (o *WorkflowServiceItemActionDefinition) SetPostCoreWorkflowsNil(b bool)`

 SetPostCoreWorkflowsNil sets the value for PostCoreWorkflows to be an explicit nil

### UnsetPostCoreWorkflows
`func (o *WorkflowServiceItemActionDefinition) UnsetPostCoreWorkflows()`

UnsetPostCoreWorkflows ensures that no value is present for PostCoreWorkflows, not even an explicit nil
### GetPreCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) GetPreCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetPreCoreWorkflows returns the PreCoreWorkflows field if non-nil, zero value otherwise.

### GetPreCoreWorkflowsOk

`func (o *WorkflowServiceItemActionDefinition) GetPreCoreWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetPreCoreWorkflowsOk returns a tuple with the PreCoreWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) SetPreCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetPreCoreWorkflows sets PreCoreWorkflows field to given value.

### HasPreCoreWorkflows

`func (o *WorkflowServiceItemActionDefinition) HasPreCoreWorkflows() bool`

HasPreCoreWorkflows returns a boolean if a field has been set.

### SetPreCoreWorkflowsNil

`func (o *WorkflowServiceItemActionDefinition) SetPreCoreWorkflowsNil(b bool)`

 SetPreCoreWorkflowsNil sets the value for PreCoreWorkflows to be an explicit nil

### UnsetPreCoreWorkflows
`func (o *WorkflowServiceItemActionDefinition) UnsetPreCoreWorkflows()`

UnsetPreCoreWorkflows ensures that no value is present for PreCoreWorkflows, not even an explicit nil
### GetRestrictOnPrivateAppliance

`func (o *WorkflowServiceItemActionDefinition) GetRestrictOnPrivateAppliance() bool`

GetRestrictOnPrivateAppliance returns the RestrictOnPrivateAppliance field if non-nil, zero value otherwise.

### GetRestrictOnPrivateApplianceOk

`func (o *WorkflowServiceItemActionDefinition) GetRestrictOnPrivateApplianceOk() (*bool, bool)`

GetRestrictOnPrivateApplianceOk returns a tuple with the RestrictOnPrivateAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictOnPrivateAppliance

`func (o *WorkflowServiceItemActionDefinition) SetRestrictOnPrivateAppliance(v bool)`

SetRestrictOnPrivateAppliance sets RestrictOnPrivateAppliance field to given value.

### HasRestrictOnPrivateAppliance

`func (o *WorkflowServiceItemActionDefinition) HasRestrictOnPrivateAppliance() bool`

HasRestrictOnPrivateAppliance returns a boolean if a field has been set.

### GetStopWorkflows

`func (o *WorkflowServiceItemActionDefinition) GetStopWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetStopWorkflows returns the StopWorkflows field if non-nil, zero value otherwise.

### GetStopWorkflowsOk

`func (o *WorkflowServiceItemActionDefinition) GetStopWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetStopWorkflowsOk returns a tuple with the StopWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWorkflows

`func (o *WorkflowServiceItemActionDefinition) SetStopWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetStopWorkflows sets StopWorkflows field to given value.

### HasStopWorkflows

`func (o *WorkflowServiceItemActionDefinition) HasStopWorkflows() bool`

HasStopWorkflows returns a boolean if a field has been set.

### SetStopWorkflowsNil

`func (o *WorkflowServiceItemActionDefinition) SetStopWorkflowsNil(b bool)`

 SetStopWorkflowsNil sets the value for StopWorkflows to be an explicit nil

### UnsetStopWorkflows
`func (o *WorkflowServiceItemActionDefinition) UnsetStopWorkflows()`

UnsetStopWorkflows ensures that no value is present for StopWorkflows, not even an explicit nil
### GetUserIdOrEmail

`func (o *WorkflowServiceItemActionDefinition) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *WorkflowServiceItemActionDefinition) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *WorkflowServiceItemActionDefinition) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *WorkflowServiceItemActionDefinition) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetValidationInformation

`func (o *WorkflowServiceItemActionDefinition) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkflowServiceItemActionDefinition) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkflowServiceItemActionDefinition) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkflowServiceItemActionDefinition) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkflowServiceItemActionDefinition) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkflowServiceItemActionDefinition) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetValidationWorkflows

`func (o *WorkflowServiceItemActionDefinition) GetValidationWorkflows() []WorkflowServiceItemActionWorkflowDefinition`

GetValidationWorkflows returns the ValidationWorkflows field if non-nil, zero value otherwise.

### GetValidationWorkflowsOk

`func (o *WorkflowServiceItemActionDefinition) GetValidationWorkflowsOk() (*[]WorkflowServiceItemActionWorkflowDefinition, bool)`

GetValidationWorkflowsOk returns a tuple with the ValidationWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationWorkflows

`func (o *WorkflowServiceItemActionDefinition) SetValidationWorkflows(v []WorkflowServiceItemActionWorkflowDefinition)`

SetValidationWorkflows sets ValidationWorkflows field to given value.

### HasValidationWorkflows

`func (o *WorkflowServiceItemActionDefinition) HasValidationWorkflows() bool`

HasValidationWorkflows returns a boolean if a field has been set.

### SetValidationWorkflowsNil

`func (o *WorkflowServiceItemActionDefinition) SetValidationWorkflowsNil(b bool)`

 SetValidationWorkflowsNil sets the value for ValidationWorkflows to be an explicit nil

### UnsetValidationWorkflows
`func (o *WorkflowServiceItemActionDefinition) UnsetValidationWorkflows()`

UnsetValidationWorkflows ensures that no value is present for ValidationWorkflows, not even an explicit nil
### GetAssociatedRoles

`func (o *WorkflowServiceItemActionDefinition) GetAssociatedRoles() []IamRoleRelationship`

GetAssociatedRoles returns the AssociatedRoles field if non-nil, zero value otherwise.

### GetAssociatedRolesOk

`func (o *WorkflowServiceItemActionDefinition) GetAssociatedRolesOk() (*[]IamRoleRelationship, bool)`

GetAssociatedRolesOk returns a tuple with the AssociatedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRoles

`func (o *WorkflowServiceItemActionDefinition) SetAssociatedRoles(v []IamRoleRelationship)`

SetAssociatedRoles sets AssociatedRoles field to given value.

### HasAssociatedRoles

`func (o *WorkflowServiceItemActionDefinition) HasAssociatedRoles() bool`

HasAssociatedRoles returns a boolean if a field has been set.

### SetAssociatedRolesNil

`func (o *WorkflowServiceItemActionDefinition) SetAssociatedRolesNil(b bool)`

 SetAssociatedRolesNil sets the value for AssociatedRoles to be an explicit nil

### UnsetAssociatedRoles
`func (o *WorkflowServiceItemActionDefinition) UnsetAssociatedRoles()`

UnsetAssociatedRoles ensures that no value is present for AssociatedRoles, not even an explicit nil
### GetServiceItemDefinition

`func (o *WorkflowServiceItemActionDefinition) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemActionDefinition) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemActionDefinition) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemActionDefinition) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.

### SetServiceItemDefinitionNil

`func (o *WorkflowServiceItemActionDefinition) SetServiceItemDefinitionNil(b bool)`

 SetServiceItemDefinitionNil sets the value for ServiceItemDefinition to be an explicit nil

### UnsetServiceItemDefinition
`func (o *WorkflowServiceItemActionDefinition) UnsetServiceItemDefinition()`

UnsetServiceItemDefinition ensures that no value is present for ServiceItemDefinition, not even an explicit nil
### GetWorkflowDefinition

`func (o *WorkflowServiceItemActionDefinition) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowServiceItemActionDefinition) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowServiceItemActionDefinition) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowServiceItemActionDefinition) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.

### SetWorkflowDefinitionNil

`func (o *WorkflowServiceItemActionDefinition) SetWorkflowDefinitionNil(b bool)`

 SetWorkflowDefinitionNil sets the value for WorkflowDefinition to be an explicit nil

### UnsetWorkflowDefinition
`func (o *WorkflowServiceItemActionDefinition) UnsetWorkflowDefinition()`

UnsetWorkflowDefinition ensures that no value is present for WorkflowDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


