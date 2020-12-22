# WorkflowPropertiesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.Properties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.Properties"]
**ExternalMeta** | Pointer to **bool** | When set to false the task definition can only be used by internal system workflows. When set to true then the task can be included in user defined workflows. | [optional] [default to false]
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**OutputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**RetryCount** | Pointer to **int64** | The number of times a task should be tried before marking as failed. | [optional] [default to 3]
**RetryDelay** | Pointer to **int64** | The delay in seconds after which the the task is re-tried. | [optional] [default to 60]
**RetryPolicy** | Pointer to **string** | The retry policy for the task. * &#x60;Fixed&#x60; - The enum specifies the option as Fixed where the task retry happens after fixed time specified by RetryDelay. | [optional] [default to "Fixed"]
**SupportStatus** | Pointer to **string** | Supported status of the definition. * &#x60;Supported&#x60; - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * &#x60;Beta&#x60; - The definition is a Beta version and this version can under go changes until the version is marked supported. * &#x60;Deprecated&#x60; - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. | [optional] [default to "Supported"]
**Timeout** | Pointer to **int64** | The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days. | [optional] [default to 600]
**TimeoutPolicy** | Pointer to **string** | The timeout policy for the task. * &#x60;Timeout&#x60; - The enum specifies the option as Timeout where task will be timed out after the specified time in Timeout property. * &#x60;Retry&#x60; - The enum specifies the option as Retry where task will be re-tried. | [optional] [default to "Timeout"]

## Methods

### NewWorkflowPropertiesAllOf

`func NewWorkflowPropertiesAllOf(classId string, objectType string, ) *WorkflowPropertiesAllOf`

NewWorkflowPropertiesAllOf instantiates a new WorkflowPropertiesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPropertiesAllOfWithDefaults

`func NewWorkflowPropertiesAllOfWithDefaults() *WorkflowPropertiesAllOf`

NewWorkflowPropertiesAllOfWithDefaults instantiates a new WorkflowPropertiesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowPropertiesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowPropertiesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowPropertiesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowPropertiesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowPropertiesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowPropertiesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalMeta

`func (o *WorkflowPropertiesAllOf) GetExternalMeta() bool`

GetExternalMeta returns the ExternalMeta field if non-nil, zero value otherwise.

### GetExternalMetaOk

`func (o *WorkflowPropertiesAllOf) GetExternalMetaOk() (*bool, bool)`

GetExternalMetaOk returns a tuple with the ExternalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMeta

`func (o *WorkflowPropertiesAllOf) SetExternalMeta(v bool)`

SetExternalMeta sets ExternalMeta field to given value.

### HasExternalMeta

`func (o *WorkflowPropertiesAllOf) HasExternalMeta() bool`

HasExternalMeta returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowPropertiesAllOf) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowPropertiesAllOf) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowPropertiesAllOf) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowPropertiesAllOf) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowPropertiesAllOf) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowPropertiesAllOf) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetOutputDefinition

`func (o *WorkflowPropertiesAllOf) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowPropertiesAllOf) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowPropertiesAllOf) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowPropertiesAllOf) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### SetOutputDefinitionNil

`func (o *WorkflowPropertiesAllOf) SetOutputDefinitionNil(b bool)`

 SetOutputDefinitionNil sets the value for OutputDefinition to be an explicit nil

### UnsetOutputDefinition
`func (o *WorkflowPropertiesAllOf) UnsetOutputDefinition()`

UnsetOutputDefinition ensures that no value is present for OutputDefinition, not even an explicit nil
### GetRetryCount

`func (o *WorkflowPropertiesAllOf) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowPropertiesAllOf) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowPropertiesAllOf) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowPropertiesAllOf) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelay

`func (o *WorkflowPropertiesAllOf) GetRetryDelay() int64`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *WorkflowPropertiesAllOf) GetRetryDelayOk() (*int64, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *WorkflowPropertiesAllOf) SetRetryDelay(v int64)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *WorkflowPropertiesAllOf) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetRetryPolicy

`func (o *WorkflowPropertiesAllOf) GetRetryPolicy() string`

GetRetryPolicy returns the RetryPolicy field if non-nil, zero value otherwise.

### GetRetryPolicyOk

`func (o *WorkflowPropertiesAllOf) GetRetryPolicyOk() (*string, bool)`

GetRetryPolicyOk returns a tuple with the RetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryPolicy

`func (o *WorkflowPropertiesAllOf) SetRetryPolicy(v string)`

SetRetryPolicy sets RetryPolicy field to given value.

### HasRetryPolicy

`func (o *WorkflowPropertiesAllOf) HasRetryPolicy() bool`

HasRetryPolicy returns a boolean if a field has been set.

### GetSupportStatus

`func (o *WorkflowPropertiesAllOf) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *WorkflowPropertiesAllOf) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *WorkflowPropertiesAllOf) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *WorkflowPropertiesAllOf) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.

### GetTimeout

`func (o *WorkflowPropertiesAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WorkflowPropertiesAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WorkflowPropertiesAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WorkflowPropertiesAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTimeoutPolicy

`func (o *WorkflowPropertiesAllOf) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *WorkflowPropertiesAllOf) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *WorkflowPropertiesAllOf) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *WorkflowPropertiesAllOf) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


