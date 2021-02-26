# VmediaMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vmedia.Mapping"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vmedia.Mapping"]
**AuthenticationProtocol** | Pointer to **string** | Type of Authentication protocol when CIFS is used for communication with the remote server. * &#x60;none&#x60; - No authentication is used. * &#x60;ntlm&#x60; - NT LAN Manager (NTLM) security protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2. * &#x60;ntlmi&#x60; - NTLMi security protocol. Use this option only when you enable Digital Signing in the CIFS Windows server. * &#x60;ntlmv2&#x60; - NTLMv2 security protocol. Use this option only with Samba Linux. * &#x60;ntlmv2i&#x60; - NTLMv2i security protocol. Use this option only with Samba Linux. * &#x60;ntlmssp&#x60; - NT LAN Manager Security Support Provider (NTLMSSP) protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2. * &#x60;ntlmsspi&#x60; - NTLMSSPi protocol. Use this option only when you enable Digital Signing in the CIFS Windows server. | [optional] [default to "none"]
**DeviceType** | Pointer to **string** | Type of remote Virtual Media device. * &#x60;cdd&#x60; - Uses compact disc drive as the virtual media mount device. * &#x60;hdd&#x60; - Uses hard disk drive as the virtual media mount device. | [optional] [default to "cdd"]
**FileLocation** | Pointer to **string** | Remote location of image. Preferred format is &#39;hostname/filePath/fileName&#39;. | [optional] 
**HostName** | Pointer to **string** | IP address or hostname of the remote server. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**MountOptions** | Pointer to **string** | Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\\n For NFS, supported options are ro, rw, nolock, noexec, soft, port&#x3D;VALUE, timeo&#x3D;VALUE, retry&#x3D;VALUE.\\n For CIFS, supported options are soft, nounix, noserverino, guest.\\n For CIFS version &lt; 3.0, vers&#x3D;VALUE is mandatory. e.g. vers&#x3D;2.0\\n For HTTP/HTTPS, the only supported option is noauto. | [optional] 
**MountProtocol** | Pointer to **string** | Protocol to use to communicate with the remote server. * &#x60;nfs&#x60; - NFS protocol for vmedia mount. * &#x60;cifs&#x60; - CIFS protocol for vmedia mount. * &#x60;http&#x60; - HTTP protocol for vmedia mount. * &#x60;https&#x60; - HTTPS protocol for vmedia mount. | [optional] [default to "nfs"]
**Password** | Pointer to **string** | Password associated with the username. | [optional] 
**RemoteFile** | Pointer to **string** | The remote file location path for the virtual media mapping. Accepted formats are: HDD for CIFS/NFS: hostname-or-IP/filePath/fileName.img. CDD for CIFS/NFS: hostname-or-IP/filePath/fileName.iso. HDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.img. CDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.iso. | [optional] 
**RemotePath** | Pointer to **string** | URL path to the location of the image on the remote server. The preferred format is &#39;/path&#39;. | [optional] 
**SanitizedFileLocation** | Pointer to **string** | File Location in standard format &#39;hostname/filePath/fileName&#39;. This field should be used to calculate config drift. User input format may vary while inventory will return data in format in compliance with mount option for the mount. Both will be converged to this standard format for comparison. | [optional] [readonly] 
**Username** | Pointer to **string** | Username to log in to the remote server. | [optional] 
**VolumeName** | Pointer to **string** | Identity of the image for Virtual Media mapping. | [optional] 

## Methods

### NewVmediaMapping

`func NewVmediaMapping(classId string, objectType string, ) *VmediaMapping`

NewVmediaMapping instantiates a new VmediaMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmediaMappingWithDefaults

`func NewVmediaMappingWithDefaults() *VmediaMapping`

NewVmediaMappingWithDefaults instantiates a new VmediaMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VmediaMapping) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VmediaMapping) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VmediaMapping) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VmediaMapping) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VmediaMapping) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VmediaMapping) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthenticationProtocol

`func (o *VmediaMapping) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *VmediaMapping) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *VmediaMapping) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *VmediaMapping) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetDeviceType

`func (o *VmediaMapping) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *VmediaMapping) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *VmediaMapping) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *VmediaMapping) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetFileLocation

`func (o *VmediaMapping) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *VmediaMapping) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *VmediaMapping) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *VmediaMapping) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetHostName

`func (o *VmediaMapping) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VmediaMapping) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VmediaMapping) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *VmediaMapping) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *VmediaMapping) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *VmediaMapping) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *VmediaMapping) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *VmediaMapping) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMountOptions

`func (o *VmediaMapping) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *VmediaMapping) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *VmediaMapping) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *VmediaMapping) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetMountProtocol

`func (o *VmediaMapping) GetMountProtocol() string`

GetMountProtocol returns the MountProtocol field if non-nil, zero value otherwise.

### GetMountProtocolOk

`func (o *VmediaMapping) GetMountProtocolOk() (*string, bool)`

GetMountProtocolOk returns a tuple with the MountProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountProtocol

`func (o *VmediaMapping) SetMountProtocol(v string)`

SetMountProtocol sets MountProtocol field to given value.

### HasMountProtocol

`func (o *VmediaMapping) HasMountProtocol() bool`

HasMountProtocol returns a boolean if a field has been set.

### GetPassword

`func (o *VmediaMapping) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VmediaMapping) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VmediaMapping) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VmediaMapping) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRemoteFile

`func (o *VmediaMapping) GetRemoteFile() string`

GetRemoteFile returns the RemoteFile field if non-nil, zero value otherwise.

### GetRemoteFileOk

`func (o *VmediaMapping) GetRemoteFileOk() (*string, bool)`

GetRemoteFileOk returns a tuple with the RemoteFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFile

`func (o *VmediaMapping) SetRemoteFile(v string)`

SetRemoteFile sets RemoteFile field to given value.

### HasRemoteFile

`func (o *VmediaMapping) HasRemoteFile() bool`

HasRemoteFile returns a boolean if a field has been set.

### GetRemotePath

`func (o *VmediaMapping) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *VmediaMapping) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *VmediaMapping) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *VmediaMapping) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetSanitizedFileLocation

`func (o *VmediaMapping) GetSanitizedFileLocation() string`

GetSanitizedFileLocation returns the SanitizedFileLocation field if non-nil, zero value otherwise.

### GetSanitizedFileLocationOk

`func (o *VmediaMapping) GetSanitizedFileLocationOk() (*string, bool)`

GetSanitizedFileLocationOk returns a tuple with the SanitizedFileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitizedFileLocation

`func (o *VmediaMapping) SetSanitizedFileLocation(v string)`

SetSanitizedFileLocation sets SanitizedFileLocation field to given value.

### HasSanitizedFileLocation

`func (o *VmediaMapping) HasSanitizedFileLocation() bool`

HasSanitizedFileLocation returns a boolean if a field has been set.

### GetUsername

`func (o *VmediaMapping) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VmediaMapping) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VmediaMapping) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VmediaMapping) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVolumeName

`func (o *VmediaMapping) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VmediaMapping) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VmediaMapping) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *VmediaMapping) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


