# RecoveryBackupProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.BackupProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.BackupProfile"]
**Enabled** | Pointer to **bool** | Enables/Disables the schedule on the endpoint. | [optional] [default to true]
**BackupConfig** | Pointer to [**RecoveryBackupConfigPolicyRelationship**](RecoveryBackupConfigPolicyRelationship.md) |  | [optional] 
**ConfigResult** | Pointer to [**RecoveryConfigResultRelationship**](RecoveryConfigResultRelationship.md) |  | [optional] 
**DeviceId** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ScheduleConfig** | Pointer to [**RecoveryScheduleConfigPolicyRelationship**](RecoveryScheduleConfigPolicyRelationship.md) |  | [optional] 

## Methods

### NewRecoveryBackupProfileAllOf

`func NewRecoveryBackupProfileAllOf(classId string, objectType string, ) *RecoveryBackupProfileAllOf`

NewRecoveryBackupProfileAllOf instantiates a new RecoveryBackupProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupProfileAllOfWithDefaults

`func NewRecoveryBackupProfileAllOfWithDefaults() *RecoveryBackupProfileAllOf`

NewRecoveryBackupProfileAllOfWithDefaults instantiates a new RecoveryBackupProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryBackupProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryBackupProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryBackupProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryBackupProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryBackupProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryBackupProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *RecoveryBackupProfileAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecoveryBackupProfileAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecoveryBackupProfileAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RecoveryBackupProfileAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBackupConfig

`func (o *RecoveryBackupProfileAllOf) GetBackupConfig() RecoveryBackupConfigPolicyRelationship`

GetBackupConfig returns the BackupConfig field if non-nil, zero value otherwise.

### GetBackupConfigOk

`func (o *RecoveryBackupProfileAllOf) GetBackupConfigOk() (*RecoveryBackupConfigPolicyRelationship, bool)`

GetBackupConfigOk returns a tuple with the BackupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfig

`func (o *RecoveryBackupProfileAllOf) SetBackupConfig(v RecoveryBackupConfigPolicyRelationship)`

SetBackupConfig sets BackupConfig field to given value.

### HasBackupConfig

`func (o *RecoveryBackupProfileAllOf) HasBackupConfig() bool`

HasBackupConfig returns a boolean if a field has been set.

### GetConfigResult

`func (o *RecoveryBackupProfileAllOf) GetConfigResult() RecoveryConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *RecoveryBackupProfileAllOf) GetConfigResultOk() (*RecoveryConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *RecoveryBackupProfileAllOf) SetConfigResult(v RecoveryConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *RecoveryBackupProfileAllOf) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetDeviceId

`func (o *RecoveryBackupProfileAllOf) GetDeviceId() AssetDeviceRegistrationRelationship`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RecoveryBackupProfileAllOf) GetDeviceIdOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RecoveryBackupProfileAllOf) SetDeviceId(v AssetDeviceRegistrationRelationship)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *RecoveryBackupProfileAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetOrganization

`func (o *RecoveryBackupProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryBackupProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryBackupProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryBackupProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScheduleConfig

`func (o *RecoveryBackupProfileAllOf) GetScheduleConfig() RecoveryScheduleConfigPolicyRelationship`

GetScheduleConfig returns the ScheduleConfig field if non-nil, zero value otherwise.

### GetScheduleConfigOk

`func (o *RecoveryBackupProfileAllOf) GetScheduleConfigOk() (*RecoveryScheduleConfigPolicyRelationship, bool)`

GetScheduleConfigOk returns a tuple with the ScheduleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleConfig

`func (o *RecoveryBackupProfileAllOf) SetScheduleConfig(v RecoveryScheduleConfigPolicyRelationship)`

SetScheduleConfig sets ScheduleConfig field to given value.

### HasScheduleConfig

`func (o *RecoveryBackupProfileAllOf) HasScheduleConfig() bool`

HasScheduleConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


