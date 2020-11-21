# I18nMessageParamAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "i18n.MessageParam"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "i18n.MessageParam"]
**Name** | Pointer to **string** | The name of a variable which is referenced in a i18n text template. | [optional] [readonly] 
**Value** | Pointer to **string** | The value of a variable which is substituted in a i18n text template. | [optional] [readonly] 

## Methods

### NewI18nMessageParamAllOf

`func NewI18nMessageParamAllOf(classId string, objectType string, ) *I18nMessageParamAllOf`

NewI18nMessageParamAllOf instantiates a new I18nMessageParamAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewI18nMessageParamAllOfWithDefaults

`func NewI18nMessageParamAllOfWithDefaults() *I18nMessageParamAllOf`

NewI18nMessageParamAllOfWithDefaults instantiates a new I18nMessageParamAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *I18nMessageParamAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *I18nMessageParamAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *I18nMessageParamAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *I18nMessageParamAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *I18nMessageParamAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *I18nMessageParamAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *I18nMessageParamAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *I18nMessageParamAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *I18nMessageParamAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *I18nMessageParamAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *I18nMessageParamAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *I18nMessageParamAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *I18nMessageParamAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *I18nMessageParamAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


