# IamSsoSessionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpSessionExpiration** | Pointer to **string** | SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid. | [optional] [readonly] 
**IdpSessionIndex** | Pointer to **string** | SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest. | [optional] [readonly] 

## Methods

### NewIamSsoSessionAttributes

`func NewIamSsoSessionAttributes() *IamSsoSessionAttributes`

NewIamSsoSessionAttributes instantiates a new IamSsoSessionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSsoSessionAttributesWithDefaults

`func NewIamSsoSessionAttributesWithDefaults() *IamSsoSessionAttributes`

NewIamSsoSessionAttributesWithDefaults instantiates a new IamSsoSessionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpSessionExpiration

`func (o *IamSsoSessionAttributes) GetIdpSessionExpiration() string`

GetIdpSessionExpiration returns the IdpSessionExpiration field if non-nil, zero value otherwise.

### GetIdpSessionExpirationOk

`func (o *IamSsoSessionAttributes) GetIdpSessionExpirationOk() (*string, bool)`

GetIdpSessionExpirationOk returns a tuple with the IdpSessionExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSessionExpiration

`func (o *IamSsoSessionAttributes) SetIdpSessionExpiration(v string)`

SetIdpSessionExpiration sets IdpSessionExpiration field to given value.

### HasIdpSessionExpiration

`func (o *IamSsoSessionAttributes) HasIdpSessionExpiration() bool`

HasIdpSessionExpiration returns a boolean if a field has been set.

### GetIdpSessionIndex

`func (o *IamSsoSessionAttributes) GetIdpSessionIndex() string`

GetIdpSessionIndex returns the IdpSessionIndex field if non-nil, zero value otherwise.

### GetIdpSessionIndexOk

`func (o *IamSsoSessionAttributes) GetIdpSessionIndexOk() (*string, bool)`

GetIdpSessionIndexOk returns a tuple with the IdpSessionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSessionIndex

`func (o *IamSsoSessionAttributes) SetIdpSessionIndex(v string)`

SetIdpSessionIndex sets IdpSessionIndex field to given value.

### HasIdpSessionIndex

`func (o *IamSsoSessionAttributes) HasIdpSessionIndex() bool`

HasIdpSessionIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


