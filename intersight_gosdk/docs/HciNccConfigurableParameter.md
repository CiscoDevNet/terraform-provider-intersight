# HciNccConfigurableParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.NccConfigurableParameter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.NccConfigurableParameter"]
**Name** | Pointer to **string** | Unique identifier name for the parameter. | [optional] [readonly] 
**ParamValue** | Pointer to [**NullableHciConfigurableParamValue**](HciConfigurableParamValue.md) |  | [optional] 
**Unit** | Pointer to **string** | Unit of measurement for the parameter. | [optional] [readonly] 

## Methods

### NewHciNccConfigurableParameter

`func NewHciNccConfigurableParameter(classId string, objectType string, ) *HciNccConfigurableParameter`

NewHciNccConfigurableParameter instantiates a new HciNccConfigurableParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciNccConfigurableParameterWithDefaults

`func NewHciNccConfigurableParameterWithDefaults() *HciNccConfigurableParameter`

NewHciNccConfigurableParameterWithDefaults instantiates a new HciNccConfigurableParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciNccConfigurableParameter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciNccConfigurableParameter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciNccConfigurableParameter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciNccConfigurableParameter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciNccConfigurableParameter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciNccConfigurableParameter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HciNccConfigurableParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciNccConfigurableParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciNccConfigurableParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciNccConfigurableParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParamValue

`func (o *HciNccConfigurableParameter) GetParamValue() HciConfigurableParamValue`

GetParamValue returns the ParamValue field if non-nil, zero value otherwise.

### GetParamValueOk

`func (o *HciNccConfigurableParameter) GetParamValueOk() (*HciConfigurableParamValue, bool)`

GetParamValueOk returns a tuple with the ParamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamValue

`func (o *HciNccConfigurableParameter) SetParamValue(v HciConfigurableParamValue)`

SetParamValue sets ParamValue field to given value.

### HasParamValue

`func (o *HciNccConfigurableParameter) HasParamValue() bool`

HasParamValue returns a boolean if a field has been set.

### SetParamValueNil

`func (o *HciNccConfigurableParameter) SetParamValueNil(b bool)`

 SetParamValueNil sets the value for ParamValue to be an explicit nil

### UnsetParamValue
`func (o *HciNccConfigurableParameter) UnsetParamValue()`

UnsetParamValue ensures that no value is present for ParamValue, not even an explicit nil
### GetUnit

`func (o *HciNccConfigurableParameter) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HciNccConfigurableParameter) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HciNccConfigurableParameter) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HciNccConfigurableParameter) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


