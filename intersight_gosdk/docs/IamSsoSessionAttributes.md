# IamSsoSessionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SsoSessionAttributes"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SsoSessionAttributes"]
**IdpSessionExpiration** | Pointer to **string** | SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid. | [optional] [readonly] 
**IdpSessionIndex** | Pointer to **string** | SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest. | [optional] [readonly] 
**IsIdpInitiatedSso** | Pointer to **bool** | Sign-in is SP-Intitiated or IdP-Intitiated. | [optional] [readonly] 
**SubjectName** | Pointer to **string** | SAML Subject NameID attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest. | [optional] [readonly] 

## Methods

### NewIamSsoSessionAttributes

`func NewIamSsoSessionAttributes(classId string, objectType string, ) *IamSsoSessionAttributes`

NewIamSsoSessionAttributes instantiates a new IamSsoSessionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSsoSessionAttributesWithDefaults

`func NewIamSsoSessionAttributesWithDefaults() *IamSsoSessionAttributes`

NewIamSsoSessionAttributesWithDefaults instantiates a new IamSsoSessionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSsoSessionAttributes) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSsoSessionAttributes) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSsoSessionAttributes) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSsoSessionAttributes) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSsoSessionAttributes) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSsoSessionAttributes) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetIsIdpInitiatedSso

`func (o *IamSsoSessionAttributes) GetIsIdpInitiatedSso() bool`

GetIsIdpInitiatedSso returns the IsIdpInitiatedSso field if non-nil, zero value otherwise.

### GetIsIdpInitiatedSsoOk

`func (o *IamSsoSessionAttributes) GetIsIdpInitiatedSsoOk() (*bool, bool)`

GetIsIdpInitiatedSsoOk returns a tuple with the IsIdpInitiatedSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdpInitiatedSso

`func (o *IamSsoSessionAttributes) SetIsIdpInitiatedSso(v bool)`

SetIsIdpInitiatedSso sets IsIdpInitiatedSso field to given value.

### HasIsIdpInitiatedSso

`func (o *IamSsoSessionAttributes) HasIsIdpInitiatedSso() bool`

HasIsIdpInitiatedSso returns a boolean if a field has been set.

### GetSubjectName

`func (o *IamSsoSessionAttributes) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *IamSsoSessionAttributes) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *IamSsoSessionAttributes) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *IamSsoSessionAttributes) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


