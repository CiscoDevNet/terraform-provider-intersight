# MgmtConfigBackupInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigBackupInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigBackupInstance"]
**AdminAction** | Pointer to **string** | Action to be performed on the backup instance. * &#x60;None&#x60; - Nothing to be done. The default state of the backup instance when there is no action to be performed. * &#x60;EncryptWithPublicKey&#x60; - Encrypt the backup using the user given public key and make it available for download. * &#x60;EncryptWithPassword&#x60; - Encrypt the backup using the user provided password and make it available for download. * &#x60;GenerateDownloadUrl&#x60; - Generate a pre-signed URL that can be used to download a backup archive. | [optional] [default to "None"]
**BackupCreationTime** | Pointer to **time.Time** | The time at which the backup file was created. | [optional] [readonly] 
**BackupName** | Pointer to **string** | Name of the backup for which the backup instance is created. | [optional] [readonly] 
**BackupSource** | Pointer to **string** | Indicates whether the backup was created locally or imported. * &#x60;Local&#x60; - A local Intersight location. * &#x60;Uploaded&#x60; - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set. * &#x60;Remote&#x60; - A remote location hosted in the user&#39;s datacenter. | [optional] [readonly] [default to "Local"]
**Checksum** | Pointer to **string** | SHA-256 checksum of the encrypted backup archive, returned as a hex-encoded string. Users can compare the checksum of a downloaded archive with this value to verify file integrity and confirm the archive was not corrupted or altered in transit. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the backup operation that created this instance. | [optional] 
**DownloadUrl** | Pointer to **string** | Pre-signed URL that can be used to download the backup archive. This URL is generated when the adminAction is set to GenerateDownloadUrl. The URL is valid for 5 minutes. | [optional] [readonly] 
**IsUserPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;userPassword&#39; property has been set. | [optional] [readonly] [default to false]
**Path** | Pointer to **string** | Path to the backup file in the Intersight repository. | [optional] [readonly] 
**PreserveIdentities** | Pointer to **bool** | The flag set by the user during a configuration backup to preserve static or dynamic IDs assigned to an export item. | [optional] [readonly] 
**RetentionLock** | Pointer to **bool** | When set, ensures that the backup archive is protected from deletion and rollover operations. | [optional] 
**Size** | Pointer to **int64** | Size of the backup instance, represents size of the compressed backup file. | [optional] [readonly] 
**SourceAccountMoid** | Pointer to **string** | The moid of the account from which the backup was taken. | [optional] [readonly] 
**SourceName** | Pointer to **string** | The name of the account from which the backup was taken. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the backup instance. * &#x60;Unknown&#x60; - The status is not known. When the backup instances are created, they are in this state until the backup is ready. * &#x60;Ready&#x60; - Ready for use in an import operation. * &#x60;RemovalInProgress&#x60; - The backup is being removed. * &#x60;Corrupted&#x60; - The backup is corrupted and is not usable. * &#x60;EncryptionInProgress&#x60; - The backup is being encrypted using the user provided public key. * &#x60;EncryptionFailed&#x60; - The backup could not be encrypted using the user provided public key. * &#x60;Encrypted&#x60; - The backup has been encrypted using the user provided public key. | [optional] [readonly] [default to "Unknown"]
**SystemAdminTriggered** | Pointer to **bool** | Set when a system administrator initiated the backup operation. | [optional] [readonly] 
**UserPassword** | Pointer to **string** | The password provided by the user to encrypt the backup before download. This password is not persisted in Intersight and this field will be reset after the backup archive is encrypted. The password encrypts the backup file before download. The user needs to provide the same password to decrypt when they upload the backup file to restore. | [optional] 
**UserPublicKey** | Pointer to **string** | The public key provided by the user to encrypt the backup before download. The key needs to be in PKCS#1 format encoded in PEM. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewMgmtConfigBackupInstance

`func NewMgmtConfigBackupInstance(classId string, objectType string, ) *MgmtConfigBackupInstance`

NewMgmtConfigBackupInstance instantiates a new MgmtConfigBackupInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigBackupInstanceWithDefaults

`func NewMgmtConfigBackupInstanceWithDefaults() *MgmtConfigBackupInstance`

NewMgmtConfigBackupInstanceWithDefaults instantiates a new MgmtConfigBackupInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigBackupInstance) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigBackupInstance) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigBackupInstance) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigBackupInstance) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigBackupInstance) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigBackupInstance) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *MgmtConfigBackupInstance) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *MgmtConfigBackupInstance) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *MgmtConfigBackupInstance) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *MgmtConfigBackupInstance) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetBackupCreationTime

`func (o *MgmtConfigBackupInstance) GetBackupCreationTime() time.Time`

GetBackupCreationTime returns the BackupCreationTime field if non-nil, zero value otherwise.

### GetBackupCreationTimeOk

`func (o *MgmtConfigBackupInstance) GetBackupCreationTimeOk() (*time.Time, bool)`

GetBackupCreationTimeOk returns a tuple with the BackupCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCreationTime

`func (o *MgmtConfigBackupInstance) SetBackupCreationTime(v time.Time)`

SetBackupCreationTime sets BackupCreationTime field to given value.

### HasBackupCreationTime

`func (o *MgmtConfigBackupInstance) HasBackupCreationTime() bool`

HasBackupCreationTime returns a boolean if a field has been set.

### GetBackupName

`func (o *MgmtConfigBackupInstance) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *MgmtConfigBackupInstance) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *MgmtConfigBackupInstance) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *MgmtConfigBackupInstance) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetBackupSource

`func (o *MgmtConfigBackupInstance) GetBackupSource() string`

GetBackupSource returns the BackupSource field if non-nil, zero value otherwise.

### GetBackupSourceOk

`func (o *MgmtConfigBackupInstance) GetBackupSourceOk() (*string, bool)`

GetBackupSourceOk returns a tuple with the BackupSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSource

`func (o *MgmtConfigBackupInstance) SetBackupSource(v string)`

SetBackupSource sets BackupSource field to given value.

### HasBackupSource

`func (o *MgmtConfigBackupInstance) HasBackupSource() bool`

HasBackupSource returns a boolean if a field has been set.

### GetChecksum

`func (o *MgmtConfigBackupInstance) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *MgmtConfigBackupInstance) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *MgmtConfigBackupInstance) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *MgmtConfigBackupInstance) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetDescription

`func (o *MgmtConfigBackupInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MgmtConfigBackupInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MgmtConfigBackupInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MgmtConfigBackupInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *MgmtConfigBackupInstance) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *MgmtConfigBackupInstance) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *MgmtConfigBackupInstance) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *MgmtConfigBackupInstance) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetIsUserPasswordSet

`func (o *MgmtConfigBackupInstance) GetIsUserPasswordSet() bool`

GetIsUserPasswordSet returns the IsUserPasswordSet field if non-nil, zero value otherwise.

### GetIsUserPasswordSetOk

`func (o *MgmtConfigBackupInstance) GetIsUserPasswordSetOk() (*bool, bool)`

GetIsUserPasswordSetOk returns a tuple with the IsUserPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserPasswordSet

`func (o *MgmtConfigBackupInstance) SetIsUserPasswordSet(v bool)`

SetIsUserPasswordSet sets IsUserPasswordSet field to given value.

### HasIsUserPasswordSet

`func (o *MgmtConfigBackupInstance) HasIsUserPasswordSet() bool`

HasIsUserPasswordSet returns a boolean if a field has been set.

### GetPath

`func (o *MgmtConfigBackupInstance) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MgmtConfigBackupInstance) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MgmtConfigBackupInstance) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MgmtConfigBackupInstance) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPreserveIdentities

`func (o *MgmtConfigBackupInstance) GetPreserveIdentities() bool`

GetPreserveIdentities returns the PreserveIdentities field if non-nil, zero value otherwise.

### GetPreserveIdentitiesOk

`func (o *MgmtConfigBackupInstance) GetPreserveIdentitiesOk() (*bool, bool)`

GetPreserveIdentitiesOk returns a tuple with the PreserveIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveIdentities

`func (o *MgmtConfigBackupInstance) SetPreserveIdentities(v bool)`

SetPreserveIdentities sets PreserveIdentities field to given value.

### HasPreserveIdentities

`func (o *MgmtConfigBackupInstance) HasPreserveIdentities() bool`

HasPreserveIdentities returns a boolean if a field has been set.

### GetRetentionLock

`func (o *MgmtConfigBackupInstance) GetRetentionLock() bool`

GetRetentionLock returns the RetentionLock field if non-nil, zero value otherwise.

### GetRetentionLockOk

`func (o *MgmtConfigBackupInstance) GetRetentionLockOk() (*bool, bool)`

GetRetentionLockOk returns a tuple with the RetentionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionLock

`func (o *MgmtConfigBackupInstance) SetRetentionLock(v bool)`

SetRetentionLock sets RetentionLock field to given value.

### HasRetentionLock

`func (o *MgmtConfigBackupInstance) HasRetentionLock() bool`

HasRetentionLock returns a boolean if a field has been set.

### GetSize

`func (o *MgmtConfigBackupInstance) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MgmtConfigBackupInstance) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MgmtConfigBackupInstance) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MgmtConfigBackupInstance) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceAccountMoid

`func (o *MgmtConfigBackupInstance) GetSourceAccountMoid() string`

GetSourceAccountMoid returns the SourceAccountMoid field if non-nil, zero value otherwise.

### GetSourceAccountMoidOk

`func (o *MgmtConfigBackupInstance) GetSourceAccountMoidOk() (*string, bool)`

GetSourceAccountMoidOk returns a tuple with the SourceAccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountMoid

`func (o *MgmtConfigBackupInstance) SetSourceAccountMoid(v string)`

SetSourceAccountMoid sets SourceAccountMoid field to given value.

### HasSourceAccountMoid

`func (o *MgmtConfigBackupInstance) HasSourceAccountMoid() bool`

HasSourceAccountMoid returns a boolean if a field has been set.

### GetSourceName

`func (o *MgmtConfigBackupInstance) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *MgmtConfigBackupInstance) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *MgmtConfigBackupInstance) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *MgmtConfigBackupInstance) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetStatus

`func (o *MgmtConfigBackupInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MgmtConfigBackupInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MgmtConfigBackupInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MgmtConfigBackupInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemAdminTriggered

`func (o *MgmtConfigBackupInstance) GetSystemAdminTriggered() bool`

GetSystemAdminTriggered returns the SystemAdminTriggered field if non-nil, zero value otherwise.

### GetSystemAdminTriggeredOk

`func (o *MgmtConfigBackupInstance) GetSystemAdminTriggeredOk() (*bool, bool)`

GetSystemAdminTriggeredOk returns a tuple with the SystemAdminTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAdminTriggered

`func (o *MgmtConfigBackupInstance) SetSystemAdminTriggered(v bool)`

SetSystemAdminTriggered sets SystemAdminTriggered field to given value.

### HasSystemAdminTriggered

`func (o *MgmtConfigBackupInstance) HasSystemAdminTriggered() bool`

HasSystemAdminTriggered returns a boolean if a field has been set.

### GetUserPassword

`func (o *MgmtConfigBackupInstance) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *MgmtConfigBackupInstance) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *MgmtConfigBackupInstance) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *MgmtConfigBackupInstance) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetUserPublicKey

`func (o *MgmtConfigBackupInstance) GetUserPublicKey() string`

GetUserPublicKey returns the UserPublicKey field if non-nil, zero value otherwise.

### GetUserPublicKeyOk

`func (o *MgmtConfigBackupInstance) GetUserPublicKeyOk() (*string, bool)`

GetUserPublicKeyOk returns a tuple with the UserPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicKey

`func (o *MgmtConfigBackupInstance) SetUserPublicKey(v string)`

SetUserPublicKey sets UserPublicKey field to given value.

### HasUserPublicKey

`func (o *MgmtConfigBackupInstance) HasUserPublicKey() bool`

HasUserPublicKey returns a boolean if a field has been set.

### GetAccount

`func (o *MgmtConfigBackupInstance) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MgmtConfigBackupInstance) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MgmtConfigBackupInstance) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MgmtConfigBackupInstance) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *MgmtConfigBackupInstance) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *MgmtConfigBackupInstance) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


