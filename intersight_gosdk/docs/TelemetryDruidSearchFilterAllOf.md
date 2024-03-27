# TelemetryDruidSearchFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Query** | [**TelemetryDruidQuerySpec**](TelemetryDruidQuerySpec.md) |  | 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidSearchFilterAllOf

`func NewTelemetryDruidSearchFilterAllOf(type_ string, dimension string, query TelemetryDruidQuerySpec, ) *TelemetryDruidSearchFilterAllOf`

NewTelemetryDruidSearchFilterAllOf instantiates a new TelemetryDruidSearchFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSearchFilterAllOfWithDefaults

`func NewTelemetryDruidSearchFilterAllOfWithDefaults() *TelemetryDruidSearchFilterAllOf`

NewTelemetryDruidSearchFilterAllOfWithDefaults instantiates a new TelemetryDruidSearchFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidSearchFilterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidSearchFilterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidSearchFilterAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidSearchFilterAllOf) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidSearchFilterAllOf) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidSearchFilterAllOf) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetQuery

`func (o *TelemetryDruidSearchFilterAllOf) GetQuery() TelemetryDruidQuerySpec`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TelemetryDruidSearchFilterAllOf) GetQueryOk() (*TelemetryDruidQuerySpec, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TelemetryDruidSearchFilterAllOf) SetQuery(v TelemetryDruidQuerySpec)`

SetQuery sets Query field to given value.


### GetExtractionFn

`func (o *TelemetryDruidSearchFilterAllOf) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidSearchFilterAllOf) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidSearchFilterAllOf) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidSearchFilterAllOf) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


