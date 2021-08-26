# WorkflowInternalPropertiesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.InternalProperties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.InternalProperties"]
**BaseTaskType** | Pointer to **string** | This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask. | [optional] [readonly] 
**Constraints** | Pointer to [**NullableWorkflowTaskConstraints**](WorkflowTaskConstraints.md) |  | [optional] 
**Internal** | Pointer to **bool** | Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow. | [optional] [readonly] 
**Owner** | Pointer to **string** | The service that owns and is responsible for execution of the task. | [optional] [readonly] 

## Methods

### NewWorkflowInternalPropertiesAllOf

`func NewWorkflowInternalPropertiesAllOf(classId string, objectType string, ) *WorkflowInternalPropertiesAllOf`

NewWorkflowInternalPropertiesAllOf instantiates a new WorkflowInternalPropertiesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInternalPropertiesAllOfWithDefaults

`func NewWorkflowInternalPropertiesAllOfWithDefaults() *WorkflowInternalPropertiesAllOf`

NewWorkflowInternalPropertiesAllOfWithDefaults instantiates a new WorkflowInternalPropertiesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowInternalPropertiesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowInternalPropertiesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowInternalPropertiesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowInternalPropertiesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowInternalPropertiesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowInternalPropertiesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetConstraintsNil

`func (o *WorkflowInternalPropertiesAllOf) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *WorkflowInternalPropertiesAllOf) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
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


