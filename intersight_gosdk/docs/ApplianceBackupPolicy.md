# ApplianceBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.BackupPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.BackupPolicy"]
**BackupTime** | Pointer to **time.Time** | The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**ManualBackup** | Pointer to **bool** | Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to &#39;true&#39; and the backup schedule field is ignored. | [optional] 
**Password** | Pointer to **string** | Password to authenticate the file server. | [optional] 
**RetentionCount** | Pointer to **int64** | The number of backups before earliest backup is overwritten. Requires cleanup policy to be enabled. | [optional] [default to 1]
**RetentionPolicyEnabled** | Pointer to **bool** | If backup rotate policy is set, older backups will automatically be overwritten. The number of backups before overwriting is defined by the retentionCount property. | [optional] 
**Schedule** | Pointer to [**NullableOnpremSchedule**](OnpremSchedule.md) |  | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackupPolicy

`func NewApplianceBackupPolicy(classId string, objectType string, ) *ApplianceBackupPolicy`

NewApplianceBackupPolicy instantiates a new ApplianceBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupPolicyWithDefaults

`func NewApplianceBackupPolicyWithDefaults() *ApplianceBackupPolicy`

NewApplianceBackupPolicyWithDefaults instantiates a new ApplianceBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupTime

`func (o *ApplianceBackupPolicy) GetBackupTime() time.Time`

GetBackupTime returns the BackupTime field if non-nil, zero value otherwise.

### GetBackupTimeOk

`func (o *ApplianceBackupPolicy) GetBackupTimeOk() (*time.Time, bool)`

GetBackupTimeOk returns a tuple with the BackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTime

`func (o *ApplianceBackupPolicy) SetBackupTime(v time.Time)`

SetBackupTime sets BackupTime field to given value.

### HasBackupTime

`func (o *ApplianceBackupPolicy) HasBackupTime() bool`

HasBackupTime returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *ApplianceBackupPolicy) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ApplianceBackupPolicy) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ApplianceBackupPolicy) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ApplianceBackupPolicy) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetManualBackup

`func (o *ApplianceBackupPolicy) GetManualBackup() bool`

GetManualBackup returns the ManualBackup field if non-nil, zero value otherwise.

### GetManualBackupOk

`func (o *ApplianceBackupPolicy) GetManualBackupOk() (*bool, bool)`

GetManualBackupOk returns a tuple with the ManualBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualBackup

`func (o *ApplianceBackupPolicy) SetManualBackup(v bool)`

SetManualBackup sets ManualBackup field to given value.

### HasManualBackup

`func (o *ApplianceBackupPolicy) HasManualBackup() bool`

HasManualBackup returns a boolean if a field has been set.

### GetPassword

`func (o *ApplianceBackupPolicy) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApplianceBackupPolicy) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApplianceBackupPolicy) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApplianceBackupPolicy) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRetentionCount

`func (o *ApplianceBackupPolicy) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *ApplianceBackupPolicy) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *ApplianceBackupPolicy) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *ApplianceBackupPolicy) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetRetentionPolicyEnabled

`func (o *ApplianceBackupPolicy) GetRetentionPolicyEnabled() bool`

GetRetentionPolicyEnabled returns the RetentionPolicyEnabled field if non-nil, zero value otherwise.

### GetRetentionPolicyEnabledOk

`func (o *ApplianceBackupPolicy) GetRetentionPolicyEnabledOk() (*bool, bool)`

GetRetentionPolicyEnabledOk returns a tuple with the RetentionPolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyEnabled

`func (o *ApplianceBackupPolicy) SetRetentionPolicyEnabled(v bool)`

SetRetentionPolicyEnabled sets RetentionPolicyEnabled field to given value.

### HasRetentionPolicyEnabled

`func (o *ApplianceBackupPolicy) HasRetentionPolicyEnabled() bool`

HasRetentionPolicyEnabled returns a boolean if a field has been set.

### GetSchedule

`func (o *ApplianceBackupPolicy) GetSchedule() OnpremSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ApplianceBackupPolicy) GetScheduleOk() (*OnpremSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ApplianceBackupPolicy) SetSchedule(v OnpremSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ApplianceBackupPolicy) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *ApplianceBackupPolicy) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ApplianceBackupPolicy) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetAccount

`func (o *ApplianceBackupPolicy) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackupPolicy) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackupPolicy) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackupPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceBackupPolicy) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceBackupPolicy) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


