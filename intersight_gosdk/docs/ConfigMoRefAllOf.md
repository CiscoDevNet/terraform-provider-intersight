# ConfigMoRefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "config.MoRef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "config.MoRef"]
**Moid** | Pointer to **string** | Moid represents the MoId of the object. | [optional] 

## Methods

### NewConfigMoRefAllOf

`func NewConfigMoRefAllOf(classId string, objectType string, ) *ConfigMoRefAllOf`

NewConfigMoRefAllOf instantiates a new ConfigMoRefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMoRefAllOfWithDefaults

`func NewConfigMoRefAllOfWithDefaults() *ConfigMoRefAllOf`

NewConfigMoRefAllOfWithDefaults instantiates a new ConfigMoRefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConfigMoRefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConfigMoRefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConfigMoRefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConfigMoRefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConfigMoRefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConfigMoRefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoid

`func (o *ConfigMoRefAllOf) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ConfigMoRefAllOf) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ConfigMoRefAllOf) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ConfigMoRefAllOf) HasMoid() bool`

HasMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


