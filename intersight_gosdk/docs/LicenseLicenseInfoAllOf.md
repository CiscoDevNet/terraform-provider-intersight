# LicenseLicenseInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAdmin** | Pointer to **bool** | The license administrative state. Set this property to &#39;true&#39; to activate the license entitlements. | [optional] [readonly] 
**DaysLeft** | Pointer to **int64** | The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The date and time when the trial period expires. The value of the &#39;endTime&#39; property is set when the account enters the TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**EnforceMode** | Pointer to **string** | The entitlement mode reported by Cisco Smart Software Manager. | [optional] [readonly] 
**ErrorDesc** | Pointer to **string** | The detailed error message when there is any error related to this licensing entitlement. | [optional] [readonly] 
**EvaluationPeriod** | Pointer to **int64** | The default Trial or Grace period customer is entitled to. | [optional] 
**ExtraEvaluation** | Pointer to **int64** | The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. | [optional] 
**LicenseCount** | Pointer to **int64** | The total number of devices claimed in the Intersight account. | [optional] [readonly] 
**LicenseState** | Pointer to **string** | The license state defined by Intersight. The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired. * &#x60;NotLicensed&#x60; - The license token is neither activated nor registered. * &#x60;GraceExpired&#x60; - The license grace period has expired. * &#x60;TrialPeriod&#x60; - The 90 days of trial period. * &#x60;OutOfCompliance&#x60; - The license is out of compliance. * &#x60;Compliance&#x60; - The license is in compliance. * &#x60;TrialExpired&#x60; - The trial period of 90 days has expired. | [optional] [readonly] [default to "NotLicensed"]
**LicenseType** | Pointer to **string** | The name of the Intersight license entitlement. For example, this property may be set to &#39;Essential&#39;. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. | [optional] [readonly] [default to "Base"]
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**TrialAdmin** | Pointer to **bool** | The administrative state of the trial license. When the LicenseState is set to &#39;NotLicensed&#39;, &#39;trialAdmin&#39; can be set to true to start the trial period, i.e. licenseState is set to be TrialPeriod. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](license.AccountLicenseData.Relationship.md) |  | [optional] 

## Methods

### NewLicenseLicenseInfoAllOf

`func NewLicenseLicenseInfoAllOf() *LicenseLicenseInfoAllOf`

NewLicenseLicenseInfoAllOf instantiates a new LicenseLicenseInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseInfoAllOfWithDefaults

`func NewLicenseLicenseInfoAllOfWithDefaults() *LicenseLicenseInfoAllOf`

NewLicenseLicenseInfoAllOfWithDefaults instantiates a new LicenseLicenseInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAdmin

`func (o *LicenseLicenseInfoAllOf) GetActiveAdmin() bool`

GetActiveAdmin returns the ActiveAdmin field if non-nil, zero value otherwise.

### GetActiveAdminOk

`func (o *LicenseLicenseInfoAllOf) GetActiveAdminOk() (*bool, bool)`

GetActiveAdminOk returns a tuple with the ActiveAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdmin

`func (o *LicenseLicenseInfoAllOf) SetActiveAdmin(v bool)`

SetActiveAdmin sets ActiveAdmin field to given value.

### HasActiveAdmin

`func (o *LicenseLicenseInfoAllOf) HasActiveAdmin() bool`

HasActiveAdmin returns a boolean if a field has been set.

### GetDaysLeft

`func (o *LicenseLicenseInfoAllOf) GetDaysLeft() int64`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *LicenseLicenseInfoAllOf) GetDaysLeftOk() (*int64, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *LicenseLicenseInfoAllOf) SetDaysLeft(v int64)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *LicenseLicenseInfoAllOf) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetEndTime

`func (o *LicenseLicenseInfoAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LicenseLicenseInfoAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LicenseLicenseInfoAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LicenseLicenseInfoAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnforceMode

`func (o *LicenseLicenseInfoAllOf) GetEnforceMode() string`

GetEnforceMode returns the EnforceMode field if non-nil, zero value otherwise.

### GetEnforceModeOk

`func (o *LicenseLicenseInfoAllOf) GetEnforceModeOk() (*string, bool)`

GetEnforceModeOk returns a tuple with the EnforceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMode

`func (o *LicenseLicenseInfoAllOf) SetEnforceMode(v string)`

SetEnforceMode sets EnforceMode field to given value.

### HasEnforceMode

`func (o *LicenseLicenseInfoAllOf) HasEnforceMode() bool`

HasEnforceMode returns a boolean if a field has been set.

### GetErrorDesc

`func (o *LicenseLicenseInfoAllOf) GetErrorDesc() string`

GetErrorDesc returns the ErrorDesc field if non-nil, zero value otherwise.

### GetErrorDescOk

`func (o *LicenseLicenseInfoAllOf) GetErrorDescOk() (*string, bool)`

GetErrorDescOk returns a tuple with the ErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDesc

`func (o *LicenseLicenseInfoAllOf) SetErrorDesc(v string)`

SetErrorDesc sets ErrorDesc field to given value.

### HasErrorDesc

`func (o *LicenseLicenseInfoAllOf) HasErrorDesc() bool`

HasErrorDesc returns a boolean if a field has been set.

### GetEvaluationPeriod

`func (o *LicenseLicenseInfoAllOf) GetEvaluationPeriod() int64`

GetEvaluationPeriod returns the EvaluationPeriod field if non-nil, zero value otherwise.

### GetEvaluationPeriodOk

`func (o *LicenseLicenseInfoAllOf) GetEvaluationPeriodOk() (*int64, bool)`

GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriod

`func (o *LicenseLicenseInfoAllOf) SetEvaluationPeriod(v int64)`

SetEvaluationPeriod sets EvaluationPeriod field to given value.

### HasEvaluationPeriod

`func (o *LicenseLicenseInfoAllOf) HasEvaluationPeriod() bool`

HasEvaluationPeriod returns a boolean if a field has been set.

### GetExtraEvaluation

`func (o *LicenseLicenseInfoAllOf) GetExtraEvaluation() int64`

GetExtraEvaluation returns the ExtraEvaluation field if non-nil, zero value otherwise.

### GetExtraEvaluationOk

`func (o *LicenseLicenseInfoAllOf) GetExtraEvaluationOk() (*int64, bool)`

GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEvaluation

`func (o *LicenseLicenseInfoAllOf) SetExtraEvaluation(v int64)`

SetExtraEvaluation sets ExtraEvaluation field to given value.

### HasExtraEvaluation

`func (o *LicenseLicenseInfoAllOf) HasExtraEvaluation() bool`

HasExtraEvaluation returns a boolean if a field has been set.

### GetLicenseCount

`func (o *LicenseLicenseInfoAllOf) GetLicenseCount() int64`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *LicenseLicenseInfoAllOf) GetLicenseCountOk() (*int64, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *LicenseLicenseInfoAllOf) SetLicenseCount(v int64)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *LicenseLicenseInfoAllOf) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetLicenseState

`func (o *LicenseLicenseInfoAllOf) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *LicenseLicenseInfoAllOf) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *LicenseLicenseInfoAllOf) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *LicenseLicenseInfoAllOf) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLicenseType

`func (o *LicenseLicenseInfoAllOf) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicenseLicenseInfoAllOf) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicenseLicenseInfoAllOf) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicenseLicenseInfoAllOf) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetStartTime

`func (o *LicenseLicenseInfoAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LicenseLicenseInfoAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LicenseLicenseInfoAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *LicenseLicenseInfoAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTrialAdmin

`func (o *LicenseLicenseInfoAllOf) GetTrialAdmin() bool`

GetTrialAdmin returns the TrialAdmin field if non-nil, zero value otherwise.

### GetTrialAdminOk

`func (o *LicenseLicenseInfoAllOf) GetTrialAdminOk() (*bool, bool)`

GetTrialAdminOk returns a tuple with the TrialAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAdmin

`func (o *LicenseLicenseInfoAllOf) SetTrialAdmin(v bool)`

SetTrialAdmin sets TrialAdmin field to given value.

### HasTrialAdmin

`func (o *LicenseLicenseInfoAllOf) HasTrialAdmin() bool`

HasTrialAdmin returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseLicenseInfoAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseLicenseInfoAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseLicenseInfoAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseLicenseInfoAllOf) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


