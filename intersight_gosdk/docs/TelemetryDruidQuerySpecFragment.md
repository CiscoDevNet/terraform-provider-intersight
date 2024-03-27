# TelemetryDruidQuerySpecFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Values** | **[]string** | A JSON array of string values to search. | 
**CaseSensitive** | Pointer to **bool** | Whether the string comparison is case-sensitive or not. | [optional] 

## Methods

### NewTelemetryDruidQuerySpecFragment

`func NewTelemetryDruidQuerySpecFragment(type_ string, values []string, ) *TelemetryDruidQuerySpecFragment`

NewTelemetryDruidQuerySpecFragment instantiates a new TelemetryDruidQuerySpecFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidQuerySpecFragmentWithDefaults

`func NewTelemetryDruidQuerySpecFragmentWithDefaults() *TelemetryDruidQuerySpecFragment`

NewTelemetryDruidQuerySpecFragmentWithDefaults instantiates a new TelemetryDruidQuerySpecFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidQuerySpecFragment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidQuerySpecFragment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidQuerySpecFragment) SetType(v string)`

SetType sets Type field to given value.


### GetValues

`func (o *TelemetryDruidQuerySpecFragment) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidQuerySpecFragment) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidQuerySpecFragment) SetValues(v []string)`

SetValues sets Values field to given value.


### GetCaseSensitive

`func (o *TelemetryDruidQuerySpecFragment) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *TelemetryDruidQuerySpecFragment) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *TelemetryDruidQuerySpecFragment) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *TelemetryDruidQuerySpecFragment) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


