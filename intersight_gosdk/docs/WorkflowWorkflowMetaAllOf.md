# WorkflowWorkflowMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description for the workflow. | [optional] 
**InputParameters** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | The name given to the workflow. | [optional] 
**OutputParameters** | Pointer to **interface{}** | The workflow output parameters. | [optional] 
**Retryable** | Pointer to **bool** | When true, this workflow can be retried for 2 weeks since the last modification of the workflow. | [optional] 
**SchemaVersion** | Pointer to **int64** | The Conductor schema version that decides what attribute can be supported. | [optional] 
**Src** | Pointer to **string** | The src is workflow owner service. | [optional] 
**Tasks** | Pointer to **interface{}** | The tasks contained inside of the workflow. | [optional] 
**Type** | Pointer to **string** | The type of workflow definition. * &#x60;SystemDefined&#x60; - System defined workflow definition. * &#x60;UserDefined&#x60; - User defined workflow definition. * &#x60;Dynamic&#x60; - Dynamically defined workflow definition. | [optional] [default to "SystemDefined"]
**Version** | Pointer to **int64** | The version for the workflow so we can support multiple versions for the same workflow name. | [optional] 
**WaitOnDuplicate** | Pointer to **bool** | Parameter decides if workflows will wait for a duplicate to finish before starting a new one. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowWorkflowMetaAllOf

`func NewWorkflowWorkflowMetaAllOf() *WorkflowWorkflowMetaAllOf`

NewWorkflowWorkflowMetaAllOf instantiates a new WorkflowWorkflowMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowMetaAllOfWithDefaults

`func NewWorkflowWorkflowMetaAllOfWithDefaults() *WorkflowWorkflowMetaAllOf`

NewWorkflowWorkflowMetaAllOfWithDefaults instantiates a new WorkflowWorkflowMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowWorkflowMetaAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowWorkflowMetaAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowWorkflowMetaAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowWorkflowMetaAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowWorkflowMetaAllOf) GetInputParameters() []string`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowWorkflowMetaAllOf) GetInputParametersOk() (*[]string, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowWorkflowMetaAllOf) SetInputParameters(v []string)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowWorkflowMetaAllOf) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWorkflowMetaAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowMetaAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowMetaAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowMetaAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputParameters

`func (o *WorkflowWorkflowMetaAllOf) GetOutputParameters() interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *WorkflowWorkflowMetaAllOf) GetOutputParametersOk() (*interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *WorkflowWorkflowMetaAllOf) SetOutputParameters(v interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *WorkflowWorkflowMetaAllOf) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *WorkflowWorkflowMetaAllOf) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *WorkflowWorkflowMetaAllOf) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetRetryable

`func (o *WorkflowWorkflowMetaAllOf) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *WorkflowWorkflowMetaAllOf) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *WorkflowWorkflowMetaAllOf) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *WorkflowWorkflowMetaAllOf) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *WorkflowWorkflowMetaAllOf) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WorkflowWorkflowMetaAllOf) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WorkflowWorkflowMetaAllOf) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *WorkflowWorkflowMetaAllOf) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowWorkflowMetaAllOf) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowWorkflowMetaAllOf) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowWorkflowMetaAllOf) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowWorkflowMetaAllOf) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetTasks

`func (o *WorkflowWorkflowMetaAllOf) GetTasks() interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowWorkflowMetaAllOf) GetTasksOk() (*interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowWorkflowMetaAllOf) SetTasks(v interface{})`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowWorkflowMetaAllOf) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *WorkflowWorkflowMetaAllOf) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *WorkflowWorkflowMetaAllOf) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetType

`func (o *WorkflowWorkflowMetaAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowWorkflowMetaAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowWorkflowMetaAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowWorkflowMetaAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowWorkflowMetaAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowWorkflowMetaAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowWorkflowMetaAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowWorkflowMetaAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWaitOnDuplicate

`func (o *WorkflowWorkflowMetaAllOf) GetWaitOnDuplicate() bool`

GetWaitOnDuplicate returns the WaitOnDuplicate field if non-nil, zero value otherwise.

### GetWaitOnDuplicateOk

`func (o *WorkflowWorkflowMetaAllOf) GetWaitOnDuplicateOk() (*bool, bool)`

GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitOnDuplicate

`func (o *WorkflowWorkflowMetaAllOf) SetWaitOnDuplicate(v bool)`

SetWaitOnDuplicate sets WaitOnDuplicate field to given value.

### HasWaitOnDuplicate

`func (o *WorkflowWorkflowMetaAllOf) HasWaitOnDuplicate() bool`

HasWaitOnDuplicate returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowWorkflowMetaAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowWorkflowMetaAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowWorkflowMetaAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowWorkflowMetaAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


