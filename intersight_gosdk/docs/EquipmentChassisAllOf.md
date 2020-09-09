# EquipmentChassisAllOf

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

### NewEquipmentChassisAllOf

`func NewEquipmentChassisAllOf() *EquipmentChassisAllOf`

NewEquipmentChassisAllOf instantiates a new EquipmentChassisAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisAllOfWithDefaults

`func NewEquipmentChassisAllOfWithDefaults() *EquipmentChassisAllOf`

NewEquipmentChassisAllOfWithDefaults instantiates a new EquipmentChassisAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmSummary

`func (o *EquipmentChassisAllOf) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *EquipmentChassisAllOf) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *EquipmentChassisAllOf) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *EquipmentChassisAllOf) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### GetChassisId

`func (o *EquipmentChassisAllOf) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *EquipmentChassisAllOf) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *EquipmentChassisAllOf) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *EquipmentChassisAllOf) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetConnectionPath

`func (o *EquipmentChassisAllOf) GetConnectionPath() string`

GetConnectionPath returns the ConnectionPath field if non-nil, zero value otherwise.

### GetConnectionPathOk

`func (o *EquipmentChassisAllOf) GetConnectionPathOk() (*string, bool)`

GetConnectionPathOk returns a tuple with the ConnectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPath

`func (o *EquipmentChassisAllOf) SetConnectionPath(v string)`

SetConnectionPath sets ConnectionPath field to given value.

### HasConnectionPath

`func (o *EquipmentChassisAllOf) HasConnectionPath() bool`

HasConnectionPath returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *EquipmentChassisAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *EquipmentChassisAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *EquipmentChassisAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *EquipmentChassisAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDescription

`func (o *EquipmentChassisAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentChassisAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentChassisAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentChassisAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFaultSummary

`func (o *EquipmentChassisAllOf) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *EquipmentChassisAllOf) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *EquipmentChassisAllOf) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *EquipmentChassisAllOf) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetManagementMode

`func (o *EquipmentChassisAllOf) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *EquipmentChassisAllOf) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *EquipmentChassisAllOf) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *EquipmentChassisAllOf) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetName

`func (o *EquipmentChassisAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentChassisAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentChassisAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentChassisAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentChassisAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentChassisAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentChassisAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentChassisAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentChassisAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentChassisAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentChassisAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentChassisAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentChassisAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentChassisAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentChassisAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentChassisAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformType

`func (o *EquipmentChassisAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *EquipmentChassisAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *EquipmentChassisAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *EquipmentChassisAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetProductName

`func (o *EquipmentChassisAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EquipmentChassisAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EquipmentChassisAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *EquipmentChassisAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentChassisAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentChassisAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentChassisAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentChassisAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentChassisAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentChassisAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentChassisAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentChassisAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetBlades

`func (o *EquipmentChassisAllOf) GetBlades() []ComputeBladeRelationship`

GetBlades returns the Blades field if non-nil, zero value otherwise.

### GetBladesOk

`func (o *EquipmentChassisAllOf) GetBladesOk() (*[]ComputeBladeRelationship, bool)`

GetBladesOk returns a tuple with the Blades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlades

`func (o *EquipmentChassisAllOf) SetBlades(v []ComputeBladeRelationship)`

SetBlades sets Blades field to given value.

### HasBlades

`func (o *EquipmentChassisAllOf) HasBlades() bool`

HasBlades returns a boolean if a field has been set.

### SetBladesNil

`func (o *EquipmentChassisAllOf) SetBladesNil(b bool)`

 SetBladesNil sets the value for Blades to be an explicit nil

### UnsetBlades
`func (o *EquipmentChassisAllOf) UnsetBlades()`

UnsetBlades ensures that no value is present for Blades, not even an explicit nil
### GetFanmodules

`func (o *EquipmentChassisAllOf) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *EquipmentChassisAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *EquipmentChassisAllOf) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *EquipmentChassisAllOf) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *EquipmentChassisAllOf) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *EquipmentChassisAllOf) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentChassisAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentChassisAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentChassisAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentChassisAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetIoms

`func (o *EquipmentChassisAllOf) GetIoms() []EquipmentIoCardRelationship`

GetIoms returns the Ioms field if non-nil, zero value otherwise.

### GetIomsOk

`func (o *EquipmentChassisAllOf) GetIomsOk() (*[]EquipmentIoCardRelationship, bool)`

GetIomsOk returns a tuple with the Ioms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoms

`func (o *EquipmentChassisAllOf) SetIoms(v []EquipmentIoCardRelationship)`

SetIoms sets Ioms field to given value.

### HasIoms

`func (o *EquipmentChassisAllOf) HasIoms() bool`

HasIoms returns a boolean if a field has been set.

### SetIomsNil

`func (o *EquipmentChassisAllOf) SetIomsNil(b bool)`

 SetIomsNil sets the value for Ioms to be an explicit nil

### UnsetIoms
`func (o *EquipmentChassisAllOf) UnsetIoms()`

UnsetIoms ensures that no value is present for Ioms, not even an explicit nil
### GetLocatorLed

`func (o *EquipmentChassisAllOf) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *EquipmentChassisAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *EquipmentChassisAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *EquipmentChassisAllOf) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetPsuControl

`func (o *EquipmentChassisAllOf) GetPsuControl() EquipmentPsuControlRelationship`

GetPsuControl returns the PsuControl field if non-nil, zero value otherwise.

### GetPsuControlOk

`func (o *EquipmentChassisAllOf) GetPsuControlOk() (*EquipmentPsuControlRelationship, bool)`

GetPsuControlOk returns a tuple with the PsuControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuControl

`func (o *EquipmentChassisAllOf) SetPsuControl(v EquipmentPsuControlRelationship)`

SetPsuControl sets PsuControl field to given value.

### HasPsuControl

`func (o *EquipmentChassisAllOf) HasPsuControl() bool`

HasPsuControl returns a boolean if a field has been set.

### GetPsus

`func (o *EquipmentChassisAllOf) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *EquipmentChassisAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *EquipmentChassisAllOf) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *EquipmentChassisAllOf) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *EquipmentChassisAllOf) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *EquipmentChassisAllOf) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentChassisAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentChassisAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentChassisAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentChassisAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSasexpanders

`func (o *EquipmentChassisAllOf) GetSasexpanders() []StorageSasExpanderRelationship`

GetSasexpanders returns the Sasexpanders field if non-nil, zero value otherwise.

### GetSasexpandersOk

`func (o *EquipmentChassisAllOf) GetSasexpandersOk() (*[]StorageSasExpanderRelationship, bool)`

GetSasexpandersOk returns a tuple with the Sasexpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasexpanders

`func (o *EquipmentChassisAllOf) SetSasexpanders(v []StorageSasExpanderRelationship)`

SetSasexpanders sets Sasexpanders field to given value.

### HasSasexpanders

`func (o *EquipmentChassisAllOf) HasSasexpanders() bool`

HasSasexpanders returns a boolean if a field has been set.

### SetSasexpandersNil

`func (o *EquipmentChassisAllOf) SetSasexpandersNil(b bool)`

 SetSasexpandersNil sets the value for Sasexpanders to be an explicit nil

### UnsetSasexpanders
`func (o *EquipmentChassisAllOf) UnsetSasexpanders()`

UnsetSasexpanders ensures that no value is present for Sasexpanders, not even an explicit nil
### GetSiocs

`func (o *EquipmentChassisAllOf) GetSiocs() []EquipmentSystemIoControllerRelationship`

GetSiocs returns the Siocs field if non-nil, zero value otherwise.

### GetSiocsOk

`func (o *EquipmentChassisAllOf) GetSiocsOk() (*[]EquipmentSystemIoControllerRelationship, bool)`

GetSiocsOk returns a tuple with the Siocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiocs

`func (o *EquipmentChassisAllOf) SetSiocs(v []EquipmentSystemIoControllerRelationship)`

SetSiocs sets Siocs field to given value.

### HasSiocs

`func (o *EquipmentChassisAllOf) HasSiocs() bool`

HasSiocs returns a boolean if a field has been set.

### SetSiocsNil

`func (o *EquipmentChassisAllOf) SetSiocsNil(b bool)`

 SetSiocsNil sets the value for Siocs to be an explicit nil

### UnsetSiocs
`func (o *EquipmentChassisAllOf) UnsetSiocs()`

UnsetSiocs ensures that no value is present for Siocs, not even an explicit nil
### GetStorageEnclosures

`func (o *EquipmentChassisAllOf) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *EquipmentChassisAllOf) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *EquipmentChassisAllOf) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *EquipmentChassisAllOf) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *EquipmentChassisAllOf) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *EquipmentChassisAllOf) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


