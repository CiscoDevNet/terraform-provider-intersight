# FabricPfcWatchDog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.PfcWatchDog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.PfcWatchDog"]
**IsWatchdogEnabled** | Pointer to **bool** | Enables or disables the Priority-based Flow Control (PFC) watchdog feature. When enabled, the watchdog actively monitors PFC pause frames. By default Priority Flow Control is enabled for new QoS policies. Existing policies remain unaffected. | [optional] [default to false]
**ShutdownMultiplier** | Pointer to **int64** | The Shutdown Multiplier, multiplied by the watchdog timer, determines the total duration a Priority-based Flow Control (PFC)-enabled queue remains in shutdown mode. The maximum Watchdog Shutdown Multiplier is 10. However, if the Watchdog Interval exceeds 500 milliseconds, the Multiplier limit is reduced to 2. | [optional] [default to 1]
**WatchdogInterval** | Pointer to **int64** | Time in milliseconds for the PFC Watchdog. The Watchdog product (Watchdog Interval (in milliseconds) × Shutdown Multiplier) cannot exceed 1000ms. | [optional] [default to 500]

## Methods

### NewFabricPfcWatchDog

`func NewFabricPfcWatchDog(classId string, objectType string, ) *FabricPfcWatchDog`

NewFabricPfcWatchDog instantiates a new FabricPfcWatchDog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPfcWatchDogWithDefaults

`func NewFabricPfcWatchDogWithDefaults() *FabricPfcWatchDog`

NewFabricPfcWatchDogWithDefaults instantiates a new FabricPfcWatchDog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricPfcWatchDog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricPfcWatchDog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricPfcWatchDog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricPfcWatchDog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricPfcWatchDog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricPfcWatchDog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsWatchdogEnabled

`func (o *FabricPfcWatchDog) GetIsWatchdogEnabled() bool`

GetIsWatchdogEnabled returns the IsWatchdogEnabled field if non-nil, zero value otherwise.

### GetIsWatchdogEnabledOk

`func (o *FabricPfcWatchDog) GetIsWatchdogEnabledOk() (*bool, bool)`

GetIsWatchdogEnabledOk returns a tuple with the IsWatchdogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWatchdogEnabled

`func (o *FabricPfcWatchDog) SetIsWatchdogEnabled(v bool)`

SetIsWatchdogEnabled sets IsWatchdogEnabled field to given value.

### HasIsWatchdogEnabled

`func (o *FabricPfcWatchDog) HasIsWatchdogEnabled() bool`

HasIsWatchdogEnabled returns a boolean if a field has been set.

### GetShutdownMultiplier

`func (o *FabricPfcWatchDog) GetShutdownMultiplier() int64`

GetShutdownMultiplier returns the ShutdownMultiplier field if non-nil, zero value otherwise.

### GetShutdownMultiplierOk

`func (o *FabricPfcWatchDog) GetShutdownMultiplierOk() (*int64, bool)`

GetShutdownMultiplierOk returns a tuple with the ShutdownMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMultiplier

`func (o *FabricPfcWatchDog) SetShutdownMultiplier(v int64)`

SetShutdownMultiplier sets ShutdownMultiplier field to given value.

### HasShutdownMultiplier

`func (o *FabricPfcWatchDog) HasShutdownMultiplier() bool`

HasShutdownMultiplier returns a boolean if a field has been set.

### GetWatchdogInterval

`func (o *FabricPfcWatchDog) GetWatchdogInterval() int64`

GetWatchdogInterval returns the WatchdogInterval field if non-nil, zero value otherwise.

### GetWatchdogIntervalOk

`func (o *FabricPfcWatchDog) GetWatchdogIntervalOk() (*int64, bool)`

GetWatchdogIntervalOk returns a tuple with the WatchdogInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdogInterval

`func (o *FabricPfcWatchDog) SetWatchdogInterval(v int64)`

SetWatchdogInterval sets WatchdogInterval field to given value.

### HasWatchdogInterval

`func (o *FabricPfcWatchDog) HasWatchdogInterval() bool`

HasWatchdogInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


