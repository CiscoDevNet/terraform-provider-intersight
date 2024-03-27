# IamLdapBasePropertiesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapBaseProperties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapBaseProperties"]
**Attribute** | Pointer to **string** | Role and locale information of the user. | [optional] 
**BaseDn** | Pointer to **string** | Base Distinguished Name (DN). Starting point from where server will search for users and groups. | [optional] 
**BindDn** | Pointer to **string** | Distinguished Name (DN) of the user, that is used to authenticate against LDAP servers. | [optional] 
**BindMethod** | Pointer to **string** | Authentication method to access LDAP servers. * &#x60;LoginCredentials&#x60; - Requires the user credentials. If the bind process fails, then user is denied access. * &#x60;Anonymous&#x60; - Requires no username and password. If this option is selected and the LDAP server is configured for Anonymous logins, then the user gains access. * &#x60;ConfiguredCredentials&#x60; - Requires a known set of credentials to be specified for the initial bind process. If the initial bind process succeeds, then the distinguished name (DN) of the user name is queried and re-used for the re-binding process. If the re-binding process fails, then the user is denied access. | [optional] [default to "LoginCredentials"]
**Domain** | Pointer to **string** | The IPv4 domain that all users must be in. | [optional] 
**EnableEncryption** | Pointer to **bool** | If enabled, the endpoint encrypts all information it sends to the LDAP server. | [optional] 
**EnableGroupAuthorization** | Pointer to **bool** | If enabled, user authorization is also done at the group level for LDAP users not in the local user database. | [optional] 
**EnableNestedGroupSearch** | Pointer to **bool** | If enabled, an extended search walks the chain of ancestry all the way to the root and returns all the groups and subgroups, each of those groups belong to recursively. | [optional] [default to false]
**Filter** | Pointer to **string** | Criteria to identify entries in search requests. | [optional] 
**GroupAttribute** | Pointer to **string** | Groups to which an LDAP entry belongs. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**NestedGroupSearchDepth** | Pointer to **int64** | Search depth to look for a nested LDAP group in an LDAP group map. | [optional] [default to 128]
**Password** | Pointer to **string** | The password of the user for initial bind process. It can be any string that adheres to the following constraints. It can have character except spaces, tabs, line breaks. It cannot be more than 254 characters. | [optional] 
**Timeout** | Pointer to **int64** | LDAP authentication timeout duration, in seconds. | [optional] [default to 0]

## Methods

### NewIamLdapBasePropertiesAllOf

`func NewIamLdapBasePropertiesAllOf(classId string, objectType string, ) *IamLdapBasePropertiesAllOf`

NewIamLdapBasePropertiesAllOf instantiates a new IamLdapBasePropertiesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapBasePropertiesAllOfWithDefaults

`func NewIamLdapBasePropertiesAllOfWithDefaults() *IamLdapBasePropertiesAllOf`

NewIamLdapBasePropertiesAllOfWithDefaults instantiates a new IamLdapBasePropertiesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapBasePropertiesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapBasePropertiesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapBasePropertiesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapBasePropertiesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapBasePropertiesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapBasePropertiesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttribute

`func (o *IamLdapBasePropertiesAllOf) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *IamLdapBasePropertiesAllOf) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *IamLdapBasePropertiesAllOf) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *IamLdapBasePropertiesAllOf) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetBaseDn

`func (o *IamLdapBasePropertiesAllOf) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *IamLdapBasePropertiesAllOf) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *IamLdapBasePropertiesAllOf) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *IamLdapBasePropertiesAllOf) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindDn

`func (o *IamLdapBasePropertiesAllOf) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *IamLdapBasePropertiesAllOf) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *IamLdapBasePropertiesAllOf) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *IamLdapBasePropertiesAllOf) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindMethod

`func (o *IamLdapBasePropertiesAllOf) GetBindMethod() string`

GetBindMethod returns the BindMethod field if non-nil, zero value otherwise.

### GetBindMethodOk

`func (o *IamLdapBasePropertiesAllOf) GetBindMethodOk() (*string, bool)`

GetBindMethodOk returns a tuple with the BindMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMethod

`func (o *IamLdapBasePropertiesAllOf) SetBindMethod(v string)`

SetBindMethod sets BindMethod field to given value.

### HasBindMethod

`func (o *IamLdapBasePropertiesAllOf) HasBindMethod() bool`

HasBindMethod returns a boolean if a field has been set.

### GetDomain

`func (o *IamLdapBasePropertiesAllOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IamLdapBasePropertiesAllOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IamLdapBasePropertiesAllOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IamLdapBasePropertiesAllOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnableEncryption

`func (o *IamLdapBasePropertiesAllOf) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *IamLdapBasePropertiesAllOf) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *IamLdapBasePropertiesAllOf) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.

### HasEnableEncryption

`func (o *IamLdapBasePropertiesAllOf) HasEnableEncryption() bool`

HasEnableEncryption returns a boolean if a field has been set.

### GetEnableGroupAuthorization

`func (o *IamLdapBasePropertiesAllOf) GetEnableGroupAuthorization() bool`

GetEnableGroupAuthorization returns the EnableGroupAuthorization field if non-nil, zero value otherwise.

### GetEnableGroupAuthorizationOk

`func (o *IamLdapBasePropertiesAllOf) GetEnableGroupAuthorizationOk() (*bool, bool)`

GetEnableGroupAuthorizationOk returns a tuple with the EnableGroupAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGroupAuthorization

`func (o *IamLdapBasePropertiesAllOf) SetEnableGroupAuthorization(v bool)`

SetEnableGroupAuthorization sets EnableGroupAuthorization field to given value.

### HasEnableGroupAuthorization

`func (o *IamLdapBasePropertiesAllOf) HasEnableGroupAuthorization() bool`

HasEnableGroupAuthorization returns a boolean if a field has been set.

### GetEnableNestedGroupSearch

`func (o *IamLdapBasePropertiesAllOf) GetEnableNestedGroupSearch() bool`

GetEnableNestedGroupSearch returns the EnableNestedGroupSearch field if non-nil, zero value otherwise.

### GetEnableNestedGroupSearchOk

`func (o *IamLdapBasePropertiesAllOf) GetEnableNestedGroupSearchOk() (*bool, bool)`

GetEnableNestedGroupSearchOk returns a tuple with the EnableNestedGroupSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNestedGroupSearch

`func (o *IamLdapBasePropertiesAllOf) SetEnableNestedGroupSearch(v bool)`

SetEnableNestedGroupSearch sets EnableNestedGroupSearch field to given value.

### HasEnableNestedGroupSearch

`func (o *IamLdapBasePropertiesAllOf) HasEnableNestedGroupSearch() bool`

HasEnableNestedGroupSearch returns a boolean if a field has been set.

### GetFilter

`func (o *IamLdapBasePropertiesAllOf) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IamLdapBasePropertiesAllOf) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IamLdapBasePropertiesAllOf) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IamLdapBasePropertiesAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupAttribute

`func (o *IamLdapBasePropertiesAllOf) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *IamLdapBasePropertiesAllOf) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *IamLdapBasePropertiesAllOf) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *IamLdapBasePropertiesAllOf) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *IamLdapBasePropertiesAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *IamLdapBasePropertiesAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *IamLdapBasePropertiesAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *IamLdapBasePropertiesAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetNestedGroupSearchDepth

`func (o *IamLdapBasePropertiesAllOf) GetNestedGroupSearchDepth() int64`

GetNestedGroupSearchDepth returns the NestedGroupSearchDepth field if non-nil, zero value otherwise.

### GetNestedGroupSearchDepthOk

`func (o *IamLdapBasePropertiesAllOf) GetNestedGroupSearchDepthOk() (*int64, bool)`

GetNestedGroupSearchDepthOk returns a tuple with the NestedGroupSearchDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedGroupSearchDepth

`func (o *IamLdapBasePropertiesAllOf) SetNestedGroupSearchDepth(v int64)`

SetNestedGroupSearchDepth sets NestedGroupSearchDepth field to given value.

### HasNestedGroupSearchDepth

`func (o *IamLdapBasePropertiesAllOf) HasNestedGroupSearchDepth() bool`

HasNestedGroupSearchDepth returns a boolean if a field has been set.

### GetPassword

`func (o *IamLdapBasePropertiesAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamLdapBasePropertiesAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamLdapBasePropertiesAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamLdapBasePropertiesAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTimeout

`func (o *IamLdapBasePropertiesAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IamLdapBasePropertiesAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IamLdapBasePropertiesAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IamLdapBasePropertiesAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


