# HciConfigurableParamValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ConfigurableParamValue"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ConfigurableParamValue"]
**CurrentValue** | Pointer to **string** | Current value of the parameter. | [optional] [readonly] 
**DefaultValue** | Pointer to **string** | Default value of the parameter. | [optional] [readonly] 
**ParamValueType** | Pointer to **string** | The object type of the configurable parameter value. | [optional] [readonly] 
**ValidValueRanges** | Pointer to [**[]HciConfigurableParamValueRange**](HciConfigurableParamValueRange.md) |  | [optional] 

## Methods

### NewHciConfigurableParamValue

`func NewHciConfigurableParamValue(classId string, objectType string, ) *HciConfigurableParamValue`

NewHciConfigurableParamValue instantiates a new HciConfigurableParamValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciConfigurableParamValueWithDefaults

`func NewHciConfigurableParamValueWithDefaults() *HciConfigurableParamValue`

NewHciConfigurableParamValueWithDefaults instantiates a new HciConfigurableParamValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciConfigurableParamValue) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciConfigurableParamValue) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciConfigurableParamValue) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciConfigurableParamValue) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciConfigurableParamValue) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciConfigurableParamValue) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentValue

`func (o *HciConfigurableParamValue) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *HciConfigurableParamValue) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *HciConfigurableParamValue) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *HciConfigurableParamValue) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetDefaultValue

`func (o *HciConfigurableParamValue) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *HciConfigurableParamValue) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *HciConfigurableParamValue) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *HciConfigurableParamValue) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetParamValueType

`func (o *HciConfigurableParamValue) GetParamValueType() string`

GetParamValueType returns the ParamValueType field if non-nil, zero value otherwise.

### GetParamValueTypeOk

`func (o *HciConfigurableParamValue) GetParamValueTypeOk() (*string, bool)`

GetParamValueTypeOk returns a tuple with the ParamValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamValueType

`func (o *HciConfigurableParamValue) SetParamValueType(v string)`

SetParamValueType sets ParamValueType field to given value.

### HasParamValueType

`func (o *HciConfigurableParamValue) HasParamValueType() bool`

HasParamValueType returns a boolean if a field has been set.

### GetValidValueRanges

`func (o *HciConfigurableParamValue) GetValidValueRanges() []HciConfigurableParamValueRange`

GetValidValueRanges returns the ValidValueRanges field if non-nil, zero value otherwise.

### GetValidValueRangesOk

`func (o *HciConfigurableParamValue) GetValidValueRangesOk() (*[]HciConfigurableParamValueRange, bool)`

GetValidValueRangesOk returns a tuple with the ValidValueRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidValueRanges

`func (o *HciConfigurableParamValue) SetValidValueRanges(v []HciConfigurableParamValueRange)`

SetValidValueRanges sets ValidValueRanges field to given value.

### HasValidValueRanges

`func (o *HciConfigurableParamValue) HasValidValueRanges() bool`

HasValidValueRanges returns a boolean if a field has been set.

### SetValidValueRangesNil

`func (o *HciConfigurableParamValue) SetValidValueRangesNil(b bool)`

 SetValidValueRangesNil sets the value for ValidValueRanges to be an explicit nil

### UnsetValidValueRanges
`func (o *HciConfigurableParamValue) UnsetValidValueRanges()`

UnsetValidValueRanges ensures that no value is present for ValidValueRanges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


