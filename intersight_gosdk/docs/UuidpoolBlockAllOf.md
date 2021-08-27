# UuidpoolBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "uuidpool.Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "uuidpool.Block"]
**UuidSuffixBlock** | Pointer to [**UuidpoolUuidBlock**](UuidpoolUuidBlock.md) |  | [optional] 
**Pool** | Pointer to [**UuidpoolPoolRelationship**](UuidpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewUuidpoolBlockAllOf

`func NewUuidpoolBlockAllOf(classId string, objectType string, ) *UuidpoolBlockAllOf`

NewUuidpoolBlockAllOf instantiates a new UuidpoolBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolBlockAllOfWithDefaults

`func NewUuidpoolBlockAllOfWithDefaults() *UuidpoolBlockAllOf`

NewUuidpoolBlockAllOfWithDefaults instantiates a new UuidpoolBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UuidpoolBlockAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UuidpoolBlockAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UuidpoolBlockAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UuidpoolBlockAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UuidpoolBlockAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UuidpoolBlockAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUuidSuffixBlock

`func (o *UuidpoolBlockAllOf) GetUuidSuffixBlock() UuidpoolUuidBlock`

GetUuidSuffixBlock returns the UuidSuffixBlock field if non-nil, zero value otherwise.

### GetUuidSuffixBlockOk

`func (o *UuidpoolBlockAllOf) GetUuidSuffixBlockOk() (*UuidpoolUuidBlock, bool)`

GetUuidSuffixBlockOk returns a tuple with the UuidSuffixBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidSuffixBlock

`func (o *UuidpoolBlockAllOf) SetUuidSuffixBlock(v UuidpoolUuidBlock)`

SetUuidSuffixBlock sets UuidSuffixBlock field to given value.

### HasUuidSuffixBlock

`func (o *UuidpoolBlockAllOf) HasUuidSuffixBlock() bool`

HasUuidSuffixBlock returns a boolean if a field has been set.

### GetPool

`func (o *UuidpoolBlockAllOf) GetPool() UuidpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UuidpoolBlockAllOf) GetPoolOk() (*UuidpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UuidpoolBlockAllOf) SetPool(v UuidpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UuidpoolBlockAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


