# VmediaMappingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationProtocol** | Pointer to **string** | Type of Authentication protocol when CIFS is used for communication with the remote server. * &#x60;none&#x60; - No authentication is used. * &#x60;ntlm&#x60; - NT LAN Manager (NTLM) security protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2. * &#x60;ntlmi&#x60; - NTLMi security protocol. Use this option only when you enable Digital Signing in the CIFS Windows server. * &#x60;ntlmv2&#x60; - NTLMv2 security protocol. Use this option only with Samba Linux. * &#x60;ntlmv2i&#x60; - NTLMv2i security protocol. Use this option only with Samba Linux. * &#x60;ntlmssp&#x60; - NT LAN Manager Security Support Provider (NTLMSSP) protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2. * &#x60;ntlmsspi&#x60; - NTLMSSPi protocol. Use this option only when you enable Digital Signing in the CIFS Windows server. | [optional] [default to "none"]
**DeviceType** | Pointer to **string** | Type of remote Virtual Media device. * &#x60;cdd&#x60; - Uses compact disc drive as the virtual media mount device. * &#x60;hdd&#x60; - Uses hard disk drive as the virtual media mount device. | [optional] [default to "cdd"]
**HostName** | Pointer to **string** | IP address or hostname of the remote server. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**MountOptions** | Pointer to **string** | Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\\n For NFS, supported options are ro, rw, nolock, noexec, soft, port&#x3D;VALUE, timeo&#x3D;VALUE, retry&#x3D;VALUE.\\n For CIFS, supported options are soft, nounix, noserverino, guest.\\n For CIFS version &lt; 3.0, vers&#x3D;VALUE is mandatory. e.g. vers&#x3D;2.0\\n For HTTP/HTTPS, the only supported option is noauto. | [optional] 
**MountProtocol** | Pointer to **string** | Protocol to use to communicate with the remote server. * &#x60;nfs&#x60; - NFS protocol for vmedia mount. * &#x60;cifs&#x60; - CIFS protocol for vmedia mount. * &#x60;http&#x60; - HTTP protocol for vmedia mount. * &#x60;https&#x60; - HTTPS protocol for vmedia mount. | [optional] [default to "nfs"]
**Password** | Pointer to **string** | Password associated with the username. | [optional] 
**RemoteFile** | Pointer to **string** | Name of the remote file. It should be of .img format for HDD Virtual Media mapping and .iso format for CDD Virtual Media mapping. | [optional] 
**RemotePath** | Pointer to **string** | URL path to the location of the image on the remote server. The preferred format is &#39;/path&#39;. | [optional] 
**Username** | Pointer to **string** | Username to log in to the remote server. | [optional] 
**VolumeName** | Pointer to **string** | Identity of the image for Virtual Media mapping. | [optional] 

## Methods

### NewVmediaMappingAllOf

`func NewVmediaMappingAllOf() *VmediaMappingAllOf`

NewVmediaMappingAllOf instantiates a new VmediaMappingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmediaMappingAllOfWithDefaults

`func NewVmediaMappingAllOfWithDefaults() *VmediaMappingAllOf`

NewVmediaMappingAllOfWithDefaults instantiates a new VmediaMappingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationProtocol

`func (o *VmediaMappingAllOf) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *VmediaMappingAllOf) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *VmediaMappingAllOf) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *VmediaMappingAllOf) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetDeviceType

`func (o *VmediaMappingAllOf) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *VmediaMappingAllOf) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *VmediaMappingAllOf) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *VmediaMappingAllOf) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetHostName

`func (o *VmediaMappingAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VmediaMappingAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VmediaMappingAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *VmediaMappingAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *VmediaMappingAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *VmediaMappingAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *VmediaMappingAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *VmediaMappingAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMountOptions

`func (o *VmediaMappingAllOf) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *VmediaMappingAllOf) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *VmediaMappingAllOf) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *VmediaMappingAllOf) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetMountProtocol

`func (o *VmediaMappingAllOf) GetMountProtocol() string`

GetMountProtocol returns the MountProtocol field if non-nil, zero value otherwise.

### GetMountProtocolOk

`func (o *VmediaMappingAllOf) GetMountProtocolOk() (*string, bool)`

GetMountProtocolOk returns a tuple with the MountProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountProtocol

`func (o *VmediaMappingAllOf) SetMountProtocol(v string)`

SetMountProtocol sets MountProtocol field to given value.

### HasMountProtocol

`func (o *VmediaMappingAllOf) HasMountProtocol() bool`

HasMountProtocol returns a boolean if a field has been set.

### GetPassword

`func (o *VmediaMappingAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VmediaMappingAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VmediaMappingAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VmediaMappingAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRemoteFile

`func (o *VmediaMappingAllOf) GetRemoteFile() string`

GetRemoteFile returns the RemoteFile field if non-nil, zero value otherwise.

### GetRemoteFileOk

`func (o *VmediaMappingAllOf) GetRemoteFileOk() (*string, bool)`

GetRemoteFileOk returns a tuple with the RemoteFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFile

`func (o *VmediaMappingAllOf) SetRemoteFile(v string)`

SetRemoteFile sets RemoteFile field to given value.

### HasRemoteFile

`func (o *VmediaMappingAllOf) HasRemoteFile() bool`

HasRemoteFile returns a boolean if a field has been set.

### GetRemotePath

`func (o *VmediaMappingAllOf) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *VmediaMappingAllOf) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *VmediaMappingAllOf) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *VmediaMappingAllOf) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetUsername

`func (o *VmediaMappingAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VmediaMappingAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VmediaMappingAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VmediaMappingAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVolumeName

`func (o *VmediaMappingAllOf) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VmediaMappingAllOf) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VmediaMappingAllOf) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *VmediaMappingAllOf) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


