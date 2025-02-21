# WorkflowBaseOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**OperationType** | Pointer to **string** | Type of action operation to be performed on the service item. * &#x60;PostDeployment&#x60; - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * &#x60;Deployment&#x60; - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * &#x60;Decommission&#x60; - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item. * &#x60;Migration&#x60; - This represents the migration action, used to migrate service item instance from one service item definition version to another service item definition version. All valid service items can have up to one operation type marked as Migration. Once a migration action is running on a service item instance, no further operations are allowed on that service item instance during the migration process. | [optional] [default to "PostDeployment"]

## Methods

### NewWorkflowBaseOperation

`func NewWorkflowBaseOperation(classId string, objectType string, ) *WorkflowBaseOperation`

NewWorkflowBaseOperation instantiates a new WorkflowBaseOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBaseOperationWithDefaults

`func NewWorkflowBaseOperationWithDefaults() *WorkflowBaseOperation`

NewWorkflowBaseOperationWithDefaults instantiates a new WorkflowBaseOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBaseOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBaseOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBaseOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBaseOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBaseOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBaseOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationType

`func (o *WorkflowBaseOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *WorkflowBaseOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *WorkflowBaseOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *WorkflowBaseOperation) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


