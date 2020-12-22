# RecoveryAbstractBackupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**FileNamePrefix** | Pointer to **string** | The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**LocationType** | Pointer to **string** | Specifies whether the backup will be stored locally or remotely. * &#x60;Network Share&#x60; - The backup is stored remotely on a separate server. * &#x60;Local Storage&#x60; - The backup is stored locally on the endpoint. | [optional] [default to "Network Share"]
**Password** | Pointer to **string** | Password of Backup server. | [optional] 
**Path** | Pointer to **string** | The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/. | [optional] 
**Protocol** | Pointer to **string** | Protocol for transferring the backup image to the network share location. * &#x60;SCP&#x60; - Secure Copy Protocol (SCP) to access the file server. * &#x60;SFTP&#x60; - SSH File Transfer Protocol (SFTP) to access file server. * &#x60;FTP&#x60; - File Transfer Protocol (FTP) to access file server. | [optional] [default to "SCP"]
**RetentionCount** | Pointer to **int64** | Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner. | [optional] [default to 10]
**UserName** | Pointer to **string** | Username for the backup server. | [optional] 

## Methods

### NewRecoveryAbstractBackupConfig

`func NewRecoveryAbstractBackupConfig(classId string, objectType string, ) *RecoveryAbstractBackupConfig`

NewRecoveryAbstractBackupConfig instantiates a new RecoveryAbstractBackupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryAbstractBackupConfigWithDefaults

`func NewRecoveryAbstractBackupConfigWithDefaults() *RecoveryAbstractBackupConfig`

NewRecoveryAbstractBackupConfigWithDefaults instantiates a new RecoveryAbstractBackupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryAbstractBackupConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryAbstractBackupConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryAbstractBackupConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryAbstractBackupConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryAbstractBackupConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryAbstractBackupConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileNamePrefix

`func (o *RecoveryAbstractBackupConfig) GetFileNamePrefix() string`

GetFileNamePrefix returns the FileNamePrefix field if non-nil, zero value otherwise.

### GetFileNamePrefixOk

`func (o *RecoveryAbstractBackupConfig) GetFileNamePrefixOk() (*string, bool)`

GetFileNamePrefixOk returns a tuple with the FileNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePrefix

`func (o *RecoveryAbstractBackupConfig) SetFileNamePrefix(v string)`

SetFileNamePrefix sets FileNamePrefix field to given value.

### HasFileNamePrefix

`func (o *RecoveryAbstractBackupConfig) HasFileNamePrefix() bool`

HasFileNamePrefix returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *RecoveryAbstractBackupConfig) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *RecoveryAbstractBackupConfig) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *RecoveryAbstractBackupConfig) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *RecoveryAbstractBackupConfig) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetLocationType

`func (o *RecoveryAbstractBackupConfig) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *RecoveryAbstractBackupConfig) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *RecoveryAbstractBackupConfig) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *RecoveryAbstractBackupConfig) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetPassword

`func (o *RecoveryAbstractBackupConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RecoveryAbstractBackupConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RecoveryAbstractBackupConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RecoveryAbstractBackupConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *RecoveryAbstractBackupConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RecoveryAbstractBackupConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RecoveryAbstractBackupConfig) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RecoveryAbstractBackupConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocol

`func (o *RecoveryAbstractBackupConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RecoveryAbstractBackupConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RecoveryAbstractBackupConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RecoveryAbstractBackupConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRetentionCount

`func (o *RecoveryAbstractBackupConfig) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *RecoveryAbstractBackupConfig) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *RecoveryAbstractBackupConfig) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *RecoveryAbstractBackupConfig) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetUserName

`func (o *RecoveryAbstractBackupConfig) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RecoveryAbstractBackupConfig) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RecoveryAbstractBackupConfig) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *RecoveryAbstractBackupConfig) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


