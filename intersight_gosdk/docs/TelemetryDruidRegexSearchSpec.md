# TelemetryDruidRegexSearchSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | null | 
**Regex** | **string** | The regular expression to match.  If any part of a dimension value contains the pattern specified in this search query a \&quot;match\&quot; occurs. | 

## Methods

### NewTelemetryDruidRegexSearchSpec

`func NewTelemetryDruidRegexSearchSpec(type_ string, regex string, ) *TelemetryDruidRegexSearchSpec`

NewTelemetryDruidRegexSearchSpec instantiates a new TelemetryDruidRegexSearchSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidRegexSearchSpecWithDefaults

`func NewTelemetryDruidRegexSearchSpecWithDefaults() *TelemetryDruidRegexSearchSpec`

NewTelemetryDruidRegexSearchSpecWithDefaults instantiates a new TelemetryDruidRegexSearchSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidRegexSearchSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidRegexSearchSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidRegexSearchSpec) SetType(v string)`

SetType sets Type field to given value.


### GetRegex

`func (o *TelemetryDruidRegexSearchSpec) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *TelemetryDruidRegexSearchSpec) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *TelemetryDruidRegexSearchSpec) SetRegex(v string)`

SetRegex sets Regex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


