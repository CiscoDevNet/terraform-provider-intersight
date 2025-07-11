# UcsdBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ucsd.BackupInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ucsd.BackupInfo"]
**BackupFileName** | Pointer to **string** | Auto generated backup file name with user defined prefix and timestamp. | [optional] [readonly] 
**BackupLocation** | Pointer to **string** | Backup location for restore operation. | [optional] [readonly] 
**BackupServerIp** | Pointer to **string** | Backup server for storing backup images. | [optional] [readonly] 
**BackupSize** | Pointer to **int64** | Backup image size in bytes. | [optional] [readonly] 
**Connectors** | Pointer to [**[]UcsdConnectorPack**](UcsdConnectorPack.md) |  | [optional] 
**Duration** | Pointer to **int64** | Time taken to complete the backup. | [optional] [readonly] 
**EncryptionKey** | Pointer to **string** | Encryption key for backup file. | [optional] 
**FailureReason** | Pointer to **string** | The cause of the backup failure. | [optional] [readonly] 
**IsPurged** | Pointer to **bool** | Backup image purge status based on retention policy. | [optional] [readonly] 
**LastModified** | Pointer to **time.Time** | Backup record last modified time. | [optional] [readonly] 
**PercentageCompletion** | Pointer to **int64** | Backup completion percentage status. | [optional] [readonly] 
**ProductVersion** | Pointer to **string** | End device product version at the backup time. | [optional] 
**Protocol** | Pointer to **string** | Supported remote backup protocol (FTP, SCP and SFTP); not applicable for localhost (127.0.0.1). | [optional] [readonly] 
**StageCompletion** | Pointer to **string** | Backup status stage information. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Backup initiation start time. | [optional] [readonly] 
**Status** | Pointer to **string** | The current backup status. | [optional] [readonly] 

## Methods

### NewUcsdBackupInfo

`func NewUcsdBackupInfo(classId string, objectType string, ) *UcsdBackupInfo`

NewUcsdBackupInfo instantiates a new UcsdBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcsdBackupInfoWithDefaults

`func NewUcsdBackupInfoWithDefaults() *UcsdBackupInfo`

NewUcsdBackupInfoWithDefaults instantiates a new UcsdBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UcsdBackupInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UcsdBackupInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UcsdBackupInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UcsdBackupInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UcsdBackupInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UcsdBackupInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupFileName

`func (o *UcsdBackupInfo) GetBackupFileName() string`

GetBackupFileName returns the BackupFileName field if non-nil, zero value otherwise.

### GetBackupFileNameOk

`func (o *UcsdBackupInfo) GetBackupFileNameOk() (*string, bool)`

GetBackupFileNameOk returns a tuple with the BackupFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFileName

`func (o *UcsdBackupInfo) SetBackupFileName(v string)`

SetBackupFileName sets BackupFileName field to given value.

### HasBackupFileName

`func (o *UcsdBackupInfo) HasBackupFileName() bool`

HasBackupFileName returns a boolean if a field has been set.

### GetBackupLocation

`func (o *UcsdBackupInfo) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *UcsdBackupInfo) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *UcsdBackupInfo) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.

### HasBackupLocation

`func (o *UcsdBackupInfo) HasBackupLocation() bool`

HasBackupLocation returns a boolean if a field has been set.

### GetBackupServerIp

`func (o *UcsdBackupInfo) GetBackupServerIp() string`

GetBackupServerIp returns the BackupServerIp field if non-nil, zero value otherwise.

### GetBackupServerIpOk

`func (o *UcsdBackupInfo) GetBackupServerIpOk() (*string, bool)`

GetBackupServerIpOk returns a tuple with the BackupServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServerIp

`func (o *UcsdBackupInfo) SetBackupServerIp(v string)`

SetBackupServerIp sets BackupServerIp field to given value.

### HasBackupServerIp

`func (o *UcsdBackupInfo) HasBackupServerIp() bool`

HasBackupServerIp returns a boolean if a field has been set.

### GetBackupSize

`func (o *UcsdBackupInfo) GetBackupSize() int64`

GetBackupSize returns the BackupSize field if non-nil, zero value otherwise.

### GetBackupSizeOk

`func (o *UcsdBackupInfo) GetBackupSizeOk() (*int64, bool)`

GetBackupSizeOk returns a tuple with the BackupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSize

`func (o *UcsdBackupInfo) SetBackupSize(v int64)`

SetBackupSize sets BackupSize field to given value.

### HasBackupSize

`func (o *UcsdBackupInfo) HasBackupSize() bool`

HasBackupSize returns a boolean if a field has been set.

### GetConnectors

`func (o *UcsdBackupInfo) GetConnectors() []UcsdConnectorPack`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *UcsdBackupInfo) GetConnectorsOk() (*[]UcsdConnectorPack, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *UcsdBackupInfo) SetConnectors(v []UcsdConnectorPack)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *UcsdBackupInfo) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### SetConnectorsNil

`func (o *UcsdBackupInfo) SetConnectorsNil(b bool)`

 SetConnectorsNil sets the value for Connectors to be an explicit nil

### UnsetConnectors
`func (o *UcsdBackupInfo) UnsetConnectors()`

UnsetConnectors ensures that no value is present for Connectors, not even an explicit nil
### GetDuration

`func (o *UcsdBackupInfo) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UcsdBackupInfo) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UcsdBackupInfo) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UcsdBackupInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *UcsdBackupInfo) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *UcsdBackupInfo) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *UcsdBackupInfo) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *UcsdBackupInfo) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetFailureReason

`func (o *UcsdBackupInfo) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UcsdBackupInfo) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UcsdBackupInfo) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UcsdBackupInfo) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetIsPurged

`func (o *UcsdBackupInfo) GetIsPurged() bool`

GetIsPurged returns the IsPurged field if non-nil, zero value otherwise.

### GetIsPurgedOk

`func (o *UcsdBackupInfo) GetIsPurgedOk() (*bool, bool)`

GetIsPurgedOk returns a tuple with the IsPurged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPurged

`func (o *UcsdBackupInfo) SetIsPurged(v bool)`

SetIsPurged sets IsPurged field to given value.

### HasIsPurged

`func (o *UcsdBackupInfo) HasIsPurged() bool`

HasIsPurged returns a boolean if a field has been set.

### GetLastModified

`func (o *UcsdBackupInfo) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *UcsdBackupInfo) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *UcsdBackupInfo) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *UcsdBackupInfo) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetPercentageCompletion

`func (o *UcsdBackupInfo) GetPercentageCompletion() int64`

GetPercentageCompletion returns the PercentageCompletion field if non-nil, zero value otherwise.

### GetPercentageCompletionOk

`func (o *UcsdBackupInfo) GetPercentageCompletionOk() (*int64, bool)`

GetPercentageCompletionOk returns a tuple with the PercentageCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompletion

`func (o *UcsdBackupInfo) SetPercentageCompletion(v int64)`

SetPercentageCompletion sets PercentageCompletion field to given value.

### HasPercentageCompletion

`func (o *UcsdBackupInfo) HasPercentageCompletion() bool`

HasPercentageCompletion returns a boolean if a field has been set.

### GetProductVersion

`func (o *UcsdBackupInfo) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *UcsdBackupInfo) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *UcsdBackupInfo) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *UcsdBackupInfo) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetProtocol

`func (o *UcsdBackupInfo) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UcsdBackupInfo) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UcsdBackupInfo) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UcsdBackupInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStageCompletion

`func (o *UcsdBackupInfo) GetStageCompletion() string`

GetStageCompletion returns the StageCompletion field if non-nil, zero value otherwise.

### GetStageCompletionOk

`func (o *UcsdBackupInfo) GetStageCompletionOk() (*string, bool)`

GetStageCompletionOk returns a tuple with the StageCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageCompletion

`func (o *UcsdBackupInfo) SetStageCompletion(v string)`

SetStageCompletion sets StageCompletion field to given value.

### HasStageCompletion

`func (o *UcsdBackupInfo) HasStageCompletion() bool`

HasStageCompletion returns a boolean if a field has been set.

### GetStartTime

`func (o *UcsdBackupInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UcsdBackupInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UcsdBackupInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UcsdBackupInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UcsdBackupInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UcsdBackupInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UcsdBackupInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UcsdBackupInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


