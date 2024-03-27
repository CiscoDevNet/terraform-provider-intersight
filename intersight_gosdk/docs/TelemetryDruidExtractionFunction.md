# TelemetryDruidExtractionFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Expr** | **string** |  | 
**Index** | **int32** |  | 
**ReplaceMissingValue** | Pointer to **bool** | If the replaceMissingValue property is true, the extraction function will transform dimension values that do not match the regex pattern to a user-specified String. Default value is false. | [optional] 
**ReplaceMissingValueWith** | Pointer to **string** | Provides a hint how to handle missing values. Setting replaceMissingValueWith to \&quot;\&quot; has the same effect as setting it to null or omitting the property. | [optional] 
**Query** | **string** |  | 
**Length** | Pointer to **int32** | The length may be omitted for substring to return the remainder of the dimension value starting from index, or null if index greater than the length of the dimension value. | [optional] 
**Format** | **string** | A sprintf expression. For example if you want to concat \&quot;[\&quot; and \&quot;]\&quot; before and after the actual dimension value, you need to specify \&quot;[%s]\&quot; as format string. | 
**Locale** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** | Time zone to use in IANA tz database format, e.g. Europe/Berlin (this can possibly be different than the aggregation time-zone) | [optional] 
**Granularity** | Pointer to [**TelemetryDruidGranularity**](TelemetryDruidGranularity.md) |  | [optional] 
**AsMillis** | Pointer to **bool** | boolean value, set to true to treat input strings as millis rather than ISO8601 strings. Additionally, if format is null or not specified, output will be in millis rather than ISO8601. | [optional] 
**TimeFormat** | Pointer to **string** |  | [optional] 
**ResultFormat** | Pointer to **string** |  | [optional] 
**Joda** | Pointer to **bool** |  | [optional] 
**Lookup** | Pointer to [**TelemetryDruidExtractionFunctionInlineLookupAllOfLookup**](TelemetryDruidExtractionFunctionInlineLookupAllOfLookup.md) |  | [optional] 
**RetainMissingValue** | Pointer to **bool** | Provides a hint how to handle missing values. Setting retainMissingValue to true will use the dimension&#39;s original value if it is not found in the lookup. The default values are replaceMissingValueWith &#x3D; null and retainMissingValue &#x3D; false which causes missing values to be treated as missing. It is illegal to set retainMissingValue &#x3D; true and also specify a replaceMissingValueWith. | [optional] 
**Injective** | Pointer to **bool** | Override the lookup&#39;s own sense of whether or not it is injective. | [optional] 
**Optimize** | Pointer to **bool** | Allow optimization of lookup based extraction filter (by default optimize &#x3D; true). The optimization layer will run on the Broker and it will rewrite the extraction filter as clause of selector filters. | [optional] 
**ExtractionFns** | Pointer to [**[]TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 
**NullHandling** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** | The size of the buckets (optional, default 1). | [optional] 
**Offset** | Pointer to **int32** | The offset for the buckets (optional, default 0). | [optional] 

## Methods

### NewTelemetryDruidExtractionFunction

`func NewTelemetryDruidExtractionFunction(type_ string, expr string, index int32, query string, format string, ) *TelemetryDruidExtractionFunction`

NewTelemetryDruidExtractionFunction instantiates a new TelemetryDruidExtractionFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionWithDefaults

`func NewTelemetryDruidExtractionFunctionWithDefaults() *TelemetryDruidExtractionFunction`

NewTelemetryDruidExtractionFunctionWithDefaults instantiates a new TelemetryDruidExtractionFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunction) SetType(v string)`

SetType sets Type field to given value.


### GetExpr

`func (o *TelemetryDruidExtractionFunction) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *TelemetryDruidExtractionFunction) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *TelemetryDruidExtractionFunction) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetIndex

`func (o *TelemetryDruidExtractionFunction) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TelemetryDruidExtractionFunction) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TelemetryDruidExtractionFunction) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetReplaceMissingValue

`func (o *TelemetryDruidExtractionFunction) GetReplaceMissingValue() bool`

GetReplaceMissingValue returns the ReplaceMissingValue field if non-nil, zero value otherwise.

### GetReplaceMissingValueOk

`func (o *TelemetryDruidExtractionFunction) GetReplaceMissingValueOk() (*bool, bool)`

GetReplaceMissingValueOk returns a tuple with the ReplaceMissingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingValue

`func (o *TelemetryDruidExtractionFunction) SetReplaceMissingValue(v bool)`

SetReplaceMissingValue sets ReplaceMissingValue field to given value.

### HasReplaceMissingValue

`func (o *TelemetryDruidExtractionFunction) HasReplaceMissingValue() bool`

HasReplaceMissingValue returns a boolean if a field has been set.

### GetReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunction) GetReplaceMissingValueWith() string`

GetReplaceMissingValueWith returns the ReplaceMissingValueWith field if non-nil, zero value otherwise.

### GetReplaceMissingValueWithOk

`func (o *TelemetryDruidExtractionFunction) GetReplaceMissingValueWithOk() (*string, bool)`

GetReplaceMissingValueWithOk returns a tuple with the ReplaceMissingValueWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunction) SetReplaceMissingValueWith(v string)`

SetReplaceMissingValueWith sets ReplaceMissingValueWith field to given value.

### HasReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunction) HasReplaceMissingValueWith() bool`

HasReplaceMissingValueWith returns a boolean if a field has been set.

### GetQuery

`func (o *TelemetryDruidExtractionFunction) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TelemetryDruidExtractionFunction) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TelemetryDruidExtractionFunction) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetLength

`func (o *TelemetryDruidExtractionFunction) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *TelemetryDruidExtractionFunction) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *TelemetryDruidExtractionFunction) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *TelemetryDruidExtractionFunction) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetFormat

`func (o *TelemetryDruidExtractionFunction) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TelemetryDruidExtractionFunction) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TelemetryDruidExtractionFunction) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetLocale

`func (o *TelemetryDruidExtractionFunction) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TelemetryDruidExtractionFunction) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TelemetryDruidExtractionFunction) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TelemetryDruidExtractionFunction) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetTimeZone

`func (o *TelemetryDruidExtractionFunction) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TelemetryDruidExtractionFunction) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TelemetryDruidExtractionFunction) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TelemetryDruidExtractionFunction) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetGranularity

`func (o *TelemetryDruidExtractionFunction) GetGranularity() TelemetryDruidGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TelemetryDruidExtractionFunction) GetGranularityOk() (*TelemetryDruidGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TelemetryDruidExtractionFunction) SetGranularity(v TelemetryDruidGranularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *TelemetryDruidExtractionFunction) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetAsMillis

`func (o *TelemetryDruidExtractionFunction) GetAsMillis() bool`

GetAsMillis returns the AsMillis field if non-nil, zero value otherwise.

### GetAsMillisOk

`func (o *TelemetryDruidExtractionFunction) GetAsMillisOk() (*bool, bool)`

GetAsMillisOk returns a tuple with the AsMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsMillis

`func (o *TelemetryDruidExtractionFunction) SetAsMillis(v bool)`

SetAsMillis sets AsMillis field to given value.

### HasAsMillis

`func (o *TelemetryDruidExtractionFunction) HasAsMillis() bool`

HasAsMillis returns a boolean if a field has been set.

### GetTimeFormat

`func (o *TelemetryDruidExtractionFunction) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *TelemetryDruidExtractionFunction) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *TelemetryDruidExtractionFunction) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *TelemetryDruidExtractionFunction) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### GetResultFormat

`func (o *TelemetryDruidExtractionFunction) GetResultFormat() string`

GetResultFormat returns the ResultFormat field if non-nil, zero value otherwise.

### GetResultFormatOk

`func (o *TelemetryDruidExtractionFunction) GetResultFormatOk() (*string, bool)`

GetResultFormatOk returns a tuple with the ResultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFormat

`func (o *TelemetryDruidExtractionFunction) SetResultFormat(v string)`

SetResultFormat sets ResultFormat field to given value.

### HasResultFormat

`func (o *TelemetryDruidExtractionFunction) HasResultFormat() bool`

HasResultFormat returns a boolean if a field has been set.

### GetJoda

`func (o *TelemetryDruidExtractionFunction) GetJoda() bool`

GetJoda returns the Joda field if non-nil, zero value otherwise.

### GetJodaOk

`func (o *TelemetryDruidExtractionFunction) GetJodaOk() (*bool, bool)`

GetJodaOk returns a tuple with the Joda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoda

`func (o *TelemetryDruidExtractionFunction) SetJoda(v bool)`

SetJoda sets Joda field to given value.

### HasJoda

`func (o *TelemetryDruidExtractionFunction) HasJoda() bool`

HasJoda returns a boolean if a field has been set.

### GetLookup

`func (o *TelemetryDruidExtractionFunction) GetLookup() TelemetryDruidExtractionFunctionInlineLookupAllOfLookup`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *TelemetryDruidExtractionFunction) GetLookupOk() (*TelemetryDruidExtractionFunctionInlineLookupAllOfLookup, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *TelemetryDruidExtractionFunction) SetLookup(v TelemetryDruidExtractionFunctionInlineLookupAllOfLookup)`

SetLookup sets Lookup field to given value.

### HasLookup

`func (o *TelemetryDruidExtractionFunction) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### GetRetainMissingValue

`func (o *TelemetryDruidExtractionFunction) GetRetainMissingValue() bool`

GetRetainMissingValue returns the RetainMissingValue field if non-nil, zero value otherwise.

### GetRetainMissingValueOk

`func (o *TelemetryDruidExtractionFunction) GetRetainMissingValueOk() (*bool, bool)`

GetRetainMissingValueOk returns a tuple with the RetainMissingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainMissingValue

`func (o *TelemetryDruidExtractionFunction) SetRetainMissingValue(v bool)`

SetRetainMissingValue sets RetainMissingValue field to given value.

### HasRetainMissingValue

`func (o *TelemetryDruidExtractionFunction) HasRetainMissingValue() bool`

HasRetainMissingValue returns a boolean if a field has been set.

### GetInjective

`func (o *TelemetryDruidExtractionFunction) GetInjective() bool`

GetInjective returns the Injective field if non-nil, zero value otherwise.

### GetInjectiveOk

`func (o *TelemetryDruidExtractionFunction) GetInjectiveOk() (*bool, bool)`

GetInjectiveOk returns a tuple with the Injective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjective

`func (o *TelemetryDruidExtractionFunction) SetInjective(v bool)`

SetInjective sets Injective field to given value.

### HasInjective

`func (o *TelemetryDruidExtractionFunction) HasInjective() bool`

HasInjective returns a boolean if a field has been set.

### GetOptimize

`func (o *TelemetryDruidExtractionFunction) GetOptimize() bool`

GetOptimize returns the Optimize field if non-nil, zero value otherwise.

### GetOptimizeOk

`func (o *TelemetryDruidExtractionFunction) GetOptimizeOk() (*bool, bool)`

GetOptimizeOk returns a tuple with the Optimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimize

`func (o *TelemetryDruidExtractionFunction) SetOptimize(v bool)`

SetOptimize sets Optimize field to given value.

### HasOptimize

`func (o *TelemetryDruidExtractionFunction) HasOptimize() bool`

HasOptimize returns a boolean if a field has been set.

### GetExtractionFns

`func (o *TelemetryDruidExtractionFunction) GetExtractionFns() []TelemetryDruidExtractionFunction`

GetExtractionFns returns the ExtractionFns field if non-nil, zero value otherwise.

### GetExtractionFnsOk

`func (o *TelemetryDruidExtractionFunction) GetExtractionFnsOk() (*[]TelemetryDruidExtractionFunction, bool)`

GetExtractionFnsOk returns a tuple with the ExtractionFns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFns

`func (o *TelemetryDruidExtractionFunction) SetExtractionFns(v []TelemetryDruidExtractionFunction)`

SetExtractionFns sets ExtractionFns field to given value.

### HasExtractionFns

`func (o *TelemetryDruidExtractionFunction) HasExtractionFns() bool`

HasExtractionFns returns a boolean if a field has been set.

### GetNullHandling

`func (o *TelemetryDruidExtractionFunction) GetNullHandling() string`

GetNullHandling returns the NullHandling field if non-nil, zero value otherwise.

### GetNullHandlingOk

`func (o *TelemetryDruidExtractionFunction) GetNullHandlingOk() (*string, bool)`

GetNullHandlingOk returns a tuple with the NullHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullHandling

`func (o *TelemetryDruidExtractionFunction) SetNullHandling(v string)`

SetNullHandling sets NullHandling field to given value.

### HasNullHandling

`func (o *TelemetryDruidExtractionFunction) HasNullHandling() bool`

HasNullHandling returns a boolean if a field has been set.

### GetSize

`func (o *TelemetryDruidExtractionFunction) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TelemetryDruidExtractionFunction) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TelemetryDruidExtractionFunction) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TelemetryDruidExtractionFunction) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetOffset

`func (o *TelemetryDruidExtractionFunction) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TelemetryDruidExtractionFunction) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TelemetryDruidExtractionFunction) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TelemetryDruidExtractionFunction) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


