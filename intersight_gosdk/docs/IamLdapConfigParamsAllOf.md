# IamLdapConfigParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapConfigParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapConfigParams"]
**ConfigHost** | Pointer to **string** | Stores the hostname based on which the issuer URL should be generated. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamLdapConfigParamsAllOf

`func NewIamLdapConfigParamsAllOf(classId string, objectType string, ) *IamLdapConfigParamsAllOf`

NewIamLdapConfigParamsAllOf instantiates a new IamLdapConfigParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapConfigParamsAllOfWithDefaults

`func NewIamLdapConfigParamsAllOfWithDefaults() *IamLdapConfigParamsAllOf`

NewIamLdapConfigParamsAllOfWithDefaults instantiates a new IamLdapConfigParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapConfigParamsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapConfigParamsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapConfigParamsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapConfigParamsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapConfigParamsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapConfigParamsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigHost

`func (o *IamLdapConfigParamsAllOf) GetConfigHost() string`

GetConfigHost returns the ConfigHost field if non-nil, zero value otherwise.

### GetConfigHostOk

`func (o *IamLdapConfigParamsAllOf) GetConfigHostOk() (*string, bool)`

GetConfigHostOk returns a tuple with the ConfigHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigHost

`func (o *IamLdapConfigParamsAllOf) SetConfigHost(v string)`

SetConfigHost sets ConfigHost field to given value.

### HasConfigHost

`func (o *IamLdapConfigParamsAllOf) HasConfigHost() bool`

HasConfigHost returns a boolean if a field has been set.

### GetAccount

`func (o *IamLdapConfigParamsAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamLdapConfigParamsAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamLdapConfigParamsAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamLdapConfigParamsAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


