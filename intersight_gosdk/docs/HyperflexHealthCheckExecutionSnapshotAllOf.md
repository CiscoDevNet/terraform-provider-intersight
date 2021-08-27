# HyperflexHealthCheckExecutionSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HealthCheckExecutionSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HealthCheckExecutionSnapshot"]
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
**SuggestedResolution** | Pointer to **string** | Information detailing a suggegsted resolution for the healthcheck failure, if the check fails. | [optional] 
**HealthCheckDefinition** | Pointer to [**HyperflexHealthCheckDefinitionRelationship**](HyperflexHealthCheckDefinitionRelationship.md) |  | [optional] 
**HxCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHealthCheckExecutionSnapshotAllOf

`func NewHyperflexHealthCheckExecutionSnapshotAllOf(classId string, objectType string, ) *HyperflexHealthCheckExecutionSnapshotAllOf`

NewHyperflexHealthCheckExecutionSnapshotAllOf instantiates a new HyperflexHealthCheckExecutionSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckExecutionSnapshotAllOfWithDefaults

`func NewHyperflexHealthCheckExecutionSnapshotAllOfWithDefaults() *HyperflexHealthCheckExecutionSnapshotAllOf`

NewHyperflexHealthCheckExecutionSnapshotAllOfWithDefaults instantiates a new HyperflexHealthCheckExecutionSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCause

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCompletionTime

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetHealthCheckDetails

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDetails() string`

GetHealthCheckDetails returns the HealthCheckDetails field if non-nil, zero value otherwise.

### GetHealthCheckDetailsOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDetailsOk() (*string, bool)`

GetHealthCheckDetailsOk returns a tuple with the HealthCheckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDetails

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckDetails(v string)`

SetHealthCheckDetails sets HealthCheckDetails field to given value.

### HasHealthCheckDetails

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckDetails() bool`

HasHealthCheckDetails returns a boolean if a field has been set.

### GetHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorDetails() string`

GetHealthCheckExecutionErrorDetails returns the HealthCheckExecutionErrorDetails field if non-nil, zero value otherwise.

### GetHealthCheckExecutionErrorDetailsOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorDetailsOk() (*string, bool)`

GetHealthCheckExecutionErrorDetailsOk returns a tuple with the HealthCheckExecutionErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckExecutionErrorDetails(v string)`

SetHealthCheckExecutionErrorDetails sets HealthCheckExecutionErrorDetails field to given value.

### HasHealthCheckExecutionErrorDetails

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckExecutionErrorDetails() bool`

HasHealthCheckExecutionErrorDetails returns a boolean if a field has been set.

### GetHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorSummary() string`

GetHealthCheckExecutionErrorSummary returns the HealthCheckExecutionErrorSummary field if non-nil, zero value otherwise.

### GetHealthCheckExecutionErrorSummaryOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorSummaryOk() (*string, bool)`

GetHealthCheckExecutionErrorSummaryOk returns a tuple with the HealthCheckExecutionErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckExecutionErrorSummary(v string)`

SetHealthCheckExecutionErrorSummary sets HealthCheckExecutionErrorSummary field to given value.

### HasHealthCheckExecutionErrorSummary

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckExecutionErrorSummary() bool`

HasHealthCheckExecutionErrorSummary returns a boolean if a field has been set.

### GetHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionStatus() string`

GetHealthCheckExecutionStatus returns the HealthCheckExecutionStatus field if non-nil, zero value otherwise.

### GetHealthCheckExecutionStatusOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionStatusOk() (*string, bool)`

GetHealthCheckExecutionStatusOk returns a tuple with the HealthCheckExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckExecutionStatus(v string)`

SetHealthCheckExecutionStatus sets HealthCheckExecutionStatus field to given value.

### HasHealthCheckExecutionStatus

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckExecutionStatus() bool`

HasHealthCheckExecutionStatus returns a boolean if a field has been set.

### GetHealthCheckResult

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckResult() string`

GetHealthCheckResult returns the HealthCheckResult field if non-nil, zero value otherwise.

### GetHealthCheckResultOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckResultOk() (*string, bool)`

GetHealthCheckResultOk returns a tuple with the HealthCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckResult

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckResult(v string)`

SetHealthCheckResult sets HealthCheckResult field to given value.

### HasHealthCheckResult

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckResult() bool`

HasHealthCheckResult returns a boolean if a field has been set.

### GetHealthCheckSummary

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckSummary() string`

GetHealthCheckSummary returns the HealthCheckSummary field if non-nil, zero value otherwise.

### GetHealthCheckSummaryOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckSummaryOk() (*string, bool)`

GetHealthCheckSummaryOk returns a tuple with the HealthCheckSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckSummary

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckSummary(v string)`

SetHealthCheckSummary sets HealthCheckSummary field to given value.

### HasHealthCheckSummary

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckSummary() bool`

HasHealthCheckSummary returns a boolean if a field has been set.

### GetHxDeviceName

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxDeviceName() string`

GetHxDeviceName returns the HxDeviceName field if non-nil, zero value otherwise.

### GetHxDeviceNameOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxDeviceNameOk() (*string, bool)`

GetHxDeviceNameOk returns a tuple with the HxDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxDeviceName

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHxDeviceName(v string)`

SetHxDeviceName sets HxDeviceName field to given value.

### HasHxDeviceName

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHxDeviceName() bool`

HasHxDeviceName returns a boolean if a field has been set.

### GetSuggestedResolution

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetSuggestedResolution() string`

GetSuggestedResolution returns the SuggestedResolution field if non-nil, zero value otherwise.

### GetSuggestedResolutionOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetSuggestedResolutionOk() (*string, bool)`

GetSuggestedResolutionOk returns a tuple with the SuggestedResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedResolution

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetSuggestedResolution(v string)`

SetSuggestedResolution sets SuggestedResolution field to given value.

### HasSuggestedResolution

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasSuggestedResolution() bool`

HasSuggestedResolution returns a boolean if a field has been set.

### GetHealthCheckDefinition

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDefinition() HyperflexHealthCheckDefinitionRelationship`

GetHealthCheckDefinition returns the HealthCheckDefinition field if non-nil, zero value otherwise.

### GetHealthCheckDefinitionOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDefinitionOk() (*HyperflexHealthCheckDefinitionRelationship, bool)`

GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDefinition

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckDefinition(v HyperflexHealthCheckDefinitionRelationship)`

SetHealthCheckDefinition sets HealthCheckDefinition field to given value.

### HasHealthCheckDefinition

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckDefinition() bool`

HasHealthCheckDefinition returns a boolean if a field has been set.

### GetHxCluster

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxCluster() HyperflexClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxClusterOk() (*HyperflexClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHxCluster(v HyperflexClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflow

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


