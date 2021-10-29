# ConvergedinfraHealthCheckExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.HealthCheckExecution"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.HealthCheckExecution"]
**Error** | Pointer to **string** | Error details of a health check execution failure. | [optional] [readonly] 
**Result** | Pointer to **string** | Health check execution result. * &#x60;Unknown&#x60; - Indicates that the health check results could not be determined. * &#x60;Pass&#x60; - Indicates that the health check has passed. * &#x60;Fail&#x60; - Indicates that the health check has failed. * &#x60;Warning&#x60; - Indicates that the health check completed with a warning. * &#x60;NotApplicable&#x60; - Indicates that the health check is either unsupported, or not applicable for the pod. | [optional] [readonly] [default to "Unknown"]
**Status** | Pointer to **string** | Status of the health check execution. * &#x60;Unknown&#x60; - Indicates that the health heck execution status is unknown. This mostly happens in case where health check could not be performed due to connectivity issues. * &#x60;Succeeded&#x60; - Indicates that the health check execution has succeeded. * &#x60;Failed&#x60; - Indicates that the health check execution has failed. * &#x60;Timedout&#x60; - Indicates that the health check execution timed out before completion. | [optional] [readonly] [default to "Unknown"]
**Summary** | Pointer to **string** | A brief summary of health check results. | [optional] [readonly] 
**HealthCheckDefinition** | Pointer to [**ConvergedinfraHealthCheckDefinitionRelationship**](ConvergedinfraHealthCheckDefinitionRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraHealthCheckExecution

`func NewConvergedinfraHealthCheckExecution(classId string, objectType string, ) *ConvergedinfraHealthCheckExecution`

NewConvergedinfraHealthCheckExecution instantiates a new ConvergedinfraHealthCheckExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraHealthCheckExecutionWithDefaults

`func NewConvergedinfraHealthCheckExecutionWithDefaults() *ConvergedinfraHealthCheckExecution`

NewConvergedinfraHealthCheckExecutionWithDefaults instantiates a new ConvergedinfraHealthCheckExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraHealthCheckExecution) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraHealthCheckExecution) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraHealthCheckExecution) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraHealthCheckExecution) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraHealthCheckExecution) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraHealthCheckExecution) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetError

`func (o *ConvergedinfraHealthCheckExecution) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConvergedinfraHealthCheckExecution) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConvergedinfraHealthCheckExecution) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConvergedinfraHealthCheckExecution) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *ConvergedinfraHealthCheckExecution) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConvergedinfraHealthCheckExecution) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConvergedinfraHealthCheckExecution) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ConvergedinfraHealthCheckExecution) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *ConvergedinfraHealthCheckExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConvergedinfraHealthCheckExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConvergedinfraHealthCheckExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConvergedinfraHealthCheckExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *ConvergedinfraHealthCheckExecution) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ConvergedinfraHealthCheckExecution) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ConvergedinfraHealthCheckExecution) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ConvergedinfraHealthCheckExecution) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetHealthCheckDefinition

`func (o *ConvergedinfraHealthCheckExecution) GetHealthCheckDefinition() ConvergedinfraHealthCheckDefinitionRelationship`

GetHealthCheckDefinition returns the HealthCheckDefinition field if non-nil, zero value otherwise.

### GetHealthCheckDefinitionOk

`func (o *ConvergedinfraHealthCheckExecution) GetHealthCheckDefinitionOk() (*ConvergedinfraHealthCheckDefinitionRelationship, bool)`

GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDefinition

`func (o *ConvergedinfraHealthCheckExecution) SetHealthCheckDefinition(v ConvergedinfraHealthCheckDefinitionRelationship)`

SetHealthCheckDefinition sets HealthCheckDefinition field to given value.

### HasHealthCheckDefinition

`func (o *ConvergedinfraHealthCheckExecution) HasHealthCheckDefinition() bool`

HasHealthCheckDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


