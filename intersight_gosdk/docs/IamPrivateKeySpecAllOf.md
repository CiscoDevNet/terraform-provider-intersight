# IamPrivateKeySpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.PrivateKeySpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.PrivateKeySpec"]
**Algorithm** | Pointer to [**NullablePkixKeyGenerationSpec**](PkixKeyGenerationSpec.md) |  | [optional] 
**CertificateRequest** | Pointer to [**IamCertificateRequestRelationship**](IamCertificateRequestRelationship.md) |  | [optional] 

## Methods

### NewIamPrivateKeySpecAllOf

`func NewIamPrivateKeySpecAllOf(classId string, objectType string, ) *IamPrivateKeySpecAllOf`

NewIamPrivateKeySpecAllOf instantiates a new IamPrivateKeySpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivateKeySpecAllOfWithDefaults

`func NewIamPrivateKeySpecAllOfWithDefaults() *IamPrivateKeySpecAllOf`

NewIamPrivateKeySpecAllOfWithDefaults instantiates a new IamPrivateKeySpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPrivateKeySpecAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPrivateKeySpecAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPrivateKeySpecAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPrivateKeySpecAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPrivateKeySpecAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPrivateKeySpecAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetAlgorithmNil

`func (o *IamPrivateKeySpecAllOf) SetAlgorithmNil(b bool)`

 SetAlgorithmNil sets the value for Algorithm to be an explicit nil

### UnsetAlgorithm
`func (o *IamPrivateKeySpecAllOf) UnsetAlgorithm()`

UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
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


