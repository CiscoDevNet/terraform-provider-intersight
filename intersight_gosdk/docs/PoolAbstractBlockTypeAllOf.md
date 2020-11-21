# PoolAbstractBlockTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Size** | Pointer to **int64** | Number of identifiers this block can hold. | [optional] [readonly] 

## Methods

### NewPoolAbstractBlockTypeAllOf

`func NewPoolAbstractBlockTypeAllOf(classId string, objectType string, ) *PoolAbstractBlockTypeAllOf`

NewPoolAbstractBlockTypeAllOf instantiates a new PoolAbstractBlockTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractBlockTypeAllOfWithDefaults

`func NewPoolAbstractBlockTypeAllOfWithDefaults() *PoolAbstractBlockTypeAllOf`

NewPoolAbstractBlockTypeAllOfWithDefaults instantiates a new PoolAbstractBlockTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolAbstractBlockTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolAbstractBlockTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolAbstractBlockTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolAbstractBlockTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolAbstractBlockTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolAbstractBlockTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSize

`func (o *PoolAbstractBlockTypeAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PoolAbstractBlockTypeAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PoolAbstractBlockTypeAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PoolAbstractBlockTypeAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


