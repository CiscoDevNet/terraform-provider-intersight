# RecoveryOnDemandBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigResult** | Pointer to [**RecoveryConfigResultRelationship**](recovery.ConfigResult.Relationship.md) |  | [optional] 
**DeviceId** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewRecoveryOnDemandBackup

`func NewRecoveryOnDemandBackup() *RecoveryOnDemandBackup`

NewRecoveryOnDemandBackup instantiates a new RecoveryOnDemandBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryOnDemandBackupWithDefaults

`func NewRecoveryOnDemandBackupWithDefaults() *RecoveryOnDemandBackup`

NewRecoveryOnDemandBackupWithDefaults instantiates a new RecoveryOnDemandBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigResult

`func (o *RecoveryOnDemandBackup) GetConfigResult() RecoveryConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *RecoveryOnDemandBackup) GetConfigResultOk() (*RecoveryConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *RecoveryOnDemandBackup) SetConfigResult(v RecoveryConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *RecoveryOnDemandBackup) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetDeviceId

`func (o *RecoveryOnDemandBackup) GetDeviceId() AssetDeviceRegistrationRelationship`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RecoveryOnDemandBackup) GetDeviceIdOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RecoveryOnDemandBackup) SetDeviceId(v AssetDeviceRegistrationRelationship)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *RecoveryOnDemandBackup) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetOrganization

`func (o *RecoveryOnDemandBackup) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryOnDemandBackup) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryOnDemandBackup) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryOnDemandBackup) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


