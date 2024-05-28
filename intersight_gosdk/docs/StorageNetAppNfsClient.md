# StorageNetAppNfsClient

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
**StorageContainer** | Pointer to [**NullableStorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNfsClient

`func NewStorageNetAppNfsClient(classId string, objectType string, ) *StorageNetAppNfsClient`

NewStorageNetAppNfsClient instantiates a new StorageNetAppNfsClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNfsClientWithDefaults

`func NewStorageNetAppNfsClientWithDefaults() *StorageNetAppNfsClient`

NewStorageNetAppNfsClientWithDefaults instantiates a new StorageNetAppNfsClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNfsClient) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNfsClient) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNfsClient) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNfsClient) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNfsClient) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNfsClient) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientIp

`func (o *StorageNetAppNfsClient) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *StorageNetAppNfsClient) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *StorageNetAppNfsClient) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *StorageNetAppNfsClient) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetIdleDuration

`func (o *StorageNetAppNfsClient) GetIdleDuration() string`

GetIdleDuration returns the IdleDuration field if non-nil, zero value otherwise.

### GetIdleDurationOk

`func (o *StorageNetAppNfsClient) GetIdleDurationOk() (*string, bool)`

GetIdleDurationOk returns a tuple with the IdleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleDuration

`func (o *StorageNetAppNfsClient) SetIdleDuration(v string)`

SetIdleDuration sets IdleDuration field to given value.

### HasIdleDuration

`func (o *StorageNetAppNfsClient) HasIdleDuration() bool`

HasIdleDuration returns a boolean if a field has been set.

### GetProtocol

`func (o *StorageNetAppNfsClient) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageNetAppNfsClient) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageNetAppNfsClient) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageNetAppNfsClient) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServerIp

`func (o *StorageNetAppNfsClient) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *StorageNetAppNfsClient) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *StorageNetAppNfsClient) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *StorageNetAppNfsClient) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetSvmName

`func (o *StorageNetAppNfsClient) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppNfsClient) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppNfsClient) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppNfsClient) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetSvmUuid

`func (o *StorageNetAppNfsClient) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppNfsClient) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppNfsClient) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppNfsClient) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageNetAppNfsClient) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageNetAppNfsClient) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageNetAppNfsClient) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageNetAppNfsClient) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeUuid

`func (o *StorageNetAppNfsClient) GetVolumeUuid() string`

GetVolumeUuid returns the VolumeUuid field if non-nil, zero value otherwise.

### GetVolumeUuidOk

`func (o *StorageNetAppNfsClient) GetVolumeUuidOk() (*string, bool)`

GetVolumeUuidOk returns a tuple with the VolumeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUuid

`func (o *StorageNetAppNfsClient) SetVolumeUuid(v string)`

SetVolumeUuid sets VolumeUuid field to given value.

### HasVolumeUuid

`func (o *StorageNetAppNfsClient) HasVolumeUuid() bool`

HasVolumeUuid returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageNetAppNfsClient) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppNfsClient) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppNfsClient) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppNfsClient) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### SetStorageContainerNil

`func (o *StorageNetAppNfsClient) SetStorageContainerNil(b bool)`

 SetStorageContainerNil sets the value for StorageContainer to be an explicit nil

### UnsetStorageContainer
`func (o *StorageNetAppNfsClient) UnsetStorageContainer()`

UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
### GetTenant

`func (o *StorageNetAppNfsClient) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppNfsClient) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppNfsClient) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppNfsClient) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *StorageNetAppNfsClient) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *StorageNetAppNfsClient) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


