# IamTrustPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.TrustPoint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.TrustPoint"]
**Certificates** | Pointer to [**[]X509Certificate**](X509Certificate.md) |  | [optional] 
**Chain** | Pointer to **string** | The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamTrustPoint

`func NewIamTrustPoint(classId string, objectType string, ) *IamTrustPoint`

NewIamTrustPoint instantiates a new IamTrustPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTrustPointWithDefaults

`func NewIamTrustPointWithDefaults() *IamTrustPoint`

NewIamTrustPointWithDefaults instantiates a new IamTrustPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamTrustPoint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamTrustPoint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamTrustPoint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamTrustPoint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamTrustPoint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamTrustPoint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificates

`func (o *IamTrustPoint) GetCertificates() []X509Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *IamTrustPoint) GetCertificatesOk() (*[]X509Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *IamTrustPoint) SetCertificates(v []X509Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *IamTrustPoint) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *IamTrustPoint) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *IamTrustPoint) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetChain

`func (o *IamTrustPoint) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *IamTrustPoint) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *IamTrustPoint) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *IamTrustPoint) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetAccount

`func (o *IamTrustPoint) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamTrustPoint) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamTrustPoint) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamTrustPoint) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


