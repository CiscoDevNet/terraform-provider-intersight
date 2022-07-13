# StorageBatteryBackupUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.BatteryBackupUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.BatteryBackupUnit"]
**CapacitanceInPercent** | Pointer to **int64** | This holds the capacitance (in percent) of the battery backup unit of the storage controller. | [optional] [readonly] 
**ChargingState** | Pointer to **string** | This holds the charging state of the battery backup unit of the storage controller. | [optional] [readonly] 
**CurrentInAmps** | Pointer to **float32** | This holds the current (in Amps) of the battery backup unit of the storage controller. | [optional] [readonly] 
**DesignCapacityInJoules** | Pointer to **string** | This holds the design Capacity (in joules) of the battery backup unit of the storage controller. | [optional] [readonly] 
**DesignVoltageInVolts** | Pointer to **float32** | This holds the design volatage (in Volts) of the battery backup unit of the storage controller. | [optional] [readonly] 
**DeviceName** | Pointer to **string** | This refers to the device name of the battery backup unit of the storage controller. | [optional] [readonly] 
**IsBatteryPresent** | Pointer to **bool** | This indicates whether the battery is present for the battery backup unit of the storage controller. | [optional] [readonly] 
**IsCapacitor** | Pointer to **bool** | This indicates the capacitor for the battery backup unit of the storage controller. | [optional] [readonly] 
**IsLearnCycleRequested** | Pointer to **bool** | This indicates learn cycle request of the battery backup unit of the storage controller. | [optional] [readonly] 
**IsLearnCycleTransparent** | Pointer to **bool** | This indicates the learn cycle transparent for the battery backup unit of the storage controller. | [optional] [readonly] 
**IsTemperatureHigh** | Pointer to **bool** | This indicates the temperature is high for the battery backup unit of the storage controller. | [optional] [readonly] 
**IsVoltageLow** | Pointer to **bool** | This indicates the voltage is Low for the battery backup unit of the storage controller. | [optional] [readonly] 
**LearnCycleProgressEndTimeStamp** | Pointer to **string** | This refers to learn cycle progress end time of the battery backup unit of the storage controller. | [optional] [readonly] 
**LearnCycleProgressStartTimeStamp** | Pointer to **string** | This refers to learn cycle progress start time of the battery backup unit of the storage controller. | [optional] [readonly] 
**LearnCycleProgressStatus** | Pointer to **string** | This refers to learn cycle progress status of the battery backup unit of the storage controller. | [optional] [readonly] 
**LearnMode** | Pointer to **string** | This refers to the learn mode of the battery backup unit of the storage controller. | [optional] [readonly] 
**ManufacturingDate** | Pointer to **string** | This refers to the manufacture date of the battery backup unit of the storage controller. | [optional] [readonly] 
**ModuleVersion** | Pointer to **string** | This refers to the current module version of the battery backup unit of the storage controller. | [optional] [readonly] 
**NextLearnCycleTimeStamp** | Pointer to **string** | This refers to next learn cycle timestamp of the battery backup unit of the storage controller. | [optional] [readonly] 
**PackEnergyInJoules** | Pointer to **string** | This holds the pack energy (in joules) of the battery backup unit of the storage controller. | [optional] [readonly] 
**RemainingPoolSpaceInPercent** | Pointer to **int64** | This holds the remaining pool space (in percent) of the battery backup unit of the storage controller. | [optional] [readonly] 
**Status** | Pointer to **string** | This holds the current status of the battery backup unit of the storage controller. | [optional] [readonly] 
**TemperatureInCel** | Pointer to **int64** | This holds the temperature (in Celsius) of the battery backup unit of the storage controller. | [optional] [readonly] 
**Type** | Pointer to **string** | This refers to the type of the battery backup unit of the storage controller. | [optional] [readonly] 
**VoltageInVolts** | Pointer to **string** | This holds the volatage (in Volts) of the battery backup unit of the storage controller. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](StorageControllerRelationship.md) |  | [optional] 

## Methods

### NewStorageBatteryBackupUnit

`func NewStorageBatteryBackupUnit(classId string, objectType string, ) *StorageBatteryBackupUnit`

NewStorageBatteryBackupUnit instantiates a new StorageBatteryBackupUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBatteryBackupUnitWithDefaults

`func NewStorageBatteryBackupUnitWithDefaults() *StorageBatteryBackupUnit`

NewStorageBatteryBackupUnitWithDefaults instantiates a new StorageBatteryBackupUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBatteryBackupUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBatteryBackupUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBatteryBackupUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBatteryBackupUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBatteryBackupUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBatteryBackupUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacitanceInPercent

`func (o *StorageBatteryBackupUnit) GetCapacitanceInPercent() int64`

GetCapacitanceInPercent returns the CapacitanceInPercent field if non-nil, zero value otherwise.

### GetCapacitanceInPercentOk

`func (o *StorageBatteryBackupUnit) GetCapacitanceInPercentOk() (*int64, bool)`

GetCapacitanceInPercentOk returns a tuple with the CapacitanceInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacitanceInPercent

`func (o *StorageBatteryBackupUnit) SetCapacitanceInPercent(v int64)`

SetCapacitanceInPercent sets CapacitanceInPercent field to given value.

### HasCapacitanceInPercent

`func (o *StorageBatteryBackupUnit) HasCapacitanceInPercent() bool`

HasCapacitanceInPercent returns a boolean if a field has been set.

### GetChargingState

`func (o *StorageBatteryBackupUnit) GetChargingState() string`

GetChargingState returns the ChargingState field if non-nil, zero value otherwise.

### GetChargingStateOk

`func (o *StorageBatteryBackupUnit) GetChargingStateOk() (*string, bool)`

GetChargingStateOk returns a tuple with the ChargingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingState

`func (o *StorageBatteryBackupUnit) SetChargingState(v string)`

SetChargingState sets ChargingState field to given value.

### HasChargingState

`func (o *StorageBatteryBackupUnit) HasChargingState() bool`

HasChargingState returns a boolean if a field has been set.

### GetCurrentInAmps

`func (o *StorageBatteryBackupUnit) GetCurrentInAmps() float32`

GetCurrentInAmps returns the CurrentInAmps field if non-nil, zero value otherwise.

### GetCurrentInAmpsOk

`func (o *StorageBatteryBackupUnit) GetCurrentInAmpsOk() (*float32, bool)`

GetCurrentInAmpsOk returns a tuple with the CurrentInAmps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentInAmps

`func (o *StorageBatteryBackupUnit) SetCurrentInAmps(v float32)`

SetCurrentInAmps sets CurrentInAmps field to given value.

### HasCurrentInAmps

`func (o *StorageBatteryBackupUnit) HasCurrentInAmps() bool`

HasCurrentInAmps returns a boolean if a field has been set.

### GetDesignCapacityInJoules

`func (o *StorageBatteryBackupUnit) GetDesignCapacityInJoules() string`

GetDesignCapacityInJoules returns the DesignCapacityInJoules field if non-nil, zero value otherwise.

### GetDesignCapacityInJoulesOk

`func (o *StorageBatteryBackupUnit) GetDesignCapacityInJoulesOk() (*string, bool)`

GetDesignCapacityInJoulesOk returns a tuple with the DesignCapacityInJoules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignCapacityInJoules

`func (o *StorageBatteryBackupUnit) SetDesignCapacityInJoules(v string)`

SetDesignCapacityInJoules sets DesignCapacityInJoules field to given value.

### HasDesignCapacityInJoules

`func (o *StorageBatteryBackupUnit) HasDesignCapacityInJoules() bool`

HasDesignCapacityInJoules returns a boolean if a field has been set.

### GetDesignVoltageInVolts

`func (o *StorageBatteryBackupUnit) GetDesignVoltageInVolts() float32`

GetDesignVoltageInVolts returns the DesignVoltageInVolts field if non-nil, zero value otherwise.

### GetDesignVoltageInVoltsOk

`func (o *StorageBatteryBackupUnit) GetDesignVoltageInVoltsOk() (*float32, bool)`

GetDesignVoltageInVoltsOk returns a tuple with the DesignVoltageInVolts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignVoltageInVolts

`func (o *StorageBatteryBackupUnit) SetDesignVoltageInVolts(v float32)`

SetDesignVoltageInVolts sets DesignVoltageInVolts field to given value.

### HasDesignVoltageInVolts

`func (o *StorageBatteryBackupUnit) HasDesignVoltageInVolts() bool`

HasDesignVoltageInVolts returns a boolean if a field has been set.

### GetDeviceName

`func (o *StorageBatteryBackupUnit) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StorageBatteryBackupUnit) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StorageBatteryBackupUnit) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *StorageBatteryBackupUnit) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetIsBatteryPresent

`func (o *StorageBatteryBackupUnit) GetIsBatteryPresent() bool`

GetIsBatteryPresent returns the IsBatteryPresent field if non-nil, zero value otherwise.

### GetIsBatteryPresentOk

`func (o *StorageBatteryBackupUnit) GetIsBatteryPresentOk() (*bool, bool)`

GetIsBatteryPresentOk returns a tuple with the IsBatteryPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBatteryPresent

`func (o *StorageBatteryBackupUnit) SetIsBatteryPresent(v bool)`

SetIsBatteryPresent sets IsBatteryPresent field to given value.

### HasIsBatteryPresent

`func (o *StorageBatteryBackupUnit) HasIsBatteryPresent() bool`

HasIsBatteryPresent returns a boolean if a field has been set.

### GetIsCapacitor

`func (o *StorageBatteryBackupUnit) GetIsCapacitor() bool`

GetIsCapacitor returns the IsCapacitor field if non-nil, zero value otherwise.

### GetIsCapacitorOk

`func (o *StorageBatteryBackupUnit) GetIsCapacitorOk() (*bool, bool)`

GetIsCapacitorOk returns a tuple with the IsCapacitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCapacitor

`func (o *StorageBatteryBackupUnit) SetIsCapacitor(v bool)`

SetIsCapacitor sets IsCapacitor field to given value.

### HasIsCapacitor

`func (o *StorageBatteryBackupUnit) HasIsCapacitor() bool`

HasIsCapacitor returns a boolean if a field has been set.

### GetIsLearnCycleRequested

`func (o *StorageBatteryBackupUnit) GetIsLearnCycleRequested() bool`

GetIsLearnCycleRequested returns the IsLearnCycleRequested field if non-nil, zero value otherwise.

### GetIsLearnCycleRequestedOk

`func (o *StorageBatteryBackupUnit) GetIsLearnCycleRequestedOk() (*bool, bool)`

GetIsLearnCycleRequestedOk returns a tuple with the IsLearnCycleRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLearnCycleRequested

`func (o *StorageBatteryBackupUnit) SetIsLearnCycleRequested(v bool)`

SetIsLearnCycleRequested sets IsLearnCycleRequested field to given value.

### HasIsLearnCycleRequested

`func (o *StorageBatteryBackupUnit) HasIsLearnCycleRequested() bool`

HasIsLearnCycleRequested returns a boolean if a field has been set.

### GetIsLearnCycleTransparent

`func (o *StorageBatteryBackupUnit) GetIsLearnCycleTransparent() bool`

GetIsLearnCycleTransparent returns the IsLearnCycleTransparent field if non-nil, zero value otherwise.

### GetIsLearnCycleTransparentOk

`func (o *StorageBatteryBackupUnit) GetIsLearnCycleTransparentOk() (*bool, bool)`

GetIsLearnCycleTransparentOk returns a tuple with the IsLearnCycleTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLearnCycleTransparent

`func (o *StorageBatteryBackupUnit) SetIsLearnCycleTransparent(v bool)`

SetIsLearnCycleTransparent sets IsLearnCycleTransparent field to given value.

### HasIsLearnCycleTransparent

`func (o *StorageBatteryBackupUnit) HasIsLearnCycleTransparent() bool`

HasIsLearnCycleTransparent returns a boolean if a field has been set.

### GetIsTemperatureHigh

`func (o *StorageBatteryBackupUnit) GetIsTemperatureHigh() bool`

GetIsTemperatureHigh returns the IsTemperatureHigh field if non-nil, zero value otherwise.

### GetIsTemperatureHighOk

`func (o *StorageBatteryBackupUnit) GetIsTemperatureHighOk() (*bool, bool)`

GetIsTemperatureHighOk returns a tuple with the IsTemperatureHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemperatureHigh

`func (o *StorageBatteryBackupUnit) SetIsTemperatureHigh(v bool)`

SetIsTemperatureHigh sets IsTemperatureHigh field to given value.

### HasIsTemperatureHigh

`func (o *StorageBatteryBackupUnit) HasIsTemperatureHigh() bool`

HasIsTemperatureHigh returns a boolean if a field has been set.

### GetIsVoltageLow

`func (o *StorageBatteryBackupUnit) GetIsVoltageLow() bool`

GetIsVoltageLow returns the IsVoltageLow field if non-nil, zero value otherwise.

### GetIsVoltageLowOk

`func (o *StorageBatteryBackupUnit) GetIsVoltageLowOk() (*bool, bool)`

GetIsVoltageLowOk returns a tuple with the IsVoltageLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoltageLow

`func (o *StorageBatteryBackupUnit) SetIsVoltageLow(v bool)`

SetIsVoltageLow sets IsVoltageLow field to given value.

### HasIsVoltageLow

`func (o *StorageBatteryBackupUnit) HasIsVoltageLow() bool`

HasIsVoltageLow returns a boolean if a field has been set.

### GetLearnCycleProgressEndTimeStamp

`func (o *StorageBatteryBackupUnit) GetLearnCycleProgressEndTimeStamp() string`

GetLearnCycleProgressEndTimeStamp returns the LearnCycleProgressEndTimeStamp field if non-nil, zero value otherwise.

### GetLearnCycleProgressEndTimeStampOk

`func (o *StorageBatteryBackupUnit) GetLearnCycleProgressEndTimeStampOk() (*string, bool)`

GetLearnCycleProgressEndTimeStampOk returns a tuple with the LearnCycleProgressEndTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnCycleProgressEndTimeStamp

`func (o *StorageBatteryBackupUnit) SetLearnCycleProgressEndTimeStamp(v string)`

SetLearnCycleProgressEndTimeStamp sets LearnCycleProgressEndTimeStamp field to given value.

### HasLearnCycleProgressEndTimeStamp

`func (o *StorageBatteryBackupUnit) HasLearnCycleProgressEndTimeStamp() bool`

HasLearnCycleProgressEndTimeStamp returns a boolean if a field has been set.

### GetLearnCycleProgressStartTimeStamp

`func (o *StorageBatteryBackupUnit) GetLearnCycleProgressStartTimeStamp() string`

GetLearnCycleProgressStartTimeStamp returns the LearnCycleProgressStartTimeStamp field if non-nil, zero value otherwise.

### GetLearnCycleProgressStartTimeStampOk

`func (o *StorageBatteryBackupUnit) GetLearnCycleProgressStartTimeStampOk() (*string, bool)`

GetLearnCycleProgressStartTimeStampOk returns a tuple with the LearnCycleProgressStartTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnCycleProgressStartTimeStamp

`func (o *StorageBatteryBackupUnit) SetLearnCycleProgressStartTimeStamp(v string)`

SetLearnCycleProgressStartTimeStamp sets LearnCycleProgressStartTimeStamp field to given value.

### HasLearnCycleProgressStartTimeStamp

`func (o *StorageBatteryBackupUnit) HasLearnCycleProgressStartTimeStamp() bool`

HasLearnCycleProgressStartTimeStamp returns a boolean if a field has been set.

### GetLearnCycleProgressStatus

`func (o *StorageBatteryBackupUnit) GetLearnCycleProgressStatus() string`

GetLearnCycleProgressStatus returns the LearnCycleProgressStatus field if non-nil, zero value otherwise.

### GetLearnCycleProgressStatusOk

`func (o *StorageBatteryBackupUnit) GetLearnCycleProgressStatusOk() (*string, bool)`

GetLearnCycleProgressStatusOk returns a tuple with the LearnCycleProgressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnCycleProgressStatus

`func (o *StorageBatteryBackupUnit) SetLearnCycleProgressStatus(v string)`

SetLearnCycleProgressStatus sets LearnCycleProgressStatus field to given value.

### HasLearnCycleProgressStatus

`func (o *StorageBatteryBackupUnit) HasLearnCycleProgressStatus() bool`

HasLearnCycleProgressStatus returns a boolean if a field has been set.

### GetLearnMode

`func (o *StorageBatteryBackupUnit) GetLearnMode() string`

GetLearnMode returns the LearnMode field if non-nil, zero value otherwise.

### GetLearnModeOk

`func (o *StorageBatteryBackupUnit) GetLearnModeOk() (*string, bool)`

GetLearnModeOk returns a tuple with the LearnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnMode

`func (o *StorageBatteryBackupUnit) SetLearnMode(v string)`

SetLearnMode sets LearnMode field to given value.

### HasLearnMode

`func (o *StorageBatteryBackupUnit) HasLearnMode() bool`

HasLearnMode returns a boolean if a field has been set.

### GetManufacturingDate

`func (o *StorageBatteryBackupUnit) GetManufacturingDate() string`

GetManufacturingDate returns the ManufacturingDate field if non-nil, zero value otherwise.

### GetManufacturingDateOk

`func (o *StorageBatteryBackupUnit) GetManufacturingDateOk() (*string, bool)`

GetManufacturingDateOk returns a tuple with the ManufacturingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturingDate

`func (o *StorageBatteryBackupUnit) SetManufacturingDate(v string)`

SetManufacturingDate sets ManufacturingDate field to given value.

### HasManufacturingDate

`func (o *StorageBatteryBackupUnit) HasManufacturingDate() bool`

HasManufacturingDate returns a boolean if a field has been set.

### GetModuleVersion

`func (o *StorageBatteryBackupUnit) GetModuleVersion() string`

GetModuleVersion returns the ModuleVersion field if non-nil, zero value otherwise.

### GetModuleVersionOk

`func (o *StorageBatteryBackupUnit) GetModuleVersionOk() (*string, bool)`

GetModuleVersionOk returns a tuple with the ModuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleVersion

`func (o *StorageBatteryBackupUnit) SetModuleVersion(v string)`

SetModuleVersion sets ModuleVersion field to given value.

### HasModuleVersion

`func (o *StorageBatteryBackupUnit) HasModuleVersion() bool`

HasModuleVersion returns a boolean if a field has been set.

### GetNextLearnCycleTimeStamp

`func (o *StorageBatteryBackupUnit) GetNextLearnCycleTimeStamp() string`

GetNextLearnCycleTimeStamp returns the NextLearnCycleTimeStamp field if non-nil, zero value otherwise.

### GetNextLearnCycleTimeStampOk

`func (o *StorageBatteryBackupUnit) GetNextLearnCycleTimeStampOk() (*string, bool)`

GetNextLearnCycleTimeStampOk returns a tuple with the NextLearnCycleTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLearnCycleTimeStamp

`func (o *StorageBatteryBackupUnit) SetNextLearnCycleTimeStamp(v string)`

SetNextLearnCycleTimeStamp sets NextLearnCycleTimeStamp field to given value.

### HasNextLearnCycleTimeStamp

`func (o *StorageBatteryBackupUnit) HasNextLearnCycleTimeStamp() bool`

HasNextLearnCycleTimeStamp returns a boolean if a field has been set.

### GetPackEnergyInJoules

`func (o *StorageBatteryBackupUnit) GetPackEnergyInJoules() string`

GetPackEnergyInJoules returns the PackEnergyInJoules field if non-nil, zero value otherwise.

### GetPackEnergyInJoulesOk

`func (o *StorageBatteryBackupUnit) GetPackEnergyInJoulesOk() (*string, bool)`

GetPackEnergyInJoulesOk returns a tuple with the PackEnergyInJoules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackEnergyInJoules

`func (o *StorageBatteryBackupUnit) SetPackEnergyInJoules(v string)`

SetPackEnergyInJoules sets PackEnergyInJoules field to given value.

### HasPackEnergyInJoules

`func (o *StorageBatteryBackupUnit) HasPackEnergyInJoules() bool`

HasPackEnergyInJoules returns a boolean if a field has been set.

### GetRemainingPoolSpaceInPercent

`func (o *StorageBatteryBackupUnit) GetRemainingPoolSpaceInPercent() int64`

GetRemainingPoolSpaceInPercent returns the RemainingPoolSpaceInPercent field if non-nil, zero value otherwise.

### GetRemainingPoolSpaceInPercentOk

`func (o *StorageBatteryBackupUnit) GetRemainingPoolSpaceInPercentOk() (*int64, bool)`

GetRemainingPoolSpaceInPercentOk returns a tuple with the RemainingPoolSpaceInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingPoolSpaceInPercent

`func (o *StorageBatteryBackupUnit) SetRemainingPoolSpaceInPercent(v int64)`

SetRemainingPoolSpaceInPercent sets RemainingPoolSpaceInPercent field to given value.

### HasRemainingPoolSpaceInPercent

`func (o *StorageBatteryBackupUnit) HasRemainingPoolSpaceInPercent() bool`

HasRemainingPoolSpaceInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBatteryBackupUnit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBatteryBackupUnit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBatteryBackupUnit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBatteryBackupUnit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemperatureInCel

`func (o *StorageBatteryBackupUnit) GetTemperatureInCel() int64`

GetTemperatureInCel returns the TemperatureInCel field if non-nil, zero value otherwise.

### GetTemperatureInCelOk

`func (o *StorageBatteryBackupUnit) GetTemperatureInCelOk() (*int64, bool)`

GetTemperatureInCelOk returns a tuple with the TemperatureInCel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureInCel

`func (o *StorageBatteryBackupUnit) SetTemperatureInCel(v int64)`

SetTemperatureInCel sets TemperatureInCel field to given value.

### HasTemperatureInCel

`func (o *StorageBatteryBackupUnit) HasTemperatureInCel() bool`

HasTemperatureInCel returns a boolean if a field has been set.

### GetType

`func (o *StorageBatteryBackupUnit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageBatteryBackupUnit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageBatteryBackupUnit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageBatteryBackupUnit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVoltageInVolts

`func (o *StorageBatteryBackupUnit) GetVoltageInVolts() string`

GetVoltageInVolts returns the VoltageInVolts field if non-nil, zero value otherwise.

### GetVoltageInVoltsOk

`func (o *StorageBatteryBackupUnit) GetVoltageInVoltsOk() (*string, bool)`

GetVoltageInVoltsOk returns a tuple with the VoltageInVolts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltageInVolts

`func (o *StorageBatteryBackupUnit) SetVoltageInVolts(v string)`

SetVoltageInVolts sets VoltageInVolts field to given value.

### HasVoltageInVolts

`func (o *StorageBatteryBackupUnit) HasVoltageInVolts() bool`

HasVoltageInVolts returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageBatteryBackupUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageBatteryBackupUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageBatteryBackupUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageBatteryBackupUnit) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageBatteryBackupUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageBatteryBackupUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageBatteryBackupUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageBatteryBackupUnit) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageBatteryBackupUnit) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageBatteryBackupUnit) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageBatteryBackupUnit) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageBatteryBackupUnit) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


