# AaaRetentionConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "aaa.RetentionConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "aaa.RetentionConfig"]
**RetentionPeriod** | Pointer to **int64** | The default retention period in months for audit log retention for accounts without a retention policy. | [optional] 

## Methods

### NewAaaRetentionConfigAllOf

`func NewAaaRetentionConfigAllOf(classId string, objectType string, ) *AaaRetentionConfigAllOf`

NewAaaRetentionConfigAllOf instantiates a new AaaRetentionConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaaRetentionConfigAllOfWithDefaults

`func NewAaaRetentionConfigAllOfWithDefaults() *AaaRetentionConfigAllOf`

NewAaaRetentionConfigAllOfWithDefaults instantiates a new AaaRetentionConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AaaRetentionConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AaaRetentionConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AaaRetentionConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AaaRetentionConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AaaRetentionConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AaaRetentionConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRetentionPeriod

`func (o *AaaRetentionConfigAllOf) GetRetentionPeriod() int64`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *AaaRetentionConfigAllOf) GetRetentionPeriodOk() (*int64, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *AaaRetentionConfigAllOf) SetRetentionPeriod(v int64)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *AaaRetentionConfigAllOf) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


