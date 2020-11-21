# IamSsoSessionAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SsoSessionAttributes"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SsoSessionAttributes"]
**IdpSessionExpiration** | Pointer to **string** | SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid. | [optional] [readonly] 
**IdpSessionIndex** | Pointer to **string** | SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest. | [optional] [readonly] 

## Methods

### NewIamSsoSessionAttributesAllOf

`func NewIamSsoSessionAttributesAllOf(classId string, objectType string, ) *IamSsoSessionAttributesAllOf`

NewIamSsoSessionAttributesAllOf instantiates a new IamSsoSessionAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSsoSessionAttributesAllOfWithDefaults

`func NewIamSsoSessionAttributesAllOfWithDefaults() *IamSsoSessionAttributesAllOf`

NewIamSsoSessionAttributesAllOfWithDefaults instantiates a new IamSsoSessionAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSsoSessionAttributesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSsoSessionAttributesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSsoSessionAttributesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSsoSessionAttributesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSsoSessionAttributesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSsoSessionAttributesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdpSessionExpiration

`func (o *IamSsoSessionAttributesAllOf) GetIdpSessionExpiration() string`

GetIdpSessionExpiration returns the IdpSessionExpiration field if non-nil, zero value otherwise.

### GetIdpSessionExpirationOk

`func (o *IamSsoSessionAttributesAllOf) GetIdpSessionExpirationOk() (*string, bool)`

GetIdpSessionExpirationOk returns a tuple with the IdpSessionExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSessionExpiration

`func (o *IamSsoSessionAttributesAllOf) SetIdpSessionExpiration(v string)`

SetIdpSessionExpiration sets IdpSessionExpiration field to given value.

### HasIdpSessionExpiration

`func (o *IamSsoSessionAttributesAllOf) HasIdpSessionExpiration() bool`

HasIdpSessionExpiration returns a boolean if a field has been set.

### GetIdpSessionIndex

`func (o *IamSsoSessionAttributesAllOf) GetIdpSessionIndex() string`

GetIdpSessionIndex returns the IdpSessionIndex field if non-nil, zero value otherwise.

### GetIdpSessionIndexOk

`func (o *IamSsoSessionAttributesAllOf) GetIdpSessionIndexOk() (*string, bool)`

GetIdpSessionIndexOk returns a tuple with the IdpSessionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSessionIndex

`func (o *IamSsoSessionAttributesAllOf) SetIdpSessionIndex(v string)`

SetIdpSessionIndex sets IdpSessionIndex field to given value.

### HasIdpSessionIndex

`func (o *IamSsoSessionAttributesAllOf) HasIdpSessionIndex() bool`

HasIdpSessionIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


