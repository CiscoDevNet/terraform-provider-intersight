# RecoveryBackupProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enables/Disables the schedule on the endpoint. | [optional] 
**BackupConfig** | Pointer to [**RecoveryBackupConfigPolicyRelationship**](recovery.BackupConfigPolicy.Relationship.md) |  | [optional] 
**ConfigResult** | Pointer to [**RecoveryConfigResultRelationship**](recovery.ConfigResult.Relationship.md) |  | [optional] 
**DeviceId** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**ScheduleConfig** | Pointer to [**RecoveryScheduleConfigPolicyRelationship**](recovery.ScheduleConfigPolicy.Relationship.md) |  | [optional] 

## Methods

### NewRecoveryBackupProfile

`func NewRecoveryBackupProfile() *RecoveryBackupProfile`

NewRecoveryBackupProfile instantiates a new RecoveryBackupProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupProfileWithDefaults

`func NewRecoveryBackupProfileWithDefaults() *RecoveryBackupProfile`

NewRecoveryBackupProfileWithDefaults instantiates a new RecoveryBackupProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RecoveryBackupProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecoveryBackupProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecoveryBackupProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RecoveryBackupProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBackupConfig

`func (o *RecoveryBackupProfile) GetBackupConfig() RecoveryBackupConfigPolicyRelationship`

GetBackupConfig returns the BackupConfig field if non-nil, zero value otherwise.

### GetBackupConfigOk

`func (o *RecoveryBackupProfile) GetBackupConfigOk() (*RecoveryBackupConfigPolicyRelationship, bool)`

GetBackupConfigOk returns a tuple with the BackupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfig

`func (o *RecoveryBackupProfile) SetBackupConfig(v RecoveryBackupConfigPolicyRelationship)`

SetBackupConfig sets BackupConfig field to given value.

### HasBackupConfig

`func (o *RecoveryBackupProfile) HasBackupConfig() bool`

HasBackupConfig returns a boolean if a field has been set.

### GetConfigResult

`func (o *RecoveryBackupProfile) GetConfigResult() RecoveryConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *RecoveryBackupProfile) GetConfigResultOk() (*RecoveryConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *RecoveryBackupProfile) SetConfigResult(v RecoveryConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *RecoveryBackupProfile) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetDeviceId

`func (o *RecoveryBackupProfile) GetDeviceId() AssetDeviceRegistrationRelationship`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RecoveryBackupProfile) GetDeviceIdOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RecoveryBackupProfile) SetDeviceId(v AssetDeviceRegistrationRelationship)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *RecoveryBackupProfile) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetOrganization

`func (o *RecoveryBackupProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryBackupProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryBackupProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryBackupProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScheduleConfig

`func (o *RecoveryBackupProfile) GetScheduleConfig() RecoveryScheduleConfigPolicyRelationship`

GetScheduleConfig returns the ScheduleConfig field if non-nil, zero value otherwise.

### GetScheduleConfigOk

`func (o *RecoveryBackupProfile) GetScheduleConfigOk() (*RecoveryScheduleConfigPolicyRelationship, bool)`

GetScheduleConfigOk returns a tuple with the ScheduleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleConfig

`func (o *RecoveryBackupProfile) SetScheduleConfig(v RecoveryScheduleConfigPolicyRelationship)`

SetScheduleConfig sets ScheduleConfig field to given value.

### HasScheduleConfig

`func (o *RecoveryBackupProfile) HasScheduleConfig() bool`

HasScheduleConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


