# TelemetryDruidOrFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableString** | The filter type. | 
**Fields** | [**[]TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 

## Methods

### NewTelemetryDruidOrFilter

`func NewTelemetryDruidOrFilter(type_ NullableString, fields []TelemetryDruidFilter, ) *TelemetryDruidOrFilter`

NewTelemetryDruidOrFilter instantiates a new TelemetryDruidOrFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidOrFilterWithDefaults

`func NewTelemetryDruidOrFilterWithDefaults() *TelemetryDruidOrFilter`

NewTelemetryDruidOrFilterWithDefaults instantiates a new TelemetryDruidOrFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidOrFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidOrFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidOrFilter) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TelemetryDruidOrFilter) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TelemetryDruidOrFilter) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFields

`func (o *TelemetryDruidOrFilter) GetFields() []TelemetryDruidFilter`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidOrFilter) GetFieldsOk() (*[]TelemetryDruidFilter, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidOrFilter) SetFields(v []TelemetryDruidFilter)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


