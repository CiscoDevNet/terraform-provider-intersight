# WorkflowWorkflowInfoProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowInfoProperties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowInfoProperties"]
**Retryable** | Pointer to **bool** | When true, this workflow can be retried if has not been modified for more than a period of 2 weeks. | [optional] [default to false]
**RollbackAction** | Pointer to **string** | Status of rollback for this workflow instance. The rollback action of the workflow can be enabled, disabled, completed. * &#x60;Disabled&#x60; - Status of the rollback action when workflow is disabled for rollback. * &#x60;Enabled&#x60; - Status of the rollback action when workflow is enabled for rollback. * &#x60;Completed&#x60; - Status of the rollback action once workflow completes the rollback for all eligible tasks. | [optional] [readonly] [default to "Disabled"]

## Methods

### NewWorkflowWorkflowInfoProperties

`func NewWorkflowWorkflowInfoProperties(classId string, objectType string, ) *WorkflowWorkflowInfoProperties`

NewWorkflowWorkflowInfoProperties instantiates a new WorkflowWorkflowInfoProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowInfoPropertiesWithDefaults

`func NewWorkflowWorkflowInfoPropertiesWithDefaults() *WorkflowWorkflowInfoProperties`

NewWorkflowWorkflowInfoPropertiesWithDefaults instantiates a new WorkflowWorkflowInfoProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowInfoProperties) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowInfoProperties) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowInfoProperties) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowInfoProperties) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowInfoProperties) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowInfoProperties) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRetryable

`func (o *WorkflowWorkflowInfoProperties) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *WorkflowWorkflowInfoProperties) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *WorkflowWorkflowInfoProperties) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *WorkflowWorkflowInfoProperties) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRollbackAction

`func (o *WorkflowWorkflowInfoProperties) GetRollbackAction() string`

GetRollbackAction returns the RollbackAction field if non-nil, zero value otherwise.

### GetRollbackActionOk

`func (o *WorkflowWorkflowInfoProperties) GetRollbackActionOk() (*string, bool)`

GetRollbackActionOk returns a tuple with the RollbackAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackAction

`func (o *WorkflowWorkflowInfoProperties) SetRollbackAction(v string)`

SetRollbackAction sets RollbackAction field to given value.

### HasRollbackAction

`func (o *WorkflowWorkflowInfoProperties) HasRollbackAction() bool`

HasRollbackAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


