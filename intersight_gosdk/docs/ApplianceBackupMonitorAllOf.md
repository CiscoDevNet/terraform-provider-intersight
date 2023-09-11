# ApplianceBackupMonitorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.BackupMonitor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.BackupMonitor"]
**LastBackupStatus** | Pointer to **string** | Status of the most recent Intersight Appliance backup. * &#x60;BackupFound&#x60; - Backup is successful and complete. * &#x60;BackupFailed&#x60; - The current Backup failed. * &#x60;BackupOutdated&#x60; - Backup is old and outdated. | [optional] [readonly] [default to "BackupFound"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackupMonitorAllOf

`func NewApplianceBackupMonitorAllOf(classId string, objectType string, ) *ApplianceBackupMonitorAllOf`

NewApplianceBackupMonitorAllOf instantiates a new ApplianceBackupMonitorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupMonitorAllOfWithDefaults

`func NewApplianceBackupMonitorAllOfWithDefaults() *ApplianceBackupMonitorAllOf`

NewApplianceBackupMonitorAllOfWithDefaults instantiates a new ApplianceBackupMonitorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupMonitorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupMonitorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupMonitorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupMonitorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupMonitorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupMonitorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLastBackupStatus

`func (o *ApplianceBackupMonitorAllOf) GetLastBackupStatus() string`

GetLastBackupStatus returns the LastBackupStatus field if non-nil, zero value otherwise.

### GetLastBackupStatusOk

`func (o *ApplianceBackupMonitorAllOf) GetLastBackupStatusOk() (*string, bool)`

GetLastBackupStatusOk returns a tuple with the LastBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupStatus

`func (o *ApplianceBackupMonitorAllOf) SetLastBackupStatus(v string)`

SetLastBackupStatus sets LastBackupStatus field to given value.

### HasLastBackupStatus

`func (o *ApplianceBackupMonitorAllOf) HasLastBackupStatus() bool`

HasLastBackupStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceBackupMonitorAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackupMonitorAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackupMonitorAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackupMonitorAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


