# GraphicsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "graphics.Card"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "graphics.Card"]
**CardId** | Pointer to **int64** | The id of the graphics card. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The device id of the graphics card. | [optional] [readonly] 
**ExpanderSlot** | Pointer to **string** | The expander slot information of the card. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The firmware version of the graphics card. | [optional] [readonly] 
**Mode** | Pointer to **string** | The current mode of the graphics card. | [optional] [readonly] 
**NumGpus** | Pointer to **string** | The number of controllers under each card. | [optional] 
**OperState** | Pointer to **string** | The current operational state of the graphics card. | [optional] [readonly] 
**PciAddress** | Pointer to **string** | The PCI address of the graphics card. | [optional] [readonly] 
**PciAddressList** | Pointer to **string** | This list contains the PCI address of all controllers for corresponding card. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The PCI slot name of the graphics card. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**GraphicsControllers** | Pointer to [**[]GraphicsControllerRelationship**](GraphicsControllerRelationship.md) | An array of relationships to graphicsController resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
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


