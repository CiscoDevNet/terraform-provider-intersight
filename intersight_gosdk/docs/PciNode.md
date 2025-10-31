# PciNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pci.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pci.Node"]
**AlarmSummary** | Pointer to [**NullableComputeAlarmSummary**](ComputeAlarmSummary.md) |  | [optional] 
**ChassisId** | Pointer to **string** | The id of the chassis that the PCIe node is currently located in. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The version of the PCIe node CIMC firmware. | [optional] [readonly] 
**InventoryReady** | Pointer to **bool** | The inventory ready field indicates whether the PCIe node management controller has completed inventory of the installed PCIe devices. | [optional] [readonly] 
**Lifecycle** | Pointer to **string** | The lifecycle state of the PCIe node. This will map to the discovery lifecycle as represented in the Identity object. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;SecureEraseInProgress&#x60; - Secure Erase is in progress on given physical entity. * &#x60;ScrubInProgress&#x60; - Scrub is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;SlotMismatch&#x60; - The blade server is detected in a different chassis/slot than it was previously. * &#x60;Removed&#x60; - The blade server has been removed from its discovered slot, and not detected anywhere else. Blade inventory can be cleaned up by performing a software remove operation on the physically removed blade. * &#x60;Moved&#x60; - The blade server has been moved from its discovered location to a new location. Blade inventory can be updated by performing a rediscover operation on the moved blade. * &#x60;Replaced&#x60; - The blade server has been removed from its discovered location and another blade has been inserted in that location. Blade inventory can be cleaned up and updated by doing a software remove operation on the physically removed blade. * &#x60;MovedAndReplaced&#x60; - The blade server has been moved from its discovered location to a new location and another blade has been inserted into the old discovered location. Blade inventory can be updated by performing a rediscover operation on the moved blade. | [optional] [readonly] [default to "None"]
**Name** | Pointer to **string** | The name of the PCIe node, the value of this property is the name of the UCS Domain along with the chassis and node slot ids. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational state of the PCIe node. | [optional] [readonly] 
**PackageVersion** | Pointer to **string** | Bundle version which the CIMC firmware belongs to. | [optional] [readonly] 
**SlotId** | Pointer to **string** | The slot number in the chassis that the PCIe node is currently located in. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](GraphicsCardRelationship.md) | An array of relationships to graphicsCard resources. | [optional] [readonly] 
**Interconnects** | Pointer to [**[]EquipmentInterconnectRelationship**](EquipmentInterconnectRelationship.md) | An array of relationships to equipmentInterconnect resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**NullableEquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SharedGraphicsCards** | Pointer to [**[]EquipmentSharedGraphicsCardRelationship**](EquipmentSharedGraphicsCardRelationship.md) | An array of relationships to equipmentSharedGraphicsCard resources. | [optional] [readonly] 

## Methods

### NewPciNode

`func NewPciNode(classId string, objectType string, ) *PciNode`

NewPciNode instantiates a new PciNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciNodeWithDefaults

`func NewPciNodeWithDefaults() *PciNode`

NewPciNodeWithDefaults instantiates a new PciNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PciNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PciNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PciNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PciNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PciNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PciNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmSummary

`func (o *PciNode) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *PciNode) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *PciNode) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *PciNode) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *PciNode) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *PciNode) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetChassisId

`func (o *PciNode) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *PciNode) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *PciNode) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *PciNode) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *PciNode) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *PciNode) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *PciNode) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *PciNode) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetInventoryReady

`func (o *PciNode) GetInventoryReady() bool`

GetInventoryReady returns the InventoryReady field if non-nil, zero value otherwise.

### GetInventoryReadyOk

`func (o *PciNode) GetInventoryReadyOk() (*bool, bool)`

GetInventoryReadyOk returns a tuple with the InventoryReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReady

`func (o *PciNode) SetInventoryReady(v bool)`

SetInventoryReady sets InventoryReady field to given value.

### HasInventoryReady

`func (o *PciNode) HasInventoryReady() bool`

HasInventoryReady returns a boolean if a field has been set.

### GetLifecycle

`func (o *PciNode) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *PciNode) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *PciNode) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *PciNode) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetName

`func (o *PciNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PciNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PciNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PciNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *PciNode) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *PciNode) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *PciNode) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *PciNode) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *PciNode) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *PciNode) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *PciNode) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PciNode) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PciNode) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PciNode) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPackageVersion

`func (o *PciNode) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *PciNode) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *PciNode) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *PciNode) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetSlotId

`func (o *PciNode) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *PciNode) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *PciNode) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *PciNode) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetComputeBlade

`func (o *PciNode) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *PciNode) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *PciNode) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *PciNode) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### SetComputeBladeNil

`func (o *PciNode) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *PciNode) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
### GetEquipmentChassis

`func (o *PciNode) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *PciNode) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *PciNode) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *PciNode) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *PciNode) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *PciNode) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetGraphicsCards

`func (o *PciNode) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *PciNode) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *PciNode) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *PciNode) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *PciNode) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *PciNode) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetInterconnects

`func (o *PciNode) GetInterconnects() []EquipmentInterconnectRelationship`

GetInterconnects returns the Interconnects field if non-nil, zero value otherwise.

### GetInterconnectsOk

`func (o *PciNode) GetInterconnectsOk() (*[]EquipmentInterconnectRelationship, bool)`

GetInterconnectsOk returns a tuple with the Interconnects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnects

`func (o *PciNode) SetInterconnects(v []EquipmentInterconnectRelationship)`

SetInterconnects sets Interconnects field to given value.

### HasInterconnects

`func (o *PciNode) HasInterconnects() bool`

HasInterconnects returns a boolean if a field has been set.

### SetInterconnectsNil

`func (o *PciNode) SetInterconnectsNil(b bool)`

 SetInterconnectsNil sets the value for Interconnects to be an explicit nil

### UnsetInterconnects
`func (o *PciNode) UnsetInterconnects()`

UnsetInterconnects ensures that no value is present for Interconnects, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *PciNode) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciNode) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciNode) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciNode) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *PciNode) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *PciNode) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetLocatorLed

`func (o *PciNode) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *PciNode) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *PciNode) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *PciNode) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### SetLocatorLedNil

`func (o *PciNode) SetLocatorLedNil(b bool)`

 SetLocatorLedNil sets the value for LocatorLed to be an explicit nil

### UnsetLocatorLed
`func (o *PciNode) UnsetLocatorLed()`

UnsetLocatorLed ensures that no value is present for LocatorLed, not even an explicit nil
### GetRegisteredDevice

`func (o *PciNode) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciNode) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciNode) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciNode) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *PciNode) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *PciNode) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSharedGraphicsCards

`func (o *PciNode) GetSharedGraphicsCards() []EquipmentSharedGraphicsCardRelationship`

GetSharedGraphicsCards returns the SharedGraphicsCards field if non-nil, zero value otherwise.

### GetSharedGraphicsCardsOk

`func (o *PciNode) GetSharedGraphicsCardsOk() (*[]EquipmentSharedGraphicsCardRelationship, bool)`

GetSharedGraphicsCardsOk returns a tuple with the SharedGraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedGraphicsCards

`func (o *PciNode) SetSharedGraphicsCards(v []EquipmentSharedGraphicsCardRelationship)`

SetSharedGraphicsCards sets SharedGraphicsCards field to given value.

### HasSharedGraphicsCards

`func (o *PciNode) HasSharedGraphicsCards() bool`

HasSharedGraphicsCards returns a boolean if a field has been set.

### SetSharedGraphicsCardsNil

`func (o *PciNode) SetSharedGraphicsCardsNil(b bool)`

 SetSharedGraphicsCardsNil sets the value for SharedGraphicsCards to be an explicit nil

### UnsetSharedGraphicsCards
`func (o *PciNode) UnsetSharedGraphicsCards()`

UnsetSharedGraphicsCards ensures that no value is present for SharedGraphicsCards, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


