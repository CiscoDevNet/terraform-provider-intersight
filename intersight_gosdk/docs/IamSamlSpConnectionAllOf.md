# IamSamlSpConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SamlSpConnection"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SamlSpConnection"]
**IdentityProviderSloBinding** | Pointer to **string** | SingleLogout Services binding in IdP Metadata. | [optional] 
**IdentityProviderSloUrl** | Pointer to **string** | SingleLogOut Services location in IdP Metadata. | [optional] 
**IdentityProviderSsoBinding** | Pointer to **string** | SingleSignOn Services binding in IdP Metadata. | [optional] 
**IdentityProviderSsoUrl** | Pointer to **string** | SingleSignOn Services location in IdP Metadata. | [optional] 
**IdpCertificateStore** | Pointer to **interface{}** | Decoded Certificate from IdP Metatdata. | [optional] 
**SignAuthnRequests** | Pointer to **bool** | WantAuthnRequestsSigned from IdP Metadata. | [optional] 

## Methods

### NewIamSamlSpConnectionAllOf

`func NewIamSamlSpConnectionAllOf(classId string, objectType string, ) *IamSamlSpConnectionAllOf`

NewIamSamlSpConnectionAllOf instantiates a new IamSamlSpConnectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSamlSpConnectionAllOfWithDefaults

`func NewIamSamlSpConnectionAllOfWithDefaults() *IamSamlSpConnectionAllOf`

NewIamSamlSpConnectionAllOfWithDefaults instantiates a new IamSamlSpConnectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSamlSpConnectionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSamlSpConnectionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSamlSpConnectionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSamlSpConnectionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSamlSpConnectionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSamlSpConnectionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentityProviderSloBinding

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloBinding() string`

GetIdentityProviderSloBinding returns the IdentityProviderSloBinding field if non-nil, zero value otherwise.

### GetIdentityProviderSloBindingOk

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloBindingOk() (*string, bool)`

GetIdentityProviderSloBindingOk returns a tuple with the IdentityProviderSloBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderSloBinding

`func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSloBinding(v string)`

SetIdentityProviderSloBinding sets IdentityProviderSloBinding field to given value.

### HasIdentityProviderSloBinding

`func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSloBinding() bool`

HasIdentityProviderSloBinding returns a boolean if a field has been set.

### GetIdentityProviderSloUrl

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloUrl() string`

GetIdentityProviderSloUrl returns the IdentityProviderSloUrl field if non-nil, zero value otherwise.

### GetIdentityProviderSloUrlOk

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloUrlOk() (*string, bool)`

GetIdentityProviderSloUrlOk returns a tuple with the IdentityProviderSloUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderSloUrl

`func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSloUrl(v string)`

SetIdentityProviderSloUrl sets IdentityProviderSloUrl field to given value.

### HasIdentityProviderSloUrl

`func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSloUrl() bool`

HasIdentityProviderSloUrl returns a boolean if a field has been set.

### GetIdentityProviderSsoBinding

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoBinding() string`

GetIdentityProviderSsoBinding returns the IdentityProviderSsoBinding field if non-nil, zero value otherwise.

### GetIdentityProviderSsoBindingOk

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoBindingOk() (*string, bool)`

GetIdentityProviderSsoBindingOk returns a tuple with the IdentityProviderSsoBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderSsoBinding

`func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSsoBinding(v string)`

SetIdentityProviderSsoBinding sets IdentityProviderSsoBinding field to given value.

### HasIdentityProviderSsoBinding

`func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSsoBinding() bool`

HasIdentityProviderSsoBinding returns a boolean if a field has been set.

### GetIdentityProviderSsoUrl

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoUrl() string`

GetIdentityProviderSsoUrl returns the IdentityProviderSsoUrl field if non-nil, zero value otherwise.

### GetIdentityProviderSsoUrlOk

`func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoUrlOk() (*string, bool)`

GetIdentityProviderSsoUrlOk returns a tuple with the IdentityProviderSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderSsoUrl

`func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSsoUrl(v string)`

SetIdentityProviderSsoUrl sets IdentityProviderSsoUrl field to given value.

### HasIdentityProviderSsoUrl

`func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSsoUrl() bool`

HasIdentityProviderSsoUrl returns a boolean if a field has been set.

### GetIdpCertificateStore

`func (o *IamSamlSpConnectionAllOf) GetIdpCertificateStore() interface{}`

GetIdpCertificateStore returns the IdpCertificateStore field if non-nil, zero value otherwise.

### GetIdpCertificateStoreOk

`func (o *IamSamlSpConnectionAllOf) GetIdpCertificateStoreOk() (*interface{}, bool)`

GetIdpCertificateStoreOk returns a tuple with the IdpCertificateStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertificateStore

`func (o *IamSamlSpConnectionAllOf) SetIdpCertificateStore(v interface{})`

SetIdpCertificateStore sets IdpCertificateStore field to given value.

### HasIdpCertificateStore

`func (o *IamSamlSpConnectionAllOf) HasIdpCertificateStore() bool`

HasIdpCertificateStore returns a boolean if a field has been set.

### SetIdpCertificateStoreNil

`func (o *IamSamlSpConnectionAllOf) SetIdpCertificateStoreNil(b bool)`

 SetIdpCertificateStoreNil sets the value for IdpCertificateStore to be an explicit nil

### UnsetIdpCertificateStore
`func (o *IamSamlSpConnectionAllOf) UnsetIdpCertificateStore()`

UnsetIdpCertificateStore ensures that no value is present for IdpCertificateStore, not even an explicit nil
### GetSignAuthnRequests

`func (o *IamSamlSpConnectionAllOf) GetSignAuthnRequests() bool`

GetSignAuthnRequests returns the SignAuthnRequests field if non-nil, zero value otherwise.

### GetSignAuthnRequestsOk

`func (o *IamSamlSpConnectionAllOf) GetSignAuthnRequestsOk() (*bool, bool)`

GetSignAuthnRequestsOk returns a tuple with the SignAuthnRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthnRequests

`func (o *IamSamlSpConnectionAllOf) SetSignAuthnRequests(v bool)`

SetSignAuthnRequests sets SignAuthnRequests field to given value.

### HasSignAuthnRequests

`func (o *IamSamlSpConnectionAllOf) HasSignAuthnRequests() bool`

HasSignAuthnRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


