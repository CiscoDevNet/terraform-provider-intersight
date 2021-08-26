# TelemetryDruidAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The aggregator type. | 
**Name** | **string** | A String for the output (result) name of the calculation. | 
**FieldName** | **string** | A String for the name of the aggregator used at ingestion time. | 
**MaxStringBytes** | Pointer to **int32** | null | [optional] [default to 1024]
**Size** | Pointer to **int32** | Must be a power of 2. Internally, size refers to the maximum number of entries sketch object will retain. Higher size means higher accuracy but more space to store sketches. Note that after you index with a particular size, druid will persist sketch in segments and you will use size greater or equal to that at query time. In general, We recommend just sticking to default size. | [optional] [default to 16384]
**Filter** | [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 
**Aggregator** | [**TelemetryDruidAggregator**](TelemetryDruidAggregator.md) |  | 

## Methods

### NewTelemetryDruidAggregator

`func NewTelemetryDruidAggregator(type_ string, name string, fieldName string, filter TelemetryDruidFilter, aggregator TelemetryDruidAggregator, ) *TelemetryDruidAggregator`

NewTelemetryDruidAggregator instantiates a new TelemetryDruidAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidAggregatorWithDefaults

`func NewTelemetryDruidAggregatorWithDefaults() *TelemetryDruidAggregator`

NewTelemetryDruidAggregatorWithDefaults instantiates a new TelemetryDruidAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidAggregator) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetMaxStringBytes

`func (o *TelemetryDruidAggregator) GetMaxStringBytes() int32`

GetMaxStringBytes returns the MaxStringBytes field if non-nil, zero value otherwise.

### GetMaxStringBytesOk

`func (o *TelemetryDruidAggregator) GetMaxStringBytesOk() (*int32, bool)`

GetMaxStringBytesOk returns a tuple with the MaxStringBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStringBytes

`func (o *TelemetryDruidAggregator) SetMaxStringBytes(v int32)`

SetMaxStringBytes sets MaxStringBytes field to given value.

### HasMaxStringBytes

`func (o *TelemetryDruidAggregator) HasMaxStringBytes() bool`

HasMaxStringBytes returns a boolean if a field has been set.

### GetSize

`func (o *TelemetryDruidAggregator) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TelemetryDruidAggregator) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TelemetryDruidAggregator) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TelemetryDruidAggregator) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetFilter

`func (o *TelemetryDruidAggregator) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidAggregator) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidAggregator) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.


### GetAggregator

`func (o *TelemetryDruidAggregator) GetAggregator() TelemetryDruidAggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *TelemetryDruidAggregator) GetAggregatorOk() (*TelemetryDruidAggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *TelemetryDruidAggregator) SetAggregator(v TelemetryDruidAggregator)`

SetAggregator sets Aggregator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


