# ApplianceBackupBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Backup filename to backup or restore. | [optional] 
**Protocol** | Pointer to **string** | Communication protocol used by the file server (e.g. scp or sftp). * &#x60;scp&#x60; - Secure Copy Protocol (SCP) to access the file server. * &#x60;sftp&#x60; - SSH File Transfer Protocol (SFTP) to access file server. | [optional] [default to "scp"]
**RemoteHost** | Pointer to **string** | Hostname of the remote file server. | [optional] 
**RemotePath** | Pointer to **string** | File server directory to copy the file. | [optional] 
**RemotePort** | Pointer to **int64** | Remote TCP port on the file server (e.g. 22 for scp). | [optional] 
**Username** | Pointer to **string** | Username to authenticate the fileserver. | [optional] 

## Methods

### NewApplianceBackupBase

`func NewApplianceBackupBase() *ApplianceBackupBase`

NewApplianceBackupBase instantiates a new ApplianceBackupBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupBaseWithDefaults

`func NewApplianceBackupBaseWithDefaults() *ApplianceBackupBase`

NewApplianceBackupBaseWithDefaults instantiates a new ApplianceBackupBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *ApplianceBackupBase) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ApplianceBackupBase) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ApplianceBackupBase) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ApplianceBackupBase) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetProtocol

`func (o *ApplianceBackupBase) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplianceBackupBase) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplianceBackupBase) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplianceBackupBase) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteHost

`func (o *ApplianceBackupBase) GetRemoteHost() string`

GetRemoteHost returns the RemoteHost field if non-nil, zero value otherwise.

### GetRemoteHostOk

`func (o *ApplianceBackupBase) GetRemoteHostOk() (*string, bool)`

GetRemoteHostOk returns a tuple with the RemoteHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHost

`func (o *ApplianceBackupBase) SetRemoteHost(v string)`

SetRemoteHost sets RemoteHost field to given value.

### HasRemoteHost

`func (o *ApplianceBackupBase) HasRemoteHost() bool`

HasRemoteHost returns a boolean if a field has been set.

### GetRemotePath

`func (o *ApplianceBackupBase) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *ApplianceBackupBase) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *ApplianceBackupBase) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *ApplianceBackupBase) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetRemotePort

`func (o *ApplianceBackupBase) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *ApplianceBackupBase) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *ApplianceBackupBase) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *ApplianceBackupBase) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetUsername

`func (o *ApplianceBackupBase) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApplianceBackupBase) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApplianceBackupBase) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApplianceBackupBase) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


