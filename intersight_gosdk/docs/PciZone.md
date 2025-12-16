# PciZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pci.Zone"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pci.Zone"]
**Name** | Pointer to **string** | The name of the PCIe endpoint zone, as reported by the platform software (BMC). | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational state of the pcie node. | [optional] [readonly] 
**ZoneId** | Pointer to **string** | The identifier of the PCIe endpoint zone where all PCIe devices are logically connected. | [optional] [readonly] 
**EquipmentExpanderModule** | Pointer to [**NullableEquipmentExpanderModuleRelationship**](EquipmentExpanderModuleRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PciEndpoints** | Pointer to [**[]PciEndpointRelationship**](PciEndpointRelationship.md) | An array of relationships to pciEndpoint resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPciZone

`func NewPciZone(classId string, objectType string, ) *PciZone`

NewPciZone instantiates a new PciZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciZoneWithDefaults

`func NewPciZoneWithDefaults() *PciZone`

NewPciZoneWithDefaults instantiates a new PciZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PciZone) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PciZone) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PciZone) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PciZone) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PciZone) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PciZone) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *PciZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PciZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PciZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PciZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *PciZone) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *PciZone) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *PciZone) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *PciZone) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *PciZone) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *PciZone) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *PciZone) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PciZone) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PciZone) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PciZone) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetZoneId

`func (o *PciZone) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *PciZone) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *PciZone) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *PciZone) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetEquipmentExpanderModule

`func (o *PciZone) GetEquipmentExpanderModule() EquipmentExpanderModuleRelationship`

GetEquipmentExpanderModule returns the EquipmentExpanderModule field if non-nil, zero value otherwise.

### GetEquipmentExpanderModuleOk

`func (o *PciZone) GetEquipmentExpanderModuleOk() (*EquipmentExpanderModuleRelationship, bool)`

GetEquipmentExpanderModuleOk returns a tuple with the EquipmentExpanderModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentExpanderModule

`func (o *PciZone) SetEquipmentExpanderModule(v EquipmentExpanderModuleRelationship)`

SetEquipmentExpanderModule sets EquipmentExpanderModule field to given value.

### HasEquipmentExpanderModule

`func (o *PciZone) HasEquipmentExpanderModule() bool`

HasEquipmentExpanderModule returns a boolean if a field has been set.

### SetEquipmentExpanderModuleNil

`func (o *PciZone) SetEquipmentExpanderModuleNil(b bool)`

 SetEquipmentExpanderModuleNil sets the value for EquipmentExpanderModule to be an explicit nil

### UnsetEquipmentExpanderModule
`func (o *PciZone) UnsetEquipmentExpanderModule()`

UnsetEquipmentExpanderModule ensures that no value is present for EquipmentExpanderModule, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *PciZone) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciZone) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciZone) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciZone) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *PciZone) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *PciZone) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetPciEndpoints

`func (o *PciZone) GetPciEndpoints() []PciEndpointRelationship`

GetPciEndpoints returns the PciEndpoints field if non-nil, zero value otherwise.

### GetPciEndpointsOk

`func (o *PciZone) GetPciEndpointsOk() (*[]PciEndpointRelationship, bool)`

GetPciEndpointsOk returns a tuple with the PciEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciEndpoints

`func (o *PciZone) SetPciEndpoints(v []PciEndpointRelationship)`

SetPciEndpoints sets PciEndpoints field to given value.

### HasPciEndpoints

`func (o *PciZone) HasPciEndpoints() bool`

HasPciEndpoints returns a boolean if a field has been set.

### SetPciEndpointsNil

`func (o *PciZone) SetPciEndpointsNil(b bool)`

 SetPciEndpointsNil sets the value for PciEndpoints to be an explicit nil

### UnsetPciEndpoints
`func (o *PciZone) UnsetPciEndpoints()`

UnsetPciEndpoints ensures that no value is present for PciEndpoints, not even an explicit nil
### GetRegisteredDevice

`func (o *PciZone) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciZone) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciZone) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciZone) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *PciZone) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *PciZone) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


