# PciLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pci.Link"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pci.Link"]
**Adapter** | Pointer to **string** | The name of the PCI device. | [optional] [readonly] 
**LinkSpeed** | Pointer to **string** | The upstream link speed of the PCI device. | [optional] [readonly] 
**LinkStatus** | Pointer to **string** | The upstream link status of the PCI device. | [optional] [readonly] 
**LinkWidth** | Pointer to **string** | The upstream link width of the PCI device. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The slot name of the PCI device. | [optional] [readonly] 
**SlotStatus** | Pointer to **string** | The health information of the PCI device. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PciSwitch** | Pointer to [**PciSwitchRelationship**](PciSwitchRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPciLink

`func NewPciLink(classId string, objectType string, ) *PciLink`

NewPciLink instantiates a new PciLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciLinkWithDefaults

`func NewPciLinkWithDefaults() *PciLink`

NewPciLinkWithDefaults instantiates a new PciLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PciLink) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PciLink) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PciLink) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PciLink) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PciLink) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PciLink) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapter

`func (o *PciLink) GetAdapter() string`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *PciLink) GetAdapterOk() (*string, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *PciLink) SetAdapter(v string)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *PciLink) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *PciLink) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *PciLink) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *PciLink) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *PciLink) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *PciLink) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *PciLink) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *PciLink) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *PciLink) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetLinkWidth

`func (o *PciLink) GetLinkWidth() string`

GetLinkWidth returns the LinkWidth field if non-nil, zero value otherwise.

### GetLinkWidthOk

`func (o *PciLink) GetLinkWidthOk() (*string, bool)`

GetLinkWidthOk returns a tuple with the LinkWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkWidth

`func (o *PciLink) SetLinkWidth(v string)`

SetLinkWidth sets LinkWidth field to given value.

### HasLinkWidth

`func (o *PciLink) HasLinkWidth() bool`

HasLinkWidth returns a boolean if a field has been set.

### GetPciSlot

`func (o *PciLink) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *PciLink) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *PciLink) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *PciLink) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetSlotStatus

`func (o *PciLink) GetSlotStatus() string`

GetSlotStatus returns the SlotStatus field if non-nil, zero value otherwise.

### GetSlotStatusOk

`func (o *PciLink) GetSlotStatusOk() (*string, bool)`

GetSlotStatusOk returns a tuple with the SlotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotStatus

`func (o *PciLink) SetSlotStatus(v string)`

SetSlotStatus sets SlotStatus field to given value.

### HasSlotStatus

`func (o *PciLink) HasSlotStatus() bool`

HasSlotStatus returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *PciLink) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciLink) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciLink) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciLink) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPciSwitch

`func (o *PciLink) GetPciSwitch() PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *PciLink) GetPciSwitchOk() (*PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *PciLink) SetPciSwitch(v PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *PciLink) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PciLink) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciLink) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciLink) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciLink) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


