# HyperflexNamedVsanAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the VSAN. The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. | [optional] 
**VsanId** | Pointer to **int64** | The ID of the named VSAN. The ID can be any number between 1 and 4093, inclusive. | [optional] 

## Methods

### NewHyperflexNamedVsanAllOf

`func NewHyperflexNamedVsanAllOf() *HyperflexNamedVsanAllOf`

NewHyperflexNamedVsanAllOf instantiates a new HyperflexNamedVsanAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNamedVsanAllOfWithDefaults

`func NewHyperflexNamedVsanAllOfWithDefaults() *HyperflexNamedVsanAllOf`

NewHyperflexNamedVsanAllOfWithDefaults instantiates a new HyperflexNamedVsanAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HyperflexNamedVsanAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexNamedVsanAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexNamedVsanAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexNamedVsanAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVsanId

`func (o *HyperflexNamedVsanAllOf) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *HyperflexNamedVsanAllOf) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *HyperflexNamedVsanAllOf) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *HyperflexNamedVsanAllOf) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


