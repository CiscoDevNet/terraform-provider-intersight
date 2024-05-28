# VirtualizationBaseVirtualMachinePciDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareVirtualMachineGpu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareVirtualMachineGpu"]
**BackingPciId** | Pointer to **string** | The backing physical host PCI device Id for this device. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of this virtual machine PCI device. | [optional] [readonly] 
**Passthrough** | Pointer to **bool** | Indicates if this virtual machine PCI device is enabled via passthrough from the host. | [optional] [readonly] 
**BackingPciDevice** | Pointer to [**NullableVirtualizationBaseHostPciDeviceRelationship**](VirtualizationBaseHostPciDeviceRelationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**NullableVirtualizationBaseVirtualMachineRelationship**](VirtualizationBaseVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationBaseVirtualMachinePciDevice

`func NewVirtualizationBaseVirtualMachinePciDevice(classId string, objectType string, ) *VirtualizationBaseVirtualMachinePciDevice`

NewVirtualizationBaseVirtualMachinePciDevice instantiates a new VirtualizationBaseVirtualMachinePciDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseVirtualMachinePciDeviceWithDefaults

`func NewVirtualizationBaseVirtualMachinePciDeviceWithDefaults() *VirtualizationBaseVirtualMachinePciDevice`

NewVirtualizationBaseVirtualMachinePciDeviceWithDefaults instantiates a new VirtualizationBaseVirtualMachinePciDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackingPciId

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciId() string`

GetBackingPciId returns the BackingPciId field if non-nil, zero value otherwise.

### GetBackingPciIdOk

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciIdOk() (*string, bool)`

GetBackingPciIdOk returns a tuple with the BackingPciId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingPciId

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetBackingPciId(v string)`

SetBackingPciId sets BackingPciId field to given value.

### HasBackingPciId

`func (o *VirtualizationBaseVirtualMachinePciDevice) HasBackingPciId() bool`

HasBackingPciId returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseVirtualMachinePciDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassthrough

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetPassthrough() bool`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetPassthroughOk() (*bool, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetPassthrough(v bool)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *VirtualizationBaseVirtualMachinePciDevice) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetBackingPciDevice

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciDevice() VirtualizationBaseHostPciDeviceRelationship`

GetBackingPciDevice returns the BackingPciDevice field if non-nil, zero value otherwise.

### GetBackingPciDeviceOk

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciDeviceOk() (*VirtualizationBaseHostPciDeviceRelationship, bool)`

GetBackingPciDeviceOk returns a tuple with the BackingPciDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingPciDevice

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetBackingPciDevice(v VirtualizationBaseHostPciDeviceRelationship)`

SetBackingPciDevice sets BackingPciDevice field to given value.

### HasBackingPciDevice

`func (o *VirtualizationBaseVirtualMachinePciDevice) HasBackingPciDevice() bool`

HasBackingPciDevice returns a boolean if a field has been set.

### SetBackingPciDeviceNil

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetBackingPciDeviceNil(b bool)`

 SetBackingPciDeviceNil sets the value for BackingPciDevice to be an explicit nil

### UnsetBackingPciDevice
`func (o *VirtualizationBaseVirtualMachinePciDevice) UnsetBackingPciDevice()`

UnsetBackingPciDevice ensures that no value is present for BackingPciDevice, not even an explicit nil
### GetVirtualMachine

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetVirtualMachine() VirtualizationBaseVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualizationBaseVirtualMachinePciDevice) GetVirtualMachineOk() (*VirtualizationBaseVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetVirtualMachine(v VirtualizationBaseVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VirtualizationBaseVirtualMachinePciDevice) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### SetVirtualMachineNil

`func (o *VirtualizationBaseVirtualMachinePciDevice) SetVirtualMachineNil(b bool)`

 SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil

### UnsetVirtualMachine
`func (o *VirtualizationBaseVirtualMachinePciDevice) UnsetVirtualMachine()`

UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


