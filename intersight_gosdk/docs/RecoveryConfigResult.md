# RecoveryConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.ConfigResult"]
**BackupProfile** | Pointer to [**RecoveryBackupProfileRelationship**](recovery.BackupProfile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]RecoveryConfigResultEntryRelationship**](RecoveryConfigResultEntryRelationship.md) | An array of relationships to recoveryConfigResultEntry resources. | [optional] 

## Methods

### NewRecoveryConfigResult

`func NewRecoveryConfigResult(classId string, objectType string, ) *RecoveryConfigResult`

NewRecoveryConfigResult instantiates a new RecoveryConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryConfigResultWithDefaults

`func NewRecoveryConfigResultWithDefaults() *RecoveryConfigResult`

NewRecoveryConfigResultWithDefaults instantiates a new RecoveryConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryConfigResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryConfigResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryConfigResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryConfigResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryConfigResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryConfigResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupProfile

`func (o *RecoveryConfigResult) GetBackupProfile() RecoveryBackupProfileRelationship`

GetBackupProfile returns the BackupProfile field if non-nil, zero value otherwise.

### GetBackupProfileOk

`func (o *RecoveryConfigResult) GetBackupProfileOk() (*RecoveryBackupProfileRelationship, bool)`

GetBackupProfileOk returns a tuple with the BackupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProfile

`func (o *RecoveryConfigResult) SetBackupProfile(v RecoveryBackupProfileRelationship)`

SetBackupProfile sets BackupProfile field to given value.

### HasBackupProfile

`func (o *RecoveryConfigResult) HasBackupProfile() bool`

HasBackupProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *RecoveryConfigResult) GetResultEntries() []RecoveryConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *RecoveryConfigResult) GetResultEntriesOk() (*[]RecoveryConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *RecoveryConfigResult) SetResultEntries(v []RecoveryConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *RecoveryConfigResult) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *RecoveryConfigResult) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *RecoveryConfigResult) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


