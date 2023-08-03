# BulkMoRefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.MoRef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.MoRef"]
**Moid** | Pointer to **string** | Moid represents the MoId of the object. | [optional] 

## Methods

### NewBulkMoRefAllOf

`func NewBulkMoRefAllOf(classId string, objectType string, ) *BulkMoRefAllOf`

NewBulkMoRefAllOf instantiates a new BulkMoRefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMoRefAllOfWithDefaults

`func NewBulkMoRefAllOfWithDefaults() *BulkMoRefAllOf`

NewBulkMoRefAllOfWithDefaults instantiates a new BulkMoRefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkMoRefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkMoRefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkMoRefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkMoRefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkMoRefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkMoRefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoid

`func (o *BulkMoRefAllOf) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *BulkMoRefAllOf) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *BulkMoRefAllOf) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *BulkMoRefAllOf) HasMoid() bool`

HasMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


