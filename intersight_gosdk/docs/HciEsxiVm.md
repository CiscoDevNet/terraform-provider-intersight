# HciEsxiVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.EsxiVm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.EsxiVm"]
**GuestTools** | Pointer to [**NullableHciEsxiGuestTools**](HciEsxiGuestTools.md) |  | [optional] 
**VirtualHardwareVersion** | Pointer to **int32** | The virtual hardware version of the VM. | [optional] [readonly] 
**Disks** | Pointer to [**[]HciEsxiVmDiskRelationship**](HciEsxiVmDiskRelationship.md) | An array of relationships to hciEsxiVmDisk resources. | [optional] [readonly] 
**Nics** | Pointer to [**[]HciEsxiVmNicRelationship**](HciEsxiVmNicRelationship.md) | An array of relationships to hciEsxiVmNic resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciEsxiVm

`func NewHciEsxiVm(classId string, objectType string, ) *HciEsxiVm`

NewHciEsxiVm instantiates a new HciEsxiVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciEsxiVmWithDefaults

`func NewHciEsxiVmWithDefaults() *HciEsxiVm`

NewHciEsxiVmWithDefaults instantiates a new HciEsxiVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciEsxiVm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciEsxiVm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciEsxiVm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciEsxiVm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciEsxiVm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciEsxiVm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGuestTools

`func (o *HciEsxiVm) GetGuestTools() HciEsxiGuestTools`

GetGuestTools returns the GuestTools field if non-nil, zero value otherwise.

### GetGuestToolsOk

`func (o *HciEsxiVm) GetGuestToolsOk() (*HciEsxiGuestTools, bool)`

GetGuestToolsOk returns a tuple with the GuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestTools

`func (o *HciEsxiVm) SetGuestTools(v HciEsxiGuestTools)`

SetGuestTools sets GuestTools field to given value.

### HasGuestTools

`func (o *HciEsxiVm) HasGuestTools() bool`

HasGuestTools returns a boolean if a field has been set.

### SetGuestToolsNil

`func (o *HciEsxiVm) SetGuestToolsNil(b bool)`

 SetGuestToolsNil sets the value for GuestTools to be an explicit nil

### UnsetGuestTools
`func (o *HciEsxiVm) UnsetGuestTools()`

UnsetGuestTools ensures that no value is present for GuestTools, not even an explicit nil
### GetVirtualHardwareVersion

`func (o *HciEsxiVm) GetVirtualHardwareVersion() int32`

GetVirtualHardwareVersion returns the VirtualHardwareVersion field if non-nil, zero value otherwise.

### GetVirtualHardwareVersionOk

`func (o *HciEsxiVm) GetVirtualHardwareVersionOk() (*int32, bool)`

GetVirtualHardwareVersionOk returns a tuple with the VirtualHardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHardwareVersion

`func (o *HciEsxiVm) SetVirtualHardwareVersion(v int32)`

SetVirtualHardwareVersion sets VirtualHardwareVersion field to given value.

### HasVirtualHardwareVersion

`func (o *HciEsxiVm) HasVirtualHardwareVersion() bool`

HasVirtualHardwareVersion returns a boolean if a field has been set.

### GetDisks

`func (o *HciEsxiVm) GetDisks() []HciEsxiVmDiskRelationship`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *HciEsxiVm) GetDisksOk() (*[]HciEsxiVmDiskRelationship, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *HciEsxiVm) SetDisks(v []HciEsxiVmDiskRelationship)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *HciEsxiVm) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *HciEsxiVm) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *HciEsxiVm) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetNics

`func (o *HciEsxiVm) GetNics() []HciEsxiVmNicRelationship`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *HciEsxiVm) GetNicsOk() (*[]HciEsxiVmNicRelationship, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *HciEsxiVm) SetNics(v []HciEsxiVmNicRelationship)`

SetNics sets Nics field to given value.

### HasNics

`func (o *HciEsxiVm) HasNics() bool`

HasNics returns a boolean if a field has been set.

### SetNicsNil

`func (o *HciEsxiVm) SetNicsNil(b bool)`

 SetNicsNil sets the value for Nics to be an explicit nil

### UnsetNics
`func (o *HciEsxiVm) UnsetNics()`

UnsetNics ensures that no value is present for Nics, not even an explicit nil
### GetRegisteredDevice

`func (o *HciEsxiVm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciEsxiVm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciEsxiVm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciEsxiVm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciEsxiVm) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciEsxiVm) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


