# MemoryPersistentMemoryGoal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryModePercentage** | Pointer to **int32** | Volatile memory percentage. | [optional] 
**PersistentMemoryType** | Pointer to **string** | Type of the Persistent Memory configuration where the Persistent Memory Modules are combined in an interleaved set or not. * &#x60;app-direct&#x60; - The App Direct interleaved Persistent Memory type. * &#x60;app-direct-non-interleaved&#x60; - The App Direct non-interleaved Persistent Memory type. | [optional] [default to "app-direct"]
**SocketId** | Pointer to **string** | CPU Socket ID to which this goal will be applied. * &#x60;All Sockets&#x60; - All the CPU socket IDs in a server. | [optional] [default to "All Sockets"]

## Methods

### NewMemoryPersistentMemoryGoal

`func NewMemoryPersistentMemoryGoal() *MemoryPersistentMemoryGoal`

NewMemoryPersistentMemoryGoal instantiates a new MemoryPersistentMemoryGoal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryGoalWithDefaults

`func NewMemoryPersistentMemoryGoalWithDefaults() *MemoryPersistentMemoryGoal`

NewMemoryPersistentMemoryGoalWithDefaults instantiates a new MemoryPersistentMemoryGoal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryModePercentage

`func (o *MemoryPersistentMemoryGoal) GetMemoryModePercentage() int32`

GetMemoryModePercentage returns the MemoryModePercentage field if non-nil, zero value otherwise.

### GetMemoryModePercentageOk

`func (o *MemoryPersistentMemoryGoal) GetMemoryModePercentageOk() (*int32, bool)`

GetMemoryModePercentageOk returns a tuple with the MemoryModePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryModePercentage

`func (o *MemoryPersistentMemoryGoal) SetMemoryModePercentage(v int32)`

SetMemoryModePercentage sets MemoryModePercentage field to given value.

### HasMemoryModePercentage

`func (o *MemoryPersistentMemoryGoal) HasMemoryModePercentage() bool`

HasMemoryModePercentage returns a boolean if a field has been set.

### GetPersistentMemoryType

`func (o *MemoryPersistentMemoryGoal) GetPersistentMemoryType() string`

GetPersistentMemoryType returns the PersistentMemoryType field if non-nil, zero value otherwise.

### GetPersistentMemoryTypeOk

`func (o *MemoryPersistentMemoryGoal) GetPersistentMemoryTypeOk() (*string, bool)`

GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryType

`func (o *MemoryPersistentMemoryGoal) SetPersistentMemoryType(v string)`

SetPersistentMemoryType sets PersistentMemoryType field to given value.

### HasPersistentMemoryType

`func (o *MemoryPersistentMemoryGoal) HasPersistentMemoryType() bool`

HasPersistentMemoryType returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryGoal) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryGoal) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryGoal) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryGoal) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


