# PortMacBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "port.MacBinding"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "port.MacBinding"]
**AggregatePortId** | Pointer to **int64** | Aggregate Port ID of the local Switch Interface. | [optional] 
**ChassisId** | Pointer to **int64** | Chassis/FEX device idetifier that is local to a cluster. | [optional] 
**ChassisModel** | Pointer to **string** | Chassis/Rack Model that is associated with the Switch/FEX interface. | [optional] 
**ChassisSerial** | Pointer to **string** | Chassis/Rack Serial that is associated with the Switch/FEX interface. | [optional] 
**ChassisVendor** | Pointer to **string** | Chassis/Rack Vendor that is associated with the Switch/FEX interface. | [optional] 
**DeviceMac** | Pointer to **string** | Device ID value that is advertised and available as a part of LLDP TLV. | [optional] 
**ModuleMode** | Pointer to **int64** | IOM/SIOC/Adapter Mode that is associated with the Switch/FEX interface. | [optional] 
**ModuleModel** | Pointer to **string** | IOM/SIOC/Adapter Model that is associated with the Switch/FEX interface. | [optional] 
**ModulePortId** | Pointer to **int64** | Uplink port identifier of the VIC that is associated with the Switch/FEX interface. | [optional] 
**ModuleSerial** | Pointer to **string** | IOM/SIOC/Adapter Serial that is associated with the Switch/FEX interface. | [optional] 
**ModuleSide** | Pointer to **int64** | IOM/SIOC/Adapter Side that is associated with the Switch/FEX interface. | [optional] 
**ModuleSlot** | Pointer to **int64** | IOM/SIOC/Adapter Slot that is associated with the Switch/FEX interface. | [optional] 
**ModuleVendor** | Pointer to **string** | IOM/SIOC/Adapter Vendor that is associated with the Switch/FEX interface. | [optional] 
**PortId** | Pointer to **int64** | Port ID of the local Switch Interface. | [optional] 
**PortMac** | Pointer to **string** | Port ID value that is advertised and available as a part of LLDP TLV. | [optional] 
**SlotId** | Pointer to **int64** | Slot ID of the local Switch slot Interface. | [optional] 
**SwitchId** | Pointer to **int64** | Switch Identifier that is local to a cluster. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewPortMacBinding

`func NewPortMacBinding(classId string, objectType string, ) *PortMacBinding`

NewPortMacBinding instantiates a new PortMacBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMacBindingWithDefaults

`func NewPortMacBindingWithDefaults() *PortMacBinding`

NewPortMacBindingWithDefaults instantiates a new PortMacBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PortMacBinding) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PortMacBinding) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PortMacBinding) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PortMacBinding) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PortMacBinding) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PortMacBinding) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregatePortId

`func (o *PortMacBinding) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *PortMacBinding) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *PortMacBinding) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *PortMacBinding) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetChassisId

`func (o *PortMacBinding) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *PortMacBinding) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *PortMacBinding) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *PortMacBinding) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetChassisModel

`func (o *PortMacBinding) GetChassisModel() string`

GetChassisModel returns the ChassisModel field if non-nil, zero value otherwise.

### GetChassisModelOk

`func (o *PortMacBinding) GetChassisModelOk() (*string, bool)`

GetChassisModelOk returns a tuple with the ChassisModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisModel

`func (o *PortMacBinding) SetChassisModel(v string)`

SetChassisModel sets ChassisModel field to given value.

### HasChassisModel

`func (o *PortMacBinding) HasChassisModel() bool`

HasChassisModel returns a boolean if a field has been set.

### GetChassisSerial

`func (o *PortMacBinding) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *PortMacBinding) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *PortMacBinding) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *PortMacBinding) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### GetChassisVendor

`func (o *PortMacBinding) GetChassisVendor() string`

GetChassisVendor returns the ChassisVendor field if non-nil, zero value otherwise.

### GetChassisVendorOk

`func (o *PortMacBinding) GetChassisVendorOk() (*string, bool)`

GetChassisVendorOk returns a tuple with the ChassisVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisVendor

`func (o *PortMacBinding) SetChassisVendor(v string)`

SetChassisVendor sets ChassisVendor field to given value.

### HasChassisVendor

`func (o *PortMacBinding) HasChassisVendor() bool`

HasChassisVendor returns a boolean if a field has been set.

### GetDeviceMac

`func (o *PortMacBinding) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *PortMacBinding) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *PortMacBinding) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *PortMacBinding) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetModuleMode

`func (o *PortMacBinding) GetModuleMode() int64`

GetModuleMode returns the ModuleMode field if non-nil, zero value otherwise.

### GetModuleModeOk

`func (o *PortMacBinding) GetModuleModeOk() (*int64, bool)`

GetModuleModeOk returns a tuple with the ModuleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleMode

`func (o *PortMacBinding) SetModuleMode(v int64)`

SetModuleMode sets ModuleMode field to given value.

### HasModuleMode

`func (o *PortMacBinding) HasModuleMode() bool`

HasModuleMode returns a boolean if a field has been set.

### GetModuleModel

`func (o *PortMacBinding) GetModuleModel() string`

GetModuleModel returns the ModuleModel field if non-nil, zero value otherwise.

### GetModuleModelOk

`func (o *PortMacBinding) GetModuleModelOk() (*string, bool)`

GetModuleModelOk returns a tuple with the ModuleModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleModel

`func (o *PortMacBinding) SetModuleModel(v string)`

SetModuleModel sets ModuleModel field to given value.

### HasModuleModel

`func (o *PortMacBinding) HasModuleModel() bool`

HasModuleModel returns a boolean if a field has been set.

### GetModulePortId

`func (o *PortMacBinding) GetModulePortId() int64`

GetModulePortId returns the ModulePortId field if non-nil, zero value otherwise.

### GetModulePortIdOk

`func (o *PortMacBinding) GetModulePortIdOk() (*int64, bool)`

GetModulePortIdOk returns a tuple with the ModulePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulePortId

`func (o *PortMacBinding) SetModulePortId(v int64)`

SetModulePortId sets ModulePortId field to given value.

### HasModulePortId

`func (o *PortMacBinding) HasModulePortId() bool`

HasModulePortId returns a boolean if a field has been set.

### GetModuleSerial

`func (o *PortMacBinding) GetModuleSerial() string`

GetModuleSerial returns the ModuleSerial field if non-nil, zero value otherwise.

### GetModuleSerialOk

`func (o *PortMacBinding) GetModuleSerialOk() (*string, bool)`

GetModuleSerialOk returns a tuple with the ModuleSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleSerial

`func (o *PortMacBinding) SetModuleSerial(v string)`

SetModuleSerial sets ModuleSerial field to given value.

### HasModuleSerial

`func (o *PortMacBinding) HasModuleSerial() bool`

HasModuleSerial returns a boolean if a field has been set.

### GetModuleSide

`func (o *PortMacBinding) GetModuleSide() int64`

GetModuleSide returns the ModuleSide field if non-nil, zero value otherwise.

### GetModuleSideOk

`func (o *PortMacBinding) GetModuleSideOk() (*int64, bool)`

GetModuleSideOk returns a tuple with the ModuleSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleSide

`func (o *PortMacBinding) SetModuleSide(v int64)`

SetModuleSide sets ModuleSide field to given value.

### HasModuleSide

`func (o *PortMacBinding) HasModuleSide() bool`

HasModuleSide returns a boolean if a field has been set.

### GetModuleSlot

`func (o *PortMacBinding) GetModuleSlot() int64`

GetModuleSlot returns the ModuleSlot field if non-nil, zero value otherwise.

### GetModuleSlotOk

`func (o *PortMacBinding) GetModuleSlotOk() (*int64, bool)`

GetModuleSlotOk returns a tuple with the ModuleSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleSlot

`func (o *PortMacBinding) SetModuleSlot(v int64)`

SetModuleSlot sets ModuleSlot field to given value.

### HasModuleSlot

`func (o *PortMacBinding) HasModuleSlot() bool`

HasModuleSlot returns a boolean if a field has been set.

### GetModuleVendor

`func (o *PortMacBinding) GetModuleVendor() string`

GetModuleVendor returns the ModuleVendor field if non-nil, zero value otherwise.

### GetModuleVendorOk

`func (o *PortMacBinding) GetModuleVendorOk() (*string, bool)`

GetModuleVendorOk returns a tuple with the ModuleVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleVendor

`func (o *PortMacBinding) SetModuleVendor(v string)`

SetModuleVendor sets ModuleVendor field to given value.

### HasModuleVendor

`func (o *PortMacBinding) HasModuleVendor() bool`

HasModuleVendor returns a boolean if a field has been set.

### GetPortId

`func (o *PortMacBinding) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PortMacBinding) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PortMacBinding) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PortMacBinding) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortMac

`func (o *PortMacBinding) GetPortMac() string`

GetPortMac returns the PortMac field if non-nil, zero value otherwise.

### GetPortMacOk

`func (o *PortMacBinding) GetPortMacOk() (*string, bool)`

GetPortMacOk returns a tuple with the PortMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMac

`func (o *PortMacBinding) SetPortMac(v string)`

SetPortMac sets PortMac field to given value.

### HasPortMac

`func (o *PortMacBinding) HasPortMac() bool`

HasPortMac returns a boolean if a field has been set.

### GetSlotId

`func (o *PortMacBinding) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *PortMacBinding) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *PortMacBinding) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *PortMacBinding) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *PortMacBinding) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *PortMacBinding) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *PortMacBinding) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *PortMacBinding) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *PortMacBinding) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *PortMacBinding) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *PortMacBinding) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *PortMacBinding) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PortMacBinding) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PortMacBinding) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PortMacBinding) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PortMacBinding) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


