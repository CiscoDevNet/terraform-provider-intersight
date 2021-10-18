# BulkHttpHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.HttpHeader"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.HttpHeader"]
**Name** | Pointer to **string** | The name of the http header. | [optional] [readonly] 
**Value** | Pointer to **string** | The value of the http header. | [optional] [readonly] 

## Methods

### NewBulkHttpHeader

`func NewBulkHttpHeader(classId string, objectType string, ) *BulkHttpHeader`

NewBulkHttpHeader instantiates a new BulkHttpHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkHttpHeaderWithDefaults

`func NewBulkHttpHeaderWithDefaults() *BulkHttpHeader`

NewBulkHttpHeaderWithDefaults instantiates a new BulkHttpHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkHttpHeader) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkHttpHeader) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkHttpHeader) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkHttpHeader) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkHttpHeader) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkHttpHeader) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *BulkHttpHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkHttpHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkHttpHeader) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkHttpHeader) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *BulkHttpHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkHttpHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkHttpHeader) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BulkHttpHeader) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


