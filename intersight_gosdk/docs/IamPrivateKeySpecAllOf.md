# IamPrivateKeySpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to [**PkixKeyGenerationSpec**](pkix.KeyGenerationSpec.md) |  | [optional] 
**CertificateRequest** | Pointer to [**IamCertificateRequestRelationship**](iam.CertificateRequest.Relationship.md) |  | [optional] 

## Methods

### NewIamPrivateKeySpecAllOf

`func NewIamPrivateKeySpecAllOf() *IamPrivateKeySpecAllOf`

NewIamPrivateKeySpecAllOf instantiates a new IamPrivateKeySpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivateKeySpecAllOfWithDefaults

`func NewIamPrivateKeySpecAllOfWithDefaults() *IamPrivateKeySpecAllOf`

NewIamPrivateKeySpecAllOfWithDefaults instantiates a new IamPrivateKeySpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *IamPrivateKeySpecAllOf) GetAlgorithm() PkixKeyGenerationSpec`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *IamPrivateKeySpecAllOf) GetAlgorithmOk() (*PkixKeyGenerationSpec, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *IamPrivateKeySpecAllOf) SetAlgorithm(v PkixKeyGenerationSpec)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *IamPrivateKeySpecAllOf) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCertificateRequest

`func (o *IamPrivateKeySpecAllOf) GetCertificateRequest() IamCertificateRequestRelationship`

GetCertificateRequest returns the CertificateRequest field if non-nil, zero value otherwise.

### GetCertificateRequestOk

`func (o *IamPrivateKeySpecAllOf) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool)`

GetCertificateRequestOk returns a tuple with the CertificateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRequest

`func (o *IamPrivateKeySpecAllOf) SetCertificateRequest(v IamCertificateRequestRelationship)`

SetCertificateRequest sets CertificateRequest field to given value.

### HasCertificateRequest

`func (o *IamPrivateKeySpecAllOf) HasCertificateRequest() bool`

HasCertificateRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


