# LicenseLicenseReservationOpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.LicenseReservationOp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.LicenseReservationOp"]
**AuthCode** | Pointer to **string** | Revervation code used to install the license. | [optional] 
**AuthCodeInstalled** | Pointer to **bool** | Flag to indicate whether authorization code is installed. | [optional] [readonly] 
**ConfirmCode** | Pointer to **string** | Confirm code used to complete the license update on smart license account. | [optional] [readonly] 
**GenerateRequestCode** | Pointer to **bool** | Trigger the generation of request code for specific license reservation. | [optional] 
**GenerateReturnCode** | Pointer to **bool** | Trigger the generation of return code for specific license reservation. | [optional] 
**RequestCode** | Pointer to **string** | Revervation code used to generate authorization code from CSSM. | [optional] [readonly] 
**ReturnCode** | Pointer to **string** | Return code used to return the reserved license to smart license account. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewLicenseLicenseReservationOpAllOf

`func NewLicenseLicenseReservationOpAllOf(classId string, objectType string, ) *LicenseLicenseReservationOpAllOf`

NewLicenseLicenseReservationOpAllOf instantiates a new LicenseLicenseReservationOpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseReservationOpAllOfWithDefaults

`func NewLicenseLicenseReservationOpAllOfWithDefaults() *LicenseLicenseReservationOpAllOf`

NewLicenseLicenseReservationOpAllOfWithDefaults instantiates a new LicenseLicenseReservationOpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseLicenseReservationOpAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseLicenseReservationOpAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseLicenseReservationOpAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseLicenseReservationOpAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseLicenseReservationOpAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseLicenseReservationOpAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthCode

`func (o *LicenseLicenseReservationOpAllOf) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *LicenseLicenseReservationOpAllOf) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *LicenseLicenseReservationOpAllOf) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *LicenseLicenseReservationOpAllOf) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAuthCodeInstalled

`func (o *LicenseLicenseReservationOpAllOf) GetAuthCodeInstalled() bool`

GetAuthCodeInstalled returns the AuthCodeInstalled field if non-nil, zero value otherwise.

### GetAuthCodeInstalledOk

`func (o *LicenseLicenseReservationOpAllOf) GetAuthCodeInstalledOk() (*bool, bool)`

GetAuthCodeInstalledOk returns a tuple with the AuthCodeInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCodeInstalled

`func (o *LicenseLicenseReservationOpAllOf) SetAuthCodeInstalled(v bool)`

SetAuthCodeInstalled sets AuthCodeInstalled field to given value.

### HasAuthCodeInstalled

`func (o *LicenseLicenseReservationOpAllOf) HasAuthCodeInstalled() bool`

HasAuthCodeInstalled returns a boolean if a field has been set.

### GetConfirmCode

`func (o *LicenseLicenseReservationOpAllOf) GetConfirmCode() string`

GetConfirmCode returns the ConfirmCode field if non-nil, zero value otherwise.

### GetConfirmCodeOk

`func (o *LicenseLicenseReservationOpAllOf) GetConfirmCodeOk() (*string, bool)`

GetConfirmCodeOk returns a tuple with the ConfirmCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmCode

`func (o *LicenseLicenseReservationOpAllOf) SetConfirmCode(v string)`

SetConfirmCode sets ConfirmCode field to given value.

### HasConfirmCode

`func (o *LicenseLicenseReservationOpAllOf) HasConfirmCode() bool`

HasConfirmCode returns a boolean if a field has been set.

### GetGenerateRequestCode

`func (o *LicenseLicenseReservationOpAllOf) GetGenerateRequestCode() bool`

GetGenerateRequestCode returns the GenerateRequestCode field if non-nil, zero value otherwise.

### GetGenerateRequestCodeOk

`func (o *LicenseLicenseReservationOpAllOf) GetGenerateRequestCodeOk() (*bool, bool)`

GetGenerateRequestCodeOk returns a tuple with the GenerateRequestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRequestCode

`func (o *LicenseLicenseReservationOpAllOf) SetGenerateRequestCode(v bool)`

SetGenerateRequestCode sets GenerateRequestCode field to given value.

### HasGenerateRequestCode

`func (o *LicenseLicenseReservationOpAllOf) HasGenerateRequestCode() bool`

HasGenerateRequestCode returns a boolean if a field has been set.

### GetGenerateReturnCode

`func (o *LicenseLicenseReservationOpAllOf) GetGenerateReturnCode() bool`

GetGenerateReturnCode returns the GenerateReturnCode field if non-nil, zero value otherwise.

### GetGenerateReturnCodeOk

`func (o *LicenseLicenseReservationOpAllOf) GetGenerateReturnCodeOk() (*bool, bool)`

GetGenerateReturnCodeOk returns a tuple with the GenerateReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateReturnCode

`func (o *LicenseLicenseReservationOpAllOf) SetGenerateReturnCode(v bool)`

SetGenerateReturnCode sets GenerateReturnCode field to given value.

### HasGenerateReturnCode

`func (o *LicenseLicenseReservationOpAllOf) HasGenerateReturnCode() bool`

HasGenerateReturnCode returns a boolean if a field has been set.

### GetRequestCode

`func (o *LicenseLicenseReservationOpAllOf) GetRequestCode() string`

GetRequestCode returns the RequestCode field if non-nil, zero value otherwise.

### GetRequestCodeOk

`func (o *LicenseLicenseReservationOpAllOf) GetRequestCodeOk() (*string, bool)`

GetRequestCodeOk returns a tuple with the RequestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCode

`func (o *LicenseLicenseReservationOpAllOf) SetRequestCode(v string)`

SetRequestCode sets RequestCode field to given value.

### HasRequestCode

`func (o *LicenseLicenseReservationOpAllOf) HasRequestCode() bool`

HasRequestCode returns a boolean if a field has been set.

### GetReturnCode

`func (o *LicenseLicenseReservationOpAllOf) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *LicenseLicenseReservationOpAllOf) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *LicenseLicenseReservationOpAllOf) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *LicenseLicenseReservationOpAllOf) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetAccount

`func (o *LicenseLicenseReservationOpAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LicenseLicenseReservationOpAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LicenseLicenseReservationOpAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LicenseLicenseReservationOpAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


