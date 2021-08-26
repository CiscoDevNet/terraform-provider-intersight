# X509Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "x509.Certificate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "x509.Certificate"]
**Issuer** | Pointer to [**NullablePkixDistinguishedName**](PkixDistinguishedName.md) |  | [optional] 
**NotAfter** | Pointer to **time.Time** | The date on which the certificate&#39;s validity period ends. | [optional] [readonly] 
**NotBefore** | Pointer to **time.Time** | The date on which the certificate&#39;s validity period begins. | [optional] [readonly] 
**PemCertificate** | Pointer to **string** | The base64 encoded certificate in PEM format. | [optional] 
**Sha256Fingerprint** | Pointer to **string** | The computed SHA-256 fingerprint of the certificate. Equivalent to &#39;openssl x509 -fingerprint -sha256&#39;. | [optional] [readonly] 
**SignatureAlgorithm** | Pointer to **string** | Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280). | [optional] [readonly] 
**Subject** | Pointer to [**NullablePkixDistinguishedName**](PkixDistinguishedName.md) |  | [optional] 

## Methods

### NewX509Certificate

`func NewX509Certificate(classId string, objectType string, ) *X509Certificate`

NewX509Certificate instantiates a new X509Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateWithDefaults

`func NewX509CertificateWithDefaults() *X509Certificate`

NewX509CertificateWithDefaults instantiates a new X509Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *X509Certificate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *X509Certificate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *X509Certificate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *X509Certificate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *X509Certificate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *X509Certificate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIssuer

`func (o *X509Certificate) GetIssuer() PkixDistinguishedName`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *X509Certificate) GetIssuerOk() (*PkixDistinguishedName, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *X509Certificate) SetIssuer(v PkixDistinguishedName)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *X509Certificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *X509Certificate) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *X509Certificate) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetNotAfter

`func (o *X509Certificate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *X509Certificate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *X509Certificate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *X509Certificate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *X509Certificate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *X509Certificate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetPemCertificate

`func (o *X509Certificate) GetPemCertificate() string`

GetPemCertificate returns the PemCertificate field if non-nil, zero value otherwise.

### GetPemCertificateOk

`func (o *X509Certificate) GetPemCertificateOk() (*string, bool)`

GetPemCertificateOk returns a tuple with the PemCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemCertificate

`func (o *X509Certificate) SetPemCertificate(v string)`

SetPemCertificate sets PemCertificate field to given value.

### HasPemCertificate

`func (o *X509Certificate) HasPemCertificate() bool`

HasPemCertificate returns a boolean if a field has been set.

### GetSha256Fingerprint

`func (o *X509Certificate) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *X509Certificate) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *X509Certificate) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.

### HasSha256Fingerprint

`func (o *X509Certificate) HasSha256Fingerprint() bool`

HasSha256Fingerprint returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *X509Certificate) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *X509Certificate) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *X509Certificate) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *X509Certificate) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSubject

`func (o *X509Certificate) GetSubject() PkixDistinguishedName`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *X509Certificate) GetSubjectOk() (*PkixDistinguishedName, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *X509Certificate) SetSubject(v PkixDistinguishedName)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *X509Certificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *X509Certificate) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *X509Certificate) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


