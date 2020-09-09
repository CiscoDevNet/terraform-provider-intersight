# FirmwareCifsServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileLocation** | Pointer to **string** | The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. | [optional] 
**MountOptions** | Pointer to **string** | Mount option (Authentication Protocol) as configured on the CIFS Server. Example:ntlmv2. * &#x60;none&#x60; - The default authentication protocol is decided by the endpoint. * &#x60;ntlm&#x60; - The external CIFS server is configured with ntlm as the authentication protocol. * &#x60;ntlmi&#x60; - Mount options of CIFS file server is ntlmi. * &#x60;ntlmv2&#x60; - Mount options of CIFS file server is ntlmv2. * &#x60;ntlmv2i&#x60; - Mount options of CIFS file server is ntlmv2i. * &#x60;ntlmssp&#x60; - Mount options of CIFS file server is ntlmssp. * &#x60;ntlmsspi&#x60; - Mount options of CIFS file server is ntlmsspi. | [optional] [default to "none"]
**RemoteFile** | Pointer to **string** | Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso. | [optional] [readonly] 
**RemoteIp** | Pointer to **string** | CIFS Server Hostname or IP Address. For example:CIFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. | [optional] [readonly] 
**RemoteShare** | Pointer to **string** | Directory where the image is stored. Example:share/subfolder. | [optional] [readonly] 

## Methods

### NewFirmwareCifsServerAllOf

`func NewFirmwareCifsServerAllOf() *FirmwareCifsServerAllOf`

NewFirmwareCifsServerAllOf instantiates a new FirmwareCifsServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareCifsServerAllOfWithDefaults

`func NewFirmwareCifsServerAllOfWithDefaults() *FirmwareCifsServerAllOf`

NewFirmwareCifsServerAllOfWithDefaults instantiates a new FirmwareCifsServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileLocation

`func (o *FirmwareCifsServerAllOf) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *FirmwareCifsServerAllOf) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *FirmwareCifsServerAllOf) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *FirmwareCifsServerAllOf) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetMountOptions

`func (o *FirmwareCifsServerAllOf) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *FirmwareCifsServerAllOf) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *FirmwareCifsServerAllOf) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *FirmwareCifsServerAllOf) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetRemoteFile

`func (o *FirmwareCifsServerAllOf) GetRemoteFile() string`

GetRemoteFile returns the RemoteFile field if non-nil, zero value otherwise.

### GetRemoteFileOk

`func (o *FirmwareCifsServerAllOf) GetRemoteFileOk() (*string, bool)`

GetRemoteFileOk returns a tuple with the RemoteFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFile

`func (o *FirmwareCifsServerAllOf) SetRemoteFile(v string)`

SetRemoteFile sets RemoteFile field to given value.

### HasRemoteFile

`func (o *FirmwareCifsServerAllOf) HasRemoteFile() bool`

HasRemoteFile returns a boolean if a field has been set.

### GetRemoteIp

`func (o *FirmwareCifsServerAllOf) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *FirmwareCifsServerAllOf) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *FirmwareCifsServerAllOf) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *FirmwareCifsServerAllOf) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteShare

`func (o *FirmwareCifsServerAllOf) GetRemoteShare() string`

GetRemoteShare returns the RemoteShare field if non-nil, zero value otherwise.

### GetRemoteShareOk

`func (o *FirmwareCifsServerAllOf) GetRemoteShareOk() (*string, bool)`

GetRemoteShareOk returns a tuple with the RemoteShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteShare

`func (o *FirmwareCifsServerAllOf) SetRemoteShare(v string)`

SetRemoteShare sets RemoteShare field to given value.

### HasRemoteShare

`func (o *FirmwareCifsServerAllOf) HasRemoteShare() bool`

HasRemoteShare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


