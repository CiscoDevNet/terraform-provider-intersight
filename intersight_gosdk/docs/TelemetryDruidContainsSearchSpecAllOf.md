# TelemetryDruidContainsSearchSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The value to match.  If any part of a dimension value contains the value specified in this search query spec, regardless of case, a \&quot;match\&quot; occurs. | 
**CaseSensitive** | Pointer to **bool** | Whether or not search is case sensitive | [optional] 

## Methods

### NewTelemetryDruidContainsSearchSpecAllOf

`func NewTelemetryDruidContainsSearchSpecAllOf(value string, ) *TelemetryDruidContainsSearchSpecAllOf`

NewTelemetryDruidContainsSearchSpecAllOf instantiates a new TelemetryDruidContainsSearchSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidContainsSearchSpecAllOfWithDefaults

`func NewTelemetryDruidContainsSearchSpecAllOfWithDefaults() *TelemetryDruidContainsSearchSpecAllOf`

NewTelemetryDruidContainsSearchSpecAllOfWithDefaults instantiates a new TelemetryDruidContainsSearchSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TelemetryDruidContainsSearchSpecAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidContainsSearchSpecAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidContainsSearchSpecAllOf) SetValue(v string)`

SetValue sets Value field to given value.


### GetCaseSensitive

`func (o *TelemetryDruidContainsSearchSpecAllOf) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *TelemetryDruidContainsSearchSpecAllOf) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *TelemetryDruidContainsSearchSpecAllOf) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *TelemetryDruidContainsSearchSpecAllOf) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


