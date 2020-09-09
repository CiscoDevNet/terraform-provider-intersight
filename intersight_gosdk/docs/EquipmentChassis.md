# EquipmentChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmSummary** | Pointer to [**ComputeAlarmSummary**](compute.AlarmSummary.md) |  | [optional] 
**ChassisId** | Pointer to **int64** | The assigned identifier for a chassis. | [optional] [readonly] 
**ConnectionPath** | Pointer to **string** | This field identifies the connectivity path for the chassis enclosure. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | This field identifies the connectivity status for the chassis enclosure. | [optional] [readonly] 
**Description** | Pointer to **string** | This field is to provide description for chassis model. | [optional] [readonly] 
**FaultSummary** | Pointer to **int64** | This field summarizes the faults on the chassis enclosure. | [optional] 
**ManagementMode** | Pointer to **string** | The management mode of the blade server chassis. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [readonly] [default to "IntersightStandalone"]
**Name** | Pointer to **string** | This field identifies the name for the chassis enclosure. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the Chassis Operational State. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part Number identifier for the chassis enclosure. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the chassis enclosure. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type that the chassis is a part of. | [optional] 
**ProductName** | Pointer to **string** | This field identifies the Product Name for the chassis enclosure. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stock Keeping Unit for the chassis enclosure. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for the chassis enclosure. | [optional] [readonly] 
**Blades** | Pointer to [**[]ComputeBladeRelationship**](compute.Blade.Relationship.md) | An array of relationships to computeBlade resources. | [optional] [readonly] 
**Fanmodules** | Pointer to [**[]EquipmentFanModuleRelationship**](equipment.FanModule.Relationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**Ioms** | Pointer to [**[]EquipmentIoCardRelationship**](equipment.IoCard.Relationship.md) | An array of relationships to equipmentIoCard resources. | [optional] [readonly] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**PsuControl** | Pointer to [**EquipmentPsuControlRelationship**](equipment.PsuControl.Relationship.md) |  | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](equipment.Psu.Relationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Sasexpanders** | Pointer to [**[]StorageSasExpanderRelationship**](storage.SasExpander.Relationship.md) | An array of relationships to storageSasExpander resources. | [optional] [readonly] 
**Siocs** | Pointer to [**[]EquipmentSystemIoControllerRelationship**](equipment.SystemIoController.Relationship.md) | An array of relationships to equipmentSystemIoController resources. | [optional] [readonly] 
**StorageEnclosures** | Pointer to [**[]StorageEnclosureRelationship**](storage.Enclosure.Relationship.md) | An array of relationships to storageEnclosure resources. | [optional] [readonly] 

## Methods

### NewEquipmentChassis

`func NewEquipmentChassis() *EquipmentChassis`

NewEquipmentChassis instantiates a new EquipmentChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisWithDefaults

`func NewEquipmentChassisWithDefaults() *EquipmentChassis`

NewEquipmentChassisWithDefaults instantiates a new EquipmentChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmSummary

`func (o *EquipmentChassis) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *EquipmentChassis) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *EquipmentChassis) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *EquipmentChassis) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### GetChassisId

`func (o *EquipmentChassis) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *EquipmentChassis) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *EquipmentChassis) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *EquipmentChassis) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetConnectionPath

`func (o *EquipmentChassis) GetConnectionPath() string`

GetConnectionPath returns the ConnectionPath field if non-nil, zero value otherwise.

### GetConnectionPathOk

`func (o *EquipmentChassis) GetConnectionPathOk() (*string, bool)`

GetConnectionPathOk returns a tuple with the ConnectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPath

`func (o *EquipmentChassis) SetConnectionPath(v string)`

SetConnectionPath sets ConnectionPath field to given value.

### HasConnectionPath

`func (o *EquipmentChassis) HasConnectionPath() bool`

HasConnectionPath returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *EquipmentChassis) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *EquipmentChassis) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *EquipmentChassis) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *EquipmentChassis) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDescription

`func (o *EquipmentChassis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentChassis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentChassis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentChassis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFaultSummary

`func (o *EquipmentChassis) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *EquipmentChassis) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *EquipmentChassis) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *EquipmentChassis) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetManagementMode

`func (o *EquipmentChassis) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *EquipmentChassis) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *EquipmentChassis) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *EquipmentChassis) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetName

`func (o *EquipmentChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentChassis) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentChassis) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentChassis) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentChassis) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentChassis) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentChassis) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentChassis) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentChassis) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentChassis) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentChassis) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentChassis) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentChassis) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentChassis) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentChassis) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformType

`func (o *EquipmentChassis) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *EquipmentChassis) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *EquipmentChassis) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *EquipmentChassis) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetProductName

`func (o *EquipmentChassis) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EquipmentChassis) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EquipmentChassis) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *EquipmentChassis) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentChassis) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentChassis) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentChassis) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentChassis) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentChassis) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentChassis) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentChassis) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentChassis) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetBlades

`func (o *EquipmentChassis) GetBlades() []ComputeBladeRelationship`

GetBlades returns the Blades field if non-nil, zero value otherwise.

### GetBladesOk

`func (o *EquipmentChassis) GetBladesOk() (*[]ComputeBladeRelationship, bool)`

GetBladesOk returns a tuple with the Blades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlades

`func (o *EquipmentChassis) SetBlades(v []ComputeBladeRelationship)`

SetBlades sets Blades field to given value.

### HasBlades

`func (o *EquipmentChassis) HasBlades() bool`

HasBlades returns a boolean if a field has been set.

### SetBladesNil

`func (o *EquipmentChassis) SetBladesNil(b bool)`

 SetBladesNil sets the value for Blades to be an explicit nil

### UnsetBlades
`func (o *EquipmentChassis) UnsetBlades()`

UnsetBlades ensures that no value is present for Blades, not even an explicit nil
### GetFanmodules

`func (o *EquipmentChassis) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *EquipmentChassis) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *EquipmentChassis) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *EquipmentChassis) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *EquipmentChassis) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *EquipmentChassis) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentChassis) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentChassis) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentChassis) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentChassis) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetIoms

`func (o *EquipmentChassis) GetIoms() []EquipmentIoCardRelationship`

GetIoms returns the Ioms field if non-nil, zero value otherwise.

### GetIomsOk

`func (o *EquipmentChassis) GetIomsOk() (*[]EquipmentIoCardRelationship, bool)`

GetIomsOk returns a tuple with the Ioms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoms

`func (o *EquipmentChassis) SetIoms(v []EquipmentIoCardRelationship)`

SetIoms sets Ioms field to given value.

### HasIoms

`func (o *EquipmentChassis) HasIoms() bool`

HasIoms returns a boolean if a field has been set.

### SetIomsNil

`func (o *EquipmentChassis) SetIomsNil(b bool)`

 SetIomsNil sets the value for Ioms to be an explicit nil

### UnsetIoms
`func (o *EquipmentChassis) UnsetIoms()`

UnsetIoms ensures that no value is present for Ioms, not even an explicit nil
### GetLocatorLed

`func (o *EquipmentChassis) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *EquipmentChassis) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *EquipmentChassis) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *EquipmentChassis) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetPsuControl

`func (o *EquipmentChassis) GetPsuControl() EquipmentPsuControlRelationship`

GetPsuControl returns the PsuControl field if non-nil, zero value otherwise.

### GetPsuControlOk

`func (o *EquipmentChassis) GetPsuControlOk() (*EquipmentPsuControlRelationship, bool)`

GetPsuControlOk returns a tuple with the PsuControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuControl

`func (o *EquipmentChassis) SetPsuControl(v EquipmentPsuControlRelationship)`

SetPsuControl sets PsuControl field to given value.

### HasPsuControl

`func (o *EquipmentChassis) HasPsuControl() bool`

HasPsuControl returns a boolean if a field has been set.

### GetPsus

`func (o *EquipmentChassis) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *EquipmentChassis) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *EquipmentChassis) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *EquipmentChassis) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *EquipmentChassis) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *EquipmentChassis) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentChassis) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentChassis) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentChassis) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentChassis) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSasexpanders

`func (o *EquipmentChassis) GetSasexpanders() []StorageSasExpanderRelationship`

GetSasexpanders returns the Sasexpanders field if non-nil, zero value otherwise.

### GetSasexpandersOk

`func (o *EquipmentChassis) GetSasexpandersOk() (*[]StorageSasExpanderRelationship, bool)`

GetSasexpandersOk returns a tuple with the Sasexpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasexpanders

`func (o *EquipmentChassis) SetSasexpanders(v []StorageSasExpanderRelationship)`

SetSasexpanders sets Sasexpanders field to given value.

### HasSasexpanders

`func (o *EquipmentChassis) HasSasexpanders() bool`

HasSasexpanders returns a boolean if a field has been set.

### SetSasexpandersNil

`func (o *EquipmentChassis) SetSasexpandersNil(b bool)`

 SetSasexpandersNil sets the value for Sasexpanders to be an explicit nil

### UnsetSasexpanders
`func (o *EquipmentChassis) UnsetSasexpanders()`

UnsetSasexpanders ensures that no value is present for Sasexpanders, not even an explicit nil
### GetSiocs

`func (o *EquipmentChassis) GetSiocs() []EquipmentSystemIoControllerRelationship`

GetSiocs returns the Siocs field if non-nil, zero value otherwise.

### GetSiocsOk

`func (o *EquipmentChassis) GetSiocsOk() (*[]EquipmentSystemIoControllerRelationship, bool)`

GetSiocsOk returns a tuple with the Siocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiocs

`func (o *EquipmentChassis) SetSiocs(v []EquipmentSystemIoControllerRelationship)`

SetSiocs sets Siocs field to given value.

### HasSiocs

`func (o *EquipmentChassis) HasSiocs() bool`

HasSiocs returns a boolean if a field has been set.

### SetSiocsNil

`func (o *EquipmentChassis) SetSiocsNil(b bool)`

 SetSiocsNil sets the value for Siocs to be an explicit nil

### UnsetSiocs
`func (o *EquipmentChassis) UnsetSiocs()`

UnsetSiocs ensures that no value is present for Siocs, not even an explicit nil
### GetStorageEnclosures

`func (o *EquipmentChassis) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *EquipmentChassis) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *EquipmentChassis) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *EquipmentChassis) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *EquipmentChassis) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *EquipmentChassis) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


