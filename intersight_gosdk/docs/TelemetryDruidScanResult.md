# TelemetryDruidScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentId** | Pointer to **string** | The identifier for the row(s)&#39; segement | [optional] 
**Columns** | Pointer to **[]string** | A list of columns returned in the row(s) | [optional] 
**Events** | Pointer to **[]map[string]interface{}** | Row results | [optional] 

## Methods

### NewTelemetryDruidScanResult

`func NewTelemetryDruidScanResult() *TelemetryDruidScanResult`

NewTelemetryDruidScanResult instantiates a new TelemetryDruidScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidScanResultWithDefaults

`func NewTelemetryDruidScanResultWithDefaults() *TelemetryDruidScanResult`

NewTelemetryDruidScanResultWithDefaults instantiates a new TelemetryDruidScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentId

`func (o *TelemetryDruidScanResult) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *TelemetryDruidScanResult) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *TelemetryDruidScanResult) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *TelemetryDruidScanResult) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

### GetColumns

`func (o *TelemetryDruidScanResult) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryDruidScanResult) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryDruidScanResult) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TelemetryDruidScanResult) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEvents

`func (o *TelemetryDruidScanResult) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TelemetryDruidScanResult) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TelemetryDruidScanResult) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TelemetryDruidScanResult) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


