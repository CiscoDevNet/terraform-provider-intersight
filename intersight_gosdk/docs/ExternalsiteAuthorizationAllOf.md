# ExternalsiteAuthorizationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "externalsite.Authorization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "externalsite.Authorization"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**IsUserIdSet** | Pointer to **bool** | Indicates whether the value of the &#39;userId&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password of the given username to download the image from external repository like cisco.com. | [optional] 
**RepositoryType** | Pointer to **string** | The repository type to which this authorization will be requested. Cisco is the only available repository today. * &#x60;cisco&#x60; - Cisco as an external site from where the resources like image will be downloaded. | [optional] [default to "cisco"]
**UserId** | Pointer to **string** | The username that has permission to download the image from external repository like cisco.com. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewExternalsiteAuthorizationAllOf

`func NewExternalsiteAuthorizationAllOf(classId string, objectType string, ) *ExternalsiteAuthorizationAllOf`

NewExternalsiteAuthorizationAllOf instantiates a new ExternalsiteAuthorizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalsiteAuthorizationAllOfWithDefaults

`func NewExternalsiteAuthorizationAllOfWithDefaults() *ExternalsiteAuthorizationAllOf`

NewExternalsiteAuthorizationAllOfWithDefaults instantiates a new ExternalsiteAuthorizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ExternalsiteAuthorizationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ExternalsiteAuthorizationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ExternalsiteAuthorizationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ExternalsiteAuthorizationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExternalsiteAuthorizationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExternalsiteAuthorizationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *ExternalsiteAuthorizationAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ExternalsiteAuthorizationAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ExternalsiteAuthorizationAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ExternalsiteAuthorizationAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetIsUserIdSet

`func (o *ExternalsiteAuthorizationAllOf) GetIsUserIdSet() bool`

GetIsUserIdSet returns the IsUserIdSet field if non-nil, zero value otherwise.

### GetIsUserIdSetOk

`func (o *ExternalsiteAuthorizationAllOf) GetIsUserIdSetOk() (*bool, bool)`

GetIsUserIdSetOk returns a tuple with the IsUserIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserIdSet

`func (o *ExternalsiteAuthorizationAllOf) SetIsUserIdSet(v bool)`

SetIsUserIdSet sets IsUserIdSet field to given value.

### HasIsUserIdSet

`func (o *ExternalsiteAuthorizationAllOf) HasIsUserIdSet() bool`

HasIsUserIdSet returns a boolean if a field has been set.

### GetPassword

`func (o *ExternalsiteAuthorizationAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ExternalsiteAuthorizationAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ExternalsiteAuthorizationAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ExternalsiteAuthorizationAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRepositoryType

`func (o *ExternalsiteAuthorizationAllOf) GetRepositoryType() string`

GetRepositoryType returns the RepositoryType field if non-nil, zero value otherwise.

### GetRepositoryTypeOk

`func (o *ExternalsiteAuthorizationAllOf) GetRepositoryTypeOk() (*string, bool)`

GetRepositoryTypeOk returns a tuple with the RepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryType

`func (o *ExternalsiteAuthorizationAllOf) SetRepositoryType(v string)`

SetRepositoryType sets RepositoryType field to given value.

### HasRepositoryType

`func (o *ExternalsiteAuthorizationAllOf) HasRepositoryType() bool`

HasRepositoryType returns a boolean if a field has been set.

### GetUserId

`func (o *ExternalsiteAuthorizationAllOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalsiteAuthorizationAllOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalsiteAuthorizationAllOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExternalsiteAuthorizationAllOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccount

`func (o *ExternalsiteAuthorizationAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExternalsiteAuthorizationAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExternalsiteAuthorizationAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExternalsiteAuthorizationAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


