# PciEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pci.Endpoint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pci.Endpoint"]
**DeviceEnclosureId** | Pointer to **int64** | The identifier of the enclosure device of the actual physical device to which the logical PCIe endpoint is pointing to. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The identifier of the actual physical device to which the logical PCIe endpoint is pointing to. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | The type of the actual physical device to which the logical PCIe endpoint is pointing to. * &#x60;Unknown&#x60; - The device type of the physical device is unknown. * &#x60;NetworkAdapter&#x60; - The device type of the physical device is a NIC adapter. * &#x60;CPU&#x60; - The device type of the physical device is CPU. * &#x60;GPU&#x60; - The device type of the physical device is GPU. | [optional] [readonly] [default to "Unknown"]
**EndpointId** | Pointer to **string** | The identifier of the PCIe endpoint within a X-Fabric module PCIe switch. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the PCIe endpoint, as reported by the XFM platform software (BMC). | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational state of the PCIe endpoint. | [optional] [readonly] 
**Uri** | Pointer to **string** | The unique identifier of the PCIe endpoint as reported by the chassis expander management controller. | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PciSwitch** | Pointer to [**NullablePciSwitchRelationship**](PciSwitchRelationship.md) |  | [optional] 
**PciSwitchPort** | Pointer to [**NullablePciPortRelationship**](PciPortRelationship.md) |  | [optional] 
**PciZone** | Pointer to [**NullablePciZoneRelationship**](PciZoneRelationship.md) |  | [optional] 
**ProcessorUnit** | Pointer to [**NullableProcessorUnitRelationship**](ProcessorUnitRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SharedAdapterUnit** | Pointer to [**NullableEquipmentSharedAdapterUnitRelationship**](EquipmentSharedAdapterUnitRelationship.md) |  | [optional] 
**SharedGraphicsCard** | Pointer to [**NullableEquipmentSharedGraphicsCardRelationship**](EquipmentSharedGraphicsCardRelationship.md) |  | [optional] 

## Methods

### NewPciEndpoint

`func NewPciEndpoint(classId string, objectType string, ) *PciEndpoint`

NewPciEndpoint instantiates a new PciEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciEndpointWithDefaults

`func NewPciEndpointWithDefaults() *PciEndpoint`

NewPciEndpointWithDefaults instantiates a new PciEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PciEndpoint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PciEndpoint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PciEndpoint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PciEndpoint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PciEndpoint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PciEndpoint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceEnclosureId

`func (o *PciEndpoint) GetDeviceEnclosureId() int64`

GetDeviceEnclosureId returns the DeviceEnclosureId field if non-nil, zero value otherwise.

### GetDeviceEnclosureIdOk

`func (o *PciEndpoint) GetDeviceEnclosureIdOk() (*int64, bool)`

GetDeviceEnclosureIdOk returns a tuple with the DeviceEnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnclosureId

`func (o *PciEndpoint) SetDeviceEnclosureId(v int64)`

SetDeviceEnclosureId sets DeviceEnclosureId field to given value.

### HasDeviceEnclosureId

`func (o *PciEndpoint) HasDeviceEnclosureId() bool`

HasDeviceEnclosureId returns a boolean if a field has been set.

### GetDeviceId

`func (o *PciEndpoint) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PciEndpoint) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PciEndpoint) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *PciEndpoint) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *PciEndpoint) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PciEndpoint) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PciEndpoint) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PciEndpoint) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetEndpointId

`func (o *PciEndpoint) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *PciEndpoint) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *PciEndpoint) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *PciEndpoint) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetName

`func (o *PciEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PciEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PciEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PciEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *PciEndpoint) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *PciEndpoint) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *PciEndpoint) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *PciEndpoint) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *PciEndpoint) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *PciEndpoint) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *PciEndpoint) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PciEndpoint) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PciEndpoint) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PciEndpoint) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetUri

`func (o *PciEndpoint) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PciEndpoint) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PciEndpoint) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PciEndpoint) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *PciEndpoint) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciEndpoint) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciEndpoint) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciEndpoint) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *PciEndpoint) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *PciEndpoint) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetPciSwitch

`func (o *PciEndpoint) GetPciSwitch() PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *PciEndpoint) GetPciSwitchOk() (*PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *PciEndpoint) SetPciSwitch(v PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *PciEndpoint) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### SetPciSwitchNil

`func (o *PciEndpoint) SetPciSwitchNil(b bool)`

 SetPciSwitchNil sets the value for PciSwitch to be an explicit nil

### UnsetPciSwitch
`func (o *PciEndpoint) UnsetPciSwitch()`

UnsetPciSwitch ensures that no value is present for PciSwitch, not even an explicit nil
### GetPciSwitchPort

`func (o *PciEndpoint) GetPciSwitchPort() PciPortRelationship`

GetPciSwitchPort returns the PciSwitchPort field if non-nil, zero value otherwise.

### GetPciSwitchPortOk

`func (o *PciEndpoint) GetPciSwitchPortOk() (*PciPortRelationship, bool)`

GetPciSwitchPortOk returns a tuple with the PciSwitchPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitchPort

`func (o *PciEndpoint) SetPciSwitchPort(v PciPortRelationship)`

SetPciSwitchPort sets PciSwitchPort field to given value.

### HasPciSwitchPort

`func (o *PciEndpoint) HasPciSwitchPort() bool`

HasPciSwitchPort returns a boolean if a field has been set.

### SetPciSwitchPortNil

`func (o *PciEndpoint) SetPciSwitchPortNil(b bool)`

 SetPciSwitchPortNil sets the value for PciSwitchPort to be an explicit nil

### UnsetPciSwitchPort
`func (o *PciEndpoint) UnsetPciSwitchPort()`

UnsetPciSwitchPort ensures that no value is present for PciSwitchPort, not even an explicit nil
### GetPciZone

`func (o *PciEndpoint) GetPciZone() PciZoneRelationship`

GetPciZone returns the PciZone field if non-nil, zero value otherwise.

### GetPciZoneOk

`func (o *PciEndpoint) GetPciZoneOk() (*PciZoneRelationship, bool)`

GetPciZoneOk returns a tuple with the PciZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciZone

`func (o *PciEndpoint) SetPciZone(v PciZoneRelationship)`

SetPciZone sets PciZone field to given value.

### HasPciZone

`func (o *PciEndpoint) HasPciZone() bool`

HasPciZone returns a boolean if a field has been set.

### SetPciZoneNil

`func (o *PciEndpoint) SetPciZoneNil(b bool)`

 SetPciZoneNil sets the value for PciZone to be an explicit nil

### UnsetPciZone
`func (o *PciEndpoint) UnsetPciZone()`

UnsetPciZone ensures that no value is present for PciZone, not even an explicit nil
### GetProcessorUnit

`func (o *PciEndpoint) GetProcessorUnit() ProcessorUnitRelationship`

GetProcessorUnit returns the ProcessorUnit field if non-nil, zero value otherwise.

### GetProcessorUnitOk

`func (o *PciEndpoint) GetProcessorUnitOk() (*ProcessorUnitRelationship, bool)`

GetProcessorUnitOk returns a tuple with the ProcessorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorUnit

`func (o *PciEndpoint) SetProcessorUnit(v ProcessorUnitRelationship)`

SetProcessorUnit sets ProcessorUnit field to given value.

### HasProcessorUnit

`func (o *PciEndpoint) HasProcessorUnit() bool`

HasProcessorUnit returns a boolean if a field has been set.

### SetProcessorUnitNil

`func (o *PciEndpoint) SetProcessorUnitNil(b bool)`

 SetProcessorUnitNil sets the value for ProcessorUnit to be an explicit nil

### UnsetProcessorUnit
`func (o *PciEndpoint) UnsetProcessorUnit()`

UnsetProcessorUnit ensures that no value is present for ProcessorUnit, not even an explicit nil
### GetRegisteredDevice

`func (o *PciEndpoint) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciEndpoint) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciEndpoint) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciEndpoint) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *PciEndpoint) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *PciEndpoint) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSharedAdapterUnit

`func (o *PciEndpoint) GetSharedAdapterUnit() EquipmentSharedAdapterUnitRelationship`

GetSharedAdapterUnit returns the SharedAdapterUnit field if non-nil, zero value otherwise.

### GetSharedAdapterUnitOk

`func (o *PciEndpoint) GetSharedAdapterUnitOk() (*EquipmentSharedAdapterUnitRelationship, bool)`

GetSharedAdapterUnitOk returns a tuple with the SharedAdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAdapterUnit

`func (o *PciEndpoint) SetSharedAdapterUnit(v EquipmentSharedAdapterUnitRelationship)`

SetSharedAdapterUnit sets SharedAdapterUnit field to given value.

### HasSharedAdapterUnit

`func (o *PciEndpoint) HasSharedAdapterUnit() bool`

HasSharedAdapterUnit returns a boolean if a field has been set.

### SetSharedAdapterUnitNil

`func (o *PciEndpoint) SetSharedAdapterUnitNil(b bool)`

 SetSharedAdapterUnitNil sets the value for SharedAdapterUnit to be an explicit nil

### UnsetSharedAdapterUnit
`func (o *PciEndpoint) UnsetSharedAdapterUnit()`

UnsetSharedAdapterUnit ensures that no value is present for SharedAdapterUnit, not even an explicit nil
### GetSharedGraphicsCard

`func (o *PciEndpoint) GetSharedGraphicsCard() EquipmentSharedGraphicsCardRelationship`

GetSharedGraphicsCard returns the SharedGraphicsCard field if non-nil, zero value otherwise.

### GetSharedGraphicsCardOk

`func (o *PciEndpoint) GetSharedGraphicsCardOk() (*EquipmentSharedGraphicsCardRelationship, bool)`

GetSharedGraphicsCardOk returns a tuple with the SharedGraphicsCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedGraphicsCard

`func (o *PciEndpoint) SetSharedGraphicsCard(v EquipmentSharedGraphicsCardRelationship)`

SetSharedGraphicsCard sets SharedGraphicsCard field to given value.

### HasSharedGraphicsCard

`func (o *PciEndpoint) HasSharedGraphicsCard() bool`

HasSharedGraphicsCard returns a boolean if a field has been set.

### SetSharedGraphicsCardNil

`func (o *PciEndpoint) SetSharedGraphicsCardNil(b bool)`

 SetSharedGraphicsCardNil sets the value for SharedGraphicsCard to be an explicit nil

### UnsetSharedGraphicsCard
`func (o *PciEndpoint) UnsetSharedGraphicsCard()`

UnsetSharedGraphicsCard ensures that no value is present for SharedGraphicsCard, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


