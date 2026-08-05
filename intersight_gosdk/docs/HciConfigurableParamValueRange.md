# HciConfigurableParamValueRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ConfigurableParamValueRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ConfigurableParamValueRange"]
**MaximumValue** | Pointer to **string** | Maximum value the parameter can have. | [optional] [readonly] 
**MinimumValue** | Pointer to **string** | Minimum value the parameter can have. | [optional] [readonly] 

## Methods

### NewHciConfigurableParamValueRange

`func NewHciConfigurableParamValueRange(classId string, objectType string, ) *HciConfigurableParamValueRange`

NewHciConfigurableParamValueRange instantiates a new HciConfigurableParamValueRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciConfigurableParamValueRangeWithDefaults

`func NewHciConfigurableParamValueRangeWithDefaults() *HciConfigurableParamValueRange`

NewHciConfigurableParamValueRangeWithDefaults instantiates a new HciConfigurableParamValueRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciConfigurableParamValueRange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciConfigurableParamValueRange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciConfigurableParamValueRange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciConfigurableParamValueRange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciConfigurableParamValueRange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciConfigurableParamValueRange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaximumValue

`func (o *HciConfigurableParamValueRange) GetMaximumValue() string`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *HciConfigurableParamValueRange) GetMaximumValueOk() (*string, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *HciConfigurableParamValueRange) SetMaximumValue(v string)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *HciConfigurableParamValueRange) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetMinimumValue

`func (o *HciConfigurableParamValueRange) GetMinimumValue() string`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *HciConfigurableParamValueRange) GetMinimumValueOk() (*string, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *HciConfigurableParamValueRange) SetMinimumValue(v string)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *HciConfigurableParamValueRange) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


