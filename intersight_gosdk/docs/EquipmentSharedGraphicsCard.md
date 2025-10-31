# EquipmentSharedGraphicsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.SharedGraphicsCard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.SharedGraphicsCard"]
**Description** | Pointer to **string** | The description of the GPU card. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The unique device identifier assigned by the vendor to a specific model of GPU. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The version of the GPU firmware. | [optional] [readonly] 
**GpuId** | Pointer to **string** | The identifier of the graphics card. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational state of graphics card. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number identifier for the graphics card. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | PCIe slot of the GPU in the PCIe node. | [optional] [readonly] 
**Pid** | Pointer to **string** | The unique product ID associated with the GPU card. | [optional] [readonly] 
**SubDeviceId** | Pointer to **int64** | The subsystem device identifier assigned by the subsystem vendor to a specific model of GPU. | [optional] [readonly] 
**SubVendorId** | Pointer to **int64** | The unique vendor identifier assigned to the organization which integrates the GPU. | [optional] [readonly] 
**VendorId** | Pointer to **int64** | The unique vendor identifier assigned to the manufacturer of the GPU. | [optional] [readonly] 
**EquipmentInterconnect** | Pointer to [**NullableEquipmentInterconnectRelationship**](EquipmentInterconnectRelationship.md) |  | [optional] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](GraphicsCardRelationship.md) | An array of relationships to graphicsCard resources. | [optional] [readonly] 
**PciEndpoint** | Pointer to [**NullablePciEndpointRelationship**](PciEndpointRelationship.md) |  | [optional] 
**PciNode** | Pointer to [**NullablePciNodeRelationship**](PciNodeRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSharedGraphicsCard

`func NewEquipmentSharedGraphicsCard(classId string, objectType string, ) *EquipmentSharedGraphicsCard`

NewEquipmentSharedGraphicsCard instantiates a new EquipmentSharedGraphicsCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSharedGraphicsCardWithDefaults

`func NewEquipmentSharedGraphicsCardWithDefaults() *EquipmentSharedGraphicsCard`

NewEquipmentSharedGraphicsCardWithDefaults instantiates a new EquipmentSharedGraphicsCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSharedGraphicsCard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSharedGraphicsCard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSharedGraphicsCard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSharedGraphicsCard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSharedGraphicsCard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSharedGraphicsCard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *EquipmentSharedGraphicsCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentSharedGraphicsCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentSharedGraphicsCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentSharedGraphicsCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceId

`func (o *EquipmentSharedGraphicsCard) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *EquipmentSharedGraphicsCard) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *EquipmentSharedGraphicsCard) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *EquipmentSharedGraphicsCard) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *EquipmentSharedGraphicsCard) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *EquipmentSharedGraphicsCard) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *EquipmentSharedGraphicsCard) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *EquipmentSharedGraphicsCard) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetGpuId

`func (o *EquipmentSharedGraphicsCard) GetGpuId() string`

GetGpuId returns the GpuId field if non-nil, zero value otherwise.

### GetGpuIdOk

`func (o *EquipmentSharedGraphicsCard) GetGpuIdOk() (*string, bool)`

GetGpuIdOk returns a tuple with the GpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuId

`func (o *EquipmentSharedGraphicsCard) SetGpuId(v string)`

SetGpuId sets GpuId field to given value.

### HasGpuId

`func (o *EquipmentSharedGraphicsCard) HasGpuId() bool`

HasGpuId returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentSharedGraphicsCard) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentSharedGraphicsCard) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentSharedGraphicsCard) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentSharedGraphicsCard) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentSharedGraphicsCard) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentSharedGraphicsCard) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentSharedGraphicsCard) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentSharedGraphicsCard) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentSharedGraphicsCard) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentSharedGraphicsCard) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentSharedGraphicsCard) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentSharedGraphicsCard) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentSharedGraphicsCard) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentSharedGraphicsCard) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPciSlot

`func (o *EquipmentSharedGraphicsCard) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *EquipmentSharedGraphicsCard) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *EquipmentSharedGraphicsCard) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *EquipmentSharedGraphicsCard) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentSharedGraphicsCard) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentSharedGraphicsCard) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentSharedGraphicsCard) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentSharedGraphicsCard) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSubDeviceId

`func (o *EquipmentSharedGraphicsCard) GetSubDeviceId() int64`

GetSubDeviceId returns the SubDeviceId field if non-nil, zero value otherwise.

### GetSubDeviceIdOk

`func (o *EquipmentSharedGraphicsCard) GetSubDeviceIdOk() (*int64, bool)`

GetSubDeviceIdOk returns a tuple with the SubDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDeviceId

`func (o *EquipmentSharedGraphicsCard) SetSubDeviceId(v int64)`

SetSubDeviceId sets SubDeviceId field to given value.

### HasSubDeviceId

`func (o *EquipmentSharedGraphicsCard) HasSubDeviceId() bool`

HasSubDeviceId returns a boolean if a field has been set.

### GetSubVendorId

`func (o *EquipmentSharedGraphicsCard) GetSubVendorId() int64`

GetSubVendorId returns the SubVendorId field if non-nil, zero value otherwise.

### GetSubVendorIdOk

`func (o *EquipmentSharedGraphicsCard) GetSubVendorIdOk() (*int64, bool)`

GetSubVendorIdOk returns a tuple with the SubVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVendorId

`func (o *EquipmentSharedGraphicsCard) SetSubVendorId(v int64)`

SetSubVendorId sets SubVendorId field to given value.

### HasSubVendorId

`func (o *EquipmentSharedGraphicsCard) HasSubVendorId() bool`

HasSubVendorId returns a boolean if a field has been set.

### GetVendorId

`func (o *EquipmentSharedGraphicsCard) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *EquipmentSharedGraphicsCard) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *EquipmentSharedGraphicsCard) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *EquipmentSharedGraphicsCard) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetEquipmentInterconnect

`func (o *EquipmentSharedGraphicsCard) GetEquipmentInterconnect() EquipmentInterconnectRelationship`

GetEquipmentInterconnect returns the EquipmentInterconnect field if non-nil, zero value otherwise.

### GetEquipmentInterconnectOk

`func (o *EquipmentSharedGraphicsCard) GetEquipmentInterconnectOk() (*EquipmentInterconnectRelationship, bool)`

GetEquipmentInterconnectOk returns a tuple with the EquipmentInterconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentInterconnect

`func (o *EquipmentSharedGraphicsCard) SetEquipmentInterconnect(v EquipmentInterconnectRelationship)`

SetEquipmentInterconnect sets EquipmentInterconnect field to given value.

### HasEquipmentInterconnect

`func (o *EquipmentSharedGraphicsCard) HasEquipmentInterconnect() bool`

HasEquipmentInterconnect returns a boolean if a field has been set.

### SetEquipmentInterconnectNil

`func (o *EquipmentSharedGraphicsCard) SetEquipmentInterconnectNil(b bool)`

 SetEquipmentInterconnectNil sets the value for EquipmentInterconnect to be an explicit nil

### UnsetEquipmentInterconnect
`func (o *EquipmentSharedGraphicsCard) UnsetEquipmentInterconnect()`

UnsetEquipmentInterconnect ensures that no value is present for EquipmentInterconnect, not even an explicit nil
### GetGraphicsCards

`func (o *EquipmentSharedGraphicsCard) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *EquipmentSharedGraphicsCard) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *EquipmentSharedGraphicsCard) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *EquipmentSharedGraphicsCard) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *EquipmentSharedGraphicsCard) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *EquipmentSharedGraphicsCard) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetPciEndpoint

`func (o *EquipmentSharedGraphicsCard) GetPciEndpoint() PciEndpointRelationship`

GetPciEndpoint returns the PciEndpoint field if non-nil, zero value otherwise.

### GetPciEndpointOk

`func (o *EquipmentSharedGraphicsCard) GetPciEndpointOk() (*PciEndpointRelationship, bool)`

GetPciEndpointOk returns a tuple with the PciEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciEndpoint

`func (o *EquipmentSharedGraphicsCard) SetPciEndpoint(v PciEndpointRelationship)`

SetPciEndpoint sets PciEndpoint field to given value.

### HasPciEndpoint

`func (o *EquipmentSharedGraphicsCard) HasPciEndpoint() bool`

HasPciEndpoint returns a boolean if a field has been set.

### SetPciEndpointNil

`func (o *EquipmentSharedGraphicsCard) SetPciEndpointNil(b bool)`

 SetPciEndpointNil sets the value for PciEndpoint to be an explicit nil

### UnsetPciEndpoint
`func (o *EquipmentSharedGraphicsCard) UnsetPciEndpoint()`

UnsetPciEndpoint ensures that no value is present for PciEndpoint, not even an explicit nil
### GetPciNode

`func (o *EquipmentSharedGraphicsCard) GetPciNode() PciNodeRelationship`

GetPciNode returns the PciNode field if non-nil, zero value otherwise.

### GetPciNodeOk

`func (o *EquipmentSharedGraphicsCard) GetPciNodeOk() (*PciNodeRelationship, bool)`

GetPciNodeOk returns a tuple with the PciNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciNode

`func (o *EquipmentSharedGraphicsCard) SetPciNode(v PciNodeRelationship)`

SetPciNode sets PciNode field to given value.

### HasPciNode

`func (o *EquipmentSharedGraphicsCard) HasPciNode() bool`

HasPciNode returns a boolean if a field has been set.

### SetPciNodeNil

`func (o *EquipmentSharedGraphicsCard) SetPciNodeNil(b bool)`

 SetPciNodeNil sets the value for PciNode to be an explicit nil

### UnsetPciNode
`func (o *EquipmentSharedGraphicsCard) UnsetPciNode()`

UnsetPciNode ensures that no value is present for PciNode, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSharedGraphicsCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSharedGraphicsCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSharedGraphicsCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSharedGraphicsCard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentSharedGraphicsCard) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentSharedGraphicsCard) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


