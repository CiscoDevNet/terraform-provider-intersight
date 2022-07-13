# StorageBatteryBackupUnitAllOf

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

### NewStorageBatteryBackupUnitAllOf

`func NewStorageBatteryBackupUnitAllOf(classId string, objectType string, ) *StorageBatteryBackupUnitAllOf`

NewStorageBatteryBackupUnitAllOf instantiates a new StorageBatteryBackupUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBatteryBackupUnitAllOfWithDefaults

`func NewStorageBatteryBackupUnitAllOfWithDefaults() *StorageBatteryBackupUnitAllOf`

NewStorageBatteryBackupUnitAllOfWithDefaults instantiates a new StorageBatteryBackupUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBatteryBackupUnitAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBatteryBackupUnitAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBatteryBackupUnitAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBatteryBackupUnitAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBatteryBackupUnitAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBatteryBackupUnitAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacitanceInPercent

`func (o *StorageBatteryBackupUnitAllOf) GetCapacitanceInPercent() int64`

GetCapacitanceInPercent returns the CapacitanceInPercent field if non-nil, zero value otherwise.

### GetCapacitanceInPercentOk

`func (o *StorageBatteryBackupUnitAllOf) GetCapacitanceInPercentOk() (*int64, bool)`

GetCapacitanceInPercentOk returns a tuple with the CapacitanceInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacitanceInPercent

`func (o *StorageBatteryBackupUnitAllOf) SetCapacitanceInPercent(v int64)`

SetCapacitanceInPercent sets CapacitanceInPercent field to given value.

### HasCapacitanceInPercent

`func (o *StorageBatteryBackupUnitAllOf) HasCapacitanceInPercent() bool`

HasCapacitanceInPercent returns a boolean if a field has been set.

### GetChargingState

`func (o *StorageBatteryBackupUnitAllOf) GetChargingState() string`

GetChargingState returns the ChargingState field if non-nil, zero value otherwise.

### GetChargingStateOk

`func (o *StorageBatteryBackupUnitAllOf) GetChargingStateOk() (*string, bool)`

GetChargingStateOk returns a tuple with the ChargingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingState

`func (o *StorageBatteryBackupUnitAllOf) SetChargingState(v string)`

SetChargingState sets ChargingState field to given value.

### HasChargingState

`func (o *StorageBatteryBackupUnitAllOf) HasChargingState() bool`

HasChargingState returns a boolean if a field has been set.

### GetCurrentInAmps

`func (o *StorageBatteryBackupUnitAllOf) GetCurrentInAmps() float32`

GetCurrentInAmps returns the CurrentInAmps field if non-nil, zero value otherwise.

### GetCurrentInAmpsOk

`func (o *StorageBatteryBackupUnitAllOf) GetCurrentInAmpsOk() (*float32, bool)`

GetCurrentInAmpsOk returns a tuple with the CurrentInAmps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentInAmps

`func (o *StorageBatteryBackupUnitAllOf) SetCurrentInAmps(v float32)`

SetCurrentInAmps sets CurrentInAmps field to given value.

### HasCurrentInAmps

`func (o *StorageBatteryBackupUnitAllOf) HasCurrentInAmps() bool`

HasCurrentInAmps returns a boolean if a field has been set.

### GetDesignCapacityInJoules

`func (o *StorageBatteryBackupUnitAllOf) GetDesignCapacityInJoules() string`

GetDesignCapacityInJoules returns the DesignCapacityInJoules field if non-nil, zero value otherwise.

### GetDesignCapacityInJoulesOk

`func (o *StorageBatteryBackupUnitAllOf) GetDesignCapacityInJoulesOk() (*string, bool)`

GetDesignCapacityInJoulesOk returns a tuple with the DesignCapacityInJoules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignCapacityInJoules

`func (o *StorageBatteryBackupUnitAllOf) SetDesignCapacityInJoules(v string)`

SetDesignCapacityInJoules sets DesignCapacityInJoules field to given value.

### HasDesignCapacityInJoules

`func (o *StorageBatteryBackupUnitAllOf) HasDesignCapacityInJoules() bool`

HasDesignCapacityInJoules returns a boolean if a field has been set.

### GetDesignVoltageInVolts

`func (o *StorageBatteryBackupUnitAllOf) GetDesignVoltageInVolts() float32`

GetDesignVoltageInVolts returns the DesignVoltageInVolts field if non-nil, zero value otherwise.

### GetDesignVoltageInVoltsOk

`func (o *StorageBatteryBackupUnitAllOf) GetDesignVoltageInVoltsOk() (*float32, bool)`

GetDesignVoltageInVoltsOk returns a tuple with the DesignVoltageInVolts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignVoltageInVolts

`func (o *StorageBatteryBackupUnitAllOf) SetDesignVoltageInVolts(v float32)`

SetDesignVoltageInVolts sets DesignVoltageInVolts field to given value.

### HasDesignVoltageInVolts

`func (o *StorageBatteryBackupUnitAllOf) HasDesignVoltageInVolts() bool`

HasDesignVoltageInVolts returns a boolean if a field has been set.

### GetDeviceName

`func (o *StorageBatteryBackupUnitAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StorageBatteryBackupUnitAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StorageBatteryBackupUnitAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *StorageBatteryBackupUnitAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetIsBatteryPresent

`func (o *StorageBatteryBackupUnitAllOf) GetIsBatteryPresent() bool`

GetIsBatteryPresent returns the IsBatteryPresent field if non-nil, zero value otherwise.

### GetIsBatteryPresentOk

`func (o *StorageBatteryBackupUnitAllOf) GetIsBatteryPresentOk() (*bool, bool)`

GetIsBatteryPresentOk returns a tuple with the IsBatteryPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBatteryPresent

`func (o *StorageBatteryBackupUnitAllOf) SetIsBatteryPresent(v bool)`

SetIsBatteryPresent sets IsBatteryPresent field to given value.

### HasIsBatteryPresent

`func (o *StorageBatteryBackupUnitAllOf) HasIsBatteryPresent() bool`

HasIsBatteryPresent returns a boolean if a field has been set.

### GetIsCapacitor

`func (o *StorageBatteryBackupUnitAllOf) GetIsCapacitor() bool`

GetIsCapacitor returns the IsCapacitor field if non-nil, zero value otherwise.

### GetIsCapacitorOk

`func (o *StorageBatteryBackupUnitAllOf) GetIsCapacitorOk() (*bool, bool)`

GetIsCapacitorOk returns a tuple with the IsCapacitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCapacitor

`func (o *StorageBatteryBackupUnitAllOf) SetIsCapacitor(v bool)`

SetIsCapacitor sets IsCapacitor field to given value.

### HasIsCapacitor

`func (o *StorageBatteryBackupUnitAllOf) HasIsCapacitor() bool`

HasIsCapacitor returns a boolean if a field has been set.

### GetIsLearnCycleRequested

`func (o *StorageBatteryBackupUnitAllOf) GetIsLearnCycleRequested() bool`

GetIsLearnCycleRequested returns the IsLearnCycleRequested field if non-nil, zero value otherwise.

### GetIsLearnCycleRequestedOk

`func (o *StorageBatteryBackupUnitAllOf) GetIsLearnCycleRequestedOk() (*bool, bool)`

GetIsLearnCycleRequestedOk returns a tuple with the IsLearnCycleRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLearnCycleRequested

`func (o *StorageBatteryBackupUnitAllOf) SetIsLearnCycleRequested(v bool)`

SetIsLearnCycleRequested sets IsLearnCycleRequested field to given value.

### HasIsLearnCycleRequested

`func (o *StorageBatteryBackupUnitAllOf) HasIsLearnCycleRequested() bool`

HasIsLearnCycleRequested returns a boolean if a field has been set.

### GetIsLearnCycleTransparent

`func (o *StorageBatteryBackupUnitAllOf) GetIsLearnCycleTransparent() bool`

GetIsLearnCycleTransparent returns the IsLearnCycleTransparent field if non-nil, zero value otherwise.

### GetIsLearnCycleTransparentOk

`func (o *StorageBatteryBackupUnitAllOf) GetIsLearnCycleTransparentOk() (*bool, bool)`

GetIsLearnCycleTransparentOk returns a tuple with the IsLearnCycleTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLearnCycleTransparent

`func (o *StorageBatteryBackupUnitAllOf) SetIsLearnCycleTransparent(v bool)`

SetIsLearnCycleTransparent sets IsLearnCycleTransparent field to given value.

### HasIsLearnCycleTransparent

`func (o *StorageBatteryBackupUnitAllOf) HasIsLearnCycleTransparent() bool`

HasIsLearnCycleTransparent returns a boolean if a field has been set.

### GetIsTemperatureHigh

`func (o *StorageBatteryBackupUnitAllOf) GetIsTemperatureHigh() bool`

GetIsTemperatureHigh returns the IsTemperatureHigh field if non-nil, zero value otherwise.

### GetIsTemperatureHighOk

`func (o *StorageBatteryBackupUnitAllOf) GetIsTemperatureHighOk() (*bool, bool)`

GetIsTemperatureHighOk returns a tuple with the IsTemperatureHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemperatureHigh

`func (o *StorageBatteryBackupUnitAllOf) SetIsTemperatureHigh(v bool)`

SetIsTemperatureHigh sets IsTemperatureHigh field to given value.

### HasIsTemperatureHigh

`func (o *StorageBatteryBackupUnitAllOf) HasIsTemperatureHigh() bool`

HasIsTemperatureHigh returns a boolean if a field has been set.

### GetIsVoltageLow

`func (o *StorageBatteryBackupUnitAllOf) GetIsVoltageLow() bool`

GetIsVoltageLow returns the IsVoltageLow field if non-nil, zero value otherwise.

### GetIsVoltageLowOk

`func (o *StorageBatteryBackupUnitAllOf) GetIsVoltageLowOk() (*bool, bool)`

GetIsVoltageLowOk returns a tuple with the IsVoltageLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoltageLow

`func (o *StorageBatteryBackupUnitAllOf) SetIsVoltageLow(v bool)`

SetIsVoltageLow sets IsVoltageLow field to given value.

### HasIsVoltageLow

`func (o *StorageBatteryBackupUnitAllOf) HasIsVoltageLow() bool`

HasIsVoltageLow returns a boolean if a field has been set.

### GetLearnCycleProgressEndTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) GetLearnCycleProgressEndTimeStamp() string`

GetLearnCycleProgressEndTimeStamp returns the LearnCycleProgressEndTimeStamp field if non-nil, zero value otherwise.

### GetLearnCycleProgressEndTimeStampOk

`func (o *StorageBatteryBackupUnitAllOf) GetLearnCycleProgressEndTimeStampOk() (*string, bool)`

GetLearnCycleProgressEndTimeStampOk returns a tuple with the LearnCycleProgressEndTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnCycleProgressEndTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) SetLearnCycleProgressEndTimeStamp(v string)`

SetLearnCycleProgressEndTimeStamp sets LearnCycleProgressEndTimeStamp field to given value.

### HasLearnCycleProgressEndTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) HasLearnCycleProgressEndTimeStamp() bool`

HasLearnCycleProgressEndTimeStamp returns a boolean if a field has been set.

### GetLearnCycleProgressStartTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) GetLearnCycleProgressStartTimeStamp() string`

GetLearnCycleProgressStartTimeStamp returns the LearnCycleProgressStartTimeStamp field if non-nil, zero value otherwise.

### GetLearnCycleProgressStartTimeStampOk

`func (o *StorageBatteryBackupUnitAllOf) GetLearnCycleProgressStartTimeStampOk() (*string, bool)`

GetLearnCycleProgressStartTimeStampOk returns a tuple with the LearnCycleProgressStartTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnCycleProgressStartTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) SetLearnCycleProgressStartTimeStamp(v string)`

SetLearnCycleProgressStartTimeStamp sets LearnCycleProgressStartTimeStamp field to given value.

### HasLearnCycleProgressStartTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) HasLearnCycleProgressStartTimeStamp() bool`

HasLearnCycleProgressStartTimeStamp returns a boolean if a field has been set.

### GetLearnCycleProgressStatus

`func (o *StorageBatteryBackupUnitAllOf) GetLearnCycleProgressStatus() string`

GetLearnCycleProgressStatus returns the LearnCycleProgressStatus field if non-nil, zero value otherwise.

### GetLearnCycleProgressStatusOk

`func (o *StorageBatteryBackupUnitAllOf) GetLearnCycleProgressStatusOk() (*string, bool)`

GetLearnCycleProgressStatusOk returns a tuple with the LearnCycleProgressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnCycleProgressStatus

`func (o *StorageBatteryBackupUnitAllOf) SetLearnCycleProgressStatus(v string)`

SetLearnCycleProgressStatus sets LearnCycleProgressStatus field to given value.

### HasLearnCycleProgressStatus

`func (o *StorageBatteryBackupUnitAllOf) HasLearnCycleProgressStatus() bool`

HasLearnCycleProgressStatus returns a boolean if a field has been set.

### GetLearnMode

`func (o *StorageBatteryBackupUnitAllOf) GetLearnMode() string`

GetLearnMode returns the LearnMode field if non-nil, zero value otherwise.

### GetLearnModeOk

`func (o *StorageBatteryBackupUnitAllOf) GetLearnModeOk() (*string, bool)`

GetLearnModeOk returns a tuple with the LearnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnMode

`func (o *StorageBatteryBackupUnitAllOf) SetLearnMode(v string)`

SetLearnMode sets LearnMode field to given value.

### HasLearnMode

`func (o *StorageBatteryBackupUnitAllOf) HasLearnMode() bool`

HasLearnMode returns a boolean if a field has been set.

### GetManufacturingDate

`func (o *StorageBatteryBackupUnitAllOf) GetManufacturingDate() string`

GetManufacturingDate returns the ManufacturingDate field if non-nil, zero value otherwise.

### GetManufacturingDateOk

`func (o *StorageBatteryBackupUnitAllOf) GetManufacturingDateOk() (*string, bool)`

GetManufacturingDateOk returns a tuple with the ManufacturingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturingDate

`func (o *StorageBatteryBackupUnitAllOf) SetManufacturingDate(v string)`

SetManufacturingDate sets ManufacturingDate field to given value.

### HasManufacturingDate

`func (o *StorageBatteryBackupUnitAllOf) HasManufacturingDate() bool`

HasManufacturingDate returns a boolean if a field has been set.

### GetModuleVersion

`func (o *StorageBatteryBackupUnitAllOf) GetModuleVersion() string`

GetModuleVersion returns the ModuleVersion field if non-nil, zero value otherwise.

### GetModuleVersionOk

`func (o *StorageBatteryBackupUnitAllOf) GetModuleVersionOk() (*string, bool)`

GetModuleVersionOk returns a tuple with the ModuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleVersion

`func (o *StorageBatteryBackupUnitAllOf) SetModuleVersion(v string)`

SetModuleVersion sets ModuleVersion field to given value.

### HasModuleVersion

`func (o *StorageBatteryBackupUnitAllOf) HasModuleVersion() bool`

HasModuleVersion returns a boolean if a field has been set.

### GetNextLearnCycleTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) GetNextLearnCycleTimeStamp() string`

GetNextLearnCycleTimeStamp returns the NextLearnCycleTimeStamp field if non-nil, zero value otherwise.

### GetNextLearnCycleTimeStampOk

`func (o *StorageBatteryBackupUnitAllOf) GetNextLearnCycleTimeStampOk() (*string, bool)`

GetNextLearnCycleTimeStampOk returns a tuple with the NextLearnCycleTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLearnCycleTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) SetNextLearnCycleTimeStamp(v string)`

SetNextLearnCycleTimeStamp sets NextLearnCycleTimeStamp field to given value.

### HasNextLearnCycleTimeStamp

`func (o *StorageBatteryBackupUnitAllOf) HasNextLearnCycleTimeStamp() bool`

HasNextLearnCycleTimeStamp returns a boolean if a field has been set.

### GetPackEnergyInJoules

`func (o *StorageBatteryBackupUnitAllOf) GetPackEnergyInJoules() string`

GetPackEnergyInJoules returns the PackEnergyInJoules field if non-nil, zero value otherwise.

### GetPackEnergyInJoulesOk

`func (o *StorageBatteryBackupUnitAllOf) GetPackEnergyInJoulesOk() (*string, bool)`

GetPackEnergyInJoulesOk returns a tuple with the PackEnergyInJoules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackEnergyInJoules

`func (o *StorageBatteryBackupUnitAllOf) SetPackEnergyInJoules(v string)`

SetPackEnergyInJoules sets PackEnergyInJoules field to given value.

### HasPackEnergyInJoules

`func (o *StorageBatteryBackupUnitAllOf) HasPackEnergyInJoules() bool`

HasPackEnergyInJoules returns a boolean if a field has been set.

### GetRemainingPoolSpaceInPercent

`func (o *StorageBatteryBackupUnitAllOf) GetRemainingPoolSpaceInPercent() int64`

GetRemainingPoolSpaceInPercent returns the RemainingPoolSpaceInPercent field if non-nil, zero value otherwise.

### GetRemainingPoolSpaceInPercentOk

`func (o *StorageBatteryBackupUnitAllOf) GetRemainingPoolSpaceInPercentOk() (*int64, bool)`

GetRemainingPoolSpaceInPercentOk returns a tuple with the RemainingPoolSpaceInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingPoolSpaceInPercent

`func (o *StorageBatteryBackupUnitAllOf) SetRemainingPoolSpaceInPercent(v int64)`

SetRemainingPoolSpaceInPercent sets RemainingPoolSpaceInPercent field to given value.

### HasRemainingPoolSpaceInPercent

`func (o *StorageBatteryBackupUnitAllOf) HasRemainingPoolSpaceInPercent() bool`

HasRemainingPoolSpaceInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBatteryBackupUnitAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBatteryBackupUnitAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBatteryBackupUnitAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBatteryBackupUnitAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemperatureInCel

`func (o *StorageBatteryBackupUnitAllOf) GetTemperatureInCel() int64`

GetTemperatureInCel returns the TemperatureInCel field if non-nil, zero value otherwise.

### GetTemperatureInCelOk

`func (o *StorageBatteryBackupUnitAllOf) GetTemperatureInCelOk() (*int64, bool)`

GetTemperatureInCelOk returns a tuple with the TemperatureInCel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureInCel

`func (o *StorageBatteryBackupUnitAllOf) SetTemperatureInCel(v int64)`

SetTemperatureInCel sets TemperatureInCel field to given value.

### HasTemperatureInCel

`func (o *StorageBatteryBackupUnitAllOf) HasTemperatureInCel() bool`

HasTemperatureInCel returns a boolean if a field has been set.

### GetType

`func (o *StorageBatteryBackupUnitAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageBatteryBackupUnitAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageBatteryBackupUnitAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageBatteryBackupUnitAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVoltageInVolts

`func (o *StorageBatteryBackupUnitAllOf) GetVoltageInVolts() string`

GetVoltageInVolts returns the VoltageInVolts field if non-nil, zero value otherwise.

### GetVoltageInVoltsOk

`func (o *StorageBatteryBackupUnitAllOf) GetVoltageInVoltsOk() (*string, bool)`

GetVoltageInVoltsOk returns a tuple with the VoltageInVolts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltageInVolts

`func (o *StorageBatteryBackupUnitAllOf) SetVoltageInVolts(v string)`

SetVoltageInVolts sets VoltageInVolts field to given value.

### HasVoltageInVolts

`func (o *StorageBatteryBackupUnitAllOf) HasVoltageInVolts() bool`

HasVoltageInVolts returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageBatteryBackupUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageBatteryBackupUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageBatteryBackupUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageBatteryBackupUnitAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageBatteryBackupUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageBatteryBackupUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageBatteryBackupUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageBatteryBackupUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageBatteryBackupUnitAllOf) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageBatteryBackupUnitAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageBatteryBackupUnitAllOf) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageBatteryBackupUnitAllOf) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


