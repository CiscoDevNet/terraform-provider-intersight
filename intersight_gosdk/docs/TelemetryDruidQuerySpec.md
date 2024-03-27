# TelemetryDruidQuerySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **string** | A String value to search. | 
**CaseSensitive** | Pointer to **bool** | Whether the string comparison is case-sensitive or not. | [optional] 
**Values** | **[]string** | A JSON array of string values to search. | 

## Methods

### NewTelemetryDruidQuerySpec

`func NewTelemetryDruidQuerySpec(type_ string, value string, values []string, ) *TelemetryDruidQuerySpec`

NewTelemetryDruidQuerySpec instantiates a new TelemetryDruidQuerySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidQuerySpecWithDefaults

`func NewTelemetryDruidQuerySpecWithDefaults() *TelemetryDruidQuerySpec`

NewTelemetryDruidQuerySpecWithDefaults instantiates a new TelemetryDruidQuerySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidQuerySpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidQuerySpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidQuerySpec) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *TelemetryDruidQuerySpec) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidQuerySpec) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidQuerySpec) SetValue(v string)`

SetValue sets Value field to given value.


### GetCaseSensitive

`func (o *TelemetryDruidQuerySpec) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *TelemetryDruidQuerySpec) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *TelemetryDruidQuerySpec) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *TelemetryDruidQuerySpec) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetValues

`func (o *TelemetryDruidQuerySpec) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidQuerySpec) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidQuerySpec) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


