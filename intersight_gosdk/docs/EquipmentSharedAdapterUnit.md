# EquipmentSharedAdapterUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.SharedAdapterUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.SharedAdapterUnit"]
**AdapterId** | Pointer to **string** | Unique Identifier of a shared adapter unit within a expander module. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | Firmware version of the shared adapter unit. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational state of the shared adapter unit. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number of a shared adapter unit. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | PCIe slot of the adapter in the chassis expander module. | [optional] [readonly] 
**AdapterUnits** | Pointer to [**[]AdapterUnitRelationship**](AdapterUnitRelationship.md) | An array of relationships to adapterUnit resources. | [optional] [readonly] 
**EquipmentExpanderModule** | Pointer to [**NullableEquipmentExpanderModuleRelationship**](EquipmentExpanderModuleRelationship.md) |  | [optional] 
**PciEndpoint** | Pointer to [**NullablePciEndpointRelationship**](PciEndpointRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSharedAdapterUnit

`func NewEquipmentSharedAdapterUnit(classId string, objectType string, ) *EquipmentSharedAdapterUnit`

NewEquipmentSharedAdapterUnit instantiates a new EquipmentSharedAdapterUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSharedAdapterUnitWithDefaults

`func NewEquipmentSharedAdapterUnitWithDefaults() *EquipmentSharedAdapterUnit`

NewEquipmentSharedAdapterUnitWithDefaults instantiates a new EquipmentSharedAdapterUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSharedAdapterUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSharedAdapterUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSharedAdapterUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSharedAdapterUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSharedAdapterUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSharedAdapterUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterId

`func (o *EquipmentSharedAdapterUnit) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *EquipmentSharedAdapterUnit) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *EquipmentSharedAdapterUnit) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *EquipmentSharedAdapterUnit) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *EquipmentSharedAdapterUnit) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *EquipmentSharedAdapterUnit) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *EquipmentSharedAdapterUnit) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *EquipmentSharedAdapterUnit) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentSharedAdapterUnit) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentSharedAdapterUnit) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentSharedAdapterUnit) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentSharedAdapterUnit) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentSharedAdapterUnit) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentSharedAdapterUnit) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentSharedAdapterUnit) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentSharedAdapterUnit) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentSharedAdapterUnit) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentSharedAdapterUnit) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentSharedAdapterUnit) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentSharedAdapterUnit) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentSharedAdapterUnit) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentSharedAdapterUnit) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPciSlot

`func (o *EquipmentSharedAdapterUnit) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *EquipmentSharedAdapterUnit) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *EquipmentSharedAdapterUnit) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *EquipmentSharedAdapterUnit) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetAdapterUnits

`func (o *EquipmentSharedAdapterUnit) GetAdapterUnits() []AdapterUnitRelationship`

GetAdapterUnits returns the AdapterUnits field if non-nil, zero value otherwise.

### GetAdapterUnitsOk

`func (o *EquipmentSharedAdapterUnit) GetAdapterUnitsOk() (*[]AdapterUnitRelationship, bool)`

GetAdapterUnitsOk returns a tuple with the AdapterUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnits

`func (o *EquipmentSharedAdapterUnit) SetAdapterUnits(v []AdapterUnitRelationship)`

SetAdapterUnits sets AdapterUnits field to given value.

### HasAdapterUnits

`func (o *EquipmentSharedAdapterUnit) HasAdapterUnits() bool`

HasAdapterUnits returns a boolean if a field has been set.

### SetAdapterUnitsNil

`func (o *EquipmentSharedAdapterUnit) SetAdapterUnitsNil(b bool)`

 SetAdapterUnitsNil sets the value for AdapterUnits to be an explicit nil

### UnsetAdapterUnits
`func (o *EquipmentSharedAdapterUnit) UnsetAdapterUnits()`

UnsetAdapterUnits ensures that no value is present for AdapterUnits, not even an explicit nil
### GetEquipmentExpanderModule

`func (o *EquipmentSharedAdapterUnit) GetEquipmentExpanderModule() EquipmentExpanderModuleRelationship`

GetEquipmentExpanderModule returns the EquipmentExpanderModule field if non-nil, zero value otherwise.

### GetEquipmentExpanderModuleOk

`func (o *EquipmentSharedAdapterUnit) GetEquipmentExpanderModuleOk() (*EquipmentExpanderModuleRelationship, bool)`

GetEquipmentExpanderModuleOk returns a tuple with the EquipmentExpanderModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentExpanderModule

`func (o *EquipmentSharedAdapterUnit) SetEquipmentExpanderModule(v EquipmentExpanderModuleRelationship)`

SetEquipmentExpanderModule sets EquipmentExpanderModule field to given value.

### HasEquipmentExpanderModule

`func (o *EquipmentSharedAdapterUnit) HasEquipmentExpanderModule() bool`

HasEquipmentExpanderModule returns a boolean if a field has been set.

### SetEquipmentExpanderModuleNil

`func (o *EquipmentSharedAdapterUnit) SetEquipmentExpanderModuleNil(b bool)`

 SetEquipmentExpanderModuleNil sets the value for EquipmentExpanderModule to be an explicit nil

### UnsetEquipmentExpanderModule
`func (o *EquipmentSharedAdapterUnit) UnsetEquipmentExpanderModule()`

UnsetEquipmentExpanderModule ensures that no value is present for EquipmentExpanderModule, not even an explicit nil
### GetPciEndpoint

`func (o *EquipmentSharedAdapterUnit) GetPciEndpoint() PciEndpointRelationship`

GetPciEndpoint returns the PciEndpoint field if non-nil, zero value otherwise.

### GetPciEndpointOk

`func (o *EquipmentSharedAdapterUnit) GetPciEndpointOk() (*PciEndpointRelationship, bool)`

GetPciEndpointOk returns a tuple with the PciEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciEndpoint

`func (o *EquipmentSharedAdapterUnit) SetPciEndpoint(v PciEndpointRelationship)`

SetPciEndpoint sets PciEndpoint field to given value.

### HasPciEndpoint

`func (o *EquipmentSharedAdapterUnit) HasPciEndpoint() bool`

HasPciEndpoint returns a boolean if a field has been set.

### SetPciEndpointNil

`func (o *EquipmentSharedAdapterUnit) SetPciEndpointNil(b bool)`

 SetPciEndpointNil sets the value for PciEndpoint to be an explicit nil

### UnsetPciEndpoint
`func (o *EquipmentSharedAdapterUnit) UnsetPciEndpoint()`

UnsetPciEndpoint ensures that no value is present for PciEndpoint, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSharedAdapterUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSharedAdapterUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSharedAdapterUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSharedAdapterUnit) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentSharedAdapterUnit) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentSharedAdapterUnit) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


