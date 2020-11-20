# TelemetryDruidAggregateSearchSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | null | 
**Value** | **string** | The value to match.  If any part of a dimension value contains the value specified in this search query spec, regardless of case, a \&quot;match\&quot; occurs. | 
**Values** | **[]string** | The value to match.  If any part of a dimension value contains all of the values specified in this search query spec a \&quot;match\&quot; occurs. | 
**CaseSensitive** | Pointer to **bool** | Whether or not search is case sensitive | [optional] 
**Regex** | **string** | The regular expression to match.  If any part of a dimension value contains the pattern specified in this search query a \&quot;match\&quot; occurs. | 

## Methods

### NewTelemetryDruidAggregateSearchSpec

`func NewTelemetryDruidAggregateSearchSpec(type_ string, value string, values []string, regex string, ) *TelemetryDruidAggregateSearchSpec`

NewTelemetryDruidAggregateSearchSpec instantiates a new TelemetryDruidAggregateSearchSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidAggregateSearchSpecWithDefaults

`func NewTelemetryDruidAggregateSearchSpecWithDefaults() *TelemetryDruidAggregateSearchSpec`

NewTelemetryDruidAggregateSearchSpecWithDefaults instantiates a new TelemetryDruidAggregateSearchSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidAggregateSearchSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidAggregateSearchSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidAggregateSearchSpec) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *TelemetryDruidAggregateSearchSpec) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidAggregateSearchSpec) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidAggregateSearchSpec) SetValue(v string)`

SetValue sets Value field to given value.


### GetValues

`func (o *TelemetryDruidAggregateSearchSpec) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidAggregateSearchSpec) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidAggregateSearchSpec) SetValues(v []string)`

SetValues sets Values field to given value.


### GetCaseSensitive

`func (o *TelemetryDruidAggregateSearchSpec) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *TelemetryDruidAggregateSearchSpec) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *TelemetryDruidAggregateSearchSpec) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *TelemetryDruidAggregateSearchSpec) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetRegex

`func (o *TelemetryDruidAggregateSearchSpec) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *TelemetryDruidAggregateSearchSpec) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *TelemetryDruidAggregateSearchSpec) SetRegex(v string)`

SetRegex sets Regex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


