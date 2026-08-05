# MgmtConfigOperationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigOperationSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigOperationSetting"]
**IsAccountKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;accountKey&#39; property has been set. | [optional] [readonly] [default to false]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewMgmtConfigOperationSetting

`func NewMgmtConfigOperationSetting(classId string, objectType string, ) *MgmtConfigOperationSetting`

NewMgmtConfigOperationSetting instantiates a new MgmtConfigOperationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigOperationSettingWithDefaults

`func NewMgmtConfigOperationSettingWithDefaults() *MgmtConfigOperationSetting`

NewMgmtConfigOperationSettingWithDefaults instantiates a new MgmtConfigOperationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigOperationSetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigOperationSetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigOperationSetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigOperationSetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigOperationSetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigOperationSetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsAccountKeySet

`func (o *MgmtConfigOperationSetting) GetIsAccountKeySet() bool`

GetIsAccountKeySet returns the IsAccountKeySet field if non-nil, zero value otherwise.

### GetIsAccountKeySetOk

`func (o *MgmtConfigOperationSetting) GetIsAccountKeySetOk() (*bool, bool)`

GetIsAccountKeySetOk returns a tuple with the IsAccountKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountKeySet

`func (o *MgmtConfigOperationSetting) SetIsAccountKeySet(v bool)`

SetIsAccountKeySet sets IsAccountKeySet field to given value.

### HasIsAccountKeySet

`func (o *MgmtConfigOperationSetting) HasIsAccountKeySet() bool`

HasIsAccountKeySet returns a boolean if a field has been set.

### GetAccount

`func (o *MgmtConfigOperationSetting) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MgmtConfigOperationSetting) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MgmtConfigOperationSetting) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MgmtConfigOperationSetting) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *MgmtConfigOperationSetting) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *MgmtConfigOperationSetting) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


