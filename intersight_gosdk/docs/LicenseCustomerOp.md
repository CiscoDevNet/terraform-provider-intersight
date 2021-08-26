# LicenseCustomerOp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.CustomerOp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.CustomerOp"]
**ActiveAdmin** | Pointer to **bool** | The license administrative state. Set this property to &#39;true&#39; to activate the license entitlements. | [optional] 
**AllDevicesToDefaultTier** | Pointer to **bool** | Move all licensed devices to default license tier. | [optional] 
**DeregisterDevice** | Pointer to **bool** | Trigger de-registration/disable. | [optional] 
**EnableTrial** | Pointer to **bool** | Enable trial for Intersight licensing. | [optional] 
**EvaluationPeriod** | Pointer to **int64** | The default Trial or Grace period customer is entitled to. | [optional] 
**ExtraEvaluation** | Pointer to **int64** | The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. | [optional] 
**RenewAuthorization** | Pointer to **bool** | Trigger renew authorization. | [optional] 
**RenewIdCertificate** | Pointer to **bool** | Trigger renew registration. | [optional] 
**ShowAgentTechSupport** | Pointer to **bool** | Trigger show tech support feature. | [optional] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseCustomerOp

`func NewLicenseCustomerOp(classId string, objectType string, ) *LicenseCustomerOp`

NewLicenseCustomerOp instantiates a new LicenseCustomerOp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseCustomerOpWithDefaults

`func NewLicenseCustomerOpWithDefaults() *LicenseCustomerOp`

NewLicenseCustomerOpWithDefaults instantiates a new LicenseCustomerOp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseCustomerOp) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseCustomerOp) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseCustomerOp) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseCustomerOp) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseCustomerOp) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseCustomerOp) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveAdmin

`func (o *LicenseCustomerOp) GetActiveAdmin() bool`

GetActiveAdmin returns the ActiveAdmin field if non-nil, zero value otherwise.

### GetActiveAdminOk

`func (o *LicenseCustomerOp) GetActiveAdminOk() (*bool, bool)`

GetActiveAdminOk returns a tuple with the ActiveAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdmin

`func (o *LicenseCustomerOp) SetActiveAdmin(v bool)`

SetActiveAdmin sets ActiveAdmin field to given value.

### HasActiveAdmin

`func (o *LicenseCustomerOp) HasActiveAdmin() bool`

HasActiveAdmin returns a boolean if a field has been set.

### GetAllDevicesToDefaultTier

`func (o *LicenseCustomerOp) GetAllDevicesToDefaultTier() bool`

GetAllDevicesToDefaultTier returns the AllDevicesToDefaultTier field if non-nil, zero value otherwise.

### GetAllDevicesToDefaultTierOk

`func (o *LicenseCustomerOp) GetAllDevicesToDefaultTierOk() (*bool, bool)`

GetAllDevicesToDefaultTierOk returns a tuple with the AllDevicesToDefaultTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDevicesToDefaultTier

`func (o *LicenseCustomerOp) SetAllDevicesToDefaultTier(v bool)`

SetAllDevicesToDefaultTier sets AllDevicesToDefaultTier field to given value.

### HasAllDevicesToDefaultTier

`func (o *LicenseCustomerOp) HasAllDevicesToDefaultTier() bool`

HasAllDevicesToDefaultTier returns a boolean if a field has been set.

### GetDeregisterDevice

`func (o *LicenseCustomerOp) GetDeregisterDevice() bool`

GetDeregisterDevice returns the DeregisterDevice field if non-nil, zero value otherwise.

### GetDeregisterDeviceOk

`func (o *LicenseCustomerOp) GetDeregisterDeviceOk() (*bool, bool)`

GetDeregisterDeviceOk returns a tuple with the DeregisterDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregisterDevice

`func (o *LicenseCustomerOp) SetDeregisterDevice(v bool)`

SetDeregisterDevice sets DeregisterDevice field to given value.

### HasDeregisterDevice

`func (o *LicenseCustomerOp) HasDeregisterDevice() bool`

HasDeregisterDevice returns a boolean if a field has been set.

### GetEnableTrial

`func (o *LicenseCustomerOp) GetEnableTrial() bool`

GetEnableTrial returns the EnableTrial field if non-nil, zero value otherwise.

### GetEnableTrialOk

`func (o *LicenseCustomerOp) GetEnableTrialOk() (*bool, bool)`

GetEnableTrialOk returns a tuple with the EnableTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrial

`func (o *LicenseCustomerOp) SetEnableTrial(v bool)`

SetEnableTrial sets EnableTrial field to given value.

### HasEnableTrial

`func (o *LicenseCustomerOp) HasEnableTrial() bool`

HasEnableTrial returns a boolean if a field has been set.

### GetEvaluationPeriod

`func (o *LicenseCustomerOp) GetEvaluationPeriod() int64`

GetEvaluationPeriod returns the EvaluationPeriod field if non-nil, zero value otherwise.

### GetEvaluationPeriodOk

`func (o *LicenseCustomerOp) GetEvaluationPeriodOk() (*int64, bool)`

GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriod

`func (o *LicenseCustomerOp) SetEvaluationPeriod(v int64)`

SetEvaluationPeriod sets EvaluationPeriod field to given value.

### HasEvaluationPeriod

`func (o *LicenseCustomerOp) HasEvaluationPeriod() bool`

HasEvaluationPeriod returns a boolean if a field has been set.

### GetExtraEvaluation

`func (o *LicenseCustomerOp) GetExtraEvaluation() int64`

GetExtraEvaluation returns the ExtraEvaluation field if non-nil, zero value otherwise.

### GetExtraEvaluationOk

`func (o *LicenseCustomerOp) GetExtraEvaluationOk() (*int64, bool)`

GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEvaluation

`func (o *LicenseCustomerOp) SetExtraEvaluation(v int64)`

SetExtraEvaluation sets ExtraEvaluation field to given value.

### HasExtraEvaluation

`func (o *LicenseCustomerOp) HasExtraEvaluation() bool`

HasExtraEvaluation returns a boolean if a field has been set.

### GetRenewAuthorization

`func (o *LicenseCustomerOp) GetRenewAuthorization() bool`

GetRenewAuthorization returns the RenewAuthorization field if non-nil, zero value otherwise.

### GetRenewAuthorizationOk

`func (o *LicenseCustomerOp) GetRenewAuthorizationOk() (*bool, bool)`

GetRenewAuthorizationOk returns a tuple with the RenewAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewAuthorization

`func (o *LicenseCustomerOp) SetRenewAuthorization(v bool)`

SetRenewAuthorization sets RenewAuthorization field to given value.

### HasRenewAuthorization

`func (o *LicenseCustomerOp) HasRenewAuthorization() bool`

HasRenewAuthorization returns a boolean if a field has been set.

### GetRenewIdCertificate

`func (o *LicenseCustomerOp) GetRenewIdCertificate() bool`

GetRenewIdCertificate returns the RenewIdCertificate field if non-nil, zero value otherwise.

### GetRenewIdCertificateOk

`func (o *LicenseCustomerOp) GetRenewIdCertificateOk() (*bool, bool)`

GetRenewIdCertificateOk returns a tuple with the RenewIdCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewIdCertificate

`func (o *LicenseCustomerOp) SetRenewIdCertificate(v bool)`

SetRenewIdCertificate sets RenewIdCertificate field to given value.

### HasRenewIdCertificate

`func (o *LicenseCustomerOp) HasRenewIdCertificate() bool`

HasRenewIdCertificate returns a boolean if a field has been set.

### GetShowAgentTechSupport

`func (o *LicenseCustomerOp) GetShowAgentTechSupport() bool`

GetShowAgentTechSupport returns the ShowAgentTechSupport field if non-nil, zero value otherwise.

### GetShowAgentTechSupportOk

`func (o *LicenseCustomerOp) GetShowAgentTechSupportOk() (*bool, bool)`

GetShowAgentTechSupportOk returns a tuple with the ShowAgentTechSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAgentTechSupport

`func (o *LicenseCustomerOp) SetShowAgentTechSupport(v bool)`

SetShowAgentTechSupport sets ShowAgentTechSupport field to given value.

### HasShowAgentTechSupport

`func (o *LicenseCustomerOp) HasShowAgentTechSupport() bool`

HasShowAgentTechSupport returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseCustomerOp) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseCustomerOp) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseCustomerOp) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseCustomerOp) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


