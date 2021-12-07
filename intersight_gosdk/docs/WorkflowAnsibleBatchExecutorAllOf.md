# WorkflowAnsibleBatchExecutorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.AnsibleBatchExecutor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.AnsibleBatchExecutor"]
**TaskDefinition** | Pointer to [**WorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowAnsibleBatchExecutorAllOf

`func NewWorkflowAnsibleBatchExecutorAllOf(classId string, objectType string, ) *WorkflowAnsibleBatchExecutorAllOf`

NewWorkflowAnsibleBatchExecutorAllOf instantiates a new WorkflowAnsibleBatchExecutorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAnsibleBatchExecutorAllOfWithDefaults

`func NewWorkflowAnsibleBatchExecutorAllOfWithDefaults() *WorkflowAnsibleBatchExecutorAllOf`

NewWorkflowAnsibleBatchExecutorAllOfWithDefaults instantiates a new WorkflowAnsibleBatchExecutorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowAnsibleBatchExecutorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowAnsibleBatchExecutorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowAnsibleBatchExecutorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowAnsibleBatchExecutorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAnsibleBatchExecutorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAnsibleBatchExecutorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTaskDefinition

`func (o *WorkflowAnsibleBatchExecutorAllOf) GetTaskDefinition() WorkflowTaskDefinitionRelationship`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *WorkflowAnsibleBatchExecutorAllOf) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *WorkflowAnsibleBatchExecutorAllOf) SetTaskDefinition(v WorkflowTaskDefinitionRelationship)`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *WorkflowAnsibleBatchExecutorAllOf) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


