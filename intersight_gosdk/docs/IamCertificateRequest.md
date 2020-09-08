# IamCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | User input email address, an optional part of the subject of the certificate request. | [optional] 
**Name** | Pointer to **string** | Name of the certificate request. | [optional] 
**Request** | Pointer to **string** | Generated certificate signing request (CSR) in PEM format. | [optional] [readonly] 
**SelfSigned** | Pointer to **bool** | Whether the user wants the generated CSR to be self-signed by the appliance. | [optional] 
**Subject** | Pointer to [**PkixDistinguishedName**](pkix.DistinguishedName.md) |  | [optional] 
**SubjectAlternateName** | Pointer to [**PkixSubjectAlternateName**](pkix.SubjectAlternateName.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Certificate** | Pointer to [**IamCertificateRelationship**](iam.Certificate.Relationship.md) |  | [optional] 
**PrivateKeySpec** | Pointer to [**IamPrivateKeySpecRelationship**](iam.PrivateKeySpec.Relationship.md) |  | [optional] 

## Methods

### NewIamCertificateRequest

`func NewIamCertificateRequest() *IamCertificateRequest`

NewIamCertificateRequest instantiates a new IamCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCertificateRequestWithDefaults

`func NewIamCertificateRequestWithDefaults() *IamCertificateRequest`

NewIamCertificateRequestWithDefaults instantiates a new IamCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *IamCertificateRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *IamCertificateRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *IamCertificateRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *IamCertificateRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetName

`func (o *IamCertificateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamCertificateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamCertificateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamCertificateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequest

`func (o *IamCertificateRequest) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *IamCertificateRequest) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *IamCertificateRequest) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *IamCertificateRequest) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSelfSigned

`func (o *IamCertificateRequest) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *IamCertificateRequest) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *IamCertificateRequest) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *IamCertificateRequest) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetSubject

`func (o *IamCertificateRequest) GetSubject() PkixDistinguishedName`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IamCertificateRequest) GetSubjectOk() (*PkixDistinguishedName, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IamCertificateRequest) SetSubject(v PkixDistinguishedName)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IamCertificateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectAlternateName

`func (o *IamCertificateRequest) GetSubjectAlternateName() PkixSubjectAlternateName`

GetSubjectAlternateName returns the SubjectAlternateName field if non-nil, zero value otherwise.

### GetSubjectAlternateNameOk

`func (o *IamCertificateRequest) GetSubjectAlternateNameOk() (*PkixSubjectAlternateName, bool)`

GetSubjectAlternateNameOk returns a tuple with the SubjectAlternateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternateName

`func (o *IamCertificateRequest) SetSubjectAlternateName(v PkixSubjectAlternateName)`

SetSubjectAlternateName sets SubjectAlternateName field to given value.

### HasSubjectAlternateName

`func (o *IamCertificateRequest) HasSubjectAlternateName() bool`

HasSubjectAlternateName returns a boolean if a field has been set.

### GetAccount

`func (o *IamCertificateRequest) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamCertificateRequest) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamCertificateRequest) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamCertificateRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCertificate

`func (o *IamCertificateRequest) GetCertificate() IamCertificateRelationship`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IamCertificateRequest) GetCertificateOk() (*IamCertificateRelationship, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IamCertificateRequest) SetCertificate(v IamCertificateRelationship)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IamCertificateRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKeySpec

`func (o *IamCertificateRequest) GetPrivateKeySpec() IamPrivateKeySpecRelationship`

GetPrivateKeySpec returns the PrivateKeySpec field if non-nil, zero value otherwise.

### GetPrivateKeySpecOk

`func (o *IamCertificateRequest) GetPrivateKeySpecOk() (*IamPrivateKeySpecRelationship, bool)`

GetPrivateKeySpecOk returns a tuple with the PrivateKeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeySpec

`func (o *IamCertificateRequest) SetPrivateKeySpec(v IamPrivateKeySpecRelationship)`

SetPrivateKeySpec sets PrivateKeySpec field to given value.

### HasPrivateKeySpec

`func (o *IamCertificateRequest) HasPrivateKeySpec() bool`

HasPrivateKeySpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


