# WorkflowInternalPropertiesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseTaskType** | Pointer to **string** | This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask. | [optional] [readonly] 
**Constraints** | Pointer to [**WorkflowTaskConstraints**](workflow.TaskConstraints.md) |  | [optional] 
**Internal** | Pointer to **bool** | Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow. | [optional] [readonly] 
**Owner** | Pointer to **string** | The service that owns and is responsible for execution of the task. | [optional] [readonly] 

## Methods

### NewWorkflowInternalPropertiesAllOf

`func NewWorkflowInternalPropertiesAllOf() *WorkflowInternalPropertiesAllOf`

NewWorkflowInternalPropertiesAllOf instantiates a new WorkflowInternalPropertiesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInternalPropertiesAllOfWithDefaults

`func NewWorkflowInternalPropertiesAllOfWithDefaults() *WorkflowInternalPropertiesAllOf`

NewWorkflowInternalPropertiesAllOfWithDefaults instantiates a new WorkflowInternalPropertiesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseTaskType

`func (o *WorkflowInternalPropertiesAllOf) GetBaseTaskType() string`

GetBaseTaskType returns the BaseTaskType field if non-nil, zero value otherwise.

### GetBaseTaskTypeOk

`func (o *WorkflowInternalPropertiesAllOf) GetBaseTaskTypeOk() (*string, bool)`

GetBaseTaskTypeOk returns a tuple with the BaseTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTaskType

`func (o *WorkflowInternalPropertiesAllOf) SetBaseTaskType(v string)`

SetBaseTaskType sets BaseTaskType field to given value.

### HasBaseTaskType

`func (o *WorkflowInternalPropertiesAllOf) HasBaseTaskType() bool`

HasBaseTaskType returns a boolean if a field has been set.

### GetConstraints

`func (o *WorkflowInternalPropertiesAllOf) GetConstraints() WorkflowTaskConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WorkflowInternalPropertiesAllOf) GetConstraintsOk() (*WorkflowTaskConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WorkflowInternalPropertiesAllOf) SetConstraints(v WorkflowTaskConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WorkflowInternalPropertiesAllOf) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetInternal

`func (o *WorkflowInternalPropertiesAllOf) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowInternalPropertiesAllOf) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowInternalPropertiesAllOf) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowInternalPropertiesAllOf) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetOwner

`func (o *WorkflowInternalPropertiesAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowInternalPropertiesAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowInternalPropertiesAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowInternalPropertiesAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


