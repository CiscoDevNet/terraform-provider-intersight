# WorkflowServiceItemInputDefinitionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemInputDefinitionType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemInputDefinitionType"]
**ActionName** | Pointer to **string** | The name of the service item action. | [optional] [readonly] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**InputParameters** | Pointer to **interface{}** | Capture the mapping of service item ActionDefinition inputDefinition to catalog service request. | [optional] 
**OperationType** | Pointer to **string** | Type of action operation to be executed on the service item. * &#x60;PostDeployment&#x60; - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * &#x60;Deployment&#x60; - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * &#x60;Decommission&#x60; - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item. * &#x60;Migration&#x60; - This represents the migration action, used to migrate service item instance from one service item definition version to another service item definition version. All valid service items can have up to one operation type marked as Migration. Once a migration action is running on a service item instance, no further operations are allowed on that service item instance during the migration process. | [optional] [readonly] [default to "PostDeployment"]
**ServiceItemActionDefinition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemInputDefinitionType

`func NewWorkflowServiceItemInputDefinitionType(classId string, objectType string, ) *WorkflowServiceItemInputDefinitionType`

NewWorkflowServiceItemInputDefinitionType instantiates a new WorkflowServiceItemInputDefinitionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemInputDefinitionTypeWithDefaults

`func NewWorkflowServiceItemInputDefinitionTypeWithDefaults() *WorkflowServiceItemInputDefinitionType`

NewWorkflowServiceItemInputDefinitionTypeWithDefaults instantiates a new WorkflowServiceItemInputDefinitionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemInputDefinitionType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemInputDefinitionType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemInputDefinitionType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemInputDefinitionType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemInputDefinitionType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemInputDefinitionType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionName

`func (o *WorkflowServiceItemInputDefinitionType) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *WorkflowServiceItemInputDefinitionType) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *WorkflowServiceItemInputDefinitionType) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *WorkflowServiceItemInputDefinitionType) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowServiceItemInputDefinitionType) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowServiceItemInputDefinitionType) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowServiceItemInputDefinitionType) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowServiceItemInputDefinitionType) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowServiceItemInputDefinitionType) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowServiceItemInputDefinitionType) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetInputParameters

`func (o *WorkflowServiceItemInputDefinitionType) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowServiceItemInputDefinitionType) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowServiceItemInputDefinitionType) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowServiceItemInputDefinitionType) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowServiceItemInputDefinitionType) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowServiceItemInputDefinitionType) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetOperationType

`func (o *WorkflowServiceItemInputDefinitionType) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *WorkflowServiceItemInputDefinitionType) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *WorkflowServiceItemInputDefinitionType) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *WorkflowServiceItemInputDefinitionType) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetServiceItemActionDefinition

`func (o *WorkflowServiceItemInputDefinitionType) GetServiceItemActionDefinition() MoMoRef`

GetServiceItemActionDefinition returns the ServiceItemActionDefinition field if non-nil, zero value otherwise.

### GetServiceItemActionDefinitionOk

`func (o *WorkflowServiceItemInputDefinitionType) GetServiceItemActionDefinitionOk() (*MoMoRef, bool)`

GetServiceItemActionDefinitionOk returns a tuple with the ServiceItemActionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemActionDefinition

`func (o *WorkflowServiceItemInputDefinitionType) SetServiceItemActionDefinition(v MoMoRef)`

SetServiceItemActionDefinition sets ServiceItemActionDefinition field to given value.

### HasServiceItemActionDefinition

`func (o *WorkflowServiceItemInputDefinitionType) HasServiceItemActionDefinition() bool`

HasServiceItemActionDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


