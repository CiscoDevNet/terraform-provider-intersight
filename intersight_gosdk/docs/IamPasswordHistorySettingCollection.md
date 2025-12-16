# IamPasswordHistorySettingCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.PasswordHistorySettingCollection"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.PasswordHistorySettingCollection"]
**Orphaned** | Pointer to **bool** | Orphaned indicates whether policy, server or profile changed or deleted related to this object. Such objects would be marked as orphaned for deletion. | [optional] [readonly] [default to false]
**PasswordHistoryObjects** | Pointer to [**[]IamPasswordHistorySetting**](IamPasswordHistorySetting.md) |  | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Policy** | Pointer to [**NullableIamEndPointUserPolicyRelationship**](IamEndPointUserPolicyRelationship.md) |  | [optional] 
**Profile** | Pointer to [**NullablePolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewIamPasswordHistorySettingCollection

`func NewIamPasswordHistorySettingCollection(classId string, objectType string, ) *IamPasswordHistorySettingCollection`

NewIamPasswordHistorySettingCollection instantiates a new IamPasswordHistorySettingCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPasswordHistorySettingCollectionWithDefaults

`func NewIamPasswordHistorySettingCollectionWithDefaults() *IamPasswordHistorySettingCollection`

NewIamPasswordHistorySettingCollectionWithDefaults instantiates a new IamPasswordHistorySettingCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPasswordHistorySettingCollection) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPasswordHistorySettingCollection) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPasswordHistorySettingCollection) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPasswordHistorySettingCollection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPasswordHistorySettingCollection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPasswordHistorySettingCollection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOrphaned

`func (o *IamPasswordHistorySettingCollection) GetOrphaned() bool`

GetOrphaned returns the Orphaned field if non-nil, zero value otherwise.

### GetOrphanedOk

`func (o *IamPasswordHistorySettingCollection) GetOrphanedOk() (*bool, bool)`

GetOrphanedOk returns a tuple with the Orphaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaned

`func (o *IamPasswordHistorySettingCollection) SetOrphaned(v bool)`

SetOrphaned sets Orphaned field to given value.

### HasOrphaned

`func (o *IamPasswordHistorySettingCollection) HasOrphaned() bool`

HasOrphaned returns a boolean if a field has been set.

### GetPasswordHistoryObjects

`func (o *IamPasswordHistorySettingCollection) GetPasswordHistoryObjects() []IamPasswordHistorySetting`

GetPasswordHistoryObjects returns the PasswordHistoryObjects field if non-nil, zero value otherwise.

### GetPasswordHistoryObjectsOk

`func (o *IamPasswordHistorySettingCollection) GetPasswordHistoryObjectsOk() (*[]IamPasswordHistorySetting, bool)`

GetPasswordHistoryObjectsOk returns a tuple with the PasswordHistoryObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryObjects

`func (o *IamPasswordHistorySettingCollection) SetPasswordHistoryObjects(v []IamPasswordHistorySetting)`

SetPasswordHistoryObjects sets PasswordHistoryObjects field to given value.

### HasPasswordHistoryObjects

`func (o *IamPasswordHistorySettingCollection) HasPasswordHistoryObjects() bool`

HasPasswordHistoryObjects returns a boolean if a field has been set.

### SetPasswordHistoryObjectsNil

`func (o *IamPasswordHistorySettingCollection) SetPasswordHistoryObjectsNil(b bool)`

 SetPasswordHistoryObjectsNil sets the value for PasswordHistoryObjects to be an explicit nil

### UnsetPasswordHistoryObjects
`func (o *IamPasswordHistorySettingCollection) UnsetPasswordHistoryObjects()`

UnsetPasswordHistoryObjects ensures that no value is present for PasswordHistoryObjects, not even an explicit nil
### GetAccount

`func (o *IamPasswordHistorySettingCollection) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamPasswordHistorySettingCollection) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamPasswordHistorySettingCollection) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamPasswordHistorySettingCollection) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *IamPasswordHistorySettingCollection) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamPasswordHistorySettingCollection) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetPolicy

`func (o *IamPasswordHistorySettingCollection) GetPolicy() IamEndPointUserPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IamPasswordHistorySettingCollection) GetPolicyOk() (*IamEndPointUserPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IamPasswordHistorySettingCollection) SetPolicy(v IamEndPointUserPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IamPasswordHistorySettingCollection) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *IamPasswordHistorySettingCollection) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *IamPasswordHistorySettingCollection) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetProfile

`func (o *IamPasswordHistorySettingCollection) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IamPasswordHistorySettingCollection) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IamPasswordHistorySettingCollection) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *IamPasswordHistorySettingCollection) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *IamPasswordHistorySettingCollection) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *IamPasswordHistorySettingCollection) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetServer

`func (o *IamPasswordHistorySettingCollection) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IamPasswordHistorySettingCollection) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IamPasswordHistorySettingCollection) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *IamPasswordHistorySettingCollection) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *IamPasswordHistorySettingCollection) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *IamPasswordHistorySettingCollection) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


