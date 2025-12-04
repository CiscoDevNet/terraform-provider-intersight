# WorkflowWorkflowProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowProperties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowProperties"]
**Cloneable** | Pointer to **bool** | When set to false workflow is not cloneable. It is set to true only if Workflow is not internal and it does not have any internal tasks. | [optional] [readonly] [default to true]
**EnableDebug** | Pointer to **bool** | Enabling this flag will capture request and response details as debug logs for tasks that are using workflow.BatchApi for implementation. For other tasks in the workflow which are not based on workflow.BatchApi logs will not be generated. | [optional] [default to false]
**EnablePublishStatus** | Pointer to **bool** | This flag determines if this workflow publish status is enforced or not. | [optional] [default to false]
**ExternalMeta** | Pointer to **bool** | When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities. | [optional] [readonly] [default to false]
**PublishStatus** | Pointer to **string** | The workflow publish status (Draft, Published, Archived), this property is relevant only when enablePublishStatus is set to true. * &#x60;Draft&#x60; - The enum specifies the option as Draft which means the meta definition is being designed and tested. * &#x60;Published&#x60; - The enum specifies the option as Published which means the meta definition is ready for consumption. * &#x60;Archived&#x60; - The enum specifies the option as Archived which means the meta definition is archived and can no longer be consumed. | [optional] [default to "Draft"]
**Retryable** | Pointer to **bool** | When set to true, the failed workflow executions from this workflow definition can be retried for up to 2 weeks since the last modification time. After two weeks of inactivity on the workflow execution, the option to retry the failed workflow will be disabled. | [optional] [default to false]
**RollbackOnCancel** | Pointer to **bool** | When set to true, the changes are automatically rolled back if the workflow execution is canceled. | [optional] [default to false]
**RollbackOnFailure** | Pointer to **bool** | When set to true, the changes are automatically rolled back if the workflow fails to execute. | [optional] [default to false]
**SupportStatus** | Pointer to **string** | Supported status of the definition. * &#x60;Supported&#x60; - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * &#x60;Beta&#x60; - The definition is a Beta version and this version can under go changes until the version is marked supported. * &#x60;Deprecated&#x60; - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. | [optional] [default to "Supported"]
**TargetCtxPolicy** | Pointer to **string** | The target context policy that controls how target context is set for this workflow when it is run as a subworkflow. * &#x60;InheritFromParent&#x60; - The target context is inherited from the parent workflow and additional targets from this current workflow are also added. Use this setting when the parent workflow is operating on a target set which is still applicable for this current subworkflow. * &#x60;OnlyThisWorkflowTargets&#x60; - The target context is set only from the targets defined in the current workflow and nothing is inherited from parent workflow. Use this setting if the parent workflow is operating on a superset of targets and this current workflow is operating on a subset of those targets or a completely different set of targets. | [optional] [default to "InheritFromParent"]

## Methods

### NewWorkflowWorkflowProperties

`func NewWorkflowWorkflowProperties(classId string, objectType string, ) *WorkflowWorkflowProperties`

NewWorkflowWorkflowProperties instantiates a new WorkflowWorkflowProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowPropertiesWithDefaults

`func NewWorkflowWorkflowPropertiesWithDefaults() *WorkflowWorkflowProperties`

NewWorkflowWorkflowPropertiesWithDefaults instantiates a new WorkflowWorkflowProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowProperties) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowProperties) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowProperties) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowProperties) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowProperties) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowProperties) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCloneable

`func (o *WorkflowWorkflowProperties) GetCloneable() bool`

GetCloneable returns the Cloneable field if non-nil, zero value otherwise.

### GetCloneableOk

`func (o *WorkflowWorkflowProperties) GetCloneableOk() (*bool, bool)`

GetCloneableOk returns a tuple with the Cloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneable

`func (o *WorkflowWorkflowProperties) SetCloneable(v bool)`

SetCloneable sets Cloneable field to given value.

### HasCloneable

`func (o *WorkflowWorkflowProperties) HasCloneable() bool`

HasCloneable returns a boolean if a field has been set.

### GetEnableDebug

`func (o *WorkflowWorkflowProperties) GetEnableDebug() bool`

GetEnableDebug returns the EnableDebug field if non-nil, zero value otherwise.

### GetEnableDebugOk

`func (o *WorkflowWorkflowProperties) GetEnableDebugOk() (*bool, bool)`

GetEnableDebugOk returns a tuple with the EnableDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebug

`func (o *WorkflowWorkflowProperties) SetEnableDebug(v bool)`

SetEnableDebug sets EnableDebug field to given value.

### HasEnableDebug

`func (o *WorkflowWorkflowProperties) HasEnableDebug() bool`

HasEnableDebug returns a boolean if a field has been set.

### GetEnablePublishStatus

`func (o *WorkflowWorkflowProperties) GetEnablePublishStatus() bool`

GetEnablePublishStatus returns the EnablePublishStatus field if non-nil, zero value otherwise.

### GetEnablePublishStatusOk

`func (o *WorkflowWorkflowProperties) GetEnablePublishStatusOk() (*bool, bool)`

GetEnablePublishStatusOk returns a tuple with the EnablePublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublishStatus

`func (o *WorkflowWorkflowProperties) SetEnablePublishStatus(v bool)`

SetEnablePublishStatus sets EnablePublishStatus field to given value.

### HasEnablePublishStatus

`func (o *WorkflowWorkflowProperties) HasEnablePublishStatus() bool`

HasEnablePublishStatus returns a boolean if a field has been set.

### GetExternalMeta

`func (o *WorkflowWorkflowProperties) GetExternalMeta() bool`

GetExternalMeta returns the ExternalMeta field if non-nil, zero value otherwise.

### GetExternalMetaOk

`func (o *WorkflowWorkflowProperties) GetExternalMetaOk() (*bool, bool)`

GetExternalMetaOk returns a tuple with the ExternalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMeta

`func (o *WorkflowWorkflowProperties) SetExternalMeta(v bool)`

SetExternalMeta sets ExternalMeta field to given value.

### HasExternalMeta

`func (o *WorkflowWorkflowProperties) HasExternalMeta() bool`

HasExternalMeta returns a boolean if a field has been set.

### GetPublishStatus

`func (o *WorkflowWorkflowProperties) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *WorkflowWorkflowProperties) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *WorkflowWorkflowProperties) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *WorkflowWorkflowProperties) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetRetryable

`func (o *WorkflowWorkflowProperties) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *WorkflowWorkflowProperties) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *WorkflowWorkflowProperties) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *WorkflowWorkflowProperties) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRollbackOnCancel

`func (o *WorkflowWorkflowProperties) GetRollbackOnCancel() bool`

GetRollbackOnCancel returns the RollbackOnCancel field if non-nil, zero value otherwise.

### GetRollbackOnCancelOk

`func (o *WorkflowWorkflowProperties) GetRollbackOnCancelOk() (*bool, bool)`

GetRollbackOnCancelOk returns a tuple with the RollbackOnCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackOnCancel

`func (o *WorkflowWorkflowProperties) SetRollbackOnCancel(v bool)`

SetRollbackOnCancel sets RollbackOnCancel field to given value.

### HasRollbackOnCancel

`func (o *WorkflowWorkflowProperties) HasRollbackOnCancel() bool`

HasRollbackOnCancel returns a boolean if a field has been set.

### GetRollbackOnFailure

`func (o *WorkflowWorkflowProperties) GetRollbackOnFailure() bool`

GetRollbackOnFailure returns the RollbackOnFailure field if non-nil, zero value otherwise.

### GetRollbackOnFailureOk

`func (o *WorkflowWorkflowProperties) GetRollbackOnFailureOk() (*bool, bool)`

GetRollbackOnFailureOk returns a tuple with the RollbackOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackOnFailure

`func (o *WorkflowWorkflowProperties) SetRollbackOnFailure(v bool)`

SetRollbackOnFailure sets RollbackOnFailure field to given value.

### HasRollbackOnFailure

`func (o *WorkflowWorkflowProperties) HasRollbackOnFailure() bool`

HasRollbackOnFailure returns a boolean if a field has been set.

### GetSupportStatus

`func (o *WorkflowWorkflowProperties) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *WorkflowWorkflowProperties) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *WorkflowWorkflowProperties) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *WorkflowWorkflowProperties) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.

### GetTargetCtxPolicy

`func (o *WorkflowWorkflowProperties) GetTargetCtxPolicy() string`

GetTargetCtxPolicy returns the TargetCtxPolicy field if non-nil, zero value otherwise.

### GetTargetCtxPolicyOk

`func (o *WorkflowWorkflowProperties) GetTargetCtxPolicyOk() (*string, bool)`

GetTargetCtxPolicyOk returns a tuple with the TargetCtxPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCtxPolicy

`func (o *WorkflowWorkflowProperties) SetTargetCtxPolicy(v string)`

SetTargetCtxPolicy sets TargetCtxPolicy field to given value.

### HasTargetCtxPolicy

`func (o *WorkflowWorkflowProperties) HasTargetCtxPolicy() bool`

HasTargetCtxPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


