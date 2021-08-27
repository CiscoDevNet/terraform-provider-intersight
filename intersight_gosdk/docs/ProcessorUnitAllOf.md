# ProcessorUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "processor.Unit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "processor.Unit"]
**Architecture** | Pointer to **string** | The architecture of the installed processor. | [optional] [readonly] 
**NumCores** | Pointer to **int64** | The number of cores present in a given processor. | [optional] [readonly] 
**NumCoresEnabled** | Pointer to **string** | The number of enabled cores in the installed processor. | [optional] [readonly] 
**NumThreads** | Pointer to **string** | The maximum number of threads available in the installed processor. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state of the processor. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | The health indicator of the processor, &#39;OK&#39; indicates the processor is operatinal. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of the central processing unit. | [optional] [readonly] 
**ProcessorId** | Pointer to **int64** | The ID number of a given processor. | [optional] [readonly] 
**SocketDesignation** | Pointer to **string** | The socket ID of the installed processor. | [optional] [readonly] 
**Speed** | Pointer to **float32** | The maximum speed of the installed processor in GHz. | [optional] [readonly] 
**Stepping** | Pointer to **string** | The CPU stepping of the installed processor. | [optional] [readonly] 
**Thermal** | Pointer to **string** | The temperature of the processor in centigrade. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewProcessorUnitAllOf

`func NewProcessorUnitAllOf(classId string, objectType string, ) *ProcessorUnitAllOf`

NewProcessorUnitAllOf instantiates a new ProcessorUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorUnitAllOfWithDefaults

`func NewProcessorUnitAllOfWithDefaults() *ProcessorUnitAllOf`

NewProcessorUnitAllOfWithDefaults instantiates a new ProcessorUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ProcessorUnitAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ProcessorUnitAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ProcessorUnitAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ProcessorUnitAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProcessorUnitAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProcessorUnitAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArchitecture

`func (o *ProcessorUnitAllOf) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ProcessorUnitAllOf) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ProcessorUnitAllOf) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ProcessorUnitAllOf) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetNumCores

`func (o *ProcessorUnitAllOf) GetNumCores() int64`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *ProcessorUnitAllOf) GetNumCoresOk() (*int64, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *ProcessorUnitAllOf) SetNumCores(v int64)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *ProcessorUnitAllOf) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.

### GetNumCoresEnabled

`func (o *ProcessorUnitAllOf) GetNumCoresEnabled() string`

GetNumCoresEnabled returns the NumCoresEnabled field if non-nil, zero value otherwise.

### GetNumCoresEnabledOk

`func (o *ProcessorUnitAllOf) GetNumCoresEnabledOk() (*string, bool)`

GetNumCoresEnabledOk returns a tuple with the NumCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCoresEnabled

`func (o *ProcessorUnitAllOf) SetNumCoresEnabled(v string)`

SetNumCoresEnabled sets NumCoresEnabled field to given value.

### HasNumCoresEnabled

`func (o *ProcessorUnitAllOf) HasNumCoresEnabled() bool`

HasNumCoresEnabled returns a boolean if a field has been set.

### GetNumThreads

`func (o *ProcessorUnitAllOf) GetNumThreads() string`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ProcessorUnitAllOf) GetNumThreadsOk() (*string, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ProcessorUnitAllOf) SetNumThreads(v string)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ProcessorUnitAllOf) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ProcessorUnitAllOf) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ProcessorUnitAllOf) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ProcessorUnitAllOf) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ProcessorUnitAllOf) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperReason

`func (o *ProcessorUnitAllOf) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *ProcessorUnitAllOf) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *ProcessorUnitAllOf) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *ProcessorUnitAllOf) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *ProcessorUnitAllOf) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *ProcessorUnitAllOf) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *ProcessorUnitAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ProcessorUnitAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ProcessorUnitAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ProcessorUnitAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ProcessorUnitAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ProcessorUnitAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ProcessorUnitAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ProcessorUnitAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetProcessorId

`func (o *ProcessorUnitAllOf) GetProcessorId() int64`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *ProcessorUnitAllOf) GetProcessorIdOk() (*int64, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *ProcessorUnitAllOf) SetProcessorId(v int64)`

SetProcessorId sets ProcessorId field to given value.

### HasProcessorId

`func (o *ProcessorUnitAllOf) HasProcessorId() bool`

HasProcessorId returns a boolean if a field has been set.

### GetSocketDesignation

`func (o *ProcessorUnitAllOf) GetSocketDesignation() string`

GetSocketDesignation returns the SocketDesignation field if non-nil, zero value otherwise.

### GetSocketDesignationOk

`func (o *ProcessorUnitAllOf) GetSocketDesignationOk() (*string, bool)`

GetSocketDesignationOk returns a tuple with the SocketDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketDesignation

`func (o *ProcessorUnitAllOf) SetSocketDesignation(v string)`

SetSocketDesignation sets SocketDesignation field to given value.

### HasSocketDesignation

`func (o *ProcessorUnitAllOf) HasSocketDesignation() bool`

HasSocketDesignation returns a boolean if a field has been set.

### GetSpeed

`func (o *ProcessorUnitAllOf) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ProcessorUnitAllOf) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ProcessorUnitAllOf) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ProcessorUnitAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStepping

`func (o *ProcessorUnitAllOf) GetStepping() string`

GetStepping returns the Stepping field if non-nil, zero value otherwise.

### GetSteppingOk

`func (o *ProcessorUnitAllOf) GetSteppingOk() (*string, bool)`

GetSteppingOk returns a tuple with the Stepping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepping

`func (o *ProcessorUnitAllOf) SetStepping(v string)`

SetStepping sets Stepping field to given value.

### HasStepping

`func (o *ProcessorUnitAllOf) HasStepping() bool`

HasStepping returns a boolean if a field has been set.

### GetThermal

`func (o *ProcessorUnitAllOf) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *ProcessorUnitAllOf) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *ProcessorUnitAllOf) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *ProcessorUnitAllOf) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ProcessorUnitAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ProcessorUnitAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ProcessorUnitAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ProcessorUnitAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *ProcessorUnitAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *ProcessorUnitAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *ProcessorUnitAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *ProcessorUnitAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ProcessorUnitAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ProcessorUnitAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ProcessorUnitAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ProcessorUnitAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ProcessorUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ProcessorUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ProcessorUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ProcessorUnitAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ProcessorUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ProcessorUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ProcessorUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ProcessorUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


