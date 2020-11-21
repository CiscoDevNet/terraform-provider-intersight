# ApplianceBackupBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Filename** | Pointer to **string** | Backup filename to backup or restore. | [optional] 
**Protocol** | Pointer to **string** | Communication protocol used by the file server (e.g. scp or sftp). * &#x60;scp&#x60; - Secure Copy Protocol (SCP) to access the file server. * &#x60;sftp&#x60; - SSH File Transfer Protocol (SFTP) to access file server. | [optional] [default to "scp"]
**RemoteHost** | Pointer to **string** | Hostname of the remote file server. | [optional] 
**RemotePath** | Pointer to **string** | File server directory to copy the file. | [optional] 
**RemotePort** | Pointer to **int64** | Remote TCP port on the file server (e.g. 22 for scp). | [optional] 
**Username** | Pointer to **string** | Username to authenticate the fileserver. | [optional] 

## Methods

### NewApplianceBackupBaseAllOf

`func NewApplianceBackupBaseAllOf(classId string, objectType string, ) *ApplianceBackupBaseAllOf`

NewApplianceBackupBaseAllOf instantiates a new ApplianceBackupBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupBaseAllOfWithDefaults

`func NewApplianceBackupBaseAllOfWithDefaults() *ApplianceBackupBaseAllOf`

NewApplianceBackupBaseAllOfWithDefaults instantiates a new ApplianceBackupBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupBaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupBaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupBaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupBaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupBaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupBaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilename

`func (o *ApplianceBackupBaseAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ApplianceBackupBaseAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ApplianceBackupBaseAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ApplianceBackupBaseAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetProtocol

`func (o *ApplianceBackupBaseAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplianceBackupBaseAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplianceBackupBaseAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplianceBackupBaseAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteHost

`func (o *ApplianceBackupBaseAllOf) GetRemoteHost() string`

GetRemoteHost returns the RemoteHost field if non-nil, zero value otherwise.

### GetRemoteHostOk

`func (o *ApplianceBackupBaseAllOf) GetRemoteHostOk() (*string, bool)`

GetRemoteHostOk returns a tuple with the RemoteHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHost

`func (o *ApplianceBackupBaseAllOf) SetRemoteHost(v string)`

SetRemoteHost sets RemoteHost field to given value.

### HasRemoteHost

`func (o *ApplianceBackupBaseAllOf) HasRemoteHost() bool`

HasRemoteHost returns a boolean if a field has been set.

### GetRemotePath

`func (o *ApplianceBackupBaseAllOf) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *ApplianceBackupBaseAllOf) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *ApplianceBackupBaseAllOf) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *ApplianceBackupBaseAllOf) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetRemotePort

`func (o *ApplianceBackupBaseAllOf) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *ApplianceBackupBaseAllOf) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *ApplianceBackupBaseAllOf) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *ApplianceBackupBaseAllOf) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetUsername

`func (o *ApplianceBackupBaseAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApplianceBackupBaseAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApplianceBackupBaseAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApplianceBackupBaseAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


