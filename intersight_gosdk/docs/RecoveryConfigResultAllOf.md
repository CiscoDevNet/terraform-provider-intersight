# RecoveryConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.ConfigResult"]
**Var0OnDemandBackup** | Pointer to [**RecoveryOnDemandBackupRelationship**](recovery.OnDemandBackup.Relationship.md) |  | [optional] 
**BackupProfile** | Pointer to [**RecoveryBackupProfileRelationship**](recovery.BackupProfile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]RecoveryConfigResultEntryRelationship**](recovery.ConfigResultEntry.Relationship.md) | An array of relationships to recoveryConfigResultEntry resources. | [optional] 

## Methods

### NewRecoveryConfigResultAllOf

`func NewRecoveryConfigResultAllOf(classId string, objectType string, ) *RecoveryConfigResultAllOf`

NewRecoveryConfigResultAllOf instantiates a new RecoveryConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryConfigResultAllOfWithDefaults

`func NewRecoveryConfigResultAllOfWithDefaults() *RecoveryConfigResultAllOf`

NewRecoveryConfigResultAllOfWithDefaults instantiates a new RecoveryConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryConfigResultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryConfigResultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryConfigResultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryConfigResultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryConfigResultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryConfigResultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVar0OnDemandBackup

`func (o *RecoveryConfigResultAllOf) GetVar0OnDemandBackup() RecoveryOnDemandBackupRelationship`

GetVar0OnDemandBackup returns the Var0OnDemandBackup field if non-nil, zero value otherwise.

### GetVar0OnDemandBackupOk

`func (o *RecoveryConfigResultAllOf) GetVar0OnDemandBackupOk() (*RecoveryOnDemandBackupRelationship, bool)`

GetVar0OnDemandBackupOk returns a tuple with the Var0OnDemandBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0OnDemandBackup

`func (o *RecoveryConfigResultAllOf) SetVar0OnDemandBackup(v RecoveryOnDemandBackupRelationship)`

SetVar0OnDemandBackup sets Var0OnDemandBackup field to given value.

### HasVar0OnDemandBackup

`func (o *RecoveryConfigResultAllOf) HasVar0OnDemandBackup() bool`

HasVar0OnDemandBackup returns a boolean if a field has been set.

### GetBackupProfile

`func (o *RecoveryConfigResultAllOf) GetBackupProfile() RecoveryBackupProfileRelationship`

GetBackupProfile returns the BackupProfile field if non-nil, zero value otherwise.

### GetBackupProfileOk

`func (o *RecoveryConfigResultAllOf) GetBackupProfileOk() (*RecoveryBackupProfileRelationship, bool)`

GetBackupProfileOk returns a tuple with the BackupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProfile

`func (o *RecoveryConfigResultAllOf) SetBackupProfile(v RecoveryBackupProfileRelationship)`

SetBackupProfile sets BackupProfile field to given value.

### HasBackupProfile

`func (o *RecoveryConfigResultAllOf) HasBackupProfile() bool`

HasBackupProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *RecoveryConfigResultAllOf) GetResultEntries() []RecoveryConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *RecoveryConfigResultAllOf) GetResultEntriesOk() (*[]RecoveryConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *RecoveryConfigResultAllOf) SetResultEntries(v []RecoveryConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *RecoveryConfigResultAllOf) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *RecoveryConfigResultAllOf) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *RecoveryConfigResultAllOf) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


