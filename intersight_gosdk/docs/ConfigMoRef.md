# ConfigMoRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "config.MoRef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "config.MoRef"]
**Moid** | Pointer to **string** | Moid represents the MoId of the object. | [optional] 

## Methods

### NewConfigMoRef

`func NewConfigMoRef(classId string, objectType string, ) *ConfigMoRef`

NewConfigMoRef instantiates a new ConfigMoRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMoRefWithDefaults

`func NewConfigMoRefWithDefaults() *ConfigMoRef`

NewConfigMoRefWithDefaults instantiates a new ConfigMoRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConfigMoRef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConfigMoRef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConfigMoRef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConfigMoRef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConfigMoRef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConfigMoRef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoid

`func (o *ConfigMoRef) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ConfigMoRef) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ConfigMoRef) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ConfigMoRef) HasMoid() bool`

HasMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


