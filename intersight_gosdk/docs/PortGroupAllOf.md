# PortGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "port.Group"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "port.Group"]
**Transport** | Pointer to **string** | Type of port group. Values are Eth or Fc. | [optional] [readonly] 
**EquipmentSharedIoModule** | Pointer to [**EquipmentSharedIoModuleRelationship**](EquipmentSharedIoModuleRelationship.md) |  | [optional] 
**EquipmentSwitchCard** | Pointer to [**EquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) |  | [optional] 
**EthernetPorts** | Pointer to [**[]EtherPhysicalPortRelationship**](EtherPhysicalPortRelationship.md) | An array of relationships to etherPhysicalPort resources. | [optional] [readonly] 
**FcPorts** | Pointer to [**[]FcPhysicalPortRelationship**](FcPhysicalPortRelationship.md) | An array of relationships to fcPhysicalPort resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SubGroups** | Pointer to [**[]PortSubGroupRelationship**](PortSubGroupRelationship.md) | An array of relationships to portSubGroup resources. | [optional] [readonly] 

## Methods

### NewPortGroupAllOf

`func NewPortGroupAllOf(classId string, objectType string, ) *PortGroupAllOf`

NewPortGroupAllOf instantiates a new PortGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortGroupAllOfWithDefaults

`func NewPortGroupAllOfWithDefaults() *PortGroupAllOf`

NewPortGroupAllOfWithDefaults instantiates a new PortGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PortGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PortGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PortGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PortGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PortGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PortGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTransport

`func (o *PortGroupAllOf) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *PortGroupAllOf) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *PortGroupAllOf) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *PortGroupAllOf) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetEquipmentSharedIoModule

`func (o *PortGroupAllOf) GetEquipmentSharedIoModule() EquipmentSharedIoModuleRelationship`

GetEquipmentSharedIoModule returns the EquipmentSharedIoModule field if non-nil, zero value otherwise.

### GetEquipmentSharedIoModuleOk

`func (o *PortGroupAllOf) GetEquipmentSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool)`

GetEquipmentSharedIoModuleOk returns a tuple with the EquipmentSharedIoModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSharedIoModule

`func (o *PortGroupAllOf) SetEquipmentSharedIoModule(v EquipmentSharedIoModuleRelationship)`

SetEquipmentSharedIoModule sets EquipmentSharedIoModule field to given value.

### HasEquipmentSharedIoModule

`func (o *PortGroupAllOf) HasEquipmentSharedIoModule() bool`

HasEquipmentSharedIoModule returns a boolean if a field has been set.

### GetEquipmentSwitchCard

`func (o *PortGroupAllOf) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *PortGroupAllOf) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *PortGroupAllOf) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *PortGroupAllOf) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### GetEthernetPorts

`func (o *PortGroupAllOf) GetEthernetPorts() []EtherPhysicalPortRelationship`

GetEthernetPorts returns the EthernetPorts field if non-nil, zero value otherwise.

### GetEthernetPortsOk

`func (o *PortGroupAllOf) GetEthernetPortsOk() (*[]EtherPhysicalPortRelationship, bool)`

GetEthernetPortsOk returns a tuple with the EthernetPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPorts

`func (o *PortGroupAllOf) SetEthernetPorts(v []EtherPhysicalPortRelationship)`

SetEthernetPorts sets EthernetPorts field to given value.

### HasEthernetPorts

`func (o *PortGroupAllOf) HasEthernetPorts() bool`

HasEthernetPorts returns a boolean if a field has been set.

### SetEthernetPortsNil

`func (o *PortGroupAllOf) SetEthernetPortsNil(b bool)`

 SetEthernetPortsNil sets the value for EthernetPorts to be an explicit nil

### UnsetEthernetPorts
`func (o *PortGroupAllOf) UnsetEthernetPorts()`

UnsetEthernetPorts ensures that no value is present for EthernetPorts, not even an explicit nil
### GetFcPorts

`func (o *PortGroupAllOf) GetFcPorts() []FcPhysicalPortRelationship`

GetFcPorts returns the FcPorts field if non-nil, zero value otherwise.

### GetFcPortsOk

`func (o *PortGroupAllOf) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool)`

GetFcPortsOk returns a tuple with the FcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPorts

`func (o *PortGroupAllOf) SetFcPorts(v []FcPhysicalPortRelationship)`

SetFcPorts sets FcPorts field to given value.

### HasFcPorts

`func (o *PortGroupAllOf) HasFcPorts() bool`

HasFcPorts returns a boolean if a field has been set.

### SetFcPortsNil

`func (o *PortGroupAllOf) SetFcPortsNil(b bool)`

 SetFcPortsNil sets the value for FcPorts to be an explicit nil

### UnsetFcPorts
`func (o *PortGroupAllOf) UnsetFcPorts()`

UnsetFcPorts ensures that no value is present for FcPorts, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *PortGroupAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PortGroupAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PortGroupAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PortGroupAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PortGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PortGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PortGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PortGroupAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSubGroups

`func (o *PortGroupAllOf) GetSubGroups() []PortSubGroupRelationship`

GetSubGroups returns the SubGroups field if non-nil, zero value otherwise.

### GetSubGroupsOk

`func (o *PortGroupAllOf) GetSubGroupsOk() (*[]PortSubGroupRelationship, bool)`

GetSubGroupsOk returns a tuple with the SubGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGroups

`func (o *PortGroupAllOf) SetSubGroups(v []PortSubGroupRelationship)`

SetSubGroups sets SubGroups field to given value.

### HasSubGroups

`func (o *PortGroupAllOf) HasSubGroups() bool`

HasSubGroups returns a boolean if a field has been set.

### SetSubGroupsNil

`func (o *PortGroupAllOf) SetSubGroupsNil(b bool)`

 SetSubGroupsNil sets the value for SubGroups to be an explicit nil

### UnsetSubGroups
`func (o *PortGroupAllOf) UnsetSubGroups()`

UnsetSubGroups ensures that no value is present for SubGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


