# WorkflowServiceItemActionProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemActionProperties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemActionProperties"]
**OperationType** | Pointer to **string** | Type of action operation to be executed on the service item. * &#x60;PostDeployment&#x60; - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * &#x60;Deployment&#x60; - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * &#x60;Decommission&#x60; - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item. * &#x60;Migration&#x60; - This represents the migration action, used to migrate service item instance from one service item definition version to another service item definition version. All valid service items can have up to one operation type marked as Migration. Once a migration action is running on a service item instance, no further operations are allowed on that service item instance during the migration process. | [optional] [default to "PostDeployment"]
**Properties** | Pointer to **interface{}** | The properties of the action. The actual structure of properties can vary based on the operationType. | [optional] 
**StopOnFailure** | Pointer to **bool** | When true, the action on the service item will be stopped when it reaches a failure by either calling the configured stop workflow or by calling the rollback workflow. By default value is set to true. | [optional] [default to true]

## Methods

### NewWorkflowServiceItemActionProperties

`func NewWorkflowServiceItemActionProperties(classId string, objectType string, ) *WorkflowServiceItemActionProperties`

NewWorkflowServiceItemActionProperties instantiates a new WorkflowServiceItemActionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemActionPropertiesWithDefaults

`func NewWorkflowServiceItemActionPropertiesWithDefaults() *WorkflowServiceItemActionProperties`

NewWorkflowServiceItemActionPropertiesWithDefaults instantiates a new WorkflowServiceItemActionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemActionProperties) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemActionProperties) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemActionProperties) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemActionProperties) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemActionProperties) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemActionProperties) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationType

`func (o *WorkflowServiceItemActionProperties) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *WorkflowServiceItemActionProperties) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *WorkflowServiceItemActionProperties) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *WorkflowServiceItemActionProperties) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowServiceItemActionProperties) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowServiceItemActionProperties) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowServiceItemActionProperties) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowServiceItemActionProperties) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowServiceItemActionProperties) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowServiceItemActionProperties) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetStopOnFailure

`func (o *WorkflowServiceItemActionProperties) GetStopOnFailure() bool`

GetStopOnFailure returns the StopOnFailure field if non-nil, zero value otherwise.

### GetStopOnFailureOk

`func (o *WorkflowServiceItemActionProperties) GetStopOnFailureOk() (*bool, bool)`

GetStopOnFailureOk returns a tuple with the StopOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnFailure

`func (o *WorkflowServiceItemActionProperties) SetStopOnFailure(v bool)`

SetStopOnFailure sets StopOnFailure field to given value.

### HasStopOnFailure

`func (o *WorkflowServiceItemActionProperties) HasStopOnFailure() bool`

HasStopOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


