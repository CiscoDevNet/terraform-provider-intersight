# PoolAbstractLeaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AllocationType** | Pointer to **string** | Type of the lease allocation either static or dynamic (i.e via pool). * &#x60;dynamic&#x60; - Identifiers to be allocated by system. * &#x60;static&#x60; - Identifiers are assigned by the user. | [optional] [default to "dynamic"]

## Methods

### NewPoolAbstractLeaseAllOf

`func NewPoolAbstractLeaseAllOf(classId string, objectType string, ) *PoolAbstractLeaseAllOf`

NewPoolAbstractLeaseAllOf instantiates a new PoolAbstractLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractLeaseAllOfWithDefaults

`func NewPoolAbstractLeaseAllOfWithDefaults() *PoolAbstractLeaseAllOf`

NewPoolAbstractLeaseAllOfWithDefaults instantiates a new PoolAbstractLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolAbstractLeaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolAbstractLeaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolAbstractLeaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolAbstractLeaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolAbstractLeaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolAbstractLeaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocationType

`func (o *PoolAbstractLeaseAllOf) GetAllocationType() string`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *PoolAbstractLeaseAllOf) GetAllocationTypeOk() (*string, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *PoolAbstractLeaseAllOf) SetAllocationType(v string)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *PoolAbstractLeaseAllOf) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


