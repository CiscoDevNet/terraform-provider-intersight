# LicenseIksCustomerOpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.IksCustomerOp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.IksCustomerOp"]
**ActiveAdmin** | Pointer to **bool** | The Intersight Kubernetes Service license administrative state. Set this property to &#39;true&#39; to activate the IKS license entitlements. | [optional] 
**EnableTrial** | Pointer to **bool** | Enable trial for IKS licensing. | [optional] 
**EvaluationPeriod** | Pointer to **int64** | The default Trial or Grace period the customer is entitled to. | [optional] 
**ExtraEvaluation** | Pointer to **int64** | The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. | [optional] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseIksCustomerOpAllOf

`func NewLicenseIksCustomerOpAllOf(classId string, objectType string, ) *LicenseIksCustomerOpAllOf`

NewLicenseIksCustomerOpAllOf instantiates a new LicenseIksCustomerOpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIksCustomerOpAllOfWithDefaults

`func NewLicenseIksCustomerOpAllOfWithDefaults() *LicenseIksCustomerOpAllOf`

NewLicenseIksCustomerOpAllOfWithDefaults instantiates a new LicenseIksCustomerOpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseIksCustomerOpAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseIksCustomerOpAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseIksCustomerOpAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseIksCustomerOpAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseIksCustomerOpAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseIksCustomerOpAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveAdmin

`func (o *LicenseIksCustomerOpAllOf) GetActiveAdmin() bool`

GetActiveAdmin returns the ActiveAdmin field if non-nil, zero value otherwise.

### GetActiveAdminOk

`func (o *LicenseIksCustomerOpAllOf) GetActiveAdminOk() (*bool, bool)`

GetActiveAdminOk returns a tuple with the ActiveAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdmin

`func (o *LicenseIksCustomerOpAllOf) SetActiveAdmin(v bool)`

SetActiveAdmin sets ActiveAdmin field to given value.

### HasActiveAdmin

`func (o *LicenseIksCustomerOpAllOf) HasActiveAdmin() bool`

HasActiveAdmin returns a boolean if a field has been set.

### GetEnableTrial

`func (o *LicenseIksCustomerOpAllOf) GetEnableTrial() bool`

GetEnableTrial returns the EnableTrial field if non-nil, zero value otherwise.

### GetEnableTrialOk

`func (o *LicenseIksCustomerOpAllOf) GetEnableTrialOk() (*bool, bool)`

GetEnableTrialOk returns a tuple with the EnableTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrial

`func (o *LicenseIksCustomerOpAllOf) SetEnableTrial(v bool)`

SetEnableTrial sets EnableTrial field to given value.

### HasEnableTrial

`func (o *LicenseIksCustomerOpAllOf) HasEnableTrial() bool`

HasEnableTrial returns a boolean if a field has been set.

### GetEvaluationPeriod

`func (o *LicenseIksCustomerOpAllOf) GetEvaluationPeriod() int64`

GetEvaluationPeriod returns the EvaluationPeriod field if non-nil, zero value otherwise.

### GetEvaluationPeriodOk

`func (o *LicenseIksCustomerOpAllOf) GetEvaluationPeriodOk() (*int64, bool)`

GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriod

`func (o *LicenseIksCustomerOpAllOf) SetEvaluationPeriod(v int64)`

SetEvaluationPeriod sets EvaluationPeriod field to given value.

### HasEvaluationPeriod

`func (o *LicenseIksCustomerOpAllOf) HasEvaluationPeriod() bool`

HasEvaluationPeriod returns a boolean if a field has been set.

### GetExtraEvaluation

`func (o *LicenseIksCustomerOpAllOf) GetExtraEvaluation() int64`

GetExtraEvaluation returns the ExtraEvaluation field if non-nil, zero value otherwise.

### GetExtraEvaluationOk

`func (o *LicenseIksCustomerOpAllOf) GetExtraEvaluationOk() (*int64, bool)`

GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEvaluation

`func (o *LicenseIksCustomerOpAllOf) SetExtraEvaluation(v int64)`

SetExtraEvaluation sets ExtraEvaluation field to given value.

### HasExtraEvaluation

`func (o *LicenseIksCustomerOpAllOf) HasExtraEvaluation() bool`

HasExtraEvaluation returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseIksCustomerOpAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseIksCustomerOpAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseIksCustomerOpAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseIksCustomerOpAllOf) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


