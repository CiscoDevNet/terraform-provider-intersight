# FabricServerRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fec** | Pointer to **string** | Forward error correction configuration for the port. * &#x60;Auto&#x60; - Forward error correction option &#39;Auto&#39;. * &#x60;Cl91&#x60; - Forward error correction option &#39;cl91&#39;. * &#x60;Cl74&#x60; - Forward error correction option &#39;cl74&#39;. | [optional] [default to "Auto"]

## Methods

### NewFabricServerRole

`func NewFabricServerRole() *FabricServerRole`

NewFabricServerRole instantiates a new FabricServerRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricServerRoleWithDefaults

`func NewFabricServerRoleWithDefaults() *FabricServerRole`

NewFabricServerRoleWithDefaults instantiates a new FabricServerRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFec

`func (o *FabricServerRole) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricServerRole) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricServerRole) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricServerRole) HasFec() bool`

HasFec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


