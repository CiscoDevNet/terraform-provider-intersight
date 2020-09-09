# IamTrustPointAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]X509Certificate**](x509.Certificate.md) |  | [optional] 
**Chain** | Pointer to **string** | The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewIamTrustPointAllOf

`func NewIamTrustPointAllOf() *IamTrustPointAllOf`

NewIamTrustPointAllOf instantiates a new IamTrustPointAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTrustPointAllOfWithDefaults

`func NewIamTrustPointAllOfWithDefaults() *IamTrustPointAllOf`

NewIamTrustPointAllOfWithDefaults instantiates a new IamTrustPointAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *IamTrustPointAllOf) GetCertificates() []X509Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *IamTrustPointAllOf) GetCertificatesOk() (*[]X509Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *IamTrustPointAllOf) SetCertificates(v []X509Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *IamTrustPointAllOf) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetChain

`func (o *IamTrustPointAllOf) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *IamTrustPointAllOf) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *IamTrustPointAllOf) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *IamTrustPointAllOf) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetAccount

`func (o *IamTrustPointAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamTrustPointAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamTrustPointAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamTrustPointAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


