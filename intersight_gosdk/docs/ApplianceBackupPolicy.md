# ApplianceBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupTime** | Pointer to [**time.Time**](time.Time.md) | The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**ManualBackup** | Pointer to **bool** | Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to &#39;true&#39; and the backup schedule field is ignored. | [optional] 
**Password** | Pointer to **string** | Password to authenticate the file server. | [optional] 
**Schedule** | Pointer to [**OnpremSchedule**](onprem.Schedule.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewApplianceBackupPolicy

`func NewApplianceBackupPolicy() *ApplianceBackupPolicy`

NewApplianceBackupPolicy instantiates a new ApplianceBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupPolicyWithDefaults

`func NewApplianceBackupPolicyWithDefaults() *ApplianceBackupPolicy`

NewApplianceBackupPolicyWithDefaults instantiates a new ApplianceBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


