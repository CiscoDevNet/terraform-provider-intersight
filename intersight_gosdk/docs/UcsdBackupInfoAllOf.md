# UcsdBackupInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupFileName** | Pointer to **string** | Auto generated backup File Name with combination of file prefix given an user input and the timestamp. | [optional] [readonly] 
**BackupLocation** | Pointer to **string** | Backup location that contains the backup images for end device which can be used for restore operation. | [optional] [readonly] 
**BackupServerIp** | Pointer to **string** | Backup server where backup images are maintained. | [optional] [readonly] 
**BackupSize** | Pointer to **int64** | Size of the backup image in bytes. | [optional] [readonly] 
**Connectors** | Pointer to [**[]UcsdConnectorPack**](ucsd.ConnectorPack.md) |  | [optional] 
**Duration** | Pointer to **int64** | Time taken for the backup to be completed. | [optional] [readonly] 
**EncryptionKey** | Pointer to **string** | The key used for encrypting the backup file. | [optional] 
**FailureReason** | Pointer to **string** | Reason for backup failure. | [optional] [readonly] 
**IsPurged** | Pointer to **bool** | Backup image got purged or not. The backup images get purged based on the retention count set by the user in the backup config policy. | [optional] [readonly] 
**LastModified** | Pointer to [**time.Time**](time.Time.md) | Last modified time when this backup record got updated. | [optional] [readonly] 
**PercentageCompletion** | Pointer to **int64** | Backup current precentage completion status information. | [optional] [readonly] 
**ProductVersion** | Pointer to **string** | The end device product version when the backup image was taken. | [optional] 
**Protocol** | Pointer to **string** | Protocol used for the remote backup. possible values are FTP, SCP and SFTP. Not applicable for the localhost (127.0.0.1). | [optional] [readonly] 
**StageCompletion** | Pointer to **string** | Backup current status stage information. | [optional] [readonly] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Start time of backup when it got initiated. | [optional] [readonly] 
**Status** | Pointer to **string** | Current status of Backup current. | [optional] [readonly] 

## Methods

### NewUcsdBackupInfoAllOf

`func NewUcsdBackupInfoAllOf() *UcsdBackupInfoAllOf`

NewUcsdBackupInfoAllOf instantiates a new UcsdBackupInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcsdBackupInfoAllOfWithDefaults

`func NewUcsdBackupInfoAllOfWithDefaults() *UcsdBackupInfoAllOf`

NewUcsdBackupInfoAllOfWithDefaults instantiates a new UcsdBackupInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupFileName

`func (o *UcsdBackupInfoAllOf) GetBackupFileName() string`

GetBackupFileName returns the BackupFileName field if non-nil, zero value otherwise.

### GetBackupFileNameOk

`func (o *UcsdBackupInfoAllOf) GetBackupFileNameOk() (*string, bool)`

GetBackupFileNameOk returns a tuple with the BackupFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFileName

`func (o *UcsdBackupInfoAllOf) SetBackupFileName(v string)`

SetBackupFileName sets BackupFileName field to given value.

### HasBackupFileName

`func (o *UcsdBackupInfoAllOf) HasBackupFileName() bool`

HasBackupFileName returns a boolean if a field has been set.

### GetBackupLocation

`func (o *UcsdBackupInfoAllOf) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *UcsdBackupInfoAllOf) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *UcsdBackupInfoAllOf) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.

### HasBackupLocation

`func (o *UcsdBackupInfoAllOf) HasBackupLocation() bool`

HasBackupLocation returns a boolean if a field has been set.

### GetBackupServerIp

`func (o *UcsdBackupInfoAllOf) GetBackupServerIp() string`

GetBackupServerIp returns the BackupServerIp field if non-nil, zero value otherwise.

### GetBackupServerIpOk

`func (o *UcsdBackupInfoAllOf) GetBackupServerIpOk() (*string, bool)`

GetBackupServerIpOk returns a tuple with the BackupServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServerIp

`func (o *UcsdBackupInfoAllOf) SetBackupServerIp(v string)`

SetBackupServerIp sets BackupServerIp field to given value.

### HasBackupServerIp

`func (o *UcsdBackupInfoAllOf) HasBackupServerIp() bool`

HasBackupServerIp returns a boolean if a field has been set.

### GetBackupSize

`func (o *UcsdBackupInfoAllOf) GetBackupSize() int64`

GetBackupSize returns the BackupSize field if non-nil, zero value otherwise.

### GetBackupSizeOk

`func (o *UcsdBackupInfoAllOf) GetBackupSizeOk() (*int64, bool)`

GetBackupSizeOk returns a tuple with the BackupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSize

`func (o *UcsdBackupInfoAllOf) SetBackupSize(v int64)`

SetBackupSize sets BackupSize field to given value.

### HasBackupSize

`func (o *UcsdBackupInfoAllOf) HasBackupSize() bool`

HasBackupSize returns a boolean if a field has been set.

### GetConnectors

`func (o *UcsdBackupInfoAllOf) GetConnectors() []UcsdConnectorPack`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *UcsdBackupInfoAllOf) GetConnectorsOk() (*[]UcsdConnectorPack, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *UcsdBackupInfoAllOf) SetConnectors(v []UcsdConnectorPack)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *UcsdBackupInfoAllOf) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetDuration

`func (o *UcsdBackupInfoAllOf) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UcsdBackupInfoAllOf) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UcsdBackupInfoAllOf) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UcsdBackupInfoAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *UcsdBackupInfoAllOf) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *UcsdBackupInfoAllOf) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *UcsdBackupInfoAllOf) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *UcsdBackupInfoAllOf) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetFailureReason

`func (o *UcsdBackupInfoAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UcsdBackupInfoAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UcsdBackupInfoAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UcsdBackupInfoAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetIsPurged

`func (o *UcsdBackupInfoAllOf) GetIsPurged() bool`

GetIsPurged returns the IsPurged field if non-nil, zero value otherwise.

### GetIsPurgedOk

`func (o *UcsdBackupInfoAllOf) GetIsPurgedOk() (*bool, bool)`

GetIsPurgedOk returns a tuple with the IsPurged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPurged

`func (o *UcsdBackupInfoAllOf) SetIsPurged(v bool)`

SetIsPurged sets IsPurged field to given value.

### HasIsPurged

`func (o *UcsdBackupInfoAllOf) HasIsPurged() bool`

HasIsPurged returns a boolean if a field has been set.

### GetLastModified

`func (o *UcsdBackupInfoAllOf) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *UcsdBackupInfoAllOf) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *UcsdBackupInfoAllOf) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *UcsdBackupInfoAllOf) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetPercentageCompletion

`func (o *UcsdBackupInfoAllOf) GetPercentageCompletion() int64`

GetPercentageCompletion returns the PercentageCompletion field if non-nil, zero value otherwise.

### GetPercentageCompletionOk

`func (o *UcsdBackupInfoAllOf) GetPercentageCompletionOk() (*int64, bool)`

GetPercentageCompletionOk returns a tuple with the PercentageCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompletion

`func (o *UcsdBackupInfoAllOf) SetPercentageCompletion(v int64)`

SetPercentageCompletion sets PercentageCompletion field to given value.

### HasPercentageCompletion

`func (o *UcsdBackupInfoAllOf) HasPercentageCompletion() bool`

HasPercentageCompletion returns a boolean if a field has been set.

### GetProductVersion

`func (o *UcsdBackupInfoAllOf) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *UcsdBackupInfoAllOf) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *UcsdBackupInfoAllOf) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *UcsdBackupInfoAllOf) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetProtocol

`func (o *UcsdBackupInfoAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UcsdBackupInfoAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UcsdBackupInfoAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UcsdBackupInfoAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStageCompletion

`func (o *UcsdBackupInfoAllOf) GetStageCompletion() string`

GetStageCompletion returns the StageCompletion field if non-nil, zero value otherwise.

### GetStageCompletionOk

`func (o *UcsdBackupInfoAllOf) GetStageCompletionOk() (*string, bool)`

GetStageCompletionOk returns a tuple with the StageCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageCompletion

`func (o *UcsdBackupInfoAllOf) SetStageCompletion(v string)`

SetStageCompletion sets StageCompletion field to given value.

### HasStageCompletion

`func (o *UcsdBackupInfoAllOf) HasStageCompletion() bool`

HasStageCompletion returns a boolean if a field has been set.

### GetStartTime

`func (o *UcsdBackupInfoAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UcsdBackupInfoAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UcsdBackupInfoAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UcsdBackupInfoAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UcsdBackupInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UcsdBackupInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UcsdBackupInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UcsdBackupInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


