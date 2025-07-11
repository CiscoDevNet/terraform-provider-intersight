# HciEsxiVmDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.EsxiVmDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.EsxiVmDisk"]
**BusType** | Pointer to **string** | The bus type of the disk. Possible values are &#39;IDE&#39;, &#39;SCSI&#39;, &#39;SATA&#39;, &#39;NVME&#39;. | [optional] [readonly] 
**DiskExtId** | Pointer to **string** | The unique identifier of the disk. | [optional] [readonly] 
**DiskSizeBytes** | Pointer to **int64** | The size of the disk in bytes. | [optional] [readonly] 
**Index** | Pointer to **int32** | The index of the disk, similar to a slot number on physical machine. | [optional] [readonly] 
**IsFlashModeEnabled** | Pointer to **bool** | Indicates whether the virtual disk is pinned to the hot tier or not. | [optional] [readonly] 
**StorageContainerExtId** | Pointer to **string** | The extId of the storage container which backs this disk. | [optional] [readonly] 
**VmExtId** | Pointer to **string** | The unique identifier of the VM. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vm** | Pointer to [**NullableHciEsxiVmRelationship**](HciEsxiVmRelationship.md) |  | [optional] 

## Methods

### NewHciEsxiVmDisk

`func NewHciEsxiVmDisk(classId string, objectType string, ) *HciEsxiVmDisk`

NewHciEsxiVmDisk instantiates a new HciEsxiVmDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciEsxiVmDiskWithDefaults

`func NewHciEsxiVmDiskWithDefaults() *HciEsxiVmDisk`

NewHciEsxiVmDiskWithDefaults instantiates a new HciEsxiVmDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciEsxiVmDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciEsxiVmDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciEsxiVmDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciEsxiVmDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciEsxiVmDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciEsxiVmDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBusType

`func (o *HciEsxiVmDisk) GetBusType() string`

GetBusType returns the BusType field if non-nil, zero value otherwise.

### GetBusTypeOk

`func (o *HciEsxiVmDisk) GetBusTypeOk() (*string, bool)`

GetBusTypeOk returns a tuple with the BusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusType

`func (o *HciEsxiVmDisk) SetBusType(v string)`

SetBusType sets BusType field to given value.

### HasBusType

`func (o *HciEsxiVmDisk) HasBusType() bool`

HasBusType returns a boolean if a field has been set.

### GetDiskExtId

`func (o *HciEsxiVmDisk) GetDiskExtId() string`

GetDiskExtId returns the DiskExtId field if non-nil, zero value otherwise.

### GetDiskExtIdOk

`func (o *HciEsxiVmDisk) GetDiskExtIdOk() (*string, bool)`

GetDiskExtIdOk returns a tuple with the DiskExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskExtId

`func (o *HciEsxiVmDisk) SetDiskExtId(v string)`

SetDiskExtId sets DiskExtId field to given value.

### HasDiskExtId

`func (o *HciEsxiVmDisk) HasDiskExtId() bool`

HasDiskExtId returns a boolean if a field has been set.

### GetDiskSizeBytes

`func (o *HciEsxiVmDisk) GetDiskSizeBytes() int64`

GetDiskSizeBytes returns the DiskSizeBytes field if non-nil, zero value otherwise.

### GetDiskSizeBytesOk

`func (o *HciEsxiVmDisk) GetDiskSizeBytesOk() (*int64, bool)`

GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeBytes

`func (o *HciEsxiVmDisk) SetDiskSizeBytes(v int64)`

SetDiskSizeBytes sets DiskSizeBytes field to given value.

### HasDiskSizeBytes

`func (o *HciEsxiVmDisk) HasDiskSizeBytes() bool`

HasDiskSizeBytes returns a boolean if a field has been set.

### GetIndex

`func (o *HciEsxiVmDisk) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *HciEsxiVmDisk) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *HciEsxiVmDisk) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *HciEsxiVmDisk) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIsFlashModeEnabled

`func (o *HciEsxiVmDisk) GetIsFlashModeEnabled() bool`

GetIsFlashModeEnabled returns the IsFlashModeEnabled field if non-nil, zero value otherwise.

### GetIsFlashModeEnabledOk

`func (o *HciEsxiVmDisk) GetIsFlashModeEnabledOk() (*bool, bool)`

GetIsFlashModeEnabledOk returns a tuple with the IsFlashModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlashModeEnabled

`func (o *HciEsxiVmDisk) SetIsFlashModeEnabled(v bool)`

SetIsFlashModeEnabled sets IsFlashModeEnabled field to given value.

### HasIsFlashModeEnabled

`func (o *HciEsxiVmDisk) HasIsFlashModeEnabled() bool`

HasIsFlashModeEnabled returns a boolean if a field has been set.

### GetStorageContainerExtId

`func (o *HciEsxiVmDisk) GetStorageContainerExtId() string`

GetStorageContainerExtId returns the StorageContainerExtId field if non-nil, zero value otherwise.

### GetStorageContainerExtIdOk

`func (o *HciEsxiVmDisk) GetStorageContainerExtIdOk() (*string, bool)`

GetStorageContainerExtIdOk returns a tuple with the StorageContainerExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerExtId

`func (o *HciEsxiVmDisk) SetStorageContainerExtId(v string)`

SetStorageContainerExtId sets StorageContainerExtId field to given value.

### HasStorageContainerExtId

`func (o *HciEsxiVmDisk) HasStorageContainerExtId() bool`

HasStorageContainerExtId returns a boolean if a field has been set.

### GetVmExtId

`func (o *HciEsxiVmDisk) GetVmExtId() string`

GetVmExtId returns the VmExtId field if non-nil, zero value otherwise.

### GetVmExtIdOk

`func (o *HciEsxiVmDisk) GetVmExtIdOk() (*string, bool)`

GetVmExtIdOk returns a tuple with the VmExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmExtId

`func (o *HciEsxiVmDisk) SetVmExtId(v string)`

SetVmExtId sets VmExtId field to given value.

### HasVmExtId

`func (o *HciEsxiVmDisk) HasVmExtId() bool`

HasVmExtId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HciEsxiVmDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciEsxiVmDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciEsxiVmDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciEsxiVmDisk) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciEsxiVmDisk) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciEsxiVmDisk) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVm

`func (o *HciEsxiVmDisk) GetVm() HciEsxiVmRelationship`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *HciEsxiVmDisk) GetVmOk() (*HciEsxiVmRelationship, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *HciEsxiVmDisk) SetVm(v HciEsxiVmRelationship)`

SetVm sets Vm field to given value.

### HasVm

`func (o *HciEsxiVmDisk) HasVm() bool`

HasVm returns a boolean if a field has been set.

### SetVmNil

`func (o *HciEsxiVmDisk) SetVmNil(b bool)`

 SetVmNil sets the value for Vm to be an explicit nil

### UnsetVm
`func (o *HciEsxiVmDisk) UnsetVm()`

UnsetVm ensures that no value is present for Vm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


