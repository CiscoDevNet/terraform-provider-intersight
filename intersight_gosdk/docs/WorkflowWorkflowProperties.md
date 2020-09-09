# WorkflowWorkflowProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalMeta** | Pointer to **bool** | When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities. | [optional] 
**Retryable** | Pointer to **bool** | When true, this workflow can be retried if has not been modified for more than a period of 2 weeks. | [optional] 
**SupportStatus** | Pointer to **string** | Supported status of the definition. * &#x60;Supported&#x60; - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * &#x60;Beta&#x60; - The definition is a Beta version and this version can under go changes until the version is marked supported. * &#x60;Deprecated&#x60; - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. | [optional] [default to "Supported"]

## Methods

### NewWorkflowWorkflowProperties

`func NewWorkflowWorkflowProperties() *WorkflowWorkflowProperties`

NewWorkflowWorkflowProperties instantiates a new WorkflowWorkflowProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowPropertiesWithDefaults

`func NewWorkflowWorkflowPropertiesWithDefaults() *WorkflowWorkflowProperties`

NewWorkflowWorkflowPropertiesWithDefaults instantiates a new WorkflowWorkflowProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


