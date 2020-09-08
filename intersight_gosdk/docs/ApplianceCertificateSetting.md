# ApplianceCertificateSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Certificate** | Pointer to [**IamCertificateRelationship**](iam.Certificate.Relationship.md) |  | [optional] 

## Methods

### NewApplianceCertificateSetting

`func NewApplianceCertificateSetting() *ApplianceCertificateSetting`

NewApplianceCertificateSetting instantiates a new ApplianceCertificateSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceCertificateSettingWithDefaults

`func NewApplianceCertificateSettingWithDefaults() *ApplianceCertificateSetting`

NewApplianceCertificateSettingWithDefaults instantiates a new ApplianceCertificateSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ApplianceCertificateSetting) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceCertificateSetting) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceCertificateSetting) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceCertificateSetting) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCertificate

`func (o *ApplianceCertificateSetting) GetCertificate() IamCertificateRelationship`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ApplianceCertificateSetting) GetCertificateOk() (*IamCertificateRelationship, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ApplianceCertificateSetting) SetCertificate(v IamCertificateRelationship)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ApplianceCertificateSetting) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


