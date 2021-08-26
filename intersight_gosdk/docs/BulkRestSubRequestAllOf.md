# BulkRestSubRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.RestSubRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.RestSubRequest"]
**Body** | Pointer to [**MoBaseMo**](MoBaseMo.md) |  | [optional] 
**TargetMoid** | Pointer to **string** | Used with PATCH &amp; DELETE actions. The moid of an existing object instance. | [optional] 

## Methods

### NewBulkRestSubRequestAllOf

`func NewBulkRestSubRequestAllOf(classId string, objectType string, ) *BulkRestSubRequestAllOf`

NewBulkRestSubRequestAllOf instantiates a new BulkRestSubRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRestSubRequestAllOfWithDefaults

`func NewBulkRestSubRequestAllOfWithDefaults() *BulkRestSubRequestAllOf`

NewBulkRestSubRequestAllOfWithDefaults instantiates a new BulkRestSubRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkRestSubRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkRestSubRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkRestSubRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkRestSubRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkRestSubRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkRestSubRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *BulkRestSubRequestAllOf) GetBody() MoBaseMo`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BulkRestSubRequestAllOf) GetBodyOk() (*MoBaseMo, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BulkRestSubRequestAllOf) SetBody(v MoBaseMo)`

SetBody sets Body field to given value.

### HasBody

`func (o *BulkRestSubRequestAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTargetMoid

`func (o *BulkRestSubRequestAllOf) GetTargetMoid() string`

GetTargetMoid returns the TargetMoid field if non-nil, zero value otherwise.

### GetTargetMoidOk

`func (o *BulkRestSubRequestAllOf) GetTargetMoidOk() (*string, bool)`

GetTargetMoidOk returns a tuple with the TargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoid

`func (o *BulkRestSubRequestAllOf) SetTargetMoid(v string)`

SetTargetMoid sets TargetMoid field to given value.

### HasTargetMoid

`func (o *BulkRestSubRequestAllOf) HasTargetMoid() bool`

HasTargetMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


