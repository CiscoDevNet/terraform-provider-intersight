# PortSubGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "port.SubGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "port.SubGroup"]
**Transport** | Pointer to **string** | Type of port sub-group. Values are Eth or Fc. | [optional] [readonly] 
**EthernetPorts** | Pointer to [**[]EtherPhysicalPortRelationship**](EtherPhysicalPortRelationship.md) | An array of relationships to etherPhysicalPort resources. | [optional] [readonly] 
**FcPorts** | Pointer to [**[]FcPhysicalPortRelationship**](FcPhysicalPortRelationship.md) | An array of relationships to fcPhysicalPort resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PortGroup** | Pointer to [**PortGroupRelationship**](PortGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPortSubGroup

`func NewPortSubGroup(classId string, objectType string, ) *PortSubGroup`

NewPortSubGroup instantiates a new PortSubGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortSubGroupWithDefaults

`func NewPortSubGroupWithDefaults() *PortSubGroup`

NewPortSubGroupWithDefaults instantiates a new PortSubGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PortSubGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PortSubGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PortSubGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PortSubGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PortSubGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PortSubGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTransport

`func (o *PortSubGroup) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *PortSubGroup) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *PortSubGroup) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *PortSubGroup) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetEthernetPorts

`func (o *PortSubGroup) GetEthernetPorts() []EtherPhysicalPortRelationship`

GetEthernetPorts returns the EthernetPorts field if non-nil, zero value otherwise.

### GetEthernetPortsOk

`func (o *PortSubGroup) GetEthernetPortsOk() (*[]EtherPhysicalPortRelationship, bool)`

GetEthernetPortsOk returns a tuple with the EthernetPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPorts

`func (o *PortSubGroup) SetEthernetPorts(v []EtherPhysicalPortRelationship)`

SetEthernetPorts sets EthernetPorts field to given value.

### HasEthernetPorts

`func (o *PortSubGroup) HasEthernetPorts() bool`

HasEthernetPorts returns a boolean if a field has been set.

### SetEthernetPortsNil

`func (o *PortSubGroup) SetEthernetPortsNil(b bool)`

 SetEthernetPortsNil sets the value for EthernetPorts to be an explicit nil

### UnsetEthernetPorts
`func (o *PortSubGroup) UnsetEthernetPorts()`

UnsetEthernetPorts ensures that no value is present for EthernetPorts, not even an explicit nil
### GetFcPorts

`func (o *PortSubGroup) GetFcPorts() []FcPhysicalPortRelationship`

GetFcPorts returns the FcPorts field if non-nil, zero value otherwise.

### GetFcPortsOk

`func (o *PortSubGroup) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool)`

GetFcPortsOk returns a tuple with the FcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPorts

`func (o *PortSubGroup) SetFcPorts(v []FcPhysicalPortRelationship)`

SetFcPorts sets FcPorts field to given value.

### HasFcPorts

`func (o *PortSubGroup) HasFcPorts() bool`

HasFcPorts returns a boolean if a field has been set.

### SetFcPortsNil

`func (o *PortSubGroup) SetFcPortsNil(b bool)`

 SetFcPortsNil sets the value for FcPorts to be an explicit nil

### UnsetFcPorts
`func (o *PortSubGroup) UnsetFcPorts()`

UnsetFcPorts ensures that no value is present for FcPorts, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *PortSubGroup) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PortSubGroup) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PortSubGroup) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PortSubGroup) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroup

`func (o *PortSubGroup) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *PortSubGroup) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *PortSubGroup) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *PortSubGroup) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PortSubGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PortSubGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PortSubGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PortSubGroup) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


