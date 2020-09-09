# ExternalsiteAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**IsUserIdSet** | Pointer to **bool** | Indicates whether the value of the &#39;userId&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | The password of the given username to download the image from external repository like cisco.com. | [optional] 
**RepositoryType** | Pointer to **string** | The repository type to which this authorization will be requested. Cisco is the only available repository today. * &#x60;cisco&#x60; - Cisco as an external site from where the resources like image will be downloaded. | [optional] [default to "cisco"]
**UserId** | Pointer to **string** | The username that has permission to download the image from external repository like cisco.com. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewExternalsiteAuthorization

`func NewExternalsiteAuthorization() *ExternalsiteAuthorization`

NewExternalsiteAuthorization instantiates a new ExternalsiteAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalsiteAuthorizationWithDefaults

`func NewExternalsiteAuthorizationWithDefaults() *ExternalsiteAuthorization`

NewExternalsiteAuthorizationWithDefaults instantiates a new ExternalsiteAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPasswordSet

`func (o *ExternalsiteAuthorization) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ExternalsiteAuthorization) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ExternalsiteAuthorization) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ExternalsiteAuthorization) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetIsUserIdSet

`func (o *ExternalsiteAuthorization) GetIsUserIdSet() bool`

GetIsUserIdSet returns the IsUserIdSet field if non-nil, zero value otherwise.

### GetIsUserIdSetOk

`func (o *ExternalsiteAuthorization) GetIsUserIdSetOk() (*bool, bool)`

GetIsUserIdSetOk returns a tuple with the IsUserIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserIdSet

`func (o *ExternalsiteAuthorization) SetIsUserIdSet(v bool)`

SetIsUserIdSet sets IsUserIdSet field to given value.

### HasIsUserIdSet

`func (o *ExternalsiteAuthorization) HasIsUserIdSet() bool`

HasIsUserIdSet returns a boolean if a field has been set.

### GetPassword

`func (o *ExternalsiteAuthorization) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ExternalsiteAuthorization) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ExternalsiteAuthorization) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ExternalsiteAuthorization) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRepositoryType

`func (o *ExternalsiteAuthorization) GetRepositoryType() string`

GetRepositoryType returns the RepositoryType field if non-nil, zero value otherwise.

### GetRepositoryTypeOk

`func (o *ExternalsiteAuthorization) GetRepositoryTypeOk() (*string, bool)`

GetRepositoryTypeOk returns a tuple with the RepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryType

`func (o *ExternalsiteAuthorization) SetRepositoryType(v string)`

SetRepositoryType sets RepositoryType field to given value.

### HasRepositoryType

`func (o *ExternalsiteAuthorization) HasRepositoryType() bool`

HasRepositoryType returns a boolean if a field has been set.

### GetUserId

`func (o *ExternalsiteAuthorization) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalsiteAuthorization) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalsiteAuthorization) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExternalsiteAuthorization) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccount

`func (o *ExternalsiteAuthorization) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExternalsiteAuthorization) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExternalsiteAuthorization) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExternalsiteAuthorization) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


