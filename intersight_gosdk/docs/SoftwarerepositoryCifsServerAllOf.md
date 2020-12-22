# SoftwarerepositoryCifsServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.CifsServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.CifsServer"]
**FileLocation** | Pointer to **string** | The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**MountOption** | Pointer to **string** | For CIFS, leave the field blank or enter one or more comma seperated options from the following. For Example, \&quot; \&quot; , \&quot; soft \&quot; , \&quot; soft , nounix \&quot; . * soft. * nounix. * noserviceino. * guest. * USERNAME&#x3D;VALUE. * PASSWORD&#x3D;VALUE. * sec&#x3D;VALUE (VALUE could be None, Ntlm, Ntlmi, Ntlmssp, Ntlmsspi, Ntlmv2, Ntlmv2i). | [optional] 
**Password** | Pointer to **string** | Password configured on the CIFS server. | [optional] 
**RemoteFile** | Pointer to **string** | Filename of the image in the CIFS server. For example:ucs-c220m5-huu-3.1.2c.iso. | [optional] [readonly] 
**RemoteIp** | Pointer to **string** | Hostname or IP Address of the CIFS server. | [optional] [readonly] 
**RemoteShare** | Pointer to **string** | Remote directory where the image is present. For example:/share/subfolder. | [optional] [readonly] 
**Username** | Pointer to **string** | Username configured on the CIFS server. | [optional] 

## Methods

### NewSoftwarerepositoryCifsServerAllOf

`func NewSoftwarerepositoryCifsServerAllOf(classId string, objectType string, ) *SoftwarerepositoryCifsServerAllOf`

NewSoftwarerepositoryCifsServerAllOf instantiates a new SoftwarerepositoryCifsServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCifsServerAllOfWithDefaults

`func NewSoftwarerepositoryCifsServerAllOfWithDefaults() *SoftwarerepositoryCifsServerAllOf`

NewSoftwarerepositoryCifsServerAllOfWithDefaults instantiates a new SoftwarerepositoryCifsServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryCifsServerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryCifsServerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryCifsServerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCifsServerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileLocation

`func (o *SoftwarerepositoryCifsServerAllOf) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *SoftwarerepositoryCifsServerAllOf) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *SoftwarerepositoryCifsServerAllOf) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *SoftwarerepositoryCifsServerAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *SoftwarerepositoryCifsServerAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *SoftwarerepositoryCifsServerAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMountOption

`func (o *SoftwarerepositoryCifsServerAllOf) GetMountOption() string`

GetMountOption returns the MountOption field if non-nil, zero value otherwise.

### GetMountOptionOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetMountOptionOk() (*string, bool)`

GetMountOptionOk returns a tuple with the MountOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOption

`func (o *SoftwarerepositoryCifsServerAllOf) SetMountOption(v string)`

SetMountOption sets MountOption field to given value.

### HasMountOption

`func (o *SoftwarerepositoryCifsServerAllOf) HasMountOption() bool`

HasMountOption returns a boolean if a field has been set.

### GetPassword

`func (o *SoftwarerepositoryCifsServerAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SoftwarerepositoryCifsServerAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SoftwarerepositoryCifsServerAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRemoteFile

`func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteFile() string`

GetRemoteFile returns the RemoteFile field if non-nil, zero value otherwise.

### GetRemoteFileOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteFileOk() (*string, bool)`

GetRemoteFileOk returns a tuple with the RemoteFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFile

`func (o *SoftwarerepositoryCifsServerAllOf) SetRemoteFile(v string)`

SetRemoteFile sets RemoteFile field to given value.

### HasRemoteFile

`func (o *SoftwarerepositoryCifsServerAllOf) HasRemoteFile() bool`

HasRemoteFile returns a boolean if a field has been set.

### GetRemoteIp

`func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *SoftwarerepositoryCifsServerAllOf) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *SoftwarerepositoryCifsServerAllOf) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteShare

`func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteShare() string`

GetRemoteShare returns the RemoteShare field if non-nil, zero value otherwise.

### GetRemoteShareOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteShareOk() (*string, bool)`

GetRemoteShareOk returns a tuple with the RemoteShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteShare

`func (o *SoftwarerepositoryCifsServerAllOf) SetRemoteShare(v string)`

SetRemoteShare sets RemoteShare field to given value.

### HasRemoteShare

`func (o *SoftwarerepositoryCifsServerAllOf) HasRemoteShare() bool`

HasRemoteShare returns a boolean if a field has been set.

### GetUsername

`func (o *SoftwarerepositoryCifsServerAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SoftwarerepositoryCifsServerAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SoftwarerepositoryCifsServerAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SoftwarerepositoryCifsServerAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


