# WorkflowInternalProperties

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

### NewWorkflowInternalProperties

`func NewWorkflowInternalProperties(classId string, objectType string, ) *WorkflowInternalProperties`

NewWorkflowInternalProperties instantiates a new WorkflowInternalProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInternalPropertiesWithDefaults

`func NewWorkflowInternalPropertiesWithDefaults() *WorkflowInternalProperties`

NewWorkflowInternalPropertiesWithDefaults instantiates a new WorkflowInternalProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowInternalProperties) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowInternalProperties) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowInternalProperties) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowInternalProperties) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowInternalProperties) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowInternalProperties) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBaseTaskType

`func (o *WorkflowInternalProperties) GetBaseTaskType() string`

GetBaseTaskType returns the BaseTaskType field if non-nil, zero value otherwise.

### GetBaseTaskTypeOk

`func (o *WorkflowInternalProperties) GetBaseTaskTypeOk() (*string, bool)`

GetBaseTaskTypeOk returns a tuple with the BaseTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTaskType

`func (o *WorkflowInternalProperties) SetBaseTaskType(v string)`

SetBaseTaskType sets BaseTaskType field to given value.

### HasBaseTaskType

`func (o *WorkflowInternalProperties) HasBaseTaskType() bool`

HasBaseTaskType returns a boolean if a field has been set.

### GetConstraints

`func (o *WorkflowInternalProperties) GetConstraints() WorkflowTaskConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WorkflowInternalProperties) GetConstraintsOk() (*WorkflowTaskConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WorkflowInternalProperties) SetConstraints(v WorkflowTaskConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WorkflowInternalProperties) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *WorkflowInternalProperties) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *WorkflowInternalProperties) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetInternal

`func (o *WorkflowInternalProperties) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowInternalProperties) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowInternalProperties) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowInternalProperties) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetOwner

`func (o *WorkflowInternalProperties) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowInternalProperties) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowInternalProperties) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowInternalProperties) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


