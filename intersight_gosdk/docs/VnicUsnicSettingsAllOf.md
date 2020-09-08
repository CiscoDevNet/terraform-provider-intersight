# VnicUsnicSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cos** | Pointer to **int64** | Class of Service to be used for traffic on the usNIC. | [optional] 
**Count** | Pointer to **int64** | Number of usNIC interfaces to be created. | [optional] 
**UsnicAdapterPolicy** | Pointer to **string** | Ethernet Adapter policy to be associated with the usNICs. | [optional] 

## Methods

### NewVnicUsnicSettingsAllOf

`func NewVnicUsnicSettingsAllOf() *VnicUsnicSettingsAllOf`

NewVnicUsnicSettingsAllOf instantiates a new VnicUsnicSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicUsnicSettingsAllOfWithDefaults

`func NewVnicUsnicSettingsAllOfWithDefaults() *VnicUsnicSettingsAllOf`

NewVnicUsnicSettingsAllOfWithDefaults instantiates a new VnicUsnicSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCos

`func (o *VnicUsnicSettingsAllOf) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicUsnicSettingsAllOf) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicUsnicSettingsAllOf) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicUsnicSettingsAllOf) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetCount

`func (o *VnicUsnicSettingsAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VnicUsnicSettingsAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VnicUsnicSettingsAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VnicUsnicSettingsAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetUsnicAdapterPolicy

`func (o *VnicUsnicSettingsAllOf) GetUsnicAdapterPolicy() string`

GetUsnicAdapterPolicy returns the UsnicAdapterPolicy field if non-nil, zero value otherwise.

### GetUsnicAdapterPolicyOk

`func (o *VnicUsnicSettingsAllOf) GetUsnicAdapterPolicyOk() (*string, bool)`

GetUsnicAdapterPolicyOk returns a tuple with the UsnicAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicAdapterPolicy

`func (o *VnicUsnicSettingsAllOf) SetUsnicAdapterPolicy(v string)`

SetUsnicAdapterPolicy sets UsnicAdapterPolicy field to given value.

### HasUsnicAdapterPolicy

`func (o *VnicUsnicSettingsAllOf) HasUsnicAdapterPolicy() bool`

HasUsnicAdapterPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


