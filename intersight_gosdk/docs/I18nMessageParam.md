# I18nMessageParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of a variable which is referenced in a i18n text template. | [optional] [readonly] 
**Value** | Pointer to **string** | The value of a variable which is substituted in a i18n text template. | [optional] [readonly] 

## Methods

### NewI18nMessageParam

`func NewI18nMessageParam() *I18nMessageParam`

NewI18nMessageParam instantiates a new I18nMessageParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewI18nMessageParamWithDefaults

`func NewI18nMessageParamWithDefaults() *I18nMessageParam`

NewI18nMessageParamWithDefaults instantiates a new I18nMessageParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *I18nMessageParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *I18nMessageParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *I18nMessageParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *I18nMessageParam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *I18nMessageParam) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *I18nMessageParam) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *I18nMessageParam) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *I18nMessageParam) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


