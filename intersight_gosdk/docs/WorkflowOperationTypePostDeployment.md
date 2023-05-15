# WorkflowOperationTypePostDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.OperationTypePostDeployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.OperationTypePostDeployment"]
**ServiceItemActionDefinition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**ServiceItemInstance** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkflowOperationTypePostDeployment

`func NewWorkflowOperationTypePostDeployment(classId string, objectType string, ) *WorkflowOperationTypePostDeployment`

NewWorkflowOperationTypePostDeployment instantiates a new WorkflowOperationTypePostDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOperationTypePostDeploymentWithDefaults

`func NewWorkflowOperationTypePostDeploymentWithDefaults() *WorkflowOperationTypePostDeployment`

NewWorkflowOperationTypePostDeploymentWithDefaults instantiates a new WorkflowOperationTypePostDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowOperationTypePostDeployment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowOperationTypePostDeployment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowOperationTypePostDeployment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowOperationTypePostDeployment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowOperationTypePostDeployment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowOperationTypePostDeployment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetServiceItemActionDefinition

`func (o *WorkflowOperationTypePostDeployment) GetServiceItemActionDefinition() MoMoRef`

GetServiceItemActionDefinition returns the ServiceItemActionDefinition field if non-nil, zero value otherwise.

### GetServiceItemActionDefinitionOk

`func (o *WorkflowOperationTypePostDeployment) GetServiceItemActionDefinitionOk() (*MoMoRef, bool)`

GetServiceItemActionDefinitionOk returns a tuple with the ServiceItemActionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemActionDefinition

`func (o *WorkflowOperationTypePostDeployment) SetServiceItemActionDefinition(v MoMoRef)`

SetServiceItemActionDefinition sets ServiceItemActionDefinition field to given value.

### HasServiceItemActionDefinition

`func (o *WorkflowOperationTypePostDeployment) HasServiceItemActionDefinition() bool`

HasServiceItemActionDefinition returns a boolean if a field has been set.

### GetServiceItemInstance

`func (o *WorkflowOperationTypePostDeployment) GetServiceItemInstance() MoMoRef`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkflowOperationTypePostDeployment) GetServiceItemInstanceOk() (*MoMoRef, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkflowOperationTypePostDeployment) SetServiceItemInstance(v MoMoRef)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkflowOperationTypePostDeployment) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


