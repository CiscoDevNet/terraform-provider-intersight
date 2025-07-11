# HciAhvVmDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.AhvVmDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.AhvVmDisk"]
**AdsfVolumeGroupExtId** | Pointer to **string** | The extId of the volume group which backs this disk for block access. ADSF (Acropolis Distributed Storage  Fabric) is the storage layer of the Nutanix platform that provides a distributed  storage system. It abstracts and pools the storage resources across the cluster. | [optional] [readonly] 
**BusType** | Pointer to **string** | The bus type of the disk. Possible values are &#39;IDE&#39;, &#39;PCI&#39;, &#39;SCSI&#39;, &#39;SATA&#39;, &#39;SPAPR&#39;. | [optional] [readonly] 
**DiskExtId** | Pointer to **string** | The unique identifier of the disk. | [optional] [readonly] 
**DiskSizeBytes** | Pointer to **int64** | The size of the disk in bytes. | [optional] [readonly] 
**Index** | Pointer to **int32** | The index of the disk, similar to a slot number on physical machine. | [optional] [readonly] 
**IsFlashModeEnabled** | Pointer to **bool** | Indicates whether the virtual disk is pinned to the hot tier or not. | [optional] [readonly] 
**IsMigrationInProgress** | Pointer to **bool** | Indicates if the disk is being migrated. | [optional] [readonly] 
**StorageContainerExtId** | Pointer to **string** | The extId of the storage container which backs this disk. | [optional] [readonly] 
**VmExtId** | Pointer to **string** | The unique identifier of the VM. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vm** | Pointer to [**NullableHciAhvVmRelationship**](HciAhvVmRelationship.md) |  | [optional] 

## Methods

### NewHciAhvVmDisk

`func NewHciAhvVmDisk(classId string, objectType string, ) *HciAhvVmDisk`

NewHciAhvVmDisk instantiates a new HciAhvVmDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciAhvVmDiskWithDefaults

`func NewHciAhvVmDiskWithDefaults() *HciAhvVmDisk`

NewHciAhvVmDiskWithDefaults instantiates a new HciAhvVmDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciAhvVmDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciAhvVmDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciAhvVmDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciAhvVmDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciAhvVmDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciAhvVmDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdsfVolumeGroupExtId

`func (o *HciAhvVmDisk) GetAdsfVolumeGroupExtId() string`

GetAdsfVolumeGroupExtId returns the AdsfVolumeGroupExtId field if non-nil, zero value otherwise.

### GetAdsfVolumeGroupExtIdOk

`func (o *HciAhvVmDisk) GetAdsfVolumeGroupExtIdOk() (*string, bool)`

GetAdsfVolumeGroupExtIdOk returns a tuple with the AdsfVolumeGroupExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsfVolumeGroupExtId

`func (o *HciAhvVmDisk) SetAdsfVolumeGroupExtId(v string)`

SetAdsfVolumeGroupExtId sets AdsfVolumeGroupExtId field to given value.

### HasAdsfVolumeGroupExtId

`func (o *HciAhvVmDisk) HasAdsfVolumeGroupExtId() bool`

HasAdsfVolumeGroupExtId returns a boolean if a field has been set.

### GetBusType

`func (o *HciAhvVmDisk) GetBusType() string`

GetBusType returns the BusType field if non-nil, zero value otherwise.

### GetBusTypeOk

`func (o *HciAhvVmDisk) GetBusTypeOk() (*string, bool)`

GetBusTypeOk returns a tuple with the BusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusType

`func (o *HciAhvVmDisk) SetBusType(v string)`

SetBusType sets BusType field to given value.

### HasBusType

`func (o *HciAhvVmDisk) HasBusType() bool`

HasBusType returns a boolean if a field has been set.

### GetDiskExtId

`func (o *HciAhvVmDisk) GetDiskExtId() string`

GetDiskExtId returns the DiskExtId field if non-nil, zero value otherwise.

### GetDiskExtIdOk

`func (o *HciAhvVmDisk) GetDiskExtIdOk() (*string, bool)`

GetDiskExtIdOk returns a tuple with the DiskExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskExtId

`func (o *HciAhvVmDisk) SetDiskExtId(v string)`

SetDiskExtId sets DiskExtId field to given value.

### HasDiskExtId

`func (o *HciAhvVmDisk) HasDiskExtId() bool`

HasDiskExtId returns a boolean if a field has been set.

### GetDiskSizeBytes

`func (o *HciAhvVmDisk) GetDiskSizeBytes() int64`

GetDiskSizeBytes returns the DiskSizeBytes field if non-nil, zero value otherwise.

### GetDiskSizeBytesOk

`func (o *HciAhvVmDisk) GetDiskSizeBytesOk() (*int64, bool)`

GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeBytes

`func (o *HciAhvVmDisk) SetDiskSizeBytes(v int64)`

SetDiskSizeBytes sets DiskSizeBytes field to given value.

### HasDiskSizeBytes

`func (o *HciAhvVmDisk) HasDiskSizeBytes() bool`

HasDiskSizeBytes returns a boolean if a field has been set.

### GetIndex

`func (o *HciAhvVmDisk) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *HciAhvVmDisk) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *HciAhvVmDisk) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *HciAhvVmDisk) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIsFlashModeEnabled

`func (o *HciAhvVmDisk) GetIsFlashModeEnabled() bool`

GetIsFlashModeEnabled returns the IsFlashModeEnabled field if non-nil, zero value otherwise.

### GetIsFlashModeEnabledOk

`func (o *HciAhvVmDisk) GetIsFlashModeEnabledOk() (*bool, bool)`

GetIsFlashModeEnabledOk returns a tuple with the IsFlashModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlashModeEnabled

`func (o *HciAhvVmDisk) SetIsFlashModeEnabled(v bool)`

SetIsFlashModeEnabled sets IsFlashModeEnabled field to given value.

### HasIsFlashModeEnabled

`func (o *HciAhvVmDisk) HasIsFlashModeEnabled() bool`

HasIsFlashModeEnabled returns a boolean if a field has been set.

### GetIsMigrationInProgress

`func (o *HciAhvVmDisk) GetIsMigrationInProgress() bool`

GetIsMigrationInProgress returns the IsMigrationInProgress field if non-nil, zero value otherwise.

### GetIsMigrationInProgressOk

`func (o *HciAhvVmDisk) GetIsMigrationInProgressOk() (*bool, bool)`

GetIsMigrationInProgressOk returns a tuple with the IsMigrationInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMigrationInProgress

`func (o *HciAhvVmDisk) SetIsMigrationInProgress(v bool)`

SetIsMigrationInProgress sets IsMigrationInProgress field to given value.

### HasIsMigrationInProgress

`func (o *HciAhvVmDisk) HasIsMigrationInProgress() bool`

HasIsMigrationInProgress returns a boolean if a field has been set.

### GetStorageContainerExtId

`func (o *HciAhvVmDisk) GetStorageContainerExtId() string`

GetStorageContainerExtId returns the StorageContainerExtId field if non-nil, zero value otherwise.

### GetStorageContainerExtIdOk

`func (o *HciAhvVmDisk) GetStorageContainerExtIdOk() (*string, bool)`

GetStorageContainerExtIdOk returns a tuple with the StorageContainerExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerExtId

`func (o *HciAhvVmDisk) SetStorageContainerExtId(v string)`

SetStorageContainerExtId sets StorageContainerExtId field to given value.

### HasStorageContainerExtId

`func (o *HciAhvVmDisk) HasStorageContainerExtId() bool`

HasStorageContainerExtId returns a boolean if a field has been set.

### GetVmExtId

`func (o *HciAhvVmDisk) GetVmExtId() string`

GetVmExtId returns the VmExtId field if non-nil, zero value otherwise.

### GetVmExtIdOk

`func (o *HciAhvVmDisk) GetVmExtIdOk() (*string, bool)`

GetVmExtIdOk returns a tuple with the VmExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmExtId

`func (o *HciAhvVmDisk) SetVmExtId(v string)`

SetVmExtId sets VmExtId field to given value.

### HasVmExtId

`func (o *HciAhvVmDisk) HasVmExtId() bool`

HasVmExtId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HciAhvVmDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciAhvVmDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciAhvVmDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciAhvVmDisk) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciAhvVmDisk) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciAhvVmDisk) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVm

`func (o *HciAhvVmDisk) GetVm() HciAhvVmRelationship`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *HciAhvVmDisk) GetVmOk() (*HciAhvVmRelationship, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *HciAhvVmDisk) SetVm(v HciAhvVmRelationship)`

SetVm sets Vm field to given value.

### HasVm

`func (o *HciAhvVmDisk) HasVm() bool`

HasVm returns a boolean if a field has been set.

### SetVmNil

`func (o *HciAhvVmDisk) SetVmNil(b bool)`

 SetVmNil sets the value for Vm to be an explicit nil

### UnsetVm
`func (o *HciAhvVmDisk) UnsetVm()`

UnsetVm ensures that no value is present for Vm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


