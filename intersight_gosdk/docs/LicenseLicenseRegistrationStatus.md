# LicenseLicenseRegistrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.LicenseRegistrationStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.LicenseRegistrationStatus"]
**AccountCreationState** | Pointer to **bool** | Stores information on whether account has gone through the registration wizard. True if has not. | [optional] 
**IsNewAccount** | Pointer to **bool** | Stores information on whether account is new. This data is used for UI theme upgrade, where existing users will be shown a slightly different screen. True if new. | [optional] [readonly] 
**LicenseRegistrationState** | Pointer to **string** | Stores information on the current flow of license registration. * &#x60;RegistrationNotStarted&#x60; - The license registration state to chose between trial and registration. * &#x60;RegistrationStarted&#x60; - The license registration state during set up flow. * &#x60;RegistrationComplete&#x60; - The license registration state after completion. | [optional] [default to "RegistrationNotStarted"]
**TrialRegistrationComplete** | Pointer to **bool** | Stores information on whether trial flow has been completed. True if trial registration finish. | [optional] 
**AccountLicenseData** | Pointer to [**NullableLicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseLicenseRegistrationStatus

`func NewLicenseLicenseRegistrationStatus(classId string, objectType string, ) *LicenseLicenseRegistrationStatus`

NewLicenseLicenseRegistrationStatus instantiates a new LicenseLicenseRegistrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseRegistrationStatusWithDefaults

`func NewLicenseLicenseRegistrationStatusWithDefaults() *LicenseLicenseRegistrationStatus`

NewLicenseLicenseRegistrationStatusWithDefaults instantiates a new LicenseLicenseRegistrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseLicenseRegistrationStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseLicenseRegistrationStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseLicenseRegistrationStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseLicenseRegistrationStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseLicenseRegistrationStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseLicenseRegistrationStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountCreationState

`func (o *LicenseLicenseRegistrationStatus) GetAccountCreationState() bool`

GetAccountCreationState returns the AccountCreationState field if non-nil, zero value otherwise.

### GetAccountCreationStateOk

`func (o *LicenseLicenseRegistrationStatus) GetAccountCreationStateOk() (*bool, bool)`

GetAccountCreationStateOk returns a tuple with the AccountCreationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationState

`func (o *LicenseLicenseRegistrationStatus) SetAccountCreationState(v bool)`

SetAccountCreationState sets AccountCreationState field to given value.

### HasAccountCreationState

`func (o *LicenseLicenseRegistrationStatus) HasAccountCreationState() bool`

HasAccountCreationState returns a boolean if a field has been set.

### GetIsNewAccount

`func (o *LicenseLicenseRegistrationStatus) GetIsNewAccount() bool`

GetIsNewAccount returns the IsNewAccount field if non-nil, zero value otherwise.

### GetIsNewAccountOk

`func (o *LicenseLicenseRegistrationStatus) GetIsNewAccountOk() (*bool, bool)`

GetIsNewAccountOk returns a tuple with the IsNewAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewAccount

`func (o *LicenseLicenseRegistrationStatus) SetIsNewAccount(v bool)`

SetIsNewAccount sets IsNewAccount field to given value.

### HasIsNewAccount

`func (o *LicenseLicenseRegistrationStatus) HasIsNewAccount() bool`

HasIsNewAccount returns a boolean if a field has been set.

### GetLicenseRegistrationState

`func (o *LicenseLicenseRegistrationStatus) GetLicenseRegistrationState() string`

GetLicenseRegistrationState returns the LicenseRegistrationState field if non-nil, zero value otherwise.

### GetLicenseRegistrationStateOk

`func (o *LicenseLicenseRegistrationStatus) GetLicenseRegistrationStateOk() (*string, bool)`

GetLicenseRegistrationStateOk returns a tuple with the LicenseRegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegistrationState

`func (o *LicenseLicenseRegistrationStatus) SetLicenseRegistrationState(v string)`

SetLicenseRegistrationState sets LicenseRegistrationState field to given value.

### HasLicenseRegistrationState

`func (o *LicenseLicenseRegistrationStatus) HasLicenseRegistrationState() bool`

HasLicenseRegistrationState returns a boolean if a field has been set.

### GetTrialRegistrationComplete

`func (o *LicenseLicenseRegistrationStatus) GetTrialRegistrationComplete() bool`

GetTrialRegistrationComplete returns the TrialRegistrationComplete field if non-nil, zero value otherwise.

### GetTrialRegistrationCompleteOk

`func (o *LicenseLicenseRegistrationStatus) GetTrialRegistrationCompleteOk() (*bool, bool)`

GetTrialRegistrationCompleteOk returns a tuple with the TrialRegistrationComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialRegistrationComplete

`func (o *LicenseLicenseRegistrationStatus) SetTrialRegistrationComplete(v bool)`

SetTrialRegistrationComplete sets TrialRegistrationComplete field to given value.

### HasTrialRegistrationComplete

`func (o *LicenseLicenseRegistrationStatus) HasTrialRegistrationComplete() bool`

HasTrialRegistrationComplete returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseLicenseRegistrationStatus) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseLicenseRegistrationStatus) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseLicenseRegistrationStatus) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseLicenseRegistrationStatus) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.

### SetAccountLicenseDataNil

`func (o *LicenseLicenseRegistrationStatus) SetAccountLicenseDataNil(b bool)`

 SetAccountLicenseDataNil sets the value for AccountLicenseData to be an explicit nil

### UnsetAccountLicenseData
`func (o *LicenseLicenseRegistrationStatus) UnsetAccountLicenseData()`

UnsetAccountLicenseData ensures that no value is present for AccountLicenseData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


