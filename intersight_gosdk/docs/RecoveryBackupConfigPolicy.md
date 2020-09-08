# RecoveryBackupConfigPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupProfiles** | Pointer to [**[]RecoveryBackupProfileRelationship**](recovery.BackupProfile.Relationship.md) | An array of relationships to recoveryBackupProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewRecoveryBackupConfigPolicy

`func NewRecoveryBackupConfigPolicy() *RecoveryBackupConfigPolicy`

NewRecoveryBackupConfigPolicy instantiates a new RecoveryBackupConfigPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupConfigPolicyWithDefaults

`func NewRecoveryBackupConfigPolicyWithDefaults() *RecoveryBackupConfigPolicy`

NewRecoveryBackupConfigPolicyWithDefaults instantiates a new RecoveryBackupConfigPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupProfiles

`func (o *RecoveryBackupConfigPolicy) GetBackupProfiles() []RecoveryBackupProfileRelationship`

GetBackupProfiles returns the BackupProfiles field if non-nil, zero value otherwise.

### GetBackupProfilesOk

`func (o *RecoveryBackupConfigPolicy) GetBackupProfilesOk() (*[]RecoveryBackupProfileRelationship, bool)`

GetBackupProfilesOk returns a tuple with the BackupProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProfiles

`func (o *RecoveryBackupConfigPolicy) SetBackupProfiles(v []RecoveryBackupProfileRelationship)`

SetBackupProfiles sets BackupProfiles field to given value.

### HasBackupProfiles

`func (o *RecoveryBackupConfigPolicy) HasBackupProfiles() bool`

HasBackupProfiles returns a boolean if a field has been set.

### SetBackupProfilesNil

`func (o *RecoveryBackupConfigPolicy) SetBackupProfilesNil(b bool)`

 SetBackupProfilesNil sets the value for BackupProfiles to be an explicit nil

### UnsetBackupProfiles
`func (o *RecoveryBackupConfigPolicy) UnsetBackupProfiles()`

UnsetBackupProfiles ensures that no value is present for BackupProfiles, not even an explicit nil
### GetOrganization

`func (o *RecoveryBackupConfigPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryBackupConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryBackupConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryBackupConfigPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


