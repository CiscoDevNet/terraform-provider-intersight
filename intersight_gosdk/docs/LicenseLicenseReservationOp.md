# LicenseLicenseReservationOp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | Pointer to **string** | Revervation code used to install the license. | [optional] 
**AuthCodeInstalled** | Pointer to **bool** | Flag to indicate whether authorization code is installed. | [optional] [readonly] 
**ConfirmCode** | Pointer to **string** | Confirm code used to complete the license update on smart license account. | [optional] [readonly] 
**GenerateRequestCode** | Pointer to **bool** | Trigger the generation of request code for specific license reservation. | [optional] 
**GenerateReturnCode** | Pointer to **bool** | Trigger the generation of return code for specific license reservation. | [optional] 
**RequestCode** | Pointer to **string** | Revervation code used to generate authorization code from CSSM. | [optional] [readonly] 
**ReturnCode** | Pointer to **string** | Return code used to return the reserved license to smart license account. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewLicenseLicenseReservationOp

`func NewLicenseLicenseReservationOp() *LicenseLicenseReservationOp`

NewLicenseLicenseReservationOp instantiates a new LicenseLicenseReservationOp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseReservationOpWithDefaults

`func NewLicenseLicenseReservationOpWithDefaults() *LicenseLicenseReservationOp`

NewLicenseLicenseReservationOpWithDefaults instantiates a new LicenseLicenseReservationOp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *LicenseLicenseReservationOp) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *LicenseLicenseReservationOp) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *LicenseLicenseReservationOp) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *LicenseLicenseReservationOp) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAuthCodeInstalled

`func (o *LicenseLicenseReservationOp) GetAuthCodeInstalled() bool`

GetAuthCodeInstalled returns the AuthCodeInstalled field if non-nil, zero value otherwise.

### GetAuthCodeInstalledOk

`func (o *LicenseLicenseReservationOp) GetAuthCodeInstalledOk() (*bool, bool)`

GetAuthCodeInstalledOk returns a tuple with the AuthCodeInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCodeInstalled

`func (o *LicenseLicenseReservationOp) SetAuthCodeInstalled(v bool)`

SetAuthCodeInstalled sets AuthCodeInstalled field to given value.

### HasAuthCodeInstalled

`func (o *LicenseLicenseReservationOp) HasAuthCodeInstalled() bool`

HasAuthCodeInstalled returns a boolean if a field has been set.

### GetConfirmCode

`func (o *LicenseLicenseReservationOp) GetConfirmCode() string`

GetConfirmCode returns the ConfirmCode field if non-nil, zero value otherwise.

### GetConfirmCodeOk

`func (o *LicenseLicenseReservationOp) GetConfirmCodeOk() (*string, bool)`

GetConfirmCodeOk returns a tuple with the ConfirmCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmCode

`func (o *LicenseLicenseReservationOp) SetConfirmCode(v string)`

SetConfirmCode sets ConfirmCode field to given value.

### HasConfirmCode

`func (o *LicenseLicenseReservationOp) HasConfirmCode() bool`

HasConfirmCode returns a boolean if a field has been set.

### GetGenerateRequestCode

`func (o *LicenseLicenseReservationOp) GetGenerateRequestCode() bool`

GetGenerateRequestCode returns the GenerateRequestCode field if non-nil, zero value otherwise.

### GetGenerateRequestCodeOk

`func (o *LicenseLicenseReservationOp) GetGenerateRequestCodeOk() (*bool, bool)`

GetGenerateRequestCodeOk returns a tuple with the GenerateRequestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRequestCode

`func (o *LicenseLicenseReservationOp) SetGenerateRequestCode(v bool)`

SetGenerateRequestCode sets GenerateRequestCode field to given value.

### HasGenerateRequestCode

`func (o *LicenseLicenseReservationOp) HasGenerateRequestCode() bool`

HasGenerateRequestCode returns a boolean if a field has been set.

### GetGenerateReturnCode

`func (o *LicenseLicenseReservationOp) GetGenerateReturnCode() bool`

GetGenerateReturnCode returns the GenerateReturnCode field if non-nil, zero value otherwise.

### GetGenerateReturnCodeOk

`func (o *LicenseLicenseReservationOp) GetGenerateReturnCodeOk() (*bool, bool)`

GetGenerateReturnCodeOk returns a tuple with the GenerateReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateReturnCode

`func (o *LicenseLicenseReservationOp) SetGenerateReturnCode(v bool)`

SetGenerateReturnCode sets GenerateReturnCode field to given value.

### HasGenerateReturnCode

`func (o *LicenseLicenseReservationOp) HasGenerateReturnCode() bool`

HasGenerateReturnCode returns a boolean if a field has been set.

### GetRequestCode

`func (o *LicenseLicenseReservationOp) GetRequestCode() string`

GetRequestCode returns the RequestCode field if non-nil, zero value otherwise.

### GetRequestCodeOk

`func (o *LicenseLicenseReservationOp) GetRequestCodeOk() (*string, bool)`

GetRequestCodeOk returns a tuple with the RequestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCode

`func (o *LicenseLicenseReservationOp) SetRequestCode(v string)`

SetRequestCode sets RequestCode field to given value.

### HasRequestCode

`func (o *LicenseLicenseReservationOp) HasRequestCode() bool`

HasRequestCode returns a boolean if a field has been set.

### GetReturnCode

`func (o *LicenseLicenseReservationOp) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *LicenseLicenseReservationOp) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *LicenseLicenseReservationOp) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *LicenseLicenseReservationOp) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetAccount

`func (o *LicenseLicenseReservationOp) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LicenseLicenseReservationOp) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LicenseLicenseReservationOp) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LicenseLicenseReservationOp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


