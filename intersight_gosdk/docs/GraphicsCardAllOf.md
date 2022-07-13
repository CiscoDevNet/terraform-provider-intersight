# GraphicsCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "graphics.Card"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "graphics.Card"]
**CardId** | Pointer to **int64** | The id of the graphics card. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The device id of the graphics card. | [optional] [readonly] 
**ExpanderSlot** | Pointer to **string** | The expander slot information of the card. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The firmware version of the graphics card. | [optional] [readonly] 
**GpuId** | Pointer to **string** | The identifier of the graphics processor unit. | [optional] [readonly] 
**Mode** | Pointer to **string** | The current mode of the graphics card. | [optional] [readonly] 
**NumGpus** | Pointer to **string** | The number of controllers under each card. | [optional] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | The current operational state of the graphics card. | [optional] [readonly] 
**PciAddress** | Pointer to **string** | The PCI address of the graphics card. | [optional] [readonly] 
**PciAddressList** | Pointer to **string** | This list contains the PCI address of all controllers for corresponding card. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The PCI slot name of the graphics card. | [optional] [readonly] 
**SubDeviceId** | Pointer to **int64** | The sub device id of the graphics processor unit. | [optional] [readonly] 
**SubVendorId** | Pointer to **int64** | The sub vendor id of the graphics processor unit. | [optional] [readonly] 
**VendorId** | Pointer to **int64** | The vendor id of the graphics processor unit. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**GraphicsControllers** | Pointer to [**[]GraphicsControllerRelationship**](GraphicsControllerRelationship.md) | An array of relationships to graphicsController resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PciDevice** | Pointer to [**PciDeviceRelationship**](PciDeviceRelationship.md) |  | [optional] 
**PciNode** | Pointer to [**PciNodeRelationship**](PciNodeRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 

## Methods

### NewGraphicsCardAllOf

`func NewGraphicsCardAllOf(classId string, objectType string, ) *GraphicsCardAllOf`

NewGraphicsCardAllOf instantiates a new GraphicsCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphicsCardAllOfWithDefaults

`func NewGraphicsCardAllOfWithDefaults() *GraphicsCardAllOf`

NewGraphicsCardAllOfWithDefaults instantiates a new GraphicsCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *GraphicsCardAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *GraphicsCardAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *GraphicsCardAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *GraphicsCardAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GraphicsCardAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GraphicsCardAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCardId

`func (o *GraphicsCardAllOf) GetCardId() int64`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *GraphicsCardAllOf) GetCardIdOk() (*int64, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *GraphicsCardAllOf) SetCardId(v int64)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *GraphicsCardAllOf) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetDeviceId

`func (o *GraphicsCardAllOf) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GraphicsCardAllOf) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GraphicsCardAllOf) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GraphicsCardAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExpanderSlot

`func (o *GraphicsCardAllOf) GetExpanderSlot() string`

GetExpanderSlot returns the ExpanderSlot field if non-nil, zero value otherwise.

### GetExpanderSlotOk

`func (o *GraphicsCardAllOf) GetExpanderSlotOk() (*string, bool)`

GetExpanderSlotOk returns a tuple with the ExpanderSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderSlot

`func (o *GraphicsCardAllOf) SetExpanderSlot(v string)`

SetExpanderSlot sets ExpanderSlot field to given value.

### HasExpanderSlot

`func (o *GraphicsCardAllOf) HasExpanderSlot() bool`

HasExpanderSlot returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *GraphicsCardAllOf) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *GraphicsCardAllOf) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *GraphicsCardAllOf) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *GraphicsCardAllOf) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetGpuId

`func (o *GraphicsCardAllOf) GetGpuId() string`

GetGpuId returns the GpuId field if non-nil, zero value otherwise.

### GetGpuIdOk

`func (o *GraphicsCardAllOf) GetGpuIdOk() (*string, bool)`

GetGpuIdOk returns a tuple with the GpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuId

`func (o *GraphicsCardAllOf) SetGpuId(v string)`

SetGpuId sets GpuId field to given value.

### HasGpuId

`func (o *GraphicsCardAllOf) HasGpuId() bool`

HasGpuId returns a boolean if a field has been set.

### GetMode

`func (o *GraphicsCardAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GraphicsCardAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GraphicsCardAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GraphicsCardAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumGpus

`func (o *GraphicsCardAllOf) GetNumGpus() string`

GetNumGpus returns the NumGpus field if non-nil, zero value otherwise.

### GetNumGpusOk

`func (o *GraphicsCardAllOf) GetNumGpusOk() (*string, bool)`

GetNumGpusOk returns a tuple with the NumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGpus

`func (o *GraphicsCardAllOf) SetNumGpus(v string)`

SetNumGpus sets NumGpus field to given value.

### HasNumGpus

`func (o *GraphicsCardAllOf) HasNumGpus() bool`

HasNumGpus returns a boolean if a field has been set.

### GetOperReason

`func (o *GraphicsCardAllOf) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *GraphicsCardAllOf) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *GraphicsCardAllOf) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *GraphicsCardAllOf) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *GraphicsCardAllOf) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *GraphicsCardAllOf) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *GraphicsCardAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *GraphicsCardAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *GraphicsCardAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *GraphicsCardAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPciAddress

`func (o *GraphicsCardAllOf) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *GraphicsCardAllOf) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *GraphicsCardAllOf) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *GraphicsCardAllOf) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPciAddressList

`func (o *GraphicsCardAllOf) GetPciAddressList() string`

GetPciAddressList returns the PciAddressList field if non-nil, zero value otherwise.

### GetPciAddressListOk

`func (o *GraphicsCardAllOf) GetPciAddressListOk() (*string, bool)`

GetPciAddressListOk returns a tuple with the PciAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddressList

`func (o *GraphicsCardAllOf) SetPciAddressList(v string)`

SetPciAddressList sets PciAddressList field to given value.

### HasPciAddressList

`func (o *GraphicsCardAllOf) HasPciAddressList() bool`

HasPciAddressList returns a boolean if a field has been set.

### GetPciSlot

`func (o *GraphicsCardAllOf) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *GraphicsCardAllOf) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *GraphicsCardAllOf) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *GraphicsCardAllOf) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetSubDeviceId

`func (o *GraphicsCardAllOf) GetSubDeviceId() int64`

GetSubDeviceId returns the SubDeviceId field if non-nil, zero value otherwise.

### GetSubDeviceIdOk

`func (o *GraphicsCardAllOf) GetSubDeviceIdOk() (*int64, bool)`

GetSubDeviceIdOk returns a tuple with the SubDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDeviceId

`func (o *GraphicsCardAllOf) SetSubDeviceId(v int64)`

SetSubDeviceId sets SubDeviceId field to given value.

### HasSubDeviceId

`func (o *GraphicsCardAllOf) HasSubDeviceId() bool`

HasSubDeviceId returns a boolean if a field has been set.

### GetSubVendorId

`func (o *GraphicsCardAllOf) GetSubVendorId() int64`

GetSubVendorId returns the SubVendorId field if non-nil, zero value otherwise.

### GetSubVendorIdOk

`func (o *GraphicsCardAllOf) GetSubVendorIdOk() (*int64, bool)`

GetSubVendorIdOk returns a tuple with the SubVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVendorId

`func (o *GraphicsCardAllOf) SetSubVendorId(v int64)`

SetSubVendorId sets SubVendorId field to given value.

### HasSubVendorId

`func (o *GraphicsCardAllOf) HasSubVendorId() bool`

HasSubVendorId returns a boolean if a field has been set.

### GetVendorId

`func (o *GraphicsCardAllOf) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *GraphicsCardAllOf) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *GraphicsCardAllOf) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *GraphicsCardAllOf) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetComputeBlade

`func (o *GraphicsCardAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *GraphicsCardAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *GraphicsCardAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *GraphicsCardAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *GraphicsCardAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *GraphicsCardAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *GraphicsCardAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *GraphicsCardAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *GraphicsCardAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *GraphicsCardAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *GraphicsCardAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *GraphicsCardAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetGraphicsControllers

`func (o *GraphicsCardAllOf) GetGraphicsControllers() []GraphicsControllerRelationship`

GetGraphicsControllers returns the GraphicsControllers field if non-nil, zero value otherwise.

### GetGraphicsControllersOk

`func (o *GraphicsCardAllOf) GetGraphicsControllersOk() (*[]GraphicsControllerRelationship, bool)`

GetGraphicsControllersOk returns a tuple with the GraphicsControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsControllers

`func (o *GraphicsCardAllOf) SetGraphicsControllers(v []GraphicsControllerRelationship)`

SetGraphicsControllers sets GraphicsControllers field to given value.

### HasGraphicsControllers

`func (o *GraphicsCardAllOf) HasGraphicsControllers() bool`

HasGraphicsControllers returns a boolean if a field has been set.

### SetGraphicsControllersNil

`func (o *GraphicsCardAllOf) SetGraphicsControllersNil(b bool)`

 SetGraphicsControllersNil sets the value for GraphicsControllers to be an explicit nil

### UnsetGraphicsControllers
`func (o *GraphicsCardAllOf) UnsetGraphicsControllers()`

UnsetGraphicsControllers ensures that no value is present for GraphicsControllers, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *GraphicsCardAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *GraphicsCardAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *GraphicsCardAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *GraphicsCardAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPciDevice

`func (o *GraphicsCardAllOf) GetPciDevice() PciDeviceRelationship`

GetPciDevice returns the PciDevice field if non-nil, zero value otherwise.

### GetPciDeviceOk

`func (o *GraphicsCardAllOf) GetPciDeviceOk() (*PciDeviceRelationship, bool)`

GetPciDeviceOk returns a tuple with the PciDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevice

`func (o *GraphicsCardAllOf) SetPciDevice(v PciDeviceRelationship)`

SetPciDevice sets PciDevice field to given value.

### HasPciDevice

`func (o *GraphicsCardAllOf) HasPciDevice() bool`

HasPciDevice returns a boolean if a field has been set.

### GetPciNode

`func (o *GraphicsCardAllOf) GetPciNode() PciNodeRelationship`

GetPciNode returns the PciNode field if non-nil, zero value otherwise.

### GetPciNodeOk

`func (o *GraphicsCardAllOf) GetPciNodeOk() (*PciNodeRelationship, bool)`

GetPciNodeOk returns a tuple with the PciNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciNode

`func (o *GraphicsCardAllOf) SetPciNode(v PciNodeRelationship)`

SetPciNode sets PciNode field to given value.

### HasPciNode

`func (o *GraphicsCardAllOf) HasPciNode() bool`

HasPciNode returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *GraphicsCardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *GraphicsCardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *GraphicsCardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *GraphicsCardAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *GraphicsCardAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *GraphicsCardAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *GraphicsCardAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *GraphicsCardAllOf) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *GraphicsCardAllOf) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *GraphicsCardAllOf) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


