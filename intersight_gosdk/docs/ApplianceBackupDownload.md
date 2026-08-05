# ApplianceBackupDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.BackupDownload"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.BackupDownload"]
**Backup** | Pointer to [**NullableApplianceBackupRelationship**](ApplianceBackupRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackupDownload

`func NewApplianceBackupDownload(classId string, objectType string, ) *ApplianceBackupDownload`

NewApplianceBackupDownload instantiates a new ApplianceBackupDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupDownloadWithDefaults

`func NewApplianceBackupDownloadWithDefaults() *ApplianceBackupDownload`

NewApplianceBackupDownloadWithDefaults instantiates a new ApplianceBackupDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupDownload) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupDownload) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupDownload) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupDownload) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupDownload) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupDownload) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackup

`func (o *ApplianceBackupDownload) GetBackup() ApplianceBackupRelationship`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ApplianceBackupDownload) GetBackupOk() (*ApplianceBackupRelationship, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ApplianceBackupDownload) SetBackup(v ApplianceBackupRelationship)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ApplianceBackupDownload) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### SetBackupNil

`func (o *ApplianceBackupDownload) SetBackupNil(b bool)`

 SetBackupNil sets the value for Backup to be an explicit nil

### UnsetBackup
`func (o *ApplianceBackupDownload) UnsetBackup()`

UnsetBackup ensures that no value is present for Backup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


