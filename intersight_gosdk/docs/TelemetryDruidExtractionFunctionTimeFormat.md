# TelemetryDruidExtractionFunctionTimeFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Format** | Pointer to **string** | Date time format for the resulting dimension value, in Joda Time DateTimeFormat, or null to use the default ISO8601 format. | [optional] 
**Locale** | Pointer to **string** | Locale (language and country) to use, given as a IETF BCP 47 language tag, e.g. en-US, en-GB, fr-FR, fr-CA, etc. | [optional] 
**TimeZone** | Pointer to **string** | Time zone to use in IANA tz database format, e.g. Europe/Berlin (this can possibly be different than the aggregation time-zone) | [optional] 
**Granularity** | Pointer to [**TelemetryDruidGranularity**](TelemetryDruidGranularity.md) |  | [optional] 
**AsMillis** | Pointer to **bool** | boolean value, set to true to treat input strings as millis rather than ISO8601 strings. Additionally, if format is null or not specified, output will be in millis rather than ISO8601. | [optional] 

## Methods

### NewTelemetryDruidExtractionFunctionTimeFormat

`func NewTelemetryDruidExtractionFunctionTimeFormat(type_ string, ) *TelemetryDruidExtractionFunctionTimeFormat`

NewTelemetryDruidExtractionFunctionTimeFormat instantiates a new TelemetryDruidExtractionFunctionTimeFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionTimeFormatWithDefaults

`func NewTelemetryDruidExtractionFunctionTimeFormatWithDefaults() *TelemetryDruidExtractionFunctionTimeFormat`

NewTelemetryDruidExtractionFunctionTimeFormatWithDefaults instantiates a new TelemetryDruidExtractionFunctionTimeFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunctionTimeFormat) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TelemetryDruidExtractionFunctionTimeFormat) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TelemetryDruidExtractionFunctionTimeFormat) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLocale

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TelemetryDruidExtractionFunctionTimeFormat) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TelemetryDruidExtractionFunctionTimeFormat) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetTimeZone

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TelemetryDruidExtractionFunctionTimeFormat) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TelemetryDruidExtractionFunctionTimeFormat) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetGranularity

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetGranularity() TelemetryDruidGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetGranularityOk() (*TelemetryDruidGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TelemetryDruidExtractionFunctionTimeFormat) SetGranularity(v TelemetryDruidGranularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *TelemetryDruidExtractionFunctionTimeFormat) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetAsMillis

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetAsMillis() bool`

GetAsMillis returns the AsMillis field if non-nil, zero value otherwise.

### GetAsMillisOk

`func (o *TelemetryDruidExtractionFunctionTimeFormat) GetAsMillisOk() (*bool, bool)`

GetAsMillisOk returns a tuple with the AsMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsMillis

`func (o *TelemetryDruidExtractionFunctionTimeFormat) SetAsMillis(v bool)`

SetAsMillis sets AsMillis field to given value.

### HasAsMillis

`func (o *TelemetryDruidExtractionFunctionTimeFormat) HasAsMillis() bool`

HasAsMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


