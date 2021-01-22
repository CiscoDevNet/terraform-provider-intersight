# StorageNetAppExportPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppExportPolicyRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppExportPolicyRule"]
**ClientMatch** | Pointer to **[]string** |  | [optional] 
**Index** | Pointer to **int64** | Position of export rule in the list of rules. | [optional] 
**RoRule** | Pointer to **[]string** |  | [optional] 
**RwRule** | Pointer to **[]string** |  | [optional] 
**SuperUser** | Pointer to **[]string** |  | [optional] 
**User** | Pointer to **string** | Export Policy rule that are mapped to this User ID. | [optional] 

## Methods

### NewStorageNetAppExportPolicyRule

`func NewStorageNetAppExportPolicyRule(classId string, objectType string, ) *StorageNetAppExportPolicyRule`

NewStorageNetAppExportPolicyRule instantiates a new StorageNetAppExportPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppExportPolicyRuleWithDefaults

`func NewStorageNetAppExportPolicyRuleWithDefaults() *StorageNetAppExportPolicyRule`

NewStorageNetAppExportPolicyRuleWithDefaults instantiates a new StorageNetAppExportPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppExportPolicyRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppExportPolicyRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppExportPolicyRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppExportPolicyRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppExportPolicyRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppExportPolicyRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientMatch

`func (o *StorageNetAppExportPolicyRule) GetClientMatch() []string`

GetClientMatch returns the ClientMatch field if non-nil, zero value otherwise.

### GetClientMatchOk

`func (o *StorageNetAppExportPolicyRule) GetClientMatchOk() (*[]string, bool)`

GetClientMatchOk returns a tuple with the ClientMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMatch

`func (o *StorageNetAppExportPolicyRule) SetClientMatch(v []string)`

SetClientMatch sets ClientMatch field to given value.

### HasClientMatch

`func (o *StorageNetAppExportPolicyRule) HasClientMatch() bool`

HasClientMatch returns a boolean if a field has been set.

### SetClientMatchNil

`func (o *StorageNetAppExportPolicyRule) SetClientMatchNil(b bool)`

 SetClientMatchNil sets the value for ClientMatch to be an explicit nil

### UnsetClientMatch
`func (o *StorageNetAppExportPolicyRule) UnsetClientMatch()`

UnsetClientMatch ensures that no value is present for ClientMatch, not even an explicit nil
### GetIndex

`func (o *StorageNetAppExportPolicyRule) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *StorageNetAppExportPolicyRule) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *StorageNetAppExportPolicyRule) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *StorageNetAppExportPolicyRule) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetRoRule

`func (o *StorageNetAppExportPolicyRule) GetRoRule() []string`

GetRoRule returns the RoRule field if non-nil, zero value otherwise.

### GetRoRuleOk

`func (o *StorageNetAppExportPolicyRule) GetRoRuleOk() (*[]string, bool)`

GetRoRuleOk returns a tuple with the RoRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoRule

`func (o *StorageNetAppExportPolicyRule) SetRoRule(v []string)`

SetRoRule sets RoRule field to given value.

### HasRoRule

`func (o *StorageNetAppExportPolicyRule) HasRoRule() bool`

HasRoRule returns a boolean if a field has been set.

### SetRoRuleNil

`func (o *StorageNetAppExportPolicyRule) SetRoRuleNil(b bool)`

 SetRoRuleNil sets the value for RoRule to be an explicit nil

### UnsetRoRule
`func (o *StorageNetAppExportPolicyRule) UnsetRoRule()`

UnsetRoRule ensures that no value is present for RoRule, not even an explicit nil
### GetRwRule

`func (o *StorageNetAppExportPolicyRule) GetRwRule() []string`

GetRwRule returns the RwRule field if non-nil, zero value otherwise.

### GetRwRuleOk

`func (o *StorageNetAppExportPolicyRule) GetRwRuleOk() (*[]string, bool)`

GetRwRuleOk returns a tuple with the RwRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRwRule

`func (o *StorageNetAppExportPolicyRule) SetRwRule(v []string)`

SetRwRule sets RwRule field to given value.

### HasRwRule

`func (o *StorageNetAppExportPolicyRule) HasRwRule() bool`

HasRwRule returns a boolean if a field has been set.

### SetRwRuleNil

`func (o *StorageNetAppExportPolicyRule) SetRwRuleNil(b bool)`

 SetRwRuleNil sets the value for RwRule to be an explicit nil

### UnsetRwRule
`func (o *StorageNetAppExportPolicyRule) UnsetRwRule()`

UnsetRwRule ensures that no value is present for RwRule, not even an explicit nil
### GetSuperUser

`func (o *StorageNetAppExportPolicyRule) GetSuperUser() []string`

GetSuperUser returns the SuperUser field if non-nil, zero value otherwise.

### GetSuperUserOk

`func (o *StorageNetAppExportPolicyRule) GetSuperUserOk() (*[]string, bool)`

GetSuperUserOk returns a tuple with the SuperUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperUser

`func (o *StorageNetAppExportPolicyRule) SetSuperUser(v []string)`

SetSuperUser sets SuperUser field to given value.

### HasSuperUser

`func (o *StorageNetAppExportPolicyRule) HasSuperUser() bool`

HasSuperUser returns a boolean if a field has been set.

### SetSuperUserNil

`func (o *StorageNetAppExportPolicyRule) SetSuperUserNil(b bool)`

 SetSuperUserNil sets the value for SuperUser to be an explicit nil

### UnsetSuperUser
`func (o *StorageNetAppExportPolicyRule) UnsetSuperUser()`

UnsetSuperUser ensures that no value is present for SuperUser, not even an explicit nil
### GetUser

`func (o *StorageNetAppExportPolicyRule) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StorageNetAppExportPolicyRule) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StorageNetAppExportPolicyRule) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *StorageNetAppExportPolicyRule) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


