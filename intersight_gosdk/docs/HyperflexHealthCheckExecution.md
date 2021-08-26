# HyperflexHealthCheckExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HealthCheckExecution"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HealthCheckExecution"]
**Category** | Pointer to **string** | Category that the HyperFlex health check Definition belongs to. | [optional] [readonly] 
**Cause** | Pointer to **string** | Information detailing the possible cause of the healthcheck failure, if the check fails. | [optional] 
**CompletionTime** | Pointer to **time.Time** | Health check execution completion time. | [optional] [readonly] 
**HealthCheckDetails** | Pointer to **string** | Details of the health check execution result. | [optional] [readonly] 
**HealthCheckExecutionErrorDetails** | Pointer to **string** | Error details of a script execution failure. | [optional] [readonly] 
**HealthCheckExecutionErrorSummary** | Pointer to **string** | Error summary of a script execution failure. | [optional] [readonly] 
**HealthCheckExecutionStatus** | Pointer to **string** | Status of the health check execution. * &#x60;UNKNOWN&#x60; - Indicates that the health heck execution results are unknown. * &#x60;SUCCEEDED&#x60; - Indicates that the health check execution succeeded. * &#x60;FAILED&#x60; - Indicates that the health check execution failed. * &#x60;TIMED_OUT&#x60; - Indicates that the health check execution timed out before completion. | [optional] [readonly] [default to "UNKNOWN"]
**HealthCheckResult** | Pointer to **string** | Health check execution result. Valid only if HealthCheckExecutionStatus is SUCCEEDED. * &#x60;UNKNOWN&#x60; - Indicates that the health check results could not be determined. * &#x60;PASS&#x60; - Indicates that the health check passed. * &#x60;FAIL&#x60; - Indicates that the health check failed. * &#x60;WARN&#x60; - Indicates that the health check completed with a warning. * &#x60;NOT_APPLICABLE&#x60; - Indicates that the health check is either unsupported, or not applicable on the Cluster. | [optional] [readonly] [default to "UNKNOWN"]
**HealthCheckSummary** | Pointer to **string** | A brief summary of health check results. | [optional] [readonly] 
**HxDeviceName** | Pointer to **string** | HyperFlex Device Name where the healthcheck is executed. | [optional] [readonly] 
**SuggestedResolution** | Pointer to **string** | Information detailing a suggested resolution for the healthcheck failure, if the check fails. | [optional] 
**Uuid** | Pointer to **string** | UUID of an instance of health check execution. | [optional] [readonly] 
**HealthCheckDefinition** | Pointer to [**HyperflexHealthCheckDefinitionRelationship**](HyperflexHealthCheckDefinitionRelationship.md) |  | [optional] 
**HxCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHealthCheckExecution

`func NewHyperflexHealthCheckExecution(classId string, objectType string, ) *HyperflexHealthCheckExecution`

NewHyperflexHealthCheckExecution instantiates a new HyperflexHealthCheckExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckExecutionWithDefaults

`func NewHyperflexHealthCheckExecutionWithDefaults() *HyperflexHealthCheckExecution`

NewHyperflexHealthCheckExecutionWithDefaults instantiates a new HyperflexHealthCheckExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckExecution) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckExecution) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckExecution) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckExecution) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckExecution) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckExecution) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HyperflexHealthCheckExecution) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HyperflexHealthCheckExecution) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HyperflexHealthCheckExecution) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HyperflexHealthCheckExecution) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCause

`func (o *HyperflexHealthCheckExecution) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *HyperflexHealthCheckExecution) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *HyperflexHealthCheckExecution) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *HyperflexHealthCheckExecution) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCompletionTime

`func (o *HyperflexHealthCheckExecution) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *HyperflexHealthCheckExecution) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *HyperflexHealthCheckExecution) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *HyperflexHealthCheckExecution) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetHealthCheckDetails

`func (o *HyperflexHealthCheckExecution) GetHealthCheckDetails() string`

GetHealthCheckDetails returns the HealthCheckDetails field if non-nil, zero value otherwise.

### GetHealthCheckDetailsOk

`func (o *HyperflexHealthCheckExecution) GetHealthCheckDetailsOk() (*string, bool)`

GetHealthCheckDetailsOk returns a tuple with the HealthCheckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDetails

`func (o *HyperflexHealthCheckExecution) SetHealthCheckDetails(v string)`

SetHealthCheckDetails sets HealthCheckDetails field to given value.

### HasHealthCheckDetails

`func (o *HyperflexHealthCheckExecution) HasHealthCheckDetails() bool`

HasHealthCheckDetails returns a boolean if a field has been set.

### GetHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecution) GetHealthCheckExecutionErrorDetails() string`

GetHealthCheckExecutionErrorDetails returns the HealthCheckExecutionErrorDetails field if non-nil, zero value otherwise.

### GetHealthCheckExecutionErrorDetailsOk

`func (o *HyperflexHealthCheckExecution) GetHealthCheckExecutionErrorDetailsOk() (*string, bool)`

GetHealthCheckExecutionErrorDetailsOk returns a tuple with the HealthCheckExecutionErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecution) SetHealthCheckExecutionErrorDetails(v string)`

SetHealthCheckExecutionErrorDetails sets HealthCheckExecutionErrorDetails field to given value.

### HasHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecution) HasHealthCheckExecutionErrorDetails() bool`

HasHealthCheckExecutionErrorDetails returns a boolean if a field has been set.

### GetHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecution) GetHealthCheckExecutionErrorSummary() string`

GetHealthCheckExecutionErrorSummary returns the HealthCheckExecutionErrorSummary field if non-nil, zero value otherwise.

### GetHealthCheckExecutionErrorSummaryOk

`func (o *HyperflexHealthCheckExecution) GetHealthCheckExecutionErrorSummaryOk() (*string, bool)`

GetHealthCheckExecutionErrorSummaryOk returns a tuple with the HealthCheckExecutionErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecution) SetHealthCheckExecutionErrorSummary(v string)`

SetHealthCheckExecutionErrorSummary sets HealthCheckExecutionErrorSummary field to given value.

### HasHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecution) HasHealthCheckExecutionErrorSummary() bool`

HasHealthCheckExecutionErrorSummary returns a boolean if a field has been set.

### GetHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecution) GetHealthCheckExecutionStatus() string`

GetHealthCheckExecutionStatus returns the HealthCheckExecutionStatus field if non-nil, zero value otherwise.

### GetHealthCheckExecutionStatusOk

`func (o *HyperflexHealthCheckExecution) GetHealthCheckExecutionStatusOk() (*string, bool)`

GetHealthCheckExecutionStatusOk returns a tuple with the HealthCheckExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecution) SetHealthCheckExecutionStatus(v string)`

SetHealthCheckExecutionStatus sets HealthCheckExecutionStatus field to given value.

### HasHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecution) HasHealthCheckExecutionStatus() bool`

HasHealthCheckExecutionStatus returns a boolean if a field has been set.

### GetHealthCheckResult

`func (o *HyperflexHealthCheckExecution) GetHealthCheckResult() string`

GetHealthCheckResult returns the HealthCheckResult field if non-nil, zero value otherwise.

### GetHealthCheckResultOk

`func (o *HyperflexHealthCheckExecution) GetHealthCheckResultOk() (*string, bool)`

GetHealthCheckResultOk returns a tuple with the HealthCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckResult

`func (o *HyperflexHealthCheckExecution) SetHealthCheckResult(v string)`

SetHealthCheckResult sets HealthCheckResult field to given value.

### HasHealthCheckResult

`func (o *HyperflexHealthCheckExecution) HasHealthCheckResult() bool`

HasHealthCheckResult returns a boolean if a field has been set.

### GetHealthCheckSummary

`func (o *HyperflexHealthCheckExecution) GetHealthCheckSummary() string`

GetHealthCheckSummary returns the HealthCheckSummary field if non-nil, zero value otherwise.

### GetHealthCheckSummaryOk

`func (o *HyperflexHealthCheckExecution) GetHealthCheckSummaryOk() (*string, bool)`

GetHealthCheckSummaryOk returns a tuple with the HealthCheckSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckSummary

`func (o *HyperflexHealthCheckExecution) SetHealthCheckSummary(v string)`

SetHealthCheckSummary sets HealthCheckSummary field to given value.

### HasHealthCheckSummary

`func (o *HyperflexHealthCheckExecution) HasHealthCheckSummary() bool`

HasHealthCheckSummary returns a boolean if a field has been set.

### GetHxDeviceName

`func (o *HyperflexHealthCheckExecution) GetHxDeviceName() string`

GetHxDeviceName returns the HxDeviceName field if non-nil, zero value otherwise.

### GetHxDeviceNameOk

`func (o *HyperflexHealthCheckExecution) GetHxDeviceNameOk() (*string, bool)`

GetHxDeviceNameOk returns a tuple with the HxDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxDeviceName

`func (o *HyperflexHealthCheckExecution) SetHxDeviceName(v string)`

SetHxDeviceName sets HxDeviceName field to given value.

### HasHxDeviceName

`func (o *HyperflexHealthCheckExecution) HasHxDeviceName() bool`

HasHxDeviceName returns a boolean if a field has been set.

### GetSuggestedResolution

`func (o *HyperflexHealthCheckExecution) GetSuggestedResolution() string`

GetSuggestedResolution returns the SuggestedResolution field if non-nil, zero value otherwise.

### GetSuggestedResolutionOk

`func (o *HyperflexHealthCheckExecution) GetSuggestedResolutionOk() (*string, bool)`

GetSuggestedResolutionOk returns a tuple with the SuggestedResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedResolution

`func (o *HyperflexHealthCheckExecution) SetSuggestedResolution(v string)`

SetSuggestedResolution sets SuggestedResolution field to given value.

### HasSuggestedResolution

`func (o *HyperflexHealthCheckExecution) HasSuggestedResolution() bool`

HasSuggestedResolution returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexHealthCheckExecution) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHealthCheckExecution) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHealthCheckExecution) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHealthCheckExecution) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHealthCheckDefinition

`func (o *HyperflexHealthCheckExecution) GetHealthCheckDefinition() HyperflexHealthCheckDefinitionRelationship`

GetHealthCheckDefinition returns the HealthCheckDefinition field if non-nil, zero value otherwise.

### GetHealthCheckDefinitionOk

`func (o *HyperflexHealthCheckExecution) GetHealthCheckDefinitionOk() (*HyperflexHealthCheckDefinitionRelationship, bool)`

GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDefinition

`func (o *HyperflexHealthCheckExecution) SetHealthCheckDefinition(v HyperflexHealthCheckDefinitionRelationship)`

SetHealthCheckDefinition sets HealthCheckDefinition field to given value.

### HasHealthCheckDefinition

`func (o *HyperflexHealthCheckExecution) HasHealthCheckDefinition() bool`

HasHealthCheckDefinition returns a boolean if a field has been set.

### GetHxCluster

`func (o *HyperflexHealthCheckExecution) GetHxCluster() HyperflexClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *HyperflexHealthCheckExecution) GetHxClusterOk() (*HyperflexClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *HyperflexHealthCheckExecution) SetHxCluster(v HyperflexClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *HyperflexHealthCheckExecution) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexHealthCheckExecution) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexHealthCheckExecution) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexHealthCheckExecution) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexHealthCheckExecution) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


