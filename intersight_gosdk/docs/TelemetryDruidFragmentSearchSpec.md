# TelemetryDruidFragmentSearchSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | null | 
**Values** | **[]string** | The value to match.  If any part of a dimension value contains all of the values specified in this search query spec a \&quot;match\&quot; occurs. | 
**CaseSensitive** | Pointer to **bool** | Whether or not search is case sensitive | [optional] 

## Methods

### NewTelemetryDruidFragmentSearchSpec

`func NewTelemetryDruidFragmentSearchSpec(type_ string, values []string, ) *TelemetryDruidFragmentSearchSpec`

NewTelemetryDruidFragmentSearchSpec instantiates a new TelemetryDruidFragmentSearchSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidFragmentSearchSpecWithDefaults

`func NewTelemetryDruidFragmentSearchSpecWithDefaults() *TelemetryDruidFragmentSearchSpec`

NewTelemetryDruidFragmentSearchSpecWithDefaults instantiates a new TelemetryDruidFragmentSearchSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidFragmentSearchSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidFragmentSearchSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidFragmentSearchSpec) SetType(v string)`

SetType sets Type field to given value.


### GetValues

`func (o *TelemetryDruidFragmentSearchSpec) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidFragmentSearchSpec) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidFragmentSearchSpec) SetValues(v []string)`

SetValues sets Values field to given value.


### GetCaseSensitive

`func (o *TelemetryDruidFragmentSearchSpec) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *TelemetryDruidFragmentSearchSpec) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *TelemetryDruidFragmentSearchSpec) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *TelemetryDruidFragmentSearchSpec) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


