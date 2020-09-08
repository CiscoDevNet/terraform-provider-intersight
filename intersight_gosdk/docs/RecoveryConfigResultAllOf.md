# RecoveryConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var0OnDemandBackup** | Pointer to [**RecoveryOnDemandBackupRelationship**](recovery.OnDemandBackup.Relationship.md) |  | [optional] 
**BackupProfile** | Pointer to [**RecoveryBackupProfileRelationship**](recovery.BackupProfile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]RecoveryConfigResultEntryRelationship**](recovery.ConfigResultEntry.Relationship.md) | An array of relationships to recoveryConfigResultEntry resources. | [optional] 

## Methods

### NewRecoveryConfigResultAllOf

`func NewRecoveryConfigResultAllOf() *RecoveryConfigResultAllOf`

NewRecoveryConfigResultAllOf instantiates a new RecoveryConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryConfigResultAllOfWithDefaults

`func NewRecoveryConfigResultAllOfWithDefaults() *RecoveryConfigResultAllOf`

NewRecoveryConfigResultAllOfWithDefaults instantiates a new RecoveryConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


