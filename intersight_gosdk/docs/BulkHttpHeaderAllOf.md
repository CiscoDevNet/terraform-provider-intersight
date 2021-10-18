# BulkHttpHeaderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.HttpHeader"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.HttpHeader"]
**Name** | Pointer to **string** | The name of the http header. | [optional] [readonly] 
**Value** | Pointer to **string** | The value of the http header. | [optional] [readonly] 

## Methods

### NewBulkHttpHeaderAllOf

`func NewBulkHttpHeaderAllOf(classId string, objectType string, ) *BulkHttpHeaderAllOf`

NewBulkHttpHeaderAllOf instantiates a new BulkHttpHeaderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkHttpHeaderAllOfWithDefaults

`func NewBulkHttpHeaderAllOfWithDefaults() *BulkHttpHeaderAllOf`

NewBulkHttpHeaderAllOfWithDefaults instantiates a new BulkHttpHeaderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkHttpHeaderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkHttpHeaderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkHttpHeaderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkHttpHeaderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkHttpHeaderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkHttpHeaderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *BulkHttpHeaderAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkHttpHeaderAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkHttpHeaderAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkHttpHeaderAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *BulkHttpHeaderAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkHttpHeaderAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkHttpHeaderAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BulkHttpHeaderAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


