# VirtualizationVmDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmDisk"]
**BootOrder** | Pointer to **int64** | Boot order for this disk. | [optional] [readonly] 
**Bus** | Pointer to **string** | Virtual machine bridge name of interface. * &#x60;virtio&#x60; - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration. * &#x60;sata&#x60; - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives. * &#x60;scsi&#x60; - SCSI (Small Computer System Interface) bus used.. | [optional] [readonly] [default to "virtio"]
**Name** | Pointer to **string** | Disk name associated with virtual machine. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the Disk (hdd, cdrom). * &#x60;hdd&#x60; - Allows the virtual machine to mount disk from hard disk drive (hdd) image. * &#x60;cdrom&#x60; - Allows the virtual machine to mount disk from compact disk (cd) image. | [optional] [readonly] [default to "hdd"]
**VirtualDisk** | Pointer to [**NullableVirtualizationVdiskConfig**](VirtualizationVdiskConfig.md) |  | [optional] 
**VirtualDiskReference** | Pointer to **string** | Virtual disk reference name. | [optional] [readonly] 

## Methods

### NewVirtualizationVmDiskAllOf

`func NewVirtualizationVmDiskAllOf(classId string, objectType string, ) *VirtualizationVmDiskAllOf`

NewVirtualizationVmDiskAllOf instantiates a new VirtualizationVmDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmDiskAllOfWithDefaults

`func NewVirtualizationVmDiskAllOfWithDefaults() *VirtualizationVmDiskAllOf`

NewVirtualizationVmDiskAllOfWithDefaults instantiates a new VirtualizationVmDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootOrder

`func (o *VirtualizationVmDiskAllOf) GetBootOrder() int64`

GetBootOrder returns the BootOrder field if non-nil, zero value otherwise.

### GetBootOrderOk

`func (o *VirtualizationVmDiskAllOf) GetBootOrderOk() (*int64, bool)`

GetBootOrderOk returns a tuple with the BootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder

`func (o *VirtualizationVmDiskAllOf) SetBootOrder(v int64)`

SetBootOrder sets BootOrder field to given value.

### HasBootOrder

`func (o *VirtualizationVmDiskAllOf) HasBootOrder() bool`

HasBootOrder returns a boolean if a field has been set.

### GetBus

`func (o *VirtualizationVmDiskAllOf) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VirtualizationVmDiskAllOf) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VirtualizationVmDiskAllOf) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VirtualizationVmDiskAllOf) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVmDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVmDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVmDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVmDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationVmDiskAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationVmDiskAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationVmDiskAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationVmDiskAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualDisk

`func (o *VirtualizationVmDiskAllOf) GetVirtualDisk() VirtualizationVdiskConfig`

GetVirtualDisk returns the VirtualDisk field if non-nil, zero value otherwise.

### GetVirtualDiskOk

`func (o *VirtualizationVmDiskAllOf) GetVirtualDiskOk() (*VirtualizationVdiskConfig, bool)`

GetVirtualDiskOk returns a tuple with the VirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisk

`func (o *VirtualizationVmDiskAllOf) SetVirtualDisk(v VirtualizationVdiskConfig)`

SetVirtualDisk sets VirtualDisk field to given value.

### HasVirtualDisk

`func (o *VirtualizationVmDiskAllOf) HasVirtualDisk() bool`

HasVirtualDisk returns a boolean if a field has been set.

### SetVirtualDiskNil

`func (o *VirtualizationVmDiskAllOf) SetVirtualDiskNil(b bool)`

 SetVirtualDiskNil sets the value for VirtualDisk to be an explicit nil

### UnsetVirtualDisk
`func (o *VirtualizationVmDiskAllOf) UnsetVirtualDisk()`

UnsetVirtualDisk ensures that no value is present for VirtualDisk, not even an explicit nil
### GetVirtualDiskReference

`func (o *VirtualizationVmDiskAllOf) GetVirtualDiskReference() string`

GetVirtualDiskReference returns the VirtualDiskReference field if non-nil, zero value otherwise.

### GetVirtualDiskReferenceOk

`func (o *VirtualizationVmDiskAllOf) GetVirtualDiskReferenceOk() (*string, bool)`

GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskReference

`func (o *VirtualizationVmDiskAllOf) SetVirtualDiskReference(v string)`

SetVirtualDiskReference sets VirtualDiskReference field to given value.

### HasVirtualDiskReference

`func (o *VirtualizationVmDiskAllOf) HasVirtualDiskReference() bool`

HasVirtualDiskReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


