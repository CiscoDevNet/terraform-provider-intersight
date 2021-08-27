# HyperflexHealthCheckExecutionAllOf

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

### NewHyperflexHealthCheckExecutionAllOf

`func NewHyperflexHealthCheckExecutionAllOf(classId string, objectType string, ) *HyperflexHealthCheckExecutionAllOf`

NewHyperflexHealthCheckExecutionAllOf instantiates a new HyperflexHealthCheckExecutionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckExecutionAllOfWithDefaults

`func NewHyperflexHealthCheckExecutionAllOfWithDefaults() *HyperflexHealthCheckExecutionAllOf`

NewHyperflexHealthCheckExecutionAllOfWithDefaults instantiates a new HyperflexHealthCheckExecutionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckExecutionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckExecutionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckExecutionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckExecutionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HyperflexHealthCheckExecutionAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HyperflexHealthCheckExecutionAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HyperflexHealthCheckExecutionAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCause

`func (o *HyperflexHealthCheckExecutionAllOf) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *HyperflexHealthCheckExecutionAllOf) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *HyperflexHealthCheckExecutionAllOf) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCompletionTime

`func (o *HyperflexHealthCheckExecutionAllOf) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *HyperflexHealthCheckExecutionAllOf) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *HyperflexHealthCheckExecutionAllOf) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetHealthCheckDetails

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckDetails() string`

GetHealthCheckDetails returns the HealthCheckDetails field if non-nil, zero value otherwise.

### GetHealthCheckDetailsOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckDetailsOk() (*string, bool)`

GetHealthCheckDetailsOk returns a tuple with the HealthCheckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDetails

`func (o *HyperflexHealthCheckExecutionAllOf) SetHealthCheckDetails(v string)`

SetHealthCheckDetails sets HealthCheckDetails field to given value.

### HasHealthCheckDetails

`func (o *HyperflexHealthCheckExecutionAllOf) HasHealthCheckDetails() bool`

HasHealthCheckDetails returns a boolean if a field has been set.

### GetHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckExecutionErrorDetails() string`

GetHealthCheckExecutionErrorDetails returns the HealthCheckExecutionErrorDetails field if non-nil, zero value otherwise.

### GetHealthCheckExecutionErrorDetailsOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckExecutionErrorDetailsOk() (*string, bool)`

GetHealthCheckExecutionErrorDetailsOk returns a tuple with the HealthCheckExecutionErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecutionAllOf) SetHealthCheckExecutionErrorDetails(v string)`

SetHealthCheckExecutionErrorDetails sets HealthCheckExecutionErrorDetails field to given value.

### HasHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecutionAllOf) HasHealthCheckExecutionErrorDetails() bool`

HasHealthCheckExecutionErrorDetails returns a boolean if a field has been set.

### GetHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckExecutionErrorSummary() string`

GetHealthCheckExecutionErrorSummary returns the HealthCheckExecutionErrorSummary field if non-nil, zero value otherwise.

### GetHealthCheckExecutionErrorSummaryOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckExecutionErrorSummaryOk() (*string, bool)`

GetHealthCheckExecutionErrorSummaryOk returns a tuple with the HealthCheckExecutionErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecutionAllOf) SetHealthCheckExecutionErrorSummary(v string)`

SetHealthCheckExecutionErrorSummary sets HealthCheckExecutionErrorSummary field to given value.

### HasHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecutionAllOf) HasHealthCheckExecutionErrorSummary() bool`

HasHealthCheckExecutionErrorSummary returns a boolean if a field has been set.

### GetHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckExecutionStatus() string`

GetHealthCheckExecutionStatus returns the HealthCheckExecutionStatus field if non-nil, zero value otherwise.

### GetHealthCheckExecutionStatusOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckExecutionStatusOk() (*string, bool)`

GetHealthCheckExecutionStatusOk returns a tuple with the HealthCheckExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecutionAllOf) SetHealthCheckExecutionStatus(v string)`

SetHealthCheckExecutionStatus sets HealthCheckExecutionStatus field to given value.

### HasHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecutionAllOf) HasHealthCheckExecutionStatus() bool`

HasHealthCheckExecutionStatus returns a boolean if a field has been set.

### GetHealthCheckResult

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckResult() string`

GetHealthCheckResult returns the HealthCheckResult field if non-nil, zero value otherwise.

### GetHealthCheckResultOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckResultOk() (*string, bool)`

GetHealthCheckResultOk returns a tuple with the HealthCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckResult

`func (o *HyperflexHealthCheckExecutionAllOf) SetHealthCheckResult(v string)`

SetHealthCheckResult sets HealthCheckResult field to given value.

### HasHealthCheckResult

`func (o *HyperflexHealthCheckExecutionAllOf) HasHealthCheckResult() bool`

HasHealthCheckResult returns a boolean if a field has been set.

### GetHealthCheckSummary

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckSummary() string`

GetHealthCheckSummary returns the HealthCheckSummary field if non-nil, zero value otherwise.

### GetHealthCheckSummaryOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckSummaryOk() (*string, bool)`

GetHealthCheckSummaryOk returns a tuple with the HealthCheckSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckSummary

`func (o *HyperflexHealthCheckExecutionAllOf) SetHealthCheckSummary(v string)`

SetHealthCheckSummary sets HealthCheckSummary field to given value.

### HasHealthCheckSummary

`func (o *HyperflexHealthCheckExecutionAllOf) HasHealthCheckSummary() bool`

HasHealthCheckSummary returns a boolean if a field has been set.

### GetHxDeviceName

`func (o *HyperflexHealthCheckExecutionAllOf) GetHxDeviceName() string`

GetHxDeviceName returns the HxDeviceName field if non-nil, zero value otherwise.

### GetHxDeviceNameOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHxDeviceNameOk() (*string, bool)`

GetHxDeviceNameOk returns a tuple with the HxDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxDeviceName

`func (o *HyperflexHealthCheckExecutionAllOf) SetHxDeviceName(v string)`

SetHxDeviceName sets HxDeviceName field to given value.

### HasHxDeviceName

`func (o *HyperflexHealthCheckExecutionAllOf) HasHxDeviceName() bool`

HasHxDeviceName returns a boolean if a field has been set.

### GetSuggestedResolution

`func (o *HyperflexHealthCheckExecutionAllOf) GetSuggestedResolution() string`

GetSuggestedResolution returns the SuggestedResolution field if non-nil, zero value otherwise.

### GetSuggestedResolutionOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetSuggestedResolutionOk() (*string, bool)`

GetSuggestedResolutionOk returns a tuple with the SuggestedResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedResolution

`func (o *HyperflexHealthCheckExecutionAllOf) SetSuggestedResolution(v string)`

SetSuggestedResolution sets SuggestedResolution field to given value.

### HasSuggestedResolution

`func (o *HyperflexHealthCheckExecutionAllOf) HasSuggestedResolution() bool`

HasSuggestedResolution returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexHealthCheckExecutionAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHealthCheckExecutionAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHealthCheckExecutionAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHealthCheckDefinition

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckDefinition() HyperflexHealthCheckDefinitionRelationship`

GetHealthCheckDefinition returns the HealthCheckDefinition field if non-nil, zero value otherwise.

### GetHealthCheckDefinitionOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHealthCheckDefinitionOk() (*HyperflexHealthCheckDefinitionRelationship, bool)`

GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDefinition

`func (o *HyperflexHealthCheckExecutionAllOf) SetHealthCheckDefinition(v HyperflexHealthCheckDefinitionRelationship)`

SetHealthCheckDefinition sets HealthCheckDefinition field to given value.

### HasHealthCheckDefinition

`func (o *HyperflexHealthCheckExecutionAllOf) HasHealthCheckDefinition() bool`

HasHealthCheckDefinition returns a boolean if a field has been set.

### GetHxCluster

`func (o *HyperflexHealthCheckExecutionAllOf) GetHxCluster() HyperflexClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetHxClusterOk() (*HyperflexClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *HyperflexHealthCheckExecutionAllOf) SetHxCluster(v HyperflexClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *HyperflexHealthCheckExecutionAllOf) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexHealthCheckExecutionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexHealthCheckExecutionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexHealthCheckExecutionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexHealthCheckExecutionAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


