# FirmwareCifsServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileLocation** | Pointer to **string** | The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. | [optional] 
**MountOptions** | Pointer to **string** | Mount option (Authentication Protocol) as configured on the CIFS Server. Example:ntlmv2. * &#x60;none&#x60; - The default authentication protocol is decided by the endpoint. * &#x60;ntlm&#x60; - The external CIFS server is configured with ntlm as the authentication protocol. * &#x60;ntlmi&#x60; - Mount options of CIFS file server is ntlmi. * &#x60;ntlmv2&#x60; - Mount options of CIFS file server is ntlmv2. * &#x60;ntlmv2i&#x60; - Mount options of CIFS file server is ntlmv2i. * &#x60;ntlmssp&#x60; - Mount options of CIFS file server is ntlmssp. * &#x60;ntlmsspi&#x60; - Mount options of CIFS file server is ntlmsspi. | [optional] [default to "none"]
**RemoteFile** | Pointer to **string** | Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso. | [optional] [readonly] 
**RemoteIp** | Pointer to **string** | CIFS Server Hostname or IP Address. For example:CIFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. | [optional] [readonly] 
**RemoteShare** | Pointer to **string** | Directory where the image is stored. Example:share/subfolder. | [optional] [readonly] 

## Methods

### NewFirmwareCifsServer

`func NewFirmwareCifsServer() *FirmwareCifsServer`

NewFirmwareCifsServer instantiates a new FirmwareCifsServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareCifsServerWithDefaults

`func NewFirmwareCifsServerWithDefaults() *FirmwareCifsServer`

NewFirmwareCifsServerWithDefaults instantiates a new FirmwareCifsServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileLocation

`func (o *FirmwareCifsServer) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *FirmwareCifsServer) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *FirmwareCifsServer) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *FirmwareCifsServer) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetMountOptions

`func (o *FirmwareCifsServer) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *FirmwareCifsServer) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *FirmwareCifsServer) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *FirmwareCifsServer) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetRemoteFile

`func (o *FirmwareCifsServer) GetRemoteFile() string`

GetRemoteFile returns the RemoteFile field if non-nil, zero value otherwise.

### GetRemoteFileOk

`func (o *FirmwareCifsServer) GetRemoteFileOk() (*string, bool)`

GetRemoteFileOk returns a tuple with the RemoteFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFile

`func (o *FirmwareCifsServer) SetRemoteFile(v string)`

SetRemoteFile sets RemoteFile field to given value.

### HasRemoteFile

`func (o *FirmwareCifsServer) HasRemoteFile() bool`

HasRemoteFile returns a boolean if a field has been set.

### GetRemoteIp

`func (o *FirmwareCifsServer) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *FirmwareCifsServer) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *FirmwareCifsServer) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *FirmwareCifsServer) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteShare

`func (o *FirmwareCifsServer) GetRemoteShare() string`

GetRemoteShare returns the RemoteShare field if non-nil, zero value otherwise.

### GetRemoteShareOk

`func (o *FirmwareCifsServer) GetRemoteShareOk() (*string, bool)`

GetRemoteShareOk returns a tuple with the RemoteShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteShare

`func (o *FirmwareCifsServer) SetRemoteShare(v string)`

SetRemoteShare sets RemoteShare field to given value.

### HasRemoteShare

`func (o *FirmwareCifsServer) HasRemoteShare() bool`

HasRemoteShare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


