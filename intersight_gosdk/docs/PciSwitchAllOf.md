# PciSwitchAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | The device id of the switch. | [optional] [readonly] 
**Health** | Pointer to **string** | The composite health of the switch. | [optional] [readonly] 
**NumOfAdaptors** | Pointer to **string** | The number of GPUs and PCI adapters connected the switch. | [optional] [readonly] 
**PciAddress** | Pointer to **string** | The PCI address of the switch. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The PCI slot name of the switch. | [optional] [readonly] 
**ProductName** | Pointer to **string** | The model information for the switch. | [optional] [readonly] 
**ProductRevision** | Pointer to **string** | The product revision of the switch. | [optional] [readonly] 
**SubDeviceId** | Pointer to **string** | The sub device id of the switch. | [optional] [readonly] 
**SubVendorId** | Pointer to **string** | The sub vendor id of the switch. | [optional] [readonly] 
**Temperature** | Pointer to **string** | The current temperature of the switch. | [optional] [readonly] 
**Type** | Pointer to **string** | The type information of the switch. | [optional] 
**VendorId** | Pointer to **string** | The vendor id of the switch. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**Links** | Pointer to [**[]PciLinkRelationship**](pci.Link.Relationship.md) | An array of relationships to pciLink resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 

## Methods

### NewPciSwitchAllOf

`func NewPciSwitchAllOf() *PciSwitchAllOf`

NewPciSwitchAllOf instantiates a new PciSwitchAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciSwitchAllOfWithDefaults

`func NewPciSwitchAllOfWithDefaults() *PciSwitchAllOf`

NewPciSwitchAllOfWithDefaults instantiates a new PciSwitchAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *PciSwitchAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PciSwitchAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PciSwitchAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *PciSwitchAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHealth

`func (o *PciSwitchAllOf) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PciSwitchAllOf) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PciSwitchAllOf) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PciSwitchAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetNumOfAdaptors

`func (o *PciSwitchAllOf) GetNumOfAdaptors() string`

GetNumOfAdaptors returns the NumOfAdaptors field if non-nil, zero value otherwise.

### GetNumOfAdaptorsOk

`func (o *PciSwitchAllOf) GetNumOfAdaptorsOk() (*string, bool)`

GetNumOfAdaptorsOk returns a tuple with the NumOfAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfAdaptors

`func (o *PciSwitchAllOf) SetNumOfAdaptors(v string)`

SetNumOfAdaptors sets NumOfAdaptors field to given value.

### HasNumOfAdaptors

`func (o *PciSwitchAllOf) HasNumOfAdaptors() bool`

HasNumOfAdaptors returns a boolean if a field has been set.

### GetPciAddress

`func (o *PciSwitchAllOf) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *PciSwitchAllOf) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *PciSwitchAllOf) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *PciSwitchAllOf) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPciSlot

`func (o *PciSwitchAllOf) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *PciSwitchAllOf) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *PciSwitchAllOf) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *PciSwitchAllOf) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetProductName

`func (o *PciSwitchAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PciSwitchAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PciSwitchAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *PciSwitchAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductRevision

`func (o *PciSwitchAllOf) GetProductRevision() string`

GetProductRevision returns the ProductRevision field if non-nil, zero value otherwise.

### GetProductRevisionOk

`func (o *PciSwitchAllOf) GetProductRevisionOk() (*string, bool)`

GetProductRevisionOk returns a tuple with the ProductRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevision

`func (o *PciSwitchAllOf) SetProductRevision(v string)`

SetProductRevision sets ProductRevision field to given value.

### HasProductRevision

`func (o *PciSwitchAllOf) HasProductRevision() bool`

HasProductRevision returns a boolean if a field has been set.

### GetSubDeviceId

`func (o *PciSwitchAllOf) GetSubDeviceId() string`

GetSubDeviceId returns the SubDeviceId field if non-nil, zero value otherwise.

### GetSubDeviceIdOk

`func (o *PciSwitchAllOf) GetSubDeviceIdOk() (*string, bool)`

GetSubDeviceIdOk returns a tuple with the SubDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDeviceId

`func (o *PciSwitchAllOf) SetSubDeviceId(v string)`

SetSubDeviceId sets SubDeviceId field to given value.

### HasSubDeviceId

`func (o *PciSwitchAllOf) HasSubDeviceId() bool`

HasSubDeviceId returns a boolean if a field has been set.

### GetSubVendorId

`func (o *PciSwitchAllOf) GetSubVendorId() string`

GetSubVendorId returns the SubVendorId field if non-nil, zero value otherwise.

### GetSubVendorIdOk

`func (o *PciSwitchAllOf) GetSubVendorIdOk() (*string, bool)`

GetSubVendorIdOk returns a tuple with the SubVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVendorId

`func (o *PciSwitchAllOf) SetSubVendorId(v string)`

SetSubVendorId sets SubVendorId field to given value.

### HasSubVendorId

`func (o *PciSwitchAllOf) HasSubVendorId() bool`

HasSubVendorId returns a boolean if a field has been set.

### GetTemperature

`func (o *PciSwitchAllOf) GetTemperature() string`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *PciSwitchAllOf) GetTemperatureOk() (*string, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *PciSwitchAllOf) SetTemperature(v string)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *PciSwitchAllOf) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetType

`func (o *PciSwitchAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PciSwitchAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PciSwitchAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PciSwitchAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendorId

`func (o *PciSwitchAllOf) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *PciSwitchAllOf) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *PciSwitchAllOf) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *PciSwitchAllOf) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetComputeBoard

`func (o *PciSwitchAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *PciSwitchAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *PciSwitchAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *PciSwitchAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *PciSwitchAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciSwitchAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciSwitchAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciSwitchAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLinks

`func (o *PciSwitchAllOf) GetLinks() []PciLinkRelationship`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PciSwitchAllOf) GetLinksOk() (*[]PciLinkRelationship, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PciSwitchAllOf) SetLinks(v []PciLinkRelationship)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PciSwitchAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *PciSwitchAllOf) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *PciSwitchAllOf) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRegisteredDevice

`func (o *PciSwitchAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciSwitchAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciSwitchAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciSwitchAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *PciSwitchAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *PciSwitchAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *PciSwitchAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *PciSwitchAllOf) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *PciSwitchAllOf) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *PciSwitchAllOf) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


