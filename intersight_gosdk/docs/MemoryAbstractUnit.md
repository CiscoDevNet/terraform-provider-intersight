# MemoryAbstractUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AdminState** | Pointer to **string** | This represents the administrative state of the memory unit on a server. | [optional] [readonly] 
**ArrayId** | Pointer to **int64** | This represents the memory array to which the memory unit belongs to. | [optional] [readonly] 
**Bank** | Pointer to **int64** | This represents the memory bank of the memory unit on a server. | [optional] [readonly] 
**Capacity** | Pointer to **string** | This represents the memory capacity in MiB of the memory unit on a server. | [optional] [readonly] 
**Clock** | Pointer to **string** | This represents the clock of the memory unit on a server. | [optional] [readonly] 
**FormFactor** | Pointer to **string** | This represents the form factor of the memory unit on a server. | [optional] [readonly] 
**Latency** | Pointer to **string** | This represents the latency of the memory unit on a server. | [optional] [readonly] 
**Location** | Pointer to **string** | This represents the location of the memory unit on a server. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | This represents the operational power state of the memory unit on a server. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | This represents the operational state of the memory unit on a server. | [optional] [readonly] 
**Operability** | Pointer to **string** | This represents the operability of the memory unit on a server. | [optional] [readonly] 
**Set** | Pointer to **int64** | This represents the set of the memory unit on a server. | [optional] [readonly] 
**Speed** | Pointer to **string** | This represents the speed of the memory unit on a server. | [optional] [readonly] 
**Thermal** | Pointer to **string** | This represents the thremal state of the memory unit on a server. | [optional] [readonly] 
**Type** | Pointer to **string** | This represents the memory type of the memory unit on a server. | [optional] [readonly] 
**Visibility** | Pointer to **string** | This represents the visibility of the memory unit on a server. | [optional] [readonly] 
**Width** | Pointer to **string** | This represents the width of the memory unit on a server. | [optional] [readonly] 

## Methods

### NewMemoryAbstractUnit

`func NewMemoryAbstractUnit(classId string, objectType string, ) *MemoryAbstractUnit`

NewMemoryAbstractUnit instantiates a new MemoryAbstractUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryAbstractUnitWithDefaults

`func NewMemoryAbstractUnitWithDefaults() *MemoryAbstractUnit`

NewMemoryAbstractUnitWithDefaults instantiates a new MemoryAbstractUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryAbstractUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryAbstractUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryAbstractUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryAbstractUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryAbstractUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryAbstractUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *MemoryAbstractUnit) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *MemoryAbstractUnit) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *MemoryAbstractUnit) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *MemoryAbstractUnit) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetArrayId

`func (o *MemoryAbstractUnit) GetArrayId() int64`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *MemoryAbstractUnit) GetArrayIdOk() (*int64, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *MemoryAbstractUnit) SetArrayId(v int64)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *MemoryAbstractUnit) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetBank

`func (o *MemoryAbstractUnit) GetBank() int64`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *MemoryAbstractUnit) GetBankOk() (*int64, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *MemoryAbstractUnit) SetBank(v int64)`

SetBank sets Bank field to given value.

### HasBank

`func (o *MemoryAbstractUnit) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCapacity

`func (o *MemoryAbstractUnit) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MemoryAbstractUnit) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MemoryAbstractUnit) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MemoryAbstractUnit) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetClock

`func (o *MemoryAbstractUnit) GetClock() string`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *MemoryAbstractUnit) GetClockOk() (*string, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *MemoryAbstractUnit) SetClock(v string)`

SetClock sets Clock field to given value.

### HasClock

`func (o *MemoryAbstractUnit) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetFormFactor

`func (o *MemoryAbstractUnit) GetFormFactor() string`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *MemoryAbstractUnit) GetFormFactorOk() (*string, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *MemoryAbstractUnit) SetFormFactor(v string)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *MemoryAbstractUnit) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetLatency

`func (o *MemoryAbstractUnit) GetLatency() string`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *MemoryAbstractUnit) GetLatencyOk() (*string, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *MemoryAbstractUnit) SetLatency(v string)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *MemoryAbstractUnit) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLocation

`func (o *MemoryAbstractUnit) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MemoryAbstractUnit) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MemoryAbstractUnit) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MemoryAbstractUnit) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOperPowerState

`func (o *MemoryAbstractUnit) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *MemoryAbstractUnit) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *MemoryAbstractUnit) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *MemoryAbstractUnit) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperReason

`func (o *MemoryAbstractUnit) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *MemoryAbstractUnit) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *MemoryAbstractUnit) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *MemoryAbstractUnit) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *MemoryAbstractUnit) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *MemoryAbstractUnit) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *MemoryAbstractUnit) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *MemoryAbstractUnit) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *MemoryAbstractUnit) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *MemoryAbstractUnit) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *MemoryAbstractUnit) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *MemoryAbstractUnit) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *MemoryAbstractUnit) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *MemoryAbstractUnit) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetSet

`func (o *MemoryAbstractUnit) GetSet() int64`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *MemoryAbstractUnit) GetSetOk() (*int64, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *MemoryAbstractUnit) SetSet(v int64)`

SetSet sets Set field to given value.

### HasSet

`func (o *MemoryAbstractUnit) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetSpeed

`func (o *MemoryAbstractUnit) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MemoryAbstractUnit) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MemoryAbstractUnit) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MemoryAbstractUnit) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetThermal

`func (o *MemoryAbstractUnit) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *MemoryAbstractUnit) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *MemoryAbstractUnit) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *MemoryAbstractUnit) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetType

`func (o *MemoryAbstractUnit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemoryAbstractUnit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemoryAbstractUnit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemoryAbstractUnit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *MemoryAbstractUnit) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MemoryAbstractUnit) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MemoryAbstractUnit) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MemoryAbstractUnit) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetWidth

`func (o *MemoryAbstractUnit) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MemoryAbstractUnit) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MemoryAbstractUnit) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MemoryAbstractUnit) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


