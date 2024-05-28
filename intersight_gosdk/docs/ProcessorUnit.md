# ProcessorUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "processor.Unit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "processor.Unit"]
**Architecture** | Pointer to **string** | The architecture of the installed processor. | [optional] [readonly] 
**Description** | Pointer to **string** | This field displays the description of the processor. | [optional] [readonly] 
**IsPlatformSupported** | Pointer to **bool** | This field indicates whether the processor is supported on the server or not. | [optional] [readonly] [default to true]
**NumCores** | Pointer to **int64** | The number of cores present in a given processor. | [optional] [readonly] 
**NumCoresEnabled** | Pointer to **string** | The number of enabled cores in the installed processor. | [optional] [readonly] 
**NumThreads** | Pointer to **string** | The maximum number of threads available in the installed processor. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state of the processor. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | The health indicator of the processor, &#39;OK&#39; indicates the processor is operatinal. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of the central processing unit. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field displays the part number of the of the processor. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field displays the product ID of the processor. | [optional] [readonly] 
**ProcessorId** | Pointer to **int64** | The ID number of a given processor. | [optional] [readonly] 
**SocketDesignation** | Pointer to **string** | The socket ID of the installed processor. | [optional] [readonly] 
**Speed** | Pointer to **float32** | The maximum speed of the installed processor in GHz. | [optional] [readonly] 
**Stepping** | Pointer to **string** | The CPU stepping of the installed processor. | [optional] [readonly] 
**Thermal** | Pointer to **string** | The temperature of the processor in centigrade. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**NullableComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewProcessorUnit

`func NewProcessorUnit(classId string, objectType string, ) *ProcessorUnit`

NewProcessorUnit instantiates a new ProcessorUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorUnitWithDefaults

`func NewProcessorUnitWithDefaults() *ProcessorUnit`

NewProcessorUnitWithDefaults instantiates a new ProcessorUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ProcessorUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ProcessorUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ProcessorUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ProcessorUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProcessorUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProcessorUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetDescription

`func (o *ProcessorUnit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProcessorUnit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProcessorUnit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProcessorUnit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPlatformSupported

`func (o *ProcessorUnit) GetIsPlatformSupported() bool`

GetIsPlatformSupported returns the IsPlatformSupported field if non-nil, zero value otherwise.

### GetIsPlatformSupportedOk

`func (o *ProcessorUnit) GetIsPlatformSupportedOk() (*bool, bool)`

GetIsPlatformSupportedOk returns a tuple with the IsPlatformSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlatformSupported

`func (o *ProcessorUnit) SetIsPlatformSupported(v bool)`

SetIsPlatformSupported sets IsPlatformSupported field to given value.

### HasIsPlatformSupported

`func (o *ProcessorUnit) HasIsPlatformSupported() bool`

HasIsPlatformSupported returns a boolean if a field has been set.

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

### GetOperReason

`func (o *ProcessorUnit) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *ProcessorUnit) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *ProcessorUnit) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *ProcessorUnit) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *ProcessorUnit) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *ProcessorUnit) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
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

### GetPartNumber

`func (o *ProcessorUnit) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *ProcessorUnit) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *ProcessorUnit) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *ProcessorUnit) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *ProcessorUnit) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ProcessorUnit) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ProcessorUnit) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *ProcessorUnit) HasPid() bool`

HasPid returns a boolean if a field has been set.

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

### SetComputeBladeNil

`func (o *ProcessorUnit) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *ProcessorUnit) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
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

### SetComputeBoardNil

`func (o *ProcessorUnit) SetComputeBoardNil(b bool)`

 SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil

### UnsetComputeBoard
`func (o *ProcessorUnit) UnsetComputeBoard()`

UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
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

### SetComputeRackUnitNil

`func (o *ProcessorUnit) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *ProcessorUnit) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
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

### SetInventoryDeviceInfoNil

`func (o *ProcessorUnit) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *ProcessorUnit) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetNetworkElement

`func (o *ProcessorUnit) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *ProcessorUnit) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *ProcessorUnit) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *ProcessorUnit) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *ProcessorUnit) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *ProcessorUnit) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
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

### SetRegisteredDeviceNil

`func (o *ProcessorUnit) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ProcessorUnit) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


