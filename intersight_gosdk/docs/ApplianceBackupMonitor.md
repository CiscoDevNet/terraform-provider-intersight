# ApplianceBackupMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.BackupMonitor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.BackupMonitor"]
**Filename** | Pointer to **string** | Filename of the backup for the backup monitor. | [optional] [readonly] 
**LastBackupRotationStatus** | Pointer to **string** | Status of the oldest Intersight Appliance backup cleanup. * &#x60;BackupFound&#x60; - Backup is successful and complete. * &#x60;BackupFailed&#x60; - The current Backup failed. * &#x60;BackupOutdated&#x60; - Backup is old and outdated. * &#x60;BackupCleanupFailed&#x60; - Cleanup of the old backup has failed. | [optional] [readonly] [default to "BackupFound"]
**LastBackupStatus** | Pointer to **string** | Status of the most recent Intersight Appliance backup. * &#x60;BackupFound&#x60; - Backup is successful and complete. * &#x60;BackupFailed&#x60; - The current Backup failed. * &#x60;BackupOutdated&#x60; - Backup is old and outdated. * &#x60;BackupCleanupFailed&#x60; - Cleanup of the old backup has failed. | [optional] [readonly] [default to "BackupFound"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackupMonitor

`func NewApplianceBackupMonitor(classId string, objectType string, ) *ApplianceBackupMonitor`

NewApplianceBackupMonitor instantiates a new ApplianceBackupMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupMonitorWithDefaults

`func NewApplianceBackupMonitorWithDefaults() *ApplianceBackupMonitor`

NewApplianceBackupMonitorWithDefaults instantiates a new ApplianceBackupMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupMonitor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupMonitor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupMonitor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupMonitor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupMonitor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupMonitor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilename

`func (o *ApplianceBackupMonitor) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ApplianceBackupMonitor) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ApplianceBackupMonitor) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ApplianceBackupMonitor) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetLastBackupRotationStatus

`func (o *ApplianceBackupMonitor) GetLastBackupRotationStatus() string`

GetLastBackupRotationStatus returns the LastBackupRotationStatus field if non-nil, zero value otherwise.

### GetLastBackupRotationStatusOk

`func (o *ApplianceBackupMonitor) GetLastBackupRotationStatusOk() (*string, bool)`

GetLastBackupRotationStatusOk returns a tuple with the LastBackupRotationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupRotationStatus

`func (o *ApplianceBackupMonitor) SetLastBackupRotationStatus(v string)`

SetLastBackupRotationStatus sets LastBackupRotationStatus field to given value.

### HasLastBackupRotationStatus

`func (o *ApplianceBackupMonitor) HasLastBackupRotationStatus() bool`

HasLastBackupRotationStatus returns a boolean if a field has been set.

### GetLastBackupStatus

`func (o *ApplianceBackupMonitor) GetLastBackupStatus() string`

GetLastBackupStatus returns the LastBackupStatus field if non-nil, zero value otherwise.

### GetLastBackupStatusOk

`func (o *ApplianceBackupMonitor) GetLastBackupStatusOk() (*string, bool)`

GetLastBackupStatusOk returns a tuple with the LastBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupStatus

`func (o *ApplianceBackupMonitor) SetLastBackupStatus(v string)`

SetLastBackupStatus sets LastBackupStatus field to given value.

### HasLastBackupStatus

`func (o *ApplianceBackupMonitor) HasLastBackupStatus() bool`

HasLastBackupStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceBackupMonitor) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackupMonitor) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackupMonitor) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackupMonitor) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceBackupMonitor) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceBackupMonitor) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


