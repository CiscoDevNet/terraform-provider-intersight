# EquipmentExpanderModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ExpanderModule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ExpanderModule"]
**FirmwareVersion** | Pointer to **string** | Firmware Version of the Chassis expander module. | [optional] [readonly] 
**InventoryReady** | Pointer to **bool** | The inventory ready field indicates whether the chassis expander module management controller has completed inventory of the installed devices. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | Module identifier for the expander module. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational state of expander module. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number identifier for the expander module. | [optional] [readonly] 
**Side** | Pointer to **string** | Location of the expander module within a chassis. The value can be top or bottom. * &#x60;unknown&#x60; - Physical location of the module is unknown. * &#x60;top&#x60; - Physical location of the module is on the top part of the chassis. * &#x60;bottom&#x60; - Physical location of the module is on the bottom part of the chassis. * &#x60;left&#x60; - Physical location of the module is on the left side of the chassis. * &#x60;right&#x60; - Physical location of the module is on the right side of the chassis. | [optional] [readonly] [default to "unknown"]
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**FanModules** | Pointer to [**[]EquipmentFanModuleRelationship**](EquipmentFanModuleRelationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**PciSwitches** | Pointer to [**[]PciSwitchRelationship**](PciSwitchRelationship.md) | An array of relationships to pciSwitch resources. | [optional] [readonly] 
**PciZones** | Pointer to [**[]PciZoneRelationship**](PciZoneRelationship.md) | An array of relationships to pciZone resources. | [optional] [readonly] 
**PhysicalDeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SharedAdapterUnits** | Pointer to [**[]EquipmentSharedAdapterUnitRelationship**](EquipmentSharedAdapterUnitRelationship.md) | An array of relationships to equipmentSharedAdapterUnit resources. | [optional] [readonly] 

## Methods

### NewEquipmentExpanderModule

`func NewEquipmentExpanderModule(classId string, objectType string, ) *EquipmentExpanderModule`

NewEquipmentExpanderModule instantiates a new EquipmentExpanderModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentExpanderModuleWithDefaults

`func NewEquipmentExpanderModuleWithDefaults() *EquipmentExpanderModule`

NewEquipmentExpanderModuleWithDefaults instantiates a new EquipmentExpanderModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentExpanderModule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentExpanderModule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentExpanderModule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentExpanderModule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentExpanderModule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentExpanderModule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFirmwareVersion

`func (o *EquipmentExpanderModule) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *EquipmentExpanderModule) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *EquipmentExpanderModule) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *EquipmentExpanderModule) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetInventoryReady

`func (o *EquipmentExpanderModule) GetInventoryReady() bool`

GetInventoryReady returns the InventoryReady field if non-nil, zero value otherwise.

### GetInventoryReadyOk

`func (o *EquipmentExpanderModule) GetInventoryReadyOk() (*bool, bool)`

GetInventoryReadyOk returns a tuple with the InventoryReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReady

`func (o *EquipmentExpanderModule) SetInventoryReady(v bool)`

SetInventoryReady sets InventoryReady field to given value.

### HasInventoryReady

`func (o *EquipmentExpanderModule) HasInventoryReady() bool`

HasInventoryReady returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentExpanderModule) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentExpanderModule) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentExpanderModule) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentExpanderModule) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentExpanderModule) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentExpanderModule) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentExpanderModule) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentExpanderModule) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentExpanderModule) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentExpanderModule) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentExpanderModule) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentExpanderModule) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentExpanderModule) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentExpanderModule) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentExpanderModule) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentExpanderModule) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentExpanderModule) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentExpanderModule) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetSide

`func (o *EquipmentExpanderModule) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *EquipmentExpanderModule) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *EquipmentExpanderModule) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *EquipmentExpanderModule) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentExpanderModule) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentExpanderModule) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentExpanderModule) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentExpanderModule) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *EquipmentExpanderModule) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *EquipmentExpanderModule) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetFanModules

`func (o *EquipmentExpanderModule) GetFanModules() []EquipmentFanModuleRelationship`

GetFanModules returns the FanModules field if non-nil, zero value otherwise.

### GetFanModulesOk

`func (o *EquipmentExpanderModule) GetFanModulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanModulesOk returns a tuple with the FanModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanModules

`func (o *EquipmentExpanderModule) SetFanModules(v []EquipmentFanModuleRelationship)`

SetFanModules sets FanModules field to given value.

### HasFanModules

`func (o *EquipmentExpanderModule) HasFanModules() bool`

HasFanModules returns a boolean if a field has been set.

### SetFanModulesNil

`func (o *EquipmentExpanderModule) SetFanModulesNil(b bool)`

 SetFanModulesNil sets the value for FanModules to be an explicit nil

### UnsetFanModules
`func (o *EquipmentExpanderModule) UnsetFanModules()`

UnsetFanModules ensures that no value is present for FanModules, not even an explicit nil
### GetPciSwitches

`func (o *EquipmentExpanderModule) GetPciSwitches() []PciSwitchRelationship`

GetPciSwitches returns the PciSwitches field if non-nil, zero value otherwise.

### GetPciSwitchesOk

`func (o *EquipmentExpanderModule) GetPciSwitchesOk() (*[]PciSwitchRelationship, bool)`

GetPciSwitchesOk returns a tuple with the PciSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitches

`func (o *EquipmentExpanderModule) SetPciSwitches(v []PciSwitchRelationship)`

SetPciSwitches sets PciSwitches field to given value.

### HasPciSwitches

`func (o *EquipmentExpanderModule) HasPciSwitches() bool`

HasPciSwitches returns a boolean if a field has been set.

### SetPciSwitchesNil

`func (o *EquipmentExpanderModule) SetPciSwitchesNil(b bool)`

 SetPciSwitchesNil sets the value for PciSwitches to be an explicit nil

### UnsetPciSwitches
`func (o *EquipmentExpanderModule) UnsetPciSwitches()`

UnsetPciSwitches ensures that no value is present for PciSwitches, not even an explicit nil
### GetPciZones

`func (o *EquipmentExpanderModule) GetPciZones() []PciZoneRelationship`

GetPciZones returns the PciZones field if non-nil, zero value otherwise.

### GetPciZonesOk

`func (o *EquipmentExpanderModule) GetPciZonesOk() (*[]PciZoneRelationship, bool)`

GetPciZonesOk returns a tuple with the PciZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciZones

`func (o *EquipmentExpanderModule) SetPciZones(v []PciZoneRelationship)`

SetPciZones sets PciZones field to given value.

### HasPciZones

`func (o *EquipmentExpanderModule) HasPciZones() bool`

HasPciZones returns a boolean if a field has been set.

### SetPciZonesNil

`func (o *EquipmentExpanderModule) SetPciZonesNil(b bool)`

 SetPciZonesNil sets the value for PciZones to be an explicit nil

### UnsetPciZones
`func (o *EquipmentExpanderModule) UnsetPciZones()`

UnsetPciZones ensures that no value is present for PciZones, not even an explicit nil
### GetPhysicalDeviceRegistration

`func (o *EquipmentExpanderModule) GetPhysicalDeviceRegistration() AssetDeviceRegistrationRelationship`

GetPhysicalDeviceRegistration returns the PhysicalDeviceRegistration field if non-nil, zero value otherwise.

### GetPhysicalDeviceRegistrationOk

`func (o *EquipmentExpanderModule) GetPhysicalDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetPhysicalDeviceRegistrationOk returns a tuple with the PhysicalDeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDeviceRegistration

`func (o *EquipmentExpanderModule) SetPhysicalDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetPhysicalDeviceRegistration sets PhysicalDeviceRegistration field to given value.

### HasPhysicalDeviceRegistration

`func (o *EquipmentExpanderModule) HasPhysicalDeviceRegistration() bool`

HasPhysicalDeviceRegistration returns a boolean if a field has been set.

### SetPhysicalDeviceRegistrationNil

`func (o *EquipmentExpanderModule) SetPhysicalDeviceRegistrationNil(b bool)`

 SetPhysicalDeviceRegistrationNil sets the value for PhysicalDeviceRegistration to be an explicit nil

### UnsetPhysicalDeviceRegistration
`func (o *EquipmentExpanderModule) UnsetPhysicalDeviceRegistration()`

UnsetPhysicalDeviceRegistration ensures that no value is present for PhysicalDeviceRegistration, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentExpanderModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentExpanderModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentExpanderModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentExpanderModule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentExpanderModule) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentExpanderModule) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSharedAdapterUnits

`func (o *EquipmentExpanderModule) GetSharedAdapterUnits() []EquipmentSharedAdapterUnitRelationship`

GetSharedAdapterUnits returns the SharedAdapterUnits field if non-nil, zero value otherwise.

### GetSharedAdapterUnitsOk

`func (o *EquipmentExpanderModule) GetSharedAdapterUnitsOk() (*[]EquipmentSharedAdapterUnitRelationship, bool)`

GetSharedAdapterUnitsOk returns a tuple with the SharedAdapterUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAdapterUnits

`func (o *EquipmentExpanderModule) SetSharedAdapterUnits(v []EquipmentSharedAdapterUnitRelationship)`

SetSharedAdapterUnits sets SharedAdapterUnits field to given value.

### HasSharedAdapterUnits

`func (o *EquipmentExpanderModule) HasSharedAdapterUnits() bool`

HasSharedAdapterUnits returns a boolean if a field has been set.

### SetSharedAdapterUnitsNil

`func (o *EquipmentExpanderModule) SetSharedAdapterUnitsNil(b bool)`

 SetSharedAdapterUnitsNil sets the value for SharedAdapterUnits to be an explicit nil

### UnsetSharedAdapterUnits
`func (o *EquipmentExpanderModule) UnsetSharedAdapterUnits()`

UnsetSharedAdapterUnits ensures that no value is present for SharedAdapterUnits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


