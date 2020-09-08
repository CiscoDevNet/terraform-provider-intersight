# ProcessorUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | The architecture of the installed processor. | [optional] [readonly] 
**NumCores** | Pointer to **int64** | The number of cores present in a given processor. | [optional] [readonly] 
**NumCoresEnabled** | Pointer to **string** | The number of enabled cores in the installed processor. | [optional] [readonly] 
**NumThreads** | Pointer to **string** | The maximum number of threads available in the installed processor. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state of the processor. | [optional] [readonly] 
**OperState** | Pointer to **string** | The health indicator of the processor, &#39;OK&#39; indicates the processor is operatinal. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of the central processing unit. | [optional] [readonly] 
**Presence** | Pointer to **string** | The valid values are &#39;equipped&#39; and &#39;absent&#39;. | [optional] [readonly] 
**ProcessorId** | Pointer to **int64** | The ID number of a given processor. | [optional] [readonly] 
**SocketDesignation** | Pointer to **string** | The socket ID of the installed processor. | [optional] [readonly] 
**Speed** | Pointer to **float32** | The maximum speed of the installed processor in GHz. | [optional] [readonly] 
**Stepping** | Pointer to **string** | The CPU stepping of the installed processor. | [optional] [readonly] 
**Thermal** | Pointer to **string** | The temperature of the processor in centigrade. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewProcessorUnit

`func NewProcessorUnit() *ProcessorUnit`

NewProcessorUnit instantiates a new ProcessorUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorUnitWithDefaults

`func NewProcessorUnitWithDefaults() *ProcessorUnit`

NewProcessorUnitWithDefaults instantiates a new ProcessorUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ProcessorUnit) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ProcessorUnit) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ProcessorUnit) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ProcessorUnit) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetNumCores

`func (o *ProcessorUnit) GetNumCores() int64`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *ProcessorUnit) GetNumCoresOk() (*int64, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *ProcessorUnit) SetNumCores(v int64)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *ProcessorUnit) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.

### GetNumCoresEnabled

`func (o *ProcessorUnit) GetNumCoresEnabled() string`

GetNumCoresEnabled returns the NumCoresEnabled field if non-nil, zero value otherwise.

### GetNumCoresEnabledOk

`func (o *ProcessorUnit) GetNumCoresEnabledOk() (*string, bool)`

GetNumCoresEnabledOk returns a tuple with the NumCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCoresEnabled

`func (o *ProcessorUnit) SetNumCoresEnabled(v string)`

SetNumCoresEnabled sets NumCoresEnabled field to given value.

### HasNumCoresEnabled

`func (o *ProcessorUnit) HasNumCoresEnabled() bool`

HasNumCoresEnabled returns a boolean if a field has been set.

### GetNumThreads

`func (o *ProcessorUnit) GetNumThreads() string`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ProcessorUnit) GetNumThreadsOk() (*string, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ProcessorUnit) SetNumThreads(v string)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ProcessorUnit) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ProcessorUnit) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ProcessorUnit) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ProcessorUnit) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ProcessorUnit) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperState

`func (o *ProcessorUnit) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ProcessorUnit) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ProcessorUnit) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ProcessorUnit) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ProcessorUnit) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ProcessorUnit) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ProcessorUnit) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ProcessorUnit) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPresence

`func (o *ProcessorUnit) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ProcessorUnit) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ProcessorUnit) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ProcessorUnit) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetProcessorId

`func (o *ProcessorUnit) GetProcessorId() int64`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *ProcessorUnit) GetProcessorIdOk() (*int64, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *ProcessorUnit) SetProcessorId(v int64)`

SetProcessorId sets ProcessorId field to given value.

### HasProcessorId

`func (o *ProcessorUnit) HasProcessorId() bool`

HasProcessorId returns a boolean if a field has been set.

### GetSocketDesignation

`func (o *ProcessorUnit) GetSocketDesignation() string`

GetSocketDesignation returns the SocketDesignation field if non-nil, zero value otherwise.

### GetSocketDesignationOk

`func (o *ProcessorUnit) GetSocketDesignationOk() (*string, bool)`

GetSocketDesignationOk returns a tuple with the SocketDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketDesignation

`func (o *ProcessorUnit) SetSocketDesignation(v string)`

SetSocketDesignation sets SocketDesignation field to given value.

### HasSocketDesignation

`func (o *ProcessorUnit) HasSocketDesignation() bool`

HasSocketDesignation returns a boolean if a field has been set.

### GetSpeed

`func (o *ProcessorUnit) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ProcessorUnit) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ProcessorUnit) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ProcessorUnit) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStepping

`func (o *ProcessorUnit) GetStepping() string`

GetStepping returns the Stepping field if non-nil, zero value otherwise.

### GetSteppingOk

`func (o *ProcessorUnit) GetSteppingOk() (*string, bool)`

GetSteppingOk returns a tuple with the Stepping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepping

`func (o *ProcessorUnit) SetStepping(v string)`

SetStepping sets Stepping field to given value.

### HasStepping

`func (o *ProcessorUnit) HasStepping() bool`

HasStepping returns a boolean if a field has been set.

### GetThermal

`func (o *ProcessorUnit) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *ProcessorUnit) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *ProcessorUnit) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *ProcessorUnit) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ProcessorUnit) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ProcessorUnit) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ProcessorUnit) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ProcessorUnit) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *ProcessorUnit) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *ProcessorUnit) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *ProcessorUnit) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *ProcessorUnit) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ProcessorUnit) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ProcessorUnit) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ProcessorUnit) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ProcessorUnit) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ProcessorUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ProcessorUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ProcessorUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ProcessorUnit) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ProcessorUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ProcessorUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ProcessorUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ProcessorUnit) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


