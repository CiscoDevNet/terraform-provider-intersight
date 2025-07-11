# HciAhvVmGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.AhvVmGpu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.AhvVmGpu"]
**DeviceId** | Pointer to **int64** | The device id of the GPU. | [optional] [readonly] 
**Fraction** | Pointer to **int64** | The fraction of a physical GPU&#39;s resources that are allocated to a virtual GPU, e.g. 8 indicate 1/8 of a GPU. | [optional] [readonly] 
**FrameBufferSizeBytes** | Pointer to **int64** | The frame buffer size of the GPU. | [optional] [readonly] 
**GpuExtId** | Pointer to **string** | The unique identifier of the GPU. | [optional] [readonly] 
**GuestDriverVersion** | Pointer to **string** | Last determined guest driver version. | [optional] [readonly] 
**Mode** | Pointer to **string** | The mode of the GPU such as UNUSED, USED_FOR_PASSTHROUGH, USED_FOR_VIRTUAL. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the GPU allocation such as \&quot;NVIDIA_A16-1B\&quot;. | [optional] [readonly] 
**NumVirtualDisplayHeads** | Pointer to **int32** | The number of virtual display heads of the GPU. | [optional] [readonly] 
**PciAddress** | Pointer to [**NullableHciSbdf**](HciSbdf.md) |  | [optional] 
**Vendor** | Pointer to **string** | The vendor name of the GPU. | [optional] [readonly] 
**VmExtId** | Pointer to **string** | The unique identifier of the VM. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vm** | Pointer to [**NullableHciAhvVmRelationship**](HciAhvVmRelationship.md) |  | [optional] 

## Methods

### NewHciAhvVmGpu

`func NewHciAhvVmGpu(classId string, objectType string, ) *HciAhvVmGpu`

NewHciAhvVmGpu instantiates a new HciAhvVmGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciAhvVmGpuWithDefaults

`func NewHciAhvVmGpuWithDefaults() *HciAhvVmGpu`

NewHciAhvVmGpuWithDefaults instantiates a new HciAhvVmGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciAhvVmGpu) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciAhvVmGpu) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciAhvVmGpu) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciAhvVmGpu) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciAhvVmGpu) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciAhvVmGpu) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *HciAhvVmGpu) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HciAhvVmGpu) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HciAhvVmGpu) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HciAhvVmGpu) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFraction

`func (o *HciAhvVmGpu) GetFraction() int64`

GetFraction returns the Fraction field if non-nil, zero value otherwise.

### GetFractionOk

`func (o *HciAhvVmGpu) GetFractionOk() (*int64, bool)`

GetFractionOk returns a tuple with the Fraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraction

`func (o *HciAhvVmGpu) SetFraction(v int64)`

SetFraction sets Fraction field to given value.

### HasFraction

`func (o *HciAhvVmGpu) HasFraction() bool`

HasFraction returns a boolean if a field has been set.

### GetFrameBufferSizeBytes

`func (o *HciAhvVmGpu) GetFrameBufferSizeBytes() int64`

GetFrameBufferSizeBytes returns the FrameBufferSizeBytes field if non-nil, zero value otherwise.

### GetFrameBufferSizeBytesOk

`func (o *HciAhvVmGpu) GetFrameBufferSizeBytesOk() (*int64, bool)`

GetFrameBufferSizeBytesOk returns a tuple with the FrameBufferSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeBytes

`func (o *HciAhvVmGpu) SetFrameBufferSizeBytes(v int64)`

SetFrameBufferSizeBytes sets FrameBufferSizeBytes field to given value.

### HasFrameBufferSizeBytes

`func (o *HciAhvVmGpu) HasFrameBufferSizeBytes() bool`

HasFrameBufferSizeBytes returns a boolean if a field has been set.

### GetGpuExtId

`func (o *HciAhvVmGpu) GetGpuExtId() string`

GetGpuExtId returns the GpuExtId field if non-nil, zero value otherwise.

### GetGpuExtIdOk

`func (o *HciAhvVmGpu) GetGpuExtIdOk() (*string, bool)`

GetGpuExtIdOk returns a tuple with the GpuExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuExtId

`func (o *HciAhvVmGpu) SetGpuExtId(v string)`

SetGpuExtId sets GpuExtId field to given value.

### HasGpuExtId

`func (o *HciAhvVmGpu) HasGpuExtId() bool`

HasGpuExtId returns a boolean if a field has been set.

### GetGuestDriverVersion

`func (o *HciAhvVmGpu) GetGuestDriverVersion() string`

GetGuestDriverVersion returns the GuestDriverVersion field if non-nil, zero value otherwise.

### GetGuestDriverVersionOk

`func (o *HciAhvVmGpu) GetGuestDriverVersionOk() (*string, bool)`

GetGuestDriverVersionOk returns a tuple with the GuestDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestDriverVersion

`func (o *HciAhvVmGpu) SetGuestDriverVersion(v string)`

SetGuestDriverVersion sets GuestDriverVersion field to given value.

### HasGuestDriverVersion

`func (o *HciAhvVmGpu) HasGuestDriverVersion() bool`

HasGuestDriverVersion returns a boolean if a field has been set.

### GetMode

`func (o *HciAhvVmGpu) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HciAhvVmGpu) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HciAhvVmGpu) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HciAhvVmGpu) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *HciAhvVmGpu) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciAhvVmGpu) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciAhvVmGpu) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciAhvVmGpu) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumVirtualDisplayHeads

`func (o *HciAhvVmGpu) GetNumVirtualDisplayHeads() int32`

GetNumVirtualDisplayHeads returns the NumVirtualDisplayHeads field if non-nil, zero value otherwise.

### GetNumVirtualDisplayHeadsOk

`func (o *HciAhvVmGpu) GetNumVirtualDisplayHeadsOk() (*int32, bool)`

GetNumVirtualDisplayHeadsOk returns a tuple with the NumVirtualDisplayHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVirtualDisplayHeads

`func (o *HciAhvVmGpu) SetNumVirtualDisplayHeads(v int32)`

SetNumVirtualDisplayHeads sets NumVirtualDisplayHeads field to given value.

### HasNumVirtualDisplayHeads

`func (o *HciAhvVmGpu) HasNumVirtualDisplayHeads() bool`

HasNumVirtualDisplayHeads returns a boolean if a field has been set.

### GetPciAddress

`func (o *HciAhvVmGpu) GetPciAddress() HciSbdf`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *HciAhvVmGpu) GetPciAddressOk() (*HciSbdf, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *HciAhvVmGpu) SetPciAddress(v HciSbdf)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *HciAhvVmGpu) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### SetPciAddressNil

`func (o *HciAhvVmGpu) SetPciAddressNil(b bool)`

 SetPciAddressNil sets the value for PciAddress to be an explicit nil

### UnsetPciAddress
`func (o *HciAhvVmGpu) UnsetPciAddress()`

UnsetPciAddress ensures that no value is present for PciAddress, not even an explicit nil
### GetVendor

`func (o *HciAhvVmGpu) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HciAhvVmGpu) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HciAhvVmGpu) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HciAhvVmGpu) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVmExtId

`func (o *HciAhvVmGpu) GetVmExtId() string`

GetVmExtId returns the VmExtId field if non-nil, zero value otherwise.

### GetVmExtIdOk

`func (o *HciAhvVmGpu) GetVmExtIdOk() (*string, bool)`

GetVmExtIdOk returns a tuple with the VmExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmExtId

`func (o *HciAhvVmGpu) SetVmExtId(v string)`

SetVmExtId sets VmExtId field to given value.

### HasVmExtId

`func (o *HciAhvVmGpu) HasVmExtId() bool`

HasVmExtId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HciAhvVmGpu) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciAhvVmGpu) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciAhvVmGpu) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciAhvVmGpu) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciAhvVmGpu) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciAhvVmGpu) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVm

`func (o *HciAhvVmGpu) GetVm() HciAhvVmRelationship`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *HciAhvVmGpu) GetVmOk() (*HciAhvVmRelationship, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *HciAhvVmGpu) SetVm(v HciAhvVmRelationship)`

SetVm sets Vm field to given value.

### HasVm

`func (o *HciAhvVmGpu) HasVm() bool`

HasVm returns a boolean if a field has been set.

### SetVmNil

`func (o *HciAhvVmGpu) SetVmNil(b bool)`

 SetVmNil sets the value for Vm to be an explicit nil

### UnsetVm
`func (o *HciAhvVmGpu) UnsetVm()`

UnsetVm ensures that no value is present for Vm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


