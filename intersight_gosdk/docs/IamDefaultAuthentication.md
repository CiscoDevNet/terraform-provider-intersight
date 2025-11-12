# IamDefaultAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.DefaultAuthentication"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.DefaultAuthentication"]
**DefaultAuthenticationMethod** | Pointer to **string** | The default authentication method used to login into Intersight. * &#x60;Local&#x60; - Local authentication uses credentials stored within the Intersight platform for user access. * &#x60;SSO&#x60; - SSO authentication uses an external identity provider for user access. * &#x60;LDAP&#x60; - LDAP authentication uses external LDAP servers for user access. | [optional] [default to "Local"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**LdapDomain** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 

## Methods

### NewIamDefaultAuthentication

`func NewIamDefaultAuthentication(classId string, objectType string, ) *IamDefaultAuthentication`

NewIamDefaultAuthentication instantiates a new IamDefaultAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamDefaultAuthenticationWithDefaults

`func NewIamDefaultAuthenticationWithDefaults() *IamDefaultAuthentication`

NewIamDefaultAuthenticationWithDefaults instantiates a new IamDefaultAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamDefaultAuthentication) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamDefaultAuthentication) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamDefaultAuthentication) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamDefaultAuthentication) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamDefaultAuthentication) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamDefaultAuthentication) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultAuthenticationMethod

`func (o *IamDefaultAuthentication) GetDefaultAuthenticationMethod() string`

GetDefaultAuthenticationMethod returns the DefaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetDefaultAuthenticationMethodOk

`func (o *IamDefaultAuthentication) GetDefaultAuthenticationMethodOk() (*string, bool)`

GetDefaultAuthenticationMethodOk returns a tuple with the DefaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthenticationMethod

`func (o *IamDefaultAuthentication) SetDefaultAuthenticationMethod(v string)`

SetDefaultAuthenticationMethod sets DefaultAuthenticationMethod field to given value.

### HasDefaultAuthenticationMethod

`func (o *IamDefaultAuthentication) HasDefaultAuthenticationMethod() bool`

HasDefaultAuthenticationMethod returns a boolean if a field has been set.

### GetAccount

`func (o *IamDefaultAuthentication) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamDefaultAuthentication) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamDefaultAuthentication) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamDefaultAuthentication) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *IamDefaultAuthentication) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamDefaultAuthentication) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetLdapDomain

`func (o *IamDefaultAuthentication) GetLdapDomain() IamIdpRelationship`

GetLdapDomain returns the LdapDomain field if non-nil, zero value otherwise.

### GetLdapDomainOk

`func (o *IamDefaultAuthentication) GetLdapDomainOk() (*IamIdpRelationship, bool)`

GetLdapDomainOk returns a tuple with the LdapDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDomain

`func (o *IamDefaultAuthentication) SetLdapDomain(v IamIdpRelationship)`

SetLdapDomain sets LdapDomain field to given value.

### HasLdapDomain

`func (o *IamDefaultAuthentication) HasLdapDomain() bool`

HasLdapDomain returns a boolean if a field has been set.

### SetLdapDomainNil

`func (o *IamDefaultAuthentication) SetLdapDomainNil(b bool)`

 SetLdapDomainNil sets the value for LdapDomain to be an explicit nil

### UnsetLdapDomain
`func (o *IamDefaultAuthentication) UnsetLdapDomain()`

UnsetLdapDomain ensures that no value is present for LdapDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


