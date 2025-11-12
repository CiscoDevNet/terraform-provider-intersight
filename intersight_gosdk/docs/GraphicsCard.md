# GraphicsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "graphics.Card"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "graphics.Card"]
**CardId** | Pointer to **int64** | The id of the graphics card. | [optional] [readonly] 
**Description** | Pointer to **string** | This field displays the description of the Graphics Processing Unit. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The device id of the graphics card. | [optional] [readonly] 
**ExpanderSlot** | Pointer to **string** | The expander slot information of the card. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The firmware version of the graphics card. | [optional] [readonly] 
**GpuId** | Pointer to **string** | The identifier of the graphics processor unit. | [optional] [readonly] 
**IsPlatformSupported** | Pointer to **bool** | This field indicates whether the Graphics Processing Unit is supported on the server or not. | [optional] [readonly] [default to true]
**Mode** | Pointer to **string** | The current mode of the graphics card. | [optional] [readonly] 
**NumGpus** | Pointer to **string** | The number of controllers under each card. | [optional] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | The current operational state of the graphics card. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field displays the part number of the Graphics Processing Unit. | [optional] [readonly] 
**PciAddress** | Pointer to **string** | The PCI address of the graphics card. | [optional] [readonly] 
**PciAddressList** | Pointer to **string** | This list contains the PCI address of all controllers for corresponding card. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The PCI slot name of the graphics card. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field displays the product ID of the Graphics Processing Unit. | [optional] [readonly] 
**SubDeviceId** | Pointer to **int64** | The sub device id of the graphics processor unit. | [optional] [readonly] 
**SubVendorId** | Pointer to **int64** | The sub vendor id of the graphics processor unit. | [optional] [readonly] 
**VendorId** | Pointer to **int64** | The vendor id of the graphics processor unit. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**NullableComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**EquipmentEnclosureElement** | Pointer to [**NullableEquipmentEnclosureElementRelationship**](EquipmentEnclosureElementRelationship.md) |  | [optional] 
**EquipmentInterconnect** | Pointer to [**NullableEquipmentInterconnectRelationship**](EquipmentInterconnectRelationship.md) |  | [optional] 
**EquipmentSharedGraphicsCard** | Pointer to [**NullableEquipmentSharedGraphicsCardRelationship**](EquipmentSharedGraphicsCardRelationship.md) |  | [optional] 
**GraphicsControllers** | Pointer to [**[]GraphicsControllerRelationship**](GraphicsControllerRelationship.md) | An array of relationships to graphicsController resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PciDevice** | Pointer to [**NullablePciDeviceRelationship**](PciDeviceRelationship.md) |  | [optional] 
**PciNode** | Pointer to [**NullablePciNodeRelationship**](PciNodeRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 

## Methods

### NewGraphicsCard

`func NewGraphicsCard(classId string, objectType string, ) *GraphicsCard`

NewGraphicsCard instantiates a new GraphicsCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphicsCardWithDefaults

`func NewGraphicsCardWithDefaults() *GraphicsCard`

NewGraphicsCardWithDefaults instantiates a new GraphicsCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *GraphicsCard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *GraphicsCard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *GraphicsCard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *GraphicsCard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GraphicsCard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GraphicsCard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCardId

`func (o *GraphicsCard) GetCardId() int64`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *GraphicsCard) GetCardIdOk() (*int64, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *GraphicsCard) SetCardId(v int64)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *GraphicsCard) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetDescription

`func (o *GraphicsCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GraphicsCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GraphicsCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GraphicsCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceId

`func (o *GraphicsCard) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GraphicsCard) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GraphicsCard) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GraphicsCard) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExpanderSlot

`func (o *GraphicsCard) GetExpanderSlot() string`

GetExpanderSlot returns the ExpanderSlot field if non-nil, zero value otherwise.

### GetExpanderSlotOk

`func (o *GraphicsCard) GetExpanderSlotOk() (*string, bool)`

GetExpanderSlotOk returns a tuple with the ExpanderSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderSlot

`func (o *GraphicsCard) SetExpanderSlot(v string)`

SetExpanderSlot sets ExpanderSlot field to given value.

### HasExpanderSlot

`func (o *GraphicsCard) HasExpanderSlot() bool`

HasExpanderSlot returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *GraphicsCard) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *GraphicsCard) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *GraphicsCard) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *GraphicsCard) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetGpuId

`func (o *GraphicsCard) GetGpuId() string`

GetGpuId returns the GpuId field if non-nil, zero value otherwise.

### GetGpuIdOk

`func (o *GraphicsCard) GetGpuIdOk() (*string, bool)`

GetGpuIdOk returns a tuple with the GpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuId

`func (o *GraphicsCard) SetGpuId(v string)`

SetGpuId sets GpuId field to given value.

### HasGpuId

`func (o *GraphicsCard) HasGpuId() bool`

HasGpuId returns a boolean if a field has been set.

### GetIsPlatformSupported

`func (o *GraphicsCard) GetIsPlatformSupported() bool`

GetIsPlatformSupported returns the IsPlatformSupported field if non-nil, zero value otherwise.

### GetIsPlatformSupportedOk

`func (o *GraphicsCard) GetIsPlatformSupportedOk() (*bool, bool)`

GetIsPlatformSupportedOk returns a tuple with the IsPlatformSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlatformSupported

`func (o *GraphicsCard) SetIsPlatformSupported(v bool)`

SetIsPlatformSupported sets IsPlatformSupported field to given value.

### HasIsPlatformSupported

`func (o *GraphicsCard) HasIsPlatformSupported() bool`

HasIsPlatformSupported returns a boolean if a field has been set.

### GetMode

`func (o *GraphicsCard) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GraphicsCard) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GraphicsCard) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GraphicsCard) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumGpus

`func (o *GraphicsCard) GetNumGpus() string`

GetNumGpus returns the NumGpus field if non-nil, zero value otherwise.

### GetNumGpusOk

`func (o *GraphicsCard) GetNumGpusOk() (*string, bool)`

GetNumGpusOk returns a tuple with the NumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGpus

`func (o *GraphicsCard) SetNumGpus(v string)`

SetNumGpus sets NumGpus field to given value.

### HasNumGpus

`func (o *GraphicsCard) HasNumGpus() bool`

HasNumGpus returns a boolean if a field has been set.

### GetOperReason

`func (o *GraphicsCard) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *GraphicsCard) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *GraphicsCard) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *GraphicsCard) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *GraphicsCard) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *GraphicsCard) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *GraphicsCard) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *GraphicsCard) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *GraphicsCard) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *GraphicsCard) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *GraphicsCard) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *GraphicsCard) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *GraphicsCard) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *GraphicsCard) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPciAddress

`func (o *GraphicsCard) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *GraphicsCard) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *GraphicsCard) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *GraphicsCard) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPciAddressList

`func (o *GraphicsCard) GetPciAddressList() string`

GetPciAddressList returns the PciAddressList field if non-nil, zero value otherwise.

### GetPciAddressListOk

`func (o *GraphicsCard) GetPciAddressListOk() (*string, bool)`

GetPciAddressListOk returns a tuple with the PciAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddressList

`func (o *GraphicsCard) SetPciAddressList(v string)`

SetPciAddressList sets PciAddressList field to given value.

### HasPciAddressList

`func (o *GraphicsCard) HasPciAddressList() bool`

HasPciAddressList returns a boolean if a field has been set.

### GetPciSlot

`func (o *GraphicsCard) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *GraphicsCard) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *GraphicsCard) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *GraphicsCard) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPid

`func (o *GraphicsCard) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *GraphicsCard) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *GraphicsCard) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *GraphicsCard) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSubDeviceId

`func (o *GraphicsCard) GetSubDeviceId() int64`

GetSubDeviceId returns the SubDeviceId field if non-nil, zero value otherwise.

### GetSubDeviceIdOk

`func (o *GraphicsCard) GetSubDeviceIdOk() (*int64, bool)`

GetSubDeviceIdOk returns a tuple with the SubDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDeviceId

`func (o *GraphicsCard) SetSubDeviceId(v int64)`

SetSubDeviceId sets SubDeviceId field to given value.

### HasSubDeviceId

`func (o *GraphicsCard) HasSubDeviceId() bool`

HasSubDeviceId returns a boolean if a field has been set.

### GetSubVendorId

`func (o *GraphicsCard) GetSubVendorId() int64`

GetSubVendorId returns the SubVendorId field if non-nil, zero value otherwise.

### GetSubVendorIdOk

`func (o *GraphicsCard) GetSubVendorIdOk() (*int64, bool)`

GetSubVendorIdOk returns a tuple with the SubVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVendorId

`func (o *GraphicsCard) SetSubVendorId(v int64)`

SetSubVendorId sets SubVendorId field to given value.

### HasSubVendorId

`func (o *GraphicsCard) HasSubVendorId() bool`

HasSubVendorId returns a boolean if a field has been set.

### GetVendorId

`func (o *GraphicsCard) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *GraphicsCard) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *GraphicsCard) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *GraphicsCard) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetComputeBlade

`func (o *GraphicsCard) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *GraphicsCard) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *GraphicsCard) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *GraphicsCard) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### SetComputeBladeNil

`func (o *GraphicsCard) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *GraphicsCard) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
### GetComputeBoard

`func (o *GraphicsCard) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *GraphicsCard) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *GraphicsCard) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *GraphicsCard) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### SetComputeBoardNil

`func (o *GraphicsCard) SetComputeBoardNil(b bool)`

 SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil

### UnsetComputeBoard
`func (o *GraphicsCard) UnsetComputeBoard()`

UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
### GetComputeRackUnit

`func (o *GraphicsCard) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *GraphicsCard) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *GraphicsCard) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *GraphicsCard) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### SetComputeRackUnitNil

`func (o *GraphicsCard) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *GraphicsCard) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
### GetEquipmentEnclosureElement

`func (o *GraphicsCard) GetEquipmentEnclosureElement() EquipmentEnclosureElementRelationship`

GetEquipmentEnclosureElement returns the EquipmentEnclosureElement field if non-nil, zero value otherwise.

### GetEquipmentEnclosureElementOk

`func (o *GraphicsCard) GetEquipmentEnclosureElementOk() (*EquipmentEnclosureElementRelationship, bool)`

GetEquipmentEnclosureElementOk returns a tuple with the EquipmentEnclosureElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEnclosureElement

`func (o *GraphicsCard) SetEquipmentEnclosureElement(v EquipmentEnclosureElementRelationship)`

SetEquipmentEnclosureElement sets EquipmentEnclosureElement field to given value.

### HasEquipmentEnclosureElement

`func (o *GraphicsCard) HasEquipmentEnclosureElement() bool`

HasEquipmentEnclosureElement returns a boolean if a field has been set.

### SetEquipmentEnclosureElementNil

`func (o *GraphicsCard) SetEquipmentEnclosureElementNil(b bool)`

 SetEquipmentEnclosureElementNil sets the value for EquipmentEnclosureElement to be an explicit nil

### UnsetEquipmentEnclosureElement
`func (o *GraphicsCard) UnsetEquipmentEnclosureElement()`

UnsetEquipmentEnclosureElement ensures that no value is present for EquipmentEnclosureElement, not even an explicit nil
### GetEquipmentInterconnect

`func (o *GraphicsCard) GetEquipmentInterconnect() EquipmentInterconnectRelationship`

GetEquipmentInterconnect returns the EquipmentInterconnect field if non-nil, zero value otherwise.

### GetEquipmentInterconnectOk

`func (o *GraphicsCard) GetEquipmentInterconnectOk() (*EquipmentInterconnectRelationship, bool)`

GetEquipmentInterconnectOk returns a tuple with the EquipmentInterconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentInterconnect

`func (o *GraphicsCard) SetEquipmentInterconnect(v EquipmentInterconnectRelationship)`

SetEquipmentInterconnect sets EquipmentInterconnect field to given value.

### HasEquipmentInterconnect

`func (o *GraphicsCard) HasEquipmentInterconnect() bool`

HasEquipmentInterconnect returns a boolean if a field has been set.

### SetEquipmentInterconnectNil

`func (o *GraphicsCard) SetEquipmentInterconnectNil(b bool)`

 SetEquipmentInterconnectNil sets the value for EquipmentInterconnect to be an explicit nil

### UnsetEquipmentInterconnect
`func (o *GraphicsCard) UnsetEquipmentInterconnect()`

UnsetEquipmentInterconnect ensures that no value is present for EquipmentInterconnect, not even an explicit nil
### GetEquipmentSharedGraphicsCard

`func (o *GraphicsCard) GetEquipmentSharedGraphicsCard() EquipmentSharedGraphicsCardRelationship`

GetEquipmentSharedGraphicsCard returns the EquipmentSharedGraphicsCard field if non-nil, zero value otherwise.

### GetEquipmentSharedGraphicsCardOk

`func (o *GraphicsCard) GetEquipmentSharedGraphicsCardOk() (*EquipmentSharedGraphicsCardRelationship, bool)`

GetEquipmentSharedGraphicsCardOk returns a tuple with the EquipmentSharedGraphicsCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSharedGraphicsCard

`func (o *GraphicsCard) SetEquipmentSharedGraphicsCard(v EquipmentSharedGraphicsCardRelationship)`

SetEquipmentSharedGraphicsCard sets EquipmentSharedGraphicsCard field to given value.

### HasEquipmentSharedGraphicsCard

`func (o *GraphicsCard) HasEquipmentSharedGraphicsCard() bool`

HasEquipmentSharedGraphicsCard returns a boolean if a field has been set.

### SetEquipmentSharedGraphicsCardNil

`func (o *GraphicsCard) SetEquipmentSharedGraphicsCardNil(b bool)`

 SetEquipmentSharedGraphicsCardNil sets the value for EquipmentSharedGraphicsCard to be an explicit nil

### UnsetEquipmentSharedGraphicsCard
`func (o *GraphicsCard) UnsetEquipmentSharedGraphicsCard()`

UnsetEquipmentSharedGraphicsCard ensures that no value is present for EquipmentSharedGraphicsCard, not even an explicit nil
### GetGraphicsControllers

`func (o *GraphicsCard) GetGraphicsControllers() []GraphicsControllerRelationship`

GetGraphicsControllers returns the GraphicsControllers field if non-nil, zero value otherwise.

### GetGraphicsControllersOk

`func (o *GraphicsCard) GetGraphicsControllersOk() (*[]GraphicsControllerRelationship, bool)`

GetGraphicsControllersOk returns a tuple with the GraphicsControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsControllers

`func (o *GraphicsCard) SetGraphicsControllers(v []GraphicsControllerRelationship)`

SetGraphicsControllers sets GraphicsControllers field to given value.

### HasGraphicsControllers

`func (o *GraphicsCard) HasGraphicsControllers() bool`

HasGraphicsControllers returns a boolean if a field has been set.

### SetGraphicsControllersNil

`func (o *GraphicsCard) SetGraphicsControllersNil(b bool)`

 SetGraphicsControllersNil sets the value for GraphicsControllers to be an explicit nil

### UnsetGraphicsControllers
`func (o *GraphicsCard) UnsetGraphicsControllers()`

UnsetGraphicsControllers ensures that no value is present for GraphicsControllers, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *GraphicsCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *GraphicsCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *GraphicsCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *GraphicsCard) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *GraphicsCard) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *GraphicsCard) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetPciDevice

`func (o *GraphicsCard) GetPciDevice() PciDeviceRelationship`

GetPciDevice returns the PciDevice field if non-nil, zero value otherwise.

### GetPciDeviceOk

`func (o *GraphicsCard) GetPciDeviceOk() (*PciDeviceRelationship, bool)`

GetPciDeviceOk returns a tuple with the PciDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevice

`func (o *GraphicsCard) SetPciDevice(v PciDeviceRelationship)`

SetPciDevice sets PciDevice field to given value.

### HasPciDevice

`func (o *GraphicsCard) HasPciDevice() bool`

HasPciDevice returns a boolean if a field has been set.

### SetPciDeviceNil

`func (o *GraphicsCard) SetPciDeviceNil(b bool)`

 SetPciDeviceNil sets the value for PciDevice to be an explicit nil

### UnsetPciDevice
`func (o *GraphicsCard) UnsetPciDevice()`

UnsetPciDevice ensures that no value is present for PciDevice, not even an explicit nil
### GetPciNode

`func (o *GraphicsCard) GetPciNode() PciNodeRelationship`

GetPciNode returns the PciNode field if non-nil, zero value otherwise.

### GetPciNodeOk

`func (o *GraphicsCard) GetPciNodeOk() (*PciNodeRelationship, bool)`

GetPciNodeOk returns a tuple with the PciNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciNode

`func (o *GraphicsCard) SetPciNode(v PciNodeRelationship)`

SetPciNode sets PciNode field to given value.

### HasPciNode

`func (o *GraphicsCard) HasPciNode() bool`

HasPciNode returns a boolean if a field has been set.

### SetPciNodeNil

`func (o *GraphicsCard) SetPciNodeNil(b bool)`

 SetPciNodeNil sets the value for PciNode to be an explicit nil

### UnsetPciNode
`func (o *GraphicsCard) UnsetPciNode()`

UnsetPciNode ensures that no value is present for PciNode, not even an explicit nil
### GetRegisteredDevice

`func (o *GraphicsCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *GraphicsCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *GraphicsCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *GraphicsCard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *GraphicsCard) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *GraphicsCard) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetRunningFirmware

`func (o *GraphicsCard) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *GraphicsCard) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *GraphicsCard) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *GraphicsCard) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *GraphicsCard) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *GraphicsCard) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


