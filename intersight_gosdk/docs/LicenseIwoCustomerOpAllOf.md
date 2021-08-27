# LicenseIwoCustomerOpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.IwoCustomerOp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.IwoCustomerOp"]
**ActiveAdmin** | Pointer to **bool** | The workload optimizer license administrative state. Set this property to &#39;true&#39; to activate the workload optimizer license entitlements. | [optional] 
**ActiveLicenseType** | Pointer to **string** | Active workload optimizer license tier set by user. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. | [optional] [default to "Base"]
**EnableTrial** | Pointer to **bool** | Enable trial for Intersight licensing. | [optional] 
**EvaluationPeriod** | Pointer to **int64** | The default Trial or Grace period customer is entitled to. | [optional] 
**ExtraEvaluation** | Pointer to **int64** | The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. | [optional] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseIwoCustomerOpAllOf

`func NewLicenseIwoCustomerOpAllOf(classId string, objectType string, ) *LicenseIwoCustomerOpAllOf`

NewLicenseIwoCustomerOpAllOf instantiates a new LicenseIwoCustomerOpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIwoCustomerOpAllOfWithDefaults

`func NewLicenseIwoCustomerOpAllOfWithDefaults() *LicenseIwoCustomerOpAllOf`

NewLicenseIwoCustomerOpAllOfWithDefaults instantiates a new LicenseIwoCustomerOpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseIwoCustomerOpAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseIwoCustomerOpAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseIwoCustomerOpAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseIwoCustomerOpAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseIwoCustomerOpAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseIwoCustomerOpAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveAdmin

`func (o *LicenseIwoCustomerOpAllOf) GetActiveAdmin() bool`

GetActiveAdmin returns the ActiveAdmin field if non-nil, zero value otherwise.

### GetActiveAdminOk

`func (o *LicenseIwoCustomerOpAllOf) GetActiveAdminOk() (*bool, bool)`

GetActiveAdminOk returns a tuple with the ActiveAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdmin

`func (o *LicenseIwoCustomerOpAllOf) SetActiveAdmin(v bool)`

SetActiveAdmin sets ActiveAdmin field to given value.

### HasActiveAdmin

`func (o *LicenseIwoCustomerOpAllOf) HasActiveAdmin() bool`

HasActiveAdmin returns a boolean if a field has been set.

### GetActiveLicenseType

`func (o *LicenseIwoCustomerOpAllOf) GetActiveLicenseType() string`

GetActiveLicenseType returns the ActiveLicenseType field if non-nil, zero value otherwise.

### GetActiveLicenseTypeOk

`func (o *LicenseIwoCustomerOpAllOf) GetActiveLicenseTypeOk() (*string, bool)`

GetActiveLicenseTypeOk returns a tuple with the ActiveLicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLicenseType

`func (o *LicenseIwoCustomerOpAllOf) SetActiveLicenseType(v string)`

SetActiveLicenseType sets ActiveLicenseType field to given value.

### HasActiveLicenseType

`func (o *LicenseIwoCustomerOpAllOf) HasActiveLicenseType() bool`

HasActiveLicenseType returns a boolean if a field has been set.

### GetEnableTrial

`func (o *LicenseIwoCustomerOpAllOf) GetEnableTrial() bool`

GetEnableTrial returns the EnableTrial field if non-nil, zero value otherwise.

### GetEnableTrialOk

`func (o *LicenseIwoCustomerOpAllOf) GetEnableTrialOk() (*bool, bool)`

GetEnableTrialOk returns a tuple with the EnableTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrial

`func (o *LicenseIwoCustomerOpAllOf) SetEnableTrial(v bool)`

SetEnableTrial sets EnableTrial field to given value.

### HasEnableTrial

`func (o *LicenseIwoCustomerOpAllOf) HasEnableTrial() bool`

HasEnableTrial returns a boolean if a field has been set.

### GetEvaluationPeriod

`func (o *LicenseIwoCustomerOpAllOf) GetEvaluationPeriod() int64`

GetEvaluationPeriod returns the EvaluationPeriod field if non-nil, zero value otherwise.

### GetEvaluationPeriodOk

`func (o *LicenseIwoCustomerOpAllOf) GetEvaluationPeriodOk() (*int64, bool)`

GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriod

`func (o *LicenseIwoCustomerOpAllOf) SetEvaluationPeriod(v int64)`

SetEvaluationPeriod sets EvaluationPeriod field to given value.

### HasEvaluationPeriod

`func (o *LicenseIwoCustomerOpAllOf) HasEvaluationPeriod() bool`

HasEvaluationPeriod returns a boolean if a field has been set.

### GetExtraEvaluation

`func (o *LicenseIwoCustomerOpAllOf) GetExtraEvaluation() int64`

GetExtraEvaluation returns the ExtraEvaluation field if non-nil, zero value otherwise.

### GetExtraEvaluationOk

`func (o *LicenseIwoCustomerOpAllOf) GetExtraEvaluationOk() (*int64, bool)`

GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEvaluation

`func (o *LicenseIwoCustomerOpAllOf) SetExtraEvaluation(v int64)`

SetExtraEvaluation sets ExtraEvaluation field to given value.

### HasExtraEvaluation

`func (o *LicenseIwoCustomerOpAllOf) HasExtraEvaluation() bool`

HasExtraEvaluation returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseIwoCustomerOpAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseIwoCustomerOpAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseIwoCustomerOpAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseIwoCustomerOpAllOf) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


