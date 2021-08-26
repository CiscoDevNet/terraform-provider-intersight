# IamCertificateRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.CertificateRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.CertificateRequest"]
**EmailAddress** | Pointer to **string** | User input email address, an optional part of the subject of the certificate request. | [optional] 
**Name** | Pointer to **string** | Name of the certificate request. | [optional] 
**Request** | Pointer to **string** | Generated certificate signing request (CSR) in PEM format. | [optional] [readonly] 
**SelfSigned** | Pointer to **bool** | Whether the user wants the generated CSR to be self-signed by the appliance. | [optional] 
**Subject** | Pointer to [**NullablePkixDistinguishedName**](PkixDistinguishedName.md) |  | [optional] 
**SubjectAlternateName** | Pointer to [**NullablePkixSubjectAlternateName**](PkixSubjectAlternateName.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Certificate** | Pointer to [**IamCertificateRelationship**](IamCertificateRelationship.md) |  | [optional] 
**PrivateKeySpec** | Pointer to [**IamPrivateKeySpecRelationship**](IamPrivateKeySpecRelationship.md) |  | [optional] 

## Methods

### NewIamCertificateRequestAllOf

`func NewIamCertificateRequestAllOf(classId string, objectType string, ) *IamCertificateRequestAllOf`

NewIamCertificateRequestAllOf instantiates a new IamCertificateRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCertificateRequestAllOfWithDefaults

`func NewIamCertificateRequestAllOfWithDefaults() *IamCertificateRequestAllOf`

NewIamCertificateRequestAllOfWithDefaults instantiates a new IamCertificateRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamCertificateRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamCertificateRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamCertificateRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamCertificateRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamCertificateRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamCertificateRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEmailAddress

`func (o *IamCertificateRequestAllOf) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *IamCertificateRequestAllOf) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *IamCertificateRequestAllOf) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *IamCertificateRequestAllOf) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetName

`func (o *IamCertificateRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamCertificateRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamCertificateRequestAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamCertificateRequestAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequest

`func (o *IamCertificateRequestAllOf) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *IamCertificateRequestAllOf) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *IamCertificateRequestAllOf) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *IamCertificateRequestAllOf) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSelfSigned

`func (o *IamCertificateRequestAllOf) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *IamCertificateRequestAllOf) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *IamCertificateRequestAllOf) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *IamCertificateRequestAllOf) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetSubject

`func (o *IamCertificateRequestAllOf) GetSubject() PkixDistinguishedName`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IamCertificateRequestAllOf) GetSubjectOk() (*PkixDistinguishedName, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IamCertificateRequestAllOf) SetSubject(v PkixDistinguishedName)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IamCertificateRequestAllOf) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *IamCertificateRequestAllOf) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *IamCertificateRequestAllOf) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSubjectAlternateName

`func (o *IamCertificateRequestAllOf) GetSubjectAlternateName() PkixSubjectAlternateName`

GetSubjectAlternateName returns the SubjectAlternateName field if non-nil, zero value otherwise.

### GetSubjectAlternateNameOk

`func (o *IamCertificateRequestAllOf) GetSubjectAlternateNameOk() (*PkixSubjectAlternateName, bool)`

GetSubjectAlternateNameOk returns a tuple with the SubjectAlternateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternateName

`func (o *IamCertificateRequestAllOf) SetSubjectAlternateName(v PkixSubjectAlternateName)`

SetSubjectAlternateName sets SubjectAlternateName field to given value.

### HasSubjectAlternateName

`func (o *IamCertificateRequestAllOf) HasSubjectAlternateName() bool`

HasSubjectAlternateName returns a boolean if a field has been set.

### SetSubjectAlternateNameNil

`func (o *IamCertificateRequestAllOf) SetSubjectAlternateNameNil(b bool)`

 SetSubjectAlternateNameNil sets the value for SubjectAlternateName to be an explicit nil

### UnsetSubjectAlternateName
`func (o *IamCertificateRequestAllOf) UnsetSubjectAlternateName()`

UnsetSubjectAlternateName ensures that no value is present for SubjectAlternateName, not even an explicit nil
### GetAccount

`func (o *IamCertificateRequestAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamCertificateRequestAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamCertificateRequestAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamCertificateRequestAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCertificate

`func (o *IamCertificateRequestAllOf) GetCertificate() IamCertificateRelationship`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IamCertificateRequestAllOf) GetCertificateOk() (*IamCertificateRelationship, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IamCertificateRequestAllOf) SetCertificate(v IamCertificateRelationship)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IamCertificateRequestAllOf) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKeySpec

`func (o *IamCertificateRequestAllOf) GetPrivateKeySpec() IamPrivateKeySpecRelationship`

GetPrivateKeySpec returns the PrivateKeySpec field if non-nil, zero value otherwise.

### GetPrivateKeySpecOk

`func (o *IamCertificateRequestAllOf) GetPrivateKeySpecOk() (*IamPrivateKeySpecRelationship, bool)`

GetPrivateKeySpecOk returns a tuple with the PrivateKeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeySpec

`func (o *IamCertificateRequestAllOf) SetPrivateKeySpec(v IamPrivateKeySpecRelationship)`

SetPrivateKeySpec sets PrivateKeySpec field to given value.

### HasPrivateKeySpec

`func (o *IamCertificateRequestAllOf) HasPrivateKeySpec() bool`

HasPrivateKeySpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


