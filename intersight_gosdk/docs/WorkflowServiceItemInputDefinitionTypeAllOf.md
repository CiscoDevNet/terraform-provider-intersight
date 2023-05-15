# WorkflowServiceItemInputDefinitionTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemInputDefinitionType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemInputDefinitionType"]
**ActionName** | Pointer to **string** | The name of the service item action. | [optional] [readonly] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**InputParameters** | Pointer to **interface{}** | Capture the mapping of service item ActionDefinition inputDefinition to catalog service request. | [optional] 
**OperationType** | Pointer to **string** | Type of action operation to be executed on the service item. * &#x60;PostDeployment&#x60; - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * &#x60;Deployment&#x60; - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * &#x60;Decommission&#x60; - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item. | [optional] [readonly] [default to "PostDeployment"]
**ServiceItemActionDefinition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemInputDefinitionTypeAllOf

`func NewWorkflowServiceItemInputDefinitionTypeAllOf(classId string, objectType string, ) *WorkflowServiceItemInputDefinitionTypeAllOf`

NewWorkflowServiceItemInputDefinitionTypeAllOf instantiates a new WorkflowServiceItemInputDefinitionTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemInputDefinitionTypeAllOfWithDefaults

`func NewWorkflowServiceItemInputDefinitionTypeAllOfWithDefaults() *WorkflowServiceItemInputDefinitionTypeAllOf`

NewWorkflowServiceItemInputDefinitionTypeAllOfWithDefaults instantiates a new WorkflowServiceItemInputDefinitionTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionName

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetInputParameters

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetOperationType

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetServiceItemActionDefinition

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetServiceItemActionDefinition() MoMoRef`

GetServiceItemActionDefinition returns the ServiceItemActionDefinition field if non-nil, zero value otherwise.

### GetServiceItemActionDefinitionOk

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) GetServiceItemActionDefinitionOk() (*MoMoRef, bool)`

GetServiceItemActionDefinitionOk returns a tuple with the ServiceItemActionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemActionDefinition

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) SetServiceItemActionDefinition(v MoMoRef)`

SetServiceItemActionDefinition sets ServiceItemActionDefinition field to given value.

### HasServiceItemActionDefinition

`func (o *WorkflowServiceItemInputDefinitionTypeAllOf) HasServiceItemActionDefinition() bool`

HasServiceItemActionDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


