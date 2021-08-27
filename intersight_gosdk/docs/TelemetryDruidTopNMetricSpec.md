# TelemetryDruidTopNMetricSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The dimension spec type. | 
**Metric** | [**TelemetryDruidTopNMetricSpec**](TelemetryDruidTopNMetricSpec.md) |  | 
**Ordering** | Pointer to **string** | Specifies the sorting order. It can be one of the following values. \&quot;lexicographic\&quot;, \&quot;alphanumeric\&quot;, \&quot;numeric\&quot;, \&quot;strlen\&quot;. * lexicographic - Sorts values by converting Strings to their UTF-8 byte array representations and comparing lexicographically, byte-by-byte. * alphanumeric - Suitable for strings with both numeric and non-numeric content, e.g. \&quot;file12 sorts after file2\&quot;. See https://github.com/amjjd/java-alphanum for more details on how this ordering sorts values. This ordering is not suitable for numbers with decimal points or negative numbers. * numeric - Sorts values as numbers, supports integers and floating point values. Negative values are supported. This sorting order will try to parse all string values as numbers. Unparseable values are treated as nulls, and nulls precede numbers. When comparing two unparseable values (e.g., \&quot;hello\&quot; and \&quot;world\&quot;), this ordering will sort by comparing the unparsed strings lexicographically. * strlen - Sorts values by the their string lengths. When there is a tie, this comparator falls back to using the String compareTo method. * version - Sorts values as versions, e.g. \&quot;10.0 sorts after 9.0\&quot;, \&quot;1.0.0-SNAPSHOT sorts after 1.0.0\&quot;. | [optional] [default to "lexicographic"]
**PreviousStop** | Pointer to **string** | The starting point of the sort. For example, if a previousStop value is &#39;b&#39;, all values before &#39;b&#39; are discarded. This field can be used to paginate through all the dimension values. | [optional] 

## Methods

### NewTelemetryDruidTopNMetricSpec

`func NewTelemetryDruidTopNMetricSpec(type_ string, metric TelemetryDruidTopNMetricSpec, ) *TelemetryDruidTopNMetricSpec`

NewTelemetryDruidTopNMetricSpec instantiates a new TelemetryDruidTopNMetricSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidTopNMetricSpecWithDefaults

`func NewTelemetryDruidTopNMetricSpecWithDefaults() *TelemetryDruidTopNMetricSpec`

NewTelemetryDruidTopNMetricSpecWithDefaults instantiates a new TelemetryDruidTopNMetricSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidTopNMetricSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidTopNMetricSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidTopNMetricSpec) SetType(v string)`

SetType sets Type field to given value.


### GetMetric

`func (o *TelemetryDruidTopNMetricSpec) GetMetric() TelemetryDruidTopNMetricSpec`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *TelemetryDruidTopNMetricSpec) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *TelemetryDruidTopNMetricSpec) SetMetric(v TelemetryDruidTopNMetricSpec)`

SetMetric sets Metric field to given value.


### GetOrdering

`func (o *TelemetryDruidTopNMetricSpec) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidTopNMetricSpec) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidTopNMetricSpec) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidTopNMetricSpec) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### GetPreviousStop

`func (o *TelemetryDruidTopNMetricSpec) GetPreviousStop() string`

GetPreviousStop returns the PreviousStop field if non-nil, zero value otherwise.

### GetPreviousStopOk

`func (o *TelemetryDruidTopNMetricSpec) GetPreviousStopOk() (*string, bool)`

GetPreviousStopOk returns a tuple with the PreviousStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStop

`func (o *TelemetryDruidTopNMetricSpec) SetPreviousStop(v string)`

SetPreviousStop sets PreviousStop field to given value.

### HasPreviousStop

`func (o *TelemetryDruidTopNMetricSpec) HasPreviousStop() bool`

HasPreviousStop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


