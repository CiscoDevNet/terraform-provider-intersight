# StorageNetAppNfsClientAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNfsClient"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNfsClient"]
**ClientIp** | Pointer to **string** | Specifies IP address of the client. | [optional] [readonly] 
**IdleDuration** | Pointer to **string** | Specifies an ISO-8601 format of date and time to retrieve the idle time duration in hours, minutes, and seconds format. | [optional] [readonly] 
**Protocol** | Pointer to **string** | The NFS protocol version over which client is accessing the volume. | [optional] [readonly] 
**ServerIp** | Pointer to **string** | Specifies IP address of the server. | [optional] [readonly] 
**SvmName** | Pointer to **string** | The storage virtual machine name for the NFS client. | [optional] [readonly] 
**SvmUuid** | Pointer to **string** | Unique identifier for the NetApp Storage Virtual Machine. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | The parent volume name for the NFS client. | [optional] [readonly] 
**VolumeUuid** | Pointer to **string** | Unique identifier for the NetApp Volume. | [optional] [readonly] 
**StorageContainer** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNfsClientAllOf

`func NewStorageNetAppNfsClientAllOf(classId string, objectType string, ) *StorageNetAppNfsClientAllOf`

NewStorageNetAppNfsClientAllOf instantiates a new StorageNetAppNfsClientAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNfsClientAllOfWithDefaults

`func NewStorageNetAppNfsClientAllOfWithDefaults() *StorageNetAppNfsClientAllOf`

NewStorageNetAppNfsClientAllOfWithDefaults instantiates a new StorageNetAppNfsClientAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNfsClientAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNfsClientAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNfsClientAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNfsClientAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNfsClientAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNfsClientAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientIp

`func (o *StorageNetAppNfsClientAllOf) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *StorageNetAppNfsClientAllOf) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *StorageNetAppNfsClientAllOf) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *StorageNetAppNfsClientAllOf) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetIdleDuration

`func (o *StorageNetAppNfsClientAllOf) GetIdleDuration() string`

GetIdleDuration returns the IdleDuration field if non-nil, zero value otherwise.

### GetIdleDurationOk

`func (o *StorageNetAppNfsClientAllOf) GetIdleDurationOk() (*string, bool)`

GetIdleDurationOk returns a tuple with the IdleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleDuration

`func (o *StorageNetAppNfsClientAllOf) SetIdleDuration(v string)`

SetIdleDuration sets IdleDuration field to given value.

### HasIdleDuration

`func (o *StorageNetAppNfsClientAllOf) HasIdleDuration() bool`

HasIdleDuration returns a boolean if a field has been set.

### GetProtocol

`func (o *StorageNetAppNfsClientAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageNetAppNfsClientAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageNetAppNfsClientAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageNetAppNfsClientAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServerIp

`func (o *StorageNetAppNfsClientAllOf) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *StorageNetAppNfsClientAllOf) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *StorageNetAppNfsClientAllOf) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *StorageNetAppNfsClientAllOf) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetSvmName

`func (o *StorageNetAppNfsClientAllOf) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppNfsClientAllOf) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppNfsClientAllOf) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppNfsClientAllOf) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetSvmUuid

`func (o *StorageNetAppNfsClientAllOf) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppNfsClientAllOf) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppNfsClientAllOf) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppNfsClientAllOf) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageNetAppNfsClientAllOf) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageNetAppNfsClientAllOf) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageNetAppNfsClientAllOf) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageNetAppNfsClientAllOf) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeUuid

`func (o *StorageNetAppNfsClientAllOf) GetVolumeUuid() string`

GetVolumeUuid returns the VolumeUuid field if non-nil, zero value otherwise.

### GetVolumeUuidOk

`func (o *StorageNetAppNfsClientAllOf) GetVolumeUuidOk() (*string, bool)`

GetVolumeUuidOk returns a tuple with the VolumeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUuid

`func (o *StorageNetAppNfsClientAllOf) SetVolumeUuid(v string)`

SetVolumeUuid sets VolumeUuid field to given value.

### HasVolumeUuid

`func (o *StorageNetAppNfsClientAllOf) HasVolumeUuid() bool`

HasVolumeUuid returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageNetAppNfsClientAllOf) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppNfsClientAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppNfsClientAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppNfsClientAllOf) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppNfsClientAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppNfsClientAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppNfsClientAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppNfsClientAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


