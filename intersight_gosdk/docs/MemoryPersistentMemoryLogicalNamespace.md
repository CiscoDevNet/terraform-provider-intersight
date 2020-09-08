# MemoryPersistentMemoryLogicalNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** | Capacity of this Namespace that is created or modified. | [optional] 
**Mode** | Pointer to **string** | Mode of this Namespace that is created or modified. * &#x60;raw&#x60; - The raw mode of Persistent Memory Namespace. * &#x60;block&#x60; - The block mode of Persistent Memory Namespace. | [optional] [default to "raw"]
**Name** | Pointer to **string** | Name of this Namespace to be created on the server. | [optional] 
**SocketId** | Pointer to **int32** | Socket ID of the region on which this Namespace has to be created or modified. * &#x60;1&#x60; - The first CPU socket in a server. * &#x60;2&#x60; - The second CPU socket in a server. * &#x60;3&#x60; - The third CPU socket in a server. * &#x60;4&#x60; - The fourth CPU socket in a server. | [optional] [default to 1]
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the region on which this Namespace has to be created or modified. * &#x60;Not Applicable&#x60; - The socket memory ID is not applicable if app-direct persistent memory type is selected in the goal. * &#x60;2&#x60; - The second socket memory ID within a socket in a server. * &#x60;4&#x60; - The fourth socket memory ID within a socket in a server. * &#x60;6&#x60; - The sixth socket memory ID within a socket in a server. * &#x60;8&#x60; - The eighth socket memory ID within a socket in a server. * &#x60;10&#x60; - The tenth socket memory ID within a socket in a server. * &#x60;12&#x60; - The twelfth socket memory ID within a socket in a server. | [optional] [default to "Not Applicable"]

## Methods

### NewMemoryPersistentMemoryLogicalNamespace

`func NewMemoryPersistentMemoryLogicalNamespace() *MemoryPersistentMemoryLogicalNamespace`

NewMemoryPersistentMemoryLogicalNamespace instantiates a new MemoryPersistentMemoryLogicalNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryLogicalNamespaceWithDefaults

`func NewMemoryPersistentMemoryLogicalNamespaceWithDefaults() *MemoryPersistentMemoryLogicalNamespace`

NewMemoryPersistentMemoryLogicalNamespaceWithDefaults instantiates a new MemoryPersistentMemoryLogicalNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *MemoryPersistentMemoryLogicalNamespace) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MemoryPersistentMemoryLogicalNamespace) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MemoryPersistentMemoryLogicalNamespace) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MemoryPersistentMemoryLogicalNamespace) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMode

`func (o *MemoryPersistentMemoryLogicalNamespace) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MemoryPersistentMemoryLogicalNamespace) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MemoryPersistentMemoryLogicalNamespace) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MemoryPersistentMemoryLogicalNamespace) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *MemoryPersistentMemoryLogicalNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemoryPersistentMemoryLogicalNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemoryPersistentMemoryLogicalNamespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemoryPersistentMemoryLogicalNamespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryLogicalNamespace) GetSocketId() int32`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryLogicalNamespace) GetSocketIdOk() (*int32, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryLogicalNamespace) SetSocketId(v int32)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryLogicalNamespace) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryLogicalNamespace) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryLogicalNamespace) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryLogicalNamespace) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryLogicalNamespace) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


