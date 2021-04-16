# HyperflexHxLicenseAuthorizationDetailsDtAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxLicenseAuthorizationDetailsDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxLicenseAuthorizationDetailsDt"]
**CommunicationDeadlineDate** | Pointer to **string** | Deadline date for next communication | [optional] [readonly] 
**EvaluationPeriodExpiredAt** | Pointer to **string** | Evaluation period expired date | [optional] [readonly] 
**EvaluationPeriodRemaining** | Pointer to **string** | Remaining evaluation period | [optional] [readonly] 
**LastCommunicationAttemptDate** | Pointer to **string** | Last Communication Attempt Date | [optional] [readonly] 
**NextCommunicationAttemptDate** | Pointer to **string** | Timestamp of Next Communication Attempt | [optional] [readonly] 
**Status** | Pointer to **string** | License Authorization Status | [optional] [readonly] 

## Methods

### NewHyperflexHxLicenseAuthorizationDetailsDtAllOf

`func NewHyperflexHxLicenseAuthorizationDetailsDtAllOf(classId string, objectType string, ) *HyperflexHxLicenseAuthorizationDetailsDtAllOf`

NewHyperflexHxLicenseAuthorizationDetailsDtAllOf instantiates a new HyperflexHxLicenseAuthorizationDetailsDtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxLicenseAuthorizationDetailsDtAllOfWithDefaults

`func NewHyperflexHxLicenseAuthorizationDetailsDtAllOfWithDefaults() *HyperflexHxLicenseAuthorizationDetailsDtAllOf`

NewHyperflexHxLicenseAuthorizationDetailsDtAllOfWithDefaults instantiates a new HyperflexHxLicenseAuthorizationDetailsDtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommunicationDeadlineDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetCommunicationDeadlineDate() string`

GetCommunicationDeadlineDate returns the CommunicationDeadlineDate field if non-nil, zero value otherwise.

### GetCommunicationDeadlineDateOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetCommunicationDeadlineDateOk() (*string, bool)`

GetCommunicationDeadlineDateOk returns a tuple with the CommunicationDeadlineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDeadlineDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetCommunicationDeadlineDate(v string)`

SetCommunicationDeadlineDate sets CommunicationDeadlineDate field to given value.

### HasCommunicationDeadlineDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasCommunicationDeadlineDate() bool`

HasCommunicationDeadlineDate returns a boolean if a field has been set.

### GetEvaluationPeriodExpiredAt

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodExpiredAt() string`

GetEvaluationPeriodExpiredAt returns the EvaluationPeriodExpiredAt field if non-nil, zero value otherwise.

### GetEvaluationPeriodExpiredAtOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodExpiredAtOk() (*string, bool)`

GetEvaluationPeriodExpiredAtOk returns a tuple with the EvaluationPeriodExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriodExpiredAt

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetEvaluationPeriodExpiredAt(v string)`

SetEvaluationPeriodExpiredAt sets EvaluationPeriodExpiredAt field to given value.

### HasEvaluationPeriodExpiredAt

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasEvaluationPeriodExpiredAt() bool`

HasEvaluationPeriodExpiredAt returns a boolean if a field has been set.

### GetEvaluationPeriodRemaining

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodRemaining() string`

GetEvaluationPeriodRemaining returns the EvaluationPeriodRemaining field if non-nil, zero value otherwise.

### GetEvaluationPeriodRemainingOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodRemainingOk() (*string, bool)`

GetEvaluationPeriodRemainingOk returns a tuple with the EvaluationPeriodRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriodRemaining

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetEvaluationPeriodRemaining(v string)`

SetEvaluationPeriodRemaining sets EvaluationPeriodRemaining field to given value.

### HasEvaluationPeriodRemaining

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasEvaluationPeriodRemaining() bool`

HasEvaluationPeriodRemaining returns a boolean if a field has been set.

### GetLastCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetLastCommunicationAttemptDate() string`

GetLastCommunicationAttemptDate returns the LastCommunicationAttemptDate field if non-nil, zero value otherwise.

### GetLastCommunicationAttemptDateOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetLastCommunicationAttemptDateOk() (*string, bool)`

GetLastCommunicationAttemptDateOk returns a tuple with the LastCommunicationAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetLastCommunicationAttemptDate(v string)`

SetLastCommunicationAttemptDate sets LastCommunicationAttemptDate field to given value.

### HasLastCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasLastCommunicationAttemptDate() bool`

HasLastCommunicationAttemptDate returns a boolean if a field has been set.

### GetNextCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetNextCommunicationAttemptDate() string`

GetNextCommunicationAttemptDate returns the NextCommunicationAttemptDate field if non-nil, zero value otherwise.

### GetNextCommunicationAttemptDateOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetNextCommunicationAttemptDateOk() (*string, bool)`

GetNextCommunicationAttemptDateOk returns a tuple with the NextCommunicationAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetNextCommunicationAttemptDate(v string)`

SetNextCommunicationAttemptDate sets NextCommunicationAttemptDate field to given value.

### HasNextCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasNextCommunicationAttemptDate() bool`

HasNextCommunicationAttemptDate returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


