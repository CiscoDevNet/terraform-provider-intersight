# HyperflexHxLicenseAuthorizationDetailsDt

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

### NewHyperflexHxLicenseAuthorizationDetailsDt

`func NewHyperflexHxLicenseAuthorizationDetailsDt(classId string, objectType string, ) *HyperflexHxLicenseAuthorizationDetailsDt`

NewHyperflexHxLicenseAuthorizationDetailsDt instantiates a new HyperflexHxLicenseAuthorizationDetailsDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxLicenseAuthorizationDetailsDtWithDefaults

`func NewHyperflexHxLicenseAuthorizationDetailsDtWithDefaults() *HyperflexHxLicenseAuthorizationDetailsDt`

NewHyperflexHxLicenseAuthorizationDetailsDtWithDefaults instantiates a new HyperflexHxLicenseAuthorizationDetailsDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommunicationDeadlineDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetCommunicationDeadlineDate() string`

GetCommunicationDeadlineDate returns the CommunicationDeadlineDate field if non-nil, zero value otherwise.

### GetCommunicationDeadlineDateOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetCommunicationDeadlineDateOk() (*string, bool)`

GetCommunicationDeadlineDateOk returns a tuple with the CommunicationDeadlineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDeadlineDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetCommunicationDeadlineDate(v string)`

SetCommunicationDeadlineDate sets CommunicationDeadlineDate field to given value.

### HasCommunicationDeadlineDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) HasCommunicationDeadlineDate() bool`

HasCommunicationDeadlineDate returns a boolean if a field has been set.

### GetEvaluationPeriodExpiredAt

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetEvaluationPeriodExpiredAt() string`

GetEvaluationPeriodExpiredAt returns the EvaluationPeriodExpiredAt field if non-nil, zero value otherwise.

### GetEvaluationPeriodExpiredAtOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetEvaluationPeriodExpiredAtOk() (*string, bool)`

GetEvaluationPeriodExpiredAtOk returns a tuple with the EvaluationPeriodExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriodExpiredAt

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetEvaluationPeriodExpiredAt(v string)`

SetEvaluationPeriodExpiredAt sets EvaluationPeriodExpiredAt field to given value.

### HasEvaluationPeriodExpiredAt

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) HasEvaluationPeriodExpiredAt() bool`

HasEvaluationPeriodExpiredAt returns a boolean if a field has been set.

### GetEvaluationPeriodRemaining

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetEvaluationPeriodRemaining() string`

GetEvaluationPeriodRemaining returns the EvaluationPeriodRemaining field if non-nil, zero value otherwise.

### GetEvaluationPeriodRemainingOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetEvaluationPeriodRemainingOk() (*string, bool)`

GetEvaluationPeriodRemainingOk returns a tuple with the EvaluationPeriodRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriodRemaining

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetEvaluationPeriodRemaining(v string)`

SetEvaluationPeriodRemaining sets EvaluationPeriodRemaining field to given value.

### HasEvaluationPeriodRemaining

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) HasEvaluationPeriodRemaining() bool`

HasEvaluationPeriodRemaining returns a boolean if a field has been set.

### GetLastCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetLastCommunicationAttemptDate() string`

GetLastCommunicationAttemptDate returns the LastCommunicationAttemptDate field if non-nil, zero value otherwise.

### GetLastCommunicationAttemptDateOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetLastCommunicationAttemptDateOk() (*string, bool)`

GetLastCommunicationAttemptDateOk returns a tuple with the LastCommunicationAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetLastCommunicationAttemptDate(v string)`

SetLastCommunicationAttemptDate sets LastCommunicationAttemptDate field to given value.

### HasLastCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) HasLastCommunicationAttemptDate() bool`

HasLastCommunicationAttemptDate returns a boolean if a field has been set.

### GetNextCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetNextCommunicationAttemptDate() string`

GetNextCommunicationAttemptDate returns the NextCommunicationAttemptDate field if non-nil, zero value otherwise.

### GetNextCommunicationAttemptDateOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetNextCommunicationAttemptDateOk() (*string, bool)`

GetNextCommunicationAttemptDateOk returns a tuple with the NextCommunicationAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetNextCommunicationAttemptDate(v string)`

SetNextCommunicationAttemptDate sets NextCommunicationAttemptDate field to given value.

### HasNextCommunicationAttemptDate

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) HasNextCommunicationAttemptDate() bool`

HasNextCommunicationAttemptDate returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxLicenseAuthorizationDetailsDt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


