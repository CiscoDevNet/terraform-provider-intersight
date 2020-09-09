# VnicVsanSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Default VSAN ID of the virtual interface. Setting the ID to 0 will not associate any default VSAN to the traffic on the virtual interface. | [optional] 

## Methods

### NewVnicVsanSettingsAllOf

`func NewVnicVsanSettingsAllOf() *VnicVsanSettingsAllOf`

NewVnicVsanSettingsAllOf instantiates a new VnicVsanSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVsanSettingsAllOfWithDefaults

`func NewVnicVsanSettingsAllOfWithDefaults() *VnicVsanSettingsAllOf`

NewVnicVsanSettingsAllOfWithDefaults instantiates a new VnicVsanSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VnicVsanSettingsAllOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VnicVsanSettingsAllOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VnicVsanSettingsAllOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VnicVsanSettingsAllOf) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


