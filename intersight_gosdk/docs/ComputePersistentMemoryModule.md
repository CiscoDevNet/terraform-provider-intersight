# ComputePersistentMemoryModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.PersistentMemoryModule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.PersistentMemoryModule"]
**SocketId** | Pointer to **string** | Socket ID of the Persistent Memory Module on the server. | [optional] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the Persistent Memory Module on the server. | [optional] 

## Methods

### NewComputePersistentMemoryModule

`func NewComputePersistentMemoryModule(classId string, objectType string, ) *ComputePersistentMemoryModule`

NewComputePersistentMemoryModule instantiates a new ComputePersistentMemoryModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePersistentMemoryModuleWithDefaults

`func NewComputePersistentMemoryModuleWithDefaults() *ComputePersistentMemoryModule`

NewComputePersistentMemoryModuleWithDefaults instantiates a new ComputePersistentMemoryModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePersistentMemoryModule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePersistentMemoryModule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePersistentMemoryModule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePersistentMemoryModule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePersistentMemoryModule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePersistentMemoryModule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSocketId

`func (o *ComputePersistentMemoryModule) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *ComputePersistentMemoryModule) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *ComputePersistentMemoryModule) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *ComputePersistentMemoryModule) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *ComputePersistentMemoryModule) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *ComputePersistentMemoryModule) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *ComputePersistentMemoryModule) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *ComputePersistentMemoryModule) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


