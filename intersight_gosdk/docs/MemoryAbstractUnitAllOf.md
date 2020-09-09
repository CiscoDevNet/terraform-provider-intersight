# MemoryAbstractUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** | This represents the administrative state of the memory unit on a server. | [optional] [readonly] 
**ArrayId** | Pointer to **int64** | This represents the memory array to which the memory unit belongs to. | [optional] [readonly] 
**Bank** | Pointer to **int64** | This represents the memory bank of the memory unit on a server. | [optional] [readonly] 
**Capacity** | Pointer to **string** | This represents the memory capacity in MiB of the memory unit on a server. | [optional] [readonly] 
**Clock** | Pointer to **string** | This represents the clock of the memory unit on a server. | [optional] [readonly] 
**FormFactor** | Pointer to **string** | This represents the form factor of the memory unit on a server. | [optional] [readonly] 
**Latency** | Pointer to **string** | This represents the latency of the memory unit on a server. | [optional] [readonly] 
**Location** | Pointer to **string** | This represents the location of the memory unit on a server. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | This represents the operational power state of the memory unit on a server. | [optional] [readonly] 
**OperState** | Pointer to **string** | This represents the operational state of the memory unit on a server. | [optional] [readonly] 
**Operability** | Pointer to **string** | This represents the operability of the memory unit on a server. | [optional] [readonly] 
**Presence** | Pointer to **string** | This represents the presence state of the memory unit on a server. | [optional] [readonly] 
**Set** | Pointer to **int64** | This represents the set of the memory unit on a server. | [optional] [readonly] 
**Speed** | Pointer to **string** | This represents the speed of the memory unit on a server. | [optional] [readonly] 
**Thermal** | Pointer to **string** | This represents the thremal state of the memory unit on a server. | [optional] [readonly] 
**Type** | Pointer to **string** | This represents the memory type of the memory unit on a server. | [optional] [readonly] 
**Visibility** | Pointer to **string** | This represents the visibility of the memory unit on a server. | [optional] [readonly] 
**Width** | Pointer to **string** | This represents the width of the memory unit on a server. | [optional] [readonly] 

## Methods

### NewMemoryAbstractUnitAllOf

`func NewMemoryAbstractUnitAllOf() *MemoryAbstractUnitAllOf`

NewMemoryAbstractUnitAllOf instantiates a new MemoryAbstractUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryAbstractUnitAllOfWithDefaults

`func NewMemoryAbstractUnitAllOfWithDefaults() *MemoryAbstractUnitAllOf`

NewMemoryAbstractUnitAllOfWithDefaults instantiates a new MemoryAbstractUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *MemoryAbstractUnitAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *MemoryAbstractUnitAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *MemoryAbstractUnitAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *MemoryAbstractUnitAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetArrayId

`func (o *MemoryAbstractUnitAllOf) GetArrayId() int64`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *MemoryAbstractUnitAllOf) GetArrayIdOk() (*int64, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *MemoryAbstractUnitAllOf) SetArrayId(v int64)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *MemoryAbstractUnitAllOf) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetBank

`func (o *MemoryAbstractUnitAllOf) GetBank() int64`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *MemoryAbstractUnitAllOf) GetBankOk() (*int64, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *MemoryAbstractUnitAllOf) SetBank(v int64)`

SetBank sets Bank field to given value.

### HasBank

`func (o *MemoryAbstractUnitAllOf) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCapacity

`func (o *MemoryAbstractUnitAllOf) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MemoryAbstractUnitAllOf) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MemoryAbstractUnitAllOf) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MemoryAbstractUnitAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetClock

`func (o *MemoryAbstractUnitAllOf) GetClock() string`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *MemoryAbstractUnitAllOf) GetClockOk() (*string, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *MemoryAbstractUnitAllOf) SetClock(v string)`

SetClock sets Clock field to given value.

### HasClock

`func (o *MemoryAbstractUnitAllOf) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetFormFactor

`func (o *MemoryAbstractUnitAllOf) GetFormFactor() string`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *MemoryAbstractUnitAllOf) GetFormFactorOk() (*string, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *MemoryAbstractUnitAllOf) SetFormFactor(v string)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *MemoryAbstractUnitAllOf) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetLatency

`func (o *MemoryAbstractUnitAllOf) GetLatency() string`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *MemoryAbstractUnitAllOf) GetLatencyOk() (*string, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *MemoryAbstractUnitAllOf) SetLatency(v string)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *MemoryAbstractUnitAllOf) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLocation

`func (o *MemoryAbstractUnitAllOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MemoryAbstractUnitAllOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MemoryAbstractUnitAllOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MemoryAbstractUnitAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOperPowerState

`func (o *MemoryAbstractUnitAllOf) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *MemoryAbstractUnitAllOf) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *MemoryAbstractUnitAllOf) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *MemoryAbstractUnitAllOf) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperState

`func (o *MemoryAbstractUnitAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *MemoryAbstractUnitAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *MemoryAbstractUnitAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *MemoryAbstractUnitAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *MemoryAbstractUnitAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *MemoryAbstractUnitAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *MemoryAbstractUnitAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *MemoryAbstractUnitAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPresence

`func (o *MemoryAbstractUnitAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *MemoryAbstractUnitAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *MemoryAbstractUnitAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *MemoryAbstractUnitAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSet

`func (o *MemoryAbstractUnitAllOf) GetSet() int64`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *MemoryAbstractUnitAllOf) GetSetOk() (*int64, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *MemoryAbstractUnitAllOf) SetSet(v int64)`

SetSet sets Set field to given value.

### HasSet

`func (o *MemoryAbstractUnitAllOf) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetSpeed

`func (o *MemoryAbstractUnitAllOf) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MemoryAbstractUnitAllOf) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MemoryAbstractUnitAllOf) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MemoryAbstractUnitAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetThermal

`func (o *MemoryAbstractUnitAllOf) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *MemoryAbstractUnitAllOf) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *MemoryAbstractUnitAllOf) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *MemoryAbstractUnitAllOf) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetType

`func (o *MemoryAbstractUnitAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemoryAbstractUnitAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemoryAbstractUnitAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemoryAbstractUnitAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *MemoryAbstractUnitAllOf) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MemoryAbstractUnitAllOf) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MemoryAbstractUnitAllOf) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MemoryAbstractUnitAllOf) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetWidth

`func (o *MemoryAbstractUnitAllOf) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MemoryAbstractUnitAllOf) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MemoryAbstractUnitAllOf) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MemoryAbstractUnitAllOf) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


