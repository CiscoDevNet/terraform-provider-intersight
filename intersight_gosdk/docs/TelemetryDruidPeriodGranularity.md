# TelemetryDruidPeriodGranularity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the type of granularity. | 
**Period** | **string** | The period in ISO 8601 format. Examples are P2W, P3M, PT1H30M, PT0.750S. | 
**TimeZone** | Pointer to **string** | An optional value specifying the time zone. Standard [IANA time zones](http://joda-time.sourceforge.net/timezones.html) are supported. | [optional] 

## Methods

### NewTelemetryDruidPeriodGranularity

`func NewTelemetryDruidPeriodGranularity(type_ string, period string, ) *TelemetryDruidPeriodGranularity`

NewTelemetryDruidPeriodGranularity instantiates a new TelemetryDruidPeriodGranularity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidPeriodGranularityWithDefaults

`func NewTelemetryDruidPeriodGranularityWithDefaults() *TelemetryDruidPeriodGranularity`

NewTelemetryDruidPeriodGranularityWithDefaults instantiates a new TelemetryDruidPeriodGranularity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidPeriodGranularity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidPeriodGranularity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidPeriodGranularity) SetType(v string)`

SetType sets Type field to given value.


### GetPeriod

`func (o *TelemetryDruidPeriodGranularity) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TelemetryDruidPeriodGranularity) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TelemetryDruidPeriodGranularity) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetTimeZone

`func (o *TelemetryDruidPeriodGranularity) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TelemetryDruidPeriodGranularity) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TelemetryDruidPeriodGranularity) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TelemetryDruidPeriodGranularity) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


