# StorageNetAppExportPolicyRuleAllOf

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

### NewStorageNetAppExportPolicyRuleAllOf

`func NewStorageNetAppExportPolicyRuleAllOf(classId string, objectType string, ) *StorageNetAppExportPolicyRuleAllOf`

NewStorageNetAppExportPolicyRuleAllOf instantiates a new StorageNetAppExportPolicyRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppExportPolicyRuleAllOfWithDefaults

`func NewStorageNetAppExportPolicyRuleAllOfWithDefaults() *StorageNetAppExportPolicyRuleAllOf`

NewStorageNetAppExportPolicyRuleAllOfWithDefaults instantiates a new StorageNetAppExportPolicyRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppExportPolicyRuleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppExportPolicyRuleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppExportPolicyRuleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppExportPolicyRuleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientMatch

`func (o *StorageNetAppExportPolicyRuleAllOf) GetClientMatch() []string`

GetClientMatch returns the ClientMatch field if non-nil, zero value otherwise.

### GetClientMatchOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetClientMatchOk() (*[]string, bool)`

GetClientMatchOk returns a tuple with the ClientMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMatch

`func (o *StorageNetAppExportPolicyRuleAllOf) SetClientMatch(v []string)`

SetClientMatch sets ClientMatch field to given value.

### HasClientMatch

`func (o *StorageNetAppExportPolicyRuleAllOf) HasClientMatch() bool`

HasClientMatch returns a boolean if a field has been set.

### SetClientMatchNil

`func (o *StorageNetAppExportPolicyRuleAllOf) SetClientMatchNil(b bool)`

 SetClientMatchNil sets the value for ClientMatch to be an explicit nil

### UnsetClientMatch
`func (o *StorageNetAppExportPolicyRuleAllOf) UnsetClientMatch()`

UnsetClientMatch ensures that no value is present for ClientMatch, not even an explicit nil
### GetIndex

`func (o *StorageNetAppExportPolicyRuleAllOf) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *StorageNetAppExportPolicyRuleAllOf) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *StorageNetAppExportPolicyRuleAllOf) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetRoRule

`func (o *StorageNetAppExportPolicyRuleAllOf) GetRoRule() []string`

GetRoRule returns the RoRule field if non-nil, zero value otherwise.

### GetRoRuleOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetRoRuleOk() (*[]string, bool)`

GetRoRuleOk returns a tuple with the RoRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoRule

`func (o *StorageNetAppExportPolicyRuleAllOf) SetRoRule(v []string)`

SetRoRule sets RoRule field to given value.

### HasRoRule

`func (o *StorageNetAppExportPolicyRuleAllOf) HasRoRule() bool`

HasRoRule returns a boolean if a field has been set.

### SetRoRuleNil

`func (o *StorageNetAppExportPolicyRuleAllOf) SetRoRuleNil(b bool)`

 SetRoRuleNil sets the value for RoRule to be an explicit nil

### UnsetRoRule
`func (o *StorageNetAppExportPolicyRuleAllOf) UnsetRoRule()`

UnsetRoRule ensures that no value is present for RoRule, not even an explicit nil
### GetRwRule

`func (o *StorageNetAppExportPolicyRuleAllOf) GetRwRule() []string`

GetRwRule returns the RwRule field if non-nil, zero value otherwise.

### GetRwRuleOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetRwRuleOk() (*[]string, bool)`

GetRwRuleOk returns a tuple with the RwRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRwRule

`func (o *StorageNetAppExportPolicyRuleAllOf) SetRwRule(v []string)`

SetRwRule sets RwRule field to given value.

### HasRwRule

`func (o *StorageNetAppExportPolicyRuleAllOf) HasRwRule() bool`

HasRwRule returns a boolean if a field has been set.

### SetRwRuleNil

`func (o *StorageNetAppExportPolicyRuleAllOf) SetRwRuleNil(b bool)`

 SetRwRuleNil sets the value for RwRule to be an explicit nil

### UnsetRwRule
`func (o *StorageNetAppExportPolicyRuleAllOf) UnsetRwRule()`

UnsetRwRule ensures that no value is present for RwRule, not even an explicit nil
### GetSuperUser

`func (o *StorageNetAppExportPolicyRuleAllOf) GetSuperUser() []string`

GetSuperUser returns the SuperUser field if non-nil, zero value otherwise.

### GetSuperUserOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetSuperUserOk() (*[]string, bool)`

GetSuperUserOk returns a tuple with the SuperUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperUser

`func (o *StorageNetAppExportPolicyRuleAllOf) SetSuperUser(v []string)`

SetSuperUser sets SuperUser field to given value.

### HasSuperUser

`func (o *StorageNetAppExportPolicyRuleAllOf) HasSuperUser() bool`

HasSuperUser returns a boolean if a field has been set.

### SetSuperUserNil

`func (o *StorageNetAppExportPolicyRuleAllOf) SetSuperUserNil(b bool)`

 SetSuperUserNil sets the value for SuperUser to be an explicit nil

### UnsetSuperUser
`func (o *StorageNetAppExportPolicyRuleAllOf) UnsetSuperUser()`

UnsetSuperUser ensures that no value is present for SuperUser, not even an explicit nil
### GetUser

`func (o *StorageNetAppExportPolicyRuleAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StorageNetAppExportPolicyRuleAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StorageNetAppExportPolicyRuleAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *StorageNetAppExportPolicyRuleAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


