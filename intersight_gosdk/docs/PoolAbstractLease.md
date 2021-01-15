# PoolAbstractLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AllocationType** | Pointer to **string** | Type of the lease allocation either static or dynamic (i.e via pool). * &#x60;dynamic&#x60; - Identifiers to be allocated by system. * &#x60;static&#x60; - Identifiers are assigned by the user. | [optional] [default to "dynamic"]

## Methods

### NewPoolAbstractLease

`func NewPoolAbstractLease(classId string, objectType string, ) *PoolAbstractLease`

NewPoolAbstractLease instantiates a new PoolAbstractLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractLeaseWithDefaults

`func NewPoolAbstractLeaseWithDefaults() *PoolAbstractLease`

NewPoolAbstractLeaseWithDefaults instantiates a new PoolAbstractLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolAbstractLease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolAbstractLease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolAbstractLease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolAbstractLease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolAbstractLease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolAbstractLease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocationType

`func (o *PoolAbstractLease) GetAllocationType() string`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *PoolAbstractLease) GetAllocationTypeOk() (*string, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *PoolAbstractLease) SetAllocationType(v string)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *PoolAbstractLease) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


