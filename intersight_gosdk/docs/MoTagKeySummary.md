# MoTagKeySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The tag key for which usage information is provided. | [optional] 
**NumKeys** | Pointer to **int32** | The number of times this tag Key has been set in an API resource. | [optional] 
**Values** | Pointer to **[]string** | A list of all Tag values that have been assigned to this tag Key. | [optional] 

## Methods

### NewMoTagKeySummary

`func NewMoTagKeySummary() *MoTagKeySummary`

NewMoTagKeySummary instantiates a new MoTagKeySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoTagKeySummaryWithDefaults

`func NewMoTagKeySummaryWithDefaults() *MoTagKeySummary`

NewMoTagKeySummaryWithDefaults instantiates a new MoTagKeySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MoTagKeySummary) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MoTagKeySummary) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MoTagKeySummary) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MoTagKeySummary) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNumKeys

`func (o *MoTagKeySummary) GetNumKeys() int32`

GetNumKeys returns the NumKeys field if non-nil, zero value otherwise.

### GetNumKeysOk

`func (o *MoTagKeySummary) GetNumKeysOk() (*int32, bool)`

GetNumKeysOk returns a tuple with the NumKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumKeys

`func (o *MoTagKeySummary) SetNumKeys(v int32)`

SetNumKeys sets NumKeys field to given value.

### HasNumKeys

`func (o *MoTagKeySummary) HasNumKeys() bool`

HasNumKeys returns a boolean if a field has been set.

### GetValues

`func (o *MoTagKeySummary) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MoTagKeySummary) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MoTagKeySummary) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *MoTagKeySummary) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


