# StorageBaseHostLun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hlu** | Pointer to **int64** | Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. | [optional] [readonly] 
**HostName** | Pointer to **string** | Name of the host associated with LUN. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | Name of the storage volume associated with LUN. | [optional] [readonly] 

## Methods

### NewStorageBaseHostLun

`func NewStorageBaseHostLun() *StorageBaseHostLun`

NewStorageBaseHostLun instantiates a new StorageBaseHostLun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseHostLunWithDefaults

`func NewStorageBaseHostLunWithDefaults() *StorageBaseHostLun`

NewStorageBaseHostLunWithDefaults instantiates a new StorageBaseHostLun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHlu

`func (o *StorageBaseHostLun) GetHlu() int64`

GetHlu returns the Hlu field if non-nil, zero value otherwise.

### GetHluOk

`func (o *StorageBaseHostLun) GetHluOk() (*int64, bool)`

GetHluOk returns a tuple with the Hlu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlu

`func (o *StorageBaseHostLun) SetHlu(v int64)`

SetHlu sets Hlu field to given value.

### HasHlu

`func (o *StorageBaseHostLun) HasHlu() bool`

HasHlu returns a boolean if a field has been set.

### GetHostName

`func (o *StorageBaseHostLun) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *StorageBaseHostLun) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *StorageBaseHostLun) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *StorageBaseHostLun) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageBaseHostLun) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageBaseHostLun) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageBaseHostLun) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageBaseHostLun) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


