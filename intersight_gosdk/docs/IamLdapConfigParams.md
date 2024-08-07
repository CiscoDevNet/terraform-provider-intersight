# IamLdapConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapConfigParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapConfigParams"]
**ConfigHost** | Pointer to **string** | Stores the hostname based on which the issuer URL should be generated. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamLdapConfigParams

`func NewIamLdapConfigParams(classId string, objectType string, ) *IamLdapConfigParams`

NewIamLdapConfigParams instantiates a new IamLdapConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapConfigParamsWithDefaults

`func NewIamLdapConfigParamsWithDefaults() *IamLdapConfigParams`

NewIamLdapConfigParamsWithDefaults instantiates a new IamLdapConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapConfigParams) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapConfigParams) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapConfigParams) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapConfigParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapConfigParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapConfigParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigHost

`func (o *IamLdapConfigParams) GetConfigHost() string`

GetConfigHost returns the ConfigHost field if non-nil, zero value otherwise.

### GetConfigHostOk

`func (o *IamLdapConfigParams) GetConfigHostOk() (*string, bool)`

GetConfigHostOk returns a tuple with the ConfigHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigHost

`func (o *IamLdapConfigParams) SetConfigHost(v string)`

SetConfigHost sets ConfigHost field to given value.

### HasConfigHost

`func (o *IamLdapConfigParams) HasConfigHost() bool`

HasConfigHost returns a boolean if a field has been set.

### GetAccount

`func (o *IamLdapConfigParams) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamLdapConfigParams) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamLdapConfigParams) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamLdapConfigParams) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *IamLdapConfigParams) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamLdapConfigParams) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


