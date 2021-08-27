# MacpoolIdBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "macpool.IdBlock"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "macpool.IdBlock"]
**MacBlock** | Pointer to [**MacpoolBlock**](MacpoolBlock.md) |  | [optional] 
**Pool** | Pointer to [**MacpoolPoolRelationship**](MacpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewMacpoolIdBlockAllOf

`func NewMacpoolIdBlockAllOf(classId string, objectType string, ) *MacpoolIdBlockAllOf`

NewMacpoolIdBlockAllOf instantiates a new MacpoolIdBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolIdBlockAllOfWithDefaults

`func NewMacpoolIdBlockAllOfWithDefaults() *MacpoolIdBlockAllOf`

NewMacpoolIdBlockAllOfWithDefaults instantiates a new MacpoolIdBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MacpoolIdBlockAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MacpoolIdBlockAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MacpoolIdBlockAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MacpoolIdBlockAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MacpoolIdBlockAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MacpoolIdBlockAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMacBlock

`func (o *MacpoolIdBlockAllOf) GetMacBlock() MacpoolBlock`

GetMacBlock returns the MacBlock field if non-nil, zero value otherwise.

### GetMacBlockOk

`func (o *MacpoolIdBlockAllOf) GetMacBlockOk() (*MacpoolBlock, bool)`

GetMacBlockOk returns a tuple with the MacBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlock

`func (o *MacpoolIdBlockAllOf) SetMacBlock(v MacpoolBlock)`

SetMacBlock sets MacBlock field to given value.

### HasMacBlock

`func (o *MacpoolIdBlockAllOf) HasMacBlock() bool`

HasMacBlock returns a boolean if a field has been set.

### GetPool

`func (o *MacpoolIdBlockAllOf) GetPool() MacpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *MacpoolIdBlockAllOf) GetPoolOk() (*MacpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *MacpoolIdBlockAllOf) SetPool(v MacpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *MacpoolIdBlockAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


