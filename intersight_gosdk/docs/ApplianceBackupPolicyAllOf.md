# ApplianceBackupPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.BackupPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.BackupPolicy"]
**BackupTime** | Pointer to **time.Time** | The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**ManualBackup** | Pointer to **bool** | Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to &#39;true&#39; and the backup schedule field is ignored. | [optional] 
**Password** | Pointer to **string** | Password to authenticate the file server. | [optional] 
**Schedule** | Pointer to [**NullableOnpremSchedule**](OnpremSchedule.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackupPolicyAllOf

`func NewApplianceBackupPolicyAllOf(classId string, objectType string, ) *ApplianceBackupPolicyAllOf`

NewApplianceBackupPolicyAllOf instantiates a new ApplianceBackupPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupPolicyAllOfWithDefaults

`func NewApplianceBackupPolicyAllOfWithDefaults() *ApplianceBackupPolicyAllOf`

NewApplianceBackupPolicyAllOfWithDefaults instantiates a new ApplianceBackupPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupTime

`func (o *ApplianceBackupPolicyAllOf) GetBackupTime() time.Time`

GetBackupTime returns the BackupTime field if non-nil, zero value otherwise.

### GetBackupTimeOk

`func (o *ApplianceBackupPolicyAllOf) GetBackupTimeOk() (*time.Time, bool)`

GetBackupTimeOk returns a tuple with the BackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTime

`func (o *ApplianceBackupPolicyAllOf) SetBackupTime(v time.Time)`

SetBackupTime sets BackupTime field to given value.

### HasBackupTime

`func (o *ApplianceBackupPolicyAllOf) HasBackupTime() bool`

HasBackupTime returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *ApplianceBackupPolicyAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ApplianceBackupPolicyAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ApplianceBackupPolicyAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ApplianceBackupPolicyAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetManualBackup

`func (o *ApplianceBackupPolicyAllOf) GetManualBackup() bool`

GetManualBackup returns the ManualBackup field if non-nil, zero value otherwise.

### GetManualBackupOk

`func (o *ApplianceBackupPolicyAllOf) GetManualBackupOk() (*bool, bool)`

GetManualBackupOk returns a tuple with the ManualBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualBackup

`func (o *ApplianceBackupPolicyAllOf) SetManualBackup(v bool)`

SetManualBackup sets ManualBackup field to given value.

### HasManualBackup

`func (o *ApplianceBackupPolicyAllOf) HasManualBackup() bool`

HasManualBackup returns a boolean if a field has been set.

### GetPassword

`func (o *ApplianceBackupPolicyAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApplianceBackupPolicyAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApplianceBackupPolicyAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApplianceBackupPolicyAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSchedule

`func (o *ApplianceBackupPolicyAllOf) GetSchedule() OnpremSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ApplianceBackupPolicyAllOf) GetScheduleOk() (*OnpremSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ApplianceBackupPolicyAllOf) SetSchedule(v OnpremSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ApplianceBackupPolicyAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *ApplianceBackupPolicyAllOf) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ApplianceBackupPolicyAllOf) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetAccount

`func (o *ApplianceBackupPolicyAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackupPolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackupPolicyAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackupPolicyAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


