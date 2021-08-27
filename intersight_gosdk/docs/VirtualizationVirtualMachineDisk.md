# VirtualizationVirtualMachineDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VirtualMachineDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VirtualMachineDisk"]
**Bus** | Pointer to **string** | Disk bus name given for a virtual machine. * &#x60;virtio&#x60; - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration. * &#x60;sata&#x60; - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives. * &#x60;scsi&#x60; - SCSI (Small Computer System Interface) bus used.. | [optional] [default to "virtio"]
**Name** | Pointer to **string** | Virtual machine network bridge name. | [optional] 
**Order** | Pointer to **int64** | Priority order of the disk. | [optional] 
**Type** | Pointer to **string** | Disk type hdd or cdrom for a virtual machine. * &#x60;hdd&#x60; - Allows the virtual machine to mount disk from hard disk drive (hdd) image. * &#x60;cdrom&#x60; - Allows the virtual machine to mount disk from compact disk (cd) image. | [optional] [default to "hdd"]
**VirtualDisk** | Pointer to [**NullableVirtualizationVirtualDiskConfig**](VirtualizationVirtualDiskConfig.md) |  | [optional] 
**VirtualDiskReference** | Pointer to **string** | Name of the existing virtual disk to be attached to the Virtual Machine. | [optional] 

## Methods

### NewVirtualizationVirtualMachineDisk

`func NewVirtualizationVirtualMachineDisk(classId string, objectType string, ) *VirtualizationVirtualMachineDisk`

NewVirtualizationVirtualMachineDisk instantiates a new VirtualizationVirtualMachineDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualMachineDiskWithDefaults

`func NewVirtualizationVirtualMachineDiskWithDefaults() *VirtualizationVirtualMachineDisk`

NewVirtualizationVirtualMachineDiskWithDefaults instantiates a new VirtualizationVirtualMachineDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualMachineDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualMachineDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualMachineDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualMachineDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualMachineDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualMachineDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBus

`func (o *VirtualizationVirtualMachineDisk) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VirtualizationVirtualMachineDisk) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VirtualizationVirtualMachineDisk) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VirtualizationVirtualMachineDisk) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualMachineDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualMachineDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualMachineDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualMachineDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VirtualizationVirtualMachineDisk) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VirtualizationVirtualMachineDisk) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VirtualizationVirtualMachineDisk) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VirtualizationVirtualMachineDisk) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationVirtualMachineDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationVirtualMachineDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationVirtualMachineDisk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationVirtualMachineDisk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualDisk

`func (o *VirtualizationVirtualMachineDisk) GetVirtualDisk() VirtualizationVirtualDiskConfig`

GetVirtualDisk returns the VirtualDisk field if non-nil, zero value otherwise.

### GetVirtualDiskOk

`func (o *VirtualizationVirtualMachineDisk) GetVirtualDiskOk() (*VirtualizationVirtualDiskConfig, bool)`

GetVirtualDiskOk returns a tuple with the VirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisk

`func (o *VirtualizationVirtualMachineDisk) SetVirtualDisk(v VirtualizationVirtualDiskConfig)`

SetVirtualDisk sets VirtualDisk field to given value.

### HasVirtualDisk

`func (o *VirtualizationVirtualMachineDisk) HasVirtualDisk() bool`

HasVirtualDisk returns a boolean if a field has been set.

### SetVirtualDiskNil

`func (o *VirtualizationVirtualMachineDisk) SetVirtualDiskNil(b bool)`

 SetVirtualDiskNil sets the value for VirtualDisk to be an explicit nil

### UnsetVirtualDisk
`func (o *VirtualizationVirtualMachineDisk) UnsetVirtualDisk()`

UnsetVirtualDisk ensures that no value is present for VirtualDisk, not even an explicit nil
### GetVirtualDiskReference

`func (o *VirtualizationVirtualMachineDisk) GetVirtualDiskReference() string`

GetVirtualDiskReference returns the VirtualDiskReference field if non-nil, zero value otherwise.

### GetVirtualDiskReferenceOk

`func (o *VirtualizationVirtualMachineDisk) GetVirtualDiskReferenceOk() (*string, bool)`

GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskReference

`func (o *VirtualizationVirtualMachineDisk) SetVirtualDiskReference(v string)`

SetVirtualDiskReference sets VirtualDiskReference field to given value.

### HasVirtualDiskReference

`func (o *VirtualizationVirtualMachineDisk) HasVirtualDiskReference() bool`

HasVirtualDiskReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


