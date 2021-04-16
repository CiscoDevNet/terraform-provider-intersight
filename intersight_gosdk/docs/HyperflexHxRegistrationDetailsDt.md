# HyperflexHxRegistrationDetailsDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxRegistrationDetailsDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxRegistrationDetailsDt"]
**InitialRegistrationDate** | Pointer to **string** | Initial Registration Date | [optional] [readonly] 
**LastRenewalAttemptDate** | Pointer to **string** | Last License Renewal attempted Date | [optional] [readonly] 
**NextRenewalAttemptDate** | Pointer to **string** | Next Attempt Date for License Renewal | [optional] [readonly] 
**OutOfComplianceStartDate** | Pointer to **string** | Out of compliance start date | [optional] [readonly] 
**RegistrationExpireDate** | Pointer to **string** | Date when the registration will expire | [optional] [readonly] 
**RegistrationFailedReason** | Pointer to **string** | License registration success or failure reason | [optional] [readonly] 
**SmartAccount** | Pointer to **string** | Smart Account mapped to cluster | [optional] [readonly] 
**Status** | Pointer to **string** | Registration Status | [optional] [readonly] 
**VirtualAccount** | Pointer to **string** | Virtual Account mapped to cluster | [optional] [readonly] 

## Methods

### NewHyperflexHxRegistrationDetailsDt

`func NewHyperflexHxRegistrationDetailsDt(classId string, objectType string, ) *HyperflexHxRegistrationDetailsDt`

NewHyperflexHxRegistrationDetailsDt instantiates a new HyperflexHxRegistrationDetailsDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxRegistrationDetailsDtWithDefaults

`func NewHyperflexHxRegistrationDetailsDtWithDefaults() *HyperflexHxRegistrationDetailsDt`

NewHyperflexHxRegistrationDetailsDtWithDefaults instantiates a new HyperflexHxRegistrationDetailsDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxRegistrationDetailsDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxRegistrationDetailsDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxRegistrationDetailsDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxRegistrationDetailsDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxRegistrationDetailsDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxRegistrationDetailsDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInitialRegistrationDate

`func (o *HyperflexHxRegistrationDetailsDt) GetInitialRegistrationDate() string`

GetInitialRegistrationDate returns the InitialRegistrationDate field if non-nil, zero value otherwise.

### GetInitialRegistrationDateOk

`func (o *HyperflexHxRegistrationDetailsDt) GetInitialRegistrationDateOk() (*string, bool)`

GetInitialRegistrationDateOk returns a tuple with the InitialRegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialRegistrationDate

`func (o *HyperflexHxRegistrationDetailsDt) SetInitialRegistrationDate(v string)`

SetInitialRegistrationDate sets InitialRegistrationDate field to given value.

### HasInitialRegistrationDate

`func (o *HyperflexHxRegistrationDetailsDt) HasInitialRegistrationDate() bool`

HasInitialRegistrationDate returns a boolean if a field has been set.

### GetLastRenewalAttemptDate

`func (o *HyperflexHxRegistrationDetailsDt) GetLastRenewalAttemptDate() string`

GetLastRenewalAttemptDate returns the LastRenewalAttemptDate field if non-nil, zero value otherwise.

### GetLastRenewalAttemptDateOk

`func (o *HyperflexHxRegistrationDetailsDt) GetLastRenewalAttemptDateOk() (*string, bool)`

GetLastRenewalAttemptDateOk returns a tuple with the LastRenewalAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRenewalAttemptDate

`func (o *HyperflexHxRegistrationDetailsDt) SetLastRenewalAttemptDate(v string)`

SetLastRenewalAttemptDate sets LastRenewalAttemptDate field to given value.

### HasLastRenewalAttemptDate

`func (o *HyperflexHxRegistrationDetailsDt) HasLastRenewalAttemptDate() bool`

HasLastRenewalAttemptDate returns a boolean if a field has been set.

### GetNextRenewalAttemptDate

`func (o *HyperflexHxRegistrationDetailsDt) GetNextRenewalAttemptDate() string`

GetNextRenewalAttemptDate returns the NextRenewalAttemptDate field if non-nil, zero value otherwise.

### GetNextRenewalAttemptDateOk

`func (o *HyperflexHxRegistrationDetailsDt) GetNextRenewalAttemptDateOk() (*string, bool)`

GetNextRenewalAttemptDateOk returns a tuple with the NextRenewalAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRenewalAttemptDate

`func (o *HyperflexHxRegistrationDetailsDt) SetNextRenewalAttemptDate(v string)`

SetNextRenewalAttemptDate sets NextRenewalAttemptDate field to given value.

### HasNextRenewalAttemptDate

`func (o *HyperflexHxRegistrationDetailsDt) HasNextRenewalAttemptDate() bool`

HasNextRenewalAttemptDate returns a boolean if a field has been set.

### GetOutOfComplianceStartDate

`func (o *HyperflexHxRegistrationDetailsDt) GetOutOfComplianceStartDate() string`

GetOutOfComplianceStartDate returns the OutOfComplianceStartDate field if non-nil, zero value otherwise.

### GetOutOfComplianceStartDateOk

`func (o *HyperflexHxRegistrationDetailsDt) GetOutOfComplianceStartDateOk() (*string, bool)`

GetOutOfComplianceStartDateOk returns a tuple with the OutOfComplianceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfComplianceStartDate

`func (o *HyperflexHxRegistrationDetailsDt) SetOutOfComplianceStartDate(v string)`

SetOutOfComplianceStartDate sets OutOfComplianceStartDate field to given value.

### HasOutOfComplianceStartDate

`func (o *HyperflexHxRegistrationDetailsDt) HasOutOfComplianceStartDate() bool`

HasOutOfComplianceStartDate returns a boolean if a field has been set.

### GetRegistrationExpireDate

`func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationExpireDate() string`

GetRegistrationExpireDate returns the RegistrationExpireDate field if non-nil, zero value otherwise.

### GetRegistrationExpireDateOk

`func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationExpireDateOk() (*string, bool)`

GetRegistrationExpireDateOk returns a tuple with the RegistrationExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationExpireDate

`func (o *HyperflexHxRegistrationDetailsDt) SetRegistrationExpireDate(v string)`

SetRegistrationExpireDate sets RegistrationExpireDate field to given value.

### HasRegistrationExpireDate

`func (o *HyperflexHxRegistrationDetailsDt) HasRegistrationExpireDate() bool`

HasRegistrationExpireDate returns a boolean if a field has been set.

### GetRegistrationFailedReason

`func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationFailedReason() string`

GetRegistrationFailedReason returns the RegistrationFailedReason field if non-nil, zero value otherwise.

### GetRegistrationFailedReasonOk

`func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationFailedReasonOk() (*string, bool)`

GetRegistrationFailedReasonOk returns a tuple with the RegistrationFailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationFailedReason

`func (o *HyperflexHxRegistrationDetailsDt) SetRegistrationFailedReason(v string)`

SetRegistrationFailedReason sets RegistrationFailedReason field to given value.

### HasRegistrationFailedReason

`func (o *HyperflexHxRegistrationDetailsDt) HasRegistrationFailedReason() bool`

HasRegistrationFailedReason returns a boolean if a field has been set.

### GetSmartAccount

`func (o *HyperflexHxRegistrationDetailsDt) GetSmartAccount() string`

GetSmartAccount returns the SmartAccount field if non-nil, zero value otherwise.

### GetSmartAccountOk

`func (o *HyperflexHxRegistrationDetailsDt) GetSmartAccountOk() (*string, bool)`

GetSmartAccountOk returns a tuple with the SmartAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccount

`func (o *HyperflexHxRegistrationDetailsDt) SetSmartAccount(v string)`

SetSmartAccount sets SmartAccount field to given value.

### HasSmartAccount

`func (o *HyperflexHxRegistrationDetailsDt) HasSmartAccount() bool`

HasSmartAccount returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxRegistrationDetailsDt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxRegistrationDetailsDt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxRegistrationDetailsDt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxRegistrationDetailsDt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *HyperflexHxRegistrationDetailsDt) GetVirtualAccount() string`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *HyperflexHxRegistrationDetailsDt) GetVirtualAccountOk() (*string, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *HyperflexHxRegistrationDetailsDt) SetVirtualAccount(v string)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *HyperflexHxRegistrationDetailsDt) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


