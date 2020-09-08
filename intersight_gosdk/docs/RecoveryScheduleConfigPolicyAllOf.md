# RecoveryScheduleConfigPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**RecoveryBackupSchedule**](recovery.BackupSchedule.md) |  | [optional] 
**BackupProfiles** | Pointer to [**[]RecoveryBackupProfileRelationship**](recovery.BackupProfile.Relationship.md) | An array of relationships to recoveryBackupProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewRecoveryScheduleConfigPolicyAllOf

`func NewRecoveryScheduleConfigPolicyAllOf() *RecoveryScheduleConfigPolicyAllOf`

NewRecoveryScheduleConfigPolicyAllOf instantiates a new RecoveryScheduleConfigPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryScheduleConfigPolicyAllOfWithDefaults

`func NewRecoveryScheduleConfigPolicyAllOfWithDefaults() *RecoveryScheduleConfigPolicyAllOf`

NewRecoveryScheduleConfigPolicyAllOfWithDefaults instantiates a new RecoveryScheduleConfigPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *RecoveryScheduleConfigPolicyAllOf) GetSchedule() RecoveryBackupSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RecoveryScheduleConfigPolicyAllOf) GetScheduleOk() (*RecoveryBackupSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RecoveryScheduleConfigPolicyAllOf) SetSchedule(v RecoveryBackupSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *RecoveryScheduleConfigPolicyAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetBackupProfiles

`func (o *RecoveryScheduleConfigPolicyAllOf) GetBackupProfiles() []RecoveryBackupProfileRelationship`

GetBackupProfiles returns the BackupProfiles field if non-nil, zero value otherwise.

### GetBackupProfilesOk

`func (o *RecoveryScheduleConfigPolicyAllOf) GetBackupProfilesOk() (*[]RecoveryBackupProfileRelationship, bool)`

GetBackupProfilesOk returns a tuple with the BackupProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProfiles

`func (o *RecoveryScheduleConfigPolicyAllOf) SetBackupProfiles(v []RecoveryBackupProfileRelationship)`

SetBackupProfiles sets BackupProfiles field to given value.

### HasBackupProfiles

`func (o *RecoveryScheduleConfigPolicyAllOf) HasBackupProfiles() bool`

HasBackupProfiles returns a boolean if a field has been set.

### SetBackupProfilesNil

`func (o *RecoveryScheduleConfigPolicyAllOf) SetBackupProfilesNil(b bool)`

 SetBackupProfilesNil sets the value for BackupProfiles to be an explicit nil

### UnsetBackupProfiles
`func (o *RecoveryScheduleConfigPolicyAllOf) UnsetBackupProfiles()`

UnsetBackupProfiles ensures that no value is present for BackupProfiles, not even an explicit nil
### GetOrganization

`func (o *RecoveryScheduleConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryScheduleConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryScheduleConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryScheduleConfigPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


