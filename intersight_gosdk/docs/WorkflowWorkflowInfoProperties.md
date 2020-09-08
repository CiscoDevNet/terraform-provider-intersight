# WorkflowWorkflowInfoProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retryable** | Pointer to **bool** | When true, this workflow can be retried if has not been modified for more than a period of 2 weeks. | [optional] 

## Methods

### NewWorkflowWorkflowInfoProperties

`func NewWorkflowWorkflowInfoProperties() *WorkflowWorkflowInfoProperties`

NewWorkflowWorkflowInfoProperties instantiates a new WorkflowWorkflowInfoProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowInfoPropertiesWithDefaults

`func NewWorkflowWorkflowInfoPropertiesWithDefaults() *WorkflowWorkflowInfoProperties`

NewWorkflowWorkflowInfoPropertiesWithDefaults instantiates a new WorkflowWorkflowInfoProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


