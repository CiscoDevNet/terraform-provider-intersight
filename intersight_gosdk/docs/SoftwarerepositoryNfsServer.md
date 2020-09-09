# SoftwarerepositoryNfsServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileLocation** | Pointer to **string** | The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. | [optional] 
**MountOptions** | Pointer to **string** | For NFS, leave the field blank or enter one or more comma seperated options from the following.For Example, \&quot; \&quot; , \&quot; ro \&quot; , \&quot; ro , rw \&quot; . * ro. * rw. * nolock. * noexec. * soft. * PORT&#x3D;VALUE. * timeo&#x3D;VALUE. * retry&#x3D;VALUE. | [optional] [readonly] 
**RemoteFile** | Pointer to **string** | Filename of the image in the NFS server. For example:ucs-c220m5-huu-3.1.2c.iso. | [optional] [readonly] 
**RemoteIp** | Pointer to **string** | Hostname or IP Address of the NFS server. | [optional] [readonly] 
**RemoteShare** | Pointer to **string** | Remote directory where the image is present. For example:/share/subfolder. | [optional] [readonly] 

## Methods

### NewSoftwarerepositoryNfsServer

`func NewSoftwarerepositoryNfsServer() *SoftwarerepositoryNfsServer`

NewSoftwarerepositoryNfsServer instantiates a new SoftwarerepositoryNfsServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryNfsServerWithDefaults

`func NewSoftwarerepositoryNfsServerWithDefaults() *SoftwarerepositoryNfsServer`

NewSoftwarerepositoryNfsServerWithDefaults instantiates a new SoftwarerepositoryNfsServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileLocation

`func (o *SoftwarerepositoryNfsServer) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *SoftwarerepositoryNfsServer) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *SoftwarerepositoryNfsServer) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *SoftwarerepositoryNfsServer) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetMountOptions

`func (o *SoftwarerepositoryNfsServer) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *SoftwarerepositoryNfsServer) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *SoftwarerepositoryNfsServer) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *SoftwarerepositoryNfsServer) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetRemoteFile

`func (o *SoftwarerepositoryNfsServer) GetRemoteFile() string`

GetRemoteFile returns the RemoteFile field if non-nil, zero value otherwise.

### GetRemoteFileOk

`func (o *SoftwarerepositoryNfsServer) GetRemoteFileOk() (*string, bool)`

GetRemoteFileOk returns a tuple with the RemoteFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFile

`func (o *SoftwarerepositoryNfsServer) SetRemoteFile(v string)`

SetRemoteFile sets RemoteFile field to given value.

### HasRemoteFile

`func (o *SoftwarerepositoryNfsServer) HasRemoteFile() bool`

HasRemoteFile returns a boolean if a field has been set.

### GetRemoteIp

`func (o *SoftwarerepositoryNfsServer) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *SoftwarerepositoryNfsServer) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *SoftwarerepositoryNfsServer) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *SoftwarerepositoryNfsServer) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteShare

`func (o *SoftwarerepositoryNfsServer) GetRemoteShare() string`

GetRemoteShare returns the RemoteShare field if non-nil, zero value otherwise.

### GetRemoteShareOk

`func (o *SoftwarerepositoryNfsServer) GetRemoteShareOk() (*string, bool)`

GetRemoteShareOk returns a tuple with the RemoteShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteShare

`func (o *SoftwarerepositoryNfsServer) SetRemoteShare(v string)`

SetRemoteShare sets RemoteShare field to given value.

### HasRemoteShare

`func (o *SoftwarerepositoryNfsServer) HasRemoteShare() bool`

HasRemoteShare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


