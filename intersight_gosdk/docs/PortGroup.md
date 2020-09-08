# PortGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transport** | Pointer to **string** | Type of port group. Values are Eth or Fc. | [optional] [readonly] 
**EquipmentSharedIoModule** | Pointer to [**EquipmentSharedIoModuleRelationship**](equipment.SharedIoModule.Relationship.md) |  | [optional] 
**EquipmentSwitchCard** | Pointer to [**EquipmentSwitchCardRelationship**](equipment.SwitchCard.Relationship.md) |  | [optional] 
**EthernetPorts** | Pointer to [**[]EtherPhysicalPortRelationship**](ether.PhysicalPort.Relationship.md) | An array of relationships to etherPhysicalPort resources. | [optional] [readonly] 
**FcPorts** | Pointer to [**[]FcPhysicalPortRelationship**](fc.PhysicalPort.Relationship.md) | An array of relationships to fcPhysicalPort resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**SubGroups** | Pointer to [**[]PortSubGroupRelationship**](port.SubGroup.Relationship.md) | An array of relationships to portSubGroup resources. | [optional] [readonly] 

## Methods

### NewPortGroup

`func NewPortGroup() *PortGroup`

NewPortGroup instantiates a new PortGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortGroupWithDefaults

`func NewPortGroupWithDefaults() *PortGroup`

NewPortGroupWithDefaults instantiates a new PortGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransport

`func (o *PortGroup) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *PortGroup) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *PortGroup) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *PortGroup) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetEquipmentSharedIoModule

`func (o *PortGroup) GetEquipmentSharedIoModule() EquipmentSharedIoModuleRelationship`

GetEquipmentSharedIoModule returns the EquipmentSharedIoModule field if non-nil, zero value otherwise.

### GetEquipmentSharedIoModuleOk

`func (o *PortGroup) GetEquipmentSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool)`

GetEquipmentSharedIoModuleOk returns a tuple with the EquipmentSharedIoModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSharedIoModule

`func (o *PortGroup) SetEquipmentSharedIoModule(v EquipmentSharedIoModuleRelationship)`

SetEquipmentSharedIoModule sets EquipmentSharedIoModule field to given value.

### HasEquipmentSharedIoModule

`func (o *PortGroup) HasEquipmentSharedIoModule() bool`

HasEquipmentSharedIoModule returns a boolean if a field has been set.

### GetEquipmentSwitchCard

`func (o *PortGroup) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *PortGroup) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *PortGroup) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *PortGroup) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### GetEthernetPorts

`func (o *PortGroup) GetEthernetPorts() []EtherPhysicalPortRelationship`

GetEthernetPorts returns the EthernetPorts field if non-nil, zero value otherwise.

### GetEthernetPortsOk

`func (o *PortGroup) GetEthernetPortsOk() (*[]EtherPhysicalPortRelationship, bool)`

GetEthernetPortsOk returns a tuple with the EthernetPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPorts

`func (o *PortGroup) SetEthernetPorts(v []EtherPhysicalPortRelationship)`

SetEthernetPorts sets EthernetPorts field to given value.

### HasEthernetPorts

`func (o *PortGroup) HasEthernetPorts() bool`

HasEthernetPorts returns a boolean if a field has been set.

### SetEthernetPortsNil

`func (o *PortGroup) SetEthernetPortsNil(b bool)`

 SetEthernetPortsNil sets the value for EthernetPorts to be an explicit nil

### UnsetEthernetPorts
`func (o *PortGroup) UnsetEthernetPorts()`

UnsetEthernetPorts ensures that no value is present for EthernetPorts, not even an explicit nil
### GetFcPorts

`func (o *PortGroup) GetFcPorts() []FcPhysicalPortRelationship`

GetFcPorts returns the FcPorts field if non-nil, zero value otherwise.

### GetFcPortsOk

`func (o *PortGroup) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool)`

GetFcPortsOk returns a tuple with the FcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPorts

`func (o *PortGroup) SetFcPorts(v []FcPhysicalPortRelationship)`

SetFcPorts sets FcPorts field to given value.

### HasFcPorts

`func (o *PortGroup) HasFcPorts() bool`

HasFcPorts returns a boolean if a field has been set.

### SetFcPortsNil

`func (o *PortGroup) SetFcPortsNil(b bool)`

 SetFcPortsNil sets the value for FcPorts to be an explicit nil

### UnsetFcPorts
`func (o *PortGroup) UnsetFcPorts()`

UnsetFcPorts ensures that no value is present for FcPorts, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *PortGroup) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PortGroup) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PortGroup) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PortGroup) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PortGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PortGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PortGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PortGroup) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSubGroups

`func (o *PortGroup) GetSubGroups() []PortSubGroupRelationship`

GetSubGroups returns the SubGroups field if non-nil, zero value otherwise.

### GetSubGroupsOk

`func (o *PortGroup) GetSubGroupsOk() (*[]PortSubGroupRelationship, bool)`

GetSubGroupsOk returns a tuple with the SubGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGroups

`func (o *PortGroup) SetSubGroups(v []PortSubGroupRelationship)`

SetSubGroups sets SubGroups field to given value.

### HasSubGroups

`func (o *PortGroup) HasSubGroups() bool`

HasSubGroups returns a boolean if a field has been set.

### SetSubGroupsNil

`func (o *PortGroup) SetSubGroupsNil(b bool)`

 SetSubGroupsNil sets the value for SubGroups to be an explicit nil

### UnsetSubGroups
`func (o *PortGroup) UnsetSubGroups()`

UnsetSubGroups ensures that no value is present for SubGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


