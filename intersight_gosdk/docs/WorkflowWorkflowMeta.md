# WorkflowWorkflowMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowMeta"]
**Description** | Pointer to **string** | The description for the workflow. | [optional] 
**InputParameters** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | The name given to the workflow. | [optional] 
**OutputParameters** | Pointer to **interface{}** | The workflow output parameters. | [optional] 
**Retryable** | Pointer to **bool** | When true, this workflow can be retried for 2 weeks since the last modification of the workflow. | [optional] [default to false]
**SchemaVersion** | Pointer to **int64** | The Conductor schema version that decides what attribute can be supported. | [optional] 
**Src** | Pointer to **string** | The src is workflow owner service. | [optional] 
**Tasks** | Pointer to **interface{}** | The tasks contained inside of the workflow. | [optional] 
**Type** | Pointer to **string** | The type of workflow definition. * &#x60;SystemDefined&#x60; - System defined workflow definition. * &#x60;UserDefined&#x60; - User defined workflow definition. * &#x60;Dynamic&#x60; - Dynamically defined workflow definition. | [optional] [default to "SystemDefined"]
**Version** | Pointer to **int64** | The version for the workflow so we can support multiple versions for the same workflow name. | [optional] [default to 1]
**WaitOnDuplicate** | Pointer to **bool** | Parameter decides if workflows will wait for a duplicate to finish before starting a new one. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewWorkflowWorkflowMeta

`func NewWorkflowWorkflowMeta(classId string, objectType string, ) *WorkflowWorkflowMeta`

NewWorkflowWorkflowMeta instantiates a new WorkflowWorkflowMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowMetaWithDefaults

`func NewWorkflowWorkflowMetaWithDefaults() *WorkflowWorkflowMeta`

NewWorkflowWorkflowMetaWithDefaults instantiates a new WorkflowWorkflowMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowWorkflowMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowWorkflowMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowWorkflowMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowWorkflowMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowWorkflowMeta) GetInputParameters() []string`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowWorkflowMeta) GetInputParametersOk() (*[]string, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowWorkflowMeta) SetInputParameters(v []string)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowWorkflowMeta) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowWorkflowMeta) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowWorkflowMeta) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetName

`func (o *WorkflowWorkflowMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputParameters

`func (o *WorkflowWorkflowMeta) GetOutputParameters() interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *WorkflowWorkflowMeta) GetOutputParametersOk() (*interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *WorkflowWorkflowMeta) SetOutputParameters(v interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *WorkflowWorkflowMeta) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *WorkflowWorkflowMeta) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *WorkflowWorkflowMeta) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetRetryable

`func (o *WorkflowWorkflowMeta) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *WorkflowWorkflowMeta) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *WorkflowWorkflowMeta) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *WorkflowWorkflowMeta) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *WorkflowWorkflowMeta) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WorkflowWorkflowMeta) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WorkflowWorkflowMeta) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *WorkflowWorkflowMeta) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowWorkflowMeta) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowWorkflowMeta) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowWorkflowMeta) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowWorkflowMeta) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetTasks

`func (o *WorkflowWorkflowMeta) GetTasks() interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowWorkflowMeta) GetTasksOk() (*interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowWorkflowMeta) SetTasks(v interface{})`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowWorkflowMeta) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *WorkflowWorkflowMeta) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *WorkflowWorkflowMeta) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetType

`func (o *WorkflowWorkflowMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowWorkflowMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowWorkflowMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowWorkflowMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowWorkflowMeta) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowWorkflowMeta) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowWorkflowMeta) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowWorkflowMeta) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWaitOnDuplicate

`func (o *WorkflowWorkflowMeta) GetWaitOnDuplicate() bool`

GetWaitOnDuplicate returns the WaitOnDuplicate field if non-nil, zero value otherwise.

### GetWaitOnDuplicateOk

`func (o *WorkflowWorkflowMeta) GetWaitOnDuplicateOk() (*bool, bool)`

GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitOnDuplicate

`func (o *WorkflowWorkflowMeta) SetWaitOnDuplicate(v bool)`

SetWaitOnDuplicate sets WaitOnDuplicate field to given value.

### HasWaitOnDuplicate

`func (o *WorkflowWorkflowMeta) HasWaitOnDuplicate() bool`

HasWaitOnDuplicate returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowWorkflowMeta) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowWorkflowMeta) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowWorkflowMeta) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowWorkflowMeta) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


