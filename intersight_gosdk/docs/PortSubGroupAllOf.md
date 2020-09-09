# PortSubGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transport** | Pointer to **string** | Type of port sub-group. Values are Eth or Fc. | [optional] [readonly] 
**EthernetPorts** | Pointer to [**[]EtherPhysicalPortRelationship**](ether.PhysicalPort.Relationship.md) | An array of relationships to etherPhysicalPort resources. | [optional] [readonly] 
**FcPorts** | Pointer to [**[]FcPhysicalPortRelationship**](fc.PhysicalPort.Relationship.md) | An array of relationships to fcPhysicalPort resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PortGroup** | Pointer to [**PortGroupRelationship**](port.Group.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewPortSubGroupAllOf

`func NewPortSubGroupAllOf() *PortSubGroupAllOf`

NewPortSubGroupAllOf instantiates a new PortSubGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortSubGroupAllOfWithDefaults

`func NewPortSubGroupAllOfWithDefaults() *PortSubGroupAllOf`

NewPortSubGroupAllOfWithDefaults instantiates a new PortSubGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransport

`func (o *PortSubGroupAllOf) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *PortSubGroupAllOf) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *PortSubGroupAllOf) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *PortSubGroupAllOf) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetEthernetPorts

`func (o *PortSubGroupAllOf) GetEthernetPorts() []EtherPhysicalPortRelationship`

GetEthernetPorts returns the EthernetPorts field if non-nil, zero value otherwise.

### GetEthernetPortsOk

`func (o *PortSubGroupAllOf) GetEthernetPortsOk() (*[]EtherPhysicalPortRelationship, bool)`

GetEthernetPortsOk returns a tuple with the EthernetPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPorts

`func (o *PortSubGroupAllOf) SetEthernetPorts(v []EtherPhysicalPortRelationship)`

SetEthernetPorts sets EthernetPorts field to given value.

### HasEthernetPorts

`func (o *PortSubGroupAllOf) HasEthernetPorts() bool`

HasEthernetPorts returns a boolean if a field has been set.

### SetEthernetPortsNil

`func (o *PortSubGroupAllOf) SetEthernetPortsNil(b bool)`

 SetEthernetPortsNil sets the value for EthernetPorts to be an explicit nil

### UnsetEthernetPorts
`func (o *PortSubGroupAllOf) UnsetEthernetPorts()`

UnsetEthernetPorts ensures that no value is present for EthernetPorts, not even an explicit nil
### GetFcPorts

`func (o *PortSubGroupAllOf) GetFcPorts() []FcPhysicalPortRelationship`

GetFcPorts returns the FcPorts field if non-nil, zero value otherwise.

### GetFcPortsOk

`func (o *PortSubGroupAllOf) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool)`

GetFcPortsOk returns a tuple with the FcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPorts

`func (o *PortSubGroupAllOf) SetFcPorts(v []FcPhysicalPortRelationship)`

SetFcPorts sets FcPorts field to given value.

### HasFcPorts

`func (o *PortSubGroupAllOf) HasFcPorts() bool`

HasFcPorts returns a boolean if a field has been set.

### SetFcPortsNil

`func (o *PortSubGroupAllOf) SetFcPortsNil(b bool)`

 SetFcPortsNil sets the value for FcPorts to be an explicit nil

### UnsetFcPorts
`func (o *PortSubGroupAllOf) UnsetFcPorts()`

UnsetFcPorts ensures that no value is present for FcPorts, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *PortSubGroupAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PortSubGroupAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PortSubGroupAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PortSubGroupAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroup

`func (o *PortSubGroupAllOf) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *PortSubGroupAllOf) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *PortSubGroupAllOf) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *PortSubGroupAllOf) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PortSubGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PortSubGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PortSubGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PortSubGroupAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


