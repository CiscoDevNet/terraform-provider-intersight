# UuidpoolBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "uuidpool.Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "uuidpool.Block"]
**UuidSuffixBlock** | Pointer to [**UuidpoolUuidBlock**](UuidpoolUuidBlock.md) |  | [optional] 
**Pool** | Pointer to [**UuidpoolPoolRelationship**](UuidpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewUuidpoolBlock

`func NewUuidpoolBlock(classId string, objectType string, ) *UuidpoolBlock`

NewUuidpoolBlock instantiates a new UuidpoolBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolBlockWithDefaults

`func NewUuidpoolBlockWithDefaults() *UuidpoolBlock`

NewUuidpoolBlockWithDefaults instantiates a new UuidpoolBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UuidpoolBlock) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UuidpoolBlock) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UuidpoolBlock) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UuidpoolBlock) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UuidpoolBlock) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UuidpoolBlock) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUuidSuffixBlock

`func (o *UuidpoolBlock) GetUuidSuffixBlock() UuidpoolUuidBlock`

GetUuidSuffixBlock returns the UuidSuffixBlock field if non-nil, zero value otherwise.

### GetUuidSuffixBlockOk

`func (o *UuidpoolBlock) GetUuidSuffixBlockOk() (*UuidpoolUuidBlock, bool)`

GetUuidSuffixBlockOk returns a tuple with the UuidSuffixBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidSuffixBlock

`func (o *UuidpoolBlock) SetUuidSuffixBlock(v UuidpoolUuidBlock)`

SetUuidSuffixBlock sets UuidSuffixBlock field to given value.

### HasUuidSuffixBlock

`func (o *UuidpoolBlock) HasUuidSuffixBlock() bool`

HasUuidSuffixBlock returns a boolean if a field has been set.

### GetPool

`func (o *UuidpoolBlock) GetPool() UuidpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UuidpoolBlock) GetPoolOk() (*UuidpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UuidpoolBlock) SetPool(v UuidpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UuidpoolBlock) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


