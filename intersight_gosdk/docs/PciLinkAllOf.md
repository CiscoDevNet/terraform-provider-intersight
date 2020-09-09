# PciLinkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adapter** | Pointer to **string** | The name of the PCI device. | [optional] [readonly] 
**LinkSpeed** | Pointer to **string** | The upstream link speed of the PCI device. | [optional] [readonly] 
**LinkStatus** | Pointer to **string** | The upstream link status of the PCI device. | [optional] [readonly] 
**LinkWidth** | Pointer to **string** | The upstream link width of the PCI device. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The slot name of the PCI device. | [optional] [readonly] 
**SlotStatus** | Pointer to **string** | The health information of the PCI device. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PciSwitch** | Pointer to [**PciSwitchRelationship**](pci.Switch.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewPciLinkAllOf

`func NewPciLinkAllOf() *PciLinkAllOf`

NewPciLinkAllOf instantiates a new PciLinkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciLinkAllOfWithDefaults

`func NewPciLinkAllOfWithDefaults() *PciLinkAllOf`

NewPciLinkAllOfWithDefaults instantiates a new PciLinkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapter

`func (o *PciLinkAllOf) GetAdapter() string`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *PciLinkAllOf) GetAdapterOk() (*string, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *PciLinkAllOf) SetAdapter(v string)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *PciLinkAllOf) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *PciLinkAllOf) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *PciLinkAllOf) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *PciLinkAllOf) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *PciLinkAllOf) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *PciLinkAllOf) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *PciLinkAllOf) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *PciLinkAllOf) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *PciLinkAllOf) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetLinkWidth

`func (o *PciLinkAllOf) GetLinkWidth() string`

GetLinkWidth returns the LinkWidth field if non-nil, zero value otherwise.

### GetLinkWidthOk

`func (o *PciLinkAllOf) GetLinkWidthOk() (*string, bool)`

GetLinkWidthOk returns a tuple with the LinkWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkWidth

`func (o *PciLinkAllOf) SetLinkWidth(v string)`

SetLinkWidth sets LinkWidth field to given value.

### HasLinkWidth

`func (o *PciLinkAllOf) HasLinkWidth() bool`

HasLinkWidth returns a boolean if a field has been set.

### GetPciSlot

`func (o *PciLinkAllOf) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *PciLinkAllOf) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *PciLinkAllOf) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *PciLinkAllOf) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetSlotStatus

`func (o *PciLinkAllOf) GetSlotStatus() string`

GetSlotStatus returns the SlotStatus field if non-nil, zero value otherwise.

### GetSlotStatusOk

`func (o *PciLinkAllOf) GetSlotStatusOk() (*string, bool)`

GetSlotStatusOk returns a tuple with the SlotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotStatus

`func (o *PciLinkAllOf) SetSlotStatus(v string)`

SetSlotStatus sets SlotStatus field to given value.

### HasSlotStatus

`func (o *PciLinkAllOf) HasSlotStatus() bool`

HasSlotStatus returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *PciLinkAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciLinkAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciLinkAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciLinkAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPciSwitch

`func (o *PciLinkAllOf) GetPciSwitch() PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *PciLinkAllOf) GetPciSwitchOk() (*PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *PciLinkAllOf) SetPciSwitch(v PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *PciLinkAllOf) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PciLinkAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciLinkAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciLinkAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciLinkAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


