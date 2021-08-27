# IamCertificateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Certificate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Certificate"]
**Certificate** | Pointer to [**NullableX509Certificate**](X509Certificate.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the certificate. * &#x60;PendingValidation&#x60; - The certificate has not been validated. * &#x60;Valid&#x60; - The certificate is valid. * &#x60;Invalid&#x60; - Ther certificate is invalid. | [optional] [readonly] [default to "PendingValidation"]
**CertificateRequest** | Pointer to [**IamCertificateRequestRelationship**](IamCertificateRequestRelationship.md) |  | [optional] 

## Methods

### NewIamCertificateAllOf

`func NewIamCertificateAllOf(classId string, objectType string, ) *IamCertificateAllOf`

NewIamCertificateAllOf instantiates a new IamCertificateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCertificateAllOfWithDefaults

`func NewIamCertificateAllOfWithDefaults() *IamCertificateAllOf`

NewIamCertificateAllOfWithDefaults instantiates a new IamCertificateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamCertificateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamCertificateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamCertificateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamCertificateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamCertificateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamCertificateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificate

`func (o *IamCertificateAllOf) GetCertificate() X509Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IamCertificateAllOf) GetCertificateOk() (*X509Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IamCertificateAllOf) SetCertificate(v X509Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IamCertificateAllOf) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *IamCertificateAllOf) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *IamCertificateAllOf) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetStatus

`func (o *IamCertificateAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamCertificateAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamCertificateAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamCertificateAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCertificateRequest

`func (o *IamCertificateAllOf) GetCertificateRequest() IamCertificateRequestRelationship`

GetCertificateRequest returns the CertificateRequest field if non-nil, zero value otherwise.

### GetCertificateRequestOk

`func (o *IamCertificateAllOf) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool)`

GetCertificateRequestOk returns a tuple with the CertificateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRequest

`func (o *IamCertificateAllOf) SetCertificateRequest(v IamCertificateRequestRelationship)`

SetCertificateRequest sets CertificateRequest field to given value.

### HasCertificateRequest

`func (o *IamCertificateAllOf) HasCertificateRequest() bool`

HasCertificateRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


